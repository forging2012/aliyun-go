package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DetachInstanceRamRoleRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	InstanceIds          string `position:"Query" name:"InstanceIds"`
	RamRoleName          string `position:"Query" name:"RamRoleName"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DetachInstanceRamRoleRequest) Invoke(client *sdk.Client) (resp *DetachInstanceRamRoleResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DetachInstanceRamRole", "ecs", "")
	resp = &DetachInstanceRamRoleResponse{}
	err = client.DoAction(req, resp)
	return
}

type DetachInstanceRamRoleResponse struct {
	responses.BaseResponse
	RequestId                    string
	TotalCount                   int
	FailCount                    int
	RamRoleName                  string
	DetachInstanceRamRoleResults DetachInstanceRamRoleDetachInstanceRamRoleResultList
}

type DetachInstanceRamRoleDetachInstanceRamRoleResult struct {
	InstanceId          string
	Success             bool
	Code                string
	Message             string
	InstanceRamRoleSets DetachInstanceRamRoleInstanceRamRoleSetList
}

type DetachInstanceRamRoleInstanceRamRoleSet struct {
	InstanceId  string
	RamRoleName string
}

type DetachInstanceRamRoleDetachInstanceRamRoleResultList []DetachInstanceRamRoleDetachInstanceRamRoleResult

func (list *DetachInstanceRamRoleDetachInstanceRamRoleResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DetachInstanceRamRoleDetachInstanceRamRoleResult)
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

type DetachInstanceRamRoleInstanceRamRoleSetList []DetachInstanceRamRoleInstanceRamRoleSet

func (list *DetachInstanceRamRoleInstanceRamRoleSetList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DetachInstanceRamRoleInstanceRamRoleSet)
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
