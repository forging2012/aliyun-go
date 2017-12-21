package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeOptimizeAdviceOnStorageRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int    `position:"Query" name:"PageSize"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
}

func (r DescribeOptimizeAdviceOnStorageRequest) Invoke(client *sdk.Client) (response *DescribeOptimizeAdviceOnStorageResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeOptimizeAdviceOnStorageRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeOptimizeAdviceOnStorage", "rds", "")

	resp := struct {
		*responses.BaseResponse
		DescribeOptimizeAdviceOnStorageResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeOptimizeAdviceOnStorageResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeOptimizeAdviceOnStorageResponse struct {
	RequestId         string
	DBInstanceId      string
	TotalRecordsCount int
	PageNumber        int
	PageRecordCount   int
	Items             DescribeOptimizeAdviceOnStorageAdviceOnStorageList
}

type DescribeOptimizeAdviceOnStorageAdviceOnStorage struct {
	DBName        string
	TableName     string
	CurrentEngine string
	AdviseEngine  string
}

type DescribeOptimizeAdviceOnStorageAdviceOnStorageList []DescribeOptimizeAdviceOnStorageAdviceOnStorage

func (list *DescribeOptimizeAdviceOnStorageAdviceOnStorageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeOptimizeAdviceOnStorageAdviceOnStorage)
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