package vpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeEipMonitorDataRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Period               int    `position:"Query" name:"Period"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	EndTime              string `position:"Query" name:"EndTime"`
	AllocationId         string `position:"Query" name:"AllocationId"`
	StartTime            string `position:"Query" name:"StartTime"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeEipMonitorDataRequest) Invoke(client *sdk.Client) (resp *DescribeEipMonitorDataResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "DescribeEipMonitorData", "vpc", "")
	resp = &DescribeEipMonitorDataResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeEipMonitorDataResponse struct {
	responses.BaseResponse
	RequestId       string
	EipMonitorDatas DescribeEipMonitorDataEipMonitorDataList
}

type DescribeEipMonitorDataEipMonitorData struct {
	EipRX        int
	EipTX        int
	EipFlow      int
	EipBandwidth int
	EipPackets   int
	TimeStamp    string
}

type DescribeEipMonitorDataEipMonitorDataList []DescribeEipMonitorDataEipMonitorData

func (list *DescribeEipMonitorDataEipMonitorDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeEipMonitorDataEipMonitorData)
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
