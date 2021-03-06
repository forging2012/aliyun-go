package ram

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListUsersRequest struct {
	requests.RpcRequest
	Marker   string `position:"Query" name:"Marker"`
	MaxItems int    `position:"Query" name:"MaxItems"`
}

func (req *ListUsersRequest) Invoke(client *sdk.Client) (resp *ListUsersResponse, err error) {
	req.InitWithApiInfo("Ram", "2015-05-01", "ListUsers", "", "")
	resp = &ListUsersResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListUsersResponse struct {
	responses.BaseResponse
	RequestId   string
	IsTruncated bool
	Marker      string
	Users       ListUsersUserList
}

type ListUsersUser struct {
	UserId      string
	UserName    string
	DisplayName string
	MobilePhone string
	Email       string
	Comments    string
	CreateDate  string
	UpdateDate  string
}

type ListUsersUserList []ListUsersUser

func (list *ListUsersUserList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListUsersUser)
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
