package kms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type EncryptRequest struct {
	requests.RpcRequest
	KeyId             string `position:"Query" name:"KeyId"`
	Plaintext         string `position:"Query" name:"Plaintext"`
	STSToken          string `position:"Query" name:"STSToken"`
	EncryptionContext string `position:"Query" name:"EncryptionContext"`
}

func (req *EncryptRequest) Invoke(client *sdk.Client) (resp *EncryptResponse, err error) {
	req.InitWithApiInfo("Kms", "2016-01-20", "Encrypt", "kms", "")
	resp = &EncryptResponse{}
	err = client.DoAction(req, resp)
	return
}

type EncryptResponse struct {
	responses.BaseResponse
	CiphertextBlob string
	KeyId          string
	RequestId      string
}
