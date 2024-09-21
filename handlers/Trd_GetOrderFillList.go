package handlers

import (
    "fmt"
    "github.com/jerryharbour/futugg"
    "github.com/jerryharbour/futugg/pb/Trd_GetOrderFillList"

    "github.com/golang/protobuf/proto"
    "github.com/golang/protobuf/jsonpb"
)

func init() {
    futugg.SetHandlerId(uint32(2211), "Trd_GetOrderFillList")
    var err error
    err = futugg.On("send.Trd_GetOrderFillList", TrdGetOrderFillListSend)
    err = futugg.On("recv.Trd_GetOrderFillList", TrdGetOrderFillListRecv)
    if err != nil {
        fmt.Println(err)
    }
}

// TODO add TrdFilterConditions
func TrdGetOrderFillListSend(conn *futugg.FutuGG, trdEnv int32, accID uint64, trdMarket int32) error {
    pack := &futugg.FutuPack{}
    pack.SetProto(uint32(2211))

    header := setTrdHeader(trdEnv, accID, trdMarket)
    reqData := &Trd_GetOrderFillList.Request{
        C2S: &Trd_GetOrderFillList.C2S{
            Header: header,
        },
    }
    pbData, err := proto.Marshal(reqData)
    if err != nil {
        return fmt.Errorf("marshal error: %s", err)
    }

    pack.SetBody(pbData)
    err = conn.Send(pack)

    return err
}

func TrdGetOrderFillListRecv(data []byte) error {
    resp := &Trd_GetOrderFillList.Response{}
    err := proto.Unmarshal(data, resp)
    if err != nil {
        return fmt.Errorf("marshal error: %s", err)
    }

    m := jsonpb.Marshaler{}
    result, err := m.MarshalToString(resp)
    fmt.Println(result)
    return err
}
