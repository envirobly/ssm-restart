FROM golang:alpine AS builder
WORKDIR /build
COPY ssm_restart.go .
RUN go build -o ssm_restart_agent ssm_restart.go

FROM scratch AS binaries
COPY --from=builder /build/ssm_restart_agent /

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /build/ssm_restart_agent .
EXPOSE 9009
CMD ["./ssm_restart_agent"]

