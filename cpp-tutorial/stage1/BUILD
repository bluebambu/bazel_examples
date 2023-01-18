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

tconf_binary(
    name = "tconf_bin_",
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
    custom_xiang = 2324,
    visibility = ["//visibility:public"],
)
