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

Welcome to Proximal, an open-source project that empowers you to develop
[Proxy-WASM](https://github.com/proxy-wasm/spec) modules for [Envoy](https://www.envoyproxy.io)
right on your local machine (or really anywhere since it's all in a [Docker
image](https://hub.docker.com/r/apoxy/proximal)).

Proximal provides an intuitive and efficient environment to create, test, and iterate on your Envoy
Proxy extensions with ease.

### What is Proxy-WASM?

[Proxy-WASM](https://github.com/proxy-wasm/spec) (WebAssembly) is a powerful technology that enables
you to extend the functionality of modern proxies like [Envoy](https://www.envoyproxy.io). using
WebAssembly (WASI) modules. By writing Proxy-WASM modules, you can tell your L4/L7 proxy to inspect,
mutate, and route requests as they are passing through, all in a language-independent, sandboxed
environment. It works with both HTTP and TCP-based connection and SDKs are available for Rust, Go,
C++, and AssemblyScript (We're working on Javascript and Python).

It is used by popular projects such as Envoy, [Istio](https://istio.io/latest/docs/concepts/wasm/),
and
[APSIX](https://apisix.apache.org/blog/2021/11/19/apisix-supports-wasm/#how-to-use-wasm-in-apache-apisix)

## Why Proximal?

Developing Proxy-WASM modules for Envoy traditionally involves cumbersome setups, complex
toolchains, and time-consuming testing iterations on a remote environment. Proximal changes the game
by bringing the entire development process to a local machine all wrapped up into a single process
environment which comes with UI and REST API.

### Key Features:

* *Local Development*: Forget about deploying to remote environments for every code change. Proximal
  allows you to develop and test your Proxy-WASM modules locally, saving you valuable time and
  effort. It features a full-blown workflow engine to compile Proxy-WASM modules from sources into
  WebAssembly WASI binary representation (.wasm) and load them into Envoy automatically. Git and
  local filesystem sources are supported.

* *Simplified Setup*: Setting up a development environment for Proxy-WASM can be daunting. Proximal
  streamlines the process, ensuring you get started quickly with minimal configuration.

* *Rapid Iterations*: With Proximal, you can rapidly iterate on your code, making changes and seeing
  the results almost instantaneously. Proximal will continuously watch working directory (even
  through mounted volume when in Docker) and trigger a rebuild/reload automatically.

* *Observability*: Debugging Proxy-WASM modules is made easy with Proximal's logs capture support -
  you can see how requests and responses are passing through in real-time.

## Getting Started

Run via Docker container:

```shell
docker run -p 8080:8080 -p 9901:9901 -p 9088:9088 -p 18000:18000 -v $HOME:/mnt docker.io/apoxy/proximal:latest
```

The above command mounts your home directory at `/mnt` inside the container so you can ingest local
Proxy-WASM code (e.g. `/mnt/myprojects/myawesome-proxy-wasm-go/`). Adjust as needed.

Bound ports:
* `8080` - Web UI (see below) and REST API at `:8080/v1/` (see definitions under `//api` folder).
* `9901` - Envoy admin endpoint/UI.
* `9088` - Temporal UI (for easier build workflow debugging).
* `18000` - Port of the Envoy listener - you can test your proxy configurations using `localhost:18000`.

Demo:

https://github.com/apoxy-dev/proximal/assets/767232/97cea009-7f6c-47f9-b2d6-70146ef7ff3a

## Architecture

We rely on Envoy as the main [data plane](https://en.wikipedia.org/wiki/Forwarding_plane) processing
engine for request routing and its WebAssembly (WASI) extension engine which uses Proxy-WASM ABI for
interop between Envoy and WASM runtime. The default runtime is [Chromium V8](https://v8.dev) but
other runtimes such as [Wasmtime](https://wasmtime.dev),
[Wamr](https://github.com/bytecodealliance/wasm-micro-runtime), and [WAVM](https://wavm.github.io/)
can be swapped in.

The control plane server is a single Go binary that combines an Envoy control plane (using xDS
protocol), a REST API server and associated CRUD logic, a React SPA served from Go, and a
[Temporal](https://temporal.io) server (which is linked directly via awesome
[temporalite](https://github.com/temporalio/temporalite) library) for managing build workflows. The
same binary also acts as Temporal worker and manages Envoy process.

Internal state is supported by an embedded SQLite instance which produces `sqlite3.db` file on local
disk. Temporal server has its own SQLite db file `temporalite.db`. Both of these need to be exported
via Docker volume mount if you want state persisted across Docker runs.

Compiled `.wasm` binaries are stored on local disk.

HTML/CSS/Javascript assets currently live on local filesystem but will be embedded in the binary
itself in the future.

High-level Design:

![Proximal (design)](https://github.com/apoxy-dev/proximal/assets/767232/c720a290-3873-428f-b927-525cc31681fc)

## Known Limitations / Future Improvements

Known Limitations:

* The entire setup is a single instance, single binary deal designed for local experimentation.
  While it's possible to run it on remote host since it's packaged in Docker, replication features
  are rather lacking.
* TCP filters aren't yet supported.
* Currently Proximal supports re-triggering builds from git source manually. Automatic build
  triggers from GitHub commit webhooks or the like aren't suppported since they would require a
  hosted solution with a stable webhook endpoint.

TBDs:

* More supported SDKs: AssemblyScript, C++, Javascript, and Python.
* Istio integration - make this thing work with an existing Istio-enabled cluster.
* K/V store integration a la Cloudflare Workers KV.
* More capabable logging / tracing / accounting.
* TCP and possibly UDP filters.

If you're interested in any of above features (or maybe something else), feel free to drop a note to
[Apoxy Team](mailto:hello@apoxy.dev).

## Contributing

Patches Welcome! (no, really)

Proximal is an open-source project, and we welcome contributions from the community. If you find
bugs, or wish to contribute code, please check out our [contribution guidelines](DEVELOPING.md) for
detailed instructions.

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
