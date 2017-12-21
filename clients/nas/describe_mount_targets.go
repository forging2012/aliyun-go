package nas

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeMountTargetsRequest struct {
	requests.RpcRequest
	MountTargetDomain string `position:"Query" name:"MountTargetDomain"`
	PageSize          int    `position:"Query" name:"PageSize"`
	PageNumber        int    `position:"Query" name:"PageNumber"`
	FileSystemId      string `position:"Query" name:"FileSystemId"`
}

func (req *DescribeMountTargetsRequest) Invoke(client *sdk.Client) (resp *DescribeMountTargetsResponse, err error) {
	req.InitWithApiInfo("NAS", "2017-06-26", "DescribeMountTargets", "nas", "")
	resp = &DescribeMountTargetsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeMountTargetsResponse struct {
	responses.BaseResponse
	RequestId    string
	TotalCount   int
	PageSize     int
	PageNumber   int
	MountTargets DescribeMountTargetsMountTargetList
}

type DescribeMountTargetsMountTarget struct {
	MountTargetDomain string
	NetworkType       string
	VpcId             string
	VswId             string
	AccessGroup       string
	Status            string
}

type DescribeMountTargetsMountTargetList []DescribeMountTargetsMountTarget

func (list *DescribeMountTargetsMountTargetList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeMountTargetsMountTarget)
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
