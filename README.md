# gRPC-example
https://github.com/protocolbuffers/protobuf/releases/tag/v21.12 - protoc, создаем по своему GOPATH папки bin и src, в bin перекидываем protoc.exe из папки bin скачанного архива, в src - из include
добавляем в PATH свой GOPATH\bin\
после этого:

`go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28`</br>
`go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2`</br>
`go get google.golang.org/grpc`</br>

`protoc message.proto --go_out=..\\..\pkg\ --go-grpc_out=..\\..\pkg\` (команда для генерации файлов pb из proto файла, применять в папке с ним (в данном случае - в api\proto\))
