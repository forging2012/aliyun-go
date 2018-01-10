package csb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteServiceListRequest struct {
	requests.RpcRequest
	Data  string `position:"Body" name:"Data"`
	CsbId int64  `position:"Query" name:"CsbId"`
}

func (req *DeleteServiceListRequest) Invoke(client *sdk.Client) (resp *DeleteServiceListResponse, err error) {
	req.InitWithApiInfo("CSB", "2017-11-18", "DeleteServiceList", "CSB", "")
	resp = &DeleteServiceListResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteServiceListResponse struct {
	responses.BaseResponse
	Code      int
	Message   string
	RequestId string
}
