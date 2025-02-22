package dcdn

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

// DescribeRoutineCanaryEnvs invokes the dcdn.DescribeRoutineCanaryEnvs API synchronously
func (client *Client) DescribeRoutineCanaryEnvs(request *DescribeRoutineCanaryEnvsRequest) (response *DescribeRoutineCanaryEnvsResponse, err error) {
	response = CreateDescribeRoutineCanaryEnvsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeRoutineCanaryEnvsWithChan invokes the dcdn.DescribeRoutineCanaryEnvs API asynchronously
func (client *Client) DescribeRoutineCanaryEnvsWithChan(request *DescribeRoutineCanaryEnvsRequest) (<-chan *DescribeRoutineCanaryEnvsResponse, <-chan error) {
	responseChan := make(chan *DescribeRoutineCanaryEnvsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeRoutineCanaryEnvs(request)
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

// DescribeRoutineCanaryEnvsWithCallback invokes the dcdn.DescribeRoutineCanaryEnvs API asynchronously
func (client *Client) DescribeRoutineCanaryEnvsWithCallback(request *DescribeRoutineCanaryEnvsRequest, callback func(response *DescribeRoutineCanaryEnvsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeRoutineCanaryEnvsResponse
		var err error
		defer close(result)
		response, err = client.DescribeRoutineCanaryEnvs(request)
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

// DescribeRoutineCanaryEnvsRequest is the request struct for api DescribeRoutineCanaryEnvs
type DescribeRoutineCanaryEnvsRequest struct {
	*requests.RpcRequest
	OwnerId requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeRoutineCanaryEnvsResponse is the response struct for api DescribeRoutineCanaryEnvs
type DescribeRoutineCanaryEnvsResponse struct {
	*responses.BaseResponse
	Content   map[string]interface{} `json:"Content" xml:"Content"`
	RequestId string                 `json:"RequestId" xml:"RequestId"`
}

// CreateDescribeRoutineCanaryEnvsRequest creates a request to invoke DescribeRoutineCanaryEnvs API
func CreateDescribeRoutineCanaryEnvsRequest() (request *DescribeRoutineCanaryEnvsRequest) {
	request = &DescribeRoutineCanaryEnvsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "DescribeRoutineCanaryEnvs", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeRoutineCanaryEnvsResponse creates a response to parse from DescribeRoutineCanaryEnvs response
func CreateDescribeRoutineCanaryEnvsResponse() (response *DescribeRoutineCanaryEnvsResponse) {
	response = &DescribeRoutineCanaryEnvsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
