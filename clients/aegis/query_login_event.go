package aegis

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryLoginEventRequest struct {
	EndTime     string `position:"Query" name:"EndTime"`
	CurrentPage int    `position:"Query" name:"CurrentPage"`
	StartTime   string `position:"Query" name:"StartTime"`
	Uuid        string `position:"Query" name:"Uuid"`
	Status      int    `position:"Query" name:"Status"`
}

func (r QueryLoginEventRequest) Invoke(client *sdk.Client) (response *QueryLoginEventResponse, err error) {
	req := struct {
		*requests.RpcRequest
		QueryLoginEventRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("aegis", "2016-11-11", "QueryLoginEvent", "", "")

	resp := struct {
		*responses.BaseResponse
		QueryLoginEventResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.QueryLoginEventResponse

	err = client.DoAction(&req, &resp)
	return
}

type QueryLoginEventResponse struct {
	RequestId string
	Code      string
	Success   bool
	Message   string
	Data      QueryLoginEventData
}

type QueryLoginEventData struct {
	List     QueryLoginEventEntityList
	PageInfo QueryLoginEventPageInfo
}

type QueryLoginEventEntity struct {
	Uuid          string
	LoginTime     string
	LoginType     int
	LoginTypeName string
	BuyVersion    string
	LoginSourceIp string
	GroupId       int
	InstanceName  string
	InstanceId    string
	Ip            string
	Region        string
	Status        int
	StatusName    string
	Location      string
	UserName      string
}

type QueryLoginEventPageInfo struct {
	CurrentPage int
	PageSize    int
	TotalCount  int
	Count       int
}

type QueryLoginEventEntityList []QueryLoginEventEntity

func (list *QueryLoginEventEntityList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryLoginEventEntity)
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
