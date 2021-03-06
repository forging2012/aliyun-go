package nas

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeAccessRulesRequest struct {
	requests.RpcRequest
	PageSize        int    `position:"Query" name:"PageSize"`
	AccessGroupName string `position:"Query" name:"AccessGroupName"`
	AccessRuleId    string `position:"Query" name:"AccessRuleId"`
	PageNumber      int    `position:"Query" name:"PageNumber"`
}

func (req *DescribeAccessRulesRequest) Invoke(client *sdk.Client) (resp *DescribeAccessRulesResponse, err error) {
	req.InitWithApiInfo("NAS", "2017-06-26", "DescribeAccessRules", "nas", "")
	resp = &DescribeAccessRulesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeAccessRulesResponse struct {
	responses.BaseResponse
	RequestId   string
	TotalCount  int
	PageSize    int
	PageNumber  int
	AccessRules DescribeAccessRulesAccessRuleList
}

type DescribeAccessRulesAccessRule struct {
	SourceCidrIp string
	Priority     int
	AccessRuleId string
	RWAccess     string
	UserAccess   string
}

type DescribeAccessRulesAccessRuleList []DescribeAccessRulesAccessRule

func (list *DescribeAccessRulesAccessRuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeAccessRulesAccessRule)
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
