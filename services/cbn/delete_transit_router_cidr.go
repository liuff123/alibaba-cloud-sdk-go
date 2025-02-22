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

// DeleteTransitRouterCidr invokes the cbn.DeleteTransitRouterCidr API synchronously
func (client *Client) DeleteTransitRouterCidr(request *DeleteTransitRouterCidrRequest) (response *DeleteTransitRouterCidrResponse, err error) {
	response = CreateDeleteTransitRouterCidrResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteTransitRouterCidrWithChan invokes the cbn.DeleteTransitRouterCidr API asynchronously
func (client *Client) DeleteTransitRouterCidrWithChan(request *DeleteTransitRouterCidrRequest) (<-chan *DeleteTransitRouterCidrResponse, <-chan error) {
	responseChan := make(chan *DeleteTransitRouterCidrResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteTransitRouterCidr(request)
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

// DeleteTransitRouterCidrWithCallback invokes the cbn.DeleteTransitRouterCidr API asynchronously
func (client *Client) DeleteTransitRouterCidrWithCallback(request *DeleteTransitRouterCidrRequest, callback func(response *DeleteTransitRouterCidrResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteTransitRouterCidrResponse
		var err error
		defer close(result)
		response, err = client.DeleteTransitRouterCidr(request)
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

// DeleteTransitRouterCidrRequest is the request struct for api DeleteTransitRouterCidr
type DeleteTransitRouterCidrRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	TransitRouterCidrId  string           `position:"Query" name:"TransitRouterCidrId"`
	DryRun               requests.Boolean `position:"Query" name:"DryRun"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	TransitRouterId      string           `position:"Query" name:"TransitRouterId"`
}

// DeleteTransitRouterCidrResponse is the response struct for api DeleteTransitRouterCidr
type DeleteTransitRouterCidrResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteTransitRouterCidrRequest creates a request to invoke DeleteTransitRouterCidr API
func CreateDeleteTransitRouterCidrRequest() (request *DeleteTransitRouterCidrRequest) {
	request = &DeleteTransitRouterCidrRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cbn", "2017-09-12", "DeleteTransitRouterCidr", "cbn", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteTransitRouterCidrResponse creates a response to parse from DeleteTransitRouterCidr response
func CreateDeleteTransitRouterCidrResponse() (response *DeleteTransitRouterCidrResponse) {
	response = &DeleteTransitRouterCidrResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
