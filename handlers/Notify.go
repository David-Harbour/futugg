package handlers

import (
	"fmt"
	"github.com/jerryharbour/futugg"
	"github.com/jerryharbour/futugg/pb/Notify"
	"log"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/jsonpb"
)

func init() {
	futugg.SetHandlerId(uint32(1003), "Notify")

	var err error
	err = futugg.On("recv.Notify", NotifyRecv)
	if err != nil {
		log.Fatalln(err)
	}
}

func NotifyRecv(data []byte) error {
	resp := &Notify.Response{}
	err := proto.Unmarshal(data, resp)
	if err != nil {
		return fmt.Errorf("marshal error: %s", err)
	}

	m := jsonpb.Marshaler{}
	result, err := m.MarshalToString(resp)
	fmt.Println(result)
    return err
}
