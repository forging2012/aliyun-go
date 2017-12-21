package domain

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateOrderRequest struct {
	requests.RpcRequest
	SubOrderParams *CreateOrderSubOrderParamList `position:"Query" type:"Repeated" name:"SubOrderParam"`
}

func (req *CreateOrderRequest) Invoke(client *sdk.Client) (resp *CreateOrderResponse, err error) {
	req.InitWithApiInfo("Domain", "2016-05-11", "CreateOrder", "", "")
	resp = &CreateOrderResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateOrderSubOrderParam struct {
	SaleID           string `name:"SaleID"`
	RelatedName      string `name:"RelatedName"`
	Action           string `name:"Action"`
	Period           int    `name:"Period"`
	DomainTemplateID string `name:"DomainTemplateID"`
}

type CreateOrderResponse struct {
	responses.BaseResponse
	RequestId string
	OrderID   string
}

type CreateOrderSubOrderParamList []CreateOrderSubOrderParam

func (list *CreateOrderSubOrderParamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateOrderSubOrderParam)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
