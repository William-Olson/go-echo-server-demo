FROM golang:1.8

WORKDIR /go/src/server

RUN go-wrapper download -u github.com/labstack/echo/... && \
  go-wrapper download -u github.com/lib/pq && \
  go-wrapper download -u github.com/jinzhu/gorm

COPY . /go/src/server

RUN go-wrapper install .

EXPOSE 7447

CMD go-wrapper run ./*.go

