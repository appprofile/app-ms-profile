#FROM library/golang
FROM golang:1.10.3

RUN apt-get update; apt-get install -qq make gettext-base; apt-get clean

# Godep for vendoring
RUN go get github.com/tools/godep

# Recompile the standard library without CGO
RUN CGO_ENABLED=0 go install -a std

ENV APP_DIR $GOPATH/src/app-ms-profile
RUN mkdir -p $APP_DIR

# Add files.
ADD . $APP_DIR
COPY ./assets/app.conf /.setup/
COPY ./assets/run.sh /

# Expose port.
EXPOSE 8081

# Se establecen los permisos de ejecución para el fichero run.sh
RUN chmod a+x /run.sh

# Ejecución del aplicativo.
ENTRYPOINT ["/run.sh"]