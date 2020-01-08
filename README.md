## Toplin 

Toplin is a readonly web app for viewing todo lists stored in [Joplin](https://joplinapp.org/). The value add that Toplin provides is organization of your todo lists by dates for easy viewing of what's been recently done and what is currently todo.

Toplin is meant to run on the same machine as your Joplin desktop client app. To access the Joplin API, Toplin requires that Joplin clipper service be enabled on your local instance of Joplin. Toplin does not require installation of the Joplin browser extension.

## Dependencies

1. Go 1.4.x or greater with a properly configured $GOPATH
2. git

## Installation

1. Clone the repository 
    ``` 
    git clone git@github.com:geordie/toplin.git
    ```

2. Navigate to the directory containing the cloned repository
    ```
    cd toplin
    ```

3. Make a copy of the `config.sample.yaml` file named `config.yaml`
    ```
    cp config.sample.yaml config.yaml
    ```

4. Edit config.yaml with the configuration settings specific to your application. You will at least need to specify the `joplin_token` parameter. See the **Configuration** section below for all available configuration settings.

5. Run the server
    ```
    cd $GOPATH/src/github.com/geordie/toplin
    go run main.go
    ```

## Configuration

You can specify the following configuration settings in a local `config.yaml` file.

* **joplin_token:** Your Joplin API token. Required.

* **joplin_port:** The port on which the Joplin API is running. Default: `"41184`

* **http_addr:** The host and port for Toplin. Default: `":8888"`

* **http_cert_file:** Path to cert file. Default: `""`

* **http_key_file:** Path to key file. Default: `""`

* **http_drain_interval:** How long application will wait to drain old requests before restarting. Default: `"1s"`

* **cookie_secret:** Cookie secret for session. Default: Auto generated.


## Vendoring Dependencies

Vendoring is handled by a separate project: [github.com/tools/godep](https://github.com/tools/godep).

Here's a quick tutorial on how to use it. For more details, read the readme [here](https://github.com/tools/godep#godep).
```
# Save all your dependencies after running go get ./...
godep save ./...

# Building with godep
godep go build

# Running tests with godep
godep go test ./...
```
