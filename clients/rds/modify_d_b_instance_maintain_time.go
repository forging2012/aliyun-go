package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyDBInstanceMaintainTimeRequest struct {
	MaintainTime         string `position:"Query" name:"MaintainTime"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r ModifyDBInstanceMaintainTimeRequest) Invoke(client *sdk.Client) (response *ModifyDBInstanceMaintainTimeResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyDBInstanceMaintainTimeRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "ModifyDBInstanceMaintainTime", "rds", "")

	resp := struct {
		*responses.BaseResponse
		ModifyDBInstanceMaintainTimeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyDBInstanceMaintainTimeResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyDBInstanceMaintainTimeResponse struct {
	RequestId string
}
