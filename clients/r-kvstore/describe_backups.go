package r_kvstore

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeBackupsRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	BackupId             int    `position:"Query" name:"BackupId"`
	EndTime              string `position:"Query" name:"EndTime"`
	StartTime            string `position:"Query" name:"StartTime"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	PageSize             int    `position:"Query" name:"PageSize"`
}

func (req *DescribeBackupsRequest) Invoke(client *sdk.Client) (resp *DescribeBackupsResponse, err error) {
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "DescribeBackups", "redisa", "")
	resp = &DescribeBackupsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeBackupsResponse struct {
	responses.BaseResponse
	RequestId  string
	PageNumber int
	PageSize   int
	TotalCount int
	Backups    DescribeBackupsBackupList
}

type DescribeBackupsBackup struct {
	BackupId          int
	BackupDBNames     string
	BackupStatus      string
	BackupStartTime   string
	BackupEndTime     string
	BackupType        string
	BackupMode        string
	BackupMethod      string
	BackupDownloadURL string
	BackupSize        int64
}

type DescribeBackupsBackupList []DescribeBackupsBackup

func (list *DescribeBackupsBackupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeBackupsBackup)
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
