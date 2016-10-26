# go-cabot
[![GoDoc](https://godoc.org/github.com/mrsaints/go-cabot/cabot?status.svg)](https://godoc.org/github.com/mrsaints/go-cabot/cabot)

An idiomatic Go client library for interacting with [Arachnys'][arachnys] [Cabot][cabot] API.

_Cabot is a self-hosted, easily-deployable monitoring, and alerts service - like a lightweight PagerDuty._


## TODO

- [ ] Less boilerplate-y
- [ ] Tests (???)
- [ ] More examples


## Usage

1. Download, and install `go-cabot/cabot`:

    ```shell
    go get github.com/mrsaints/go-cabot/cabot
    ```

2. Import the package into your code:

    ```go
    import "github.com/mrsaints/go-cabot/cabot"
    ```

3. Construct a Cabot API client:

    ```go
    auth := cabot.WithBasicAuth("username", "password")
    baseURL := cabot.WithBaseURL("https://your-base-cabot-url/")

    // Accepts a variable number of `Option` arguments
    client = cabot.NewClient(auth, baseURL)
    ```

4. Using the newly constructed client, call methods via its respective service:

    ```go
    // Returns a list of services
    client.Services.List()

    // Returns a specific instance using its ID
    client.Instances.Get(1)

    // Creates a new ICMP check with "Hello World!" as its name
    check := &cabot.ICMPCheck{cabot.StatusCheck{Name: "Hello World!"}}
    created, _ := client.ICMPChecks.Create(check)
    log.Println(created)
    ```


[arachnys]: https://www.arachnys.com/
[cabot]: https://github.com/arachnys/cabot
