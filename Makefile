.PHONY: build
build:
	# "bazel run" instead of "bazel build" as running instructs bazel to import the image into docker after building 
	bazel run //:image --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 -- --norun

.PHONY: run
run: build
	# you could also use "bazel run", using "docker run" to prove the container runs as a non root user
	docker run --rm bazel:image
