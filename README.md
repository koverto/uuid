# Koverto - UUID

[![Go](https://github.com/koverto/uuid/workflows/Go/badge.svg)][workflow]
[![Go Report Card](https://goreportcard.com/badge/github.com/koverto/uuid)](https://goreportcard.com/report/github.com/koverto/uuid)

Wraps [github.com/google/uuid][uuid] for use as a [protobuf][] type and with
implementations of various serialization interfaces.

## Usage

```protobuf
import "github.com/koverto/uuid/uuid.proto";

message User {
    uuid.UUID id = 1;
}
```

## Copyright

Copyright © 2020 Jesse B. Hannah. Licensed under the [GNU GPL version 3 or
later][gpl].

[gpl]: LICENSE
[protobuf]: https://developers.google.com/protocol-buffers/
[uuid]: https://github.com/google/uuid
[workflow]: https://github.com/koverto/uuid/actions?query=workflow%3AGo
