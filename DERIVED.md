# ConsoleDot Go Starter App API
Welcome to your new app!

## Configuration

### Clowder & App Config
Clowdapps use the `cdappconfig.json` file for their infrustructure and resource config information. Database, kafka, storage and other config information is made available to your app through `cdappconfig.json`. When running on a cluster the app will load the Clowder provided `cdappconfig.json` and make it available on the `Config` object. When running locally your app will load the `cdappconfig.json` at the root of your application's source tree and make it available on the `Config` object.

### Environment Variables
You can put any environemnt variable definitions you want in the `local.env` file. That file and any definitions in it are loaded at run time and are available to your app through the normal Golang OS env var library.

### Config Files
There's no single idiomatic way to handle config in Golang apps. Environment variables are probably the most popular, but other options exist. This starter app includes no configuration beyond environment variables and Clowder. You are welcome to use whatever else you like.

## Run in Ephemeral

0. Make sure you are logged into the Ephemeral Cluster
1. Make sure you are docker/podman logged into quay.io
2. Check out a namespace `bonfire namespace reserve`
3. Deploy the app into your namespace `make run-ephemeral NAMESPACE=$NAMESPCACE`

## Run Local

0. Run the local dependency stack `make run-local-deps`
1. Run the app `make run`