package lmztest

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

// DemoGreyReleaseTest invokes the lmztest.DemoGreyReleaseTest API synchronously
func (client *Client) DemoGreyReleaseTest(request *DemoGreyReleaseTestRequest) (response *DemoGreyReleaseTestResponse, err error) {
	response = CreateDemoGreyReleaseTestResponse()
	err = client.DoAction(request, response)
	return
}

// DemoGreyReleaseTestWithChan invokes the lmztest.DemoGreyReleaseTest API asynchronously
func (client *Client) DemoGreyReleaseTestWithChan(request *DemoGreyReleaseTestRequest) (<-chan *DemoGreyReleaseTestResponse, <-chan error) {
	responseChan := make(chan *DemoGreyReleaseTestResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DemoGreyReleaseTest(request)
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

// DemoGreyReleaseTestWithCallback invokes the lmztest.DemoGreyReleaseTest API asynchronously
func (client *Client) DemoGreyReleaseTestWithCallback(request *DemoGreyReleaseTestRequest, callback func(response *DemoGreyReleaseTestResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DemoGreyReleaseTestResponse
		var err error
		defer close(result)
		response, err = client.DemoGreyReleaseTest(request)
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

// DemoGreyReleaseTestRequest is the request struct for api DemoGreyReleaseTest
type DemoGreyReleaseTestRequest struct {
	*requests.RpcRequest
	Number requests.Integer `position:"Query" name:"Number"`
}

// DemoGreyReleaseTestResponse is the response struct for api DemoGreyReleaseTest
type DemoGreyReleaseTestResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Id        string `json:"Id" xml:"Id"`
}

// CreateDemoGreyReleaseTestRequest creates a request to invoke DemoGreyReleaseTest API
func CreateDemoGreyReleaseTestRequest() (request *DemoGreyReleaseTestRequest) {
	request = &DemoGreyReleaseTestRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("LmzTest", "2010-10-11", "DemoGreyReleaseTest", "", "")
	request.Method = requests.POST
	return
}

// CreateDemoGreyReleaseTestResponse creates a response to parse from DemoGreyReleaseTest response
func CreateDemoGreyReleaseTestResponse() (response *DemoGreyReleaseTestResponse) {
	response = &DemoGreyReleaseTestResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
