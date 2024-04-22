FROM golang:1.7-alpine

# Copy the current directory contents into the container at /home
ADD . /home

# Set the working directory to /home
WORKDIR /home

# Install necessary packages
RUN apk add --no-cache bash git openssh

# Build the Go web application
RUN go build -o app

# Command to run the application
CMD ["./app"]
