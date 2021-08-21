.PHONY: genproto
genproto:
	protoc --proto_path=./proto \
		--go_opt=Myb/cdc/cdc_consumer.proto=github.com/radekg/yugabyte-db-go-proto/yb/api \
		--go_opt=Myb/cdc/cdc_service.proto=github.com/radekg/yugabyte-db-go-proto/yb/api \
		--go_opt=Myb/common/common.proto=github.com/radekg/yugabyte-db-go-proto/yb/api \
		--go_opt=Myb/common/pgsql_protocol.proto=github.com/radekg/yugabyte-db-go-proto/yb/api \
		--go_opt=Myb/common/ql_protocol.proto=github.com/radekg/yugabyte-db-go-proto/yb/api \
		--go_opt=Myb/common/redis_protocol.proto=github.com/radekg/yugabyte-db-go-proto/yb/api \
		--go_opt=Myb/common/wire_protocol.proto=github.com/radekg/yugabyte-db-go-proto/yb/api \
		--go_opt=Myb/consensus/consensus.proto=github.com/radekg/yugabyte-db-go-proto/yb/api \
		--go_opt=Myb/consensus/consensus_metadata.proto=github.com/radekg/yugabyte-db-go-proto/yb/api \
		--go_opt=Myb/consensus/log.proto=github.com/radekg/yugabyte-db-go-proto/yb/api \
		--go_opt=Myb/docdb/docdb.proto=github.com/radekg/yugabyte-db-go-proto/yb/api \
		--go_opt=Myb/fs/fs.proto=github.com/radekg/yugabyte-db-go-proto/yb/api \
		--go_opt=Myb/master/master.proto=github.com/radekg/yugabyte-db-go-proto/yb/api \
		--go_opt=Myb/rocksdb/db/version_edit.proto=github.com/radekg/yugabyte-db-go-proto/yb/api \
		--go_opt=Myb/rpc/rpc_header.proto=github.com/radekg/yugabyte-db-go-proto/yb/api \
		--go_opt=Myb/rpc/rpc_introspection.proto=github.com/radekg/yugabyte-db-go-proto/yb/api \
		--go_opt=Myb/rpc/rtest_diff_package.proto=github.com/radekg/yugabyte-db-go-proto/yb/api \
		--go_opt=Myb/rpc/rtest.proto=github.com/radekg/yugabyte-db-go-proto/yb/api \
		--go_opt=Myb/server/server_base.proto=github.com/radekg/yugabyte-db-go-proto/yb/api \
		--go_opt=Myb/tablet/tablet_metadata.proto=github.com/radekg/yugabyte-db-go-proto/yb/api \
		--go_opt=Myb/tablet/tablet.proto=github.com/radekg/yugabyte-db-go-proto/yb/api \
		--go_opt=Myb/tserver/backup.proto=github.com/radekg/yugabyte-db-go-proto/yb/api \
		--go_opt=Myb/tserver/remote_bootstrap.proto=github.com/radekg/yugabyte-db-go-proto/yb/api \
		--go_opt=Myb/tserver/tserver_admin.proto=github.com/radekg/yugabyte-db-go-proto/yb/api \
		--go_opt=Myb/tserver/tserver_forward_service.proto=github.com/radekg/yugabyte-db-go-proto/yb/api \
		--go_opt=Myb/tserver/tserver_service.proto=github.com/radekg/yugabyte-db-go-proto/yb/api \
		--go_opt=Myb/tserver/tserver.proto=github.com/radekg/yugabyte-db-go-proto/yb/api \
		--go_opt=Myb/util/encryption.proto=github.com/radekg/yugabyte-db-go-proto/yb/api \
		--go_opt=Myb/util/histogram.proto=github.com/radekg/yugabyte-db-go-proto/yb/api \
		--go_opt=Myb/util/jsonwriter_test.proto=github.com/radekg/yugabyte-db-go-proto/yb/api \
		--go_opt=Myb/util/opid.proto=github.com/radekg/yugabyte-db-go-proto/yb/api \
		--go_opt=Myb/util/pb_util.proto=github.com/radekg/yugabyte-db-go-proto/yb/api \
		--go_opt=Myb/util/proto_container_test.proto=github.com/radekg/yugabyte-db-go-proto/yb/api \
		--go_opt=Myb/util/proto_container_test2.proto=github.com/radekg/yugabyte-db-go-proto/yb/api \
		--go_opt=Myb/util/proto_container_test3.proto=github.com/radekg/yugabyte-db-go-proto/yb/api \
		--go_opt=Myb/util/version_info.proto=github.com/radekg/yugabyte-db-go-proto/yb/api \
		--go_opt=Myb/yql/cql/cqlserver/cql_service.proto=github.com/radekg/yugabyte-db-go-proto/yb/api/yql/cql/cqlserver \
		--go_opt=Myb/yql/redis/redisserver/redis_service.proto=github.com/radekg/yugabyte-db-go-proto/yb/api/yql/redis/redisserver \
		--go-grpc_opt=Myb/cdc/cdc_consumer.proto=github.com/radekg/yugabyte-db-go-proto/yb/api \
		--go-grpc_opt=Myb/cdc/cdc_service.proto=github.com/radekg/yugabyte-db-go-proto/yb/api \
		--go-grpc_opt=Myb/common/common.proto=github.com/radekg/yugabyte-db-go-proto/yb/api \
		--go-grpc_opt=Myb/common/pgsql_protocol.proto=github.com/radekg/yugabyte-db-go-proto/yb/api \
		--go-grpc_opt=Myb/common/ql_protocol.proto=github.com/radekg/yugabyte-db-go-proto/yb/api \
		--go-grpc_opt=Myb/common/redis_protocol.proto=github.com/radekg/yugabyte-db-go-proto/yb/api \
		--go-grpc_opt=Myb/common/wire_protocol.proto=github.com/radekg/yugabyte-db-go-proto/yb/api \
		--go-grpc_opt=Myb/consensus/consensus.proto=github.com/radekg/yugabyte-db-go-proto/yb/api \
		--go-grpc_opt=Myb/consensus/consensus_metadata.proto=github.com/radekg/yugabyte-db-go-proto/yb/api \
		--go-grpc_opt=Myb/consensus/log.proto=github.com/radekg/yugabyte-db-go-proto/yb/api \
		--go-grpc_opt=Myb/docdb/docdb.proto=github.com/radekg/yugabyte-db-go-proto/yb/api \
		--go-grpc_opt=Myb/fs/fs.proto=github.com/radekg/yugabyte-db-go-proto/yb/api \
		--go-grpc_opt=Myb/master/master.proto=github.com/radekg/yugabyte-db-go-proto/yb/api \
		--go-grpc_opt=Myb/rocksdb/db/version_edit.proto=github.com/radekg/yugabyte-db-go-proto/yb/api \
		--go-grpc_opt=Myb/rpc/rpc_header.proto=github.com/radekg/yugabyte-db-go-proto/yb/api \
		--go-grpc_opt=Myb/rpc/rpc_introspection.proto=github.com/radekg/yugabyte-db-go-proto/yb/api \
		--go-grpc_opt=Myb/rpc/rtest_diff_package.proto=github.com/radekg/yugabyte-db-go-proto/yb/api \
		--go-grpc_opt=Myb/rpc/rtest.proto=github.com/radekg/yugabyte-db-go-proto/yb/api \
		--go-grpc_opt=Myb/server/server_base.proto=github.com/radekg/yugabyte-db-go-proto/yb/api \
		--go-grpc_opt=Myb/tablet/tablet_metadata.proto=github.com/radekg/yugabyte-db-go-proto/yb/api \
		--go-grpc_opt=Myb/tablet/tablet.proto=github.com/radekg/yugabyte-db-go-proto/yb/api \
		--go-grpc_opt=Myb/tserver/backup.proto=github.com/radekg/yugabyte-db-go-proto/yb/api \
		--go-grpc_opt=Myb/tserver/remote_bootstrap.proto=github.com/radekg/yugabyte-db-go-proto/yb/api \
		--go-grpc_opt=Myb/tserver/tserver_admin.proto=github.com/radekg/yugabyte-db-go-proto/yb/api \
		--go-grpc_opt=Myb/tserver/tserver_forward_service.proto=github.com/radekg/yugabyte-db-go-proto/yb/api \
		--go-grpc_opt=Myb/tserver/tserver_service.proto=github.com/radekg/yugabyte-db-go-proto/yb/api \
		--go-grpc_opt=Myb/tserver/tserver.proto=github.com/radekg/yugabyte-db-go-proto/yb/api \
		--go-grpc_opt=Myb/util/encryption.proto=github.com/radekg/yugabyte-db-go-proto/yb/api \
		--go-grpc_opt=Myb/util/histogram.proto=github.com/radekg/yugabyte-db-go-proto/yb/api \
		--go-grpc_opt=Myb/util/jsonwriter_test.proto=github.com/radekg/yugabyte-db-go-proto/yb/api \
		--go-grpc_opt=Myb/util/opid.proto=github.com/radekg/yugabyte-db-go-proto/yb/api \
		--go-grpc_opt=Myb/util/pb_util.proto=github.com/radekg/yugabyte-db-go-proto/yb/api \
		--go-grpc_opt=Myb/util/proto_container_test.proto=github.com/radekg/yugabyte-db-go-proto/yb/api \
		--go-grpc_opt=Myb/util/proto_container_test2.proto=github.com/radekg/yugabyte-db-go-proto/yb/api \
		--go-grpc_opt=Myb/util/proto_container_test3.proto=github.com/radekg/yugabyte-db-go-proto/yb/api \
		--go-grpc_opt=Myb/util/version_info.proto=github.com/radekg/yugabyte-db-go-proto/yb/api \
		--go-grpc_opt=Myb/yql/cql/cqlserver/cql_service.proto=github.com/radekg/yugabyte-db-go-proto/yb/api/yql/cql/cqlserver \
		--go-grpc_opt=Myb/yql/redis/redisserver/redis_service.proto=github.com/radekg/yugabyte-db-go-proto/yb/api/yql/redis/redisserver \
		--go_out=${GOPATH}/src/ \
		--go-grpc_out=${GOPATH}/src/ \
		--go-grpc_opt=require_unimplemented_servers=false \
		$$(find proto/yb -iname "*.proto")
	
	$(MAKE) fix-consensus-types fix-rtest-types fix-master-types fix-cdc-service-types

