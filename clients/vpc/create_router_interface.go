package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateRouterInterfaceRequest struct {
	AccessPointId            string `position:"Query" name:"AccessPointId"`
	OppositeRouterId         string `position:"Query" name:"OppositeRouterId"`
	OppositeAccessPointId    string `position:"Query" name:"OppositeAccessPointId"`
	ResourceOwnerId          int64  `position:"Query" name:"ResourceOwnerId"`
	Role                     string `position:"Query" name:"Role"`
	ClientToken              string `position:"Query" name:"ClientToken"`
	ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
	OppositeRegionId         string `position:"Query" name:"OppositeRegionId"`
	OwnerAccount             string `position:"Query" name:"OwnerAccount"`
	HealthCheckTargetIp      string `position:"Query" name:"HealthCheckTargetIp"`
	Description              string `position:"Query" name:"Description"`
	OwnerId                  int64  `position:"Query" name:"OwnerId"`
	Spec                     string `position:"Query" name:"Spec"`
	OppositeInterfaceOwnerId string `position:"Query" name:"OppositeInterfaceOwnerId"`
	RouterType               string `position:"Query" name:"RouterType"`
	HealthCheckSourceIp      string `position:"Query" name:"HealthCheckSourceIp"`
	RouterId                 string `position:"Query" name:"RouterId"`
	OppositeRouterType       string `position:"Query" name:"OppositeRouterType"`
	Name                     string `position:"Query" name:"Name"`
	UserCidr                 string `position:"Query" name:"UserCidr"`
	OppositeInterfaceId      string `position:"Query" name:"OppositeInterfaceId"`
}

func (r CreateRouterInterfaceRequest) Invoke(client *sdk.Client) (response *CreateRouterInterfaceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateRouterInterfaceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Vpc", "2016-04-28", "CreateRouterInterface", "vpc", "")

	resp := struct {
		*responses.BaseResponse
		CreateRouterInterfaceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateRouterInterfaceResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateRouterInterfaceResponse struct {
	RequestId         string
	RouterInterfaceId string
}
