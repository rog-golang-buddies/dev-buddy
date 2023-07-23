FROM golang:1.18 as build
WORKDIR /go/src/app

# Copy the entire source code to the container's workspace
COPY . .
# Debugging step: Check current directory and contents of the 'cmd' directory
RUN pwd
RUN ls -la /go/src/app/cmd

# Static build
RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o /go/bin/app ./cmd
FROM gcr.io/distroless/static:nonroot
COPY --from=build /go/bin/app /app

USER 65532:65532

CMD ["/app"]
