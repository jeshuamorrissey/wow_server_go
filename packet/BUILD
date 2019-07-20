load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "packet",
    srcs = [
        "packet.go",
    ],
    importpath = "gitlab.com/jeshuamorrissey/wow-server-go/packet",
    visibility = ["//visibility:public"],
    deps = [
        "@com_github_gotest_tools//assert:go_default_library",
    ],
)

go_test(
    name = "packet_test",
    srcs = ["packet_test.go"],
    embed = [":packet"],
    importpath = "gitlab.com/jeshuamorrissey/wow-server-go/packet",
)
