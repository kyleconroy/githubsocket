# Github Socket

Github Hooks, real-time and in the browser

## Installation

You'll need Go and the `websocket` module.

    $ go get code.google.com/p/go.net/websocket

The server looks for the following environment varaibles

```bash
$ export GITHUB_USERNAME=""
$ export GITHUB_PASSWORD=""
}
```

## Running

    $ go run server.go
     * http://localhost:23456

You are now ready for Github hooks in your browser.
