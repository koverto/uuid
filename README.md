# Koverto - UUID

Wraps github.com/google/uuid for use as a protobuf type and with implementations
of various serialization interfaces.

## Usage

```protobuf
import "github.com/koverto/uuid/uuid.proto";

message User {
    uuid.UUID id = 1;
}
```

## Copyright

Copyright © 2020 Jesse B. Hannah. Licensed under the [GNU GPL version 3 or later
][gpl].

[gpl]: LICENSE
