package r_kvstore

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ReleaseReplicaRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	ReplicaId            string `position:"Query" name:"ReplicaId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *ReleaseReplicaRequest) Invoke(client *sdk.Client) (resp *ReleaseReplicaResponse, err error) {
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "ReleaseReplica", "redisa", "")
	resp = &ReleaseReplicaResponse{}
	err = client.DoAction(req, resp)
	return
}

type ReleaseReplicaResponse struct {
	responses.BaseResponse
	RequestId string
}
