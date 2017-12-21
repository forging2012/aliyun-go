package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddTagsToResourceRequest struct {
	Tag4value            string `position:"Query" name:"Tag.4.value"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Tag2key              string `position:"Query" name:"Tag.2.key"`
	Tag5key              string `position:"Query" name:"Tag.5.key"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Tag3key              string `position:"Query" name:"Tag.3.key"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tag5value            string `position:"Query" name:"Tag.5.value"`
	Tags                 string `position:"Query" name:"Tags"`
	Tag1key              string `position:"Query" name:"Tag.1.key"`
	Tag1value            string `position:"Query" name:"Tag.1.value"`
	Tag2value            string `position:"Query" name:"Tag.2.value"`
	Tag4key              string `position:"Query" name:"Tag.4.key"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	Tag3value            string `position:"Query" name:"Tag.3.value"`
	ProxyId              string `position:"Query" name:"ProxyId"`
}

func (r AddTagsToResourceRequest) Invoke(client *sdk.Client) (response *AddTagsToResourceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		AddTagsToResourceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "AddTagsToResource", "rds", "")

	resp := struct {
		*responses.BaseResponse
		AddTagsToResourceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.AddTagsToResourceResponse

	err = client.DoAction(&req, &resp)
	return
}

type AddTagsToResourceResponse struct {
	RequestId string
}