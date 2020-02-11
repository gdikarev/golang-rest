FROM golang:1.13-alpine
ENV APP_NAME=golang_rest
ENV PORT=8080
COPY ./ /go/src/${APP_NAME}
WORKDIR /go/src/${APP_NAME}
RUN apk update && apk add git && go get github.com/githubnemo/CompileDaemon
RUN go get github.com/gorilla/mux
RUN go get github.com/BurntSushi/toml
RUN go get github.com/sirupsen/logrus
ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main
CMD ./${APP_NAME}
EXPOSE ${PORT}