package iot

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type BatchGetDeviceStateRequest struct {
	requests.RpcRequest
	DeviceNames *BatchGetDeviceStateDeviceNameList `position:"Query" type:"Repeated" name:"DeviceName"`
	ProductKey  string                             `position:"Query" name:"ProductKey"`
}

func (req *BatchGetDeviceStateRequest) Invoke(client *sdk.Client) (resp *BatchGetDeviceStateResponse, err error) {
	req.InitWithApiInfo("Iot", "2017-04-20", "BatchGetDeviceState", "", "")
	resp = &BatchGetDeviceStateResponse{}
	err = client.DoAction(req, resp)
	return
}

type BatchGetDeviceStateResponse struct {
	responses.BaseResponse
	RequestId        string
	Success          bool
	ErrorMessage     string
	DeviceStatusList BatchGetDeviceStateDeviceStatusList
}

type BatchGetDeviceStateDeviceStatus struct {
	DeviceId       string
	DeviceName     string
	Status         string
	AsAddress      string
	LastOnlineTime string
}

type BatchGetDeviceStateDeviceNameList []string

func (list *BatchGetDeviceStateDeviceNameList) UnmarshalJSON(data []byte) error {
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

type BatchGetDeviceStateDeviceStatusList []BatchGetDeviceStateDeviceStatus

func (list *BatchGetDeviceStateDeviceStatusList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]BatchGetDeviceStateDeviceStatus)
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
