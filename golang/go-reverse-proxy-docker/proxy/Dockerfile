FROM golang:alpine
# Add Maintainer Info
LABEL maintainer="Jose Mejia"
# Install git.
# Git is required for fetching the dependencies.
RUN apk add git
# RUN apk update && apk add --no-cache git
# Create folder app
RUN mkdir /app
# Build Args
ARG LOG_DIR=/logs
# Create Log Directory
RUN mkdir -p ${LOG_DIR}
# Environment Variables
ENV LOG_FILE_LOCATION=${LOG_DIR}/
ENV ENVIRONMENT=dev
# Copy sources
COPY . /app
# Set the Current Working Directory inside the container
WORKDIR /app
# get dependencies
RUN go get -d -v
# Build the Go app
RUN go build -o main .
# Declare volumes to mount
VOLUME [${LOG_DIR}]
# This container exposes port
EXPOSE 4443 
# Run
CMD ["./main"]