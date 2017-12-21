package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteAlarmRequest struct {
	Callby_cms_owner string `position:"Query" name:"Callby_cms_owner"`
	Id               string `position:"Query" name:"Id"`
}

func (r DeleteAlarmRequest) Invoke(client *sdk.Client) (response *DeleteAlarmResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteAlarmRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cms", "2017-03-01", "DeleteAlarm", "cms", "")

	resp := struct {
		*responses.BaseResponse
		DeleteAlarmResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteAlarmResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteAlarmResponse struct {
	Success   bool
	Code      string
	Message   string
	RequestId string
}