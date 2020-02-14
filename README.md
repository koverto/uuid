# Koverto - UUID

[![Go](https://github.com/koverto/uuid/workflows/Go/badge.svg)][workflow]
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fkoverto%2Fuuid.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fkoverto%2Fuuid?ref=badge_shield)

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

## License

[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fkoverto%2Fuuid.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fkoverto%2Fuuid?ref=badge_large)
