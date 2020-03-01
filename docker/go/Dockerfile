FROM golang:1.13-alpine
ENV APP_NAME=golang_rest
ENV PORT=8080
COPY ./ /go/src/${APP_NAME}
WORKDIR /go/src/${APP_NAME}
RUN apk update && apk add git && go get github.com/githubnemo/CompileDaemon \
    && go get github.com/gorilla/mux \
    && go get github.com/BurntSushi/toml \
    && go get github.com/sirupsen/logrus \
    && go get -v golang.org/x/tools/gopls \
    && go get github.com/go-sql-driver/mysql \
    && go get -u -d github.com/golang-migrate/migrate
ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main
CMD ./${APP_NAME}
EXPOSE ${PORT}