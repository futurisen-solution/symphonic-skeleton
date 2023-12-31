FROM golang:1.21.1-alpine AS builder

ARG SSH_PRIVATE_KEY

WORKDIR /app

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV GOPRIVATE=github.com/futurisen-solution/*

RUN apk --no-cache add git openssh-client

RUN --mount=type=ssh mkdir -p -m 0700 /root/.ssh && \
    rm -rf /root/.ssh/known_hosts && \
    echo "$SSH_PRIVATE_KEY" > /root/.ssh/id_rsa && \
    chmod 600 /root/.ssh/id_rsa && \
    ssh-keyscan github.com >> /root/.ssh/known_hosts

RUN git config --global url."ssh://git@github.com/".insteadOf https://github.com/

COPY go.mod .
COPY go.sum .

RUN --mount=type=ssh go mod download
RUN go mod verify

COPY .. .

RUN \
    files=$(find internal/**/cmd -name main.go); \
    for file in $files; do \
      target=$(dirname $file); \
      go build -mod=readonly -ldflags "-s -w" -o /dist/"$target" $file; \
    done

RUN go build -mod=readonly -ldflags "-s -w" -o /dist/SERVICE_NAME .

FROM scratch

COPY --from=builder /dist/SERVICE_NAME .
COPY --from=builder /dist/internal internal

ENTRYPOINT ["/SERVICE_NAME"]