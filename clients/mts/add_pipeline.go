package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddPipelineRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Role                 string `position:"Query" name:"Role"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	NotifyConfig         string `position:"Query" name:"NotifyConfig"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	SpeedLevel           int64  `position:"Query" name:"SpeedLevel"`
	Speed                string `position:"Query" name:"Speed"`
}

func (r AddPipelineRequest) Invoke(client *sdk.Client) (response *AddPipelineResponse, err error) {
	req := struct {
		*requests.RpcRequest
		AddPipelineRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Mts", "2014-06-18", "AddPipeline", "", "")

	resp := struct {
		*responses.BaseResponse
		AddPipelineResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.AddPipelineResponse

	err = client.DoAction(&req, &resp)
	return
}

type AddPipelineResponse struct {
	RequestId string
	Pipeline  AddPipelinePipeline
}

type AddPipelinePipeline struct {
	Id           string
	Name         string
	State        string
	Speed        string
	SpeedLevel   int64
	Role         string
	NotifyConfig AddPipelineNotifyConfig
}

type AddPipelineNotifyConfig struct {
	Topic     string
	QueueName string
}
