package polardb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeParametersRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeParametersRequest) Invoke(client *sdk.Client) (response *DescribeParametersResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeParametersRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("polardb", "2017-08-01", "DescribeParameters", "polardb", "")

	resp := struct {
		*responses.BaseResponse
		DescribeParametersResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeParametersResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeParametersResponse struct {
	RequestId         string
	Engine            string
	DBType            string
	DBVersion         string
	ConfigParameters  DescribeParametersDBInstanceParameterList
	RunningParameters DescribeParametersDBInstanceParameterList
}

type DescribeParametersDBInstanceParameter struct {
	ParameterName        string
	ParameterValue       string
	ParameterDescription string
}

type DescribeParametersDBInstanceParameterList []DescribeParametersDBInstanceParameter

func (list *DescribeParametersDBInstanceParameterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeParametersDBInstanceParameter)
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
