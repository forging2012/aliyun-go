package csb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type FindOrderableListRequest struct {
	requests.RpcRequest
	ProjectName string `position:"Query" name:"ProjectName"`
	CsbId       int64  `position:"Query" name:"CsbId"`
	Alias       string `position:"Query" name:"Alias"`
	ServiceName string `position:"Query" name:"ServiceName"`
	PageNum     int    `position:"Query" name:"PageNum"`
}

func (req *FindOrderableListRequest) Invoke(client *sdk.Client) (resp *FindOrderableListResponse, err error) {
	req.InitWithApiInfo("CSB", "2017-11-18", "FindOrderableList", "CSB", "")
	resp = &FindOrderableListResponse{}
	err = client.DoAction(req, resp)
	return
}

type FindOrderableListResponse struct {
	responses.BaseResponse
	Code      int
	Message   string
	RequestId string
	Data      FindOrderableListData
}

type FindOrderableListData struct {
	CurrentPage int
	PageNumber  int
	ServiceList FindOrderableListServiceList
}

type FindOrderableListService struct {
	Alias          string
	AllVisiable    bool
	ApproveUserId  string
	CasTargets     string
	CreateTime     int64
	CsbId          int64
	Id             int64
	InterfaceName  string
	ModifiedTime   int64
	OwnerId        string
	PrincipalName  string
	ProjectId      string
	ProjectName    string
	Scope          string
	ServiceName    string
	ServiceVersion string
	SkipAuth       bool
	StatisticName  string
	Status         int
	UserId         string
}

type FindOrderableListServiceList []FindOrderableListService

func (list *FindOrderableListServiceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]FindOrderableListService)
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
