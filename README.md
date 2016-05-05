# pg-app

A small application for the Cloud Foundry that uses [cf-postgresql-broker](https://github.com/Altoros/cf-postgresql-broker). Simply shows `SELECT version()` query result.

## Installation

1. Install and configure [cf-postgresql-broker](https://github.com/Altoros/cf-postgresql-broker)
2. Push the application to the Cloud Foundry `$ cf push <app-name> --no-start -m 128M -k 256M`
3. Create a postgresql service `$ cf create-service <serivice-broker-name> <plan-name> <service-name>` if you did not.
4. Bind the application `$ cf bind-service <app-name> <service-name>` to the service
5. Set the broker name environment variable `$ cf set-env <app-name> PG_BROKER_NAME <serivice-broker-name>`
6. Start the application `$ cf start <app-name>`
7. Verify the installation `$ curl -v $(cf app <app-name> | grep urls: | awk '{print $2}')`
