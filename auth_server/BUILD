load("@io_bazel_rules_go//go:def.bzl", "go_binary")

go_binary(
    name = "auth_server",
    srcs = ["auth_server.go"],
    importpath = "gitlab.com/jeshuamorrissey/wow-server-go/auth_server",
    deps = ["//packet"],
)
