package cloudapi

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDomainsResolutionRequest struct {
	requests.RpcRequest
	GroupId     string `position:"Query" name:"GroupId"`
	DomainNames string `position:"Query" name:"DomainNames"`
}

func (req *DescribeDomainsResolutionRequest) Invoke(client *sdk.Client) (resp *DescribeDomainsResolutionResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeDomainsResolution", "apigateway", "")
	resp = &DescribeDomainsResolutionResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDomainsResolutionResponse struct {
	responses.BaseResponse
	RequestId         string
	GroupId           string
	DomainResolutions DescribeDomainsResolutionDomainResolutionList
}

type DescribeDomainsResolutionDomainResolution struct {
	DomainName             string
	DomainResolutionStatus string
}

type DescribeDomainsResolutionDomainResolutionList []DescribeDomainsResolutionDomainResolution

func (list *DescribeDomainsResolutionDomainResolutionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainsResolutionDomainResolution)
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
