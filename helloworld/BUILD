load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "hello2",
    srcs = ["hello2.go"],
    importpath = "github.com/edokimok/golang-boilerplate/helloworld_hello2",
    visibility = ["//visibility:public"],
)

go_binary(
    name = "helloworld",
    srcs = ["main.go"],
    importpath = "github.com/edokimok/golang-boilerplate/helloworld",
    visibility = ["//visibility:public"],
    deps = [
        "//helloworld/hello:go_default_library",
        "//helloworld:hello2"
    ],
)
