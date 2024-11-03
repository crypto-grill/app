# syntax=docker/dockerfile:1

ARG ALPINE_VERSION=3.20
ARG GO_VERSION=1.22
ARG DBMATE_VERSION=v2.16.0

################################################################################
# Create a stage for building the application.
FROM --platform=$BUILDPLATFORM golang:${GO_VERSION}-alpine${ALPINE_VERSION} AS build
WORKDIR /src

# Download dependencies as a separate step to take advantage of Docker's caching.
# Leverage a cache mount to /go/pkg/mod/ to speed up subsequent builds.
# Leverage bind mounts to go.sum and go.mod to avoid having to copy them into
# the container.
RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=go.sum,target=go.sum \
    --mount=type=bind,source=go.mod,target=go.mod \
    go mod download -x

# This is the architecture you’re building for, which is passed in by the builder.
# Placing it here allows the previous steps to be cached across architectures.
ARG TARGETARCH

# Build the application.
# Leverage a cache mount to /go/pkg/mod/ to speed up subsequent builds.
# Leverage a bind mount to the current directory to avoid having to copy the
# source code into the container.
RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,target=. \
    CGO_ENABLED=0 GOARCH=$TARGETARCH go build -o /usr/local/bin/svc .

################################################################################
# Create a new stage for running the application that contains the minimal
# runtime dependencies for the application. This often uses a different base
# image from the build stage where the necessary files are copied from the build
# stage.
#
# The example below uses the alpine image as the foundation for running the app.
# By specifying the "latest" tag, it will also use whatever happens to be the
# most recent version of that image when you build your Dockerfile. If
# reproducability is important, consider using a versioned tag
# (e.g., alpine:3.17.2) or SHA (e.g., alpine@sha256:c41ab5c992deb4fe7e5da09f67a8804a46bd0592bfdf0b1847dde0e0889d2bff).
FROM alpine:${ALPINE_VERSION} AS final

# Install any runtime dependencies that are needed to run your application.
# Leverage a cache mount to /var/cache/apk/ to speed up subsequent builds.
RUN --mount=type=cache,target=/var/cache/apk \
    apk --update add \
        ca-certificates \
        tzdata \
        postgresql-client \
        curl \
        && \
        update-ca-certificates

# This is the architecture you’re building for, which is passed in by the builder.
# Placing it here allows the previous steps to be cached across architectures.
ARG TARGETARCH
ARG DBMATE_VERSION

RUN curl \
    -sSL \
    -o /usr/local/bin/dbmate \
    https://github.com/amacneil/dbmate/releases/download/${DBMATE_VERSION}/dbmate-linux-${TARGETARCH}

RUN chmod \
    +x \
    /usr/local/bin/dbmate

# Copy the executable from the "build" stage.
COPY --from=build /usr/local/bin/svc /app/bin/

# Copy the scripts and migrations from source.
COPY ./scripts /app/scripts/
COPY ./migrations /app/migrations/

# Expose the port that the application listens on.
EXPOSE 80

# What the container should run when it is started.
ENTRYPOINT [ "/app/scripts/run.sh", "serve" ]
