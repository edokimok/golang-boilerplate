load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["hello.go"],
    importpath = "github.com/edokimok/golang-boilerplate/helloworld/hello",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = ["hello_test.go"],
    embed = [":go_default_library"],
)
