package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDomainUvDataRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeDomainUvDataRequest) Invoke(client *sdk.Client) (response *DescribeDomainUvDataResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeDomainUvDataRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeDomainUvData", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeDomainUvDataResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeDomainUvDataResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeDomainUvDataResponse struct {
	RequestId      string
	DomainName     string
	DataInterval   string
	StartTime      string
	EndTime        string
	UvDataInterval DescribeDomainUvDataUsageDataList
}

type DescribeDomainUvDataUsageData struct {
	TimeStamp string
	Value     string
}

type DescribeDomainUvDataUsageDataList []DescribeDomainUvDataUsageData

func (list *DescribeDomainUvDataUsageDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainUvDataUsageData)
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
