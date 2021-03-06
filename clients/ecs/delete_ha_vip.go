package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteHaVipRequest struct {
	requests.RpcRequest
	HaVipId              string `position:"Query" name:"HaVipId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DeleteHaVipRequest) Invoke(client *sdk.Client) (resp *DeleteHaVipResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DeleteHaVip", "ecs", "")
	resp = &DeleteHaVipResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteHaVipResponse struct {
	responses.BaseResponse
	RequestId string
}
