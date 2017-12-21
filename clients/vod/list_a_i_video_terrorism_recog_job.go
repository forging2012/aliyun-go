package vod

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListAIVideoTerrorismRecogJobRequest struct {
	ResourceOwnerId             string `position:"Query" name:"ResourceOwnerId"`
	AIVideoTerrorismRecogJobIds string `position:"Query" name:"AIVideoTerrorismRecogJobIds"`
	ResourceOwnerAccount        string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                string `position:"Query" name:"OwnerAccount"`
	OwnerId                     string `position:"Query" name:"OwnerId"`
}

func (r ListAIVideoTerrorismRecogJobRequest) Invoke(client *sdk.Client) (response *ListAIVideoTerrorismRecogJobResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ListAIVideoTerrorismRecogJobRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("vod", "2017-03-21", "ListAIVideoTerrorismRecogJob", "", "")

	resp := struct {
		*responses.BaseResponse
		ListAIVideoTerrorismRecogJobResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ListAIVideoTerrorismRecogJobResponse

	err = client.DoAction(&req, &resp)
	return
}

type ListAIVideoTerrorismRecogJobResponse struct {
	RequestId                    string
	AIVideoTerrorismRecogJobList ListAIVideoTerrorismRecogJobAIVideoTerrorismRecogJobList
	NonExistTerrorismRecogJobIds ListAIVideoTerrorismRecogJobNonExistTerrorismRecogJobIdList
}

type ListAIVideoTerrorismRecogJobAIVideoTerrorismRecogJob struct {
	JobId        string
	MediaId      string
	Status       string
	Code         string
	Message      string
	CreationTime string
	Data         string
}

type ListAIVideoTerrorismRecogJobAIVideoTerrorismRecogJobList []ListAIVideoTerrorismRecogJobAIVideoTerrorismRecogJob

func (list *ListAIVideoTerrorismRecogJobAIVideoTerrorismRecogJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAIVideoTerrorismRecogJobAIVideoTerrorismRecogJob)
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

type ListAIVideoTerrorismRecogJobNonExistTerrorismRecogJobIdList []string

func (list *ListAIVideoTerrorismRecogJobNonExistTerrorismRecogJobIdList) UnmarshalJSON(data []byte) error {
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
