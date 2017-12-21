package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ReportMinuteRequest struct {
	BeginDate string `position:"Query" name:"BeginDate"`
	EndDate   string `position:"Query" name:"EndDate"`
	Agsid     int64  `position:"Query" name:"Agsid"`
}

func (r ReportMinuteRequest) Invoke(client *sdk.Client) (response *ReportMinuteResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ReportMinuteRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ReportMinute", "", "")

	resp := struct {
		*responses.BaseResponse
		ReportMinuteResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ReportMinuteResponse

	err = client.DoAction(&req, &resp)
	return
}

type ReportMinuteResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}