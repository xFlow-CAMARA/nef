FROM golang:1.22

ARG GIT_BRANCH=unknown
ARG GIT_COMMIT=unknown

WORKDIR /

COPY  . app

WORKDIR /app

RUN go mod tidy
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-X main.gitBranch=${GIT_BRANCH} -X main.gitCommit=${GIT_COMMIT}" -o /traffic-influence  cmd/traffic-influence/main.go

RUN rm -rf /app
WORKDIR /

EXPOSE 8080

CMD ["/traffic-influence"]