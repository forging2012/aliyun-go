package ehpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListSoftwaresRequest struct {
	requests.RpcRequest
	EhpcVersion string `position:"Query" name:"EhpcVersion"`
}

func (req *ListSoftwaresRequest) Invoke(client *sdk.Client) (resp *ListSoftwaresResponse, err error) {
	req.InitWithApiInfo("EHPC", "2017-07-14", "ListSoftwares", "ehs", "")
	resp = &ListSoftwaresResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListSoftwaresResponse struct {
	responses.BaseResponse
	RequestId string
	Softwares ListSoftwaresSoftwareInfoList
}

type ListSoftwaresSoftwareInfo struct {
	EhpcVersion      string
	OsTag            string
	SchedulerType    string
	SchedulerVersion string
	AccountType      string
	AccountVersion   string
	Applications     ListSoftwaresApplicationInfoList
}

type ListSoftwaresApplicationInfo struct {
	Tag      string
	Name     string
	Version  string
	Required bool
}

type ListSoftwaresSoftwareInfoList []ListSoftwaresSoftwareInfo

func (list *ListSoftwaresSoftwareInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListSoftwaresSoftwareInfo)
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

type ListSoftwaresApplicationInfoList []ListSoftwaresApplicationInfo

func (list *ListSoftwaresApplicationInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListSoftwaresApplicationInfo)
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
