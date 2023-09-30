FROM golang:alpine
# testing docker pipelines
RUN mkdir /app

WORKDIR /app

ADD go.mod .
ADD go.sum .

RUN go mod download
ADD . .

RUN go get github.com/githubnemo/CompileDaemon

EXPOSE 8003
RUN go run cmd/main.go
ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main
