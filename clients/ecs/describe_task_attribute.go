package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeTaskAttributeRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	TaskId               string `position:"Query" name:"TaskId"`
}

func (r DescribeTaskAttributeRequest) Invoke(client *sdk.Client) (response *DescribeTaskAttributeResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeTaskAttributeRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeTaskAttribute", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		DescribeTaskAttributeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeTaskAttributeResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeTaskAttributeResponse struct {
	RequestId            string
	TaskId               string
	RegionId             string
	TaskAction           string
	TaskStatus           string
	TaskProcess          string
	SupportCancel        string
	TotalCount           int
	SuccessCount         int
	FailedCount          int
	CreationTime         string
	FinishedTime         string
	OperationProgressSet DescribeTaskAttributeOperationProgressList
}

type DescribeTaskAttributeOperationProgress struct {
	OperationStatus string
	ErrorCode       string
	ErrorMsg        string
	RelatedItemSet  DescribeTaskAttributeRelatedItemList
}

type DescribeTaskAttributeRelatedItem struct {
	Name  string
	Value string
}

type DescribeTaskAttributeOperationProgressList []DescribeTaskAttributeOperationProgress

func (list *DescribeTaskAttributeOperationProgressList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeTaskAttributeOperationProgress)
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

type DescribeTaskAttributeRelatedItemList []DescribeTaskAttributeRelatedItem

func (list *DescribeTaskAttributeRelatedItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeTaskAttributeRelatedItem)
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
