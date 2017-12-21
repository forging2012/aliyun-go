package dds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeSecurityIpsRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeSecurityIpsRequest) Invoke(client *sdk.Client) (response *DescribeSecurityIpsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeSecurityIpsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Dds", "2015-12-01", "DescribeSecurityIps", "dds", "")

	resp := struct {
		*responses.BaseResponse
		DescribeSecurityIpsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeSecurityIpsResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeSecurityIpsResponse struct {
	RequestId        string
	SecurityIps      string
	SecurityIpGroups DescribeSecurityIpsSecurityIpGroupList
}

type DescribeSecurityIpsSecurityIpGroup struct {
	SecurityIpGroupName      string
	SecurityIpGroupAttribute string
	SecurityIpList           string
}

type DescribeSecurityIpsSecurityIpGroupList []DescribeSecurityIpsSecurityIpGroup

func (list *DescribeSecurityIpsSecurityIpGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSecurityIpsSecurityIpGroup)
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
