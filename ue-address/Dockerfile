

FROM golang:1.23

ARG GIT_BRANCH=unknown
ARG GIT_COMMIT=unknown

WORKDIR /

COPY  . app

WORKDIR /app

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-X main.gitBranch=${GIT_BRANCH} -X main.gitCommit=${GIT_COMMIT}" -o /ue-address  cmd/3gpp-ue-address/main.go

RUN rm -rf /app
WORKDIR /

EXPOSE 8080

CMD ["/ue-address"]