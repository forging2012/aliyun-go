package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDomainCCAttackInfoRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeDomainCCAttackInfoRequest) Invoke(client *sdk.Client) (resp *DescribeDomainCCAttackInfoResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeDomainCCAttackInfo", "", "")
	resp = &DescribeDomainCCAttackInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDomainCCAttackInfoResponse struct {
	responses.BaseResponse
	RequestId           string
	DomainName          string
	StartTime           string
	EndTime             string
	AttackIpDataList    DescribeDomainCCAttackInfoAttackIpDatasList
	AttackedUrlDataList DescribeDomainCCAttackInfoAttackedUrlDatasList
}

type DescribeDomainCCAttackInfoAttackIpDatas struct {
	Ip          string
	AttackCount string
	Result      string
}

type DescribeDomainCCAttackInfoAttackedUrlDatas struct {
	Url         string
	AttackCount string
	Result      string
}

type DescribeDomainCCAttackInfoAttackIpDatasList []DescribeDomainCCAttackInfoAttackIpDatas

func (list *DescribeDomainCCAttackInfoAttackIpDatasList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainCCAttackInfoAttackIpDatas)
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

type DescribeDomainCCAttackInfoAttackedUrlDatasList []DescribeDomainCCAttackInfoAttackedUrlDatas

func (list *DescribeDomainCCAttackInfoAttackedUrlDatasList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainCCAttackInfoAttackedUrlDatas)
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
