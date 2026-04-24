#!/usr/bin/env fish
echo "Generating protobuf files..."

set PROTO_DIR "manifest/protobuf"
set API_DIR "api"
set PROTOC "/opt/homebrew/bin/protoc"

# 先找到所有 .proto 文件
set proto_files (find $PROTO_DIR -name "*.proto" | sort)

echo "Found proto files:"
for f in $proto_files
    echo "  - $f"
end

echo ""

# 运行 protoc
$PROTOC \
  --proto_path=$PROTO_DIR \
  --proto_path=/opt/homebrew/include \
  --go_out=paths=source_relative:$API_DIR \
  --go-grpc_out=paths=source_relative:$API_DIR \
  $proto_files

if test $status -eq 0
    echo ""
    echo "✅ Done! Generated files in $API_DIR/"
else
    echo ""
    echo "❌ Failed to generate protobuf files"
    exit 1
end
