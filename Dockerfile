FROM golang:1.8

WORKDIR /go/src/

RUN go-wrapper download -u github.com/labstack/echo/... && \
  go-wrapper download -u github.com/lib/pq && \
  go-wrapper download -u github.com/jinzhu/gorm && \
  go-wrapper download -u gopkg.in/matryer/try.v1 && \
  go-wrapper download -u golang.org/x/crypto/bcrypt && \
  go-wrapper download -u github.com/dgrijalva/jwt-go

COPY . /go/src/

RUN go-wrapper install ./utils
RUN go-wrapper install ./models
RUN go-wrapper install ./api
RUN go-wrapper install ./app

EXPOSE 7447

CMD app
