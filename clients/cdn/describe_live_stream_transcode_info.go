package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeLiveStreamTranscodeInfoRequest struct {
	requests.RpcRequest
	SecurityToken       string `position:"Query" name:"SecurityToken"`
	OwnerId             int64  `position:"Query" name:"OwnerId"`
	DomainTranscodeName string `position:"Query" name:"DomainTranscodeName"`
}

func (req *DescribeLiveStreamTranscodeInfoRequest) Invoke(client *sdk.Client) (resp *DescribeLiveStreamTranscodeInfoResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeLiveStreamTranscodeInfo", "", "")
	resp = &DescribeLiveStreamTranscodeInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeLiveStreamTranscodeInfoResponse struct {
	responses.BaseResponse
	RequestId           string
	DomainTranscodeList DescribeLiveStreamTranscodeInfoDomainTranscodeInfoList
}

type DescribeLiveStreamTranscodeInfoDomainTranscodeInfo struct {
	TranscodeApp      string
	TranscodeId       string
	TranscodeName     string
	TranscodeRecord   string
	TranscodeSnapshot string
	TranscodeTemplate string
}

type DescribeLiveStreamTranscodeInfoDomainTranscodeInfoList []DescribeLiveStreamTranscodeInfoDomainTranscodeInfo

func (list *DescribeLiveStreamTranscodeInfoDomainTranscodeInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamTranscodeInfoDomainTranscodeInfo)
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
