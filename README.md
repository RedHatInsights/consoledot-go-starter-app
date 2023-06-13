# ConsoleDot Go Starter App

As of this writing, this is a work in progress. It isn't quite ready to take out of the oven. That said, if you want to kick it around here's the basics.

1. Clone the app and run the setup make target to install deps and whatnot
2. Log into ephemeral
3. Checkout a namespace
4. Put your namespace into an env var
5. Run the ephemeral deploy make target to deploy to ephemeral
6. Find your OpenShift Route
7. Access your web app!

Example:
```bash

$ make setup
go build -o bin/server main.go
go install github.com/swaggo/swag/cmd/swag@latest

$ bonfire namespace reserve
2023-06-13 08:28:27 [    INFO] [          MainThread] namespace console url: https://console-openshift-console.apps.c-rh-c-eph.8p0c.p1.openshiftapps.com/k8s/cluster/projects/ephemeral-ratuiz
ephemeral-ratuiz

$ NAMESPACE=ephemeral-ratuiz

$ make run-ephemeral NAMESPACE=$NAMESPACE
oc process -f deploy/clowdapp.yaml -p NAMESPACE=ephemeral-ratuiz -p ENV_NAME=env-ephemeral-ratuiz  IMAGE=quay.io/rh_ee_addrew/consoledot-go-starter-app IMAGE_TAG=latest | oc create -f -
clowdapp.cloud.redhat.com/go-starter-app created

$ oc get route --namespace $NAMESPACE | grep go-starter-app
go-starter-app-app-t57jg                    env-ephemeral-ratuiz-8tlb3rve.apps.c-rh-c-eph.8p0c.p1.openshiftapps.com        /api/starter-app/    go-starter-app-app                    auth       edge/Redirect   None

$ curl -X GET https://env-ephemeral-ratuiz-8tlb3rve.apps.c-rh-c-eph.8p0c.p1.openshiftapps.com/api/starter-app/v1/hello
{"hello":"world"}
```

You can access the API Docs in a web browser at the route `/api/starter-app/swagger/index.html` exposed by your OpenShift route / ingress.
```
https://env-ephemeral-ratuiz-8tlb3rve.apps.c-rh-c-eph.8p0c.p1.openshiftapps.com/api/starter-app/swagger/index.html
```

More user-friendly features like a forking script and some cleanup yet to come.

## Forking the Starter App
Start by forking this repo and checking it out. After you've checked it out you can kick the tires and use it as is. As soon as you are ready to make it your own and build on it you can run the fork script. I suggest you do this in another branch and then merge after you've confirmed everything looks OK.

*NOTE: Before you run the fork script make sure you have a Quay repo set up for your project! You will need the Quay repo URL as part of setup.*

```bash
$ make fork
```

The fork script will ask you for some information and then change the files in the repo based on your input. This will do things like change the go module, info in the OpenAPI spec, the API path, change references to the Github and Quay repos and more. Once the script is complete you can remove the scripts directory and commit the changes.

Once the script is complete there are some manual steps you'll need to perform as well. Don't worry, they're not bad.
* Change this README to what makes sense for your project
* Change LICENSE to what make sense for your project
* Once you've integrated auth remove the whitelist paths from the ClowdApp public webService

