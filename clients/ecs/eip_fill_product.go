package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type EipFillProductRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Data                 string `position:"Query" name:"Data"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	UserCidr             string `position:"Query" name:"UserCidr"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *EipFillProductRequest) Invoke(client *sdk.Client) (resp *EipFillProductResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "EipFillProduct", "ecs", "")
	resp = &EipFillProductResponse{}
	err = client.DoAction(req, resp)
	return
}

type EipFillProductResponse struct {
	responses.BaseResponse
	RequestId string
	Data      string
	Code      string
	Success   bool
	Message   string
}
