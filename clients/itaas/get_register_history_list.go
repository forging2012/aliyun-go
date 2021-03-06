package itaas

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetRegisterHistoryListRequest struct {
	requests.RpcRequest
	Sysfrom     string `position:"Query" name:"Sysfrom"`
	Operator    string `position:"Query" name:"Operator"`
	Clientappid string `position:"Query" name:"Clientappid"`
}

func (req *GetRegisterHistoryListRequest) Invoke(client *sdk.Client) (resp *GetRegisterHistoryListResponse, err error) {
	req.InitWithApiInfo("ITaaS", "2017-05-05", "GetRegisterHistoryList", "", "")
	resp = &GetRegisterHistoryListResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetRegisterHistoryListResponse struct {
	responses.BaseResponse
	RequestId string
	ErrorCode int
	ErrorMsg  string
	Success   bool
	Data      GetRegisterHistoryListRegisterHistoryInfoList
	ErrorList GetRegisterHistoryListErrorMessageList
}

type GetRegisterHistoryListRegisterHistoryInfo struct {
	CreateTimeL  int64
	DrIp         string
	DrMac        string
	DrName       string
	Eventinfo    string
	Eventtype    int
	EventtypeTxt string
	Memo         string
	Screencode   string
}

type GetRegisterHistoryListErrorMessage struct {
	ErrorMessage string
}

type GetRegisterHistoryListRegisterHistoryInfoList []GetRegisterHistoryListRegisterHistoryInfo

func (list *GetRegisterHistoryListRegisterHistoryInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetRegisterHistoryListRegisterHistoryInfo)
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

type GetRegisterHistoryListErrorMessageList []GetRegisterHistoryListErrorMessage

func (list *GetRegisterHistoryListErrorMessageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetRegisterHistoryListErrorMessage)
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
