package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeOperatorPermissionRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeOperatorPermissionRequest) Invoke(client *sdk.Client) (response *DescribeOperatorPermissionResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeOperatorPermissionRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeOperatorPermission", "rds", "")

	resp := struct {
		*responses.BaseResponse
		DescribeOperatorPermissionResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeOperatorPermissionResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeOperatorPermissionResponse struct {
	RequestId   string
	Privileges  string
	CreatedTime string
	ExpiredTime string
}
