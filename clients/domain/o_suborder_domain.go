package domain

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type OSuborderDomainRequest struct {
	EndDate   string `position:"Query" name:"EndDate"`
	PageSize  int    `position:"Query" name:"PageSize"`
	Type      string `position:"Query" name:"Type"`
	StartDate string `position:"Query" name:"StartDate"`
	PageNum   int    `position:"Query" name:"PageNum"`
}

func (r OSuborderDomainRequest) Invoke(client *sdk.Client) (response *OSuborderDomainResponse, err error) {
	req := struct {
		*requests.RpcRequest
		OSuborderDomainRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Domain", "2016-05-11", "OSuborderDomain", "", "")

	resp := struct {
		*responses.BaseResponse
		OSuborderDomainResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.OSuborderDomainResponse

	err = client.DoAction(&req, &resp)
	return
}

type OSuborderDomainResponse struct {
	RequestId      string
	TotalItemNum   int
	CurrentPageNum int
	PageSize       int
	Data           OSuborderDomainObjectList
}

type OSuborderDomainObject struct {
	CommodityKey  string
	CommodityCode string
	Amount        int64
	SettleTime    string
}

type OSuborderDomainObjectList []OSuborderDomainObject

func (list *OSuborderDomainObjectList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OSuborderDomainObject)
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
