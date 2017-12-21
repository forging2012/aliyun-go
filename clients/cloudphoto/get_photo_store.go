package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetPhotoStoreRequest struct {
	StoreName string `position:"Query" name:"StoreName"`
}

func (r GetPhotoStoreRequest) Invoke(client *sdk.Client) (response *GetPhotoStoreResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetPhotoStoreRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "GetPhotoStore", "cloudphoto", "")

	resp := struct {
		*responses.BaseResponse
		GetPhotoStoreResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetPhotoStoreResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetPhotoStoreResponse struct {
	Code       string
	Message    string
	RequestId  string
	Action     string
	PhotoStore GetPhotoStorePhotoStore
}

type GetPhotoStorePhotoStore struct {
	Id               int64
	Name             string
	Remark           string
	AutoCleanEnabled bool
	AutoCleanDays    int
	DefaultQuota     int64
	Ctime            int64
	Mtime            int64
	Buckets          GetPhotoStoreBucketList
}

type GetPhotoStoreBucket struct {
	Name   string
	Region string
	State  string
	Acl    string
}

type GetPhotoStoreBucketList []GetPhotoStoreBucket

func (list *GetPhotoStoreBucketList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetPhotoStoreBucket)
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
