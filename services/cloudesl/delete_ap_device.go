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

// DeleteApDevice invokes the cloudesl.DeleteApDevice API synchronously
func (client *Client) DeleteApDevice(request *DeleteApDeviceRequest) (response *DeleteApDeviceResponse, err error) {
	response = CreateDeleteApDeviceResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteApDeviceWithChan invokes the cloudesl.DeleteApDevice API asynchronously
func (client *Client) DeleteApDeviceWithChan(request *DeleteApDeviceRequest) (<-chan *DeleteApDeviceResponse, <-chan error) {
	responseChan := make(chan *DeleteApDeviceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteApDevice(request)
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

// DeleteApDeviceWithCallback invokes the cloudesl.DeleteApDevice API asynchronously
func (client *Client) DeleteApDeviceWithCallback(request *DeleteApDeviceRequest, callback func(response *DeleteApDeviceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteApDeviceResponse
		var err error
		defer close(result)
		response, err = client.DeleteApDevice(request)
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

// DeleteApDeviceRequest is the request struct for api DeleteApDevice
type DeleteApDeviceRequest struct {
	*requests.RpcRequest
	ExtraParams string `position:"Body" name:"ExtraParams"`
	ApMac       string `position:"Body" name:"ApMac"`
	StoreId     string `position:"Body" name:"StoreId"`
}

// DeleteApDeviceResponse is the response struct for api DeleteApDevice
type DeleteApDeviceResponse struct {
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

// CreateDeleteApDeviceRequest creates a request to invoke DeleteApDevice API
func CreateDeleteApDeviceRequest() (request *DeleteApDeviceRequest) {
	request = &DeleteApDeviceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudesl", "2020-02-01", "DeleteApDevice", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteApDeviceResponse creates a response to parse from DeleteApDevice response
func CreateDeleteApDeviceResponse() (response *DeleteApDeviceResponse) {
	response = &DeleteApDeviceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
