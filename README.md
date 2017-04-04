# pg-app

An example application for the Cloud Foundry that uses [cf-postgresql-broker](https://github.com/Altoros/cf-postgresql-broker).


## Installation on the Cloud Foundry

```
$ go get github.com/altoros/pg-app
$ cd $GOPATH/github.com/altoros/pg-app
$ cf push
$ cf create-service <serivice-broker-name> <plan-name> <service-name>
$ cf bind-service <app-name> <service-name>
$ cf set-env <app-name> PG_BROKER_NAME <serivice-broker-name>
$ cf start <app-name>
```
