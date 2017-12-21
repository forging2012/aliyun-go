package ons

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type OnsMessageGetByMsgIdRequest struct {
	requests.RpcRequest
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	MsgId        string `position:"Query" name:"MsgId"`
	Topic        string `position:"Query" name:"Topic"`
}

func (req *OnsMessageGetByMsgIdRequest) Invoke(client *sdk.Client) (resp *OnsMessageGetByMsgIdResponse, err error) {
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsMessageGetByMsgId", "", "")
	resp = &OnsMessageGetByMsgIdResponse{}
	err = client.DoAction(req, resp)
	return
}

type OnsMessageGetByMsgIdResponse struct {
	responses.BaseResponse
	RequestId string
	HelpUrl   string
	Data      OnsMessageGetByMsgIdData
}

type OnsMessageGetByMsgIdData struct {
	Topic          string
	Flag           int
	Body           string
	StoreSize      int
	BornTimestamp  int64
	BornHost       string
	StoreTimestamp int64
	StoreHost      string
	MsgId          string
	OffsetId       string
	BodyCRC        int
	ReconsumeTimes int
	PropertyList   OnsMessageGetByMsgIdMessagePropertyList
}

type OnsMessageGetByMsgIdMessageProperty struct {
	Name  string
	Value string
}

type OnsMessageGetByMsgIdMessagePropertyList []OnsMessageGetByMsgIdMessageProperty

func (list *OnsMessageGetByMsgIdMessagePropertyList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsMessageGetByMsgIdMessageProperty)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
