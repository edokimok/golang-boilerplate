load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

rules_go_version = "0.42.0"
rules_go_sha = "91585017debb61982f7054c9688857a2ad1fd823fc3f9cb05048b0025c47d023"
http_archive(
    name = "io_bazel_rules_go",
    urls = [
            "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v%s/rules_go-v%s.zip" % (rules_go_version, rules_go_version),
            "https://github.com/bazelbuild/rules_go/releases/download/v%s/rules_go-v%s.zip" % (rules_go_version, rules_go_version),
        ],
    sha256 = rules_go_sha,
)

http_archive(
    name = "bazel_gazelle",
    integrity = "sha256-MpOL2hbmcABjA1R5Bj2dJMYO2o15/Uc5Vj9Q0zHLMgk=",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.35.0/bazel-gazelle-v0.35.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.35.0/bazel-gazelle-v0.35.0.tar.gz",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

go_repository(
    name = "org_golang_google_grpc",
    commit = "5e8f83304c0563d1ba74db05fee83d9c18ab9a58",  # 1.32.0
    importpath = "google.golang.org/grpc",
)

go_repository(
    name = "org_golang_x_net",
    importpath = "golang.org/x/net",
    commit = "ab34263943818b32f575efc978a3d24e80b04bd7",  # release-branch.go1.15
)

go_repository(
    name = "org_golang_x_text",
    importpath = "golang.org/x/text",
    commit = "afb9336c4530b4b18f37130eab53f245f7d6821e",  # release-branch.go1.15
)

go_repository(
    name = "org_golang_x_mod",
    commit = "859b3ef565e237f9f1a0fb6b55385c497545680d",  # release-branch.go1.15
    importpath = "golang.org/x/mod",
)

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "protobuf_java",
    urls = ["https://github.com/protocolbuffers/protobuf/releases/download/v3.18.1/protobuf-java-3.18.1.tar.gz"],
    strip_prefix = "protobuf-3.18.1",
)

go_rules_dependencies()

go_register_toolchains(version = "1.21.4")

gazelle_dependencies()

load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

git_repository(
    name = "com_google_protobuf",
    commit = "fde7cf7358ec7cd69e8db9be4f1fa6a5c431386a",  # v3.13.0
    remote = "https://github.com/protocolbuffers/protobuf",
)

load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")

protobuf_deps()
