load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
# load("@io_bazel_rules_docker//go:image.bzl", "go_image")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "github.com/edokimok/golang-boilerplate/grpc/server",
    visibility = ["//visibility:private"],
    deps = [
        "//grpc/common:go_default_library",
        "//protobufs/helloworld:go_default_library",
        "@org_golang_google_grpc//:go_default_library",
        "@org_golang_google_grpc//reflection:go_default_library",
        "@org_golang_x_net//context:go_default_library",
    ],
)

go_binary(
    name = "server",
    embed = [":go_default_library"],
    visibility = ["//visibility:public"],
)

# go_image(
#     name = "server_image",
#     embed = [":go_default_library"],
#     goarch = "amd64",
#     goos = "linux",
#     pure = "on",
# )
