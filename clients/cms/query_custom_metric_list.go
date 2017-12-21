package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryCustomMetricListRequest struct {
	Size       string `position:"Query" name:"Size"`
	GroupId    string `position:"Query" name:"GroupId"`
	Page       string `position:"Query" name:"Page"`
	MetricName string `position:"Query" name:"MetricName"`
	Dimension  string `position:"Query" name:"Dimension"`
	UUID       string `position:"Query" name:"UUID"`
	Md5        string `position:"Query" name:"Md.5"`
}

func (r QueryCustomMetricListRequest) Invoke(client *sdk.Client) (response *QueryCustomMetricListResponse, err error) {
	req := struct {
		*requests.RpcRequest
		QueryCustomMetricListRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cms", "2017-03-01", "QueryCustomMetricList", "cms", "")

	resp := struct {
		*responses.BaseResponse
		QueryCustomMetricListResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.QueryCustomMetricListResponse

	err = client.DoAction(&req, &resp)
	return
}

type QueryCustomMetricListResponse struct {
	Code      string
	Message   string
	RequestId string
	Result    string
}