.PHONY: fix-consensus-types
fix-consensus-types:
	sed -i.bak -e 's!NoOpRequestPB!ConsensusNoOpRequestPB!g' ${GOPATH}/src/github.com/radekg/yugabyte-db-go-proto/yb/api/consensus.pb.go \
		&& rm -- ${GOPATH}/src/github.com/radekg/yugabyte-db-go-proto/yb/api/consensus.pb.go.bak
	sed -i.bak -e 's!NoOpResponsePB!ConsensusNoOpResponsePB!g' ${GOPATH}/src/github.com/radekg/yugabyte-db-go-proto/yb/api/consensus.pb.go \
		&& rm -- ${GOPATH}/src/github.com/radekg/yugabyte-db-go-proto/yb/api/consensus.pb.go.bak

.PHONY: fix-rtest-types
fix-rtest-types:
	sed -i.bak -e 's!PingRequestPB!RTestPingRequestPB!g' ${GOPATH}/src/github.com/radekg/yugabyte-db-go-proto/yb/api/rtest_grpc.pb.go \
		&& rm -- ${GOPATH}/src/github.com/radekg/yugabyte-db-go-proto/yb/api/rtest_grpc.pb.go.bak
	sed -i.bak -e 's!PingResponsePB!RTestPingResponsePB!g' ${GOPATH}/src/github.com/radekg/yugabyte-db-go-proto/yb/api/rtest_grpc.pb.go \
		&& rm -- ${GOPATH}/src/github.com/radekg/yugabyte-db-go-proto/yb/api/rtest_grpc.pb.go.bak
	sed -i.bak -e 's!PingRequestPB!RTestPingRequestPB!g' ${GOPATH}/src/github.com/radekg/yugabyte-db-go-proto/yb/api/rtest.pb.go \
		&& rm -- ${GOPATH}/src/github.com/radekg/yugabyte-db-go-proto/yb/api/rtest.pb.go.bak
	sed -i.bak -e 's!PingResponsePB!RTestPingResponsePB!g' ${GOPATH}/src/github.com/radekg/yugabyte-db-go-proto/yb/api/rtest.pb.go \
		&& rm -- ${GOPATH}/src/github.com/radekg/yugabyte-db-go-proto/yb/api/rtest.pb.go.bak

