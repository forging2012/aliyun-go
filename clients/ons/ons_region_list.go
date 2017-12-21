package ons

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type OnsRegionListRequest struct {
	requests.RpcRequest
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
}

func (req *OnsRegionListRequest) Invoke(client *sdk.Client) (resp *OnsRegionListResponse, err error) {
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsRegionList", "", "")
	resp = &OnsRegionListResponse{}
	err = client.DoAction(req, resp)
	return
}

type OnsRegionListResponse struct {
	responses.BaseResponse
	RequestId string
	HelpUrl   string
	Data      OnsRegionListRegionDoList
}

type OnsRegionListRegionDo struct {
	Id          int64
	OnsRegionId string
	RegionName  string
	ChannelId   int
	ChannelName string
	CreateTime  int64
	UpdateTime  int64
}

type OnsRegionListRegionDoList []OnsRegionListRegionDo

func (list *OnsRegionListRegionDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsRegionListRegionDo)
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
