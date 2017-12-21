package yundun

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DdosLogRequest struct {
	JstOwnerId   int64  `position:"Query" name:"JstOwnerId"`
	PageNumber   int    `position:"Query" name:"PageNumber"`
	PageSize     int    `position:"Query" name:"PageSize"`
	InstanceId   string `position:"Query" name:"InstanceId"`
	InstanceType string `position:"Query" name:"InstanceType"`
}

func (r DdosLogRequest) Invoke(client *sdk.Client) (response *DdosLogResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DdosLogRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Yundun", "2015-04-16", "DdosLog", "", "")

	resp := struct {
		*responses.BaseResponse
		DdosLogResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DdosLogResponse

	err = client.DoAction(&req, &resp)
	return
}

type DdosLogResponse struct {
	RequestId    string
	AttackStatus int
	StartTime    string
	EndTime      string
	PageNumber   int
	PageSize     int
	TotalCount   int
	LogList      DdosLogDdosLogList
}

type DdosLogDdosLog struct {
	StartTime    string
	EndTime      string
	Reason       string
	Status       int
	Bps          int64
	Pps          int64
	Qps          int64
	AttackType   string
	AttackIpList string
	Type         int
}

type DdosLogDdosLogList []DdosLogDdosLog

func (list *DdosLogDdosLogList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DdosLogDdosLog)
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
