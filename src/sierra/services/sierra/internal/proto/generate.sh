mkdir -p ../../../../packages/varsoapi/src/app
if ! command -v protoc-gen-go &> /dev/null; then
  echo "src/sierra/services/sierra/internal: protoc-gen-go not found, installing..."
  echo src/services/app/internal/proto: Installing protoc-gen-go
  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
fi

mkdir -p ../../../../../electron/src/proto/app
../../../../common/proto/bin/bin/protoc -I=./ --proto_path=../../../../common/proto/bin/include --plugin=$(go env GOPATH)/bin/protoc-gen-go --plugin=../../../../common/proto/node_modules/.bin/protoc-gen-ts_proto --ts_proto_opt=esModuleInterop=true --go_out=../../ --ts_proto_out=../../../../../electron/src/proto/app appbase.proto apprequests.proto