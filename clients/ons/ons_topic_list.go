package ons

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type OnsTopicListRequest struct {
	requests.RpcRequest
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	Topic        string `position:"Query" name:"Topic"`
}

func (req *OnsTopicListRequest) Invoke(client *sdk.Client) (resp *OnsTopicListResponse, err error) {
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsTopicList", "", "")
	resp = &OnsTopicListResponse{}
	err = client.DoAction(req, resp)
	return
}

type OnsTopicListResponse struct {
	responses.BaseResponse
	RequestId string
	HelpUrl   string
	Data      OnsTopicListPublishInfoDoList
}

type OnsTopicListPublishInfoDo struct {
	Id           int64
	ChannelId    int
	ChannelName  string
	OnsRegionId  string
	RegionName   string
	Topic        string
	Owner        string
	Relation     int
	RelationName string
	Status       int
	StatusName   string
	Appkey       string
	CreateTime   int64
	UpdateTime   int64
	Remark       string
}

type OnsTopicListPublishInfoDoList []OnsTopicListPublishInfoDo

func (list *OnsTopicListPublishInfoDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsTopicListPublishInfoDo)
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
