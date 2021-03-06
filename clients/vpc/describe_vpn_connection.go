package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeVpnConnectionRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	VpnConnectionId      string `position:"Query" name:"VpnConnectionId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeVpnConnectionRequest) Invoke(client *sdk.Client) (resp *DescribeVpnConnectionResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "DescribeVpnConnection", "vpc", "")
	resp = &DescribeVpnConnectionResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeVpnConnectionResponse struct {
	responses.BaseResponse
	RequestId         string
	VpnConnectionId   string
	CustomerGatewayId string
	VpnGatewayId      string
	Name              string
	LocalSubnet       string
	RemoteSubnet      string
	CreateTime        int64
	EffectImmediately bool
	Status            string
	IkeConfig         DescribeVpnConnectionIkeConfig
	IpsecConfig       DescribeVpnConnectionIpsecConfig
}

type DescribeVpnConnectionIkeConfig struct {
	Psk         string
	IkeVersion  string
	IkeMode     string
	IkeEncAlg   string
	IkeAuthAlg  string
	IkePfs      string
	IkeLifetime int64
	LocalId     string
	RemoteId    string
}

type DescribeVpnConnectionIpsecConfig struct {
	IpsecEncAlg   string
	IpsecAuthAlg  string
	IpsecPfs      string
	IpsecLifetime int64
}
