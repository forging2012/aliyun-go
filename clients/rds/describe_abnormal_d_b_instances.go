package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeAbnormalDBInstancesRequest struct {
	requests.RpcRequest
	Tag4value            string `position:"Query" name:"Tag.4.value"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Tag2key              string `position:"Query" name:"Tag.2.key"`
	Tag5key              string `position:"Query" name:"Tag.5.key"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Tag3key              string `position:"Query" name:"Tag.3.key"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tag5value            string `position:"Query" name:"Tag.5.value"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
	Tags                 string `position:"Query" name:"Tags"`
	Tag1key              string `position:"Query" name:"Tag.1.key"`
	Tag1value            string `position:"Query" name:"Tag.1.value"`
	Tag2value            string `position:"Query" name:"Tag.2.value"`
	PageSize             int    `position:"Query" name:"PageSize"`
	Tag4key              string `position:"Query" name:"Tag.4.key"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	Tag3value            string `position:"Query" name:"Tag.3.value"`
	ProxyId              string `position:"Query" name:"ProxyId"`
}

func (req *DescribeAbnormalDBInstancesRequest) Invoke(client *sdk.Client) (resp *DescribeAbnormalDBInstancesResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeAbnormalDBInstances", "rds", "")
	resp = &DescribeAbnormalDBInstancesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeAbnormalDBInstancesResponse struct {
	responses.BaseResponse
	RequestId        string
	TotalRecordCount int
	PageNumber       int
	PageRecordCount  int
	Items            DescribeAbnormalDBInstancesInstanceResultList
}

type DescribeAbnormalDBInstancesInstanceResult struct {
	DBInstanceDescription string
	DBInstanceId          string
	AbnormalItems         DescribeAbnormalDBInstancesAbnormalItemList
}

type DescribeAbnormalDBInstancesAbnormalItem struct {
	CheckTime      string
	CheckItem      string
	AbnormalReason string
	AbnormalValue  string
	AbnormalDetail string
	AdviceKey      string
	AdviseValue    DescribeAbnormalDBInstancesAdviseValueList
}

type DescribeAbnormalDBInstancesInstanceResultList []DescribeAbnormalDBInstancesInstanceResult

func (list *DescribeAbnormalDBInstancesInstanceResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeAbnormalDBInstancesInstanceResult)
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

type DescribeAbnormalDBInstancesAbnormalItemList []DescribeAbnormalDBInstancesAbnormalItem

func (list *DescribeAbnormalDBInstancesAbnormalItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeAbnormalDBInstancesAbnormalItem)
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

type DescribeAbnormalDBInstancesAdviseValueList []string

func (list *DescribeAbnormalDBInstancesAdviseValueList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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
