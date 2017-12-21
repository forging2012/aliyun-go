package vpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeAccessPointsRequest struct {
	Filters              *DescribeAccessPointsFilterList `position:"Query" type:"Repeated" name:"Filter"`
	ResourceOwnerId      int64                           `position:"Query" name:"ResourceOwnerId"`
	HostOperator         string                          `position:"Query" name:"HostOperator"`
	ResourceOwnerAccount string                          `position:"Query" name:"ResourceOwnerAccount"`
	Name                 string                          `position:"Query" name:"Name"`
	PageSize             int                             `position:"Query" name:"PageSize"`
	OwnerId              int64                           `position:"Query" name:"OwnerId"`
	Type                 string                          `position:"Query" name:"Type"`
	PageNumber           int                             `position:"Query" name:"PageNumber"`
}

func (r DescribeAccessPointsRequest) Invoke(client *sdk.Client) (response *DescribeAccessPointsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeAccessPointsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Vpc", "2016-04-28", "DescribeAccessPoints", "vpc", "")

	resp := struct {
		*responses.BaseResponse
		DescribeAccessPointsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeAccessPointsResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeAccessPointsFilter struct {
	Key    string                         `name:"Key"`
	Values *DescribeAccessPointsValueList `type:"Repeated" name:"Value"`
}

type DescribeAccessPointsResponse struct {
	RequestId      string
	PageNumber     int
	PageSize       int
	TotalCount     int
	AccessPointSet DescribeAccessPointsAccessPointTypeList
}

type DescribeAccessPointsAccessPointType struct {
	AccessPointId    string
	Status           string
	Type             string
	AttachedRegionNo string
	Location         string
	HostOperator     string
	Name             string
	Description      string
}

type DescribeAccessPointsFilterList []DescribeAccessPointsFilter

func (list *DescribeAccessPointsFilterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeAccessPointsFilter)
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

type DescribeAccessPointsValueList []string

func (list *DescribeAccessPointsValueList) UnmarshalJSON(data []byte) error {
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

type DescribeAccessPointsAccessPointTypeList []DescribeAccessPointsAccessPointType

func (list *DescribeAccessPointsAccessPointTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeAccessPointsAccessPointType)
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
