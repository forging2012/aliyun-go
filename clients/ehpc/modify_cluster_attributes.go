package ehpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyClusterAttributesRequest struct {
	requests.RpcRequest
	Name        string `position:"Query" name:"Name"`
	Description string `position:"Query" name:"Description"`
	ClusterId   string `position:"Query" name:"ClusterId"`
}

func (req *ModifyClusterAttributesRequest) Invoke(client *sdk.Client) (resp *ModifyClusterAttributesResponse, err error) {
	req.InitWithApiInfo("EHPC", "2017-07-14", "ModifyClusterAttributes", "ehs", "")
	resp = &ModifyClusterAttributesResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyClusterAttributesResponse struct {
	responses.BaseResponse
	RequestId string
}
