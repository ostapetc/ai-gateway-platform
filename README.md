### Generate service code by service.proto file
goctl rpc protoc service.proto \         artos@artos-pc
--go_out=./grpc \    
--go-grpc_out=./grpc \
--zrpc_out=.


### Generate service code by service.api file
goctl api go -api service.api -dir . 

### Update all submodules at once
git submodule update --remote --recursive