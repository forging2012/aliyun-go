package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RemoveFacePhotosRequest struct {
	LibraryId string                       `position:"Query" name:"LibraryId"`
	PhotoIds  *RemoveFacePhotosPhotoIdList `position:"Query" type:"Repeated" name:"PhotoId"`
	StoreName string                       `position:"Query" name:"StoreName"`
	FaceId    int64                        `position:"Query" name:"FaceId"`
}

func (r RemoveFacePhotosRequest) Invoke(client *sdk.Client) (response *RemoveFacePhotosResponse, err error) {
	req := struct {
		*requests.RpcRequest
		RemoveFacePhotosRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "RemoveFacePhotos", "cloudphoto", "")

	resp := struct {
		*responses.BaseResponse
		RemoveFacePhotosResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.RemoveFacePhotosResponse

	err = client.DoAction(&req, &resp)
	return
}

type RemoveFacePhotosResponse struct {
	Code      string
	Message   string
	RequestId string
	Action    string
	Results   RemoveFacePhotosResultList
}

type RemoveFacePhotosResult struct {
	Id      int64
	Code    string
	Message string
}

type RemoveFacePhotosPhotoIdList []int64

func (list *RemoveFacePhotosPhotoIdList) UnmarshalJSON(data []byte) error {
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

type RemoveFacePhotosResultList []RemoveFacePhotosResult

func (list *RemoveFacePhotosResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]RemoveFacePhotosResult)
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
