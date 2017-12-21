package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetReqAuthConfigRequest struct {
	Key1          string `position:"Query" name:"Key.1"`
	Key2          string `position:"Query" name:"Key.2"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	TimeOut       string `position:"Query" name:"TimeOut"`
	AuthType      string `position:"Query" name:"AuthType"`
}

func (r SetReqAuthConfigRequest) Invoke(client *sdk.Client) (response *SetReqAuthConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SetReqAuthConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "SetReqAuthConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		SetReqAuthConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SetReqAuthConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type SetReqAuthConfigResponse struct {
	RequestId string
}
