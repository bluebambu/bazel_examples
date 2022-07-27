load("@rules_cc//cc:defs.bzl", "cc_binary", "cc_library")

cc_library(
    name = "hello-greet",
    srcs = ["hello-greet.cc"],
    hdrs = [
	"hello-greet.h"
	],
)

cc_library(
    name = "xg-utils",
    srcs = ["utils.cpp"],
    hdrs = [
	"utils.hpp"
	],
)

cc_binary(
    name = "hello-world",
    srcs = ["hello-world.cc"],
    deps = [
        ":hello-greet",
        ":xg-utils",
    ],
)
