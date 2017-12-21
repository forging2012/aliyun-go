package vpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeRouteTableListRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Bandwidth            string `position:"Query" name:"Bandwidth"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
	RouterType           string `position:"Query" name:"RouterType"`
	KbpsBandwidth        string `position:"Query" name:"KbpsBandwidth"`
	RouteTableName       string `position:"Query" name:"RouteTableName"`
	RouterId             string `position:"Query" name:"RouterId"`
	VpcId                string `position:"Query" name:"VpcId"`
	ResourceUid          int64  `position:"Query" name:"ResourceUid"`
	PageSize             int    `position:"Query" name:"PageSize"`
	ResourceBid          string `position:"Query" name:"ResourceBid"`
	RouteTableId         string `position:"Query" name:"RouteTableId"`
}

func (r DescribeRouteTableListRequest) Invoke(client *sdk.Client) (response *DescribeRouteTableListResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeRouteTableListRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Vpc", "2016-04-28", "DescribeRouteTableList", "vpc", "")

	resp := struct {
		*responses.BaseResponse
		DescribeRouteTableListResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeRouteTableListResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeRouteTableListResponse struct {
	RequestId       string
	Code            string
	Message         string
	Success         bool
	PageSize        int
	PageNumber      int
	TotalCount      int
	RouterTableList DescribeRouteTableListRouterTableListTypeList
}

type DescribeRouteTableListRouterTableListType struct {
	VpcId          string
	RouterType     string
	RouterId       string
	RouteTableId   string
	RouteTableName string
	RouteTableType string
	Description    string
	CreationTime   string
}

type DescribeRouteTableListRouterTableListTypeList []DescribeRouteTableListRouterTableListType

func (list *DescribeRouteTableListRouterTableListTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRouteTableListRouterTableListType)
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
