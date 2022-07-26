# wow-server-go

This project aims to create golang-based server for World of Warcraft 1.12.1. This is for education purposes only.

## Running the server

The server is built/run using [bazel](https://bazel.build).

1. `$ bazel run //:wow_server_go`
2. Change the `realmlist.wtf` file in your game client to:

```
set realmlist 127.0.0.1:5000
set patchlist 127.0.0.1:5000
```

3. Launch the game and enter the account name `test` and password `test`.

## Running the tests

Tests are also managed using bazel.

- `$ bazel test //...`
