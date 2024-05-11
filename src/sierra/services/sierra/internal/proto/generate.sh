mkdir -p ../../../../packages/varsoapi/src/app
echo src/services/app/internal/proto: Installing protoc-gen-go
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
mkdir -p ../../../../../electron/src/proto/app
protoc -I=./ --plugin=$(go env GOPATH)/bin/protoc-gen-go --plugin=../../../../common/proto/node_modules/.bin/protoc-gen-ts_proto --ts_proto_opt=esModuleInterop=true --go_out=../../ --ts_proto_out=../../../../../electron/src/proto/app appbase.proto apprequests.proto