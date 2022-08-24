# syntax=docker/dockerfile:1

# golang:1.18-buster is used
FROM golang:1.18-buster

# Create build directory INSIDE the image and Set as destination for COPY
WORKDIR /app

# Copy the source code into build image dir. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
#COPY . ./

# Create production dir/subdir on image
RUN mkdir /unl
RUN mkdir /unl/build

# Copy react build from source to production.
COPY ./build /unl/build

# Copy golang staff to WORDIR
COPY go.mod web.go ./

# Build goweb-docker production exe on image production dir
RUN go build -o /unl/goweb-docker

# This is for documentation purposes only.
# To actually open the port, runtime parameters
# must be supplied to the docker command.
EXPOSE 8080

# (Optional) environment variable that our dockerised
# application can make use of. The value of environment
# variables can also be set via parameters supplied
# to the docker command on the command line.
#ENV HTTP_PORT=8081

# Change build dir to producton dir
WORKDIR /unl

# Run goweb-docker from production dir
CMD [ "./goweb-docker" ]

##docker build -t goweb-docker .
##docker run -d -p 3000:8080 --name goweb-docker goweb-docker
