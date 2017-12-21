package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeviceCreateRequest struct {
	DeviceNum      string `position:"Query" name:"DeviceNum"`
	DevicePosition string `position:"Query" name:"DevicePosition"`
	DeviceName     string `position:"Query" name:"DeviceName"`
	DeviceType     int    `position:"Query" name:"DeviceType"`
	Sid            int64  `position:"Query" name:"Sid"`
}

func (r DeviceCreateRequest) Invoke(client *sdk.Client) (response *DeviceCreateResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeviceCreateRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "DeviceCreate", "", "")

	resp := struct {
		*responses.BaseResponse
		DeviceCreateResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeviceCreateResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeviceCreateResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
