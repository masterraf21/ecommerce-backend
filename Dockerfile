# Start from golang base image
FROM golang:1.13-alpine as builder

# ENV GO111MODULE=on
# ARG
ARG PORT
ARG MONGO_HOST
ARG MONGO_DATABASE
ARG MONGO_OPTIONS
ARG TIMEOUT_ON_SECONDS
ARG OPERATION_ON_EACH_CONTEXT
ARG SECRET
ARG REFRESH_SECRET
# ENV
ENV PORT=${PORT}
ENV MONGO_HOST=${MONGO_HOST}}
ENV MONGO_DATABASE=${MONGO_DATABASE}}
ENV MONGO_OPTIONS=${MONGO_OPTIONS}}
ENV TIMEOUT_ON_SECONDS=${TIMEOUT_ON_SECONDS}
ENV OPERATION_ON_EACH_CONTEXT=${TIMEOUT_ON_SECONDS}
ENV SECRET=${SECRET}
ENV REFRESH_SECRET=${REFRESH_SECRET}
# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

# Set the current working directory inside the container 
WORKDIR /app

# Copy go mod and sum files 
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed 
RUN go mod download 

# Copy the source from the current directory to the working Directory inside the container 
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Start a new stage from scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage. Observe we also copied the .env file
COPY --from=builder /app/main .
COPY --from=builder /app/.env .       

# Expose port 8800 to the outside world
EXPOSE 8800

#Command to run the executable
CMD ["./main"]