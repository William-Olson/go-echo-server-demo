FROM golang:1.8

WORKDIR /go/src/server

COPY . /go/src/server

RUN go-wrapper download -u github.com/labstack/echo/... && \
  go-wrapper install .

EXPOSE 7447

CMD go-wrapper run ./*.go

