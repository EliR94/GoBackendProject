# Build stage
FROM golang:1.24 AS builder

# Make the destination folder for all action in the container the /project folder
WORKDIR /project

# Copy all files in this project from source to container
COPY ./Game .

# Build binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o exec

# Execute the following command (in this case, run the 'exec' file (the API we just built))
CMD ["./exec"]