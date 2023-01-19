load("@bazel_gazelle//:def.bzl", "gazelle", "gazelle_binary")

# gazelle:prefix github.com/example/project
gazelle(
    name = "gazelle",
    gazelle = ":gazelle_local",
)

gazelle_binary(
    name = "gazelle_local",
    languages = [
        "//internal/language/test_filegroup",
        "@bazel_gazelle//language/go",  # Built-in rule from gazelle for Golang, 最好带上，否则一些 gazelle directive 无法使用。
        "@bazel_gazelle//language/proto",  # Built-in rule from gazelle for Protos.
    ],
)

filegroup(
    name = "all_files",
    testonly = True,
    srcs = [
        ".bazelrc",
        "BUILD",
        "README.md",
        "WORKSPACE",
        "bazel-bin",
        "bazel-out",
        "bazel-stage1",
        "bazel-testlogs",
        "//internal:all_files",
        "//main:all_files",
        "//xiang_test:all_files",
    ],
    visibility = ["//visibility:public"],
)
