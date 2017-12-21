package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeSecurityGroupsRequest struct {
	requests.RpcRequest
	Tag4Value            string `position:"Query" name:"Tag.4.Value"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Tag2Key              string `position:"Query" name:"Tag.2.Key"`
	FuzzyQuery           string `position:"Query" name:"FuzzyQuery"`
	SecurityGroupId      string `position:"Query" name:"SecurityGroupId"`
	Tag3Key              string `position:"Query" name:"Tag.3.Key"`
	IsQueryEcsCount      string `position:"Query" name:"IsQueryEcsCount"`
	NetworkType          string `position:"Query" name:"NetworkType"`
	SecurityGroupName    string `position:"Query" name:"SecurityGroupName"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
	Tag1Value            string `position:"Query" name:"Tag.1.Value"`
	PageSize             int    `position:"Query" name:"PageSize"`
	Tag3Value            string `position:"Query" name:"Tag.3.Value"`
	Tag5Key              string `position:"Query" name:"Tag.5.Key"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	SecurityGroupIds     string `position:"Query" name:"SecurityGroupIds"`
	Tag5Value            string `position:"Query" name:"Tag.5.Value"`
	Tag1Key              string `position:"Query" name:"Tag.1.Key"`
	VpcId                string `position:"Query" name:"VpcId"`
	Tag2Value            string `position:"Query" name:"Tag.2.Value"`
	Tag4Key              string `position:"Query" name:"Tag.4.Key"`
}

func (req *DescribeSecurityGroupsRequest) Invoke(client *sdk.Client) (resp *DescribeSecurityGroupsResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeSecurityGroups", "ecs", "")
	resp = &DescribeSecurityGroupsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeSecurityGroupsResponse struct {
	responses.BaseResponse
	RequestId      string
	RegionId       string
	TotalCount     int
	PageNumber     int
	PageSize       int
	SecurityGroups DescribeSecurityGroupsSecurityGroupList
}

type DescribeSecurityGroupsSecurityGroup struct {
	SecurityGroupId         string
	Description             string
	SecurityGroupName       string
	VpcId                   string
	CreationTime            string
	AvailableInstanceAmount int
	EcsCount                int
	Tags                    DescribeSecurityGroupsTagList
}

type DescribeSecurityGroupsTag struct {
	TagKey   string
	TagValue string
}

type DescribeSecurityGroupsSecurityGroupList []DescribeSecurityGroupsSecurityGroup

func (list *DescribeSecurityGroupsSecurityGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSecurityGroupsSecurityGroup)
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

type DescribeSecurityGroupsTagList []DescribeSecurityGroupsTag

func (list *DescribeSecurityGroupsTagList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSecurityGroupsTag)
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
