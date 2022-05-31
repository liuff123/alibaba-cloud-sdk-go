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

// BindEslDevice invokes the cloudesl.BindEslDevice API synchronously
func (client *Client) BindEslDevice(request *BindEslDeviceRequest) (response *BindEslDeviceResponse, err error) {
	response = CreateBindEslDeviceResponse()
	err = client.DoAction(request, response)
	return
}

// BindEslDeviceWithChan invokes the cloudesl.BindEslDevice API asynchronously
func (client *Client) BindEslDeviceWithChan(request *BindEslDeviceRequest) (<-chan *BindEslDeviceResponse, <-chan error) {
	responseChan := make(chan *BindEslDeviceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BindEslDevice(request)
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

// BindEslDeviceWithCallback invokes the cloudesl.BindEslDevice API asynchronously
func (client *Client) BindEslDeviceWithCallback(request *BindEslDeviceRequest, callback func(response *BindEslDeviceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BindEslDeviceResponse
		var err error
		defer close(result)
		response, err = client.BindEslDevice(request)
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

// BindEslDeviceRequest is the request struct for api BindEslDevice
type BindEslDeviceRequest struct {
	*requests.RpcRequest
	ExtraParams string           `position:"Body" name:"ExtraParams"`
	StoreId     string           `position:"Body" name:"StoreId"`
	Layer       requests.Integer `position:"Body" name:"Layer"`
	EslBarCode  string           `position:"Body" name:"EslBarCode"`
	ItemBarCode string           `position:"Body" name:"ItemBarCode"`
	Column      string           `position:"Body" name:"Column"`
	Shelf       string           `position:"Body" name:"Shelf"`
}

// BindEslDeviceResponse is the response struct for api BindEslDevice
type BindEslDeviceResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	ErrorMessage   string `json:"ErrorMessage" xml:"ErrorMessage"`
	Success        bool   `json:"Success" xml:"Success"`
	ErrorCode      string `json:"ErrorCode" xml:"ErrorCode"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	DynamicMessage string `json:"DynamicMessage" xml:"DynamicMessage"`
	DynamicCode    string `json:"DynamicCode" xml:"DynamicCode"`
}

// CreateBindEslDeviceRequest creates a request to invoke BindEslDevice API
func CreateBindEslDeviceRequest() (request *BindEslDeviceRequest) {
	request = &BindEslDeviceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudesl", "2020-02-01", "BindEslDevice", "", "")
	request.Method = requests.POST
	return
}

// CreateBindEslDeviceResponse creates a response to parse from BindEslDevice response
func CreateBindEslDeviceResponse() (response *BindEslDeviceResponse) {
	response = &BindEslDeviceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
