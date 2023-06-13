# ConsoleDot Go Starter App

As of this writing, this is a work in progress. It isn't quite ready to take out of the oven. That said, if you want to kick it around here's the basics.

1. Clone the app
2. Log into ephemeral
3. Checkout a namespace
4. Put your namespace into an env var
5. Run the ephemeral deploy make target to deploy to ephemeral
6. Find your OpenShift Route
7. Access your web app!

Example:
```bash
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
