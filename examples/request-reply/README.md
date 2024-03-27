# Request Reply Example

This example emulates the [request-reply](https://github.com/FairRootGroup/FairMQ/tree/master/examples/req-rep) pattern.
This start with a config file, [request-reply.json](request-reply.json) that you can validate as such:

```bash
go run ./cmd/validate/main.go ./examples/request-reply/request-reply.json
$ echo $?
0
```

## Example

```sh
## terminal 1
$ go run server/main.go --id server --mq-config ./request-reply.json

## terminal 2
$ go run client/main.go --id client --mq-config ./request-reply.json
```

This will run 2 [devices](https://github.com/FairRootGroup/FairMQ/blob/master/docs/Device.md#1-device), using the ZeroMQ transport.
To run with `nanomsg` as a transport layer, add `--transport nanomsg` to the invocations.
