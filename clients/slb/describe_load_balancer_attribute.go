package slb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeLoadBalancerAttributeRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	LoadBalancerId       string `position:"Query" name:"LoadBalancerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tags                 string `position:"Query" name:"Tags"`
}

func (r DescribeLoadBalancerAttributeRequest) Invoke(client *sdk.Client) (response *DescribeLoadBalancerAttributeResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeLoadBalancerAttributeRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Slb", "2014-05-15", "DescribeLoadBalancerAttribute", "slb", "")

	resp := struct {
		*responses.BaseResponse
		DescribeLoadBalancerAttributeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeLoadBalancerAttributeResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeLoadBalancerAttributeResponse struct {
	RequestId                string
	LoadBalancerId           string
	ResourceGroupId          string
	LoadBalancerName         string
	LoadBalancerStatus       string
	RegionId                 string
	RegionIdAlias            string
	Address                  string
	AddressType              string
	VpcId                    string
	VSwitchId                string
	NetworkType              string
	InternetChargeType       string
	AutoReleaseTime          int64
	Bandwidth                int
	LoadBalancerSpec         string
	CreateTime               string
	CreateTimeStamp          int64
	EndTime                  string
	EndTimeStamp             int64
	PayType                  string
	MasterZoneId             string
	SlaveZoneId              string
	ListenerPortsAndProtocal DescribeLoadBalancerAttributeListenerPortAndProtocalList
	ListenerPortsAndProtocol DescribeLoadBalancerAttributeListenerPortAndProtocolList
	BackendServers           DescribeLoadBalancerAttributeBackendServerList
	ListenerPorts            DescribeLoadBalancerAttributeListenerPortList
}

type DescribeLoadBalancerAttributeListenerPortAndProtocal struct {
	ListenerPort     int
	ListenerProtocal string
}

type DescribeLoadBalancerAttributeListenerPortAndProtocol struct {
	ListenerPort     int
	ListenerProtocol string
}

type DescribeLoadBalancerAttributeBackendServer struct {
	ServerId string
	Weight   int
}

type DescribeLoadBalancerAttributeListenerPortAndProtocalList []DescribeLoadBalancerAttributeListenerPortAndProtocal

func (list *DescribeLoadBalancerAttributeListenerPortAndProtocalList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLoadBalancerAttributeListenerPortAndProtocal)
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

type DescribeLoadBalancerAttributeListenerPortAndProtocolList []DescribeLoadBalancerAttributeListenerPortAndProtocol

func (list *DescribeLoadBalancerAttributeListenerPortAndProtocolList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLoadBalancerAttributeListenerPortAndProtocol)
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

type DescribeLoadBalancerAttributeBackendServerList []DescribeLoadBalancerAttributeBackendServer

func (list *DescribeLoadBalancerAttributeBackendServerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLoadBalancerAttributeBackendServer)
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

type DescribeLoadBalancerAttributeListenerPortList []string

func (list *DescribeLoadBalancerAttributeListenerPortList) UnmarshalJSON(data []byte) error {
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
