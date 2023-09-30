FROM golang

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -o main ./main.go

EXPOSE 8090
CMD [ "/app/main" ]