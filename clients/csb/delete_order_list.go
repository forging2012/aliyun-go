package csb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteOrderListRequest struct {
	requests.RpcRequest
	Data string `position:"Body" name:"Data"`
}

func (req *DeleteOrderListRequest) Invoke(client *sdk.Client) (resp *DeleteOrderListResponse, err error) {
	req.InitWithApiInfo("CSB", "2017-11-18", "DeleteOrderList", "CSB", "")
	resp = &DeleteOrderListResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteOrderListResponse struct {
	responses.BaseResponse
	Code      int
	Message   string
	RequestId string
}
