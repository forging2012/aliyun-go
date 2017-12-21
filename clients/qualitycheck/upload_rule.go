package qualitycheck

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UploadRuleRequest struct {
	requests.RpcRequest
	JsonStr string `position:"Query" name:"JsonStr"`
}

func (req *UploadRuleRequest) Invoke(client *sdk.Client) (resp *UploadRuleResponse, err error) {
	req.InitWithApiInfo("Qualitycheck", "2016-08-01", "UploadRule", "", "")
	resp = &UploadRuleResponse{}
	err = client.DoAction(req, resp)
	return
}

type UploadRuleResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Code      string
	Message   string
	Data      UploadRuleDatumList
}

type UploadRuleDatumList []string

func (list *UploadRuleDatumList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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
