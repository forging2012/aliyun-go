package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ClearUserDomainBlackListRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	OwnerAccount  string `position:"Query" name:"OwnerAccount"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r ClearUserDomainBlackListRequest) Invoke(client *sdk.Client) (response *ClearUserDomainBlackListResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ClearUserDomainBlackListRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "ClearUserDomainBlackList", "", "")

	resp := struct {
		*responses.BaseResponse
		ClearUserDomainBlackListResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ClearUserDomainBlackListResponse

	err = client.DoAction(&req, &resp)
	return
}

type ClearUserDomainBlackListResponse struct {
	RequestId string
}
