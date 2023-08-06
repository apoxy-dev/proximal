<br /><br />

<div align="center">

<a href="https://apoxy.dev">
  <img src="static/github-proximal.png" alt="Proximal Logo" width="350">
</a>
<br />
<br />

[![Apache 2.0 License](https://badgen.net/badge/License/Apache2.0?icon=github)](LICENSE) [![Slack Community](https://img.shields.io/badge/slack-apoxy-bde868.svg?logo=slack)](http://slack.apoxy.dev/)
  
</div>
<br />

Proximal makes it easy to develop
[Proxy-WASM](https://github.com/proxy-wasm/spec) modules for [Envoy](https://www.envoyproxy.io)
right on your local machine (or anywhere you can run our [Docker
image](https://hub.docker.com/r/apoxy/proximal)).

## TL;DR

Develop WebAssembly extensions for Envoy with ease. Try running `docker run -p 8080:8080 -p 18000:18000 docker.io/apoxy/proximal:latest`
then visit http://localhost:8080

Demo:

https://github.com/apoxy-dev/proximal/assets/767232/97cea009-7f6c-47f9-b2d6-70146ef7ff3a

### What is Proxy-WASM?

[Proxy-WASM](https://github.com/proxy-wasm/spec) (WebAssembly) is a powerful technology that enables
you to extend the functionality of modern proxies like [Envoy](https://www.envoyproxy.io) with
WebAssembly modules. By writing Proxy-WASM modules, you can write code in your L4/L7 proxy that inspects,
mutates, and routes requests as they are passing through, all in a language-independent and sandboxed
environment. It works with both HTTP and TCP-based connection and SDKs are available for
[Rust](https://github.com/proxy-wasm/proxy-wasm-rust-sdk),
[Go](https://github.com/tetratelabs/proxy-wasm-go-sdk),
[C++](https://github.com/proxy-wasm/proxy-wasm-cpp-sdk), and
[AssemblyScript](https://github.com/solo-io/proxy-runtime) (we're working on JavaScript and Python).

These standards based WebAssembly modules can be used with [Istio](https://istio.io/latest/docs/concepts/wasm/),
[MOSN](https://github.com/mosn/mosn) and
[APSIX](https://apisix.apache.org/blog/2021/11/19/apisix-supports-wasm/#how-to-use-wasm-in-apache-apisix) as well.

## Why Proximal?

Developing Proxy-WASM modules for Envoy traditionally involves cumbersome setups, complex
toolchains, and time-consuming testing iterations, frequently on a remote environment.
Proximal simplifies the development process and brings it to your local machine in a single process
environment with a friendly UI and basic REST API.

We believe that developers have been held back from adopting this incredibly powerful technology because
the developer experience for WASM and Proxy-WASM has been a little rough around the edges. Proximal is here to help.

### Key Features:

* **Local Development**: Forget about deploying to remote environments for every code change. Proximal
  allows you to develop and test your Proxy-WASM modules locally, saving you valuable time and
  effort. It features a workflow engine that compiles source code into
  WebAssembly binary (.wasm) and loads them into Envoy automatically.

* **Rapid Iterations**: Change your code and see the results almost instantaneously. Proximal continuously
  watches a working directory (even Docker mounted volumes) and triggers a rebuild/reload of your module in Envoy automatically.

* **Simplified Setup + Examples**: Setting up a development environment for Proxy-WASM can be daunting. Proximal
  streamlines the process and provides a few examples you can use to get started with minimal configuration.

* **Observability**: Debugging is easier with integrated logs capture. See requests and responses in real-time.

## Getting Started

Run via Docker container:

```shell
docker run -p 8080:8080 -p 9901:9901 -p 9088:9088 -p 18000:18000 -v `pwd`:/mnt docker.io/apoxy/proximal:latest
```

The above command mounts your current working directory at `/mnt` inside the container so you can ingest local
Proxy-WASM code (e.g. `/mnt/myprojects/myawesome-proxy-wasm-go/`). Adjust as needed.

Bound ports:
* `8080` - Web UI (see below) and REST API at `:8080/v1/` (see definitions in the [`//api`](https://github.com/apoxy-dev/proximal/tree/main/api) folder).
* `18000` - Envoy listener - test your proxy configurations by sending requests to `localhost:18000`.
* `9901` - Envoy admin UI.
* `9088` - Temporal UI (for build workflow debugging).


## Architecture

We rely on Envoy as the main [data plane](https://en.wikipedia.org/wiki/Forwarding_plane) processing
engine for request routing and its WebAssembly (WASI) extension engine that implements the Proxy-WASM
ABI. The default runtime is [Chromium V8](https://v8.dev) but other runtimes such as
[Wasmtime](https://wasmtime.dev),
[Wamr](https://github.com/bytecodealliance/wasm-micro-runtime), and
[WAVM](https://wavm.github.io/)
can be configured.

The control plane server is a single Go binary that combines an Envoy control plane (using xDS
protocol), a REST API server, a React app, and a [Temporal](https://temporal.io) server
(which is linked directly via the awesome [temporalite](https://github.com/temporalio/temporalite) library)
for managing build workflows. The same binary also acts as a Temporal worker and manages the Envoy process.

Internal state is supported by an embedded SQLite instance which produces an `sqlite3.db` file on local
disk. The Temporal server has its own SQLite db file - `temporalite.db`. Both of these need to be exported
via Docker volume mount if you want state persisted across Docker runs.

Compiled `.wasm` binaries are stored on local disk in the `/tmp/proximal/` directory.

HTML/CSS/JavaScript assets currently live on local filesystem but will be embedded in the binary
itself in the future.

High-level Design:

<div align="center">
  
![proximal-architecture](https://github.com/apoxy-dev/proximal/assets/284347/3585bbae-b014-47cd-aa38-d47a03acacc3)

</div>

## Known Limitations / Future Improvements

Known Limitations:

* The entire setup is a single instance, single binary deal designed for local experimentation.
  While it's possible to run it on remote host since it's packaged in Docker, replication features
  are rather lacking.
* TCP filters aren't yet supported.
* Currently Proximal supports re-triggering builds from a git source manually. Automatic build
  triggers from GitHub commit webhooks or the like aren't suppported since they would require a
  hosted solution with a stable webhook endpoint.

Roadmap:

* More SDKs + Examples - [AssemblyScript](https://github.com/apoxy-dev/proximal/issues/1),
  [C++](https://github.com/apoxy-dev/proximal/issues/2),
  [JavaScript](https://github.com/apoxy-dev/proximal/issues/3), and
  [Python](https://github.com/apoxy-dev/proximal/issues/4).
* Istio examples - show how you take these modules into an existing Istio-enabled cluster.
* K/V store integration.
* Improved logging / tracing / accounting.
* TCP and UDP filters.

If you're interested in any of above features (or maybe something else), feel free to drop a note to the
[Apoxy Team](mailto:hello@apoxy.dev) or open an issue on this repo!

## Contributing

Patches Welcome! (no, really)

Proximal welcomes contributions from the community. If you find bugs, or wish to contribute code, please
check out our [contribution guidelines](DEVELOPING.md) for detailed instructions.

### Support and Feedback

If you encounter any issues, have questions, or want to provide feedback, we want to hear from you!
Feel free to join our active community on Slack, raise an issue on GitHub, or shoot us an email:

* [Apoxy Community Slack](http://slack.apoxy.dev/)
* [👋 Apoxy Email](mailto:hello@apoxy.dev)

## License

Proximal is released under the [Apache 2.0 License](LICENSE).

## Credits

Proximal is developed and maintained by the Apoxy team. We want to thank the open-source community
for their contributions and support in making this project possible. Special thanks go to: [Envoy
Proxy](https://www.envoyproxy.io) community, [Proxy-WASM ABI and
SDKs](https://github.com/proxy-wasm/spec) contributors, and fine folks at
[Temporal](https://temporal.io).

<br />
<br />
<p align="center">
<a href="https://apoxy.dev">
  <img src="static/github-apoxy.png" alt="Apoxy Logo" width="350">
</a>
</p>
<br />
<br />

<p align="center">
Let's take Proxy-WASM development to new levels with Proximal! Happy Proxying! 🚀
</p>
