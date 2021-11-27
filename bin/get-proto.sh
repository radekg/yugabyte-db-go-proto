#!/bin/bash

proto_dir="${GOPATH}/src/github.com/radekg/yugabyte-db-go-proto/proto"

for f in `find ${YB_SRC_DIR} -name *.proto | grep src/yb | grep -v /ent`; do
	relative=${f#$YB_SRC_DIR/src/}
	target="${proto_dir}"/$(dirname $relative)
	mkdir -p "${target}"
	cp -v "${f}" "${target}"/$(basename $f)
done

# Include /ent directory.
# Per the license: The entire database with all its features (including the enterprise ones) are licensed under the Apache License 2.0.
for f in `find ${YB_SRC_DIR} -name *.proto | grep ent/src/yb`; do
	relative=${f#$YB_SRC_DIR/ent/src/}
	target="${proto_dir}"/$(dirname $relative)
	mkdir -p "${target}"
	cp -v "${f}" "${target}"/$(basename $f)
done

mv "${GOPATH}/src/github.com/radekg/yugabyte-db-go-proto/proto/yb/consensus/metadata.proto" \
	"${GOPATH}/src/github.com/radekg/yugabyte-db-go-proto/proto/yb/consensus/consensus_metadata.proto"

for f in `grep -rnw yb/consensus/metadata.proto ${proto_dir} | awk -F  ":" '{print $1}'`; do
	sed -i.bak -e 's!yb/consensus/metadata.proto!yb/consensus/consensus_metadata.proto!g' "${f}" \
		&& rm -- "${f}.bak"
done

mv "${GOPATH}/src/github.com/radekg/yugabyte-db-go-proto/proto/yb/tablet/metadata.proto" \
	"${GOPATH}/src/github.com/radekg/yugabyte-db-go-proto/proto/yb/tablet/tablet_metadata.proto"

for f in `grep -rnw yb/tablet/metadata.proto ${proto_dir} | awk -F  ":" '{print $1}'`; do
	sed -i.bak -e 's!yb/tablet/metadata.proto!yb/tablet/tablet_metadata.proto!' "${f}" \
		&& rm -- "${f}.bak"
done
