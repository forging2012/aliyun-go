package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyRouterInterfaceSpecRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	UserCidr             string `position:"Query" name:"UserCidr"`
	RouterInterfaceId    string `position:"Query" name:"RouterInterfaceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Spec                 string `position:"Query" name:"Spec"`
}

func (req *ModifyRouterInterfaceSpecRequest) Invoke(client *sdk.Client) (resp *ModifyRouterInterfaceSpecResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "ModifyRouterInterfaceSpec", "ecs", "")
	resp = &ModifyRouterInterfaceSpecResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyRouterInterfaceSpecResponse struct {
	responses.BaseResponse
	RequestId string
	Spec      string
}
