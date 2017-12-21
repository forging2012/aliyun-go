package aegis

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetStatisticsByUuidRequest struct {
	Uuid string `position:"Query" name:"Uuid"`
}

func (r GetStatisticsByUuidRequest) Invoke(client *sdk.Client) (response *GetStatisticsByUuidResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetStatisticsByUuidRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("aegis", "2016-11-11", "GetStatisticsByUuid", "", "")

	resp := struct {
		*responses.BaseResponse
		GetStatisticsByUuidResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetStatisticsByUuidResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetStatisticsByUuidResponse struct {
	RequestId string
	Code      string
	Success   bool
	Message   string
	Data      GetStatisticsByUuidEntityList
}

type GetStatisticsByUuidEntity struct {
	Uuid    string
	Account int
	Health  int
	Patch   int
	Trojan  int
	Online  bool
}

type GetStatisticsByUuidEntityList []GetStatisticsByUuidEntity

func (list *GetStatisticsByUuidEntityList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetStatisticsByUuidEntity)
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
