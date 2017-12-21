package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeRealtimeDiagnosesRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int    `position:"Query" name:"PageSize"`
	EndTime              string `position:"Query" name:"EndTime"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	StartTime            string `position:"Query" name:"StartTime"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
}

func (req *DescribeRealtimeDiagnosesRequest) Invoke(client *sdk.Client) (resp *DescribeRealtimeDiagnosesResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeRealtimeDiagnoses", "rds", "")
	resp = &DescribeRealtimeDiagnosesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeRealtimeDiagnosesResponse struct {
	responses.BaseResponse
	RequestId        string
	Engine           string
	TotalRecordCount int
	PageNumber       int
	PageRecordCount  int
	Tasks            DescribeRealtimeDiagnosesRealtimeDiagnoseTasksList
}

type DescribeRealtimeDiagnosesRealtimeDiagnoseTasks struct {
	CreateTime  string
	TaskId      string
	HealthScore string
	Status      string
}

type DescribeRealtimeDiagnosesRealtimeDiagnoseTasksList []DescribeRealtimeDiagnosesRealtimeDiagnoseTasks

func (list *DescribeRealtimeDiagnosesRealtimeDiagnoseTasksList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRealtimeDiagnosesRealtimeDiagnoseTasks)
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
