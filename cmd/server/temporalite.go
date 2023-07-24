package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/temporalio/temporalite"
	uiserver "github.com/temporalio/ui-server/v2/server"
	uiconfig "github.com/temporalio/ui-server/v2/server/config"
	uiserveroptions "github.com/temporalio/ui-server/v2/server/server_options"
	temporalconfig "go.temporal.io/server/common/config"
	"go.temporal.io/server/common/log/tag"
	"go.temporal.io/server/common/primitives"
	"go.temporal.io/server/temporal"
	"golang.org/x/exp/slog"
)

var (
	temporalitePort   = flag.Int("temporalite_port", 8088, "Temporalite port (local mode only).")
	temporaliteUIPort = flag.Int("temporalite_ui_port", 9088, "Temporalite UI port (local mode only).")
	temporaliteEphem  = flag.Bool(
		"temporalite_ephemeral",
		false,
		"Use an ephemeral database for Temporalite (local mode only).",
	)
	temporaliteDB = flag.String("temporalite_db", "temporalite.db", "Temporalite database file path.")
)

func createTemporaliteServer() (*temporalite.Server, error) {
	opts := []temporalite.ServerOption{
		temporalite.WithDynamicPorts(),
		temporalite.WithNamespaces(*temporalNs),
		temporalite.WithFrontendPort(*temporalitePort),
		temporalite.WithFrontendIP("0.0.0.0"),
		temporalite.WithUpstreamOptions(
			temporal.WithLogger(temporaliteLogger{}),
			temporal.ForServices([]string{
				string(primitives.FrontendService),
				string(primitives.HistoryService),
				string(primitives.MatchingService),
				string(primitives.WorkerService),
			}),
		),
		temporalite.WithBaseConfig(&temporalconfig.Config{}),
		temporalite.WithUI(uiserver.NewServer(uiserveroptions.WithConfigProvider(&uiconfig.Config{
			Host:                "0.0.0.0",
			Port:                *temporaliteUIPort,
			TemporalGRPCAddress: fmt.Sprintf("0.0.0.0:%d", *temporalitePort),
			EnableUI:            true,
		}))),
	}
	if *temporaliteEphem {
		opts = append(opts, temporalite.WithPersistenceDisabled())
	} else {
		opts = append(opts, temporalite.WithDatabaseFilePath(*temporaliteDB))
	}
	return temporalite.NewServer(opts...)
}

type temporaliteLogger struct{}

func (l temporaliteLogger) kv(msg string, tags []tag.Tag) []interface{} {
	kvs := make([]interface{}, len(tags))
	for i, tag := range tags {
		kvs[i] = slog.Any(tag.Key(), tag.Value())
	}
	return kvs
}

func (l temporaliteLogger) Debug(msg string, tags ...tag.Tag) {
	slog.Debug(msg, l.kv(msg, tags)...)
}

func (l temporaliteLogger) Info(msg string, tags ...tag.Tag) {
	slog.Info(msg, l.kv(msg, tags)...)
}

func (l temporaliteLogger) Warn(msg string, tags ...tag.Tag) {
	slog.Warn(msg, l.kv(msg, tags)...)
}

func (l temporaliteLogger) Error(msg string, tags ...tag.Tag) {
	slog.Error(msg, l.kv(msg, tags)...)
}

func (l temporaliteLogger) DPanic(msg string, tags ...tag.Tag) {
	panic(fmt.Sprintf(msg, l.kv(msg, tags)...))
}

func (l temporaliteLogger) Panic(msg string, tags ...tag.Tag) {
	panic(fmt.Sprintf(msg, l.kv(msg, tags)...))
}

func (l temporaliteLogger) Fatal(msg string, tags ...tag.Tag) {
	slog.Error(msg, l.kv(msg, tags)...)
	os.Exit(1)
}
