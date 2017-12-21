package slb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateVServerGroupRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	LoadBalancerId       string `position:"Query" name:"LoadBalancerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	BackendServers       string `position:"Query" name:"BackendServers"`
	Tags                 string `position:"Query" name:"Tags"`
	VServerGroupName     string `position:"Query" name:"VServerGroupName"`
}

func (r CreateVServerGroupRequest) Invoke(client *sdk.Client) (response *CreateVServerGroupResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateVServerGroupRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Slb", "2014-05-15", "CreateVServerGroup", "slb", "")

	resp := struct {
		*responses.BaseResponse
		CreateVServerGroupResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateVServerGroupResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateVServerGroupResponse struct {
	RequestId      string
	VServerGroupId string
	BackendServers CreateVServerGroupBackendServerList
}

type CreateVServerGroupBackendServer struct {
	ServerId string
	Port     int
	Weight   int
}

type CreateVServerGroupBackendServerList []CreateVServerGroupBackendServer

func (list *CreateVServerGroupBackendServerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateVServerGroupBackendServer)
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
