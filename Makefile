YB_SRC_DIR?=/tmp/ybdb
CURRENT_DIR=$(dir $(realpath $(firstword $(MAKEFILE_LIST))))
# do not add v2 here, the directory structure stays untouched
GO_PACKAGE=github.com/radekg/yugabyte-db-go-proto

.PHONY: clean
clean:
	rm -rf ${CURRENT_DIR}/proto/*
	rm -rf ${CURRENT_DIR}/yb

.PHONY: get-proto
get-proto:
	YB_SRC_DIR=${YB_SRC_DIR} ${CURRENT_DIR}/bin/get-proto.sh

.PHONY: gen-proto
gen-proto:
	protoc --proto_path=./proto \
		--go_opt=Myb/cdc/cdc_consumer.proto=${GO_PACKAGE}/yb/api \
		--go_opt=Myb/cdc/cdc_service.proto=${GO_PACKAGE}/yb/api \
		--go_opt=Myb/common/common.proto=${GO_PACKAGE}/yb/api \
		--go_opt=Myb/common/pgsql_protocol.proto=${GO_PACKAGE}/yb/api \
		--go_opt=Myb/common/ql_protocol.proto=${GO_PACKAGE}/yb/api \
		--go_opt=Myb/common/redis_protocol.proto=${GO_PACKAGE}/yb/api \
		--go_opt=Myb/common/wire_protocol.proto=${GO_PACKAGE}/yb/api \
		--go_opt=Myb/consensus/consensus.proto=${GO_PACKAGE}/yb/api \
		--go_opt=Myb/consensus/consensus_metadata.proto=${GO_PACKAGE}/yb/api \
		--go_opt=Myb/consensus/log.proto=${GO_PACKAGE}/yb/api \
		--go_opt=Myb/docdb/docdb.proto=${GO_PACKAGE}/yb/api \
		--go_opt=Myb/fs/fs.proto=${GO_PACKAGE}/yb/api \
		--go_opt=Myb/master/master.proto=${GO_PACKAGE}/yb/api \
		--go_opt=Myb/rocksdb/db/version_edit.proto=${GO_PACKAGE}/yb/api \
		--go_opt=Myb/rpc/rpc_header.proto=${GO_PACKAGE}/yb/api \
		--go_opt=Myb/rpc/rpc_introspection.proto=${GO_PACKAGE}/yb/api \
		--go_opt=Myb/rpc/rtest_diff_package.proto=${GO_PACKAGE}/yb/api \
		--go_opt=Myb/rpc/rtest.proto=${GO_PACKAGE}/yb/api \
		--go_opt=Myb/server/server_base.proto=${GO_PACKAGE}/yb/api \
		--go_opt=Myb/tablet/tablet_metadata.proto=${GO_PACKAGE}/yb/api \
		--go_opt=Myb/tablet/tablet.proto=${GO_PACKAGE}/yb/api \
		--go_opt=Myb/tserver/backup.proto=${GO_PACKAGE}/yb/api \
		--go_opt=Myb/tserver/remote_bootstrap.proto=${GO_PACKAGE}/yb/api \
		--go_opt=Myb/tserver/tserver_admin.proto=${GO_PACKAGE}/yb/api \
		--go_opt=Myb/tserver/tserver_forward_service.proto=${GO_PACKAGE}/yb/api \
		--go_opt=Myb/tserver/tserver_service.proto=${GO_PACKAGE}/yb/api \
		--go_opt=Myb/tserver/tserver.proto=${GO_PACKAGE}/yb/api \
		--go_opt=Myb/util/encryption.proto=${GO_PACKAGE}/yb/api \
		--go_opt=Myb/util/histogram.proto=${GO_PACKAGE}/yb/api \
		--go_opt=Myb/util/jsonwriter_test.proto=${GO_PACKAGE}/yb/api \
		--go_opt=Myb/util/opid.proto=${GO_PACKAGE}/yb/api \
		--go_opt=Myb/util/pb_util.proto=${GO_PACKAGE}/yb/api \
		--go_opt=Myb/util/proto_container_test.proto=${GO_PACKAGE}/yb/api \
		--go_opt=Myb/util/proto_container_test2.proto=${GO_PACKAGE}/yb/api \
		--go_opt=Myb/util/proto_container_test3.proto=${GO_PACKAGE}/yb/api \
		--go_opt=Myb/util/version_info.proto=${GO_PACKAGE}/yb/api \
		--go_opt=Myb/yql/cql/cqlserver/cql_service.proto=${GO_PACKAGE}/yb/api/yql/cql/cqlserver \
		--go_opt=Myb/yql/redis/redisserver/redis_service.proto=${GO_PACKAGE}/yb/api/yql/redis/redisserver \
		--go-grpc_opt=Myb/cdc/cdc_consumer.proto=${GO_PACKAGE}/yb/api \
		--go-grpc_opt=Myb/cdc/cdc_service.proto=${GO_PACKAGE}/yb/api \
		--go-grpc_opt=Myb/common/common.proto=${GO_PACKAGE}/yb/api \
		--go-grpc_opt=Myb/common/pgsql_protocol.proto=${GO_PACKAGE}/yb/api \
		--go-grpc_opt=Myb/common/ql_protocol.proto=${GO_PACKAGE}/yb/api \
		--go-grpc_opt=Myb/common/redis_protocol.proto=${GO_PACKAGE}/yb/api \
		--go-grpc_opt=Myb/common/wire_protocol.proto=${GO_PACKAGE}/yb/api \
		--go-grpc_opt=Myb/consensus/consensus.proto=${GO_PACKAGE}/yb/api \
		--go-grpc_opt=Myb/consensus/consensus_metadata.proto=${GO_PACKAGE}/yb/api \
		--go-grpc_opt=Myb/consensus/log.proto=${GO_PACKAGE}/yb/api \
		--go-grpc_opt=Myb/docdb/docdb.proto=${GO_PACKAGE}/yb/api \
		--go-grpc_opt=Myb/fs/fs.proto=${GO_PACKAGE}/yb/api \
		--go-grpc_opt=Myb/master/master.proto=${GO_PACKAGE}/yb/api \
		--go-grpc_opt=Myb/rocksdb/db/version_edit.proto=${GO_PACKAGE}/yb/api \
		--go-grpc_opt=Myb/rpc/rpc_header.proto=${GO_PACKAGE}/yb/api \
		--go-grpc_opt=Myb/rpc/rpc_introspection.proto=${GO_PACKAGE}/yb/api \
		--go-grpc_opt=Myb/rpc/rtest_diff_package.proto=${GO_PACKAGE}/yb/api \
		--go-grpc_opt=Myb/rpc/rtest.proto=${GO_PACKAGE}/yb/api \
		--go-grpc_opt=Myb/server/server_base.proto=${GO_PACKAGE}/yb/api \
		--go-grpc_opt=Myb/tablet/tablet_metadata.proto=${GO_PACKAGE}/yb/api \
		--go-grpc_opt=Myb/tablet/tablet.proto=${GO_PACKAGE}/yb/api \
		--go-grpc_opt=Myb/tserver/backup.proto=${GO_PACKAGE}/yb/api \
		--go-grpc_opt=Myb/tserver/remote_bootstrap.proto=${GO_PACKAGE}/yb/api \
		--go-grpc_opt=Myb/tserver/tserver_admin.proto=${GO_PACKAGE}/yb/api \
		--go-grpc_opt=Myb/tserver/tserver_forward_service.proto=${GO_PACKAGE}/yb/api \
		--go-grpc_opt=Myb/tserver/tserver_service.proto=${GO_PACKAGE}/yb/api \
		--go-grpc_opt=Myb/tserver/tserver.proto=${GO_PACKAGE}/yb/api \
		--go-grpc_opt=Myb/util/encryption.proto=${GO_PACKAGE}/yb/api \
		--go-grpc_opt=Myb/util/histogram.proto=${GO_PACKAGE}/yb/api \
		--go-grpc_opt=Myb/util/jsonwriter_test.proto=${GO_PACKAGE}/yb/api \
		--go-grpc_opt=Myb/util/opid.proto=${GO_PACKAGE}/yb/api \
		--go-grpc_opt=Myb/util/pb_util.proto=${GO_PACKAGE}/yb/api \
		--go-grpc_opt=Myb/util/proto_container_test.proto=${GO_PACKAGE}/yb/api \
		--go-grpc_opt=Myb/util/proto_container_test2.proto=${GO_PACKAGE}/yb/api \
		--go-grpc_opt=Myb/util/proto_container_test3.proto=${GO_PACKAGE}/yb/api \
		--go-grpc_opt=Myb/util/version_info.proto=${GO_PACKAGE}/yb/api \
		--go-grpc_opt=Myb/yql/cql/cqlserver/cql_service.proto=${GO_PACKAGE}/yb/api/yql/cql/cqlserver \
		--go-grpc_opt=Myb/yql/redis/redisserver/redis_service.proto=${GO_PACKAGE}/yb/api/yql/redis/redisserver \
		--go_out=${GOPATH}/src/ \
		--go-grpc_out=${GOPATH}/src/ \
		--go-grpc_opt=require_unimplemented_servers=false \
		$$(find proto/yb -iname "*.proto")
	$(MAKE) fix-consensus-types fix-rtest-types fix-master-types fix-cdc-service-types

.PHONY: fix-consensus-types
fix-consensus-types:
	sed -i.bak -e 's!NoOpRequestPB!ConsensusNoOpRequestPB!g' ${GOPATH}/src/${GO_PACKAGE}/yb/api/consensus.pb.go \
		&& rm -- ${GOPATH}/src/${GO_PACKAGE}/yb/api/consensus.pb.go.bak
	sed -i.bak -e 's!NoOpResponsePB!ConsensusNoOpResponsePB!g' ${GOPATH}/src/${GO_PACKAGE}/yb/api/consensus.pb.go \
		&& rm -- ${GOPATH}/src/${GO_PACKAGE}/yb/api/consensus.pb.go.bak

.PHONY: fix-rtest-types
fix-rtest-types:
	sed -i.bak -e 's!PingRequestPB!RTestPingRequestPB!g' ${GOPATH}/src/${GO_PACKAGE}/yb/api/rtest_grpc.pb.go \
		&& rm -- ${GOPATH}/src/${GO_PACKAGE}/yb/api/rtest_grpc.pb.go.bak
	sed -i.bak -e 's!PingResponsePB!RTestPingResponsePB!g' ${GOPATH}/src/${GO_PACKAGE}/yb/api/rtest_grpc.pb.go \
		&& rm -- ${GOPATH}/src/${GO_PACKAGE}/yb/api/rtest_grpc.pb.go.bak
	sed -i.bak -e 's!PingRequestPB!RTestPingRequestPB!g' ${GOPATH}/src/${GO_PACKAGE}/yb/api/rtest.pb.go \
		&& rm -- ${GOPATH}/src/${GO_PACKAGE}/yb/api/rtest.pb.go.bak
	sed -i.bak -e 's!PingResponsePB!RTestPingResponsePB!g' ${GOPATH}/src/${GO_PACKAGE}/yb/api/rtest.pb.go \
		&& rm -- ${GOPATH}/src/${GO_PACKAGE}/yb/api/rtest.pb.go.bak

.PHONY: fix-master-types
fix-master-types:
	sed -i.bak -e 's!CreateCDCStreamRequestPB!MasterCreateCDCStreamRequestPB!g' ${GOPATH}/src/${GO_PACKAGE}/yb/api/master.pb.go \
		&& rm -- ${GOPATH}/src/${GO_PACKAGE}/yb/api/master.pb.go.bak
	sed -i.bak -e 's!CreateCDCStreamResponsePB!MasterCreateCDCStreamResponsePB!g' ${GOPATH}/src/${GO_PACKAGE}/yb/api/master.pb.go \
		&& rm -- ${GOPATH}/src/${GO_PACKAGE}/yb/api/master.pb.go.bak
	sed -i.bak -e 's!DeleteCDCStreamRequestPB!MasterDeleteCDCStreamRequestPB!g' ${GOPATH}/src/${GO_PACKAGE}/yb/api/master.pb.go \
		&& rm -- ${GOPATH}/src/${GO_PACKAGE}/yb/api/master.pb.go.bak
	sed -i.bak -e 's!DeleteCDCStreamResponsePB!MasterDeleteCDCStreamResponsePB!g' ${GOPATH}/src/${GO_PACKAGE}/yb/api/master.pb.go \
		&& rm -- ${GOPATH}/src/${GO_PACKAGE}/yb/api/master.pb.go.bak
	sed -i.bak -e 's!BackfillIndexRequestPB!MasterBackfillIndexRequestPB!g' ${GOPATH}/src/${GO_PACKAGE}/yb/api/master.pb.go \
		&& rm -- ${GOPATH}/src/${GO_PACKAGE}/yb/api/master.pb.go.bak
	sed -i.bak -e 's!BackfillIndexResponsePB!MasterBackfillIndexResponsePB!g' ${GOPATH}/src/${GO_PACKAGE}/yb/api/master.pb.go \
		&& rm -- ${GOPATH}/src/${GO_PACKAGE}/yb/api/master.pb.go.bak
	sed -i.bak -e 's!DeleteTabletRequestPB!MasterDeleteTabletRequestPB!g' ${GOPATH}/src/${GO_PACKAGE}/yb/api/master.pb.go \
		&& rm -- ${GOPATH}/src/${GO_PACKAGE}/yb/api/master.pb.go.bak
	sed -i.bak -e 's!DeleteTabletResponsePB!MasterDeleteTabletResponsePB!g' ${GOPATH}/src/${GO_PACKAGE}/yb/api/master.pb.go \
		&& rm -- ${GOPATH}/src/${GO_PACKAGE}/yb/api/master.pb.go.bak
	sed -i.bak -e 's!SplitTabletRequestPB!MasterSplitTabletRequestPB!g' ${GOPATH}/src/${GO_PACKAGE}/yb/api/master.pb.go \
		&& rm -- ${GOPATH}/src/${GO_PACKAGE}/yb/api/master.pb.go.bak
	sed -i.bak -e 's!SplitTabletResponsePB!MasterSplitTabletResponsePB!g' ${GOPATH}/src/${GO_PACKAGE}/yb/api/master.pb.go \
		&& rm -- ${GOPATH}/src/${GO_PACKAGE}/yb/api/master.pb.go.bak
	sed -i.bak -e 's!CreateCDCStreamRequestPB!MasterCreateCDCStreamRequestPB!g' ${GOPATH}/src/${GO_PACKAGE}/yb/api/master_grpc.pb.go \
		&& rm -- ${GOPATH}/src/${GO_PACKAGE}/yb/api/master_grpc.pb.go.bak
	sed -i.bak -e 's!CreateCDCStreamResponsePB!MasterCreateCDCStreamResponsePB!g' ${GOPATH}/src/${GO_PACKAGE}/yb/api/master_grpc.pb.go \
		&& rm -- ${GOPATH}/src/${GO_PACKAGE}/yb/api/master_grpc.pb.go.bak
	sed -i.bak -e 's!DeleteCDCStreamRequestPB!MasterDeleteCDCStreamRequestPB!g' ${GOPATH}/src/${GO_PACKAGE}/yb/api/master_grpc.pb.go \
		&& rm -- ${GOPATH}/src/${GO_PACKAGE}/yb/api/master_grpc.pb.go.bak
	sed -i.bak -e 's!DeleteCDCStreamResponsePB!MasterDeleteCDCStreamResponsePB!g' ${GOPATH}/src/${GO_PACKAGE}/yb/api/master_grpc.pb.go \
		&& rm -- ${GOPATH}/src/${GO_PACKAGE}/yb/api/master_grpc.pb.go.bak
	sed -i.bak -e 's!BackfillIndexRequestPB!MasterBackfillIndexRequestPB!g' ${GOPATH}/src/${GO_PACKAGE}/yb/api/master_grpc.pb.go \
		&& rm -- ${GOPATH}/src/${GO_PACKAGE}/yb/api/master_grpc.pb.go.bak
	sed -i.bak -e 's!BackfillIndexResponsePB!MasterBackfillIndexResponsePB!g' ${GOPATH}/src/${GO_PACKAGE}/yb/api/master_grpc.pb.go \
		&& rm -- ${GOPATH}/src/${GO_PACKAGE}/yb/api/master_grpc.pb.go.bak
	sed -i.bak -e 's!DeleteTabletRequestPB!MasterDeleteTabletRequestPB!g' ${GOPATH}/src/${GO_PACKAGE}/yb/api/master_grpc.pb.go \
		&& rm -- ${GOPATH}/src/${GO_PACKAGE}/yb/api/master_grpc.pb.go.bak
	sed -i.bak -e 's!DeleteTabletResponsePB!MasterDeleteTabletResponsePB!g' ${GOPATH}/src/${GO_PACKAGE}/yb/api/master_grpc.pb.go \
		&& rm -- ${GOPATH}/src/${GO_PACKAGE}/yb/api/master_grpc.pb.go.bak
	sed -i.bak -e 's!SplitTabletRequestPB!MasterSplitTabletRequestPB!g' ${GOPATH}/src/${GO_PACKAGE}/yb/api/master_grpc.pb.go \
		&& rm -- ${GOPATH}/src/${GO_PACKAGE}/yb/api/master_grpc.pb.go.bak
	sed -i.bak -e 's!SplitTabletResponsePB!MasterSplitTabletResponsePB!g' ${GOPATH}/src/${GO_PACKAGE}/yb/api/master_grpc.pb.go \
		&& rm -- ${GOPATH}/src/${GO_PACKAGE}/yb/api/master_grpc.pb.go.bak

.PHONY: fix-cdc-service-types
fix-cdc-service-types:
	sed -i.bak -e 's!ListTabletsRequestPB!CDCListTabletsRequestPB!g' ${GOPATH}/src/${GO_PACKAGE}/yb/api/cdc_service.pb.go \
		&& rm -- ${GOPATH}/src/${GO_PACKAGE}/yb/api/cdc_service.pb.go.bak
	sed -i.bak -e 's!ListTabletsResponsePB!CDCListTabletsResponsePB!g' ${GOPATH}/src/${GO_PACKAGE}/yb/api/cdc_service.pb.go \
		&& rm -- ${GOPATH}/src/${GO_PACKAGE}/yb/api/cdc_service.pb.go.bak
	sed -i.bak -e 's!KeyValuePairPB!CDCKeyValuePairPB!g' ${GOPATH}/src/${GO_PACKAGE}/yb/api/cdc_service.pb.go \
		&& rm -- ${GOPATH}/src/${GO_PACKAGE}/yb/api/cdc_service.pb.go.bak
	sed -i.bak -e 's!ListTabletsRequestPB!CDCListTabletsRequestPB!g' ${GOPATH}/src/${GO_PACKAGE}/yb/api/cdc_service_grpc.pb.go \
		&& rm -- ${GOPATH}/src/${GO_PACKAGE}/yb/api/cdc_service_grpc.pb.go.bak
	sed -i.bak -e 's!ListTabletsResponsePB!CDCListTabletsResponsePB!g' ${GOPATH}/src/${GO_PACKAGE}/yb/api/cdc_service_grpc.pb.go \
		&& rm -- ${GOPATH}/src/${GO_PACKAGE}/yb/api/cdc_service_grpc.pb.go.bak
	sed -i.bak -e 's!KeyValuePairPB!CDCKeyValuePairPB!g' ${GOPATH}/src/${GO_PACKAGE}/yb/api/cdc_service_grpc.pb.go \
		&& rm -- ${GOPATH}/src/${GO_PACKAGE}/yb/api/cdc_service_grpc.pb.go.bak

.PHONY: prototools
prototools:
	go get -u github.com/golang/protobuf/proto \
		github.com/golang/protobuf/protoc-gen-go \
		google.golang.org/grpc \
		google.golang.org/protobuf/cmd/protoc-gen-go \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go install google.golang.org/protobuf/cmd/protoc-gen-go \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go mod tidy
