package csb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateOrderRequest struct {
	requests.RpcRequest
	Data  string `position:"Body" name:"Data"`
	CsbId int64  `position:"Query" name:"CsbId"`
}

func (req *CreateOrderRequest) Invoke(client *sdk.Client) (resp *CreateOrderResponse, err error) {
	req.InitWithApiInfo("CSB", "2017-11-18", "CreateOrder", "CSB", "")
	resp = &CreateOrderResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateOrderResponse struct {
	responses.BaseResponse
	Code      int
	Message   string
	RequestId string
	Data      CreateOrderData
}

type CreateOrderData struct {
	Id int64
}
