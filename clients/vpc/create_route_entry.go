package vpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateRouteEntryRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64                            `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string                           `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string                           `position:"Query" name:"ClientToken"`
	DestinationCidrBlock string                           `position:"Query" name:"DestinationCidrBlock"`
	OwnerAccount         string                           `position:"Query" name:"OwnerAccount"`
	NextHopId            string                           `position:"Query" name:"NextHopId"`
	OwnerId              int64                            `position:"Query" name:"OwnerId"`
	NextHopType          string                           `position:"Query" name:"NextHopType"`
	NextHopLists         *CreateRouteEntryNextHopListList `position:"Query" type:"Repeated" name:"NextHopList"`
	RouteTableId         string                           `position:"Query" name:"RouteTableId"`
}

func (req *CreateRouteEntryRequest) Invoke(client *sdk.Client) (resp *CreateRouteEntryResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "CreateRouteEntry", "vpc", "")
	resp = &CreateRouteEntryResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateRouteEntryNextHopList struct {
	NextHopType string `name:"NextHopType"`
	NextHopId   string `name:"NextHopId"`
	Weight      int    `name:"Weight"`
}

type CreateRouteEntryResponse struct {
	responses.BaseResponse
	RequestId string
}

type CreateRouteEntryNextHopListList []CreateRouteEntryNextHopList

func (list *CreateRouteEntryNextHopListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateRouteEntryNextHopList)
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
