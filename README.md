# 协议说明
方式一、生成方式协议文件和grpc服务
1、protoc --go_out=. --go-grpc_out=. --proto_path=. xxx.proto ，google目录与xxx.proto处于同目录，否则会报错找不到goole目录。
生成http转grpc文件
2、protoc --grpc-gateway_out=./ --proto_path=. helloworld.proto

方式二、Kratos框架
1
protoc --proto_path=. --proto_path=../thirdparty --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. --go-http_out=paths=source_relative:. --validate_out=lang=go:. *.proto
2、自定义protobuf tags
protoc --proto_path=. --proto_path=../thirdparty --gogo_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. --go-http_out=paths=source_relative:. --validate_out=lang=go:. *.proto


# 协议说明
woa 微信公众平台回调get请求默认Accept:*/*，默认回包编码采用application/json，所以woa_http.pb.go手动改成ctx.String.