package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyReadWriteSplittingConnectionRequest struct {
	ResourceOwnerId        int64  `position:"Query" name:"ResourceOwnerId"`
	ConnectionStringPrefix string `position:"Query" name:"ConnectionStringPrefix"`
	ResourceOwnerAccount   string `position:"Query" name:"ResourceOwnerAccount"`
	Port                   string `position:"Query" name:"Port"`
	DistributionType       string `position:"Query" name:"DistributionType"`
	OwnerAccount           string `position:"Query" name:"OwnerAccount"`
	Weight                 string `position:"Query" name:"Weight"`
	DBInstanceId           string `position:"Query" name:"DBInstanceId"`
	OwnerId                int64  `position:"Query" name:"OwnerId"`
	MaxDelayTime           string `position:"Query" name:"MaxDelayTime"`
}

func (r ModifyReadWriteSplittingConnectionRequest) Invoke(client *sdk.Client) (response *ModifyReadWriteSplittingConnectionResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyReadWriteSplittingConnectionRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "ModifyReadWriteSplittingConnection", "rds", "")

	resp := struct {
		*responses.BaseResponse
		ModifyReadWriteSplittingConnectionResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyReadWriteSplittingConnectionResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyReadWriteSplittingConnectionResponse struct {
	RequestId string
}
