FROM golang:alpine AS builder

ENV USER=appuser
ENV UID=1001

RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

RUN apk update && apk add --no-cache git

RUN mkdir /src
WORKDIR /src

COPY go.mod .
COPY go.sum .

RUN go mod download
RUN go mod verify

ADD . .

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -tags timetzdata -ldflags='-w -s -extldflags "-static"' -a -o /createVm

FROM scratch

COPY --from=alpine:latest /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /createVm /createVm

ENTRYPOINT ["/createVm"]