package drds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeRegionsRequest struct {
}

func (r DescribeRegionsRequest) Invoke(client *sdk.Client) (response *DescribeRegionsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeRegionsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Drds", "2017-10-16", "DescribeRegions", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeRegionsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeRegionsResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeRegionsResponse struct {
	RequestId   string
	Success     bool
	DrdsRegions DescribeRegionsDrdsRegionList
}

type DescribeRegionsDrdsRegion struct {
	RegionId           string
	RegionName         string
	ZoneId             string
	ZoneName           string
	InstanceSeriesList DescribeRegionsInstanceSeriesList
}

type DescribeRegionsInstanceSeries struct {
	SeriesName string
	SpecList   DescribeRegionsSpecListList
}

type DescribeRegionsDrdsRegionList []DescribeRegionsDrdsRegion

func (list *DescribeRegionsDrdsRegionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRegionsDrdsRegion)
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

type DescribeRegionsInstanceSeriesList []DescribeRegionsInstanceSeries

func (list *DescribeRegionsInstanceSeriesList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRegionsInstanceSeries)
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

type DescribeRegionsSpecListList []string

func (list *DescribeRegionsSpecListList) UnmarshalJSON(data []byte) error {
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
