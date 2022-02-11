package iot

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

// DisableDeviceTunnelShare invokes the iot.DisableDeviceTunnelShare API synchronously
func (client *Client) DisableDeviceTunnelShare(request *DisableDeviceTunnelShareRequest) (response *DisableDeviceTunnelShareResponse, err error) {
	response = CreateDisableDeviceTunnelShareResponse()
	err = client.DoAction(request, response)
	return
}

// DisableDeviceTunnelShareWithChan invokes the iot.DisableDeviceTunnelShare API asynchronously
func (client *Client) DisableDeviceTunnelShareWithChan(request *DisableDeviceTunnelShareRequest) (<-chan *DisableDeviceTunnelShareResponse, <-chan error) {
	responseChan := make(chan *DisableDeviceTunnelShareResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DisableDeviceTunnelShare(request)
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

// DisableDeviceTunnelShareWithCallback invokes the iot.DisableDeviceTunnelShare API asynchronously
func (client *Client) DisableDeviceTunnelShareWithCallback(request *DisableDeviceTunnelShareRequest, callback func(response *DisableDeviceTunnelShareResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DisableDeviceTunnelShareResponse
		var err error
		defer close(result)
		response, err = client.DisableDeviceTunnelShare(request)
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

// DisableDeviceTunnelShareRequest is the request struct for api DisableDeviceTunnelShare
type DisableDeviceTunnelShareRequest struct {
	*requests.RpcRequest
	IotInstanceId string `position:"Query" name:"IotInstanceId"`
	ProductKey    string `position:"Query" name:"ProductKey"`
	ApiProduct    string `position:"Body" name:"ApiProduct"`
	ApiRevision   string `position:"Body" name:"ApiRevision"`
	DeviceName    string `position:"Query" name:"DeviceName"`
}

// DisableDeviceTunnelShareResponse is the response struct for api DisableDeviceTunnelShare
type DisableDeviceTunnelShareResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Code         string `json:"Code" xml:"Code"`
}

// CreateDisableDeviceTunnelShareRequest creates a request to invoke DisableDeviceTunnelShare API
func CreateDisableDeviceTunnelShareRequest() (request *DisableDeviceTunnelShareRequest) {
	request = &DisableDeviceTunnelShareRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "DisableDeviceTunnelShare", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDisableDeviceTunnelShareResponse creates a response to parse from DisableDeviceTunnelShare response
func CreateDisableDeviceTunnelShareResponse() (response *DisableDeviceTunnelShareResponse) {
	response = &DisableDeviceTunnelShareResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
