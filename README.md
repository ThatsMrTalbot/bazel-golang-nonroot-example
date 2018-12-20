# Bazel non-root golang container example

## Running 
The Makefile can be used to build and run the docker image. 

- `make build` will build a docker image called, then import it into docker with the name `bazel:image`.
- `make run` will build the image and run it

## Files

- `WORKSPACE` is the file that pulls in dependencies (the Bazel golang library for example)
- `BUILD` is the file containing the build definition

