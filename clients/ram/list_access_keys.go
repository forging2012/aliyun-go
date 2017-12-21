package ram

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListAccessKeysRequest struct {
	UserName string `position:"Query" name:"UserName"`
}

func (r ListAccessKeysRequest) Invoke(client *sdk.Client) (response *ListAccessKeysResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ListAccessKeysRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ram", "2015-05-01", "ListAccessKeys", "", "")

	resp := struct {
		*responses.BaseResponse
		ListAccessKeysResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ListAccessKeysResponse

	err = client.DoAction(&req, &resp)
	return
}

type ListAccessKeysResponse struct {
	RequestId  string
	AccessKeys ListAccessKeysAccessKeyList
}

type ListAccessKeysAccessKey struct {
	AccessKeyId string
	Status      string
	CreateDate  string
}

type ListAccessKeysAccessKeyList []ListAccessKeysAccessKey

func (list *ListAccessKeysAccessKeyList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAccessKeysAccessKey)
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
