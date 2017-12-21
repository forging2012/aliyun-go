package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDomainReqHitRateDataRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	Interval      string `position:"Query" name:"Interval"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeDomainReqHitRateDataRequest) Invoke(client *sdk.Client) (resp *DescribeDomainReqHitRateDataResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeDomainReqHitRateData", "", "")
	resp = &DescribeDomainReqHitRateDataResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDomainReqHitRateDataResponse struct {
	responses.BaseResponse
	RequestId          string
	DomainName         string
	DataInterval       string
	StartTime          string
	EndTime            string
	ReqHitRateInterval DescribeDomainReqHitRateDataDataModuleList
}

type DescribeDomainReqHitRateDataDataModule struct {
	TimeStamp string
	Value     string
}

type DescribeDomainReqHitRateDataDataModuleList []DescribeDomainReqHitRateDataDataModule

func (list *DescribeDomainReqHitRateDataDataModuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainReqHitRateDataDataModule)
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
