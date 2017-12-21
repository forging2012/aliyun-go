package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeAutoSnapshotPolicyExRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	AutoSnapshotPolicyId string `position:"Query" name:"AutoSnapshotPolicyId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int    `position:"Query" name:"PageSize"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
}

func (r DescribeAutoSnapshotPolicyExRequest) Invoke(client *sdk.Client) (response *DescribeAutoSnapshotPolicyExResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeAutoSnapshotPolicyExRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeAutoSnapshotPolicyEx", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		DescribeAutoSnapshotPolicyExResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeAutoSnapshotPolicyExResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeAutoSnapshotPolicyExResponse struct {
	RequestId            string
	TotalCount           int
	PageNumber           int
	PageSize             int
	AutoSnapshotPolicies DescribeAutoSnapshotPolicyExAutoSnapshotPolicyList
}

type DescribeAutoSnapshotPolicyExAutoSnapshotPolicy struct {
	AutoSnapshotPolicyId   string
	RegionId               string
	AutoSnapshotPolicyName string
	TimePoints             string
	RepeatWeekdays         string
	RetentionDays          int
	DiskNums               int
	VolumeNums             int
	CreationTime           string
	Status                 string
}

type DescribeAutoSnapshotPolicyExAutoSnapshotPolicyList []DescribeAutoSnapshotPolicyExAutoSnapshotPolicy

func (list *DescribeAutoSnapshotPolicyExAutoSnapshotPolicyList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeAutoSnapshotPolicyExAutoSnapshotPolicy)
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
