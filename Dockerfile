FROM golang:1.12 as builder

ADD . /
RUN go build -o /main /main.go

FROM debian:9-slim

COPY --from=builder /main /

ENTRYPOINT /main