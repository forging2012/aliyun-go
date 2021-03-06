package itaas

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetRegisterBoxListRequest struct {
	requests.RpcRequest
	Sysfrom     string `position:"Query" name:"Sysfrom"`
	Operator    string `position:"Query" name:"Operator"`
	Clientappid string `position:"Query" name:"Clientappid"`
}

func (req *GetRegisterBoxListRequest) Invoke(client *sdk.Client) (resp *GetRegisterBoxListResponse, err error) {
	req.InitWithApiInfo("ITaaS", "2017-05-05", "GetRegisterBoxList", "", "")
	resp = &GetRegisterBoxListResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetRegisterBoxListResponse struct {
	responses.BaseResponse
	RequestId string
	ErrorCode int
	ErrorMsg  string
	Success   bool
	ErrorList GetRegisterBoxListErrorMessageList
	Data      GetRegisterBoxListData
}

type GetRegisterBoxListErrorMessage struct {
	ErrorMessage string
}

type GetRegisterBoxListData struct {
	ActivedNumber int
	BuyNumber     int
	BoxesList     GetRegisterBoxListBoxInfoList
}

type GetRegisterBoxListBoxInfo struct {
	CurVersion      string
	DrName          string
	DrSessionId     string
	DrStatus        string
	DrStatusTxt     string
	Ipaddress       string
	LastReportTimeL int64
	OnlineTimeL     int64
	Screencode      string
	SysVersion      string
}

type GetRegisterBoxListErrorMessageList []GetRegisterBoxListErrorMessage

func (list *GetRegisterBoxListErrorMessageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetRegisterBoxListErrorMessage)
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

type GetRegisterBoxListBoxInfoList []GetRegisterBoxListBoxInfo

func (list *GetRegisterBoxListBoxInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetRegisterBoxListBoxInfo)
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
