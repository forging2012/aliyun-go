package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDomainTopUrlVisitRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	SortBy        string `position:"Query" name:"SortBy"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeDomainTopUrlVisitRequest) Invoke(client *sdk.Client) (resp *DescribeDomainTopUrlVisitResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeDomainTopUrlVisit", "", "")
	resp = &DescribeDomainTopUrlVisitResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDomainTopUrlVisitResponse struct {
	responses.BaseResponse
	RequestId  string
	DomainName string
	StartTime  string
	AllUrlList DescribeDomainTopUrlVisitUrlListList
	Url200List DescribeDomainTopUrlVisitUrlListList
	Url300List DescribeDomainTopUrlVisitUrlListList
	Url400List DescribeDomainTopUrlVisitUrlListList
	Url500List DescribeDomainTopUrlVisitUrlListList
}

type DescribeDomainTopUrlVisitUrlList struct {
	UrlDetail       string
	VisitData       string
	VisitProportion float32
	Flow            string
	FlowProportion  float32
}

type DescribeDomainTopUrlVisitUrlListList []DescribeDomainTopUrlVisitUrlList

func (list *DescribeDomainTopUrlVisitUrlListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainTopUrlVisitUrlList)
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
