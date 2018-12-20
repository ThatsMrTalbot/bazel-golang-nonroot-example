load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("@io_bazel_rules_docker//go:image.bzl", "go_image")
load("@io_bazel_rules_docker//container:container.bzl", "container_push", "container_image")
load("@io_bazel_rules_docker//contrib:passwd.bzl", "passwd_entry", "passwd_tar")

# root user passwd entry
passwd_entry(
    name = "root_user",
    username = "root",
    uid = 0,
    gid = 0,
)

# non root user passwd entry
passwd_entry(
    name = "groot_user",
    username = "groot",
    info = "groot",
    uid = 1002,
)

# create tarball containing build passwd file 
passwd_tar(
    name = "passwd",
    entries = [
        ":root_user",
        ":groot_user",
    ],
    passwd_file_pkg_dir = "etc",
)

# build image using default golang base image + passwd tarball
container_image(
    name = "passwd_image",
    base = "@go_image_base//image",
    tars = [":passwd"],
    user = "groot",
)

# build image from constructed base image + compiled go program
go_image(
    name = "image",
    srcs = ["main.go"],
    visibility = ["//visibility:public"],
    pure = "on",
    base = ":passwd_image",
)
