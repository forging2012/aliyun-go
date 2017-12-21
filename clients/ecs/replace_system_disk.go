package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ReplaceSystemDiskRequest struct {
	ResourceOwnerId             int64  `position:"Query" name:"ResourceOwnerId"`
	ImageId                     string `position:"Query" name:"ImageId"`
	ResourceOwnerAccount        string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken                 string `position:"Query" name:"ClientToken"`
	OwnerAccount                string `position:"Query" name:"OwnerAccount"`
	SecurityEnhancementStrategy string `position:"Query" name:"SecurityEnhancementStrategy"`
	KeyPairName                 string `position:"Query" name:"KeyPairName"`
	OwnerId                     int64  `position:"Query" name:"OwnerId"`
	Password                    string `position:"Query" name:"Password"`
	InstanceId                  string `position:"Query" name:"InstanceId"`
	SystemDiskSize              int    `position:"Query" name:"SystemDiskSize"`
	UseAdditionalService        string `position:"Query" name:"UseAdditionalService"`
}

func (r ReplaceSystemDiskRequest) Invoke(client *sdk.Client) (response *ReplaceSystemDiskResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ReplaceSystemDiskRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "ReplaceSystemDisk", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		ReplaceSystemDiskResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ReplaceSystemDiskResponse

	err = client.DoAction(&req, &resp)
	return
}

type ReplaceSystemDiskResponse struct {
	RequestId string
	DiskId    string
}
