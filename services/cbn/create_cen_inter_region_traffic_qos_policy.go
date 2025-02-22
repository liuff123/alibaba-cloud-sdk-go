package cbn

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// CreateCenInterRegionTrafficQosPolicy invokes the cbn.CreateCenInterRegionTrafficQosPolicy API synchronously
func (client *Client) CreateCenInterRegionTrafficQosPolicy(request *CreateCenInterRegionTrafficQosPolicyRequest) (response *CreateCenInterRegionTrafficQosPolicyResponse, err error) {
	response = CreateCreateCenInterRegionTrafficQosPolicyResponse()
	err = client.DoAction(request, response)
	return
}

// CreateCenInterRegionTrafficQosPolicyWithChan invokes the cbn.CreateCenInterRegionTrafficQosPolicy API asynchronously
func (client *Client) CreateCenInterRegionTrafficQosPolicyWithChan(request *CreateCenInterRegionTrafficQosPolicyRequest) (<-chan *CreateCenInterRegionTrafficQosPolicyResponse, <-chan error) {
	responseChan := make(chan *CreateCenInterRegionTrafficQosPolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateCenInterRegionTrafficQosPolicy(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// CreateCenInterRegionTrafficQosPolicyWithCallback invokes the cbn.CreateCenInterRegionTrafficQosPolicy API asynchronously
func (client *Client) CreateCenInterRegionTrafficQosPolicyWithCallback(request *CreateCenInterRegionTrafficQosPolicyRequest, callback func(response *CreateCenInterRegionTrafficQosPolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateCenInterRegionTrafficQosPolicyResponse
		var err error
		defer close(result)
		response, err = client.CreateCenInterRegionTrafficQosPolicy(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// CreateCenInterRegionTrafficQosPolicyRequest is the request struct for api CreateCenInterRegionTrafficQosPolicy
type CreateCenInterRegionTrafficQosPolicyRequest struct {
	*requests.RpcRequest
	ResourceOwnerId             requests.Integer                                        `position:"Query" name:"ResourceOwnerId"`
	ClientToken                 string                                                  `position:"Query" name:"ClientToken"`
	TrafficQosQueues            *[]CreateCenInterRegionTrafficQosPolicyTrafficQosQueues `position:"Query" name:"TrafficQosQueues"  type:"Repeated"`
	TrafficQosPolicyName        string                                                  `position:"Query" name:"TrafficQosPolicyName"`
	DryRun                      requests.Boolean                                        `position:"Query" name:"DryRun"`
	ResourceOwnerAccount        string                                                  `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                string                                                  `position:"Query" name:"OwnerAccount"`
	TrafficQosPolicyDescription string                                                  `position:"Query" name:"TrafficQosPolicyDescription"`
	OwnerId                     requests.Integer                                        `position:"Query" name:"OwnerId"`
	TransitRouterId             string                                                  `position:"Query" name:"TransitRouterId"`
	TransitRouterAttachmentId   string                                                  `position:"Query" name:"TransitRouterAttachmentId"`
}

// CreateCenInterRegionTrafficQosPolicyTrafficQosQueues is a repeated param struct in CreateCenInterRegionTrafficQosPolicyRequest
type CreateCenInterRegionTrafficQosPolicyTrafficQosQueues struct {
	Dscps                  *[]string `name:"Dscps" type:"Repeated"`
	QosQueueName           string    `name:"QosQueueName"`
	RemainBandwidthPercent string    `name:"RemainBandwidthPercent"`
	QosQueueDescription    string    `name:"QosQueueDescription"`
}

// CreateCenInterRegionTrafficQosPolicyResponse is the response struct for api CreateCenInterRegionTrafficQosPolicy
type CreateCenInterRegionTrafficQosPolicyResponse struct {
	*responses.BaseResponse
	TrafficQosPolicyId string `json:"TrafficQosPolicyId" xml:"TrafficQosPolicyId"`
	RequestId          string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateCenInterRegionTrafficQosPolicyRequest creates a request to invoke CreateCenInterRegionTrafficQosPolicy API
func CreateCreateCenInterRegionTrafficQosPolicyRequest() (request *CreateCenInterRegionTrafficQosPolicyRequest) {
	request = &CreateCenInterRegionTrafficQosPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cbn", "2017-09-12", "CreateCenInterRegionTrafficQosPolicy", "cbn", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateCenInterRegionTrafficQosPolicyResponse creates a response to parse from CreateCenInterRegionTrafficQosPolicy response
func CreateCreateCenInterRegionTrafficQosPolicyResponse() (response *CreateCenInterRegionTrafficQosPolicyResponse) {
	response = &CreateCenInterRegionTrafficQosPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
