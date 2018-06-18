# AppMsProfile

RESTful API for profiles CRUD.

## Installing

**NOTE:** It is not necessary execute this step if you only want to run the microservice. Go to **Deployment** section.

### Installing Golang

Download golang from [here](https://golang.org/dl/). Extract the content in `$HOME/go`. 

Make sure to configure the following environment vars.

```
# Golang config.
export GOROOT=$HOME/go
export PATH=$PATH:$GOROOT/bin
export GOPATH=$HOMEgoworkspace
export PATH=$PATH:$GOPATH/bin
export GOBIN=$GOPATH/bin
```

### Installing Beego

Follow the next steps to install the framework Beego in Cloud9. See the official
guide here https://beego.me/docs/install/bee.md

1. Run the following command.

```
$ go get github.com/beego/bee
```

2. bee is installed in GOBIN by default, so we need to add GOBIN to the PATH, 
otherwise the bee command won’t work. To do that we need modify the source 
~/.profile and add the following lines at the end.

```
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN
```

3. Run the bee command to verify.

```
$ bee
```

## Deployment

### Development server

#### Run with Docker

Run `docker build -t app-ms-profile .` to build the image. 

Run `docker-compose up` to execute. 

**NOTE:** Make sure you have the right `CONFIG_SERVER` and `CONFIG_BRANCH` in the `docker-compose.yml` file.

#### Run with Beego

Run `bee run` for a development server. Navigate to `https://{HOST}:{PORT}/swagger` to see the API definition.

