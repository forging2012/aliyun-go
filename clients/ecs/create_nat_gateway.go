package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateNatGatewayRequest struct {
	ResourceOwnerId      int64                                 `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string                                `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string                                `position:"Query" name:"ClientToken"`
	OwnerAccount         string                                `position:"Query" name:"OwnerAccount"`
	VpcId                string                                `position:"Query" name:"VpcId"`
	Name                 string                                `position:"Query" name:"Name"`
	Description          string                                `position:"Query" name:"Description"`
	OwnerId              int64                                 `position:"Query" name:"OwnerId"`
	BandwidthPackages    *CreateNatGatewayBandwidthPackageList `position:"Query" type:"Repeated" name:"BandwidthPackage"`
}

func (r CreateNatGatewayRequest) Invoke(client *sdk.Client) (response *CreateNatGatewayResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateNatGatewayRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "CreateNatGateway", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		CreateNatGatewayResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateNatGatewayResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateNatGatewayBandwidthPackage struct {
	IpCount   int    `name:"IpCount"`
	Bandwidth int    `name:"Bandwidth"`
	Zone      string `name:"Zone"`
}

type CreateNatGatewayResponse struct {
	RequestId           string
	NatGatewayId        string
	ForwardTableIds     CreateNatGatewayForwardTableIdList
	BandwidthPackageIds CreateNatGatewayBandwidthPackageIdList
}

type CreateNatGatewayBandwidthPackageList []CreateNatGatewayBandwidthPackage

func (list *CreateNatGatewayBandwidthPackageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateNatGatewayBandwidthPackage)
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

type CreateNatGatewayForwardTableIdList []string

func (list *CreateNatGatewayForwardTableIdList) UnmarshalJSON(data []byte) error {
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

type CreateNatGatewayBandwidthPackageIdList []string

func (list *CreateNatGatewayBandwidthPackageIdList) UnmarshalJSON(data []byte) error {
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
