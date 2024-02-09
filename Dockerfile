FROM golang:1.21.6-alpine@sha256:fd78f2fb1e49bcf343079bbbb851c936a18fc694df993cbddaa24ace0cc724c5 AS builder
RUN apk add --no-cache ca-certificates git
RUN apk add build-base

WORKDIR /src
# restore dependencies
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Skaffold passes in debug-oriented compiler flags
ARG SKAFFOLD_GO_GCFLAGS
RUN go build -gcflags="${SKAFFOLD_GO_GCFLAGS}" -o /loadgenerator .

# Security -- let's tun the application in an empty OS
FROM alpine:3.19.0@sha256:51b67269f354137895d43f3b3d810bfacd3945438e94dc5ac55fdac340352f48
RUN apk add --no-cache ca-certificates

# todo: dependencies check on precommit

WORKDIR /src
COPY --from=builder /loadgenerator ./executable

# Security -- let's run the application from a dedicated user, not from root
RUN adduser -D applicationuser
RUN chown applicationuser:applicationuser ./executable
RUN chmod +x ./executable
USER applicationuser

# Definition of this variable is used by 'skaffold debug' to identify a golang binary.
# Default behavior - a failure prints a stack trace for the current goroutine.
# See https://golang.org/pkg/runtime/
ENV GOTRACEBACK=single

# sleep infinity is there to prevent the pod restart after the command execution
ENTRYPOINT /src/executable -h "http://${FRONTEND_ADDR}" -u "${USERS:-10}" & sleep infinity 2>&1
