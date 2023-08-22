# ConsoleDot Golang Starter App

A template for creating new ConsoleDot services in Golang.

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

## Create a New App
Creating a new app is as simple as running a command and answering a few questions. However, **there are a couple of things you need to have in place before running the template.**

**Before getting started create GitHub and Quay repos for your new application.** Even if you are just creating a test application you will need those two resources before you create your new app. 

The Quay repo is required for pushing images so your app can be deployed to ephemeral. The Github repo is required because all Golang modules needs to be namespaced via a valid repo path. **If you don't have these repos created and use fake information to complete the template your app may not function properly.**


1. Install the NASTI template processor:
```
$ pip install nasti
```

2. Create a project from the template:
```
$ nasti process git@github.com:RedHatInsights/consoledot-go-starter-app.git
```

The template will ask you multiple questions to help configure your app. We suggest using the provided defaults for most quesitons. Once completed check the README in your new project for information on how to configure and run your new app! 


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

