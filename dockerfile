FROM golang:1.18-alpine AS builder

# Set maintainer
# LABEL Maintainer="Marsk <manhtokim@gmail.com>"

# Set HEADER AND ENV FILES
ARG HEADER_FILE
ARG ENV_FILE
ENV HEADER_FILE=header_dev.go
ENV ENV_FILE=.env.dev

# Check HEADER_FILE & ENV_FILE
RUN echo "File swagger: $HEADER_FILE"
RUN echo "File env: $ENV_FILE"

RUN apk add bash ca-certificates git gcc g++ libc-dev

# Set working directory for the build
RUN mkdir -p /work/vigour
WORKDIR /work/vigour

# Copy go.mod and go.sum
COPY go.mod .
COPY go.sum .
RUN ls -la /work/vigour/

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# COPY everything else
COPY . /work/vigour/

RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init --parseDependency -g header_dev.go

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o vigour .

# EXPOSE 1900

CMD ["./vigour"]