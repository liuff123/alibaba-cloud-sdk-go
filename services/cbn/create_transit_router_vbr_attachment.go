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

// CreateTransitRouterVbrAttachment invokes the cbn.CreateTransitRouterVbrAttachment API synchronously
func (client *Client) CreateTransitRouterVbrAttachment(request *CreateTransitRouterVbrAttachmentRequest) (response *CreateTransitRouterVbrAttachmentResponse, err error) {
	response = CreateCreateTransitRouterVbrAttachmentResponse()
	err = client.DoAction(request, response)
	return
}

// CreateTransitRouterVbrAttachmentWithChan invokes the cbn.CreateTransitRouterVbrAttachment API asynchronously
func (client *Client) CreateTransitRouterVbrAttachmentWithChan(request *CreateTransitRouterVbrAttachmentRequest) (<-chan *CreateTransitRouterVbrAttachmentResponse, <-chan error) {
	responseChan := make(chan *CreateTransitRouterVbrAttachmentResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateTransitRouterVbrAttachment(request)
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

// CreateTransitRouterVbrAttachmentWithCallback invokes the cbn.CreateTransitRouterVbrAttachment API asynchronously
func (client *Client) CreateTransitRouterVbrAttachmentWithCallback(request *CreateTransitRouterVbrAttachmentRequest, callback func(response *CreateTransitRouterVbrAttachmentResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateTransitRouterVbrAttachmentResponse
		var err error
		defer close(result)
		response, err = client.CreateTransitRouterVbrAttachment(request)
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

// CreateTransitRouterVbrAttachmentRequest is the request struct for api CreateTransitRouterVbrAttachment
type CreateTransitRouterVbrAttachmentRequest struct {
	*requests.RpcRequest
	ResourceOwnerId                    requests.Integer                       `position:"Query" name:"ResourceOwnerId"`
	ClientToken                        string                                 `position:"Query" name:"ClientToken"`
	CenId                              string                                 `position:"Query" name:"CenId"`
	RouteTableAssociationEnabled       requests.Boolean                       `position:"Query" name:"RouteTableAssociationEnabled"`
	VbrId                              string                                 `position:"Query" name:"VbrId"`
	TransitRouterAttachmentName        string                                 `position:"Query" name:"TransitRouterAttachmentName"`
	Tag                                *[]CreateTransitRouterVbrAttachmentTag `position:"Query" name:"Tag"  type:"Repeated"`
	AutoPublishRouteEnabled            requests.Boolean                       `position:"Query" name:"AutoPublishRouteEnabled"`
	RouteTablePropagationEnabled       requests.Boolean                       `position:"Query" name:"RouteTablePropagationEnabled"`
	DryRun                             requests.Boolean                       `position:"Query" name:"DryRun"`
	ResourceOwnerAccount               string                                 `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                       string                                 `position:"Query" name:"OwnerAccount"`
	OwnerId                            requests.Integer                       `position:"Query" name:"OwnerId"`
	TransitRouterId                    string                                 `position:"Query" name:"TransitRouterId"`
	ResourceType                       string                                 `position:"Query" name:"ResourceType"`
	TransitRouterAttachmentDescription string                                 `position:"Query" name:"TransitRouterAttachmentDescription"`
	AssociateRouteTableId              string                                 `position:"Query" name:"AssociateRouteTableId"`
	VbrOwnerId                         requests.Integer                       `position:"Query" name:"VbrOwnerId"`
}

// CreateTransitRouterVbrAttachmentTag is a repeated param struct in CreateTransitRouterVbrAttachmentRequest
type CreateTransitRouterVbrAttachmentTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// CreateTransitRouterVbrAttachmentResponse is the response struct for api CreateTransitRouterVbrAttachment
type CreateTransitRouterVbrAttachmentResponse struct {
	*responses.BaseResponse
	TransitRouterAttachmentId string `json:"TransitRouterAttachmentId" xml:"TransitRouterAttachmentId"`
	RequestId                 string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateTransitRouterVbrAttachmentRequest creates a request to invoke CreateTransitRouterVbrAttachment API
func CreateCreateTransitRouterVbrAttachmentRequest() (request *CreateTransitRouterVbrAttachmentRequest) {
	request = &CreateTransitRouterVbrAttachmentRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cbn", "2017-09-12", "CreateTransitRouterVbrAttachment", "cbn", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateTransitRouterVbrAttachmentResponse creates a response to parse from CreateTransitRouterVbrAttachment response
func CreateCreateTransitRouterVbrAttachmentResponse() (response *CreateTransitRouterVbrAttachmentResponse) {
	response = &CreateTransitRouterVbrAttachmentResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
