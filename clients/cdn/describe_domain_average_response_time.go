package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDomainAverageResponseTimeRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	TimeMerge     string `position:"Query" name:"TimeMerge"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	Interval      string `position:"Query" name:"Interval"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeDomainAverageResponseTimeRequest) Invoke(client *sdk.Client) (resp *DescribeDomainAverageResponseTimeResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeDomainAverageResponseTime", "", "")
	resp = &DescribeDomainAverageResponseTimeResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDomainAverageResponseTimeResponse struct {
	responses.BaseResponse
	RequestId        string
	DomainName       string
	DataInterval     string
	StartTime        string
	EndTime          string
	AvgRTPerInterval DescribeDomainAverageResponseTimeDataModuleList
}

type DescribeDomainAverageResponseTimeDataModule struct {
	TimeStamp string
	Value     string
}

type DescribeDomainAverageResponseTimeDataModuleList []DescribeDomainAverageResponseTimeDataModule

func (list *DescribeDomainAverageResponseTimeDataModuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainAverageResponseTimeDataModule)
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
