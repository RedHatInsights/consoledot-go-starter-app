# ConsoleDot Golang Starter App

Getting started writing a ConsoleDot app can be confusing. There's a lot of Red Hat "special sauce" that an application developer may not be familiar with. This starter app provides developers with a simple way to bootstrap new applications that are built for the platform from the ground-up. Developers of existing applications can also benefit by using this project as a reference for how to solve common problems.

The starter app is designed to be as simple as possible. It contains only the minimal amount of code required to bootstrap an app on the ConsoleDot platform. There's no example code on how to implement APIs or work with databases; we expect you know that. Instead the project focus is on the stuff you may not know about the platform, and making those things easy. Additionally, the starter app was designed to be as unopinionated as possible. The choice of libraries is kept minimal and low level, with great care put into their selection. That said, you should be able to replace anything you want to, using what's here as an example implmentation.

## Starter App Features
* ClowdApp for [Clowder](https://github.com/RedHatInsights/clowder) managed deployments
* Integrated with Clowder AppConfig via [app-common-go](https://github.com/RedHatInsights/app-common-go)
* One-command deployment to [Ephemeral Environments](https://consoledot.pages.redhat.com/docs/dev/creating-a-new-app/using-ee/index.html)
* Develop locally; test in Ephemeral
* OpenAPI documentation generated from a GoDoc-like syntax provided by [Swag](https://github.com/swaggo/swag/)
* Clowder Provider integration
    * *Note: V1 only supports Postgres. More providers to come.*
* Web API via the [Gin Web Framework](https://gin-gonic.com/docs/)
* Liveliness and Readiness probes
* Easy configuration system that allows for config from Clowder and Environment Variables
* Unit Tests
* Logging via [zerolog](https://github.com/rs/zerolog) along with Gin integration via the [gin-zerolog](https://github.com/dn365/gin-zerolog) middleware
* Metrics via Prometheus client

## Requirements
* Golang 1.18.8 or higher
* Python 3.10 or higher
* Podman 4.5 or higher
* Podman Compose 1.0.6 or higher

## Prerequisites
Before using the starter app there are a few things you'll want to have in place:
1. [Access to the Ephemeral Environment cluster](https://consoledot.pages.redhat.com/docs/dev/creating-a-new-app/using-ee/getting-started-with-ees.html)
2. [Bonfire installed](https://consoledot.pages.redhat.com/docs/dev/creating-a-new-app/using-ee/bonfire.html)
3. [Openshift Console](https://console-openshift-console.apps.c-rh-c-eph.8p0c.p1.openshiftapps.com/command-line-tools) (oc) installed

If you are able to log into the ephmeral cluster and reserve a namespace you are good. If you have no idea what I'm talking about and think I'm just making words up reach out to DevProd and we'll help get you sorted.

## Quickstart: Clone & Run
Want to take it for a test drive before driving it home? Here's how to get started in N minutes where N is impressively small!

1. Clone the repo:
```bash
 git clone git@github.com:RedHatInsights/consoledot-go-starter-app.git
 cd consoledot-go-starter-app
```

2. Run the setup make target to install deps, build the app, and build the api docs
```bash
make setup
```

3. Reserve an ephemeral namespace and put its name in an environment variable. *Note: you need to be logged into the ephemeral cluster.*
```bash
bonfire namespace reserve
2023-06-13 08:28:27 [    INFO] [          MainThread] namespace console url: https://console-openshift-console.apps.c-rh-c-eph.8p0c.p1.openshiftapps.com/k8s/cluster/projects/ephemeral-ratuiz
ephemeral-ratuiz

NAMESPACE=ephemeral-ratuiz
```

4. Launch the starter app in your ephemeral environment:
```bash
make run-ephemeral NAMESPACE=$NAMESPACE
oc process -f deploy/clowdapp.yaml -p NAMESPACE=ephemeral-ratuiz -p ENV_NAME=env-ephemeral-ratuiz  IMAGE=quay.io/rh_ee_addrew/consoledot-go-starter-app IMAGE_TAG=latest | oc create -f -
clowdapp.cloud.redhat.com/go-starter-app created
```

5. Get the route associated with the app:
```bash
oc get route --namespace $NAMESPACE | grep go-starter-app
go-starter-app-app-t57jg                    env-ephemeral-ratuiz-8tlb3rve.apps.c-rh-c-eph.8p0c.p1.openshiftapps.com        /api/starter-app-api/    go-starter-app-app                    auth       edge/Redirect   None
```

6. Give it a minute or two for the deployment to spin up and initial health checks to pass. 

7. Hit the API:
```bash
curl -X GET https://env-ephemeral-ratuiz-8tlb3rve.apps.c-rh-c-eph.8p0c.p1.openshiftapps.com/api/starter-app-api/v1/hello
{"hello":"world"}
```

8. Open the OpenAPI docs in a browser
```bash
xdg-open https://env-ephemeral-ratuiz-8tlb3rve.apps.c-rh-c-eph.8p0c.p1.openshiftapps.com/api/starter-app-api/api_docs/index.html
```

## Fork and Make it Your Own
Intrigued? You can get started building your own app easily.

1. Fork the starter app repo
2. Create a Quay repo that is set to build from your new forked Github repo
3. Clone your fork
4. Run the fork script
```bash
$ make fork
```

The fork script will run through a bunch of questions like your Quay repo, app name, api endpoint - things like that. It will go through the source files and change out the boilerplate names for the config you specified. Commit the changes, push, and you've got your own customized app ready to build whatever you want!

After you've forked you can delete the `scripts` directory, as all it contains is the resources required for the `fork` script. You should also take a look around and see if there's anything we missed that you need to change.

## Local Development
Running in ephemeral is cool, and useful for testing - as well as showing your app is ready to run in stage and prod. But how do you do local development with this project? Easy - you can spin up dependencies like Postgres with podman-compose and then run the app in your IDE of choice to make use of all your favorite debugger features.

When you are ready to get to work you can bring local dependencies up with the `run-local-deps` make target:
``` bash
make run-local-deps
```
Now you can run your app in the IDE debugger and it will automatically connect to the required dependencies. The configuration management system abstracts away the difference between running under Clowder on a cluster or running locally.

*Note: V1 of the starter app only suppport Postgres. We'll add more providers over time. As we do you can adopt those providers by integrating updates to the config module and `env.local` file.*

## Configuration Management
The primary configuration source of truth for this app is the `cdappconfig.json` file generated by Clowder for your app at run time. This contains everything your app needs to run: routes of depedant resources, database creds, secrets, etc. The `config` module handles loading this file for you at app start and can be accessed through the top level `conf.AppConfig` var in the `main` module. 

But what if you are are running locally? We use a local version of that same `cdappconfig.json` file that ships with this project in `local_config.json`. Anything that you put in there will be available to your app, but it must conform to the `AppConfig` spec as defined in `app-common-go`. 

Finally, environment variables are made available through `local.env`. Any environment variables you need can be added to that file and they will be available both locally and when on OpenShift.

You can extend the system to take advantage of more providers by continuing the pattern you see in the `config` module and the `env.local` file.

## ClowdApp
You'll find your app's ClowdApp at `deploy/clowdapp.yaml` - this is what defines your app's config and deployment. You can tailor it to your needs to scale from the simplest database backed API all the way up to the most complex platform app you can imagine. The version we ship here is very simple. It contains only a public web API and a postgres database. We connect to the database but don't do anything with it.

## OpenAPI & Swag
This project uses [Swag](https://github.com/swaggo/swag/) to generate OpenAPI docs from a GoDoc like annotation syntax. You'll notice GoDoc structured comments above the route handlers. These are what is used to generate the OpenAPI docs which are then served out over a route. You are free to change the route, beef up what's in the spec, or modify it however you want. 

```go
// helloWorld godoc
// @Summary      Receive a greeting
// @Description  Receive a greeting from the API
// @Tags         api
// @Produce      json
// @Success      200  {object}  map[string]any
// @Router       /api/starter-app-api/v1/hello [get]
func helloWorld(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"hello": "world",
	})
}
```

When implenting OpenAPI one is invariably faced with a binary choice right from the start: should the API spec generate the API code, or should the code generate the spec? We chose the latter, because we had to make a choice and it seemed simplest, but you are free to rip it out and do it your own way. The way we implemented it isn't the "right" way, its just a good way. Other good ways abound. 

Swag is a very powerful tool and we barely scratch the surface in this project. Definitely head over to their docs and see what else you can do.

To generate new versions of the docs as you make changes run the `api-docs` make target:
```bash
make api-docs
```
## Web Framework: Gin
We had to choose what kind of app to develop for the starter app and we chose a simple web API. That said, many or maybe even most console apps are not web APIs. We chose a web API because it is simple to demo and understand, and requires very little code. "Platform App that isn't an API and doesn't do anything" was a bit wide of a canvas, so we went with a web API. If you need to develop something that isn't a web API it should be very easy to remove the gin parts and leave everything else.

We chose [Gin](https://gin-gonic.com/docs/) after looking at many frameworks. It has the following advantages:
1. One of the most mature and widely used Golang web frameworks
2. One of the fastest Golang web frameworks
3. Very simple to use
4. Ships light-weight and minimal, but supports a wide range of middlewares

That said, you are free to remove it and use what you like. Gin is a recomendation. That said, if you don't have a strong opinion on the matter Gin is a great choice and will allow you to build whatever you need in a fast, idiomatic, and easy to test way.

## Testing
Unit tests are provided for the web API. Run the `test` make target to run the tests:

```bash
make test
```

The test implementation can be found in `main_test.go` and is a very simple example of the [test pattern that Gin recommends](https://gin-gonic.com/docs/testing/). It should be easy to observe the pattern and scale it up. That said, these are just unit tests for the routes. You'll want to think about other testing strategies like end to end, smoke, and integration as you develop.


