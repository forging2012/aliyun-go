package vod

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListAIVideoCategoryJobRequest struct {
	requests.RpcRequest
	ResourceOwnerId       string `position:"Query" name:"ResourceOwnerId"`
	AIVideoCategoryJobIds string `position:"Query" name:"AIVideoCategoryJobIds"`
	ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount          string `position:"Query" name:"OwnerAccount"`
	OwnerId               string `position:"Query" name:"OwnerId"`
}

func (req *ListAIVideoCategoryJobRequest) Invoke(client *sdk.Client) (resp *ListAIVideoCategoryJobResponse, err error) {
	req.InitWithApiInfo("vod", "2017-03-21", "ListAIVideoCategoryJob", "", "")
	resp = &ListAIVideoCategoryJobResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListAIVideoCategoryJobResponse struct {
	responses.BaseResponse
	RequestId                     string
	AIVideoCategoryJobList        ListAIVideoCategoryJobAIVideoCategoryJobList
	NonExistAIVideoCategoryJobIds ListAIVideoCategoryJobNonExistAIVideoCategoryJobIdList
}

type ListAIVideoCategoryJobAIVideoCategoryJob struct {
	JobId        string
	MediaId      string
	Status       string
	Code         string
	Message      string
	CreationTime string
	Data         string
}

type ListAIVideoCategoryJobAIVideoCategoryJobList []ListAIVideoCategoryJobAIVideoCategoryJob

func (list *ListAIVideoCategoryJobAIVideoCategoryJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAIVideoCategoryJobAIVideoCategoryJob)
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

type ListAIVideoCategoryJobNonExistAIVideoCategoryJobIdList []string

func (list *ListAIVideoCategoryJobNonExistAIVideoCategoryJobIdList) UnmarshalJSON(data []byte) error {
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
