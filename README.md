# AppMsProfile

RESTful API for profiles CRUD.

## Installing

### Installing Golang



### Installing Beego

**NOTE:** It is not necessary execute this step if you only want to run the microservice.

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

#### Run

#### Run with Beego

Run `bee run` for a development server. Navigate to `https://{HOST}:{PORT}/swagger` to see the API definition.

