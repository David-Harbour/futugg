package handlers

import (
    "fmt"
    "github.com/jerryharbour/futugg"
    "github.com/jerryharbour/futugg/pb/Trd_UpdateOrderFill"

    "github.com/golang/protobuf/proto"
    "github.com/golang/protobuf/jsonpb"
)

func init() {
    futugg.SetHandlerId(uint32(2218), "Trd_UpdateOrderFill")
    var err error
    err = futugg.On("recv.Trd_UpdateOrderFill", TrdUpdateOrderFillRecv)
    if err != nil {
        fmt.Println(err)
    }
}


func TrdUpdateOrderFillRecv(data []byte) error {
    resp := &Trd_UpdateOrderFill.Response{}
    err := proto.Unmarshal(data, resp)
    if err != nil {
        return fmt.Errorf("marshal error: %s", err)
    }

    m := jsonpb.Marshaler{}
    result, err := m.MarshalToString(resp)
    fmt.Println(result)
    return err
}
