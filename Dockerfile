FROM golang:1.10 AS build

# Install tools required for project
# Run `docker build --no-cache .` to update dependencies
RUN apk add --no-cache git
RUN go get github.com/golang/dep/cmd/dep

# List project dependencies with Gopkg.toml and Gopkg.lock
# These layers are only re-built when Gopkg files are updated
COPY Gopkg.lock Gopkg.toml /go/src/go-starter-commandhandler/
WORKDIR /go/src/go-starter-commandhandler/
# Install library dependencies
RUN dep ensure -vendor-only

# Copy the entire project and build it
# This layer is rebuilt when a file changes in the project directory
COPY . /go/src/go-starter-commandhandler/
RUN go build -o /bin/go-starter-commandhandler

# This results in a single layer image
FROM scratch
COPY --from=build /bin/go-starter-commandhandler /bin/go-starter-commandhandler
ENTRYPOINT ["/bin/go-starter-commandhandler"]
CMD ["--help"]