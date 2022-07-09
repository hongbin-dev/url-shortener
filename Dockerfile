FROM golang:1.18.3

WORKDIR /src

COPY . /src

RUN go build -o Go_Echo_Default

EXPOSE 8080
CMD ["./Go_Echo_Default"]