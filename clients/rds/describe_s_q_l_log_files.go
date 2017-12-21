package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeSQLLogFilesRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	FileName             string `position:"Query" name:"FileName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int    `position:"Query" name:"PageSize"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
}

func (r DescribeSQLLogFilesRequest) Invoke(client *sdk.Client) (response *DescribeSQLLogFilesResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeSQLLogFilesRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeSQLLogFiles", "rds", "")

	resp := struct {
		*responses.BaseResponse
		DescribeSQLLogFilesResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeSQLLogFilesResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeSQLLogFilesResponse struct {
	RequestId        string
	TotalRecordCount int
	PageNumber       int
	PageRecordCount  int
	Items            DescribeSQLLogFilesLogFileList
}

type DescribeSQLLogFilesLogFile struct {
	FileID         string
	LogStatus      string
	LogDownloadURL string
	LogSize        string
	LogStartTime   string
	LogEndTime     string
}

type DescribeSQLLogFilesLogFileList []DescribeSQLLogFilesLogFile

func (list *DescribeSQLLogFilesLogFileList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSQLLogFilesLogFile)
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
