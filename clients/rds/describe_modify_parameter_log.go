package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeModifyParameterLogRequest struct {
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

func (req *DescribeModifyParameterLogRequest) Invoke(client *sdk.Client) (resp *DescribeModifyParameterLogResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeModifyParameterLog", "rds", "")
	resp = &DescribeModifyParameterLogResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeModifyParameterLogResponse struct {
	responses.BaseResponse
	RequestId        string
	Engine           string
	DBInstanceId     string
	EngineVersion    string
	TotalRecordCount int
	PageNumber       int
	PageRecordCount  int
	Items            DescribeModifyParameterLogParameterChangeLogList
}

type DescribeModifyParameterLogParameterChangeLog struct {
	ModifyTime        string
	OldParameterValue string
	NewParameterValue string
	ParameterName     string
	Status            string
}

type DescribeModifyParameterLogParameterChangeLogList []DescribeModifyParameterLogParameterChangeLog

func (list *DescribeModifyParameterLogParameterChangeLogList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeModifyParameterLogParameterChangeLog)
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
