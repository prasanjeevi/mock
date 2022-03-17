FROM golang:1.17-alpine AS builder
WORKDIR /app
COPY src/ .
RUN CGO_ENABLED=0 go build  -o mock mock.go

FROM alpine:3.15.1
WORKDIR /opt
COPY --from=builder /app .
ARG port
ARG status
ENV PORT=$port
ENV STATUS=$status
ENTRYPOINT ["/opt/mock"]
