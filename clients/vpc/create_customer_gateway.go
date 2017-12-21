package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateCustomerGatewayRequest struct {
	IpAddress            string `position:"Query" name:"IpAddress"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	Description          string `position:"Query" name:"Description"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r CreateCustomerGatewayRequest) Invoke(client *sdk.Client) (response *CreateCustomerGatewayResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateCustomerGatewayRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Vpc", "2016-04-28", "CreateCustomerGateway", "vpc", "")

	resp := struct {
		*responses.BaseResponse
		CreateCustomerGatewayResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateCustomerGatewayResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateCustomerGatewayResponse struct {
	RequestId         string
	CustomerGatewayId string
	IpAddress         string
	Name              string
	Description       string
	CreateTime        int64
}
