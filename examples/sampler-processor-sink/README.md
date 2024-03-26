# Sampler Processor Sink Example

This example emulates the [sampler-processor-sink](https://github.com/FairRootGroup/FairMQ/tree/master/examples/1-n-1) pattern.
This start with a config file, [sampler-processor-sink.json](sampler-processor-sink.json) that you
can validate as such:

```bash
go run ./cmd/validate/main.go ./examples/netmark/sampler-processor-sink.json 
$ echo $?
0
```

And this will demonstrate a few components working together:

 - [processor](processor): 
 - [sampler](sampler)
 - [sink](sink):


## Example

```sh
## terminal 1
$ go run sink/main.go --id sink1 --mq-config ./sampler-processor-sink.json

## terminal 2
$ go run sampler/main.go --id sampler1 --mq-config ./sampler-processor-sink.json

## terminal 3
$ go run processor/main.go --id processor --mq-config ./sampler-processor-sink.json
```

This will run 3 [devices](https://github.com/FairRootGroup/FairMQ/blob/master/docs/Device.md#1-device), using the ZeroMQ transport.

To run with `nanomsg` as a transport layer, add `--transport nanomsg` to the invocations.
