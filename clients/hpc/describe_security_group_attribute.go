package hpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeSecurityGroupAttributeRequest struct {
	requests.RpcRequest
	TOKEN   string `position:"Query" name:"TOKEN"`
	NicType string `position:"Query" name:"NicType"`
}

func (req *DescribeSecurityGroupAttributeRequest) Invoke(client *sdk.Client) (resp *DescribeSecurityGroupAttributeResponse, err error) {
	req.InitWithApiInfo("HPC", "2016-12-13", "DescribeSecurityGroupAttribute", "hpc", "")
	resp = &DescribeSecurityGroupAttributeResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeSecurityGroupAttributeResponse struct {
	responses.BaseResponse
	Records DescribeSecurityGroupAttributeRecords
}

type DescribeSecurityGroupAttributeRecords struct {
	RegionId    string
	Permissions DescribeSecurityGroupAttributePermissionList
}

type DescribeSecurityGroupAttributePermission struct {
	SourceIp string
	Policy   DescribeSecurityGroupAttributePolicy
	NicType  DescribeSecurityGroupAttributeNicType
	Priority string
	Time     string
}

type DescribeSecurityGroupAttributePolicy struct {
	StringValue string
}

type DescribeSecurityGroupAttributeNicType struct {
	StringValue string
}

type DescribeSecurityGroupAttributePermissionList []DescribeSecurityGroupAttributePermission

func (list *DescribeSecurityGroupAttributePermissionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSecurityGroupAttributePermission)
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
