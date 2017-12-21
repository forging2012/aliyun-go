package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type EditPhotosRequest struct {
	LibraryId       string                 `position:"Query" name:"LibraryId"`
	ShareExpireTime int64                  `position:"Query" name:"ShareExpireTime"`
	PhotoIds        *EditPhotosPhotoIdList `position:"Query" type:"Repeated" name:"PhotoId"`
	StoreName       string                 `position:"Query" name:"StoreName"`
	Remark          string                 `position:"Query" name:"Remark"`
	Title           string                 `position:"Query" name:"Title"`
}

func (r EditPhotosRequest) Invoke(client *sdk.Client) (response *EditPhotosResponse, err error) {
	req := struct {
		*requests.RpcRequest
		EditPhotosRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "EditPhotos", "cloudphoto", "")

	resp := struct {
		*responses.BaseResponse
		EditPhotosResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.EditPhotosResponse

	err = client.DoAction(&req, &resp)
	return
}

type EditPhotosResponse struct {
	Code      string
	Message   string
	RequestId string
	Action    string
	Results   EditPhotosResultList
}

type EditPhotosResult struct {
	Id      int64
	Code    string
	Message string
}

type EditPhotosPhotoIdList []int64

func (list *EditPhotosPhotoIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]int64)
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

type EditPhotosResultList []EditPhotosResult

func (list *EditPhotosResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]EditPhotosResult)
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
