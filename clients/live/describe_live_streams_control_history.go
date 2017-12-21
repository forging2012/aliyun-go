package live

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeLiveStreamsControlHistoryRequest struct {
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeLiveStreamsControlHistoryRequest) Invoke(client *sdk.Client) (response *DescribeLiveStreamsControlHistoryResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeLiveStreamsControlHistoryRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("live", "2016-11-01", "DescribeLiveStreamsControlHistory", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeLiveStreamsControlHistoryResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeLiveStreamsControlHistoryResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeLiveStreamsControlHistoryResponse struct {
	RequestId   string
	ControlInfo DescribeLiveStreamsControlHistoryLiveStreamControlInfoList
}

type DescribeLiveStreamsControlHistoryLiveStreamControlInfo struct {
	StreamName string
	ClientIP   string
	Action     string
	TimeStamp  string
}

type DescribeLiveStreamsControlHistoryLiveStreamControlInfoList []DescribeLiveStreamsControlHistoryLiveStreamControlInfo

func (list *DescribeLiveStreamsControlHistoryLiveStreamControlInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamsControlHistoryLiveStreamControlInfo)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
