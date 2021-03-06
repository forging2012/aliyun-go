package ons

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type OnsEmpowerListRequest struct {
	requests.RpcRequest
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	EmpowerUser  string `position:"Query" name:"EmpowerUser"`
	Topic        string `position:"Query" name:"Topic"`
}

func (req *OnsEmpowerListRequest) Invoke(client *sdk.Client) (resp *OnsEmpowerListResponse, err error) {
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsEmpowerList", "", "")
	resp = &OnsEmpowerListResponse{}
	err = client.DoAction(req, resp)
	return
}

type OnsEmpowerListResponse struct {
	responses.BaseResponse
	RequestId string
	HelpUrl   string
	Data      OnsEmpowerListAuthOwnerInfoDoList
}

type OnsEmpowerListAuthOwnerInfoDo struct {
	Topic    string
	Owner    int64
	Relation int
}

type OnsEmpowerListAuthOwnerInfoDoList []OnsEmpowerListAuthOwnerInfoDo

func (list *OnsEmpowerListAuthOwnerInfoDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsEmpowerListAuthOwnerInfoDo)
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
