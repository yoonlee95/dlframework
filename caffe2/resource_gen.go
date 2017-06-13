//go:generate protoc --plugin=protoc-gen-go=${GOPATH}/bin/protoc-gen-go --proto_path=./proto:.. --gogofaster_out=plugins=grpc:.  proto/caffe2.proto proto/caffe2_legacy.proto proto/hsm.proto proto/metanet.proto proto/predictor_consts.proto proto/prof_dag.proto
//go:generate go-bindata -nomemcopy -prefix builtin_models/ -pkg caffe2 -o builtin_models_static.go -ignore=.DS_Store ./builtin_models/...
package caffe2
