FROM ubuntu
RUN apt update
RUN apt install golang -y
RUN go version
COPY . /gRPC
WORKDIR /gRPC/
RUN apt install ca-certificates -y
RUN go get google.golang.org/grpc
RUN mkdir "bin"
RUN go build -o /gRPC/bin/server.bin /gRPC/cmd/server/main.go
RUN go build -o /gRPC/bin/client.bin /gRPC/cmd/client/main.go
CMD ["bash", "-c", "/gRPC/bin/server.bin & /gRPC/bin/client.bin"]