FROM golang:1.8

WORKDIR /go/src/server

RUN go-wrapper download -u github.com/labstack/echo/... && \
  go-wrapper download -u github.com/lib/pq && \
  go-wrapper download -u github.com/jinzhu/gorm && \
  go-wrapper download -u gopkg.in/matryer/try.v1

COPY . /go/src/server

RUN go-wrapper install .

EXPOSE 7447

CMD go-wrapper run ./*.go

