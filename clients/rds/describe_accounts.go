package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeAccountsRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	AccountName          string `position:"Query" name:"AccountName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeAccountsRequest) Invoke(client *sdk.Client) (resp *DescribeAccountsResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeAccounts", "rds", "")
	resp = &DescribeAccountsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeAccountsResponse struct {
	responses.BaseResponse
	RequestId string
	Accounts  DescribeAccountsDBInstanceAccountList
}

type DescribeAccountsDBInstanceAccount struct {
	DBInstanceId       string
	AccountName        string
	AccountStatus      string
	AccountType        string
	AccountDescription string
	DatabasePrivileges DescribeAccountsDatabasePrivilegeList
}

type DescribeAccountsDatabasePrivilege struct {
	DBName           string
	AccountPrivilege string
}

type DescribeAccountsDBInstanceAccountList []DescribeAccountsDBInstanceAccount

func (list *DescribeAccountsDBInstanceAccountList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeAccountsDBInstanceAccount)
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

type DescribeAccountsDatabasePrivilegeList []DescribeAccountsDatabasePrivilege

func (list *DescribeAccountsDatabasePrivilegeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeAccountsDatabasePrivilege)
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
