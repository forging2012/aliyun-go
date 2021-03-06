package domain

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteEmailVerificationRequest struct {
	requests.RpcRequest
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
	Email        string `position:"Query" name:"Email"`
}

func (req *DeleteEmailVerificationRequest) Invoke(client *sdk.Client) (resp *DeleteEmailVerificationResponse, err error) {
	req.InitWithApiInfo("Domain", "2018-01-29", "DeleteEmailVerification", "", "")
	resp = &DeleteEmailVerificationResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteEmailVerificationResponse struct {
	responses.BaseResponse
	RequestId   string
	SuccessList DeleteEmailVerificationSendResultList
	FailList    DeleteEmailVerificationSendResultList
}

type DeleteEmailVerificationSendResult struct {
	Email   string
	Code    string
	Message string
}

type DeleteEmailVerificationSendResultList []DeleteEmailVerificationSendResult

func (list *DeleteEmailVerificationSendResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DeleteEmailVerificationSendResult)
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
