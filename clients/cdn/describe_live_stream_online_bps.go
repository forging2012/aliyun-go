package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeLiveStreamOnlineBpsRequest struct {
	requests.RpcRequest
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	StreamName    string `position:"Query" name:"StreamName"`
}

func (req *DescribeLiveStreamOnlineBpsRequest) Invoke(client *sdk.Client) (resp *DescribeLiveStreamOnlineBpsResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeLiveStreamOnlineBps", "", "")
	resp = &DescribeLiveStreamOnlineBpsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeLiveStreamOnlineBpsResponse struct {
	responses.BaseResponse
	RequestId                string
	TotalUserNumber          int64
	FlvBps                   float32
	HlsBps                   float32
	LiveStreamOnlineBpsInfos DescribeLiveStreamOnlineBpsLiveStreamOnlineBpsInfoList
}

type DescribeLiveStreamOnlineBpsLiveStreamOnlineBpsInfo struct {
	StreamUrl string
	UpBps     float32
	DownBps   float32
	Time      string
}

type DescribeLiveStreamOnlineBpsLiveStreamOnlineBpsInfoList []DescribeLiveStreamOnlineBpsLiveStreamOnlineBpsInfo

func (list *DescribeLiveStreamOnlineBpsLiveStreamOnlineBpsInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamOnlineBpsLiveStreamOnlineBpsInfo)
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
