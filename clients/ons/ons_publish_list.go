package ons

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type OnsPublishListRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
}

func (r OnsPublishListRequest) Invoke(client *sdk.Client) (response *OnsPublishListResponse, err error) {
	req := struct {
		*requests.RpcRequest
		OnsPublishListRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsPublishList", "", "")

	resp := struct {
		*responses.BaseResponse
		OnsPublishListResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.OnsPublishListResponse

	err = client.DoAction(&req, &resp)
	return
}

type OnsPublishListResponse struct {
	RequestId string
	HelpUrl   string
	Data      OnsPublishListPublishInfoDoList
}

type OnsPublishListPublishInfoDo struct {
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

type OnsPublishListPublishInfoDoList []OnsPublishListPublishInfoDo

func (list *OnsPublishListPublishInfoDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsPublishListPublishInfoDo)
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
