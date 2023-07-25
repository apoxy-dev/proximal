load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "51dc53293afe317d2696d4d6433a4c33feedb7748a9e352072e2ec3c0dafd2c6",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.40.1/rules_go-v0.40.1.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.40.1/rules_go-v0.40.1.zip",
    ],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "727f3e4edd96ea20c29e8c2ca9e8d2af724d8c7778e7923a854b2c80952bc405",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.30.0/bazel-gazelle-v0.30.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.30.0/bazel-gazelle-v0.30.0.tar.gz",
    ],
)

http_archive(
    name = "bazel_skylib",
    sha256 = "af87959afe497dc8dfd4c6cb66e1279cb98ccc84284619ebfec27d9c09a903de",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-skylib/releases/download/1.2.0/bazel-skylib-1.2.0.tar.gz",
        "https://github.com/bazelbuild/bazel-skylib/releases/download/1.2.0/bazel-skylib-1.2.0.tar.gz",
    ],
)

BAZEL_ZIG_CC_VERSION = "v1.0.0"

http_archive(
    name = "bazel-zig-cc",
    sha256 = "1f4a1d1e0f6b3e5aa6e1c225fcb23c032f8849441de97b9a38d6ea37362d28e2",
    strip_prefix = "bazel-zig-cc-{}".format(BAZEL_ZIG_CC_VERSION),
    urls = ["https://git.sr.ht/~motiejus/bazel-zig-cc/archive/{}.tar.gz".format(BAZEL_ZIG_CC_VERSION)],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")
load("@bazel_skylib//:workspace.bzl", "bazel_skylib_workspace")

go_repository(
    name = "com_github_acomagu_bufpipe",
    importpath = "github.com/acomagu/bufpipe",
    sum = "h1:e3H4WUzM3npvo5uv95QuJM3cQspFNtFBzvJ2oNjKIDQ=",
    version = "v1.0.4",
)

go_repository(
    name = "com_github_anmitsu_go_shlex",
    importpath = "github.com/anmitsu/go-shlex",
    sum = "h1:9AeTilPcZAjCFIImctFaOjnTIavg87rW78vTPkQqLI8=",
    version = "v0.0.0-20200514113438-38f4b401e2be",
)

go_repository(
    name = "com_github_armon_go_socks5",
    importpath = "github.com/armon/go-socks5",
    sum = "h1:0CwZNZbxp69SHPdPJAN/hZIm0C4OItdklCFmMRWYpio=",
    version = "v0.0.0-20160902184237-e75332964ef5",
)

go_repository(
    name = "com_github_bwesterb_go_ristretto",
    importpath = "github.com/bwesterb/go-ristretto",
    sum = "h1:1w53tCkGhCQ5djbat3+MH0BAQ5Kfgbt56UZQ/JMzngw=",
    version = "v1.2.3",
)

go_repository(
    name = "com_github_cloudflare_circl",
    importpath = "github.com/cloudflare/circl",
    sum = "h1:fE/Qz0QdIGqeWfnwq0RE0R7MI51s0M2E4Ga9kq5AEMs=",
    patch_args = ["-p1"],
    patches = [
        # CGo headers need to be properly referenced for amd64 paltform.
        "//thirdparty:com_github_cloudflare_circl-amd64-cgo.patch",
    ],
    version = "v1.3.3",
)

go_repository(
    name = "com_github_elazarl_goproxy",
    importpath = "github.com/elazarl/goproxy",
    sum = "h1:RIB4cRk+lBqKK3Oy0r2gRX4ui7tuhiZq2SuTtTCi0/0=",
    version = "v0.0.0-20221015165544-a0805db90819",
)

go_repository(
    name = "com_github_gliderlabs_ssh",
    importpath = "github.com/gliderlabs/ssh",
    sum = "h1:OcaySEmAQJgyYcArR+gGGTHCyE7nvhEMTlYY+Dp8CpY=",
    version = "v0.3.5",
)

go_repository(
    name = "com_github_go_git_gcfg",
    importpath = "github.com/go-git/gcfg",
    sum = "h1:+zs/tPmkDkHx3U66DAb0lQFJrpS6731Oaa12ikc+DiI=",
    version = "v1.5.1-0.20230307220236-3a3c6141e376",
)

go_repository(
    name = "com_github_go_git_go_billy_v5",
    importpath = "github.com/go-git/go-billy/v5",
    sum = "h1:Uwp5tDRkPr+l/TnbHOQzp+tmJfLceOlbVucgpTz8ix4=",
    version = "v5.4.1",
)

go_repository(
    name = "com_github_go_git_go_git_fixtures_v4",
    importpath = "github.com/go-git/go-git-fixtures/v4",
    sum = "h1:Pz0DHeFij3XFhoBRGUDPzSJ+w2UcK5/0JvF8DRI58r8=",
    version = "v4.3.2-0.20230305113008-0c11038e723f",
)

go_repository(
    name = "com_github_go_git_go_git_v5",
    importpath = "github.com/go-git/go-git/v5",
    sum = "h1:t9AudWVLmqzlo+4bqdf7GY+46SUuRsx59SboFxkq2aE=",
    version = "v5.7.0",
)

go_repository(
    name = "com_github_imdario_mergo",
    importpath = "github.com/imdario/mergo",
    sum = "h1:M8XP7IuFNsqUx6VPK2P9OSmsYsI/YFaGil0uD21V3dM=",
    version = "v0.3.15",
)

go_repository(
    name = "com_github_jbenet_go_context",
    importpath = "github.com/jbenet/go-context",
    sum = "h1:BQSFePA1RWJOlocH6Fxy8MmwDt+yVQYULKfN0RoTN8A=",
    version = "v0.0.0-20150711004518-d14ea06fba99",
)

go_repository(
    name = "com_github_kevinburke_ssh_config",
    importpath = "github.com/kevinburke/ssh_config",
    sum = "h1:x584FjTGwHzMwvHx18PXxbBVzfnxogHaAReU4gf13a4=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_matryer_is",
    importpath = "github.com/matryer/is",
    sum = "h1:92UTHpy8CDwaJ08GqLDzhhuixiBUUD1p3AU6PHddz4A=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_mmcloughlin_avo",
    importpath = "github.com/mmcloughlin/avo",
    sum = "h1:nAco9/aI9Lg2kiuROBY6BhCI/z0t5jEvJfjWbL8qXLU=",
    version = "v0.5.0",
)

go_repository(
    name = "com_github_pjbgf_sha1cd",
    importpath = "github.com/pjbgf/sha1cd",
    sum = "h1:4D5XXmUUBUl/xQ6IjCkEAbqXskkq/4O7LmGn0AqMDs4=",
    version = "v0.3.0",
)

go_repository(
    name = "com_github_protonmail_go_crypto",
    importpath = "github.com/ProtonMail/go-crypto",
    sum = "h1:ZK3C5DtzV2nVAQTx5S5jQvMeDqWtD1By5mOoyY/xJek=",
    version = "v0.0.0-20230518184743-7afd39499903",
)

go_repository(
    name = "com_github_sergi_go_diff",
    importpath = "github.com/sergi/go-diff",
    sum = "h1:we8PVUC3FE2uYfodKH/nBHMSetSfHDR6scGdBi+erh0=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_skeema_knownhosts",
    importpath = "github.com/skeema/knownhosts",
    sum = "h1:MTk78x9FPgDFVFkDLTrsnnfCJl7g1C/nnKvePgrIngE=",
    version = "v1.1.1",
)

go_repository(
    name = "com_github_xanzy_ssh_agent",
    importpath = "github.com/xanzy/ssh-agent",
    sum = "h1:+/15pJfg/RsTxqYcX6fHqOXZwwMP+2VyYWJeWM2qQFM=",
    version = "v0.3.3",
)

go_repository(
    name = "in_gopkg_warnings_v0",
    importpath = "gopkg.in/warnings.v0",
    sum = "h1:wFXVbFY8DY5/xOe1ECiWdKCzZlxgshcYVNkBHstARME=",
    version = "v0.1.2",
)

go_repository(
    name = "com_github_fsnotify_fsnotify",
    importpath = "github.com/fsnotify/fsnotify",
    sum = "h1:n+5WquG0fcWoWp6xPWfHdbskMCQaFnG6PfBrh1Ky4HY=",
    version = "v1.6.0",
)

bazel_skylib_workspace()

load("@bazel-zig-cc//toolchain:defs.bzl", zig_toolchains = "toolchains")

# This is picking up this change https://github.com/ziglang/zig/pull/13971 or otherwise very
# cryptic, very scary linking errors appear.
# TODO(dilyevsky): Undo this later.
zig_toolchains(
    host_platform_sha256 = {
        "linux-aarch64": "b759a11993949531c692ccfc3d1a004b14df714a7a3515fe0b5c90c9a7631d61",
        "linux-x86_64": "028dad5189e02b2058679b64df16e854a1c1ca0e6044b334d4f3be6e35544f07",
        "macos-aarch64": "5709c27d581988f50f5e6fd5b69d92707787e803a1d04992e290b764617664e6",
        "macos-x86_64": "88d194adb2f3c1a9edbb4a24d018007d5f827a57d1d26b2d9f3459236da1b7b6",
        "windows-x86_64": "75e510bda108e4d78b89d5d1d09e70ea8595fac7c43b5611f280668881adb09d",
    },
    version = "0.11.0-dev.1638+7199d7c77",
)

go_repository(
    name = "com_github_grpc_ecosystem_grpc_gateway",
    build_file_proto_mode = "disable_global",
    build_naming_convention = "go_default_library",
    importpath = "github.com/grpc-ecosystem/grpc-gateway",
    patch_args = ["-p1"],
    patches = [
        # Some linking conflicts resolution for Temporal.
        "//thirdparty:com_github_grpc_ecosystem_grpc_gateway.patch",
    ],
    sum = "h1:gmcG1KaJ57LophUzW0Hy8NmPhnMZb4M0+kPpLofRdBo=",
    version = "v1.16.0",
)

go_repository(
    name = "com_github_antihax_optional",
    importpath = "github.com/antihax/optional",
    sum = "h1:xK2lYat7ZLaVVcIuj82J8kIro4V6kDe0AUDFboUCwcg=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_golang_glog",
    importpath = "github.com/golang/glog",
    sum = "h1:/d3pCKDPWNnvIWe0vVUpNP32qc8U3PDVxySP/y360qE=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_golang_protobuf",
    importpath = "github.com/golang/protobuf",
    sum = "h1:KhyjKVUg7Usr/dYsdSqoFveMYd5ko72D+zANwlG1mmg=",
    version = "v1.5.3",
)

go_repository(
    name = "com_github_google_go_cmp",
    importpath = "github.com/google/go-cmp",
    sum = "h1:O2Tfq5qg4qc4AmwVlvv0oLiVAGB7enBSJ2x2DqQFi38=",
    version = "v0.5.9",
)

go_repository(
    name = "com_github_grpc_ecosystem_grpc_gateway_v2",
    importpath = "github.com/grpc-ecosystem/grpc-gateway/v2",
    sum = "h1:gDLXvp5S9izjldquuoAhDzccbskOL6tDC5jMSyx3zxE=",
    version = "v2.15.2",
)

go_repository(
    name = "com_github_kr_pretty",
    importpath = "github.com/kr/pretty",
    sum = "h1:flRD4NNwYAUpkphVc1HcthR4KEIFJ65n8Mw5qdRn3LE=",
    version = "v0.3.1",
)

go_repository(
    name = "com_github_rogpeppe_fastuuid",
    importpath = "github.com/rogpeppe/fastuuid",
    sum = "h1:Ppwyp6VYCF1nvBTXL3trRso7mXMlRrw9ooo375wvi2s=",
    version = "v1.2.0",
)

go_repository(
    name = "in_gopkg_check_v1",
    importpath = "gopkg.in/check.v1",
    sum = "h1:Hei/4ADfdWqJk1ZMxUNpqntNwaWcugrBjAiHlqqRiVk=",
    version = "v1.0.0-20201130134442-10cb98267c6c",
)

go_repository(
    name = "in_gopkg_yaml_v3",
    importpath = "gopkg.in/yaml.v3",
    sum = "h1:fxVm/GzAzEWqLHuvctI91KS9hhNmmWOoWu0XTYJS7CA=",
    version = "v3.0.1",
)

go_repository(
    name = "org_golang_google_appengine",
    importpath = "google.golang.org/appengine",
    sum = "h1:FZR1q0exgwxzPzp/aF+VccGrSfxfPpkBqjIIEq3ru6c=",
    version = "v1.6.7",
)

go_repository(
    name = "org_golang_google_genproto",
    build_file_proto_mode = "disable_global",
    importpath = "google.golang.org/genproto",
    sum = "h1:8DyZCyvI8mE1IdLy/60bS+52xfymkE72wv1asokgtao=",
    version = "v0.0.0-20230530153820-e85fd2cbaebc",
)

go_repository(
    name = "org_golang_google_genproto_googleapis_api",
    build_file_proto_mode = "disable_global",
    importpath = "google.golang.org/genproto/googleapis/api",
    sum = "h1:kVKPf/IiYSBWEWtkIn6wZXwWGCnLKcC8oWfZvXjsGnM=",
    version = "v0.0.0-20230530153820-e85fd2cbaebc",
)

go_repository(
    name = "org_golang_google_genproto_googleapis_rpc",
    build_file_proto_mode = "disable_global",
    importpath = "google.golang.org/genproto/googleapis/rpc",
    sum = "h1:XSJ8Vk1SWuNr8S18z1NZSziL0CPIXLCCMDOEFtHBOFc=",
    version = "v0.0.0-20230530153820-e85fd2cbaebc",
)

go_repository(
    name = "org_golang_google_grpc",
    build_file_proto_mode = "disable_global",
    importpath = "google.golang.org/grpc",
    sum = "h1:3Oj82/tFSCeUrRTg/5E/7d/W5A1tj6Ky1ABAuZuv5ag=",
    version = "v1.55.0",
)

go_repository(
    name = "org_golang_google_protobuf",
    importpath = "google.golang.org/protobuf",
    sum = "h1:kPPoIgf3TsEvrm0PFe15JQ+570QVxYzEvvHqChK+cng=",
    version = "v1.30.0",
)

go_repository(
    name = "org_golang_x_net",
    importpath = "golang.org/x/net",
    sum = "h1:X2//UzNDwYmtCLn7To6G58Wr6f5ahEAQgKNzv9Y951M=",
    version = "v0.10.0",
)

go_repository(
    name = "org_golang_x_oauth2",
    importpath = "golang.org/x/oauth2",
    sum = "h1:6dkIjl3j3LtZ/O3sTgZTMsLKSftL/B8Zgq4huOIIUu8=",
    version = "v0.8.0",
)

go_repository(
    name = "org_golang_x_sys",
    importpath = "golang.org/x/sys",
    sum = "h1:EBmGv8NaZBZTWvrbjNoL6HVt+IVy3QDQpJs7VRIw3tU=",
    version = "v0.8.0",
)

go_repository(
    name = "org_golang_x_text",
    importpath = "golang.org/x/text",
    sum = "h1:2sjJmO8cDvYveuX97RDLsxlyUxLl+GHoLxBiRdHllBE=",
    version = "v0.9.0",
)

go_repository(
    name = "co_honnef_go_tools",
    importpath = "honnef.co/go/tools",
    sum = "h1:qTakTkI6ni6LFD5sBwwsdSO+AQqbSIxOauHTTQKZ/7o=",
    version = "v0.1.3",
)

go_repository(
    name = "com_github_ajstarks_svgo",
    importpath = "github.com/ajstarks/svgo",
    sum = "h1:slYM766cy2nI3BwyRiyQj/Ud48djTMtMebDqepE95rw=",
    version = "v0.0.0-20211024235047-1546f124cd8b",
)

go_repository(
    name = "com_github_alecthomas_template",
    importpath = "github.com/alecthomas/template",
    sum = "h1:JYp7IbQjafoB+tBA3gMyHYHrpOtNuDiK/uB5uXxq5wM=",
    version = "v0.0.0-20190718012654-fb15b899a751",
)

go_repository(
    name = "com_github_alecthomas_units",
    importpath = "github.com/alecthomas/units",
    sum = "h1:s6gZFSlWYmbqAuRjVTiNNhvNRfY2Wxp9nhfyel4rklc=",
    version = "v0.0.0-20211218093645-b94a6e3cc137",
)

go_repository(
    name = "com_github_apache_thrift",
    importpath = "github.com/apache/thrift",
    sum = "h1:YXuoqgVIHYiAp1WhRw59wXe86HQflof8fh3llIjRzMY=",
    version = "v0.18.0",
)

go_repository(
    name = "com_github_aws_aws_sdk_go",
    importpath = "github.com/aws/aws-sdk-go",
    sum = "h1:pcsP805b9acL3wUqa4JR2vg1k2wnItkDYNvfmcy6F+U=",
    version = "v1.44.203",
)

go_repository(
    name = "com_github_benbjohnson_clock",
    importpath = "github.com/benbjohnson/clock",
    sum = "h1:ip6w0uFQkncKQ979AypyG0ER7mqUSBdKLOgAle/AT8A=",
    version = "v1.3.0",
)

go_repository(
    name = "com_github_beorn7_perks",
    importpath = "github.com/beorn7/perks",
    sum = "h1:VlbKKnNfV8bJzeqoa4cOKqO6bYr3WgKZxO8Z16+hsOM=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_bitly_go_hostpool",
    importpath = "github.com/bitly/go-hostpool",
    sum = "h1:XKmsF6k5el6xHG3WPJ8U0Ku/ye7njX7W81Ng7O2ioR0=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_blang_semver_v4",
    importpath = "github.com/blang/semver/v4",
    sum = "h1:1PFHFE6yCCTv8C1TeyNNarDzntLi7wMI5i/pzqYIsAM=",
    version = "v4.0.0",
)

go_repository(
    name = "com_github_bmizerany_assert",
    importpath = "github.com/bmizerany/assert",
    sum = "h1:DDGfHa7BWjL4YnC6+E63dPcxHo2sUxDIu8g3QgEJdRY=",
    version = "v0.0.0-20160611221934-b7ed37b82869",
)

go_repository(
    name = "com_github_bmizerany_perks",
    importpath = "github.com/bmizerany/perks",
    sum = "h1:AP/Y7sqYicnjGDfD5VcY4CIfh1hRXBUavxrvELjTiOE=",
    version = "v0.0.0-20141205001514-d9a9656a3a4b",
)

go_repository(
    name = "com_github_brianvoe_gofakeit_v6",
    importpath = "github.com/brianvoe/gofakeit/v6",
    sum = "h1:8ihJ60OvPnPJ2W6wZR7M+TTeaZ9bml0z6oy4gvyJ/ek=",
    version = "v6.20.1",
)

go_repository(
    name = "com_github_burntsushi_toml",
    importpath = "github.com/BurntSushi/toml",
    sum = "h1:WXkYYl6Yr3qBf1K79EBnL4mak0OimBfB0XUf9Vl28OQ=",
    version = "v0.3.1",
)

go_repository(
    name = "com_github_burntsushi_xgb",
    importpath = "github.com/BurntSushi/xgb",
    sum = "h1:1BDTz0u9nC3//pOCMdNH+CiXJVYJh5UQNCOBG7jbELc=",
    version = "v0.0.0-20160522181843-27f122750802",
)

go_repository(
    name = "com_github_cactus_go_statsd_client_statsd",
    importpath = "github.com/cactus/go-statsd-client/statsd",
    sum = "h1:HIGF0r/56+7fuIZw2V4isE22MK6xpxWx7BbV8dJ290w=",
    version = "v0.0.0-20200423205355-cb0885a1018c",
)

go_repository(
    name = "com_github_cactus_go_statsd_client_v4",
    importpath = "github.com/cactus/go-statsd-client/v4",
    sum = "h1:cjO9CI3GAHtj/Vmmt1/BRq8+hf5MXy/Pr/CM6Dy4dp0=",
    version = "v4.0.0",
)

go_repository(
    name = "com_github_cenkalti_backoff_v4",
    importpath = "github.com/cenkalti/backoff/v4",
    sum = "h1:y4OZtCnogmCPw98Zjyt5a6+QwPLGkiQsYW5oUqylYbM=",
    version = "v4.2.1",
)

go_repository(
    name = "com_github_census_instrumentation_opencensus_proto",
    # Fixes missing strict dependency errors.
    build_extra_args = ["-exclude=src"],  # keep
    importpath = "github.com/census-instrumentation/opencensus-proto",
    sum = "h1:iKLQ0xPNFxR/2hzXZMrBo8f1j86j5WHzznCCQxV/b8g=",
    version = "v0.4.1",
)

go_repository(
    name = "com_github_cespare_xxhash",
    importpath = "github.com/cespare/xxhash",
    sum = "h1:a6HrQnmkObjyL+Gs60czilIUGqrzKutQD6XZog3p+ko=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_cespare_xxhash_v2",
    importpath = "github.com/cespare/xxhash/v2",
    sum = "h1:DC2CZ1Ep5Y4k3ZQ899DldepgrayRUGE6BBZ/cd9Cj44=",
    version = "v2.2.0",
)

go_repository(
    name = "com_github_chzyer_logex",
    importpath = "github.com/chzyer/logex",
    sum = "h1:Swpa1K6QvQznwJRcfTfQJmTE72DqScAa40E+fbHEXEE=",
    version = "v1.1.10",
)

go_repository(
    name = "com_github_chzyer_readline",
    importpath = "github.com/chzyer/readline",
    sum = "h1:fY5BOSpyZCqRo5OhCuC+XN+r/bBCmeuuJtjz+bCNIf8=",
    version = "v0.0.0-20180603132655-2972be24d48e",
)

go_repository(
    name = "com_github_chzyer_test",
    importpath = "github.com/chzyer/test",
    sum = "h1:q763qf9huN11kDQavWsoZXJNW3xEE4JJyHa5Q25/sd8=",
    version = "v0.0.0-20180213035817-a1ea475d72b1",
)

go_repository(
    name = "com_github_client9_misspell",
    importpath = "github.com/client9/misspell",
    sum = "h1:ta993UF76GwbvJcIo3Y68y/M3WxlpEHPWIGDkJYwzJI=",
    version = "v0.3.4",
)

go_repository(
    name = "com_github_cncf_udpa_go",
    importpath = "github.com/cncf/udpa/go",
    sum = "h1:QQ3GSy+MqSHxm/d8nCtnAiZdYFd45cYZPs8vOOIYKfk=",
    version = "v0.0.0-20220112060539-c52dc94e7fbe",
)

go_repository(
    name = "com_github_cncf_xds_go",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/cncf/xds/go",
    sum = "h1:zlUubfBUxApscKFsF4VSvvfhsBNTBu0eF/ddvpo96yk=",
    version = "v0.0.0-20230428030218-4003588d1b74",
)

go_repository(
    name = "com_github_coreos_go_oidc_v3",
    importpath = "github.com/coreos/go-oidc/v3",
    sum = "h1:6avEvcdvTa1qYsOZ6I5PRkSYHzpTNWgKYmaJfaYbrRw=",
    version = "v3.1.0",
)

go_repository(
    name = "com_github_cpuguy83_go_md2man_v2",
    importpath = "github.com/cpuguy83/go-md2man/v2",
    sum = "h1:p1EgwI/C7NhT0JmVkwCD2ZBK8j4aeHQX2pMHHBfMQ6w=",
    version = "v2.0.2",
)

go_repository(
    name = "com_github_creack_pty",
    importpath = "github.com/creack/pty",
    sum = "h1:uDmaGzcdjhF4i/plgjmEsriH11Y0o7RKapEf/LDaM3w=",
    version = "v1.1.9",
)

go_repository(
    name = "com_github_crossdock_crossdock_go",
    importpath = "github.com/crossdock/crossdock-go",
    sum = "h1:WR1qVJzbvrVywhAk4kMQKRPx09AZVI0NdEdYs59iHcA=",
    version = "v0.0.0-20160816171116-049aabb0122b",
)

go_repository(
    name = "com_github_davecgh_go_spew",
    importpath = "github.com/davecgh/go-spew",
    sum = "h1:vj9j/u1bqnvCEfJOwUhtlOARqs3+rkHYY13jYWTU97c=",
    version = "v1.1.1",
)

go_repository(
    name = "com_github_dgryski_go_farm",
    importpath = "github.com/dgryski/go-farm",
    sum = "h1:fAjc9m62+UWV/WAFKLNi6ZS0675eEUC9y3AlwSbQu1Y=",
    version = "v0.0.0-20200201041132-a6ae2369ad13",
)

go_repository(
    name = "com_github_dustin_go_humanize",
    importpath = "github.com/dustin/go-humanize",
    sum = "h1:GzkhY7T5VNhEkwH0PVJgjz+fX1rhBrR7pRT3mDkpeCY=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_emirpasic_gods",
    importpath = "github.com/emirpasic/gods",
    sum = "h1:FXtiHYKDGKCW2KzwZKx0iC0PQmdlorYgdFG9jPXJ1Bc=",
    version = "v1.18.1",
)

go_repository(
    name = "com_github_envoyproxy_go_control_plane",
    build_directives = [
        # @go_googleapis is what go_rules pass so rewrite resolution rules to avoid conflicts.
        "gazelle:resolve go google.golang.org/genproto/googleapis/api/annotations @go_googleapis//google/api:annotations_go_proto",  # keep
    ],
    build_file_proto_mode = "disable_global",
    importpath = "github.com/envoyproxy/go-control-plane",
    sum = "h1:wSUXTlLfiAQRWs2F+p+EKOY9rUyis1MyGqJ2DIk5HpM=",
    version = "v0.11.1",
)

go_repository(
    name = "com_github_envoyproxy_protoc_gen_validate",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/envoyproxy/protoc-gen-validate",
    sum = "h1:kt9FtLiooDc0vbwTLhdg3dyNX1K9Qwa1EK9LcD4jVUQ=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_facebookgo_clock",
    importpath = "github.com/facebookgo/clock",
    sum = "h1:yDWHCSQ40h88yih2JAcL6Ls/kVkSE8GFACTGVnMPruw=",
    version = "v0.0.0-20150410010913-600d898af40a",
)

go_repository(
    name = "com_github_fatih_color",
    importpath = "github.com/fatih/color",
    sum = "h1:qfhVLaG5s+nCROl1zJsZRxFeYrHLqWroPOQ8BWiNb4w=",
    version = "v1.14.1",
)

go_repository(
    name = "com_github_fogleman_gg",
    importpath = "github.com/fogleman/gg",
    sum = "h1:/7zJX8F6AaYQc57WQCyN9cAIz+4bCJGO9B+dyW29am8=",
    version = "v1.3.0",
)

go_repository(
    name = "com_github_fortytw2_leaktest",
    importpath = "github.com/fortytw2/leaktest",
    sum = "h1:u8491cBMTQ8ft8aeV+adlcytMZylmA5nnwwkRZjI8vw=",
    version = "v1.3.0",
)

go_repository(
    name = "com_github_ghodss_yaml",
    importpath = "github.com/ghodss/yaml",
    sum = "h1:wQHKEahhL6wmXdzwWG11gIVCkOv05bNOh+Rxn0yngAk=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_go_gl_glfw",
    importpath = "github.com/go-gl/glfw",
    sum = "h1:QbL/5oDUmRBzO9/Z7Seo6zf912W/a6Sr4Eu0G/3Jho0=",
    version = "v0.0.0-20190409004039-e6da0acd62b1",
)

go_repository(
    name = "com_github_go_gl_glfw_v3_3_glfw",
    importpath = "github.com/go-gl/glfw/v3.3/glfw",
    sum = "h1:WtGNWLvXpe6ZudgnXrq0barxBImvnnJoMEhXAzcbM0I=",
    version = "v0.0.0-20200222043503-6f7a984d4dc4",
)

go_repository(
    name = "com_github_go_kit_kit",
    importpath = "github.com/go-kit/kit",
    sum = "h1:wDJmvq38kDhkVxi50ni9ykkdUr1PKgqKOoi01fa0Mdk=",
    version = "v0.9.0",
)

go_repository(
    name = "com_github_go_kit_log",
    importpath = "github.com/go-kit/log",
    sum = "h1:MRVx0/zhvdseW+Gza6N9rVzU/IVzaeE1SFI4raAhmBU=",
    version = "v0.2.1",
)

go_repository(
    name = "com_github_go_logfmt_logfmt",
    importpath = "github.com/go-logfmt/logfmt",
    sum = "h1:otpy5pqBCBZ1ng9RQ0dPu4PN7ba75Y/aA+UpowDyNVA=",
    version = "v0.5.1",
)

go_repository(
    name = "com_github_go_logr_logr",
    importpath = "github.com/go-logr/logr",
    sum = "h1:g01GSCwiDw2xSZfjJ2/T9M+S6pFdcNtFYsp+Y43HYDQ=",
    version = "v1.2.4",
)

go_repository(
    name = "com_github_go_logr_stdr",
    importpath = "github.com/go-logr/stdr",
    sum = "h1:hSWxHoqTgW2S2qGc0LTAI563KZ5YKYRhT3MFKZMbjag=",
    version = "v1.2.2",
)

go_repository(
    name = "com_github_go_sql_driver_mysql",
    importpath = "github.com/go-sql-driver/mysql",
    sum = "h1:ozyZYNQW3x3HtqT1jira07DN2PArx2v7/mN66gGcHOs=",
    version = "v1.5.0",
)

go_repository(
    name = "com_github_go_stack_stack",
    importpath = "github.com/go-stack/stack",
    sum = "h1:5SgMzNM5HxrEjV0ww2lTmX6E2Izsfxas4+YHWRs3Lsk=",
    version = "v1.8.0",
)

go_repository(
    name = "com_github_gocql_gocql",
    importpath = "github.com/gocql/gocql",
    sum = "h1:NIlXAJXsjzjGvVn36njh9OLYWzS3D7FdvsifLj4eDEY=",
    version = "v1.4.0",
)

go_repository(
    name = "com_github_gogo_gateway",
    importpath = "github.com/gogo/gateway",
    sum = "h1:u0SuhL9+Il+UbjM9VIE3ntfRujKbvVpFvNB4HbjeVQ0=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_gogo_googleapis",
    build_file_proto_mode = "disable",
    importpath = "github.com/gogo/googleapis",
    sum = "h1:1Yx4Myt7BxzvUr5ldGSbwYiZG6t9wGBZ+8/fX3Wvtq0=",
    version = "v1.4.1",
)

go_repository(
    name = "com_github_gogo_protobuf",
    build_file_proto_mode = "disable",
    importpath = "github.com/gogo/protobuf",
    sum = "h1:Ov1cvc58UF3b5XjBnZv7+opcTcQFZebYjWzi34vdm4Q=",
    version = "v1.3.2",
)

go_repository(
    name = "com_github_gogo_status",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/gogo/status",
    sum = "h1:DuHXlSFHNKqTQ+/ACf5Vs6r4X/dH2EgIzR9Vr+H65kg=",
    version = "v1.1.1",
)

go_repository(
    name = "com_github_golang_freetype",
    importpath = "github.com/golang/freetype",
    sum = "h1:DACJavvAHhabrF08vX0COfcOBJRhZ8lUbR+ZWIs0Y5g=",
    version = "v0.0.0-20170609003504-e2365dfdc4a0",
)

go_repository(
    name = "com_github_golang_groupcache",
    importpath = "github.com/golang/groupcache",
    sum = "h1:oI5xCqsCo564l8iNU+DwB5epxmsaqB+rhGL0m5jtYqE=",
    version = "v0.0.0-20210331224755-41bb18bfe9da",
)

go_repository(
    name = "com_github_golang_jwt_jwt",
    importpath = "github.com/golang-jwt/jwt",
    sum = "h1:IfV12K8xAKAnZqdXVzCZ+TOjboZ2keLg81eXfW3O+oY=",
    version = "v3.2.2+incompatible",
)

go_repository(
    name = "com_github_golang_jwt_jwt_v4",
    importpath = "github.com/golang-jwt/jwt/v4",
    sum = "h1:Hxl6lhQFj4AnOX6MLrsCb/+7tCj7DxP7VA+2rDIq5AU=",
    version = "v4.4.3",
)

go_repository(
    name = "com_github_golang_mock",
    importpath = "github.com/golang/mock",
    sum = "h1:YojYx61/OLFsiv6Rw1Z96LpldJIy31o+UHmwAUMJ6/U=",
    version = "v1.7.0-rc.1",
)

go_repository(
    name = "com_github_golang_snappy",
    importpath = "github.com/golang/snappy",
    sum = "h1:yAGX7huGHXlcLOEtBnF4w7FQwA26wojNCwOYAEhLjQM=",
    version = "v0.0.4",
)

go_repository(
    name = "com_github_google_btree",
    importpath = "github.com/google/btree",
    sum = "h1:0udJVsspx3VBr5FwtLhQQtuAsVc79tTq0ocGIPAU6qo=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_google_gofuzz",
    importpath = "github.com/google/gofuzz",
    sum = "h1:A8PeW59pxE9IoFRqBp37U+mSNaQoZ46F1f0f863XSXw=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_google_martian",
    importpath = "github.com/google/martian",
    sum = "h1:/CP5g8u/VJHijgedC/Legn3BAbAaWPgecwXBIDzw5no=",
    version = "v2.1.0+incompatible",
)

go_repository(
    name = "com_github_google_martian_v3",
    importpath = "github.com/google/martian/v3",
    sum = "h1:IqNFLAmvJOgVlpdEBiQbDc2EwKW77amAycfTuWKdfvw=",
    version = "v3.3.2",
)

go_repository(
    name = "com_github_google_pprof",
    importpath = "github.com/google/pprof",
    sum = "h1:Xim43kblpZXfIBQsbuBVKCudVG457BR2GZFIz3uw3hQ=",
    version = "v0.0.0-20221118152302-e6195bd50e26",
)

go_repository(
    name = "com_github_google_renameio",
    importpath = "github.com/google/renameio",
    sum = "h1:GOZbcHa3HfsPKPlmyPyN2KEohoMXOhdMbHrvbpl2QaA=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_google_uuid",
    importpath = "github.com/google/uuid",
    sum = "h1:t6JiXgmwXMjEs8VusXIJk2BXHsn+wx8BZdTaoZ5fu7I=",
    version = "v1.3.0",
)

go_repository(
    name = "com_github_googleapis_enterprise_certificate_proxy",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/googleapis/enterprise-certificate-proxy",
    sum = "h1:yk9/cqRKtT9wXZSsRH9aurXEpJX+U6FLtpYTdC3R06k=",
    version = "v0.2.3",
)

go_repository(
    name = "com_github_googleapis_gax_go_v2",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/googleapis/gax-go/v2",
    patch_args = ["-p1"],
    patches = [
        "//thirdparty:com_github_googleapis_gax_go_v2-googleapis.patch",
    ],
    sum = "h1:gF4c0zjUP2H/s/hEGyLA3I0fA2ZWjzYiONAD6cvPr8A=",
    version = "v2.7.1",
)

go_repository(
    name = "com_github_googleapis_go_type_adapters",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/googleapis/go-type-adapters",
    sum = "h1:9XdMn+d/G57qq1s8dNc5IesGCXHf6V2HZ2JwRxfA2tA=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_gorilla_securecookie",
    importpath = "github.com/gorilla/securecookie",
    sum = "h1:miw7JPhV+b/lAHSXz4qd/nN9jRiAFV5FwjeKyCS8BvQ=",
    version = "v1.1.1",
)

go_repository(
    name = "com_github_grpc_ecosystem_go_grpc_middleware",
    importpath = "github.com/grpc-ecosystem/go-grpc-middleware",
    sum = "h1:+9834+KizmvFV7pXQGSXQTsaWhq2GjuNUt0aUU0YBYw=",
    version = "v1.3.0",
)

go_repository(
    name = "com_github_hailocab_go_hostpool",
    importpath = "github.com/hailocab/go-hostpool",
    sum = "h1:5upAirOpQc1Q53c0bnx2ufif5kANL7bfZWcc6VJWJd8=",
    version = "v0.0.0-20160125115350-e80d13ce29ed",
)

go_repository(
    name = "com_github_hashicorp_golang_lru",
    importpath = "github.com/hashicorp/golang-lru",
    sum = "h1:0hERBMJE1eitiLkihrMvRVBYAkpHzc/J3QdDN+dAcgU=",
    version = "v0.5.1",
)

go_repository(
    name = "com_github_hdrhistogram_hdrhistogram_go",
    importpath = "github.com/HdrHistogram/hdrhistogram-go",
    sum = "h1:5IcZpTvzydCQeHzK4Ef/D5rrSqwxob0t8PQPMybUNFM=",
    version = "v1.1.2",
)

go_repository(
    name = "com_github_iancoleman_strcase",
    importpath = "github.com/iancoleman/strcase",
    sum = "h1:05I4QRnGpI0m37iZQRuskXh+w77mr6Z41lwQzuHLwW0=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_ianlancetaylor_demangle",
    importpath = "github.com/ianlancetaylor/demangle",
    sum = "h1:mV02weKRL81bEnm8A0HT1/CAelMQDBuQIfLw8n+d6xI=",
    version = "v0.0.0-20200824232613-28f6c0f3b639",
)

go_repository(
    name = "com_github_jessevdk_go_flags",
    importpath = "github.com/jessevdk/go-flags",
    sum = "h1:1jKYvbxEjfUl0fmqTCOfonvskHHXMjBySTLW4y9LFvc=",
    version = "v1.5.0",
)

go_repository(
    name = "com_github_jmespath_go_jmespath",
    importpath = "github.com/jmespath/go-jmespath",
    sum = "h1:BEgLn5cpjn8UN1mAw4NjwDrS35OdebyEtFe+9YPoQUg=",
    version = "v0.4.0",
)

go_repository(
    name = "com_github_jmespath_go_jmespath_internal_testify",
    importpath = "github.com/jmespath/go-jmespath/internal/testify",
    sum = "h1:shLQSRRSCCPj3f2gpwzGwWFoC7ycTf1rcQZHOlsJ6N8=",
    version = "v1.5.1",
)

go_repository(
    name = "com_github_jmoiron_sqlx",
    importpath = "github.com/jmoiron/sqlx",
    sum = "h1:wv+0IJZfL5z0uZoUjlpKgHkgaFSYD+r9CfrXjEXsO7w=",
    version = "v1.3.4",
)

go_repository(
    name = "com_github_jonboulle_clockwork",
    importpath = "github.com/jonboulle/clockwork",
    sum = "h1:p4Cf1aMWXnXAUh8lVfewRBx1zaTSYKrKMF2g3ST4RZ4=",
    version = "v0.4.0",
)

go_repository(
    name = "com_github_josharian_intern",
    importpath = "github.com/josharian/intern",
    sum = "h1:vlS4z54oSdjm0bgjRigI+G1HpF+tI+9rE5LLzOg8HmY=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_jpillora_backoff",
    importpath = "github.com/jpillora/backoff",
    sum = "h1:uvFg412JmmHBHw7iwprIxkPMI+sGQ4kzOWsMeHnm2EA=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_json_iterator_go",
    importpath = "github.com/json-iterator/go",
    sum = "h1:PV8peI4a0ysnczrg+LtxykD8LfKY9ML6u2jnxaEnrnM=",
    version = "v1.1.12",
)

go_repository(
    name = "com_github_jstemmer_go_junit_report",
    importpath = "github.com/jstemmer/go-junit-report",
    sum = "h1:6QPYqodiu3GuPL+7mfx+NwDdp2eTkp9IfEUpgAwUN0o=",
    version = "v0.9.1",
)

go_repository(
    name = "com_github_julienschmidt_httprouter",
    importpath = "github.com/julienschmidt/httprouter",
    sum = "h1:U0609e9tgbseu3rBINet9P48AI/D3oJs4dN7jwJOQ1U=",
    version = "v1.3.0",
)

go_repository(
    name = "com_github_jung_kurt_gofpdf",
    importpath = "github.com/jung-kurt/gofpdf",
    sum = "h1:PJr+ZMXIecYc1Ey2zucXdR73SMBtgjPgwa31099IMv0=",
    version = "v1.0.3-0.20190309125859-24315acbbda5",
)

go_repository(
    name = "com_github_kballard_go_shellquote",
    importpath = "github.com/kballard/go-shellquote",
    sum = "h1:Z9n2FFNUXsshfwJMBgNA0RU6/i7WVaAegv3PtuIHPMs=",
    version = "v0.0.0-20180428030007-95032a82bc51",
)

go_repository(
    name = "com_github_kisielk_errcheck",
    importpath = "github.com/kisielk/errcheck",
    sum = "h1:e8esj/e4R+SAOwFwN+n3zr0nYeCyeweozKfO23MvHzY=",
    version = "v1.5.0",
)

go_repository(
    name = "com_github_kisielk_gotool",
    importpath = "github.com/kisielk/gotool",
    sum = "h1:AV2c/EiW3KqPNT9ZKl07ehoAGi4C5/01Cfbblndcapg=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_konsorten_go_windows_terminal_sequences",
    importpath = "github.com/konsorten/go-windows-terminal-sequences",
    sum = "h1:CE8S1cTafDpPvMhIxNJKvHsGVBgn1xWYf1NbHQhywc8=",
    version = "v1.0.3",
)

go_repository(
    name = "com_github_kr_logfmt",
    importpath = "github.com/kr/logfmt",
    sum = "h1:T+h1c/A9Gawja4Y9mFVWj2vyii2bbUNDw3kt9VxK2EY=",
    version = "v0.0.0-20140226030751-b84e30acd515",
)

go_repository(
    name = "com_github_kr_pty",
    importpath = "github.com/kr/pty",
    sum = "h1:VkoXIwSboBpnk99O/KFauAEILuNHv5DVFKZMBN/gUgw=",
    version = "v1.1.1",
)

go_repository(
    name = "com_github_kr_text",
    importpath = "github.com/kr/text",
    sum = "h1:5Nx0Ya0ZqY2ygV366QzturHI13Jq95ApcVaJBhpS+AY=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_labstack_echo_v4",
    importpath = "github.com/labstack/echo/v4",
    sum = "h1:GliPYSpzGKlyOhqIbG8nmHBo3i1saKWFOgh41AN3b+Y=",
    version = "v4.9.1",
)

go_repository(
    name = "com_github_labstack_gommon",
    importpath = "github.com/labstack/gommon",
    sum = "h1:y7cvthEAEbU0yHOf4axH8ZG2NH8knB9iNSoTO8dyIk8=",
    version = "v0.4.0",
)

go_repository(
    name = "com_github_lib_pq",
    importpath = "github.com/lib/pq",
    sum = "h1:p7ZhMD+KsSRozJr34udlUrhboJwWAgCg34+/ZZNvZZw=",
    version = "v1.10.7",
)

go_repository(
    name = "com_github_mailru_easyjson",
    importpath = "github.com/mailru/easyjson",
    sum = "h1:UGYAvKxe3sBsEDzO8ZeWOSlIQfWFlxbzLZe7hwFURr0=",
    version = "v0.7.7",
)

go_repository(
    name = "com_github_mattn_go_colorable",
    importpath = "github.com/mattn/go-colorable",
    sum = "h1:fFA4WZxdEF4tXPZVKMLwD8oUnCTTo08duU7wxecdEvA=",
    version = "v0.1.13",
)

go_repository(
    name = "com_github_mattn_go_isatty",
    importpath = "github.com/mattn/go-isatty",
    sum = "h1:BTarxUcIeDqL27Mc+vyvdWYSL28zpIhv3RoTdsLMPng=",
    version = "v0.0.17",
)

go_repository(
    name = "com_github_mattn_go_runewidth",
    importpath = "github.com/mattn/go-runewidth",
    sum = "h1:+xnbZSEeDbOIg5/mE6JF0w6n9duR1l3/WmbinWVwUuU=",
    version = "v0.0.14",
)

go_repository(
    name = "com_github_mattn_go_sqlite3",
    importpath = "github.com/mattn/go-sqlite3",
    sum = "h1:mCRHCLDUBXgpKAqIKsaAaAsrAlbkeomtRFKXh2L6YIM=",
    version = "v1.14.17",
)

go_repository(
    name = "com_github_matttproud_golang_protobuf_extensions",
    importpath = "github.com/matttproud/golang_protobuf_extensions",
    sum = "h1:mmDVorXM7PCGKw94cs5zkfA9PSy5pEvNWRP0ET0TIVo=",
    version = "v1.0.4",
)

go_repository(
    name = "com_github_modern_go_concurrent",
    importpath = "github.com/modern-go/concurrent",
    sum = "h1:TRLaZ9cD/w8PVh93nsPXa1VrQ6jlwL5oN8l14QlcNfg=",
    version = "v0.0.0-20180306012644-bacd9c7ef1dd",
)

go_repository(
    name = "com_github_modern_go_reflect2",
    importpath = "github.com/modern-go/reflect2",
    sum = "h1:xBagoLtFs94CBntxluKeaWgTMpvLxC4ur3nMaC9Gz0M=",
    version = "v1.0.2",
)

go_repository(
    name = "com_github_mwitkow_go_conntrack",
    importpath = "github.com/mwitkow/go-conntrack",
    sum = "h1:KUppIJq7/+SVif2QVs3tOP0zanoHgBEVAwHxUSIzRqU=",
    version = "v0.0.0-20190716064945-2f068394615f",
)

go_repository(
    name = "com_github_niemeyer_pretty",
    importpath = "github.com/niemeyer/pretty",
    sum = "h1:fD57ERR4JtEqsWbfPhv4DMiApHyliiK5xCTNVSPiaAs=",
    version = "v0.0.0-20200227124842-a10e7caefd8e",
)

go_repository(
    name = "com_github_olekukonko_tablewriter",
    importpath = "github.com/olekukonko/tablewriter",
    sum = "h1:P2Ga83D34wi1o9J6Wh1mRuqd4mF/x/lgBS7N7AbDhec=",
    version = "v0.0.5",
)

go_repository(
    name = "com_github_olivere_elastic_v7",
    importpath = "github.com/olivere/elastic/v7",
    sum = "h1:R7CXvbu8Eq+WlsLgxmKVKPox0oOwAE/2T9Si5BnvK6E=",
    version = "v7.0.32",
)

go_repository(
    name = "com_github_oneofone_xxhash",
    importpath = "github.com/OneOfOne/xxhash",
    sum = "h1:KMrpdQIwFcEqXDklaen+P1axHaj9BSKzvpUUfnHldSE=",
    version = "v1.2.2",
)

go_repository(
    name = "com_github_opentracing_opentracing_go",
    importpath = "github.com/opentracing/opentracing-go",
    sum = "h1:uEJPy/1a5RIPAJ0Ov+OIO8OxWu77jEv+1B0VhjKrZUs=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_pborman_uuid",
    importpath = "github.com/pborman/uuid",
    sum = "h1:+ZZIw58t/ozdjRaXh/3awHfmWRbzYxJoAdNJxe/3pvw=",
    version = "v1.2.1",
)

go_repository(
    name = "com_github_pkg_errors",
    importpath = "github.com/pkg/errors",
    sum = "h1:FEBLx1zS214owpjy7qsBeixbURkuhQAwrK5UwLGTwt4=",
    version = "v0.9.1",
)

go_repository(
    name = "com_github_pmezard_go_difflib",
    importpath = "github.com/pmezard/go-difflib",
    sum = "h1:4DBwDE0NGyQoBHbLQYPwSUPoCMWR5BEzIk/f1lZbAQM=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_prashantv_protectmem",
    importpath = "github.com/prashantv/protectmem",
    sum = "h1:AA9vgIBDjMHPC2McaGPojgV2dcI78ZC0TLNhYCXEKH8=",
    version = "v0.0.0-20171002184600-e20412882b3a",
)

go_repository(
    name = "com_github_prometheus_client_golang",
    importpath = "github.com/prometheus/client_golang",
    sum = "h1:8tXpTmJbyH5lydzFPoxSIJ0J46jdh3tylbvM1xCv0LI=",
    version = "v1.15.1",
)

go_repository(
    name = "com_github_prometheus_client_model",
    importpath = "github.com/prometheus/client_model",
    sum = "h1:5lQXD3cAg1OXBf4Wq03gTrXHeaV0TQvGfUooCfx1yqY=",
    version = "v0.4.0",
)

go_repository(
    name = "com_github_prometheus_common",
    importpath = "github.com/prometheus/common",
    sum = "h1:+5BrQJwiBB9xsMygAB3TNvpQKOwlkc25LbISbrdOOfY=",
    version = "v0.44.0",
)

go_repository(
    name = "com_github_prometheus_procfs",
    importpath = "github.com/prometheus/procfs",
    sum = "h1:kYK1Va/YMlutzCGazswoHKo//tZVlFpKYh+PymziUAg=",
    version = "v0.10.1",
)

go_repository(
    name = "com_github_rcrowley_go_metrics",
    importpath = "github.com/rcrowley/go-metrics",
    sum = "h1:N/ElC8H3+5XpJzTSTfLsJV/mx9Q9g7kxmchpfZyxgzM=",
    version = "v0.0.0-20201227073835-cf1acfcdf475",
)

go_repository(
    name = "com_github_remyoudompheng_bigfft",
    importpath = "github.com/remyoudompheng/bigfft",
    sum = "h1:W09IVJc94icq4NjY3clb7Lk8O1qJ8BdBEF8z0ibU0rE=",
    version = "v0.0.0-20230129092748-24d4a6f8daec",
)

go_repository(
    name = "com_github_rivo_uniseg",
    importpath = "github.com/rivo/uniseg",
    sum = "h1:utMvzDsuh3suAEnhH0RdHmoPbU648o6CvXxTx4SBMOw=",
    version = "v0.4.3",
)

go_repository(
    name = "com_github_robfig_cron",
    importpath = "github.com/robfig/cron",
    sum = "h1:ZjScXvvxeQ63Dbyxy76Fj3AT3Ut0aKsyd2/tl3DTMuQ=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_robfig_cron_v3",
    importpath = "github.com/robfig/cron/v3",
    sum = "h1:WdRxkvbJztn8LMz/QEvLN5sBU+xKpSqwwUO1Pjr4qDs=",
    version = "v3.0.1",
)

go_repository(
    name = "com_github_rogpeppe_go_internal",
    importpath = "github.com/rogpeppe/go-internal",
    sum = "h1:TMyTOH3F/DB16zRVcYyreMH6GnZZrwQVAoYjRBZyWFQ=",
    version = "v1.10.0",
)

go_repository(
    name = "com_github_russross_blackfriday_v2",
    importpath = "github.com/russross/blackfriday/v2",
    sum = "h1:JIOH55/0cWyOuilr9/qlrm0BSXldqnqwMsf35Ld67mk=",
    version = "v2.1.0",
)

go_repository(
    name = "com_github_samuel_go_thrift",
    importpath = "github.com/samuel/go-thrift",
    sum = "h1:gD4vkYmuoWVgdV6UwI3tPo9MtMfVoIRY+Xn9919SJBg=",
    version = "v0.0.0-20190219015601-e8b6b52668fe",
)

go_repository(
    name = "com_github_sirupsen_logrus",
    importpath = "github.com/sirupsen/logrus",
    sum = "h1:oxx1eChJGI6Uks2ZC4W1zpLlVgqB8ner4EuQwV4Ik1Y=",
    version = "v1.9.2",
)

go_repository(
    name = "com_github_smartystreets_assertions",
    importpath = "github.com/smartystreets/assertions",
    sum = "h1:T/YLemO5Yp7KPzS+lVtu+WsHn8yoSwTfItdAd1r3cck=",
    version = "v1.1.1",
)

go_repository(
    name = "com_github_smartystreets_go_aws_auth",
    importpath = "github.com/smartystreets/go-aws-auth",
    sum = "h1:hp2CYQUINdZMHdvTdXtPOY2ainKl4IoMcpAXEf2xj3Q=",
    version = "v0.0.0-20180515143844-0c1422d1fdb9",
)

go_repository(
    name = "com_github_smartystreets_gunit",
    importpath = "github.com/smartystreets/gunit",
    sum = "h1:tyWYZffdPhQPfK5VsMQXfauwnJkqg7Tv5DLuQVYxq3Q=",
    version = "v1.4.2",
)

go_repository(
    name = "com_github_spaolacci_murmur3",
    importpath = "github.com/spaolacci/murmur3",
    sum = "h1:qLC7fQah7D6K1B0ujays3HV9gkFtllcxhzImRR7ArPQ=",
    version = "v0.0.0-20180118202830-f09979ecbc72",
)

go_repository(
    name = "com_github_stretchr_objx",
    importpath = "github.com/stretchr/objx",
    sum = "h1:1zr/of2m5FGMsad5YfcqgdqdWrIhu+EBEJRhR1U7z/c=",
    version = "v0.5.0",
)

go_repository(
    name = "com_github_stretchr_testify",
    importpath = "github.com/stretchr/testify",
    sum = "h1:CcVxjf3Q8PM0mHUKJCdn+eZZtm5yQwehR5yeSVQQcUk=",
    version = "v1.8.4",
)

go_repository(
    name = "com_github_temporalio_ringpop_go",
    importpath = "github.com/temporalio/ringpop-go",
    sum = "h1:QIwUh2HCtmtB4rDM4CYnmX8ep9X7n9WZh+0rv18FySc=",
    version = "v0.0.0-20220818230611-30bf23b490b2",
)

go_repository(
    name = "com_github_temporalio_tchannel_go",
    importpath = "github.com/temporalio/tchannel-go",
    sum = "h1:Fs3LdlF7xbnOWHymbFmvIEuxIEt1dNRCfaDkoajSaZk=",
    version = "v1.22.1-0.20220818200552-1be8d8cffa5b",
)

go_repository(
    name = "com_github_temporalio_tctl_kit",
    importpath = "github.com/temporalio/tctl-kit",
    sum = "h1:N+lz4Nav7634VmiCmQlbudy2ja6s5NAqc/5dUd4sOw8=",
    version = "v0.0.0-20230213052353-2342ea1e7d14",
)

go_repository(
    name = "com_github_temporalio_temporalite",
    importpath = "github.com/temporalio/temporalite",
    patch_args = ["-p1"],
    patches = [
        "//thirdparty:com_github_temporalio_temporalite-new-server-return.patch",
    ],
    sum = "h1:qA+JJDC5tNd1ZS1zgk660QPOl6UK4LVbJOPpSX4vrCs=",
    version = "v0.3.0",
)

go_repository(
    name = "com_github_temporalio_ui_server_v2",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/temporalio/ui-server/v2",
    sum = "h1:eII8RPOa2UGLZp+bzZdZbnqcR28SIACv+RWE/N3k0rY=",
    version = "v2.8.3",
)

go_repository(
    name = "com_github_twmb_murmur3",
    importpath = "github.com/twmb/murmur3",
    sum = "h1:mqrRot1BRxm+Yct+vavLMou2/iJt0tNVTTC0QoIjaZg=",
    version = "v1.1.6",
)

go_repository(
    name = "com_github_uber_common_bark",
    importpath = "github.com/uber-common/bark",
    sum = "h1:DkuZCBaQS9LWuNAPrCO6yQVANckIX3QI0QwLemUnzCo=",
    version = "v1.3.0",
)

go_repository(
    name = "com_github_uber_go_tally",
    importpath = "github.com/uber-go/tally",
    sum = "h1:9hLSgNBP28CjIaDmAuRTq9qV+UZY+9PcvAkXO4nNMwg=",
    version = "v3.3.15+incompatible",
)

go_repository(
    name = "com_github_uber_go_tally_v4",
    importpath = "github.com/uber-go/tally/v4",
    sum = "h1:k8NwZvdfj/w398Q0H1nqi6yq2W5bgxsqn76kNwZ/hfM=",
    version = "v4.1.6",
)

go_repository(
    name = "com_github_uber_jaeger_client_go",
    importpath = "github.com/uber/jaeger-client-go",
    sum = "h1:D6wyKGCecFaSRUpo8lCVbaOOb6ThwMmTEbhRwtKR97o=",
    version = "v2.30.0+incompatible",
)

go_repository(
    name = "com_github_uber_jaeger_lib",
    importpath = "github.com/uber/jaeger-lib",
    sum = "h1:td4jdvLcExb4cBISKIpHuGoVXh+dVKhn2Um6rjCsSsg=",
    version = "v2.4.1+incompatible",
)

go_repository(
    name = "com_github_urfave_cli",
    importpath = "github.com/urfave/cli",
    sum = "h1:igJgVw1JdKH+trcLWLeLwZjU9fEfPesQ+9/e4MQ44S8=",
    version = "v1.22.12",
)

go_repository(
    name = "com_github_urfave_cli_v2",
    importpath = "github.com/urfave/cli/v2",
    sum = "h1:xbrU7tAYviSpqeR3X4nEFWUdB/uDZ6DE+HxmRU7Xtyw=",
    version = "v2.23.5",
)

go_repository(
    name = "com_github_valyala_bytebufferpool",
    importpath = "github.com/valyala/bytebufferpool",
    sum = "h1:GqA5TC/0021Y/b9FG4Oi9Mr3q7XYx6KllzawFIhcdPw=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_valyala_fasttemplate",
    importpath = "github.com/valyala/fasttemplate",
    sum = "h1:TVEnxayobAdVkhQfrfes2IzOB6o+z4roRkPF52WA1u4=",
    version = "v1.2.1",
)

go_repository(
    name = "com_github_xrash_smetrics",
    importpath = "github.com/xrash/smetrics",
    sum = "h1:bAn7/zixMGCfxrRTfdpNzjtPYqr8smhKouy9mxVdGPU=",
    version = "v0.0.0-20201216005158-039620a65673",
)

go_repository(
    name = "com_github_xwb1989_sqlparser",
    importpath = "github.com/xwb1989/sqlparser",
    sum = "h1:zzrxE1FKn5ryBNl9eKOeqQ58Y/Qpo3Q9QNxKHX5uzzQ=",
    version = "v0.0.0-20180606152119-120387863bf2",
)

go_repository(
    name = "com_github_yuin_goldmark",
    importpath = "github.com/yuin/goldmark",
    sum = "h1:fVcFKWvrslecOb/tg+Cc05dkeYx540o0FuFt3nUVDoE=",
    version = "v1.4.13",
)

go_repository(
    name = "com_google_cloud_go",
    importpath = "cloud.google.com/go",
    sum = "h1:Zc8gqp3+a9/Eyph2KDmcGaPtbKRIoqq4YTlL4NMD0Ys=",
    version = "v0.110.0",
)

go_repository(
    name = "com_google_cloud_go_accessapproval",
    importpath = "cloud.google.com/go/accessapproval",
    sum = "h1:x0cEHro/JFPd7eS4BlEWNTMecIj2HdXjOVB5BtvwER0=",
    version = "v1.6.0",
)

go_repository(
    name = "com_google_cloud_go_accesscontextmanager",
    importpath = "cloud.google.com/go/accesscontextmanager",
    sum = "h1:MG60JgnEoawHJrbWw0jGdv6HLNSf6gQvYRiXpuzqgEA=",
    version = "v1.7.0",
)

go_repository(
    name = "com_google_cloud_go_aiplatform",
    importpath = "cloud.google.com/go/aiplatform",
    sum = "h1:zTw+suCVchgZyO+k847wjzdVjWmrAuehxdvcZvJwfGg=",
    version = "v1.37.0",
)

go_repository(
    name = "com_google_cloud_go_analytics",
    importpath = "cloud.google.com/go/analytics",
    sum = "h1:LqAo3tAh2FU9+w/r7vc3hBjU23Kv7GhO/PDIW7kIYgM=",
    version = "v0.19.0",
)

go_repository(
    name = "com_google_cloud_go_apigateway",
    importpath = "cloud.google.com/go/apigateway",
    sum = "h1:ZI9mVO7x3E9RK/BURm2p1aw9YTBSCQe3klmyP1WxWEg=",
    version = "v1.5.0",
)

go_repository(
    name = "com_google_cloud_go_apigeeconnect",
    importpath = "cloud.google.com/go/apigeeconnect",
    sum = "h1:sWOmgDyAsi1AZ48XRHcATC0tsi9SkPT7DA/+VCfkaeA=",
    version = "v1.5.0",
)

go_repository(
    name = "com_google_cloud_go_apigeeregistry",
    importpath = "cloud.google.com/go/apigeeregistry",
    sum = "h1:E43RdhhCxdlV+I161gUY2rI4eOaMzHTA5kNkvRsFXvc=",
    version = "v0.6.0",
)

go_repository(
    name = "com_google_cloud_go_appengine",
    importpath = "cloud.google.com/go/appengine",
    sum = "h1:aBGDKmRIaRRoWJ2tAoN0oVSHoWLhtO9aj/NvUyP4aYs=",
    version = "v1.7.1",
)

go_repository(
    name = "com_google_cloud_go_area120",
    importpath = "cloud.google.com/go/area120",
    sum = "h1:ugckkFh4XkHJMPhTIx0CyvdoBxmOpMe8rNs4Ok8GAag=",
    version = "v0.7.1",
)

go_repository(
    name = "com_google_cloud_go_artifactregistry",
    importpath = "cloud.google.com/go/artifactregistry",
    sum = "h1:o1Q80vqEB6Qp8WLEH3b8FBLNUCrGQ4k5RFj0sn/sgO8=",
    version = "v1.13.0",
)

go_repository(
    name = "com_google_cloud_go_asset",
    importpath = "cloud.google.com/go/asset",
    sum = "h1:YAsssO08BqZ6mncbb6FPlj9h6ACS7bJQUOlzciSfbNk=",
    version = "v1.13.0",
)

go_repository(
    name = "com_google_cloud_go_assuredworkloads",
    importpath = "cloud.google.com/go/assuredworkloads",
    sum = "h1:VLGnVFta+N4WM+ASHbhc14ZOItOabDLH1MSoDv+Xuag=",
    version = "v1.10.0",
)

go_repository(
    name = "com_google_cloud_go_automl",
    importpath = "cloud.google.com/go/automl",
    sum = "h1:50VugllC+U4IGl3tDNcZaWvApHBTrn/TvyHDJ0wM+Uw=",
    version = "v1.12.0",
)

go_repository(
    name = "com_google_cloud_go_baremetalsolution",
    importpath = "cloud.google.com/go/baremetalsolution",
    sum = "h1:2AipdYXL0VxMboelTTw8c1UJ7gYu35LZYUbuRv9Q28s=",
    version = "v0.5.0",
)

go_repository(
    name = "com_google_cloud_go_batch",
    importpath = "cloud.google.com/go/batch",
    sum = "h1:YbMt0E6BtqeD5FvSv1d56jbVsWEzlGm55lYte+M6Mzs=",
    version = "v0.7.0",
)

go_repository(
    name = "com_google_cloud_go_beyondcorp",
    importpath = "cloud.google.com/go/beyondcorp",
    sum = "h1:UkY2BTZkEUAVrgqnSdOJ4p3y9ZRBPEe1LkjgC8Bj/Pc=",
    version = "v0.5.0",
)

go_repository(
    name = "com_google_cloud_go_bigquery",
    importpath = "cloud.google.com/go/bigquery",
    sum = "h1:RscMV6LbnAmhAzD893Lv9nXXy2WCaJmbxYPWDLbGqNQ=",
    version = "v1.50.0",
)

go_repository(
    name = "com_google_cloud_go_billing",
    importpath = "cloud.google.com/go/billing",
    sum = "h1:JYj28UYF5w6VBAh0gQYlgHJ/OD1oA+JgW29YZQU+UHM=",
    version = "v1.13.0",
)

go_repository(
    name = "com_google_cloud_go_binaryauthorization",
    importpath = "cloud.google.com/go/binaryauthorization",
    sum = "h1:d3pMDBCCNivxt5a4eaV7FwL7cSH0H7RrEnFrTb1QKWs=",
    version = "v1.5.0",
)

go_repository(
    name = "com_google_cloud_go_certificatemanager",
    importpath = "cloud.google.com/go/certificatemanager",
    sum = "h1:5C5UWeSt8Jkgp7OWn2rCkLmYurar/vIWIoSQ2+LaTOc=",
    version = "v1.6.0",
)

go_repository(
    name = "com_google_cloud_go_channel",
    importpath = "cloud.google.com/go/channel",
    sum = "h1:GpcQY5UJKeOekYgsX3QXbzzAc/kRGtBq43fTmyKe6Uw=",
    version = "v1.12.0",
)

go_repository(
    name = "com_google_cloud_go_cloudbuild",
    importpath = "cloud.google.com/go/cloudbuild",
    sum = "h1:GHQCjV4WlPPVU/j3Rlpc8vNIDwThhd1U9qSY/NPZdko=",
    version = "v1.9.0",
)

go_repository(
    name = "com_google_cloud_go_clouddms",
    importpath = "cloud.google.com/go/clouddms",
    sum = "h1:E7v4TpDGUyEm1C/4KIrpVSOCTm0P6vWdHT0I4mostRA=",
    version = "v1.5.0",
)

go_repository(
    name = "com_google_cloud_go_cloudtasks",
    importpath = "cloud.google.com/go/cloudtasks",
    sum = "h1:uK5k6abf4yligFgYFnG0ni8msai/dSv6mDmiBulU0hU=",
    version = "v1.10.0",
)

go_repository(
    name = "com_google_cloud_go_compute",
    importpath = "cloud.google.com/go/compute",
    sum = "h1:+9zda3WGgW1ZSTlVppLCYFIr48Pa35q1uG2N1itbCEQ=",
    version = "v1.19.0",
)

go_repository(
    name = "com_google_cloud_go_compute_metadata",
    importpath = "cloud.google.com/go/compute/metadata",
    sum = "h1:mg4jlk7mCAj6xXp9UJ4fjI9VUI5rubuGBW5aJ7UnBMY=",
    version = "v0.2.3",
)

go_repository(
    name = "com_google_cloud_go_contactcenterinsights",
    importpath = "cloud.google.com/go/contactcenterinsights",
    sum = "h1:jXIpfcH/VYSE1SYcPzO0n1VVb+sAamiLOgCw45JbOQk=",
    version = "v1.6.0",
)

go_repository(
    name = "com_google_cloud_go_container",
    importpath = "cloud.google.com/go/container",
    sum = "h1:NKlY/wCDapfVZlbVVaeuu2UZZED5Dy1z4Zx1KhEzm8c=",
    version = "v1.15.0",
)

go_repository(
    name = "com_google_cloud_go_containeranalysis",
    importpath = "cloud.google.com/go/containeranalysis",
    sum = "h1:EQ4FFxNaEAg8PqQCO7bVQfWz9NVwZCUKaM1b3ycfx3U=",
    version = "v0.9.0",
)

go_repository(
    name = "com_google_cloud_go_datacatalog",
    importpath = "cloud.google.com/go/datacatalog",
    sum = "h1:4H5IJiyUE0X6ShQBqgFFZvGGcrwGVndTwUSLP4c52gw=",
    version = "v1.13.0",
)

go_repository(
    name = "com_google_cloud_go_dataflow",
    importpath = "cloud.google.com/go/dataflow",
    sum = "h1:eYyD9o/8Nm6EttsKZaEGD84xC17bNgSKCu0ZxwqUbpg=",
    version = "v0.8.0",
)

go_repository(
    name = "com_google_cloud_go_dataform",
    importpath = "cloud.google.com/go/dataform",
    sum = "h1:Dyk+fufup1FR6cbHjFpMuP4SfPiF3LI3JtoIIALoq48=",
    version = "v0.7.0",
)

go_repository(
    name = "com_google_cloud_go_datafusion",
    importpath = "cloud.google.com/go/datafusion",
    sum = "h1:sZjRnS3TWkGsu1LjYPFD/fHeMLZNXDK6PDHi2s2s/bk=",
    version = "v1.6.0",
)

go_repository(
    name = "com_google_cloud_go_datalabeling",
    importpath = "cloud.google.com/go/datalabeling",
    sum = "h1:ch4qA2yvddGRUrlfwrNJCr79qLqhS9QBwofPHfFlDIk=",
    version = "v0.7.0",
)

go_repository(
    name = "com_google_cloud_go_dataplex",
    importpath = "cloud.google.com/go/dataplex",
    sum = "h1:RvoZ5T7gySwm1CHzAw7yY1QwwqaGswunmqEssPxU/AM=",
    version = "v1.6.0",
)

go_repository(
    name = "com_google_cloud_go_dataproc",
    importpath = "cloud.google.com/go/dataproc",
    sum = "h1:W47qHL3W4BPkAIbk4SWmIERwsWBaNnWm0P2sdx3YgGU=",
    version = "v1.12.0",
)

go_repository(
    name = "com_google_cloud_go_dataqna",
    importpath = "cloud.google.com/go/dataqna",
    sum = "h1:yFzi/YU4YAdjyo7pXkBE2FeHbgz5OQQBVDdbErEHmVQ=",
    version = "v0.7.0",
)

go_repository(
    name = "com_google_cloud_go_datastore",
    importpath = "cloud.google.com/go/datastore",
    sum = "h1:iF6I/HaLs3Ado8uRKMvZRvF/ZLkWaWE9i8AiHzbC774=",
    version = "v1.11.0",
)

go_repository(
    name = "com_google_cloud_go_datastream",
    importpath = "cloud.google.com/go/datastream",
    sum = "h1:BBCBTnWMDwwEzQQmipUXxATa7Cm7CA/gKjKcR2w35T0=",
    version = "v1.7.0",
)

go_repository(
    name = "com_google_cloud_go_deploy",
    importpath = "cloud.google.com/go/deploy",
    sum = "h1:otshdKEbmsi1ELYeCKNYppwV0UH5xD05drSdBm7ouTk=",
    version = "v1.8.0",
)

go_repository(
    name = "com_google_cloud_go_dialogflow",
    importpath = "cloud.google.com/go/dialogflow",
    sum = "h1:uVlKKzp6G/VtSW0E7IH1Y5o0H48/UOCmqksG2riYCwQ=",
    version = "v1.32.0",
)

go_repository(
    name = "com_google_cloud_go_dlp",
    importpath = "cloud.google.com/go/dlp",
    sum = "h1:1JoJqezlgu6NWCroBxr4rOZnwNFILXr4cB9dMaSKO4A=",
    version = "v1.9.0",
)

go_repository(
    name = "com_google_cloud_go_documentai",
    importpath = "cloud.google.com/go/documentai",
    sum = "h1:KM3Xh0QQyyEdC8Gs2vhZfU+rt6OCPF0dwVwxKgLmWfI=",
    version = "v1.18.0",
)

go_repository(
    name = "com_google_cloud_go_domains",
    importpath = "cloud.google.com/go/domains",
    sum = "h1:2ti/o9tlWL4N+wIuWUNH+LbfgpwxPr8J1sv9RHA4bYQ=",
    version = "v0.8.0",
)

go_repository(
    name = "com_google_cloud_go_edgecontainer",
    importpath = "cloud.google.com/go/edgecontainer",
    sum = "h1:O0YVE5v+O0Q/ODXYsQHmHb+sYM8KNjGZw2pjX2Ws41c=",
    version = "v1.0.0",
)

go_repository(
    name = "com_google_cloud_go_errorreporting",
    importpath = "cloud.google.com/go/errorreporting",
    sum = "h1:kj1XEWMu8P0qlLhm3FwcaFsUvXChV/OraZwA70trRR0=",
    version = "v0.3.0",
)

go_repository(
    name = "com_google_cloud_go_essentialcontacts",
    importpath = "cloud.google.com/go/essentialcontacts",
    sum = "h1:gIzEhCoOT7bi+6QZqZIzX1Erj4SswMPIteNvYVlu+pM=",
    version = "v1.5.0",
)

go_repository(
    name = "com_google_cloud_go_eventarc",
    importpath = "cloud.google.com/go/eventarc",
    sum = "h1:fsJmNeqvqtk74FsaVDU6cH79lyZNCYP8Rrv7EhaB/PU=",
    version = "v1.11.0",
)

go_repository(
    name = "com_google_cloud_go_filestore",
    importpath = "cloud.google.com/go/filestore",
    sum = "h1:ckTEXN5towyTMu4q0uQ1Mde/JwTHur0gXs8oaIZnKfw=",
    version = "v1.6.0",
)

go_repository(
    name = "com_google_cloud_go_firestore",
    importpath = "cloud.google.com/go/firestore",
    sum = "h1:IBlRyxgGySXu5VuW0RgGFlTtLukSnNkpDiEOMkQkmpA=",
    version = "v1.9.0",
)

go_repository(
    name = "com_google_cloud_go_functions",
    importpath = "cloud.google.com/go/functions",
    sum = "h1:pPDqtsXG2g9HeOQLoquLbmvmb82Y4Ezdo1GXuotFoWg=",
    version = "v1.13.0",
)

go_repository(
    name = "com_google_cloud_go_gaming",
    importpath = "cloud.google.com/go/gaming",
    sum = "h1:7vEhFnZmd931Mo7sZ6pJy7uQPDxF7m7v8xtBheG08tc=",
    version = "v1.9.0",
)

go_repository(
    name = "com_google_cloud_go_gkebackup",
    importpath = "cloud.google.com/go/gkebackup",
    sum = "h1:za3QZvw6ujR0uyqkhomKKKNoXDyqYGPJies3voUK8DA=",
    version = "v0.4.0",
)

go_repository(
    name = "com_google_cloud_go_gkeconnect",
    importpath = "cloud.google.com/go/gkeconnect",
    sum = "h1:gXYKciHS/Lgq0GJ5Kc9SzPA35NGc3yqu6SkjonpEr2Q=",
    version = "v0.7.0",
)

go_repository(
    name = "com_google_cloud_go_gkehub",
    importpath = "cloud.google.com/go/gkehub",
    sum = "h1:TqCSPsEBQ6oZSJgEYZ3XT8x2gUadbvfwI32YB0kuHCs=",
    version = "v0.12.0",
)

go_repository(
    name = "com_google_cloud_go_gkemulticloud",
    importpath = "cloud.google.com/go/gkemulticloud",
    sum = "h1:8I84Q4vl02rJRsFiinBxl7WCozfdLlUVBQuSrqr9Wtk=",
    version = "v0.5.0",
)

go_repository(
    name = "com_google_cloud_go_gsuiteaddons",
    importpath = "cloud.google.com/go/gsuiteaddons",
    sum = "h1:1mvhXqJzV0Vg5Fa95QwckljODJJfDFXV4pn+iL50zzA=",
    version = "v1.5.0",
)

go_repository(
    name = "com_google_cloud_go_iam",
    importpath = "cloud.google.com/go/iam",
    sum = "h1:+CmB+K0J/33d0zSQ9SlFWUeCCEn5XJA0ZMZ3pHE9u8k=",
    version = "v0.13.0",
)

go_repository(
    name = "com_google_cloud_go_iap",
    importpath = "cloud.google.com/go/iap",
    sum = "h1:PxVHFuMxmSZyfntKXHXhd8bo82WJ+LcATenq7HLdVnU=",
    version = "v1.7.1",
)

go_repository(
    name = "com_google_cloud_go_ids",
    importpath = "cloud.google.com/go/ids",
    sum = "h1:fodnCDtOXuMmS8LTC2y3h8t24U8F3eKWfhi+3LY6Qf0=",
    version = "v1.3.0",
)

go_repository(
    name = "com_google_cloud_go_iot",
    importpath = "cloud.google.com/go/iot",
    sum = "h1:39W5BFSarRNZfVG0eXI5LYux+OVQT8GkgpHCnrZL2vM=",
    version = "v1.6.0",
)

go_repository(
    name = "com_google_cloud_go_kms",
    importpath = "cloud.google.com/go/kms",
    sum = "h1:7hm1bRqGCA1GBRQUrp831TwJ9TWhP+tvLuP497CQS2g=",
    version = "v1.10.1",
)

go_repository(
    name = "com_google_cloud_go_language",
    importpath = "cloud.google.com/go/language",
    sum = "h1:7Ulo2mDk9huBoBi8zCE3ONOoBrL6UXfAI71CLQ9GEIM=",
    version = "v1.9.0",
)

go_repository(
    name = "com_google_cloud_go_lifesciences",
    importpath = "cloud.google.com/go/lifesciences",
    sum = "h1:uWrMjWTsGjLZpCTWEAzYvyXj+7fhiZST45u9AgasasI=",
    version = "v0.8.0",
)

go_repository(
    name = "com_google_cloud_go_logging",
    importpath = "cloud.google.com/go/logging",
    sum = "h1:CJYxlNNNNAMkHp9em/YEXcfJg+rPDg7YfwoRpMU+t5I=",
    version = "v1.7.0",
)

go_repository(
    name = "com_google_cloud_go_longrunning",
    importpath = "cloud.google.com/go/longrunning",
    sum = "h1:v+yFJOfKC3yZdY6ZUI933pIYdhyhV8S3NpWrXWmg7jM=",
    version = "v0.4.1",
)

go_repository(
    name = "com_google_cloud_go_managedidentities",
    importpath = "cloud.google.com/go/managedidentities",
    sum = "h1:ZRQ4k21/jAhrHBVKl/AY7SjgzeJwG1iZa+mJ82P+VNg=",
    version = "v1.5.0",
)

go_repository(
    name = "com_google_cloud_go_maps",
    importpath = "cloud.google.com/go/maps",
    sum = "h1:mv9YaczD4oZBZkM5XJl6fXQ984IkJNHPwkc8MUsdkBo=",
    version = "v0.7.0",
)

go_repository(
    name = "com_google_cloud_go_mediatranslation",
    importpath = "cloud.google.com/go/mediatranslation",
    sum = "h1:anPxH+/WWt8Yc3EdoEJhPMBRF7EhIdz426A+tuoA0OU=",
    version = "v0.7.0",
)

go_repository(
    name = "com_google_cloud_go_memcache",
    importpath = "cloud.google.com/go/memcache",
    sum = "h1:8/VEmWCpnETCrBwS3z4MhT+tIdKgR1Z4Tr2tvYH32rg=",
    version = "v1.9.0",
)

go_repository(
    name = "com_google_cloud_go_metastore",
    importpath = "cloud.google.com/go/metastore",
    sum = "h1:QCFhZVe2289KDBQ7WxaHV2rAmPrmRAdLC6gbjUd3HPo=",
    version = "v1.10.0",
)

go_repository(
    name = "com_google_cloud_go_monitoring",
    importpath = "cloud.google.com/go/monitoring",
    sum = "h1:2qsrgXGVoRXpP7otZ14eE1I568zAa92sJSDPyOJvwjM=",
    version = "v1.13.0",
)

go_repository(
    name = "com_google_cloud_go_networkconnectivity",
    importpath = "cloud.google.com/go/networkconnectivity",
    sum = "h1:ZD6b4Pk1jEtp/cx9nx0ZYcL3BKqDa+KixNDZ6Bjs1B8=",
    version = "v1.11.0",
)

go_repository(
    name = "com_google_cloud_go_networkmanagement",
    importpath = "cloud.google.com/go/networkmanagement",
    sum = "h1:8KWEUNGcpSX9WwZXq7FtciuNGPdPdPN/ruDm769yAEM=",
    version = "v1.6.0",
)

go_repository(
    name = "com_google_cloud_go_networksecurity",
    importpath = "cloud.google.com/go/networksecurity",
    sum = "h1:sOc42Ig1K2LiKlzG71GUVloeSJ0J3mffEBYmvu+P0eo=",
    version = "v0.8.0",
)

go_repository(
    name = "com_google_cloud_go_notebooks",
    importpath = "cloud.google.com/go/notebooks",
    sum = "h1:Kg2K3K7CbSXYJHZ1aGQpf1xi5x2GUvQWf2sFVuiZh8M=",
    version = "v1.8.0",
)

go_repository(
    name = "com_google_cloud_go_optimization",
    importpath = "cloud.google.com/go/optimization",
    sum = "h1:dj8O4VOJRB4CUwZXdmwNViH1OtI0WtWL867/lnYH248=",
    version = "v1.3.1",
)

go_repository(
    name = "com_google_cloud_go_orchestration",
    importpath = "cloud.google.com/go/orchestration",
    sum = "h1:Vw+CEXo8M/FZ1rb4EjcLv0gJqqw89b7+g+C/EmniTb8=",
    version = "v1.6.0",
)

go_repository(
    name = "com_google_cloud_go_orgpolicy",
    importpath = "cloud.google.com/go/orgpolicy",
    sum = "h1:XDriMWug7sd0kYT1QKofRpRHzjad0bK8Q8uA9q+XrU4=",
    version = "v1.10.0",
)

go_repository(
    name = "com_google_cloud_go_osconfig",
    importpath = "cloud.google.com/go/osconfig",
    sum = "h1:PkSQx4OHit5xz2bNyr11KGcaFccL5oqglFPdTboyqwQ=",
    version = "v1.11.0",
)

go_repository(
    name = "com_google_cloud_go_oslogin",
    importpath = "cloud.google.com/go/oslogin",
    sum = "h1:whP7vhpmc+ufZa90eVpkfbgzJRK/Xomjz+XCD4aGwWw=",
    version = "v1.9.0",
)

go_repository(
    name = "com_google_cloud_go_phishingprotection",
    importpath = "cloud.google.com/go/phishingprotection",
    sum = "h1:l6tDkT7qAEV49MNEJkEJTB6vOO/onbSOcNtAT09HPuA=",
    version = "v0.7.0",
)

go_repository(
    name = "com_google_cloud_go_policytroubleshooter",
    importpath = "cloud.google.com/go/policytroubleshooter",
    sum = "h1:yKAGC4p9O61ttZUswaq9GAn1SZnEzTd0vUYXD7ZBT7Y=",
    version = "v1.6.0",
)

go_repository(
    name = "com_google_cloud_go_privatecatalog",
    importpath = "cloud.google.com/go/privatecatalog",
    sum = "h1:EPEJ1DpEGXLDnmc7mnCAqFmkwUJbIsaLAiLHVOkkwtc=",
    version = "v0.8.0",
)

go_repository(
    name = "com_google_cloud_go_pubsub",
    importpath = "cloud.google.com/go/pubsub",
    sum = "h1:vCge8m7aUKBJYOgrZp7EsNDf6QMd2CAlXZqWTn3yq6s=",
    version = "v1.30.0",
)

go_repository(
    name = "com_google_cloud_go_pubsublite",
    importpath = "cloud.google.com/go/pubsublite",
    sum = "h1:cb9fsrtpINtETHiJ3ECeaVzrfIVhcGjhhJEjybHXHao=",
    version = "v1.7.0",
)

go_repository(
    name = "com_google_cloud_go_recaptchaenterprise_v2",
    importpath = "cloud.google.com/go/recaptchaenterprise/v2",
    sum = "h1:6iOCujSNJ0YS7oNymI64hXsjGq60T4FK1zdLugxbzvU=",
    version = "v2.7.0",
)

go_repository(
    name = "com_google_cloud_go_recommendationengine",
    importpath = "cloud.google.com/go/recommendationengine",
    sum = "h1:VibRFCwWXrFebEWKHfZAt2kta6pS7Tlimsnms0fjv7k=",
    version = "v0.7.0",
)

go_repository(
    name = "com_google_cloud_go_recommender",
    importpath = "cloud.google.com/go/recommender",
    sum = "h1:ZnFRY5R6zOVk2IDS1Jbv5Bw+DExCI5rFumsTnMXiu/A=",
    version = "v1.9.0",
)

go_repository(
    name = "com_google_cloud_go_redis",
    importpath = "cloud.google.com/go/redis",
    sum = "h1:JoAd3SkeDt3rLFAAxEvw6wV4t+8y4ZzfZcZmddqphQ8=",
    version = "v1.11.0",
)

go_repository(
    name = "com_google_cloud_go_resourcemanager",
    importpath = "cloud.google.com/go/resourcemanager",
    sum = "h1:NRM0p+RJkaQF9Ee9JMnUV9BQ2QBIOq/v8M+Pbv/wmCs=",
    version = "v1.7.0",
)

go_repository(
    name = "com_google_cloud_go_resourcesettings",
    importpath = "cloud.google.com/go/resourcesettings",
    sum = "h1:8Dua37kQt27CCWHm4h/Q1XqCF6ByD7Ouu49xg95qJzI=",
    version = "v1.5.0",
)

go_repository(
    name = "com_google_cloud_go_retail",
    importpath = "cloud.google.com/go/retail",
    sum = "h1:1Dda2OpFNzIb4qWgFZjYlpP7sxX3aLeypKG6A3H4Yys=",
    version = "v1.12.0",
)

go_repository(
    name = "com_google_cloud_go_run",
    importpath = "cloud.google.com/go/run",
    sum = "h1:ydJQo+k+MShYnBfhaRHSZYeD/SQKZzZLAROyfpeD9zw=",
    version = "v0.9.0",
)

go_repository(
    name = "com_google_cloud_go_scheduler",
    importpath = "cloud.google.com/go/scheduler",
    sum = "h1:NpQAHtx3sulByTLe2dMwWmah8PWgeoieFPpJpArwFV0=",
    version = "v1.9.0",
)

go_repository(
    name = "com_google_cloud_go_secretmanager",
    importpath = "cloud.google.com/go/secretmanager",
    sum = "h1:pu03bha7ukxF8otyPKTFdDz+rr9sE3YauS5PliDXK60=",
    version = "v1.10.0",
)

go_repository(
    name = "com_google_cloud_go_security",
    importpath = "cloud.google.com/go/security",
    sum = "h1:PYvDxopRQBfYAXKAuDpFCKBvDOWPWzp9k/H5nB3ud3o=",
    version = "v1.13.0",
)

go_repository(
    name = "com_google_cloud_go_securitycenter",
    importpath = "cloud.google.com/go/securitycenter",
    sum = "h1:AF3c2s3awNTMoBtMX3oCUoOMmGlYxGOeuXSYHNBkf14=",
    version = "v1.19.0",
)

go_repository(
    name = "com_google_cloud_go_servicedirectory",
    importpath = "cloud.google.com/go/servicedirectory",
    sum = "h1:SJwk0XX2e26o25ObYUORXx6torSFiYgsGkWSkZgkoSU=",
    version = "v1.9.0",
)

go_repository(
    name = "com_google_cloud_go_shell",
    importpath = "cloud.google.com/go/shell",
    sum = "h1:wT0Uw7ib7+AgZST9eCDygwTJn4+bHMDtZo5fh7kGWDU=",
    version = "v1.6.0",
)

go_repository(
    name = "com_google_cloud_go_spanner",
    importpath = "cloud.google.com/go/spanner",
    sum = "h1:7VdjZ8zj4sHbDw55atp5dfY6kn1j9sam9DRNpPQhqR4=",
    version = "v1.45.0",
)

go_repository(
    name = "com_google_cloud_go_speech",
    importpath = "cloud.google.com/go/speech",
    sum = "h1:JEVoWGNnTF128kNty7T4aG4eqv2z86yiMJPT9Zjp+iw=",
    version = "v1.15.0",
)

go_repository(
    name = "com_google_cloud_go_storage",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/storage",
    patch_args = ["-p1"],
    patches = [
        "//thirdparty:com_google_cloud_go_storage-go-googleapis.patch",
    ],
    sum = "h1:uOdMxAs8HExqBlnLtnQyP0YkvbiDpdGShGKtx6U/oNM=",
    version = "v1.30.1",
)

go_repository(
    name = "com_google_cloud_go_storagetransfer",
    importpath = "cloud.google.com/go/storagetransfer",
    sum = "h1:5T+PM+3ECU3EY2y9Brv0Sf3oka8pKmsCfpQ07+91G9o=",
    version = "v1.8.0",
)

go_repository(
    name = "com_google_cloud_go_talent",
    importpath = "cloud.google.com/go/talent",
    sum = "h1:nI9sVZPjMKiO2q3Uu0KhTDVov3Xrlpt63fghP9XjyEM=",
    version = "v1.5.0",
)

go_repository(
    name = "com_google_cloud_go_texttospeech",
    importpath = "cloud.google.com/go/texttospeech",
    sum = "h1:H4g1ULStsbVtalbZGktyzXzw6jP26RjVGYx9RaYjBzc=",
    version = "v1.6.0",
)

go_repository(
    name = "com_google_cloud_go_tpu",
    importpath = "cloud.google.com/go/tpu",
    sum = "h1:/34T6CbSi+kTv5E19Q9zbU/ix8IviInZpzwz3rsFE+A=",
    version = "v1.5.0",
)

go_repository(
    name = "com_google_cloud_go_trace",
    importpath = "cloud.google.com/go/trace",
    sum = "h1:olxC0QHC59zgJVALtgqfD9tGk0lfeCP5/AGXL3Px/no=",
    version = "v1.9.0",
)

go_repository(
    name = "com_google_cloud_go_translate",
    importpath = "cloud.google.com/go/translate",
    sum = "h1:GvLP4oQ4uPdChBmBaUSa/SaZxCdyWELtlAaKzpHsXdA=",
    version = "v1.7.0",
)

go_repository(
    name = "com_google_cloud_go_video",
    importpath = "cloud.google.com/go/video",
    sum = "h1:upIbnGI0ZgACm58HPjAeBMleW3sl5cT84AbYQ8PWOgM=",
    version = "v1.15.0",
)

go_repository(
    name = "com_google_cloud_go_videointelligence",
    importpath = "cloud.google.com/go/videointelligence",
    sum = "h1:Uh5BdoET8XXqXX2uXIahGb+wTKbLkGH7s4GXR58RrG8=",
    version = "v1.10.0",
)

go_repository(
    name = "com_google_cloud_go_vision_v2",
    importpath = "cloud.google.com/go/vision/v2",
    sum = "h1:8C8RXUJoflCI4yVdqhTy9tRyygSHmp60aP363z23HKg=",
    version = "v2.7.0",
)

go_repository(
    name = "com_google_cloud_go_vmmigration",
    importpath = "cloud.google.com/go/vmmigration",
    sum = "h1:Azs5WKtfOC8pxvkyrDvt7J0/4DYBch0cVbuFfCCFt5k=",
    version = "v1.6.0",
)

go_repository(
    name = "com_google_cloud_go_vmwareengine",
    importpath = "cloud.google.com/go/vmwareengine",
    sum = "h1:b0NBu7S294l0gmtrT0nOJneMYgZapr5x9tVWvgDoVEM=",
    version = "v0.3.0",
)

go_repository(
    name = "com_google_cloud_go_vpcaccess",
    importpath = "cloud.google.com/go/vpcaccess",
    sum = "h1:FOe6CuiQD3BhHJWt7E8QlbBcaIzVRddupwJlp7eqmn4=",
    version = "v1.6.0",
)

go_repository(
    name = "com_google_cloud_go_webrisk",
    importpath = "cloud.google.com/go/webrisk",
    sum = "h1:IY+L2+UwxcVm2zayMAtBhZleecdIFLiC+QJMzgb0kT0=",
    version = "v1.8.0",
)

go_repository(
    name = "com_google_cloud_go_websecurityscanner",
    importpath = "cloud.google.com/go/websecurityscanner",
    sum = "h1:AHC1xmaNMOZtNqxI9Rmm87IJEyPaRkOxeI0gpAacXGk=",
    version = "v1.5.0",
)

go_repository(
    name = "com_google_cloud_go_workflows",
    importpath = "cloud.google.com/go/workflows",
    sum = "h1:FfGp9w0cYnaKZJhUOMqCOJCYT/WlvYBfTQhFWV3sRKI=",
    version = "v1.10.0",
)

go_repository(
    name = "com_lukechampine_uint128",
    importpath = "lukechampine.com/uint128",
    sum = "h1:mBi/5l91vocEN8otkC5bDLhi2KdCticRiwbdB0O+rjI=",
    version = "v1.2.0",
)

go_repository(
    name = "com_shuralyov_dmitri_gpu_mtl",
    importpath = "dmitri.shuralyov.com/gpu/mtl",
    sum = "h1:VpgP7xuJadIUuKccphEpTJnWhS2jkQyMt6Y7pJCD7fY=",
    version = "v0.0.0-20190408044501-666a987793e9",
)

go_repository(
    name = "in_gopkg_alecthomas_kingpin_v2",
    importpath = "gopkg.in/alecthomas/kingpin.v2",
    sum = "h1:jMFz6MfLP0/4fUyZle81rXUoxOBFi19VUFKVDOQfozc=",
    version = "v2.2.6",
)

go_repository(
    name = "in_gopkg_errgo_v2",
    importpath = "gopkg.in/errgo.v2",
    sum = "h1:0vLT13EuvQ0hNvakwLuFZ/jYrLp5F3kcWHXdRggjCE8=",
    version = "v2.1.0",
)

go_repository(
    name = "in_gopkg_inf_v0",
    importpath = "gopkg.in/inf.v0",
    sum = "h1:73M5CoZyi3ZLMOyDlQh031Cx6N9NDJ2Vvfl76EDAgDc=",
    version = "v0.9.1",
)

go_repository(
    name = "in_gopkg_square_go_jose_v2",
    importpath = "gopkg.in/square/go-jose.v2",
    sum = "h1:NGk74WTnPKBNUhNzQX7PYcTLUjoq7mzKk2OKbvwk2iI=",
    version = "v2.6.0",
)

go_repository(
    name = "in_gopkg_validator_v2",
    importpath = "gopkg.in/validator.v2",
    sum = "h1:xF0KWyGWXm/LM2G1TrEjqOu4pa6coO9AlWSf3msVfDY=",
    version = "v2.0.1",
)

go_repository(
    name = "in_gopkg_yaml_v2",
    importpath = "gopkg.in/yaml.v2",
    sum = "h1:D8xgwECY7CYvx+Y2n4sBz93Jn9JRvxdiyyo8CTfuKaY=",
    version = "v2.4.0",
)

go_repository(
    name = "io_opencensus_go",
    importpath = "go.opencensus.io",
    sum = "h1:y73uSU6J157QMP2kn2r30vwW1A2W2WFwSCGnAVxeaD0=",
    version = "v0.24.0",
)

go_repository(
    name = "io_opentelemetry_go_contrib_instrumentation_google_golang_org_grpc_otelgrpc",
    importpath = "go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc",
    sum = "h1:ZOLJc06r4CB42laIXg/7udr0pbZyuAihN10A/XuiQRY=",
    version = "v0.42.0",
)

go_repository(
    name = "io_opentelemetry_go_otel",
    importpath = "go.opentelemetry.io/otel",
    sum = "h1:Z7GVAX/UkAXPKsy94IU+i6thsQS4nb7LviLpnaNeW8s=",
    version = "v1.16.0",
)

go_repository(
    name = "io_opentelemetry_go_otel_exporters_otlp_internal_retry",
    importpath = "go.opentelemetry.io/otel/exporters/otlp/internal/retry",
    sum = "h1:t4ZwRPU+emrcvM2e9DHd0Fsf0JTPVcbfa/BhTDF03d0=",
    version = "v1.16.0",
)

go_repository(
    name = "io_opentelemetry_go_otel_exporters_otlp_otlpmetric",
    importpath = "go.opentelemetry.io/otel/exporters/otlp/otlpmetric",
    sum = "h1:f6BwB2OACc3FCbYVznctQ9V6KK7Vq6CjmYXJ7DeSs4E=",
    version = "v0.39.0",
)

go_repository(
    name = "io_opentelemetry_go_otel_exporters_otlp_otlpmetric_otlpmetricgrpc",
    importpath = "go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc",
    sum = "h1:rm+Fizi7lTM2UefJ1TO347fSRcwmIsUAaZmYmIGBRAo=",
    version = "v0.39.0",
)

go_repository(
    name = "io_opentelemetry_go_otel_exporters_otlp_otlptrace",
    importpath = "go.opentelemetry.io/otel/exporters/otlp/otlptrace",
    sum = "h1:cbsD4cUcviQGXdw8+bo5x2wazq10SKz8hEbtCRPcU78=",
    version = "v1.16.0",
)

go_repository(
    name = "io_opentelemetry_go_otel_exporters_otlp_otlptrace_otlptracegrpc",
    importpath = "go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc",
    sum = "h1:TVQp/bboR4mhZSav+MdgXB8FaRho1RC8UwVn3T0vjVc=",
    version = "v1.16.0",
)

go_repository(
    name = "io_opentelemetry_go_otel_exporters_prometheus",
    importpath = "go.opentelemetry.io/otel/exporters/prometheus",
    sum = "h1:whAaiHxOatgtKd+w0dOi//1KUxj3KoPINZdtDaDj3IA=",
    version = "v0.39.0",
)

go_repository(
    name = "io_opentelemetry_go_otel_metric",
    importpath = "go.opentelemetry.io/otel/metric",
    sum = "h1:RbrpwVG1Hfv85LgnZ7+txXioPDoh6EdbZHo26Q3hqOo=",
    version = "v1.16.0",
)

go_repository(
    name = "io_opentelemetry_go_otel_sdk",
    importpath = "go.opentelemetry.io/otel/sdk",
    sum = "h1:Z1Ok1YsijYL0CSJpHt4cS3wDDh7p572grzNrBMiMWgE=",
    version = "v1.16.0",
)

go_repository(
    name = "io_opentelemetry_go_otel_sdk_metric",
    importpath = "go.opentelemetry.io/otel/sdk/metric",
    sum = "h1:Kun8i1eYf48kHH83RucG93ffz0zGV1sh46FAScOTuDI=",
    version = "v0.39.0",
)

go_repository(
    name = "io_opentelemetry_go_otel_trace",
    importpath = "go.opentelemetry.io/otel/trace",
    sum = "h1:8JRpaObFoW0pxuVPapkgH8UhHQj+bJW8jJsCZEu5MQs=",
    version = "v1.16.0",
)

go_repository(
    name = "io_opentelemetry_go_proto_otlp",
    importpath = "go.opentelemetry.io/proto/otlp",
    sum = "h1:IVN6GR+mhC4s5yfcTbmzHYODqvWAp3ZedA2SJPI1Nnw=",
    version = "v0.19.0",
)

go_repository(
    name = "io_rsc_binaryregexp",
    importpath = "rsc.io/binaryregexp",
    sum = "h1:HfqmD5MEmC0zvwBuF187nq9mdnXjXsSivRiXN7SmRkE=",
    version = "v0.2.0",
)

go_repository(
    name = "io_rsc_pdf",
    importpath = "rsc.io/pdf",
    sum = "h1:k1MczvYDUvJBe93bYd7wrZLLUEcLZAuF824/I4e5Xr4=",
    version = "v0.1.1",
)

go_repository(
    name = "io_rsc_quote_v3",
    importpath = "rsc.io/quote/v3",
    sum = "h1:9JKUTTIUgS6kzR9mK1YuGKv6Nl+DijDNIc0ghT58FaY=",
    version = "v3.1.0",
)

go_repository(
    name = "io_rsc_sampler",
    importpath = "rsc.io/sampler",
    sum = "h1:7uVkIFmeBqHfdjD+gZwtXXI+RODJ2Wc4O7MPEh/QiW4=",
    version = "v1.3.0",
)

go_repository(
    name = "io_temporal_go_api",
    build_file_proto_mode = "disable_global",
    importpath = "go.temporal.io/api",
    sum = "h1:4y9mTQjEHsE0Du0WJ2ExJUcP/1/a+B/UefzIDm4ALTE=",
    version = "v1.23.0",
)

go_repository(
    name = "io_temporal_go_sdk",
    importpath = "go.temporal.io/sdk",
    sum = "h1:oa9/1f3bbcBLiNGbYf9woIx7uWFJ153q0JOkPeZqJtQ=",
    version = "v1.23.0",
)

go_repository(
    name = "io_temporal_go_server",
    build_file_proto_mode = "disable_global",
    importpath = "go.temporal.io/server",
    sum = "h1:MLkw3wdg2fmJZZWvP+nKH9Gtqm2GL3x7CTwdlcE0bIQ=",
    version = "v1.21.1",
)

go_repository(
    name = "io_temporal_go_version",
    importpath = "go.temporal.io/version",
    sum = "h1:dMrei9l9NyHt8nG6EB8vAwDLLTwx2SvRyucCSumAiig=",
    version = "v0.3.0",
)

go_repository(
    name = "org_golang_google_api",
    importpath = "google.golang.org/api",
    sum = "h1:1xQPji6cO2E2vLiI+C/XiFAnsn1WV3mjaEwGLhi3grE=",
    version = "v0.114.0",
)

go_repository(
    name = "org_golang_google_grpc_cmd_protoc_gen_go_grpc",
    importpath = "google.golang.org/grpc/cmd/protoc-gen-go-grpc",
    sum = "h1:M1YKkFIboKNieVO5DLUEVzQfGwJD30Nv2jfUgzb5UcE=",
    version = "v1.1.0",
)

go_repository(
    name = "org_golang_google_grpc_examples",
    importpath = "google.golang.org/grpc/examples",
    sum = "h1:0/wXWRnloOk74ilxLn9ZTrNSa0pMzJ0CnYyOe4gaf8U=",
    version = "v0.0.0-20230216223317-abff344ead8f",
)

go_repository(
    name = "org_golang_x_crypto",
    importpath = "golang.org/x/crypto",
    sum = "h1:LF6fAI+IutBocDJ2OT0Q1g8plpYljMZ4+lty+dsqw3g=",
    version = "v0.9.0",
)

go_repository(
    name = "org_golang_x_exp",
    importpath = "golang.org/x/exp",
    sum = "h1:pVgRXcIictcr+lBQIFeiwuwtDIs4eL21OuM9nyAADmo=",
    version = "v0.0.0-20230315142452-642cacee5cc0",
)

go_repository(
    name = "org_golang_x_image",
    importpath = "golang.org/x/image",
    sum = "h1:TcHcE0vrmgzNH1v3ppjcMGbhG5+9fMuvOmUYwNEF4q4=",
    version = "v0.0.0-20220302094943-723b81ca9867",
)

go_repository(
    name = "org_golang_x_lint",
    importpath = "golang.org/x/lint",
    sum = "h1:VLliZ0d+/avPrXXH+OakdXhpJuEoBZuwh1m2j7U6Iug=",
    version = "v0.0.0-20210508222113-6edffad5e616",
)

go_repository(
    name = "org_golang_x_mobile",
    importpath = "golang.org/x/mobile",
    sum = "h1:4+4C/Iv2U4fMZBiMCc98MG1In4gJY5YRhtpDNeDeHWs=",
    version = "v0.0.0-20190719004257-d2bd2a29d028",
)

go_repository(
    name = "org_golang_x_mod",
    importpath = "golang.org/x/mod",
    sum = "h1:lFO9qtOdlre5W1jxS3r/4szv2/6iXxScdzjoBMXNhYk=",
    version = "v0.10.0",
)

go_repository(
    name = "org_golang_x_sync",
    importpath = "golang.org/x/sync",
    sum = "h1:PUR+T4wwASmuSTYdKjYHI5TD22Wy5ogLU5qZCOLxBrI=",
    version = "v0.2.0",
)

go_repository(
    name = "org_golang_x_term",
    importpath = "golang.org/x/term",
    sum = "h1:n5xxQn2i3PC0yLAbjTpNT85q/Kgzcr2gIoX9OrJUols=",
    version = "v0.8.0",
)

go_repository(
    name = "org_golang_x_time",
    importpath = "golang.org/x/time",
    sum = "h1:rg5rLMjNzMS1RkNLzCG38eapWhnYLFYXDXj2gOlr8j4=",
    version = "v0.3.0",
)

go_repository(
    name = "org_golang_x_tools",
    importpath = "golang.org/x/tools",
    sum = "h1:8WMNJAz3zrtPmnYC7ISf5dEn3MT0gY7jBJfw27yrrLo=",
    version = "v0.9.1",
)

go_repository(
    name = "org_golang_x_xerrors",
    importpath = "golang.org/x/xerrors",
    sum = "h1:H2TDz8ibqkAF6YGhCdN3jS9O0/s90v0rJh3X/OLHEUk=",
    version = "v0.0.0-20220907171357-04be3eba64a2",
)

go_repository(
    name = "org_gonum_v1_gonum",
    importpath = "gonum.org/v1/gonum",
    sum = "h1:f1IJhK4Km5tBJmaiJXtk/PkL4cdVX6J+tGiM187uT5E=",
    version = "v0.11.0",
)

go_repository(
    name = "org_gonum_v1_netlib",
    importpath = "gonum.org/v1/netlib",
    sum = "h1:OE9mWmgKkjJyEmDAAtGMPjXu+YNeGvK9VTSHY6+Qihc=",
    version = "v0.0.0-20190313105609-8cb42192e0e0",
)

go_repository(
    name = "org_gonum_v1_plot",
    importpath = "gonum.org/v1/plot",
    sum = "h1:dnifSs43YJuNMDzB7v8wV64O4ABBHReuAVAoBxqBqS4=",
    version = "v0.10.1",
)

go_repository(
    name = "org_modernc_cc_v3",
    importpath = "modernc.org/cc/v3",
    sum = "h1:P3g79IUS/93SYhtoeaHW+kRCIrYaxJ27MFPv+7kaTOw=",
    version = "v3.40.0",
)

go_repository(
    name = "org_modernc_ccgo_v3",
    importpath = "modernc.org/ccgo/v3",
    sum = "h1:Mkgdzl46i5F/CNR/Kj80Ri59hC8TKAhZrYSaqvkwzUw=",
    version = "v3.16.13",
)

go_repository(
    name = "org_modernc_ccorpus",
    importpath = "modernc.org/ccorpus",
    sum = "h1:J16RXiiqiCgua6+ZvQot4yUuUy8zxgqbqEEUuGPlISk=",
    version = "v1.11.6",
)

go_repository(
    name = "org_modernc_httpfs",
    importpath = "modernc.org/httpfs",
    sum = "h1:AAgIpFZRXuYnkjftxTAZwMIiwEqAfk8aVB2/oA6nAeM=",
    version = "v1.0.6",
)

go_repository(
    name = "org_modernc_libc",
    importpath = "modernc.org/libc",
    sum = "h1:D/g6O5ftAfavceqlLOFwaZuA5KYafKwmr30A6iSqoyY=",
    version = "v1.22.3",
)

go_repository(
    name = "org_modernc_mathutil",
    importpath = "modernc.org/mathutil",
    sum = "h1:rV0Ko/6SfM+8G+yKiyI830l3Wuz1zRutdslNoQ0kfiQ=",
    version = "v1.5.0",
)

go_repository(
    name = "org_modernc_memory",
    importpath = "modernc.org/memory",
    sum = "h1:N+/8c5rE6EqugZwHii4IFsaJ7MUhoWX07J5tC/iI5Ds=",
    version = "v1.5.0",
)

go_repository(
    name = "org_modernc_opt",
    importpath = "modernc.org/opt",
    sum = "h1:3XOZf2yznlhC+ibLltsDGzABUGVx8J6pnFMS3E4dcq4=",
    version = "v0.1.3",
)

go_repository(
    name = "org_modernc_sqlite",
    importpath = "modernc.org/sqlite",
    sum = "h1:4aP4MdUf15i3R3M2mx6Q90WHKz3nZLoz96zlB6tNdow=",
    version = "v1.21.0",
)

go_repository(
    name = "org_modernc_strutil",
    importpath = "modernc.org/strutil",
    sum = "h1:fNMm+oJklMGYfU9Ylcywl0CO5O6nTfaowNsh2wpPjzY=",
    version = "v1.1.3",
)

go_repository(
    name = "org_modernc_tcl",
    importpath = "modernc.org/tcl",
    sum = "h1:mOQwiEK4p7HruMZcwKTZPw/aqtGM4aY00uzWhlKKYws=",
    version = "v1.15.1",
)

go_repository(
    name = "org_modernc_token",
    importpath = "modernc.org/token",
    sum = "h1:Xl7Ap9dKaEs5kLoOQeQmPWevfnk/DM5qcLcYlA8ys6Y=",
    version = "v1.1.0",
)

go_repository(
    name = "org_modernc_z",
    importpath = "modernc.org/z",
    sum = "h1:xkDw/KepgEjeizO2sNco+hqYkU12taxQFqPEmgm1GWE=",
    version = "v1.7.0",
)

go_repository(
    name = "org_uber_go_atomic",
    importpath = "go.uber.org/atomic",
    sum = "h1:9qC72Qh0+3MqyJbAn8YU5xVq1frD8bn3JtD2oXtafVQ=",
    version = "v1.10.0",
)

go_repository(
    name = "org_uber_go_dig",
    importpath = "go.uber.org/dig",
    sum = "h1:+alNIBsl0qfY0j6epRubp/9obgtrObRAc5aD+6jbWY8=",
    version = "v1.16.1",
)

go_repository(
    name = "org_uber_go_fx",
    importpath = "go.uber.org/fx",
    sum = "h1:JwYIYAQzXBuBBwSZ1/tn/95pnQO/Sp3yE8lWj9eSAzI=",
    version = "v1.19.1",
)

go_repository(
    name = "org_uber_go_goleak",
    importpath = "go.uber.org/goleak",
    sum = "h1:NBol2c7O1ZokfZ0LEU9K6Whx/KnwvepVetCUhtKja4A=",
    version = "v1.2.1",
)

go_repository(
    name = "org_uber_go_multierr",
    importpath = "go.uber.org/multierr",
    sum = "h1:7fIwc/ZtS0q++VgcfqFDxSBZVv/Xo49/SYnDFupUwlI=",
    version = "v1.9.0",
)

go_repository(
    name = "org_uber_go_tools",
    importpath = "go.uber.org/tools",
    sum = "h1:0mgffUl7nfd+FpvXMVz4IDEaUSmT1ysygQC7qYo7sG4=",
    version = "v0.0.0-20190618225709-2cfd321de3ee",
)

go_repository(
    name = "org_uber_go_zap",
    importpath = "go.uber.org/zap",
    sum = "h1:FiJd5l1UOLj0wCgbSE0rwwXHzEdAZS6hiiSnxJN/D60=",
    version = "v1.24.0",
)

go_repository(
    name = "com_github_99designs_go_keychain",
    importpath = "github.com/99designs/go-keychain",
    sum = "h1:/vQbFIOMbk2FiG/kXiLl8BRyzTWDw7gX/Hz7Dd5eDMs=",
    version = "v0.0.0-20191008050251-8e49817e8af4",
)

go_repository(
    name = "com_github_99designs_keyring",
    importpath = "github.com/99designs/keyring",
    sum = "h1:tYLp1ULvO7i3fI5vE21ReQuj99QFSs7lGm0xWyJo87o=",
    version = "v1.2.1",
)

go_repository(
    name = "com_github_andybalholm_brotli",
    importpath = "github.com/andybalholm/brotli",
    sum = "h1:V7DdXeJtZscaqfNuAdSRuRFzuiKlHSC/Zh3zl9qY3JY=",
    version = "v1.0.4",
)

go_repository(
    name = "com_github_apache_arrow_go_v10",
    importpath = "github.com/apache/arrow/go/v10",
    sum = "h1:n9dERvixoC/1JjDmBcs9FPaEryoANa2sCgVFo6ez9cI=",
    version = "v10.0.1",
)

go_repository(
    name = "com_github_aws_aws_sdk_go_v2",
    importpath = "github.com/aws/aws-sdk-go-v2",
    sum = "h1:M1fj4FE2lB4NzRb9Y0xdWsn2P0+2UHVxwKyOa4YJNjk=",
    version = "v1.16.16",
)

go_repository(
    name = "com_github_aws_aws_sdk_go_v2_aws_protocol_eventstream",
    importpath = "github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream",
    sum = "h1:tcFliCWne+zOuUfKNRn8JdFBuWPDuISDH08wD2ULkhk=",
    version = "v1.4.8",
)

go_repository(
    name = "com_github_aws_aws_sdk_go_v2_credentials",
    importpath = "github.com/aws/aws-sdk-go-v2/credentials",
    sum = "h1:9+ZhlDY7N9dPnUmf7CDfW9In4sW5Ff3bh7oy4DzS1IE=",
    version = "v1.12.20",
)

go_repository(
    name = "com_github_aws_aws_sdk_go_v2_feature_s3_manager",
    importpath = "github.com/aws/aws-sdk-go-v2/feature/s3/manager",
    sum = "h1:fAoVmNGhir6BR+RU0/EI+6+D7abM+MCwWf8v4ip5jNI=",
    version = "v1.11.33",
)

go_repository(
    name = "com_github_aws_aws_sdk_go_v2_internal_configsources",
    importpath = "github.com/aws/aws-sdk-go-v2/internal/configsources",
    sum = "h1:s4g/wnzMf+qepSNgTvaQQHNxyMLKSawNhKCPNy++2xY=",
    version = "v1.1.23",
)

go_repository(
    name = "com_github_aws_aws_sdk_go_v2_internal_endpoints_v2",
    importpath = "github.com/aws/aws-sdk-go-v2/internal/endpoints/v2",
    sum = "h1:/K482T5A3623WJgWT8w1yRAFK4RzGzEl7y39yhtn9eA=",
    version = "v2.4.17",
)

go_repository(
    name = "com_github_aws_aws_sdk_go_v2_internal_v4a",
    importpath = "github.com/aws/aws-sdk-go-v2/internal/v4a",
    sum = "h1:ZSIPAkAsCCjYrhqfw2+lNzWDzxzHXEckFkTePL5RSWQ=",
    version = "v1.0.14",
)

go_repository(
    name = "com_github_aws_aws_sdk_go_v2_service_internal_accept_encoding",
    importpath = "github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding",
    sum = "h1:Lh1AShsuIJTwMkoxVCAYPJgNG5H+eN6SmoUn8nOZ5wE=",
    version = "v1.9.9",
)

go_repository(
    name = "com_github_aws_aws_sdk_go_v2_service_internal_checksum",
    importpath = "github.com/aws/aws-sdk-go-v2/service/internal/checksum",
    sum = "h1:BBYoNQt2kUZUUK4bIPsKrCcjVPUMNsgQpNAwhznK/zo=",
    version = "v1.1.18",
)

go_repository(
    name = "com_github_aws_aws_sdk_go_v2_service_internal_presigned_url",
    importpath = "github.com/aws/aws-sdk-go-v2/service/internal/presigned-url",
    sum = "h1:Jrd/oMh0PKQc6+BowB+pLEwLIgaQF29eYbe7E1Av9Ug=",
    version = "v1.9.17",
)

go_repository(
    name = "com_github_aws_aws_sdk_go_v2_service_internal_s3shared",
    importpath = "github.com/aws/aws-sdk-go-v2/service/internal/s3shared",
    sum = "h1:HfVVR1vItaG6le+Bpw6P4midjBDMKnjMyZnw9MXYUcE=",
    version = "v1.13.17",
)

go_repository(
    name = "com_github_aws_aws_sdk_go_v2_service_s3",
    importpath = "github.com/aws/aws-sdk-go-v2/service/s3",
    sum = "h1:3/gm/JTX9bX8CpzTgIlrtYpB3EVBDxyg/GY/QdcIEZw=",
    version = "v1.27.11",
)

go_repository(
    name = "com_github_aws_smithy_go",
    importpath = "github.com/aws/smithy-go",
    sum = "h1:l7LYxGuzK6/K+NzJ2mC+VvLUbae0sL3bXU//04MkmnA=",
    version = "v1.13.3",
)

go_repository(
    name = "com_github_azure_azure_sdk_for_go_sdk_azcore",
    importpath = "github.com/Azure/azure-sdk-for-go/sdk/azcore",
    sum = "h1:rTnT/Jrcm+figWlYz4Ixzt0SJVR2cMC8lvZcimipiEY=",
    version = "v1.4.0",
)

go_repository(
    name = "com_github_azure_azure_sdk_for_go_sdk_internal",
    importpath = "github.com/Azure/azure-sdk-for-go/sdk/internal",
    sum = "h1:+5VZ72z0Qan5Bog5C+ZkgSqUbeVUd9wgtHOrIKuc5b8=",
    version = "v1.1.2",
)

go_repository(
    name = "com_github_azure_azure_sdk_for_go_sdk_storage_azblob",
    importpath = "github.com/Azure/azure-sdk-for-go/sdk/storage/azblob",
    sum = "h1:u/LLAOFgsMv7HmNL4Qufg58y+qElGOt5qv0z1mURkRY=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_azure_go_ansiterm",
    importpath = "github.com/Azure/go-ansiterm",
    sum = "h1:L/gRVlceqvL25UVaW/CKtUDjefjrs0SPonmDGUVOYP0=",
    version = "v0.0.0-20230124172434-306776ec8161",
)

go_repository(
    name = "com_github_azure_go_autorest",
    importpath = "github.com/Azure/go-autorest",
    sum = "h1:V5VMDjClD3GiElqLWO7mz2MxNAK/vTfRHdAubSIPRgs=",
    version = "v14.2.0+incompatible",
)

go_repository(
    name = "com_github_azure_go_autorest_autorest_adal",
    importpath = "github.com/Azure/go-autorest/autorest/adal",
    sum = "h1:P8An8Z9rH1ldbOLdFpxYorgOt2sywL9V24dAwWHPuGc=",
    version = "v0.9.16",
)

go_repository(
    name = "com_github_azure_go_autorest_autorest_date",
    importpath = "github.com/Azure/go-autorest/autorest/date",
    sum = "h1:7gUk1U5M/CQbp9WoqinNzJar+8KY+LPI6wiWrP/myHw=",
    version = "v0.3.0",
)

go_repository(
    name = "com_github_azure_go_autorest_logger",
    importpath = "github.com/Azure/go-autorest/logger",
    sum = "h1:IG7i4p/mDa2Ce4TRyAO8IHnVhAVF3RFU+ZtXWSmf4Tg=",
    version = "v0.2.1",
)

go_repository(
    name = "com_github_azure_go_autorest_tracing",
    importpath = "github.com/Azure/go-autorest/tracing",
    sum = "h1:TYi4+3m5t6K48TGI9AUdb+IzbnSxvnvUMfuitfgcfuo=",
    version = "v0.6.0",
)

go_repository(
    name = "com_github_clickhouse_clickhouse_go",
    importpath = "github.com/ClickHouse/clickhouse-go",
    sum = "h1:iAFMa2UrQdR5bHJ2/yaSLffZkxpcOYQMCUuKeNXGdqc=",
    version = "v1.4.3",
)

go_repository(
    name = "com_github_cloudflare_golz4",
    importpath = "github.com/cloudflare/golz4",
    sum = "h1:F1EaeKL/ta07PY/k9Os/UFtwERei2/XzGemhpGnBKNg=",
    version = "v0.0.0-20150217214814-ef862a3cdc58",
)

go_repository(
    name = "com_github_cockroachdb_cockroach_go_v2",
    importpath = "github.com/cockroachdb/cockroach-go/v2",
    sum = "h1:3XzfSMuUT0wBe1a3o5C0eOTcArhmmFAg2Jzh/7hhKqo=",
    version = "v2.1.1",
)

go_repository(
    name = "com_github_cznic_mathutil",
    importpath = "github.com/cznic/mathutil",
    sum = "h1:XNT/Zf5l++1Pyg08/HV04ppB0gKxAqtZQBRYiYrUuYk=",
    version = "v0.0.0-20180504122225-ca4c9f2c1369",
)

go_repository(
    name = "com_github_danieljoos_wincred",
    importpath = "github.com/danieljoos/wincred",
    sum = "h1:QLdCxFs1/Yl4zduvBdcHB8goaYk9RARS2SgLLRuAyr0=",
    version = "v1.1.2",
)

go_repository(
    name = "com_github_dhui_dktest",
    importpath = "github.com/dhui/dktest",
    sum = "h1:i6gq2YQEtcrjKbeJpBkWjE8MmLZPYllcjOFbTZuPDnw=",
    version = "v0.3.16",
)

go_repository(
    name = "com_github_docker_distribution",
    importpath = "github.com/docker/distribution",
    sum = "h1:T3de5rq0dB1j30rp0sA2rER+m322EBzniBPB6ZIzuh8=",
    version = "v2.8.2+incompatible",
)

go_repository(
    name = "com_github_docker_docker",
    importpath = "github.com/docker/docker",
    sum = "h1:Ugvxm7a8+Gz6vqQYQQ2W7GYq5EUPaAiuPgIfVyI3dYE=",
    version = "v20.10.24+incompatible",
)

go_repository(
    name = "com_github_docker_go_connections",
    importpath = "github.com/docker/go-connections",
    sum = "h1:El9xVISelRB7BuFusrZozjnkIM5YnzCViNKohAFqRJQ=",
    version = "v0.4.0",
)

go_repository(
    name = "com_github_docker_go_units",
    importpath = "github.com/docker/go-units",
    sum = "h1:69rxXcBk27SvSaaxTtLh/8llcHD8vYHT7WSdRZ/jvr4=",
    version = "v0.5.0",
)

go_repository(
    name = "com_github_dvsekhvalnov_jose2go",
    importpath = "github.com/dvsekhvalnov/jose2go",
    sum = "h1:3j8ya4Z4kMCwT5nXIKFSV84YS+HdqSSO0VsTQxaLAeM=",
    version = "v1.5.0",
)

go_repository(
    name = "com_github_edsrzf_mmap_go",
    importpath = "github.com/edsrzf/mmap-go",
    sum = "h1:aaQcKT9WumO6JEJcRyTqFVq4XUZiUcKR2/GI31TOcz8=",
    version = "v0.0.0-20170320065105-0bce6a688712",
)

go_repository(
    name = "com_github_form3tech_oss_jwt_go",
    importpath = "github.com/form3tech-oss/jwt-go",
    sum = "h1:/l4kBbb4/vGSsdtB5nUe8L7B9mImVMaBPw9L/0TBHU8=",
    version = "v3.2.5+incompatible",
)

go_repository(
    name = "com_github_fsouza_fake_gcs_server",
    importpath = "github.com/fsouza/fake-gcs-server",
    sum = "h1:OeH75kBZcZa3ZE+zz/mFdJ2btt9FgqfjI7gIh9+5fvk=",
    version = "v1.17.0",
)

go_repository(
    name = "com_github_gabriel_vasile_mimetype",
    importpath = "github.com/gabriel-vasile/mimetype",
    sum = "h1:TRWk7se+TOjCYgRth7+1/OYLNiRNIotknkFtf/dnN7Q=",
    version = "v1.4.1",
)

go_repository(
    name = "com_github_gobuffalo_here",
    importpath = "github.com/gobuffalo/here",
    sum = "h1:hYrd0a6gDmWxBM4TnrGw8mQg24iSVoIkHEk7FodQcBI=",
    version = "v0.6.0",
)

go_repository(
    name = "com_github_goccy_go_json",
    importpath = "github.com/goccy/go-json",
    sum = "h1:/pAaQDLHEoCq/5FFmSKBswWmK6H0e8g4159Kc/X/nqk=",
    version = "v0.9.11",
)

go_repository(
    name = "com_github_godbus_dbus",
    importpath = "github.com/godbus/dbus",
    sum = "h1:ZpnhV/YsD2/4cESfV5+Hoeu/iUR3ruzNvZ+yQfO03a0=",
    version = "v0.0.0-20190726142602-4481cbc300e2",
)

go_repository(
    name = "com_github_golang_migrate_migrate_v4",
    importpath = "github.com/golang-migrate/migrate/v4",
    sum = "h1:8coYbMKUyInrFk1lfGfRovTLAW7PhWp8qQDT2iKfuoA=",
    version = "v4.16.2",
)

go_repository(
    name = "com_github_golang_sql_civil",
    importpath = "github.com/golang-sql/civil",
    sum = "h1:lXe2qZdvpiX5WZkZR4hgp4KJVfY3nMkvmwbVkpv1rVY=",
    version = "v0.0.0-20190719163853-cb61b32ac6fe",
)

go_repository(
    name = "com_github_golang_sql_sqlexp",
    importpath = "github.com/golang-sql/sqlexp",
    sum = "h1:ZCD6MBpcuOVfGVqsEmY5/4FtYiKz6tSyUv9LPEDei6A=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_google_flatbuffers",
    importpath = "github.com/google/flatbuffers",
    sum = "h1:ivUb1cGomAB101ZM1T0nOiWz9pSrTMoa9+EiY7igmkM=",
    version = "v2.0.8+incompatible",
)

go_repository(
    name = "com_github_google_go_github_v39",
    importpath = "github.com/google/go-github/v39",
    sum = "h1:rNNM311XtPOz5rDdsJXAp2o8F67X9FnROXTvto3aSnQ=",
    version = "v39.2.0",
)

go_repository(
    name = "com_github_google_go_querystring",
    importpath = "github.com/google/go-querystring",
    sum = "h1:AnCroh3fv4ZBgVIf1Iwtovgjaw/GiKJo8M8yD/fhyJ8=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_gorilla_handlers",
    importpath = "github.com/gorilla/handlers",
    sum = "h1:0QniY0USkHQ1RGCLfKxeNHK9bkDHGRYGNDFBCS+YARg=",
    version = "v1.4.2",
)

go_repository(
    name = "com_github_gorilla_mux",
    importpath = "github.com/gorilla/mux",
    sum = "h1:VuZ8uybHlWmqV03+zRzdwKL4tUnIp1MAQtp1mIFE1bc=",
    version = "v1.7.4",
)

go_repository(
    name = "com_github_gsterjov_go_libsecret",
    importpath = "github.com/gsterjov/go-libsecret",
    sum = "h1:6rhixN/i8ZofjG1Y75iExal34USq5p+wiN1tpie8IrU=",
    version = "v0.0.0-20161001094733-a6f4afe4910c",
)

go_repository(
    name = "com_github_hashicorp_errwrap",
    importpath = "github.com/hashicorp/errwrap",
    sum = "h1:OxrOeh75EUXMY8TBjag2fzXGZ40LB6IKw45YeGUDY2I=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_hashicorp_go_multierror",
    importpath = "github.com/hashicorp/go-multierror",
    sum = "h1:H5DkEtf6CXdFp0N0Em5UCwQpXMWke8IA0+lD48awMYo=",
    version = "v1.1.1",
)

go_repository(
    name = "com_github_jackc_chunkreader_v2",
    importpath = "github.com/jackc/chunkreader/v2",
    sum = "h1:i+RDz65UE+mmpjTfyz0MoVTnzeYxroil2G82ki7MGG8=",
    version = "v2.0.1",
)

go_repository(
    name = "com_github_jackc_pgconn",
    importpath = "github.com/jackc/pgconn",
    sum = "h1:vrbA9Ud87g6JdFWkHTJXppVce58qPIdP7N8y0Ml/A7Q=",
    version = "v1.14.0",
)

go_repository(
    name = "com_github_jackc_pgerrcode",
    importpath = "github.com/jackc/pgerrcode",
    sum = "h1:s+4MhCQ6YrzisK6hFJUX53drDT4UsSW3DEhKn0ifuHw=",
    version = "v0.0.0-20220416144525-469b46aa5efa",
)

go_repository(
    name = "com_github_jackc_pgio",
    importpath = "github.com/jackc/pgio",
    sum = "h1:g12B9UwVnzGhueNavwioyEEpAmqMe1E/BN9ES+8ovkE=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_jackc_pgpassfile",
    importpath = "github.com/jackc/pgpassfile",
    sum = "h1:/6Hmqy13Ss2zCq62VdNG8tM1wchn8zjSGOBJ6icpsIM=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_jackc_pgproto3_v2",
    importpath = "github.com/jackc/pgproto3/v2",
    sum = "h1:7eY55bdBeCz1F2fTzSz69QC+pG46jYq9/jtSPiJ5nn0=",
    version = "v2.3.2",
)

go_repository(
    name = "com_github_jackc_pgservicefile",
    importpath = "github.com/jackc/pgservicefile",
    sum = "h1:bbPeKD0xmW/Y25WS6cokEszi5g+S0QxI/d45PkRi7Nk=",
    version = "v0.0.0-20221227161230-091c0ba34f0a",
)

go_repository(
    name = "com_github_jackc_pgtype",
    importpath = "github.com/jackc/pgtype",
    sum = "h1:y+xUdabmyMkJLyApYuPj38mW+aAIqCe5uuBB51rH3Vw=",
    version = "v1.14.0",
)

go_repository(
    name = "com_github_jackc_pgx_v4",
    importpath = "github.com/jackc/pgx/v4",
    sum = "h1:YP7G1KABtKpB5IHrO9vYwSrCOhs7p3uqhvhhQBptya0=",
    version = "v4.18.1",
)

go_repository(
    name = "com_github_jackc_pgx_v5",
    importpath = "github.com/jackc/pgx/v5",
    sum = "h1:Fcr8QJ1ZeLi5zsPZqQeUZhNhxfkkKBOgJuYkJHoBOtU=",
    version = "v5.3.1",
)

go_repository(
    name = "com_github_k0kubun_pp",
    importpath = "github.com/k0kubun/pp",
    sum = "h1:EKhKbi34VQDWJtq+zpsKSEhkHHs9w2P8Izbq8IhLVSo=",
    version = "v2.3.0+incompatible",
)

go_repository(
    name = "com_github_kardianos_osext",
    importpath = "github.com/kardianos/osext",
    sum = "h1:iQTw/8FWTuc7uiaSepXwyf3o52HaUYcV+Tu66S3F5GA=",
    version = "v0.0.0-20190222173326-2bc1f35cddc0",
)

go_repository(
    name = "com_github_klauspost_asmfmt",
    importpath = "github.com/klauspost/asmfmt",
    sum = "h1:4Ri7ox3EwapiOjCki+hw14RyKk201CN4rzyCJRFLpK4=",
    version = "v1.3.2",
)

go_repository(
    name = "com_github_klauspost_compress",
    importpath = "github.com/klauspost/compress",
    sum = "h1:Lcadnb3RKGin4FYM/orgq0qde+nc15E5Cbqg4B9Sx9c=",
    version = "v1.15.11",
)

go_repository(
    name = "com_github_klauspost_cpuid_v2",
    importpath = "github.com/klauspost/cpuid/v2",
    sum = "h1:sxCkb+qR91z4vsqw4vGGZlDgPz3G7gjaLyK3V8y70BU=",
    version = "v2.2.3",
)

go_repository(
    name = "com_github_ktrysmt_go_bitbucket",
    importpath = "github.com/ktrysmt/go-bitbucket",
    sum = "h1:C8dUGp0qkwncKtAnozHCbbqhptefzEd1I0sfnuy9rYQ=",
    version = "v0.6.4",
)

go_repository(
    name = "com_github_markbates_pkger",
    importpath = "github.com/markbates/pkger",
    sum = "h1:3MPelV53RnGSW07izx5xGxl4e/sdRD6zqseIk0rMASY=",
    version = "v0.15.1",
)

go_repository(
    name = "com_github_microsoft_go_mssqldb",
    importpath = "github.com/microsoft/go-mssqldb",
    sum = "h1:k2p2uuG8T5T/7Hp7/e3vMGTnnR0sU4h8d1CcC71iLHU=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_microsoft_go_winio",
    importpath = "github.com/Microsoft/go-winio",
    sum = "h1:9/kr64B9VUZrLm5YYwbGtUJnMgqWVOdUAXu6Migciow=",
    version = "v0.6.1",
)

go_repository(
    name = "com_github_minio_asm2plan9s",
    importpath = "github.com/minio/asm2plan9s",
    sum = "h1:AMFGa4R4MiIpspGNG7Z948v4n35fFGB3RR3G/ry4FWs=",
    version = "v0.0.0-20200509001527-cdd76441f9d8",
)

go_repository(
    name = "com_github_minio_c2goasm",
    importpath = "github.com/minio/c2goasm",
    sum = "h1:+n/aFZefKZp7spd8DFdX7uMikMLXX4oubIzJF4kv/wI=",
    version = "v0.0.0-20190812172519-36a3d3bbc4f3",
)

go_repository(
    name = "com_github_mitchellh_mapstructure",
    importpath = "github.com/mitchellh/mapstructure",
    sum = "h1:fmNYVwqnSfB9mZU6OS2O6GsXM+wcskZDuKQzvN1EDeE=",
    version = "v1.1.2",
)

go_repository(
    name = "com_github_moby_term",
    importpath = "github.com/moby/term",
    sum = "h1:xt8Q1nalod/v7BqbG21f8mQPqH+xAaC9C3N3wfWbVP0=",
    version = "v0.5.0",
)

go_repository(
    name = "com_github_morikuni_aec",
    importpath = "github.com/morikuni/aec",
    sum = "h1:nP9CBfwrvYnBRgY6qfDQkygYDmYwOilePFkwzv4dU8A=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_mtibben_percent",
    importpath = "github.com/mtibben/percent",
    sum = "h1:5gssi8Nqo8QU/r2pynCm+hBQHpkB/uNK7BJCFogWdzs=",
    version = "v0.2.1",
)

go_repository(
    name = "com_github_mutecomm_go_sqlcipher_v4",
    importpath = "github.com/mutecomm/go-sqlcipher/v4",
    sum = "h1:sV1tWCWGAVlPhNGT95Q+z/txFxuhAYWwHD1afF5bMZg=",
    version = "v4.4.0",
)

go_repository(
    name = "com_github_nakagami_firebirdsql",
    importpath = "github.com/nakagami/firebirdsql",
    sum = "h1:P48LjvUQpTReR3TQRbxSeSBsMXzfK0uol7eRcr7VBYQ=",
    version = "v0.0.0-20190310045651-3c02a58cfed8",
)

go_repository(
    name = "com_github_neo4j_neo4j_go_driver",
    importpath = "github.com/neo4j/neo4j-go-driver",
    sum = "h1:fhFP5RliM2HW/8XdcO5QngSfFli9GcRIpMXvypTQt6E=",
    version = "v1.8.1-0.20200803113522-b626aa943eba",
)

go_repository(
    name = "com_github_onsi_ginkgo",
    importpath = "github.com/onsi/ginkgo",
    sum = "h1:29JGrr5oVBm5ulCWet69zQkzWipVXIol6ygQUe/EzNc=",
    version = "v1.16.4",
)

go_repository(
    name = "com_github_onsi_gomega",
    importpath = "github.com/onsi/gomega",
    sum = "h1:WjP/FQ/sk43MRmnEcT+MlDw2TFvkrXlprrPST/IudjU=",
    version = "v1.15.0",
)

go_repository(
    name = "com_github_opencontainers_go_digest",
    importpath = "github.com/opencontainers/go-digest",
    sum = "h1:apOUWs51W5PlhuyGyz9FCeeBIOUDA/6nW8Oi/yOhh5U=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_opencontainers_image_spec",
    importpath = "github.com/opencontainers/image-spec",
    sum = "h1:9yCKha/T5XdGtO0q9Q9a6T5NUCsTn/DrBg0D7ufOcFM=",
    version = "v1.0.2",
)

go_repository(
    name = "com_github_pierrec_lz4_v4",
    importpath = "github.com/pierrec/lz4/v4",
    sum = "h1:kQPfno+wyx6C5572ABwV+Uo3pDFzQ7yhyGchSyRda0c=",
    version = "v4.1.16",
)

go_repository(
    name = "com_github_pkg_browser",
    importpath = "github.com/pkg/browser",
    sum = "h1:KoWmjvw+nsYOo29YJK9vDA65RGE3NrOnUtO7a+RF9HU=",
    version = "v0.0.0-20210911075715-681adbf594b8",
)

go_repository(
    name = "com_github_shopspring_decimal",
    importpath = "github.com/shopspring/decimal",
    sum = "h1:abSATXmQEYyShuxI4/vyW3tV1MrKAJzCZ/0zLUXYbsQ=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_snowflakedb_gosnowflake",
    importpath = "github.com/snowflakedb/gosnowflake",
    sum = "h1:KSHXrQ5o7uso25hNIzi/RObXtnSGkFgie91X82KcvMY=",
    version = "v1.6.19",
)

go_repository(
    name = "com_github_xanzy_go_gitlab",
    importpath = "github.com/xanzy/go-gitlab",
    sum = "h1:rWtwKTgEnXyNUGrOArN7yyc3THRkpYcKXIXia9abywQ=",
    version = "v0.15.0",
)

go_repository(
    name = "com_github_xdg_go_pbkdf2",
    importpath = "github.com/xdg-go/pbkdf2",
    sum = "h1:Su7DPu48wXMwC3bs7MCNG+z4FhcyEuz5dlvchbq0B0c=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_xdg_go_scram",
    importpath = "github.com/xdg-go/scram",
    sum = "h1:VOMT+81stJgXW3CpHyqHN3AXDYIMsx56mEFrB37Mb/E=",
    version = "v1.1.1",
)

go_repository(
    name = "com_github_xdg_go_stringprep",
    importpath = "github.com/xdg-go/stringprep",
    sum = "h1:kdwGpVNwPFtjs98xCGkHjQtGKh86rDcRZN17QEMCOIs=",
    version = "v1.0.3",
)

go_repository(
    name = "com_github_youmark_pkcs8",
    importpath = "github.com/youmark/pkcs8",
    sum = "h1:splanxYIlg+5LfHAM6xpdFEAYOk8iySO56hMFq6uLyA=",
    version = "v0.0.0-20181117223130-1be2e3e5546d",
)

go_repository(
    name = "com_github_zeebo_xxh3",
    importpath = "github.com/zeebo/xxh3",
    sum = "h1:xZmwmqxHZA8AI603jOQ0tMqmBr9lPeFwGg6d+xy9DC0=",
    version = "v1.0.2",
)

go_repository(
    name = "com_gitlab_nyarla_go_crypt",
    importpath = "gitlab.com/nyarla/go-crypt",
    sum = "h1:7gd+rd8P3bqcn/96gOZa3F5dpJr/vEiDQYlNb/y2uNs=",
    version = "v0.0.0-20160106005555-d9a5dc2b789b",
)

go_repository(
    name = "org_modernc_b",
    importpath = "modernc.org/b",
    sum = "h1:vpvqeyp17ddcQWF29Czawql4lDdABCDRbXRAS4+aF2o=",
    version = "v1.0.0",
)

go_repository(
    name = "org_modernc_db",
    importpath = "modernc.org/db",
    sum = "h1:2c6NdCfaLnshSvY7OU09cyAY0gYXUZj4lmg5ItHyucg=",
    version = "v1.0.0",
)

go_repository(
    name = "org_modernc_file",
    importpath = "modernc.org/file",
    sum = "h1:9/PdvjVxd5+LcWUQIfapAWRGOkDLK90rloa8s/au06A=",
    version = "v1.0.0",
)

go_repository(
    name = "org_modernc_fileutil",
    importpath = "modernc.org/fileutil",
    sum = "h1:Z1AFLZwl6BO8A5NldQg/xTSjGLetp+1Ubvl4alfGx8w=",
    version = "v1.0.0",
)

go_repository(
    name = "org_modernc_golex",
    importpath = "modernc.org/golex",
    sum = "h1:wWpDlbK8ejRfSyi0frMyhilD3JBvtcx2AdGDnU+JtsE=",
    version = "v1.0.0",
)

go_repository(
    name = "org_modernc_internal",
    importpath = "modernc.org/internal",
    sum = "h1:XMDsFDcBDsibbBnHB2xzljZ+B1yrOVLEFkKL2u15Glw=",
    version = "v1.0.0",
)

go_repository(
    name = "org_modernc_lldb",
    importpath = "modernc.org/lldb",
    sum = "h1:6vjDJxQEfhlOLwl4bhpwIz00uyFK4EmSYcbwqwbynsc=",
    version = "v1.0.0",
)

go_repository(
    name = "org_modernc_ql",
    importpath = "modernc.org/ql",
    sum = "h1:bIQ/trWNVjQPlinI6jdOQsi195SIturGo3mp5hsDqVU=",
    version = "v1.0.0",
)

go_repository(
    name = "org_modernc_sortutil",
    importpath = "modernc.org/sortutil",
    sum = "h1:oP3U4uM+NT/qBQcbg/K2iqAX0Nx7B1b6YZtq3Gk/PjM=",
    version = "v1.1.0",
)

go_repository(
    name = "org_modernc_zappy",
    importpath = "modernc.org/zappy",
    sum = "h1:dPVaP+3ueIUv4guk8PuZ2wiUGcJ1WUVvIheeSSTD0yk=",
    version = "v1.0.0",
)

go_repository(
    name = "org_mongodb_go_mongo_driver",
    importpath = "go.mongodb.org/mongo-driver",
    sum = "h1:ny3p0reEpgsR2cfA5cjgwFZg3Cv/ofFh/8jbhGtz9VI=",
    version = "v1.7.5",
)

go_repository(
    name = "com_github_ajstarks_deck",
    importpath = "github.com/ajstarks/deck",
    sum = "h1:7kQgkwGRoLzC9K0oyXdJo7nve/bynv/KwUsxbiTlzAM=",
    version = "v0.0.0-20200831202436-30c9fc6549a9",
)

go_repository(
    name = "com_github_ajstarks_deck_generate",
    importpath = "github.com/ajstarks/deck/generate",
    sum = "h1:iXUgAaqDcIUGbRoy2TdeofRG/j1zpGRSEmNK05T+bi8=",
    version = "v0.0.0-20210309230005-c3f852c02e19",
)

go_repository(
    name = "com_github_alecthomas_kingpin_v2",
    importpath = "github.com/alecthomas/kingpin/v2",
    sum = "h1:H0aULhgmSzN8xQ3nX1uxtdlTHYoPLu5AhHxWrKI6ocU=",
    version = "v2.3.2",
)

go_repository(
    name = "com_github_apache_arrow_go_v11",
    importpath = "github.com/apache/arrow/go/v11",
    sum = "h1:hqauxvFQxww+0mEU/2XHG6LT7eZternCZq+A5Yly2uM=",
    version = "v11.0.0",
)

go_repository(
    name = "com_github_boombuler_barcode",
    importpath = "github.com/boombuler/barcode",
    sum = "h1:NDBbPmhS+EqABEs5Kg3n/5ZNjy73Pz7SIV+KCeqyXcs=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_docopt_docopt_go",
    importpath = "github.com/docopt/docopt-go",
    sum = "h1:bWDMxwH3px2JBh6AyO7hdCn/PkvCZXii8TGj7sbtEbQ=",
    version = "v0.0.0-20180111231733-ee0de3bc6815",
)

go_repository(
    name = "com_github_go_fonts_dejavu",
    importpath = "github.com/go-fonts/dejavu",
    sum = "h1:JSajPXURYqpr+Cu8U9bt8K+XcACIHWqWrvWCKyeFmVQ=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_go_fonts_latin_modern",
    importpath = "github.com/go-fonts/latin-modern",
    sum = "h1:5/Tv1Ek/QCr20C6ZOz15vw3g7GELYL98KWr8Hgo+3vk=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_go_fonts_liberation",
    importpath = "github.com/go-fonts/liberation",
    sum = "h1:jAkAWJP4S+OsrPLZM4/eC9iW7CtHy+HBXrEwZXWo5VM=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_go_fonts_stix",
    importpath = "github.com/go-fonts/stix",
    sum = "h1:UlZlgrvvmT/58o573ot7NFw0vZasZ5I6bcIft/oMdgg=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_go_latex_latex",
    importpath = "github.com/go-latex/latex",
    sum = "h1:6zl3BbBhdnMkpSj2YY30qV3gDcVBGtFgVsV3+/i+mKQ=",
    version = "v0.0.0-20210823091927-c0d11ff05a81",
)

go_repository(
    name = "com_github_go_pdf_fpdf",
    importpath = "github.com/go-pdf/fpdf",
    sum = "h1:MlgtGIfsdMEEQJr2le6b/HNr1ZlQwxyWr77r2aj2U/8=",
    version = "v0.6.0",
)

go_repository(
    name = "com_github_googleapis_google_cloud_go_testing",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/googleapis/google-cloud-go-testing",
    sum = "h1:tlyzajkF3030q6M8SvmJSemC9DTHL/xaMa18b65+JM4=",
    version = "v0.0.0-20200911160855-bcd43fbb19e8",
)

go_repository(
    name = "com_github_johncgriffin_overflow",
    importpath = "github.com/JohnCGriffin/overflow",
    sum = "h1:RGWPOewvKIROun94nF7v2cua9qP+thov/7M50KEoeSU=",
    version = "v0.0.0-20211019200055-46fa312c352c",
)

go_repository(
    name = "com_github_kr_fs",
    importpath = "github.com/kr/fs",
    sum = "h1:Jskdu9ieNAYnjxsi0LbQp1ulIKZV1LAFgK1tWhpZgl8=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_lyft_protoc_gen_star",
    importpath = "github.com/lyft/protoc-gen-star",
    sum = "h1:erE0rdztuaDq3bpGifD95wfoPrSZc95nGA6tbiNYh6M=",
    version = "v0.6.1",
)

go_repository(
    name = "com_github_lyft_protoc_gen_star_v2",
    importpath = "github.com/lyft/protoc-gen-star/v2",
    sum = "h1:keaAo8hRuAT0O3DfJ/wM3rufbAjGeJ1lAtWZHDjKGB0=",
    version = "v2.0.1",
)

go_repository(
    name = "com_github_phpdave11_gofpdf",
    importpath = "github.com/phpdave11/gofpdf",
    sum = "h1:KPKiIbfwbvC/wOncwhrpRdXVj2CZTCFlw4wnoyjtHfQ=",
    version = "v1.4.2",
)

go_repository(
    name = "com_github_phpdave11_gofpdi",
    importpath = "github.com/phpdave11/gofpdi",
    sum = "h1:o61duiW8M9sMlkVXWlvP92sZJtGKENvW3VExs6dZukQ=",
    version = "v1.0.13",
)

go_repository(
    name = "com_github_pkg_diff",
    importpath = "github.com/pkg/diff",
    sum = "h1:aoZm08cpOy4WuID//EZDgcC4zIxODThtZNPirFr42+A=",
    version = "v0.0.0-20210226163009-20ebb0f2a09e",
)

go_repository(
    name = "com_github_pkg_sftp",
    importpath = "github.com/pkg/sftp",
    sum = "h1:I2qBYMChEhIjOgazfJmV3/mZM256btk6wkCDRmW7JYs=",
    version = "v1.13.1",
)

go_repository(
    name = "com_github_ruudk_golang_pdf417",
    importpath = "github.com/ruudk/golang-pdf417",
    sum = "h1:K1Xf3bKttbF+koVGaX5xngRIZ5bVjbmPnaxE/dR08uY=",
    version = "v0.0.0-20201230142125-a7e3863a1245",
)

go_repository(
    name = "com_github_spf13_afero",
    importpath = "github.com/spf13/afero",
    sum = "h1:j49Hj62F0n+DaZ1dDCvhABaPNSGNkt32oRFxI33IEMw=",
    version = "v1.9.2",
)

go_repository(
    name = "com_github_xhit_go_str2duration_v2",
    importpath = "github.com/xhit/go-str2duration/v2",
    sum = "h1:lxklc02Drh6ynqX+DdPyp5pCKLUQpRT8bp8Ydu2Bstc=",
    version = "v2.1.0",
)

go_repository(
    name = "com_github_zeebo_assert",
    importpath = "github.com/zeebo/assert",
    sum = "h1:g7C04CbJuIDKNPFHmsk4hwZDO5O+kntRxzaUoNXj+IQ=",
    version = "v1.3.0",
)

go_repository(
    name = "com_google_cloud_go_apikeys",
    importpath = "cloud.google.com/go/apikeys",
    sum = "h1:B9CdHFZTFjVti89tmyXXrO+7vSNo2jvZuHG8zD5trdQ=",
    version = "v0.6.0",
)

go_repository(
    name = "com_google_cloud_go_grafeas",
    importpath = "cloud.google.com/go/grafeas",
    sum = "h1:CYjC+xzdPvbV65gi6Dr4YowKcmLo045pm18L0DhdELM=",
    version = "v0.2.0",
)

go_repository(
    name = "com_google_cloud_go_recaptchaenterprise",
    importpath = "cloud.google.com/go/recaptchaenterprise",
    sum = "h1:u6EznTGzIdsyOsvm+Xkw0aSuKFXQlyjGE9a4exk6iNQ=",
    version = "v1.3.1",
)

go_repository(
    name = "com_google_cloud_go_servicecontrol",
    importpath = "cloud.google.com/go/servicecontrol",
    sum = "h1:d0uV7Qegtfaa7Z2ClDzr9HJmnbJW7jn0WhZ7wOX6hLE=",
    version = "v1.11.1",
)

go_repository(
    name = "com_google_cloud_go_servicemanagement",
    importpath = "cloud.google.com/go/servicemanagement",
    sum = "h1:fopAQI/IAzlxnVeiKn/8WiV6zKndjFkvi+gzu+NjywY=",
    version = "v1.8.0",
)

go_repository(
    name = "com_google_cloud_go_serviceusage",
    importpath = "cloud.google.com/go/serviceusage",
    sum = "h1:rXyq+0+RSIm3HFypctp7WoXxIA563rn206CfMWdqXX4=",
    version = "v1.6.0",
)

go_repository(
    name = "com_google_cloud_go_vision",
    importpath = "cloud.google.com/go/vision",
    sum = "h1:/CsSTkbmO9HC8iQpxbK8ATms3OQaX3YQUeTMGCxlaK4=",
    version = "v1.2.0",
)

go_repository(
    name = "ht_sr_git_sbinet_gg",
    importpath = "git.sr.ht/~sbinet/gg",
    sum = "h1:LNhjNn8DerC8f9DHLz6lS0YYul/b602DUxDgGkd/Aik=",
    version = "v0.3.1",
)

go_repository(
    name = "org_gioui",
    importpath = "gioui.org",
    sum = "h1:K72hopUosKG3ntOPNG4OzzbuhxGuVf06fa2la1/H/Ho=",
    version = "v0.0.0-20210308172011-57750fc8a0a6",
)

go_repository(
    name = "org_uber_go_automaxprocs",
    importpath = "go.uber.org/automaxprocs",
    sum = "h1:2LxUOGiR3O6tw8ui5sZa2LAaHnsviZdVOUZw4fvbnME=",
    version = "v1.5.2",
)

go_repository(
    name = "com_github_martinlindhe_base36",
    importpath = "github.com/martinlindhe/base36",
    sum = "h1:1F1MZ5MGghBXDZ2KJ3QfxmiydlWOGB8HCEtkap5NkVg=",
    version = "v1.1.1",
)

go_rules_dependencies()

go_register_toolchains(version = "1.20.5")

gazelle_dependencies()

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "com_google_protobuf",
    sha256 = "543395bc2ae58e72f7be674221db08b8f14e3bd7e3a19158f76105b3b61570a0",
    strip_prefix = "protobuf-3.21.8",
    urls = ["https://github.com/protocolbuffers/protobuf/archive/v3.21.8.tar.gz"],
)

