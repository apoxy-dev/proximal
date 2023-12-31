# Developer Documentation

## Building Proximal

Since we rely on considerable amount of codegen (Protobufs, gRPC, gRPC-GW, sqlc) this project is
using [Bazel](https://bazel.build) for its build automation. Bazel version is pinned using
`.bazeliskrc` in the root so we recommend using [Bazelisk](https://github.com/bazelbuild/bazelisk)
for the least amount of surprise.

To build the Proximal server process:
```shell
bazel build --config=`arch` //cmd/server
```

This above command will use [zig cc toolchain](https://sr.ht/~motiejus/bazel-zig-cc/) to compile and
link Linux binary for amd64/aarch64 archs. That means cross-compilation when running from Mac hosts.

Not to run a locally-built image:
```shell
bazel run --config=`arch` --stamp //cmd/server:go.image -- --norun && \
    docker run -p 8080:8080 -p 9901:9901 -p 9088:9088 -p 18000:18000 -v $HOME:/mnt bazel/cmd/server:go.image
```

### Building Frontend

Run `npm run build` from the `//frontend` directory. Then make sure to `git add` the
`//frontend/build` directory to vendor all generated assets. Just re-run `bazel run` command above
to re-launch with updated frontend code.

## Create New Release

1. Build images for `linux/arm64` and `linux/amd64` architectures stamped with
   the latest Git commit SHA and push them the repository.

ARM:
```shell
bazel run --config=arm64 --stamp //cmd/server:publish
```
Output:
```shell
INFO: Build options --extra_toolchains and --platforms have changed, discarding analysis cache.
INFO: Analyzed target //cmd/server:publish (0 packages loaded, 24765 targets configured).
INFO: Found 1 target...
Target //cmd/server:publish up-to-date:
  bazel-bin/cmd/server/publish.digest
  bazel-bin/cmd/server/publish
INFO: Elapsed time: 52.532s, Critical Path: 40.43s
INFO: 983 processes: 4 internal, 979 darwin-sandbox.
INFO: Build completed successfully, 983 total actions
INFO: Running command line: bazel-bin/cmd/server/publish
2023/07/25 14:46:57 Destination docker.io/apoxy/proximal:{STABLE_GIT_SHA}-arm64 was resolved to docker.io/apoxy/proximal:193b439e50aba2c30823fc7d952b4520b49ae323-arm64 after stamping.
2023/07/25 14:47:26 Successfully pushed Docker image to docker.io/apoxy/proximal:193b439e50aba2c30823fc7d952b4520b49ae323-arm64 - docker.io/apoxy/proximal@sha256:cab49bbb6106bbc3a74a348f132e3537e9178225c994f11351899c13d2287063
```

Intel:
```shell
bazel run --config=amd64 --stamp //cmd/server:publish
```
Output:
```
INFO: Build option --stamp has changed, discarding analysis cache.
INFO: Analyzed target //cmd/server:publish (0 packages loaded, 24888 targets configured).
INFO: Found 1 target...
Target //cmd/server:publish up-to-date:
  bazel-bin/cmd/server/publish.digest
  bazel-bin/cmd/server/publish
INFO: Elapsed time: 0.809s, Critical Path: 0.11s
INFO: 7 processes: 4 internal, 3 darwin-sandbox.
INFO: Build completed successfully, 7 total actions
INFO: Running command line: bazel-bin/cmd/server/publish
2023/07/25 14:45:36 Destination docker.io/apoxy/proximal:{STABLE_GIT_SHA}amd64 was resolved to docker.io/apoxy/proximal:193b439e50aba2c30823fc7d952b4520b49ae323amd64 after stamping.
2023/07/25 14:45:42 Successfully pushed Docker image to docker.io/apoxy/proximal:193b439e50aba2c30823fc7d952b4520b49ae323amd64 - docker.io/apoxy/proximal@sha256:dac87b48e45dbca1f87309fedb16e3ad1ab26f9977f1204101789c1141404788
```

2. Use `buildx imagetools` command to create a single multi-arch manifest referencing the above outputs:

```shell
TAG=<tag>
docker buildx imagetools create -t docker.io/apoxy/proximal:$TAG \
    docker.io/apoxy/proximal:193b439e50aba2c30823fc7d952b4520b49ae323-amd64
    docker.io/apoxy/proximal:193b439e50aba2c30823fc7d952b4520b49ae323-arm64
```

3. Bump the `:latest` manifest:
```
docker buildx imagetools create -t docker.io/apoxy/proximal:latest docker.io/apoxy/proximal:$TAG
```

:tada:!
