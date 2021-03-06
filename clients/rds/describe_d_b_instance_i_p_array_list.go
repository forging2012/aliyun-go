package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDBInstanceIPArrayListRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeDBInstanceIPArrayListRequest) Invoke(client *sdk.Client) (resp *DescribeDBInstanceIPArrayListResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeDBInstanceIPArrayList", "rds", "")
	resp = &DescribeDBInstanceIPArrayListResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDBInstanceIPArrayListResponse struct {
	responses.BaseResponse
	RequestId string
	Items     DescribeDBInstanceIPArrayListDBInstanceIPArrayList
}

type DescribeDBInstanceIPArrayListDBInstanceIPArray struct {
	DBInstanceIPArrayName      string
	DBInstanceIPArrayAttribute string
	SecurityIPList             string
}

type DescribeDBInstanceIPArrayListDBInstanceIPArrayList []DescribeDBInstanceIPArrayListDBInstanceIPArray

func (list *DescribeDBInstanceIPArrayListDBInstanceIPArrayList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstanceIPArrayListDBInstanceIPArray)
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
