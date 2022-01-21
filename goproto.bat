
set input_path=../proto
set input_out=../proto_out

cd ./tools
protoc --go_out=plugins=grpc:../proto_out *.proto -I %input_path%