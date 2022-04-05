## gRPC-Go intro

Example https://earthly.dev/blog/golang-grpc-example/

### 1) Istall gRPC-Go 
Go https://github.com/grpc/grpc-go 
```
$ go get -u google.golang.org/grpc
```
### 2) Install protobuf
useful links
https://google.github.io/proto-lens/installing-protoc.html
https://developers.google.com/protocol-buffers/docs/gotutorial

for MAC:
```
brew install protobuf
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

### Install protoc-gen-go-grpc
 
```
 https://formulae.brew.sh/formula/protoc-gen-go-grpc
```

## Codegenerating test

Пример: https://golang-blog.blogspot.com/2021/03/quick-start-with-grpc-in-golang.html

Для перегенерации кода запустим:
```
$ protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    helloworld/helloworld.proto
```


### 1) --go_out
 - Код для заполнения, сериализации и получения типов сообщений 

### 2)  --go-grpc_out 
 - Сгенерированный клиентский и серверный код.



## gRPC Errors
https://avi.im/grpc-errors/
