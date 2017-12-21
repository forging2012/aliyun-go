package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddPornPipelineRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	NotifyConfig         string `position:"Query" name:"NotifyConfig"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Priority             int    `position:"Query" name:"Priority"`
}

func (r AddPornPipelineRequest) Invoke(client *sdk.Client) (response *AddPornPipelineResponse, err error) {
	req := struct {
		*requests.RpcRequest
		AddPornPipelineRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Mts", "2014-06-18", "AddPornPipeline", "", "")

	resp := struct {
		*responses.BaseResponse
		AddPornPipelineResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.AddPornPipelineResponse

	err = client.DoAction(&req, &resp)
	return
}

type AddPornPipelineResponse struct {
	RequestId string
	Pipeline  AddPornPipelinePipeline
}

type AddPornPipelinePipeline struct {
	Id           string
	Name         string
	Priority     int
	State        string
	NotifyConfig AddPornPipelineNotifyConfig
}

type AddPornPipelineNotifyConfig struct {
	Topic string
	Queue string
}
