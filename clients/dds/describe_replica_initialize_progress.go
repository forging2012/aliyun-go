package dds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeReplicaInitializeProgressRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	ReplicaId            string `position:"Query" name:"ReplicaId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeReplicaInitializeProgressRequest) Invoke(client *sdk.Client) (response *DescribeReplicaInitializeProgressResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeReplicaInitializeProgressRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Dds", "2015-12-01", "DescribeReplicaInitializeProgress", "dds", "")

	resp := struct {
		*responses.BaseResponse
		DescribeReplicaInitializeProgressResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeReplicaInitializeProgressResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeReplicaInitializeProgressResponse struct {
	RequestId string
	Items     DescribeReplicaInitializeProgressItemsItemList
}

type DescribeReplicaInitializeProgressItemsItem struct {
	ReplicaId   string
	Status      string
	Progress    string
	FinishTime  string
	CurrentStep string
}

type DescribeReplicaInitializeProgressItemsItemList []DescribeReplicaInitializeProgressItemsItem

func (list *DescribeReplicaInitializeProgressItemsItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeReplicaInitializeProgressItemsItem)
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
