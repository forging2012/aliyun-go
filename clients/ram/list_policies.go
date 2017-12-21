package ram

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListPoliciesRequest struct {
	PolicyType string `position:"Query" name:"PolicyType"`
	Marker     string `position:"Query" name:"Marker"`
	MaxItems   int    `position:"Query" name:"MaxItems"`
}

func (r ListPoliciesRequest) Invoke(client *sdk.Client) (response *ListPoliciesResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ListPoliciesRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ram", "2015-05-01", "ListPolicies", "", "")

	resp := struct {
		*responses.BaseResponse
		ListPoliciesResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ListPoliciesResponse

	err = client.DoAction(&req, &resp)
	return
}

type ListPoliciesResponse struct {
	RequestId   string
	IsTruncated bool
	Marker      string
	Policies    ListPoliciesPolicyList
}

type ListPoliciesPolicy struct {
	PolicyName      string
	PolicyType      string
	Description     string
	DefaultVersion  string
	CreateDate      string
	UpdateDate      string
	AttachmentCount int
}

type ListPoliciesPolicyList []ListPoliciesPolicy

func (list *ListPoliciesPolicyList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListPoliciesPolicy)
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
