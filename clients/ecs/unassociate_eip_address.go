package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UnassociateEipAddressRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	InstanceType         string `position:"Query" name:"InstanceType"`
	AllocationId         string `position:"Query" name:"AllocationId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *UnassociateEipAddressRequest) Invoke(client *sdk.Client) (resp *UnassociateEipAddressResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "UnassociateEipAddress", "ecs", "")
	resp = &UnassociateEipAddressResponse{}
	err = client.DoAction(req, resp)
	return
}

type UnassociateEipAddressResponse struct {
	responses.BaseResponse
	RequestId string
}
