package live

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type StartCasterRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	CasterId      string `position:"Query" name:"CasterId"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	Version       string `position:"Query" name:"Version"`
}

func (r StartCasterRequest) Invoke(client *sdk.Client) (response *StartCasterResponse, err error) {
	req := struct {
		*requests.RpcRequest
		StartCasterRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("live", "2016-11-01", "StartCaster", "", "")

	resp := struct {
		*responses.BaseResponse
		StartCasterResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.StartCasterResponse

	err = client.DoAction(&req, &resp)
	return
}

type StartCasterResponse struct {
	RequestId     string
	PvwSceneInfos StartCasterSceneInfoList
	PgmSceneInfos StartCasterSceneInfo1List
}

type StartCasterSceneInfo struct {
	SceneId   string
	StreamUrl string
}

type StartCasterSceneInfo1 struct {
	SceneId     string
	StreamUrl   string
	StreamInfos StartCasterStreamInfoList
}

type StartCasterStreamInfo struct {
	TranscodeConfig string
	VideoFormat     string
	OutputStreamUrl string
}

type StartCasterSceneInfoList []StartCasterSceneInfo

func (list *StartCasterSceneInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]StartCasterSceneInfo)
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

type StartCasterSceneInfo1List []StartCasterSceneInfo1

func (list *StartCasterSceneInfo1List) UnmarshalJSON(data []byte) error {
	m := make(map[string][]StartCasterSceneInfo1)
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

type StartCasterStreamInfoList []StartCasterStreamInfo

func (list *StartCasterStreamInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]StartCasterStreamInfo)
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
