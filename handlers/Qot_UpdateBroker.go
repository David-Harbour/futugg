package handlers

import (
    "fmt"
    "github.com/jerryharbour/futugg"
    "github.com/jerryharbour/futugg/pb/Qot_UpdateBroker"

    "github.com/golang/protobuf/proto"
    "github.com/golang/protobuf/jsonpb"
)

func init() {
    futugg.SetHandlerId(uint32(3015), "Qot_UpdateBroker")

    var err error
    err = futugg.On("recv.Qot_UpdateBroker", QotUpdateBrokerRecv)
    if err != nil {
        fmt.Println(err)
    }
}

func QotUpdateBrokerRecv(data []byte) error {
    resp := &Qot_UpdateBroker.Response{}
    err := proto.Unmarshal(data, resp)
    if err != nil {
        return fmt.Errorf("marshal error: %s", err)
    }

    m := jsonpb.Marshaler{}
    result, err := m.MarshalToString(resp)
    fmt.Println(result)
    return err
}