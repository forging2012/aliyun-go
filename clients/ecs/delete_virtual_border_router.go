package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteVirtualBorderRouterRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	UserCidr             string `position:"Query" name:"UserCidr"`
	VbrId                string `position:"Query" name:"VbrId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DeleteVirtualBorderRouterRequest) Invoke(client *sdk.Client) (response *DeleteVirtualBorderRouterResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteVirtualBorderRouterRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "DeleteVirtualBorderRouter", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		DeleteVirtualBorderRouterResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteVirtualBorderRouterResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteVirtualBorderRouterResponse struct {
	RequestId string
}
