package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteLiveSnapshotDetectPornConfigRequest struct {
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r DeleteLiveSnapshotDetectPornConfigRequest) Invoke(client *sdk.Client) (response *DeleteLiveSnapshotDetectPornConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteLiveSnapshotDetectPornConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("live", "2016-11-01", "DeleteLiveSnapshotDetectPornConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		DeleteLiveSnapshotDetectPornConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteLiveSnapshotDetectPornConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteLiveSnapshotDetectPornConfigResponse struct {
	RequestId string
}
