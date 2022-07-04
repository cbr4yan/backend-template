FROM golang:1.17-stretch
WORKDIR /app
RUN go install github.com/cespare/reflex@latest
ENTRYPOINT reflex -r '(\.go$|go\.mod)' -s -- sh -c 'go mod tidy && go run ./cmd/app'
