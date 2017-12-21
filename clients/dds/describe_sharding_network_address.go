package dds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeShardingNetworkAddressRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	NodeId               string `position:"Query" name:"NodeId"`
}

func (r DescribeShardingNetworkAddressRequest) Invoke(client *sdk.Client) (response *DescribeShardingNetworkAddressResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeShardingNetworkAddressRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Dds", "2015-12-01", "DescribeShardingNetworkAddress", "dds", "")

	resp := struct {
		*responses.BaseResponse
		DescribeShardingNetworkAddressResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeShardingNetworkAddressResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeShardingNetworkAddressResponse struct {
	RequestId        string
	NetworkAddresses DescribeShardingNetworkAddressNetworkAddressList
}

type DescribeShardingNetworkAddressNetworkAddress struct {
	NetworkAddress string
	IPAddress      string
	NetworkType    string
	Port           string
	VPCId          string
	VswitchId      string
	NodeId         string
	ExpiredTime    string
}

type DescribeShardingNetworkAddressNetworkAddressList []DescribeShardingNetworkAddressNetworkAddress

func (list *DescribeShardingNetworkAddressNetworkAddressList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeShardingNetworkAddressNetworkAddress)
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
