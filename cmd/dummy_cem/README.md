# Dummy CCEM

This dummy cem can be started to simulate a remote device setting some random load limits
in the LPC use case.

## Usage

```sh
go run dummy_cem/main.go <serverport> <remoteski> <certfile> <keyfile>
```

* `serverport` should be 4715
* `remoteski` is the SKI of the remote device or
    service you want to connect to
* `certfile` is a local file containing the
    generated certificate in the first usage run
* `keyfile` is a local file containing the generated
    key in the first usage run
