FROM golang:1.15.4 as builder
COPY . /src/
RUN cd /src && GOPROXY=https://goproxy.io go mod download
RUN cd /src/cmd/logd && CGO_ENABLED=0 go build -o logd

# final stage
FROM juxuny/alpine:3.13.5
WORKDIR /app
COPY --from=builder /src/cmd/logd/logd /app/logd
RUN mkdir -p /app/log
ENV PATH="/app:${PATH}"
ENTRYPOINT /app/logd -d ./log