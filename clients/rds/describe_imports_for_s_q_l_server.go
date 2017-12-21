package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeImportsForSQLServerRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ImportId             int    `position:"Query" name:"ImportId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int    `position:"Query" name:"PageSize"`
	EndTime              string `position:"Query" name:"EndTime"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	StartTime            string `position:"Query" name:"StartTime"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
}

func (req *DescribeImportsForSQLServerRequest) Invoke(client *sdk.Client) (resp *DescribeImportsForSQLServerResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeImportsForSQLServer", "rds", "")
	resp = &DescribeImportsForSQLServerResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeImportsForSQLServerResponse struct {
	responses.BaseResponse
	RequestId         string
	TotalRecordCounts int
	PageNumber        int
	SQLItemsCounts    int
	Items             DescribeImportsForSQLServerSQLServerImportList
}

type DescribeImportsForSQLServerSQLServerImport struct {
	ImportId     int
	FileName     string
	DBName       string
	ImportStatus string
	StartTime    string
}

type DescribeImportsForSQLServerSQLServerImportList []DescribeImportsForSQLServerSQLServerImport

func (list *DescribeImportsForSQLServerSQLServerImportList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeImportsForSQLServerSQLServerImport)
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
