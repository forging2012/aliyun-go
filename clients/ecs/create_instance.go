package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateInstanceRequest struct {
	requests.RpcRequest
	Tag4Value                   string                      `position:"Query" name:"Tag.4.Value"`
	ResourceOwnerId             int64                       `position:"Query" name:"ResourceOwnerId"`
	Tag2Key                     string                      `position:"Query" name:"Tag.2.Key"`
	HpcClusterId                string                      `position:"Query" name:"HpcClusterId"`
	Tag3Key                     string                      `position:"Query" name:"Tag.3.Key"`
	SecurityEnhancementStrategy string                      `position:"Query" name:"SecurityEnhancementStrategy"`
	KeyPairName                 string                      `position:"Query" name:"KeyPairName"`
	SpotPriceLimit              float32                     `position:"Query" name:"SpotPriceLimit"`
	Tag1Value                   string                      `position:"Query" name:"Tag.1.Value"`
	ResourceGroupId             string                      `position:"Query" name:"ResourceGroupId"`
	HostName                    string                      `position:"Query" name:"HostName"`
	Password                    string                      `position:"Query" name:"Password"`
	AutoRenewPeriod             int                         `position:"Query" name:"AutoRenewPeriod"`
	NodeControllerId            string                      `position:"Query" name:"NodeControllerId"`
	Period                      int                         `position:"Query" name:"Period"`
	DryRun                      string                      `position:"Query" name:"DryRun"`
	Tag5Key                     string                      `position:"Query" name:"Tag.5.Key"`
	OwnerId                     int64                       `position:"Query" name:"OwnerId"`
	VSwitchId                   string                      `position:"Query" name:"VSwitchId"`
	PrivateIpAddress            string                      `position:"Query" name:"PrivateIpAddress"`
	SpotStrategy                string                      `position:"Query" name:"SpotStrategy"`
	PeriodUnit                  string                      `position:"Query" name:"PeriodUnit"`
	InstanceName                string                      `position:"Query" name:"InstanceName"`
	AutoRenew                   string                      `position:"Query" name:"AutoRenew"`
	InternetChargeType          string                      `position:"Query" name:"InternetChargeType"`
	ZoneId                      string                      `position:"Query" name:"ZoneId"`
	Tag4Key                     string                      `position:"Query" name:"Tag.4.Key"`
	InternetMaxBandwidthIn      int                         `position:"Query" name:"InternetMaxBandwidthIn"`
	UseAdditionalService        string                      `position:"Query" name:"UseAdditionalService"`
	ImageId                     string                      `position:"Query" name:"ImageId"`
	ClientToken                 string                      `position:"Query" name:"ClientToken"`
	VlanId                      string                      `position:"Query" name:"VlanId"`
	IoOptimized                 string                      `position:"Query" name:"IoOptimized"`
	SecurityGroupId             string                      `position:"Query" name:"SecurityGroupId"`
	InternetMaxBandwidthOut     int                         `position:"Query" name:"InternetMaxBandwidthOut"`
	Description                 string                      `position:"Query" name:"Description"`
	SystemDiskCategory          string                      `position:"Query" name:"SystemDiskCategory"`
	UserData                    string                      `position:"Query" name:"UserData"`
	InstanceType                string                      `position:"Query" name:"InstanceType"`
	InstanceChargeType          string                      `position:"Query" name:"InstanceChargeType"`
	Tag3Value                   string                      `position:"Query" name:"Tag.3.Value"`
	DeploymentSetId             string                      `position:"Query" name:"DeploymentSetId"`
	InnerIpAddress              string                      `position:"Query" name:"InnerIpAddress"`
	ResourceOwnerAccount        string                      `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                string                      `position:"Query" name:"OwnerAccount"`
	SystemDiskDiskName          string                      `position:"Query" name:"SystemDiskDiskName"`
	RamRoleName                 string                      `position:"Query" name:"RamRoleName"`
	ClusterId                   string                      `position:"Query" name:"ClusterId"`
	DataDisks                   *CreateInstanceDataDiskList `position:"Query" type:"Repeated" name:"DataDisk"`
	Tag5Value                   string                      `position:"Query" name:"Tag.5.Value"`
	Tag1Key                     string                      `position:"Query" name:"Tag.1.Key"`
	SystemDiskSize              int                         `position:"Query" name:"SystemDiskSize"`
	Tag2Value                   string                      `position:"Query" name:"Tag.2.Value"`
	SystemDiskDescription       string                      `position:"Query" name:"SystemDiskDescription"`
}

func (req *CreateInstanceRequest) Invoke(client *sdk.Client) (resp *CreateInstanceResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "CreateInstance", "ecs", "")
	resp = &CreateInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateInstanceDataDisk struct {
	Size               int    `name:"Size"`
	SnapshotId         string `name:"SnapshotId"`
	Category           string `name:"Category"`
	DiskName           string `name:"DiskName"`
	Description        string `name:"Description"`
	Device             string `name:"Device"`
	DeleteWithInstance string `name:"DeleteWithInstance"`
	Encrypted          string `name:"Encrypted"`
}

type CreateInstanceResponse struct {
	responses.BaseResponse
	RequestId  string
	InstanceId string
}

type CreateInstanceDataDiskList []CreateInstanceDataDisk

func (list *CreateInstanceDataDiskList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateInstanceDataDisk)
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
