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

// ListTransitRouterPeerAttachments invokes the cbn.ListTransitRouterPeerAttachments API synchronously
func (client *Client) ListTransitRouterPeerAttachments(request *ListTransitRouterPeerAttachmentsRequest) (response *ListTransitRouterPeerAttachmentsResponse, err error) {
	response = CreateListTransitRouterPeerAttachmentsResponse()
	err = client.DoAction(request, response)
	return
}

// ListTransitRouterPeerAttachmentsWithChan invokes the cbn.ListTransitRouterPeerAttachments API asynchronously
func (client *Client) ListTransitRouterPeerAttachmentsWithChan(request *ListTransitRouterPeerAttachmentsRequest) (<-chan *ListTransitRouterPeerAttachmentsResponse, <-chan error) {
	responseChan := make(chan *ListTransitRouterPeerAttachmentsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListTransitRouterPeerAttachments(request)
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

// ListTransitRouterPeerAttachmentsWithCallback invokes the cbn.ListTransitRouterPeerAttachments API asynchronously
func (client *Client) ListTransitRouterPeerAttachmentsWithCallback(request *ListTransitRouterPeerAttachmentsRequest, callback func(response *ListTransitRouterPeerAttachmentsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListTransitRouterPeerAttachmentsResponse
		var err error
		defer close(result)
		response, err = client.ListTransitRouterPeerAttachments(request)
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

// ListTransitRouterPeerAttachmentsRequest is the request struct for api ListTransitRouterPeerAttachments
type ListTransitRouterPeerAttachmentsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId           requests.Integer                       `position:"Query" name:"ResourceOwnerId"`
	CenId                     string                                 `position:"Query" name:"CenId"`
	NextToken                 string                                 `position:"Query" name:"NextToken"`
	Tag                       *[]ListTransitRouterPeerAttachmentsTag `position:"Query" name:"Tag"  type:"Repeated"`
	ResourceOwnerAccount      string                                 `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount              string                                 `position:"Query" name:"OwnerAccount"`
	OwnerId                   requests.Integer                       `position:"Query" name:"OwnerId"`
	TransitRouterId           string                                 `position:"Query" name:"TransitRouterId"`
	ResourceType              string                                 `position:"Query" name:"ResourceType"`
	TransitRouterAttachmentId string                                 `position:"Query" name:"TransitRouterAttachmentId"`
	MaxResults                requests.Integer                       `position:"Query" name:"MaxResults"`
}

// ListTransitRouterPeerAttachmentsTag is a repeated param struct in ListTransitRouterPeerAttachmentsRequest
type ListTransitRouterPeerAttachmentsTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// ListTransitRouterPeerAttachmentsResponse is the response struct for api ListTransitRouterPeerAttachments
type ListTransitRouterPeerAttachmentsResponse struct {
	*responses.BaseResponse
	NextToken                string                    `json:"NextToken" xml:"NextToken"`
	RequestId                string                    `json:"RequestId" xml:"RequestId"`
	TotalCount               int                       `json:"TotalCount" xml:"TotalCount"`
	MaxResults               int                       `json:"MaxResults" xml:"MaxResults"`
	TransitRouterAttachments []TransitRouterAttachment `json:"TransitRouterAttachments" xml:"TransitRouterAttachments"`
}

// CreateListTransitRouterPeerAttachmentsRequest creates a request to invoke ListTransitRouterPeerAttachments API
func CreateListTransitRouterPeerAttachmentsRequest() (request *ListTransitRouterPeerAttachmentsRequest) {
	request = &ListTransitRouterPeerAttachmentsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cbn", "2017-09-12", "ListTransitRouterPeerAttachments", "cbn", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListTransitRouterPeerAttachmentsResponse creates a response to parse from ListTransitRouterPeerAttachments response
func CreateListTransitRouterPeerAttachmentsResponse() (response *ListTransitRouterPeerAttachmentsResponse) {
	response = &ListTransitRouterPeerAttachmentsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
