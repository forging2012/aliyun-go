package slb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetLoadBalancerAutoReleaseTimeRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	LoadBalancerId       string `position:"Query" name:"LoadBalancerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	AutoReleaseTime      int64  `position:"Query" name:"AutoReleaseTime"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tags                 string `position:"Query" name:"Tags"`
}

func (r SetLoadBalancerAutoReleaseTimeRequest) Invoke(client *sdk.Client) (response *SetLoadBalancerAutoReleaseTimeResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SetLoadBalancerAutoReleaseTimeRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Slb", "2014-05-15", "SetLoadBalancerAutoReleaseTime", "slb", "")

	resp := struct {
		*responses.BaseResponse
		SetLoadBalancerAutoReleaseTimeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SetLoadBalancerAutoReleaseTimeResponse

	err = client.DoAction(&req, &resp)
	return
}

type SetLoadBalancerAutoReleaseTimeResponse struct {
	RequestId string
}
