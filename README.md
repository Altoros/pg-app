# pg-app

An example application for the Cloud Foundry that uses [cf-postgresql-broker](https://github.com/Altoros/cf-postgresql-broker).

## Requirements

* [govendor](https://github.com/kardianos/govendor)
* [go-buildpack](https://github.com/cloudfoundry/go-buildpack) >= 1.7.7
* [cf-postgresql-broker](https://github.com/Altoros/cf-postgresql-broker)

## Installation on the Cloud Foundry

```
$ go get github.com/altoros/pg-app
$ cd $GOPATH/github.com/altoros/pg-app
$ govendor sync
$ cf push <app-name> --no-start -m 128M -k 256M [-b https://github.com/cloudfoundry/go-buildpack#v1.7.7]
$ cf create-service <serivice-broker-name> <plan-name> <service-name>
$ cf bind-service <app-name> <service-name>
$ cf set-env <app-name> PG_BROKER_NAME <serivice-broker-name>
$ cf start <app-name>
```
