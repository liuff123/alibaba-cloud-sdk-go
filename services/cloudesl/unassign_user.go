package cloudesl

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

// UnassignUser invokes the cloudesl.UnassignUser API synchronously
func (client *Client) UnassignUser(request *UnassignUserRequest) (response *UnassignUserResponse, err error) {
	response = CreateUnassignUserResponse()
	err = client.DoAction(request, response)
	return
}

// UnassignUserWithChan invokes the cloudesl.UnassignUser API asynchronously
func (client *Client) UnassignUserWithChan(request *UnassignUserRequest) (<-chan *UnassignUserResponse, <-chan error) {
	responseChan := make(chan *UnassignUserResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UnassignUser(request)
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

// UnassignUserWithCallback invokes the cloudesl.UnassignUser API asynchronously
func (client *Client) UnassignUserWithCallback(request *UnassignUserRequest, callback func(response *UnassignUserResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UnassignUserResponse
		var err error
		defer close(result)
		response, err = client.UnassignUser(request)
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

// UnassignUserRequest is the request struct for api UnassignUser
type UnassignUserRequest struct {
	*requests.RpcRequest
	ExtraParams string `position:"Body" name:"ExtraParams"`
	UserId      string `position:"Body" name:"UserId"`
}

// UnassignUserResponse is the response struct for api UnassignUser
type UnassignUserResponse struct {
	*responses.BaseResponse
	ErrorMessage   string `json:"ErrorMessage" xml:"ErrorMessage"`
	ErrorCode      string `json:"ErrorCode" xml:"ErrorCode"`
	Message        string `json:"Message" xml:"Message"`
	DynamicCode    string `json:"DynamicCode" xml:"DynamicCode"`
	Code           string `json:"Code" xml:"Code"`
	DynamicMessage string `json:"DynamicMessage" xml:"DynamicMessage"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
}

// CreateUnassignUserRequest creates a request to invoke UnassignUser API
func CreateUnassignUserRequest() (request *UnassignUserRequest) {
	request = &UnassignUserRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudesl", "2020-02-01", "UnassignUser", "", "")
	request.Method = requests.POST
	return
}

// CreateUnassignUserResponse creates a response to parse from UnassignUser response
func CreateUnassignUserResponse() (response *UnassignUserResponse) {
	response = &UnassignUserResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
