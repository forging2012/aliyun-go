package itaas

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpdateIPSegmentRequest struct {
	Sysfrom     string `position:"Query" name:"Sysfrom"`
	Operator    string `position:"Query" name:"Operator"`
	Clientappid string `position:"Query" name:"Clientappid"`
	Uuid        string `position:"Query" name:"Uuid"`
	Ipsegment   string `position:"Query" name:"Ipsegment"`
	Memo        string `position:"Query" name:"Memo"`
}

func (r UpdateIPSegmentRequest) Invoke(client *sdk.Client) (response *UpdateIPSegmentResponse, err error) {
	req := struct {
		*requests.RpcRequest
		UpdateIPSegmentRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("ITaaS", "2017-05-05", "UpdateIPSegment", "", "")

	resp := struct {
		*responses.BaseResponse
		UpdateIPSegmentResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.UpdateIPSegmentResponse

	err = client.DoAction(&req, &resp)
	return
}

type UpdateIPSegmentResponse struct {
	RequestId string
	ErrorCode int
	ErrorMsg  string
	Success   bool
	ErrorList UpdateIPSegmentErrorMessageList
}

type UpdateIPSegmentErrorMessage struct {
	ErrorMessage string
}

type UpdateIPSegmentErrorMessageList []UpdateIPSegmentErrorMessage

func (list *UpdateIPSegmentErrorMessageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]UpdateIPSegmentErrorMessage)
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
