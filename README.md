[![Build Status](https://drone.depado.eu/api/badges/Depado/govue/status.svg)](https://drone.depado.eu/Depado/govue)
# govue
An attempt to work with Vue.js and a golang gin rest server

## Quickstart

If you want to give `govue` a try :
```
go get github.com/Depado/govue
cd $GOPATH/src/Depado/govue
npm install
echo "port: 8080
  debug: true
  api_version: 1" > conf.yml
go build
./govue
```

## Configuration

The current configuration file is `conf.yml` at the root of the project. Here is an example configuration file :

```
port: 8080
debug: true
api_version: 1
```

**port :** The port on which the server should run   
**debug :** Tells gin to run in debug mode or not   
**api_version :** An integer representing the API version. The only API verison for now is `1`    

