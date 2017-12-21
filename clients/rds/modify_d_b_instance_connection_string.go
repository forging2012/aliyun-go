package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyDBInstanceConnectionStringRequest struct {
	ResourceOwnerId         int64  `position:"Query" name:"ResourceOwnerId"`
	ConnectionStringPrefix  string `position:"Query" name:"ConnectionStringPrefix"`
	ResourceOwnerAccount    string `position:"Query" name:"ResourceOwnerAccount"`
	Port                    string `position:"Query" name:"Port"`
	OwnerAccount            string `position:"Query" name:"OwnerAccount"`
	DBInstanceId            string `position:"Query" name:"DBInstanceId"`
	OwnerId                 int64  `position:"Query" name:"OwnerId"`
	CurrentConnectionString string `position:"Query" name:"CurrentConnectionString"`
}

func (r ModifyDBInstanceConnectionStringRequest) Invoke(client *sdk.Client) (response *ModifyDBInstanceConnectionStringResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyDBInstanceConnectionStringRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "ModifyDBInstanceConnectionString", "rds", "")

	resp := struct {
		*responses.BaseResponse
		ModifyDBInstanceConnectionStringResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyDBInstanceConnectionStringResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyDBInstanceConnectionStringResponse struct {
	RequestId string
}
