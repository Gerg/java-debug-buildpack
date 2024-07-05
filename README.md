# Java Debug Buildpack

## About

Proof-of-concept Cloud Foundry buildpack for collecting debug information about
java applications at runtime.

## How it Works

This buildpack adds a [sidecar
process](https://docs.cloudfoundry.org/devguide/sidecars.html) to any
application built with it. The sidecar is a simple web server. If a route is
mapped to the web server and an HTTP request is issued to it, then the sidecar
collects a java thread dump from the running java application.

## Usage

1. Build the sidecar binary
    ```
    $ GOOS=linux GOARCH=amd64 go build java_debug_sidecar.go
    ```
1. Zip it up:
    ```
    # from the root of this repo, zip up buildpack excluding the .git directory 
    $ zip -r java-debug-buildpack.zip . -x "*.git*"
    ```
1. Upload the buildpack to cloudfoundry
    ```
    $ cf create-buildpack java_debug_buildpack java-debug-buildpack.zip <index>
    ```
1. Push a java app
    ```
    $ cd <app-rootdir>
    # note the debug buildpack cannot be the final buildpack
    $ cf push my-app -b java_debug_buildpack -b java_buildpack
    ```

In a real use case, you would set up an external route to the sidecar via the
v3 API, but it is not currently well-supported via the CLI. For simplicity,
it's easier to ssh in and call the sidecar on localhost:

```
$ cf ssh my-app
...
$ curl localhost:8081/threaddump
```

The thread dump will then be added to the app's logs. 
