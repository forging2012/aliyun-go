package slb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateLoadBalancerUDPListenerRequest struct {
	Access_key_id             string `position:"Query" name:"Access_key_id"`
	HealthCheckConnectTimeout int    `position:"Query" name:"HealthCheckConnectTimeout"`
	ResourceOwnerId           int64  `position:"Query" name:"ResourceOwnerId"`
	UnhealthyThreshold        int    `position:"Query" name:"UnhealthyThreshold"`
	HealthyThreshold          int    `position:"Query" name:"HealthyThreshold"`
	Scheduler                 string `position:"Query" name:"Scheduler"`
	MaxConnection             int    `position:"Query" name:"MaxConnection"`
	PersistenceTimeout        int    `position:"Query" name:"PersistenceTimeout"`
	VServerGroupId            string `position:"Query" name:"VServerGroupId"`
	ListenerPort              int    `position:"Query" name:"ListenerPort"`
	ResourceOwnerAccount      string `position:"Query" name:"ResourceOwnerAccount"`
	Bandwidth                 int    `position:"Query" name:"Bandwidth"`
	OwnerAccount              string `position:"Query" name:"OwnerAccount"`
	OwnerId                   int64  `position:"Query" name:"OwnerId"`
	Tags                      string `position:"Query" name:"Tags"`
	LoadBalancerId            string `position:"Query" name:"LoadBalancerId"`
	MasterSlaveServerGroupId  string `position:"Query" name:"MasterSlaveServerGroupId"`
	HealthCheckReq            string `position:"Query" name:"HealthCheckReq"`
	BackendServerPort         int    `position:"Query" name:"BackendServerPort"`
	HealthCheckInterval       int    `position:"Query" name:"HealthCheckInterval"`
	HealthCheckExp            string `position:"Query" name:"HealthCheckExp"`
	HealthCheckConnectPort    int    `position:"Query" name:"HealthCheckConnectPort"`
}

func (r CreateLoadBalancerUDPListenerRequest) Invoke(client *sdk.Client) (response *CreateLoadBalancerUDPListenerResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateLoadBalancerUDPListenerRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Slb", "2014-05-15", "CreateLoadBalancerUDPListener", "slb", "")

	resp := struct {
		*responses.BaseResponse
		CreateLoadBalancerUDPListenerResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateLoadBalancerUDPListenerResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateLoadBalancerUDPListenerResponse struct {
	RequestId string
}
