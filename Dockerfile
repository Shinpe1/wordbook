# Build image
FROM golang:1.15.5-buster AS go-builder

# Install dependencies
WORKDIR /go/src/github.com/Shinpe1/wordbook_web

# Build modules
COPY main.go .
RUN GOOS=linux go build -a -o wordbook_web .

# ----------
# Production image
FROM busybox
WORKDIR /opt/wordbook/bin

# Deploy modules
COPY --from=go-builder /go/src/github.com/Shinpe1/wordbook_web .
ENTRYPOINT ["./wordbook-anywhere"]