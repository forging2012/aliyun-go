package dds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyDBInstanceNetworkTypeRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	NetworkType          string `position:"Query" name:"NetworkType"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	VSwitchId            string `position:"Query" name:"VSwitchId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	RetainClassic        string `position:"Query" name:"RetainClassic"`
	ClassicExpiredDays   int    `position:"Query" name:"ClassicExpiredDays"`
	VpcId                string `position:"Query" name:"VpcId"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
}

func (r ModifyDBInstanceNetworkTypeRequest) Invoke(client *sdk.Client) (response *ModifyDBInstanceNetworkTypeResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyDBInstanceNetworkTypeRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Dds", "2015-12-01", "ModifyDBInstanceNetworkType", "dds", "")

	resp := struct {
		*responses.BaseResponse
		ModifyDBInstanceNetworkTypeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyDBInstanceNetworkTypeResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyDBInstanceNetworkTypeResponse struct {
	RequestId string
}
