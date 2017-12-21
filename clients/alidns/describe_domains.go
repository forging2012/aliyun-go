package alidns

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDomainsRequest struct {
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	KeyWord      string `position:"Query" name:"KeyWord"`
	GroupId      string `position:"Query" name:"GroupId"`
	PageNumber   int64  `position:"Query" name:"PageNumber"`
	PageSize     int64  `position:"Query" name:"PageSize"`
}

func (r DescribeDomainsRequest) Invoke(client *sdk.Client) (response *DescribeDomainsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeDomainsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Alidns", "2015-01-09", "DescribeDomains", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeDomainsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeDomainsResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeDomainsResponse struct {
	RequestId  string
	TotalCount int64
	PageNumber int64
	PageSize   int64
	Domains    DescribeDomainsDomainList
}

type DescribeDomainsDomain struct {
	DomainId        string
	DomainName      string
	PunyCode        string
	AliDomain       bool
	RegistrantEmail string
	GroupId         string
	GroupName       string
	InstanceId      string
	VersionCode     string
	VersionName     string
	DnsServers      DescribeDomainsDnsServerList
}

type DescribeDomainsDomainList []DescribeDomainsDomain

func (list *DescribeDomainsDomainList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainsDomain)
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

type DescribeDomainsDnsServerList []string

func (list *DescribeDomainsDnsServerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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
