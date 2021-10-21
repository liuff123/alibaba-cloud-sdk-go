package facebody

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

// DetectMask invokes the facebody.DetectMask API synchronously
func (client *Client) DetectMask(request *DetectMaskRequest) (response *DetectMaskResponse, err error) {
	response = CreateDetectMaskResponse()
	err = client.DoAction(request, response)
	return
}

// DetectMaskWithChan invokes the facebody.DetectMask API asynchronously
func (client *Client) DetectMaskWithChan(request *DetectMaskRequest) (<-chan *DetectMaskResponse, <-chan error) {
	responseChan := make(chan *DetectMaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DetectMask(request)
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

// DetectMaskWithCallback invokes the facebody.DetectMask API asynchronously
func (client *Client) DetectMaskWithCallback(request *DetectMaskRequest, callback func(response *DetectMaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DetectMaskResponse
		var err error
		defer close(result)
		response, err = client.DetectMask(request)
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

// DetectMaskRequest is the request struct for api DetectMask
type DetectMaskRequest struct {
	*requests.RpcRequest
	ImageURL string `position:"Body" name:"ImageURL"`
}

// DetectMaskResponse is the response struct for api DetectMask
type DetectMaskResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateDetectMaskRequest creates a request to invoke DetectMask API
func CreateDetectMaskRequest() (request *DetectMaskRequest) {
	request = &DetectMaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("facebody", "2019-12-30", "DetectMask", "", "")
	request.Method = requests.POST
	return
}

// CreateDetectMaskResponse creates a response to parse from DetectMask response
func CreateDetectMaskResponse() (response *DetectMaskResponse) {
	response = &DetectMaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
