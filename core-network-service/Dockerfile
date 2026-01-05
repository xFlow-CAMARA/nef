ARG BUILD_IMAGE=golang:1.23
ARG TARGET_IMAGE=alpine:3.18

#---------------------------------------------------------------------
# BUILDER IMAGE
#---------------------------------------------------------------------
FROM $BUILD_IMAGE AS core-network-service-build
WORKDIR /core-network-service
ENV CGO_ENABLED=0
COPY go.* ./
RUN go mod tidy
COPY . ./
RUN go build -o bin/core-network-service cmd/core-network-service/main.go

#---------------------------------------------------------------------
# TARGET IMAGE
#---------------------------------------------------------------------
FROM $TARGET_IMAGE AS core-network-service
RUN apk update && apk add -U net-tools curl bash
COPY --from=core-network-service-build /core-network-service/bin/core-network-service ./
COPY .env ./
EXPOSE 8080/tcp
ENTRYPOINT ["./core-network-service"]
