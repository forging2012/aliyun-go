package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeNetworkInterfacesRequest struct {
	ResourceOwnerId      int64                                            `position:"Query" name:"ResourceOwnerId"`
	SecurityGroupId      string                                           `position:"Query" name:"SecurityGroupId"`
	Type                 string                                           `position:"Query" name:"Type"`
	PageNumber           int                                              `position:"Query" name:"PageNumber"`
	PageSize             int                                              `position:"Query" name:"PageSize"`
	NetworkInterfaceName string                                           `position:"Query" name:"NetworkInterfaceName"`
	ResourceOwnerAccount string                                           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string                                           `position:"Query" name:"OwnerAccount"`
	OwnerId              int64                                            `position:"Query" name:"OwnerId"`
	VSwitchId            string                                           `position:"Query" name:"VSwitchId"`
	InstanceId           string                                           `position:"Query" name:"InstanceId"`
	PrimaryIpAddress     string                                           `position:"Query" name:"PrimaryIpAddress"`
	NetworkInterfaceIds  *DescribeNetworkInterfacesNetworkInterfaceIdList `position:"Query" type:"Repeated" name:"NetworkInterfaceId"`
}

func (r DescribeNetworkInterfacesRequest) Invoke(client *sdk.Client) (response *DescribeNetworkInterfacesResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeNetworkInterfacesRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeNetworkInterfaces", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		DescribeNetworkInterfacesResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeNetworkInterfacesResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeNetworkInterfacesResponse struct {
	RequestId            string
	TotalCount           int
	PageNumber           int
	PageSize             int
	NetworkInterfaceSets DescribeNetworkInterfacesNetworkInterfaceSetList
}

type DescribeNetworkInterfacesNetworkInterfaceSet struct {
	NetworkInterfaceId   string
	Status               string
	Type                 string
	VpcId                string
	VSwitchId            string
	ZoneId               string
	PrivateIpAddress     string
	MacAddress           string
	NetworkInterfaceName string
	Description          string
	InstanceId           string
	CreationTime         string
	PrivateIpSets        DescribeNetworkInterfacesPrivateIpSetList
	SecurityGroupIds     DescribeNetworkInterfacesSecurityGroupIdList
	AssociatedPublicIp   DescribeNetworkInterfacesAssociatedPublicIp
}

type DescribeNetworkInterfacesPrivateIpSet struct {
	PrivateIpAddress    string
	Primary             bool
	AssociatedPublicIp1 DescribeNetworkInterfacesAssociatedPublicIp1
}

type DescribeNetworkInterfacesAssociatedPublicIp1 struct {
	PublicIpAddress string
	AllocationId    string
}

type DescribeNetworkInterfacesAssociatedPublicIp struct {
	PublicIpAddress string
	AllocationId    string
}

type DescribeNetworkInterfacesNetworkInterfaceIdList []string

func (list *DescribeNetworkInterfacesNetworkInterfaceIdList) UnmarshalJSON(data []byte) error {
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

type DescribeNetworkInterfacesNetworkInterfaceSetList []DescribeNetworkInterfacesNetworkInterfaceSet

func (list *DescribeNetworkInterfacesNetworkInterfaceSetList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeNetworkInterfacesNetworkInterfaceSet)
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

type DescribeNetworkInterfacesPrivateIpSetList []DescribeNetworkInterfacesPrivateIpSet

func (list *DescribeNetworkInterfacesPrivateIpSetList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeNetworkInterfacesPrivateIpSet)
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

type DescribeNetworkInterfacesSecurityGroupIdList []string

func (list *DescribeNetworkInterfacesSecurityGroupIdList) UnmarshalJSON(data []byte) error {
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
