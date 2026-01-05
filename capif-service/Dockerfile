

FROM golang:1.22

WORKDIR /app

RUN apt update
RUN apt install python3 python3-requests -y

RUN mkdir -p /etc/capif-service/certs

COPY . .
RUN go mod download


RUN CGO_ENABLED=0 GOOS=linux go build -o /capif-service cmd/main.go

EXPOSE 8080


CMD ["/capif-service"]