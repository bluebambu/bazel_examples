load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "test_filegroup",
    srcs = ["lang.go"],
    importpath = "github.com/example/project/internal/language/test_filegroup",
    visibility = ["//visibility:public"],
    deps = [
        "@bazel_gazelle//config:go_default_library",
        "@bazel_gazelle//label:go_default_library",
        "@bazel_gazelle//language:go_default_library",
        "@bazel_gazelle//repo:go_default_library",
        "@bazel_gazelle//resolve:go_default_library",
        "@bazel_gazelle//rule:go_default_library",
    ],
)

filegroup(
    name = "all_files",
    testonly = True,
    srcs = [
        "BUILD",
        "lang.go",
    ],
    visibility = ["//visibility:public"],
)
