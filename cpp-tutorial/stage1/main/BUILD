load("@rules_cc//cc:defs.bzl", "cc_binary")

cc_binary(
    name = "hello-world",
    srcs = ["hello-world.cc"],
)

filegroup(
    name = "all_files",
    testonly = True,
    srcs = [
        "BUILD",
        "hello-world.cc",
    ],
    visibility = ["//visibility:public"],
)

tconf_binary(
    name = "tconf_bin_main",
    srcs = [
        "BUILD",
        "hello-world.cc",
    ],
    custom_xiang = 2324,
    visibility = ["//visibility:public"],
)
