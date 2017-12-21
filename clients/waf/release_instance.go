package waf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ReleaseInstanceRequest struct {
	OwnerId    int64  `position:"Query" name:"OwnerId"`
	InstanceId string `position:"Query" name:"InstanceId"`
}

func (r ReleaseInstanceRequest) Invoke(client *sdk.Client) (response *ReleaseInstanceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ReleaseInstanceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("waf-openapi", "2016-11-11", "ReleaseInstance", "", "")

	resp := struct {
		*responses.BaseResponse
		ReleaseInstanceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ReleaseInstanceResponse

	err = client.DoAction(&req, &resp)
	return
}

type ReleaseInstanceResponse struct {
	RequestId string
}