.PHONY: fix-master-types
fix-master-types:
	sed -i.bak -e 's!CreateCDCStreamRequestPB!MasterCreateCDCStreamRequestPB!g' ${GOPATH}/src/github.com/radekg/yugabyte-db-go-proto/yb/api/master.pb.go \
		&& rm -- ${GOPATH}/src/github.com/radekg/yugabyte-db-go-proto/yb/api/master.pb.go.bak
	sed -i.bak -e 's!CreateCDCStreamResponsePB!MasterCreateCDCStreamResponsePB!g' ${GOPATH}/src/github.com/radekg/yugabyte-db-go-proto/yb/api/master.pb.go \
		&& rm -- ${GOPATH}/src/github.com/radekg/yugabyte-db-go-proto/yb/api/master.pb.go.bak
	sed -i.bak -e 's!DeleteCDCStreamRequestPB!MasterDeleteCDCStreamRequestPB!g' ${GOPATH}/src/github.com/radekg/yugabyte-db-go-proto/yb/api/master.pb.go \
		&& rm -- ${GOPATH}/src/github.com/radekg/yugabyte-db-go-proto/yb/api/master.pb.go.bak
	sed -i.bak -e 's!DeleteCDCStreamResponsePB!MasterDeleteCDCStreamResponsePB!g' ${GOPATH}/src/github.com/radekg/yugabyte-db-go-proto/yb/api/master.pb.go \
		&& rm -- ${GOPATH}/src/github.com/radekg/yugabyte-db-go-proto/yb/api/master.pb.go.bak
	sed -i.bak -e 's!BackfillIndexRequestPB!MasterBackfillIndexRequestPB!g' ${GOPATH}/src/github.com/radekg/yugabyte-db-go-proto/yb/api/master.pb.go \
		&& rm -- ${GOPATH}/src/github.com/radekg/yugabyte-db-go-proto/yb/api/master.pb.go.bak
	sed -i.bak -e 's!BackfillIndexResponsePB!MasterBackfillIndexResponsePB!g' ${GOPATH}/src/github.com/radekg/yugabyte-db-go-proto/yb/api/master.pb.go \
		&& rm -- ${GOPATH}/src/github.com/radekg/yugabyte-db-go-proto/yb/api/master.pb.go.bak
	sed -i.bak -e 's!DeleteTabletRequestPB!MasterDeleteTabletRequestPB!g' ${GOPATH}/src/github.com/radekg/yugabyte-db-go-proto/yb/api/master.pb.go \
		&& rm -- ${GOPATH}/src/github.com/radekg/yugabyte-db-go-proto/yb/api/master.pb.go.bak
	sed -i.bak -e 's!DeleteTabletResponsePB!MasterDeleteTabletResponsePB!g' ${GOPATH}/src/github.com/radekg/yugabyte-db-go-proto/yb/api/master.pb.go \
		&& rm -- ${GOPATH}/src/github.com/radekg/yugabyte-db-go-proto/yb/api/master.pb.go.bak
	sed -i.bak -e 's!SplitTabletRequestPB!MasterSplitTabletRequestPB!g' ${GOPATH}/src/github.com/radekg/yugabyte-db-go-proto/yb/api/master.pb.go \
		&& rm -- ${GOPATH}/src/github.com/radekg/yugabyte-db-go-proto/yb/api/master.pb.go.bak
	sed -i.bak -e 's!SplitTabletResponsePB!MasterSplitTabletResponsePB!g' ${GOPATH}/src/github.com/radekg/yugabyte-db-go-proto/yb/api/master.pb.go \
		&& rm -- ${GOPATH}/src/github.com/radekg/yugabyte-db-go-proto/yb/api/master.pb.go.bak
	sed -i.bak -e 's!CreateCDCStreamRequestPB!MasterCreateCDCStreamRequestPB!g' ${GOPATH}/src/github.com/radekg/yugabyte-db-go-proto/yb/api/master_grpc.pb.go \
		&& rm -- ${GOPATH}/src/github.com/radekg/yugabyte-db-go-proto/yb/api/master_grpc.pb.go.bak
	sed -i.bak -e 's!CreateCDCStreamResponsePB!MasterCreateCDCStreamResponsePB!g' ${GOPATH}/src/github.com/radekg/yugabyte-db-go-proto/yb/api/master_grpc.pb.go \
		&& rm -- ${GOPATH}/src/github.com/radekg/yugabyte-db-go-proto/yb/api/master_grpc.pb.go.bak
	sed -i.bak -e 's!DeleteCDCStreamRequestPB!MasterDeleteCDCStreamRequestPB!g' ${GOPATH}/src/github.com/radekg/yugabyte-db-go-proto/yb/api/master_grpc.pb.go \
		&& rm -- ${GOPATH}/src/github.com/radekg/yugabyte-db-go-proto/yb/api/master_grpc.pb.go.bak
	sed -i.bak -e 's!DeleteCDCStreamResponsePB!MasterDeleteCDCStreamResponsePB!g' ${GOPATH}/src/github.com/radekg/yugabyte-db-go-proto/yb/api/master_grpc.pb.go \
		&& rm -- ${GOPATH}/src/github.com/radekg/yugabyte-db-go-proto/yb/api/master_grpc.pb.go.bak
	sed -i.bak -e 's!BackfillIndexRequestPB!MasterBackfillIndexRequestPB!g' ${GOPATH}/src/github.com/radekg/yugabyte-db-go-proto/yb/api/master_grpc.pb.go \
		&& rm -- ${GOPATH}/src/github.com/radekg/yugabyte-db-go-proto/yb/api/master_grpc.pb.go.bak
	sed -i.bak -e 's!BackfillIndexResponsePB!MasterBackfillIndexResponsePB!g' ${GOPATH}/src/github.com/radekg/yugabyte-db-go-proto/yb/api/master_grpc.pb.go \
		&& rm -- ${GOPATH}/src/github.com/radekg/yugabyte-db-go-proto/yb/api/master_grpc.pb.go.bak
	sed -i.bak -e 's!DeleteTabletRequestPB!MasterDeleteTabletRequestPB!g' ${GOPATH}/src/github.com/radekg/yugabyte-db-go-proto/yb/api/master_grpc.pb.go \
		&& rm -- ${GOPATH}/src/github.com/radekg/yugabyte-db-go-proto/yb/api/master_grpc.pb.go.bak
	sed -i.bak -e 's!DeleteTabletResponsePB!MasterDeleteTabletResponsePB!g' ${GOPATH}/src/github.com/radekg/yugabyte-db-go-proto/yb/api/master_grpc.pb.go \
		&& rm -- ${GOPATH}/src/github.com/radekg/yugabyte-db-go-proto/yb/api/master_grpc.pb.go.bak
	sed -i.bak -e 's!SplitTabletRequestPB!MasterSplitTabletRequestPB!g' ${GOPATH}/src/github.com/radekg/yugabyte-db-go-proto/yb/api/master_grpc.pb.go \
		&& rm -- ${GOPATH}/src/github.com/radekg/yugabyte-db-go-proto/yb/api/master_grpc.pb.go.bak
	sed -i.bak -e 's!SplitTabletResponsePB!MasterSplitTabletResponsePB!g' ${GOPATH}/src/github.com/radekg/yugabyte-db-go-proto/yb/api/master_grpc.pb.go \
		&& rm -- ${GOPATH}/src/github.com/radekg/yugabyte-db-go-proto/yb/api/master_grpc.pb.go.bak

