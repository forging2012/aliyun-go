package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteCasterComponentRequest struct {
	requests.RpcRequest
	ComponentId   string `position:"Query" name:"ComponentId"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	CasterId      string `position:"Query" name:"CasterId"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	Version       string `position:"Query" name:"Version"`
}

func (req *DeleteCasterComponentRequest) Invoke(client *sdk.Client) (resp *DeleteCasterComponentResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "DeleteCasterComponent", "", "")
	resp = &DeleteCasterComponentResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteCasterComponentResponse struct {
	responses.BaseResponse
	RequestId   string
	CasterId    string
	ComponentId string
}
