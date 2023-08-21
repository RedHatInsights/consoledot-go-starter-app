# ConsoleDot Go Starter App API
Welcome to your new app! This README contains all the info you need to get started. After you are comfortable with the app feel free to replace the contents of this file.

## Quickstart

Before using the app there are a few things you'll want to have in place:
1. [Access to the Ephemeral Environment cluster](https://consoledot.pages.redhat.com/docs/dev/creating-a-new-app/using-ee/getting-started-with-ees.html)
2. [Bonfire installed](https://consoledot.pages.redhat.com/docs/dev/creating-a-new-app/using-ee/bonfire.html)
3. [Openshift Console](https://console-openshift-console.apps.c-rh-c-eph.8p0c.p1.openshiftapps.com/command-line-tools) (oc) installed




To deploy this app in ephemeral do the following:

0. Use `oc` to log in to the ephemeral cluster and your container engine to log into quay:
``` bash
$ oc login --token=... --server=https://api.c-rh-c-eph.8p0c.p1.openshiftapps.com:6443
$ podman login quay.io
```

1. Reserve a namespace and save the namespace name in an environment variable
```bash
$ bonfire namespace reserve
...
ephemeral-utemsa
$ NAMESPACE=ephemeral-utemsa
```

2. Deploy the app:
```bash
$ make run-ephemeral NAMESPACE=$NAMESPACE
...
clowdapp.cloud.redhat.com/awesome-new-app created
```

3. Get your app's ingress
```bash
$ oc get ingress -n $NAMESPACE
NAME                                  CLASS               HOSTS                                                                          ADDRESS                                                    PORTS     AGE
awesome-new-app-awesome-new-app       openshift-default   env-ephemeral-5irb9j-uqojupc5.apps.c-rh-c-eph.8p0c.p1.openshiftapps.com        router-default.apps.c-rh-c-eph.8p0c.p1.openshiftapps.com   
...
```
It'll take a minute or so for your app to spin up, so throw a ball around with a dog or something and then come back.

4. Say hi to your app!
```
$ curl -XGET https://env-ephemeral-5irb9j-uqojupc5.apps.c-rh-c-eph.8p0c.p1.openshiftapps.com/api/starter-app-api/v1/hello
{"hello":"world"}
```

5. Check out the API Docs in a browser:
```
$ xdg-open https://env-ephemeral-5irb9j-uqojupc5.apps.c-rh-c-eph.8p0c.p1.openshiftapps.com/api/starter-app-api/api-docs/index.html

```

## Configuration

### Clowder & App Config
Clowdapps use the `cdappconfig.json` file for their infrustructure and resource config information. Database, kafka, storage and other config information is made available to your app through `cdappconfig.json`. When running on a cluster the app will load the Clowder provided `cdappconfig.json` and make it available on the `Config` object. When running locally your app will load the `cdappconfig.json` at the root of your application's source tree and make it available on the `Config` object.

### Environment Variables
You can put any environemnt variable definitions you want in the `local.env` file. That file and any definitions in it are loaded at run time and are available to your app through the normal Golang OS env var library.

### Config Files
There's no single idiomatic way to handle config in Golang apps. Environment variables are probably the most popular, but other options exist. This starter app includes no configuration beyond environment variables and Clowder. You are welcome to use whatever else you like.

## Running Locally

0. Start the local dependency stack. This will start your database.
```
$ make run-local-depts
```
1. Run the app 
```
$ make run
```