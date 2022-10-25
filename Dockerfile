FROM golang:1.17 as builder

#
RUN mkdir -p $GOPATH/src/github.com/Mirobidjon/task/rest_service
WORKDIR $GOPATH/src/github.com/Mirobidjon/task/rest_service

# Copy the local package files to the container's workspace.
COPY . ./

# installing depends and build
RUN export CGO_ENABLED=0 && \
    export GOOS=linux && \
    go mod vendor && \
    make build && \
    mv ./bin/rest_service /

FROM alpine
COPY --from=builder rest_service .

ENTRYPOINT ["/rest_service"]
