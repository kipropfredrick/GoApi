# use official Golang image
FROM golang:alpine

# set working directory
WORKDIR /app

# Copy the source code
COPY . .
COPY .env .
# Download and install the dependencies
#RUN go get -d -v ./...c

# Build the Go app
RUN go build -o api .

#EXPOSE the port
# Run the executable
CMD ["./api"]