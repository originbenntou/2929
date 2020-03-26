# protobuf生成 SERVICEを変数で渡す
gen_proto:
	protoc --proto_path=${GOPATH}/src --proto_path=. --go_out=plugins=grpc:${GOPATH}/src --govalidators_out=${GOPATH}/src ./proto/${SERVICE}/*.proto

# graphql生成
gen_graph:
	cd ./graphql
	go run github.com/99designs/gqlgen generate
