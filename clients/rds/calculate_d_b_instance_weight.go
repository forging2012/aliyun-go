package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CalculateDBInstanceWeightRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *CalculateDBInstanceWeightRequest) Invoke(client *sdk.Client) (resp *CalculateDBInstanceWeightResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "CalculateDBInstanceWeight", "rds", "")
	resp = &CalculateDBInstanceWeightResponse{}
	err = client.DoAction(req, resp)
	return
}

type CalculateDBInstanceWeightResponse struct {
	responses.BaseResponse
	RequestId string
	Items     CalculateDBInstanceWeightDBInstanceWeightList
}

type CalculateDBInstanceWeightDBInstanceWeight struct {
	DBInstanceId   string
	DBInstanceType string
	Availability   string
	Weight         string
}

type CalculateDBInstanceWeightDBInstanceWeightList []CalculateDBInstanceWeightDBInstanceWeight

func (list *CalculateDBInstanceWeightDBInstanceWeightList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CalculateDBInstanceWeightDBInstanceWeight)
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
