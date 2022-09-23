FROM golang:1.18.6

WORKDIR /app

ADD . /app


EXPOSE 8080

CMD ["./hezzl"]
