package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteLiveSnapshotDetectPornConfigRequest struct {
	requests.RpcRequest
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *DeleteLiveSnapshotDetectPornConfigRequest) Invoke(client *sdk.Client) (resp *DeleteLiveSnapshotDetectPornConfigResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "DeleteLiveSnapshotDetectPornConfig", "", "")
	resp = &DeleteLiveSnapshotDetectPornConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteLiveSnapshotDetectPornConfigResponse struct {
	responses.BaseResponse
	RequestId string
}
