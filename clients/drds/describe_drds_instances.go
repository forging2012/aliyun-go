package drds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDrdsInstancesRequest struct {
	requests.RpcRequest
	Type string `position:"Query" name:"Type"`
}

func (req *DescribeDrdsInstancesRequest) Invoke(client *sdk.Client) (resp *DescribeDrdsInstancesResponse, err error) {
	req.InitWithApiInfo("Drds", "2017-10-16", "DescribeDrdsInstances", "", "")
	resp = &DescribeDrdsInstancesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDrdsInstancesResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Data      DescribeDrdsInstancesInstanceList
}

type DescribeDrdsInstancesInstance struct {
	DrdsInstanceId string
	Type           string
	RegionId       string
	ZoneId         string
	Description    string
	NetworkType    string
	Status         string
	CreateTime     int64
	Version        int64
	Vips           DescribeDrdsInstancesVipList
}

type DescribeDrdsInstancesVip struct {
	IP   string
	Port string
	Type string
}

type DescribeDrdsInstancesInstanceList []DescribeDrdsInstancesInstance

func (list *DescribeDrdsInstancesInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDrdsInstancesInstance)
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

type DescribeDrdsInstancesVipList []DescribeDrdsInstancesVip

func (list *DescribeDrdsInstancesVipList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDrdsInstancesVip)
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