.PHONY: fix-cdc-service-types
fix-cdc-service-types:
	sed -i.bak -e 's!ListTabletsRequestPB!CDCListTabletsRequestPB!g' ${GOPATH}/src/github.com/radekg/yugabyte-db-go-proto/yb/api/cdc_service.pb.go \
		&& rm -- ${GOPATH}/src/github.com/radekg/yugabyte-db-go-proto/yb/api/cdc_service.pb.go.bak
	sed -i.bak -e 's!ListTabletsResponsePB!CDCListTabletsResponsePB!g' ${GOPATH}/src/github.com/radekg/yugabyte-db-go-proto/yb/api/cdc_service.pb.go \
		&& rm -- ${GOPATH}/src/github.com/radekg/yugabyte-db-go-proto/yb/api/cdc_service.pb.go.bak
	sed -i.bak -e 's!KeyValuePairPB!CDCKeyValuePairPB!g' ${GOPATH}/src/github.com/radekg/yugabyte-db-go-proto/yb/api/cdc_service.pb.go \
		&& rm -- ${GOPATH}/src/github.com/radekg/yugabyte-db-go-proto/yb/api/cdc_service.pb.go.bak
	sed -i.bak -e 's!ListTabletsRequestPB!CDCListTabletsRequestPB!g' ${GOPATH}/src/github.com/radekg/yugabyte-db-go-proto/yb/api/cdc_service_grpc.pb.go \
		&& rm -- ${GOPATH}/src/github.com/radekg/yugabyte-db-go-proto/yb/api/cdc_service_grpc.pb.go.bak
	sed -i.bak -e 's!ListTabletsResponsePB!CDCListTabletsResponsePB!g' ${GOPATH}/src/github.com/radekg/yugabyte-db-go-proto/yb/api/cdc_service_grpc.pb.go \
		&& rm -- ${GOPATH}/src/github.com/radekg/yugabyte-db-go-proto/yb/api/cdc_service_grpc.pb.go.bak
	sed -i.bak -e 's!KeyValuePairPB!CDCKeyValuePairPB!g' ${GOPATH}/src/github.com/radekg/yugabyte-db-go-proto/yb/api/cdc_service_grpc.pb.go \
		&& rm -- ${GOPATH}/src/github.com/radekg/yugabyte-db-go-proto/yb/api/cdc_service_grpc.pb.go.bak

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
