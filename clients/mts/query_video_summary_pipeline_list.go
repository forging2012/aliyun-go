package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryVideoSummaryPipelineListRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	PipelineIds          string `position:"Query" name:"PipelineIds"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *QueryVideoSummaryPipelineListRequest) Invoke(client *sdk.Client) (resp *QueryVideoSummaryPipelineListResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "QueryVideoSummaryPipelineList", "mts", "")
	resp = &QueryVideoSummaryPipelineListResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryVideoSummaryPipelineListResponse struct {
	responses.BaseResponse
	RequestId    string
	PipelineList QueryVideoSummaryPipelineListPipelineList
	NonExistIds  QueryVideoSummaryPipelineListNonExistIdList
}

type QueryVideoSummaryPipelineListPipeline struct {
	Id           string
	Name         string
	State        string
	Priority     string
	NotifyConfig QueryVideoSummaryPipelineListNotifyConfig
}

type QueryVideoSummaryPipelineListNotifyConfig struct {
	Topic     string
	QueueName string
}

type QueryVideoSummaryPipelineListPipelineList []QueryVideoSummaryPipelineListPipeline

func (list *QueryVideoSummaryPipelineListPipelineList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryVideoSummaryPipelineListPipeline)
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

type QueryVideoSummaryPipelineListNonExistIdList []string

func (list *QueryVideoSummaryPipelineListNonExistIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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
