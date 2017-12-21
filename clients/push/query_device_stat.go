package push

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryDeviceStatRequest struct {
	EndTime    string `position:"Query" name:"EndTime"`
	AppKey     int64  `position:"Query" name:"AppKey"`
	StartTime  string `position:"Query" name:"StartTime"`
	DeviceType string `position:"Query" name:"DeviceType"`
	QueryType  string `position:"Query" name:"QueryType"`
}

func (r QueryDeviceStatRequest) Invoke(client *sdk.Client) (response *QueryDeviceStatResponse, err error) {
	req := struct {
		*requests.RpcRequest
		QueryDeviceStatRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Push", "2016-08-01", "QueryDeviceStat", "", "")

	resp := struct {
		*responses.BaseResponse
		QueryDeviceStatResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.QueryDeviceStatResponse

	err = client.DoAction(&req, &resp)
	return
}

type QueryDeviceStatResponse struct {
	RequestId      string
	AppDeviceStats QueryDeviceStatAppDeviceStatList
}

type QueryDeviceStatAppDeviceStat struct {
	Time       string
	Count      int64
	DeviceType string
}

type QueryDeviceStatAppDeviceStatList []QueryDeviceStatAppDeviceStat

func (list *QueryDeviceStatAppDeviceStatList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryDeviceStatAppDeviceStat)
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