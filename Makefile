
all: 
	echo "Doing nothing"
	
protos:
	echo "Building protos"
	go install google.golang.org/grpc
	go install github.com/golang/protobuf/protoc-gen-go
	export PATH=$$PATH:$$GOPATH/bin
	protoc -I proto/ proto/restaurant-svc.proto --go_out=plugins=grpc:restaurant
	
clean:
	rm restaurant-svc.debug