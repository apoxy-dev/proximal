# Worker Base Images

There images include all necessary toolchains to perform Wasm/WASI builds. The Proximal
server itself is mostly statically linked.

## Re-building and  the images

1. Re-build multi-arch image and push to registry:
```
REPO=<repo>
TAG=<tag>
docker buildx build --push --tag docker.io/$REPO/proximal-builder:$TAG --platform linux/amd64,linux/arm64 -f Dockerfile.build .
```

2. Update both `base_image_arm64` and `base_image_amd64` container image bases in the `WORKSPACE` file.
