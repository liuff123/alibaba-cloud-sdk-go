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

// UnlockCards invokes the cc5g.UnlockCards API synchronously
func (client *Client) UnlockCards(request *UnlockCardsRequest) (response *UnlockCardsResponse, err error) {
	response = CreateUnlockCardsResponse()
	err = client.DoAction(request, response)
	return
}

// UnlockCardsWithChan invokes the cc5g.UnlockCards API asynchronously
func (client *Client) UnlockCardsWithChan(request *UnlockCardsRequest) (<-chan *UnlockCardsResponse, <-chan error) {
	responseChan := make(chan *UnlockCardsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UnlockCards(request)
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

// UnlockCardsWithCallback invokes the cc5g.UnlockCards API asynchronously
func (client *Client) UnlockCardsWithCallback(request *UnlockCardsRequest, callback func(response *UnlockCardsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UnlockCardsResponse
		var err error
		defer close(result)
		response, err = client.UnlockCards(request)
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

// UnlockCardsRequest is the request struct for api UnlockCards
type UnlockCardsRequest struct {
	*requests.RpcRequest
	Iccids      *[]string        `position:"Query" name:"Iccids"  type:"Repeated"`
	DryRun      requests.Boolean `position:"Query" name:"DryRun"`
	ClientToken string           `position:"Query" name:"ClientToken"`
}

// UnlockCardsResponse is the response struct for api UnlockCards
type UnlockCardsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUnlockCardsRequest creates a request to invoke UnlockCards API
func CreateUnlockCardsRequest() (request *UnlockCardsRequest) {
	request = &UnlockCardsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CC5G", "2022-03-14", "UnlockCards", "fivegcc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUnlockCardsResponse creates a response to parse from UnlockCards response
func CreateUnlockCardsResponse() (response *UnlockCardsResponse) {
	response = &UnlockCardsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
