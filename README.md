# Github Socket

Github Hooks, real-time and in the browser

## Installation

You'll need Go and the `websocket` module.

    $ go get code.google.com/p/go.net/websocket

`config.json` contains your Github API credentials

```json
{
  "username": "GITHUB_USERNAME",
  "password": "GITHUB_PASSWORD"
}
```

## Running

    $ go run server.go config.json
     * http://localhost:23456

You are now ready for Github hooks in your browser.
