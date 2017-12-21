package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeLiveStreamBpsDataRequest struct {
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	StreamName    string `position:"Query" name:"StreamName"`
}

func (r DescribeLiveStreamBpsDataRequest) Invoke(client *sdk.Client) (response *DescribeLiveStreamBpsDataResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeLiveStreamBpsDataRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeLiveStreamBpsData", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeLiveStreamBpsDataResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeLiveStreamBpsDataResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeLiveStreamBpsDataResponse struct {
	RequestId string
	BpsDatas  DescribeLiveStreamBpsDataDomainBpsModelList
}

type DescribeLiveStreamBpsDataDomainBpsModel struct {
	Time string
	Bps  float32
}

type DescribeLiveStreamBpsDataDomainBpsModelList []DescribeLiveStreamBpsDataDomainBpsModel

func (list *DescribeLiveStreamBpsDataDomainBpsModelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamBpsDataDomainBpsModel)
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
