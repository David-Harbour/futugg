package handlers

import (
    "fmt"
    "github.com/jerryharbour/futugg"
    "github.com/jerryharbour/futugg/pb/Qot_GetBroker"

    "github.com/golang/protobuf/proto"
    "github.com/golang/protobuf/jsonpb"
)

func init() {
    futugg.SetHandlerId(uint32(3014), "Qot_GetBroker")
    var err error
    err = futugg.On("send.Qot_GetBroker", QotGetBrokerSend)
    err = futugg.On("recv.Qot_GetBroker", QotGetBrokerRecv)
    if err != nil {
        fmt.Println(err)
    }
}

func QotGetBrokerSend(conn *futugg.FutuGG, stockCode string) error {
    pack := &futugg.FutuPack{}
    pack.SetProto(uint32(3014))

    security := transStockCode(stockCode)

    reqData := &Qot_GetBroker.Request{
        C2S: &Qot_GetBroker.C2S{
            Security: security,
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

func QotGetBrokerRecv(data []byte) error {
    resp := &Qot_GetBroker.Response{}
    err := proto.Unmarshal(data, resp)
    if err != nil {
        return fmt.Errorf("marshal error: %s", err)
    }

    m := jsonpb.Marshaler{}
    result, err := m.MarshalToString(resp)
    fmt.Println(result)
    return err
}
