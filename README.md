# YugabyteDB protobuf and GRPC clients

Generated from the YugabyteDB provided _proto_ files, no enterprise protobufs here.

## Regenerate

Clone YugabyteDB sources somewhere:

```sh
mkdir -p /tmp/ybdb
cd /tmp/ybdb
git clone https://github.com/yugabyte/yugabyte-db.git .
git checkout v2.7.2
```

Clone this repository:

```sh
mkdir -p ${GOPATH}/src/github.com/radekg/yugabyte-db-go-proto
cd ${GOPATH}/src/github.com/radekg/yugabyte-db-go-proto
git clone https://github.com/radekg/yugabyte-db-go-proto.git .
```

Generate:

```
make YB_SRC_DIR=/tmp/ybdb clean get-proto gen-proto
```

## Using

```go
import (
    ybApi "github.com/radekg/yugabyte-db-go-proto/v2/yb/api"
)
```

For a real world example, have a look at the reference [YugabyteDB CLI client](https://github.com/radekg/yugabyte-db-go-client/tree/master/client/cli).
