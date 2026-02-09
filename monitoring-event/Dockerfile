

FROM golang:1.23

ARG GIT_BRANCH=unknown
ARG GIT_COMMIT=unknown

WORKDIR /

COPY  . app

WORKDIR /app

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-X main.gitBranch=${GIT_BRANCH} -X main.gitCommit=${GIT_COMMIT}" -o /monitoring-event  cmd/3gpp-monitoring-event/main.go

RUN rm -rf /app
WORKDIR /

EXPOSE 8080

CMD ["/monitoring-event"]