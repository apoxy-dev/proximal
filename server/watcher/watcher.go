package watcher

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"

	"github.com/apoxy-dev/proximal/core/log"

	middlewarev1 "github.com/apoxy-dev/proximal/api/middleware/v1"
)

// Watcher watches a directory for changes and triggers a build when a change is
// detected.
type Watcher struct {
	// IgnoreRegex is a regex of files to ignore.
	IgnoreRegex string

	client middlewarev1.MiddlewareServiceClient
	wr     *fsnotify.Watcher

	mu sync.Mutex
	// watches is a map of slugs to watch directories.
	watches map[string]string
	// subpaths is a map of subpaths to slugs.
	subpaths map[string]string
	// batches is a map of slugs to whether a build has been triggered.
	batches map[string]bool
}

// NewWatcher creates a new Watcher.
func NewWatcher(c middlewarev1.MiddlewareServiceClient, ignoreRegex string) *Watcher {
	return &Watcher{
		IgnoreRegex: ignoreRegex,
		client:      c,
		watches:     make(map[string]string),
		subpaths:    make(map[string]string),
		batches:     make(map[string]bool),
	}
}

// Run starts the watcher until the context is cancelled.
func (w *Watcher) Run(ctx context.Context) error {
	var err error
	w.wr, err = fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer w.wr.Close()

	r, err := regexp.Compile(w.IgnoreRegex)
	if err != nil {
		return fmt.Errorf("failed to compile ignore regex: %w", err)
	}

	errCh := make(chan error, 1)
	go func() {
		var err error
		defer func() {
			errCh <- err
		}()
		for {
			select {
			case e, ok := <-w.wr.Events:
				if !ok {
					err = errors.New("watcher closed")
					return
				}
				log.Infof("change detected in %s: %s", e.Name, e.Op)

				if e.Op&(fsnotify.Create|fsnotify.Write|fsnotify.Remove|fsnotify.Rename) != 0 {
					if r.MatchString(filepath.Base(e.Name)) {
						log.Debugf("ignoring %s", e.Name)
						continue
					}

					slug, ok := w.subpaths[filepath.Dir(e.Name)]
					if !ok {
						log.Errorf("no slug found for %s", e.Name)
						continue
					}

					log.Infof("batching re-build for slug %s due to %s", slug, e.Name)

					w.mu.Lock()
					w.batches[slug] = true
					w.mu.Unlock()
				}
			case err, ok := <-w.wr.Errors:
				if !ok {
					err = errors.New("watcher closed")
					return
				}
				log.Errorf("watcher error: %v", err)
			case <-ctx.Done():
				err = ctx.Err()
				return
			}
		}
	}()

	go func() {
		for {
			select {
			case <-time.After(2 * time.Second):
				w.mu.Lock()
				for slug := range w.batches {
					log.Infof("triggering build for %s", slug)
					if _, err := w.client.TriggerBuild(ctx, &middlewarev1.TriggerBuildRequest{
						Slug: slug,
					}); err != nil {
						log.Errorf("failed to trigger build: %v", err)
					}
				}
				w.batches = make(map[string]bool)
				w.mu.Unlock()
			case <-ctx.Done():
				return
			}
		}
	}()

	return <-errCh
}

// Add adds a new watch.
func (w *Watcher) Add(slug, dir string) error {
	w.mu.Lock()
	defer w.mu.Unlock()

	dir, err := filepath.Abs(dir)
	if err != nil {
		return err
	}

	if d, ok := w.watches[slug]; ok && d != dir {
		return fmt.Errorf("slug %s already exists with dir %s", slug, d)
	}

	w.watches[slug] = dir

	log.Infof("watching %s", dir)

	// Walk dir tree and place watchs for each dir (watches are not recursive).
	if err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() || info.Name() == ".git" {
			return nil
		}
		log.Debugf("watching %s", path)
		if err := w.wr.Add(path); err != nil {
			return err
		}
		w.subpaths[path] = slug
		return nil
	}); err != nil {
		return err
	}

	return nil
}

// Remove removes a watch.
func (w *Watcher) Remove(slug string) error {
	w.mu.Lock()
	defer w.mu.Unlock()

	dir, ok := w.watches[slug]
	if !ok {
		return fmt.Errorf("slug %s does not exist", slug)
	}

	// Removes all paths beginning with dir.
	for _, path := range w.wr.WatchList() {
		if !filepath.HasPrefix(path, dir) {
			continue
		}

		if err := w.wr.Remove(path); err != nil {
			return err
		}

		delete(w.subpaths, path)
	}

	delete(w.watches, slug)

	return nil
}
