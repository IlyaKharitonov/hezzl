FROM golang:1.18.6

WORKDIR /app

ADD . /app

#CMD "./hezzl"

EXPOSE 8080 8080

