# Sampler Processor Sink Example

This example emulates the [sampler-sink](https://github.com/FairRootGroup/FairMQ/blob/master/examples/1-1/fairmq-start-ex-1-1.sh.in) pattern.
This start with a config file, [sampler-sink.json](sampler-sink.json) that you
can validate as such:

```bash
go run ./cmd/validate/main.go ./examples/sampler-to-sink/sampler-sink.json 
$ echo $?
0
```

And this will demonstrate a few components working together:

 - [sampler](sampler): generating data
 - [sink](sink): receiving it


## Example

```sh
## terminal 1
$ go run sampler/main.go --id sampler --mq-config ./sampler-sink.json

## terminal 2
$ go run sink/main.go --id sink --mq-config ./sampler-sink.json
```

This will run 2 [devices](https://github.com/FairRootGroup/FairMQ/blob/master/docs/Device.md#1-device), using the ZeroMQ transport.
To run with `nanomsg` as a transport layer, add `--transport nanomsg` to the invocations.
