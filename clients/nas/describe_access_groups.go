package nas

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeAccessGroupsRequest struct {
	requests.RpcRequest
	PageSize        int    `position:"Query" name:"PageSize"`
	AccessGroupName string `position:"Query" name:"AccessGroupName"`
	PageNumber      int    `position:"Query" name:"PageNumber"`
}

func (req *DescribeAccessGroupsRequest) Invoke(client *sdk.Client) (resp *DescribeAccessGroupsResponse, err error) {
	req.InitWithApiInfo("NAS", "2017-06-26", "DescribeAccessGroups", "nas", "")
	resp = &DescribeAccessGroupsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeAccessGroupsResponse struct {
	responses.BaseResponse
	RequestId    string
	TotalCount   int
	PageSize     int
	PageNumber   int
	AccessGroups DescribeAccessGroupsAccessGroupList
}

type DescribeAccessGroupsAccessGroup struct {
	AccessGroupName  string
	AccessGroupType  string
	RuleCount        int
	MountTargetCount int
	Description      string
}

type DescribeAccessGroupsAccessGroupList []DescribeAccessGroupsAccessGroup

func (list *DescribeAccessGroupsAccessGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeAccessGroupsAccessGroup)
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
