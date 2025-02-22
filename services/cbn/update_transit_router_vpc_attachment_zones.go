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

// UpdateTransitRouterVpcAttachmentZones invokes the cbn.UpdateTransitRouterVpcAttachmentZones API synchronously
func (client *Client) UpdateTransitRouterVpcAttachmentZones(request *UpdateTransitRouterVpcAttachmentZonesRequest) (response *UpdateTransitRouterVpcAttachmentZonesResponse, err error) {
	response = CreateUpdateTransitRouterVpcAttachmentZonesResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateTransitRouterVpcAttachmentZonesWithChan invokes the cbn.UpdateTransitRouterVpcAttachmentZones API asynchronously
func (client *Client) UpdateTransitRouterVpcAttachmentZonesWithChan(request *UpdateTransitRouterVpcAttachmentZonesRequest) (<-chan *UpdateTransitRouterVpcAttachmentZonesResponse, <-chan error) {
	responseChan := make(chan *UpdateTransitRouterVpcAttachmentZonesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateTransitRouterVpcAttachmentZones(request)
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

// UpdateTransitRouterVpcAttachmentZonesWithCallback invokes the cbn.UpdateTransitRouterVpcAttachmentZones API asynchronously
func (client *Client) UpdateTransitRouterVpcAttachmentZonesWithCallback(request *UpdateTransitRouterVpcAttachmentZonesRequest, callback func(response *UpdateTransitRouterVpcAttachmentZonesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateTransitRouterVpcAttachmentZonesResponse
		var err error
		defer close(result)
		response, err = client.UpdateTransitRouterVpcAttachmentZones(request)
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

// UpdateTransitRouterVpcAttachmentZonesRequest is the request struct for api UpdateTransitRouterVpcAttachmentZones
type UpdateTransitRouterVpcAttachmentZonesRequest struct {
	*requests.RpcRequest
	ResourceOwnerId           requests.Integer                                           `position:"Query" name:"ResourceOwnerId"`
	ClientToken               string                                                     `position:"Query" name:"ClientToken"`
	RemoveZoneMappings        *[]UpdateTransitRouterVpcAttachmentZonesRemoveZoneMappings `position:"Query" name:"RemoveZoneMappings"  type:"Repeated"`
	AddZoneMappings           *[]UpdateTransitRouterVpcAttachmentZonesAddZoneMappings    `position:"Query" name:"AddZoneMappings"  type:"Repeated"`
	DryRun                    requests.Boolean                                           `position:"Query" name:"DryRun"`
	ResourceOwnerAccount      string                                                     `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount              string                                                     `position:"Query" name:"OwnerAccount"`
	OwnerId                   requests.Integer                                           `position:"Query" name:"OwnerId"`
	TransitRouterAttachmentId string                                                     `position:"Query" name:"TransitRouterAttachmentId"`
}

// UpdateTransitRouterVpcAttachmentZonesRemoveZoneMappings is a repeated param struct in UpdateTransitRouterVpcAttachmentZonesRequest
type UpdateTransitRouterVpcAttachmentZonesRemoveZoneMappings struct {
	VSwitchId string `name:"VSwitchId"`
	ZoneId    string `name:"ZoneId"`
}

// UpdateTransitRouterVpcAttachmentZonesAddZoneMappings is a repeated param struct in UpdateTransitRouterVpcAttachmentZonesRequest
type UpdateTransitRouterVpcAttachmentZonesAddZoneMappings struct {
	VSwitchId string `name:"VSwitchId"`
	ZoneId    string `name:"ZoneId"`
}

// UpdateTransitRouterVpcAttachmentZonesResponse is the response struct for api UpdateTransitRouterVpcAttachmentZones
type UpdateTransitRouterVpcAttachmentZonesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateTransitRouterVpcAttachmentZonesRequest creates a request to invoke UpdateTransitRouterVpcAttachmentZones API
func CreateUpdateTransitRouterVpcAttachmentZonesRequest() (request *UpdateTransitRouterVpcAttachmentZonesRequest) {
	request = &UpdateTransitRouterVpcAttachmentZonesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cbn", "2017-09-12", "UpdateTransitRouterVpcAttachmentZones", "cbn", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateTransitRouterVpcAttachmentZonesResponse creates a response to parse from UpdateTransitRouterVpcAttachmentZones response
func CreateUpdateTransitRouterVpcAttachmentZonesResponse() (response *UpdateTransitRouterVpcAttachmentZonesResponse) {
	response = &UpdateTransitRouterVpcAttachmentZonesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
