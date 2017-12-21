package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeErrorLogsRequest struct {
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

func (r DescribeErrorLogsRequest) Invoke(client *sdk.Client) (response *DescribeErrorLogsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeErrorLogsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeErrorLogs", "rds", "")

	resp := struct {
		*responses.BaseResponse
		DescribeErrorLogsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeErrorLogsResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeErrorLogsResponse struct {
	RequestId        string
	TotalRecordCount int
	PageNumber       int
	PageRecordCount  int
	Items            DescribeErrorLogsErrorLogList
}

type DescribeErrorLogsErrorLog struct {
	ErrorInfo  string
	CreateTime string
}

type DescribeErrorLogsErrorLogList []DescribeErrorLogsErrorLog

func (list *DescribeErrorLogsErrorLogList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeErrorLogsErrorLog)
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
