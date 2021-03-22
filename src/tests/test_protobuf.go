package main

import (
	"fmt"

	_ "github.com/golang/protobuf/proto"
	//pb "github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
)

func main() {
	orig := &any.Any{Value: []byte("test")}
	packed, err := ptypes.MarshalAny(orig)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(packed)
}
