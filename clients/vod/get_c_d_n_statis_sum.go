package vod

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetCDNStatisSumRequest struct {
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	StartStatisTime      string `position:"Query" name:"StartStatisTime"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Level                string `position:"Query" name:"Level"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	EndStatisTime        string `position:"Query" name:"EndStatisTime"`
}

func (r GetCDNStatisSumRequest) Invoke(client *sdk.Client) (response *GetCDNStatisSumResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetCDNStatisSumRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("vod", "2017-03-21", "GetCDNStatisSum", "", "")

	resp := struct {
		*responses.BaseResponse
		GetCDNStatisSumResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetCDNStatisSumResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetCDNStatisSumResponse struct {
	RequestId        string
	SumFlowDataValue int64
	MaxBpsDataValue  int64
	CDNStatisList    GetCDNStatisSumCDNMetricList
}

type GetCDNStatisSumCDNMetric struct {
	StatTime              string
	FlowDataDomesticValue int64
	FlowDataOverseasValue int64
	FlowDataValue         int64
	BpsDataDomesticValue  int64
	BpsDataOverseasValue  int64
	BpsDataValue          int64
}

type GetCDNStatisSumCDNMetricList []GetCDNStatisSumCDNMetric

func (list *GetCDNStatisSumCDNMetricList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetCDNStatisSumCDNMetric)
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
