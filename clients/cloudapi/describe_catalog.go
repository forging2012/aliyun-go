package cloudapi

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeCatalogRequest struct {
	CatalogId string `position:"Query" name:"CatalogId"`
}

func (r DescribeCatalogRequest) Invoke(client *sdk.Client) (response *DescribeCatalogResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeCatalogRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeCatalog", "apigateway", "")

	resp := struct {
		*responses.BaseResponse
		DescribeCatalogResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeCatalogResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeCatalogResponse struct {
	RequestId    string
	CatalogId    string
	CatalogName  string
	Description  string
	ParentId     string
	CreatedTime  string
	ModifiedTime string
	RegionId     string
	ApiIds       DescribeCatalogApiIdList
}

type DescribeCatalogApiIdList []string

func (list *DescribeCatalogApiIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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
