package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeLiveStreamRelayPushErrorsRequest struct {
	requests.RpcRequest
	RelayDomain   string `position:"Query" name:"RelayDomain"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeLiveStreamRelayPushErrorsRequest) Invoke(client *sdk.Client) (resp *DescribeLiveStreamRelayPushErrorsResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeLiveStreamRelayPushErrors", "", "")
	resp = &DescribeLiveStreamRelayPushErrorsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeLiveStreamRelayPushErrorsResponse struct {
	responses.BaseResponse
	RequestId                string
	RelayPushErrorsModelList DescribeLiveStreamRelayPushErrorsRelayPushErrorsModelList
}

type DescribeLiveStreamRelayPushErrorsRelayPushErrorsModel struct {
	ErrorCode string
}

type DescribeLiveStreamRelayPushErrorsRelayPushErrorsModelList []DescribeLiveStreamRelayPushErrorsRelayPushErrorsModel

func (list *DescribeLiveStreamRelayPushErrorsRelayPushErrorsModelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamRelayPushErrorsRelayPushErrorsModel)
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
