package arms

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type WhereInDimQueryRequest struct {
	DateStr        string                            `position:"Query" name:"DateStr"`
	MinTime        int64                             `position:"Query" name:"MinTime"`
	ReduceTail     string                            `position:"Query" name:"ReduceTail"`
	MaxTime        int64                             `position:"Query" name:"MaxTime"`
	WhereInKey     string                            `position:"Query" name:"WhereInKey"`
	Measuress      *WhereInDimQueryMeasuresList      `position:"Query" type:"Repeated" name:"Measures"`
	IntervalInSec  int                               `position:"Query" name:"IntervalInSec"`
	IsDrillDown    string                            `position:"Query" name:"IsDrillDown"`
	OrderByKey     string                            `position:"Query" name:"OrderByKey"`
	Limit          int                               `position:"Query" name:"Limit"`
	DatasetId      int64                             `position:"Query" name:"DatasetId"`
	WhereInValuess *WhereInDimQueryWhereInValuesList `position:"Query" type:"Repeated" name:"WhereInValues"`
	Dimensionss    *WhereInDimQueryDimensionsList    `position:"Query" type:"Repeated" name:"Dimensions"`
}

func (r WhereInDimQueryRequest) Invoke(client *sdk.Client) (response *WhereInDimQueryResponse, err error) {
	req := struct {
		*requests.RpcRequest
		WhereInDimQueryRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("ARMS", "2016-11-25", "WhereInDimQuery", "", "")

	resp := struct {
		*responses.BaseResponse
		WhereInDimQueryResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.WhereInDimQueryResponse

	err = client.DoAction(&req, &resp)
	return
}

type WhereInDimQueryDimensions struct {
	Key   string `name:"Key"`
	Value string `name:"Value"`
	Type  string `name:"Type"`
}

type WhereInDimQueryResponse struct {
	Data string
}

type WhereInDimQueryMeasuresList []string

func (list *WhereInDimQueryMeasuresList) UnmarshalJSON(data []byte) error {
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

type WhereInDimQueryWhereInValuesList []string

func (list *WhereInDimQueryWhereInValuesList) UnmarshalJSON(data []byte) error {
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

type WhereInDimQueryDimensionsList []WhereInDimQueryDimensions

func (list *WhereInDimQueryDimensionsList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]WhereInDimQueryDimensions)
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
