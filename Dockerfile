# syntax=docker/dockerfile:1

## Multisage dockerfile
## Build using golang:1.18-buster
FROM golang:1.18-buster AS build

# Create build directory INSIDE first stage image and Set as destination for COPY
WORKDIR /app

# Create production dir/subdir on first stage image
RUN mkdir /unl
RUN mkdir /unl/build

# Copy react build from source to production.
COPY ./build /unl/build

# Copy golang staff to WORDIR
COPY go.mod web.go ./

# Build goweb-docker production exe on first stage production dir
RUN go build -o /unl/goweb-docker

## Deploy using “distroless” base image for our Go application/
## Distroless base imageis very barebones and is meant for lean deployments of static binaries.
FROM gcr.io/distroless/base-debian10

# Create producton dir in distroless image
WORKDIR /unl

# Copy all from build prodiction to distroless production
COPY --from=build /unl /unl

EXPOSE 8080

USER nonroot:nonroot

# Start goweb-docker from distroless image production dir
ENTRYPOINT ["/unl/goweb-docker"]

##docker build -t goweb-docker:multistage .
##docker run -d -p 3000:8080 --name goweb-docker goweb-docker:multistage
