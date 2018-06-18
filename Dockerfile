FROM library/golang

# Godep for vendoring
RUN go get github.com/tools/godep

# Recompile the standard library without CGO
RUN CGO_ENABLED=0 go install -a std

ENV APP_DIR $GOPATH/src/app-ms-profile
RUN mkdir -p $APP_DIR

# Add files.
COPY /assets/app.conf /.setup/
COPY /assets/run.sh /

# Expose port.
EXPOSE 8081

# Se establecen los permisos de ejecución para el fichero run.sh
RUN chmod a+x /run.sh

# Ejecución del aplicativo.
ENTRYPOINT ["/run.sh"]
