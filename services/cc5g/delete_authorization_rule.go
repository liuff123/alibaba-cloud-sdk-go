package cc5g

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

// DeleteAuthorizationRule invokes the cc5g.DeleteAuthorizationRule API synchronously
func (client *Client) DeleteAuthorizationRule(request *DeleteAuthorizationRuleRequest) (response *DeleteAuthorizationRuleResponse, err error) {
	response = CreateDeleteAuthorizationRuleResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteAuthorizationRuleWithChan invokes the cc5g.DeleteAuthorizationRule API asynchronously
func (client *Client) DeleteAuthorizationRuleWithChan(request *DeleteAuthorizationRuleRequest) (<-chan *DeleteAuthorizationRuleResponse, <-chan error) {
	responseChan := make(chan *DeleteAuthorizationRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteAuthorizationRule(request)
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

// DeleteAuthorizationRuleWithCallback invokes the cc5g.DeleteAuthorizationRule API asynchronously
func (client *Client) DeleteAuthorizationRuleWithCallback(request *DeleteAuthorizationRuleRequest, callback func(response *DeleteAuthorizationRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteAuthorizationRuleResponse
		var err error
		defer close(result)
		response, err = client.DeleteAuthorizationRule(request)
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

// DeleteAuthorizationRuleRequest is the request struct for api DeleteAuthorizationRule
type DeleteAuthorizationRuleRequest struct {
	*requests.RpcRequest
	DryRun                   requests.Boolean `position:"Query" name:"DryRun"`
	ClientToken              string           `position:"Query" name:"ClientToken"`
	AuthorizationRuleId      string           `position:"Query" name:"AuthorizationRuleId"`
	WirelessCloudConnectorId string           `position:"Query" name:"WirelessCloudConnectorId"`
}

// DeleteAuthorizationRuleResponse is the response struct for api DeleteAuthorizationRule
type DeleteAuthorizationRuleResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteAuthorizationRuleRequest creates a request to invoke DeleteAuthorizationRule API
func CreateDeleteAuthorizationRuleRequest() (request *DeleteAuthorizationRuleRequest) {
	request = &DeleteAuthorizationRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CC5G", "2022-03-14", "DeleteAuthorizationRule", "fivegcc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteAuthorizationRuleResponse creates a response to parse from DeleteAuthorizationRule response
func CreateDeleteAuthorizationRuleResponse() (response *DeleteAuthorizationRuleResponse) {
	response = &DeleteAuthorizationRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
