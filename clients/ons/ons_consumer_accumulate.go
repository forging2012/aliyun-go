package ons

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type OnsConsumerAccumulateRequest struct {
	requests.RpcRequest
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	ConsumerId   string `position:"Query" name:"ConsumerId"`
	Detail       string `position:"Query" name:"Detail"`
}

func (req *OnsConsumerAccumulateRequest) Invoke(client *sdk.Client) (resp *OnsConsumerAccumulateResponse, err error) {
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsConsumerAccumulate", "", "")
	resp = &OnsConsumerAccumulateResponse{}
	err = client.DoAction(req, resp)
	return
}

type OnsConsumerAccumulateResponse struct {
	responses.BaseResponse
	RequestId string
	HelpUrl   string
	Data      OnsConsumerAccumulateData
}

type OnsConsumerAccumulateData struct {
	Online            bool
	TotalDiff         int64
	ConsumeTps        float32
	LastTimestamp     int64
	DelayTime         int64
	DetailInTopicList OnsConsumerAccumulateDetailInTopicDoList
}

type OnsConsumerAccumulateDetailInTopicDo struct {
	Topic         string
	TotalDiff     int64
	LastTimestamp int64
	DelayTime     int64
}

type OnsConsumerAccumulateDetailInTopicDoList []OnsConsumerAccumulateDetailInTopicDo

func (list *OnsConsumerAccumulateDetailInTopicDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsConsumerAccumulateDetailInTopicDo)
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
