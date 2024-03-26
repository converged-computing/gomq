# gomq

> FairMQ in Go

This is a continued implementation of [fer](https://github.com/alice-go/fer/tree/master?tab=readme-ov-file) that is based off of [FairMQ](https://github.com/FairRootGroup/FairMQ). 
The goal is to provide a message passing interface that is not MPI. This one uses ZeroMQ under the hood.
I want something that can be used in cloud for distributed applications, and more easily than something that requires MPI.
I am going to refactor the interface to be easy to use, and then generate some benchmarks for communication
using it.

## Examples

See the [examples](examples) directory. I'm going to be implementing examples from [FairMQ](https://github.com/FairRootGroup/FairMQ/tree/master/examples) and then some benchmarks
that will give us a sense of latency, etc.

## License

The original code "fer" is released under the `BSD-3` license, which is [included here](.github/LICENSE) per directive of the license.
I give sincere thank you to the original author [Sebastian Binet](https://github.com/sbinet) that was an early advocate
for distributed applications with Go, who I first interacted with in this [GitHub issue](https://github.com/go-hep/hep/issues/1010).
We are providing now under an MIT license, which seems [OK to do](https://opensource.stackexchange.com/a/10687) as long as the original license is preserved.

HPCIC DevTools is distributed under the terms of the MIT license.
All new contributions must be made under this license.

See [LICENSE](https://github.com/converged-computing/cloud-select/blob/main/LICENSE),
[COPYRIGHT](https://github.com/converged-computing/cloud-select/blob/main/COPYRIGHT), and
[NOTICE](https://github.com/converged-computing/cloud-select/blob/main/NOTICE) for details.

SPDX-License-Identifier: (MIT)

LLNL-CODE- 842614