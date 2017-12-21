package live

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeLiveStreamBitRateDataRequest struct {
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	StreamName    string `position:"Query" name:"StreamName"`
}

func (r DescribeLiveStreamBitRateDataRequest) Invoke(client *sdk.Client) (response *DescribeLiveStreamBitRateDataResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeLiveStreamBitRateDataRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("live", "2016-11-01", "DescribeLiveStreamBitRateData", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeLiveStreamBitRateDataResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeLiveStreamBitRateDataResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeLiveStreamBitRateDataResponse struct {
	RequestId                string
	FrameRateAndBitRateInfos DescribeLiveStreamBitRateDataFrameRateAndBitRateInfoList
}

type DescribeLiveStreamBitRateDataFrameRateAndBitRateInfo struct {
	StreamUrl      string
	VideoFrameRate float32
	AudioFrameRate float32
	BitRate        float32
	Time           string
}

type DescribeLiveStreamBitRateDataFrameRateAndBitRateInfoList []DescribeLiveStreamBitRateDataFrameRateAndBitRateInfo

func (list *DescribeLiveStreamBitRateDataFrameRateAndBitRateInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamBitRateDataFrameRateAndBitRateInfo)
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
