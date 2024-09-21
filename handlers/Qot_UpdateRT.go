package handlers

import (
    "fmt"
    "github.com/jerryharbour/futugg"
    "github.com/jerryharbour/futugg/pb/Qot_UpdateRT"

    "github.com/golang/protobuf/proto"
    "github.com/golang/protobuf/jsonpb"
)

func init() {
    futugg.SetHandlerId(uint32(3009), "Qot_UpdateRT")

    var err error
    err = futugg.On("recv.Qot_UpdateRT", QotUpdateRTRecv)
    if err != nil {
        fmt.Println(err)
    }
}

func QotUpdateRTRecv(data []byte) error {
    resp := &Qot_UpdateRT.Response{}
    err := proto.Unmarshal(data, resp)
    if err != nil {
        return fmt.Errorf("marshal error: %s", err)
    }

    m := jsonpb.Marshaler{}
    result, err := m.MarshalToString(resp)
    fmt.Println(result)
    return err
}