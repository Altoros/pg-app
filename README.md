# pg-app

A small application for the Cloud Foundry that uses [cf-postgresql-broker](https://github.com/Altoros/cf-postgresql-broker). Simply shows `SELECT version()` query result.

## Installation

1. Set up [cf-postgresql-broker](https://github.com/Altoros/cf-postgresql-broker)
2. Push the application to the Cloud Foundry `$ cf push pg-app --no-start -m 128M -k 256M`
3. Create a postgresql service `$ cf create-service p-postgresql basic pgsql` if you did not
4. Bind the application `$ cf bind-service pg-app pgsql` to the service
5. Start the application `$ cf start pg-app`
6. Verify the installation `$ curl -v $(cf app pg-app | grep urls: | awk '{print $2}')`
