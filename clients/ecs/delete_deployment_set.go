package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteDeploymentSetRequest struct {
	DeploymentSetId      string `position:"Query" name:"DeploymentSetId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DeleteDeploymentSetRequest) Invoke(client *sdk.Client) (response *DeleteDeploymentSetResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteDeploymentSetRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "DeleteDeploymentSet", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		DeleteDeploymentSetResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteDeploymentSetResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteDeploymentSetResponse struct {
	RequestId string
}
