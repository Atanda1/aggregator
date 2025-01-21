# Start from golang base image
FROM golang:1.22.11-alpine3.21

RUN apk --no-cache add curl
# Install the air binary so we get live code-reloading when we save files
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

# Run the air command in the directory where our code will live
WORKDIR /app
CMD ["air"]