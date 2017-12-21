package ons

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type OnsPublishGetRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	Topic        string `position:"Query" name:"Topic"`
	ProducerId   string `position:"Query" name:"ProducerId"`
}

func (r OnsPublishGetRequest) Invoke(client *sdk.Client) (response *OnsPublishGetResponse, err error) {
	req := struct {
		*requests.RpcRequest
		OnsPublishGetRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsPublishGet", "", "")

	resp := struct {
		*responses.BaseResponse
		OnsPublishGetResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.OnsPublishGetResponse

	err = client.DoAction(&req, &resp)
	return
}

type OnsPublishGetResponse struct {
	RequestId string
	HelpUrl   string
	Data      OnsPublishGetPublishInfoDoList
}

type OnsPublishGetPublishInfoDo struct {
	Id          int64
	ChannelId   int
	ChannelName string
	OnsRegionId string
	RegionName  string
	Owner       string
	ProducerId  string
	Topic       string
	Status      int
	StatusName  string
	CreateTime  int64
	UpdateTime  int64
}

type OnsPublishGetPublishInfoDoList []OnsPublishGetPublishInfoDo

func (list *OnsPublishGetPublishInfoDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsPublishGetPublishInfoDo)
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