load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")

protobuf_deps()

# SQLC rules.

load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

# TODO(dsky): update to official community release when merged.
git_repository(
    name = "com_plezentek_rules_sqlc",
    commit = "ad64354652ecfd4a194e264eb387b1a1d760c241",
    remote = "https://github.com/apoxy-dev/rules_sqlc.git",
)

load("@com_plezentek_rules_sqlc//sqlc:deps.bzl", "sqlc_register_toolchains", "sqlc_rules_dependencies")

sqlc_rules_dependencies()

sqlc_register_toolchains()

# Docker rules.

http_archive(
    name = "io_bazel_rules_docker",
    sha256 = "b1e80761a8a8243d03ebca8845e9cc1ba6c82ce7c5179ce2b295cd36f7e394bf",
    urls = ["https://github.com/bazelbuild/rules_docker/releases/download/v0.25.0/rules_docker-v0.25.0.tar.gz"],
)

load(
    "@io_bazel_rules_docker//toolchains/docker:toolchain.bzl",
    docker_toolchain_configure = "toolchain_configure",
)

docker_toolchain_configure(
    name = "docker_config",
)

load(
    "@io_bazel_rules_docker//repositories:repositories.bzl",
    container_repositories = "repositories",
)

container_repositories()

load(
    "@io_bazel_rules_docker//go:image.bzl",
    go_image_repos = "repositories",
)

go_image_repos()

load("@io_bazel_rules_docker//container:container.bzl", "container_pull")

container_pull(
    name = "base_image_arm64",
    architecture = "arm64",
    registry = "docker.io",
    repository = "apoxy/proximal-builder",
    tag = "768d04e11d",
)

container_pull(
    name = "base_image_amd64",
    architecture = "amd64",
    registry = "docker.io",
    repository = "apoxy/proximal-builder",
    tag = "768d04e11d",
)
