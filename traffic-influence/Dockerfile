FROM golang:1.22

WORKDIR /

COPY  . app

WORKDIR /app

RUN go mod tidy
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-X main.gitBranch=$(git rev-parse --abbrev-ref HEAD) -X main.gitCommit=$(git rev-parse HEAD)" -o /traffic-influence  cmd/traffic-influence/main.go

RUN rm -rf /app
WORKDIR /

EXPOSE 8080

CMD ["/traffic-influence"]