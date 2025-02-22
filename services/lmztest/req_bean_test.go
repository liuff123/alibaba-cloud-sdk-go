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

// ReqBeanTest invokes the lmztest.ReqBeanTest API synchronously
func (client *Client) ReqBeanTest(request *ReqBeanTestRequest) (response *ReqBeanTestResponse, err error) {
	response = CreateReqBeanTestResponse()
	err = client.DoAction(request, response)
	return
}

// ReqBeanTestWithChan invokes the lmztest.ReqBeanTest API asynchronously
func (client *Client) ReqBeanTestWithChan(request *ReqBeanTestRequest) (<-chan *ReqBeanTestResponse, <-chan error) {
	responseChan := make(chan *ReqBeanTestResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ReqBeanTest(request)
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

// ReqBeanTestWithCallback invokes the lmztest.ReqBeanTest API asynchronously
func (client *Client) ReqBeanTestWithCallback(request *ReqBeanTestRequest, callback func(response *ReqBeanTestResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ReqBeanTestResponse
		var err error
		defer close(result)
		response, err = client.ReqBeanTest(request)
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

// ReqBeanTestRequest is the request struct for api ReqBeanTest
type ReqBeanTestRequest struct {
	*requests.RpcRequest
	Codes           *[]string           `position:"Query" name:"codes"  type:"Json"`
	XHostHeaderTest string              `position:"Header" name:"x-host-header-test"`
	Id              string              `position:"Query" name:"Id"`
	Nums            *[]string           `position:"Query" name:"nums"  type:"Json"`
	Users           *[]ReqBeanTestUsers `position:"Query" name:"users"  type:"Json"`
}

// ReqBeanTestUsers is a repeated param struct in ReqBeanTestRequest
type ReqBeanTestUsers struct {
	Name string `name:"name"`
	Id   string `name:"id"`
}

// ReqBeanTestResponse is the response struct for api ReqBeanTest
type ReqBeanTestResponse struct {
	*responses.BaseResponse
	Id   string `json:"Id" xml:"Id"`
	Name string `json:"Name" xml:"Name"`
}

// CreateReqBeanTestRequest creates a request to invoke ReqBeanTest API
func CreateReqBeanTestRequest() (request *ReqBeanTestRequest) {
	request = &ReqBeanTestRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("LmzTest", "2010-10-11", "ReqBeanTest", "", "")
	request.Method = requests.POST
	return
}

// CreateReqBeanTestResponse creates a response to parse from ReqBeanTest response
func CreateReqBeanTestResponse() (response *ReqBeanTestResponse) {
	response = &ReqBeanTestResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
