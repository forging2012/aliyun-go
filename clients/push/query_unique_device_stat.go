package push

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryUniqueDeviceStatRequest struct {
	Granularity string `position:"Query" name:"Granularity"`
	EndTime     string `position:"Query" name:"EndTime"`
	AppKey      int64  `position:"Query" name:"AppKey"`
	StartTime   string `position:"Query" name:"StartTime"`
}

func (r QueryUniqueDeviceStatRequest) Invoke(client *sdk.Client) (response *QueryUniqueDeviceStatResponse, err error) {
	req := struct {
		*requests.RpcRequest
		QueryUniqueDeviceStatRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Push", "2016-08-01", "QueryUniqueDeviceStat", "", "")

	resp := struct {
		*responses.BaseResponse
		QueryUniqueDeviceStatResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.QueryUniqueDeviceStatResponse

	err = client.DoAction(&req, &resp)
	return
}

type QueryUniqueDeviceStatResponse struct {
	RequestId      string
	AppDeviceStats QueryUniqueDeviceStatAppDeviceStatList
}

type QueryUniqueDeviceStatAppDeviceStat struct {
	Time  string
	Count int64
}

type QueryUniqueDeviceStatAppDeviceStatList []QueryUniqueDeviceStatAppDeviceStat

func (list *QueryUniqueDeviceStatAppDeviceStatList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryUniqueDeviceStatAppDeviceStat)
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
