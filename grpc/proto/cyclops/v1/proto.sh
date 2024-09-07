#
PATH=/bin:/usr/bin:/etc:/usr/local/bin; export PATH
#
#rm cyclops.pb.go
#rm cyclops_grpc.pb.go
#
# gsc@lingling:19>/usr/local/bin/protoc --version
# libprotoc 25.1
#
# brew install protoc-gen-go-grpc
#
#protoc --go_out=../../../gen/cyclops/v1 --go_opt=paths=source_relative --go-grpc_out=../../../gen/cyclops/v1 --go-grpc_opt=paths=source_relative ./cyclops.proto
#
echo "obsolete, use buf instead"
#
