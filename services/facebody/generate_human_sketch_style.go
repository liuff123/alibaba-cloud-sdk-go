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

// GenerateHumanSketchStyle invokes the facebody.GenerateHumanSketchStyle API synchronously
func (client *Client) GenerateHumanSketchStyle(request *GenerateHumanSketchStyleRequest) (response *GenerateHumanSketchStyleResponse, err error) {
	response = CreateGenerateHumanSketchStyleResponse()
	err = client.DoAction(request, response)
	return
}

// GenerateHumanSketchStyleWithChan invokes the facebody.GenerateHumanSketchStyle API asynchronously
func (client *Client) GenerateHumanSketchStyleWithChan(request *GenerateHumanSketchStyleRequest) (<-chan *GenerateHumanSketchStyleResponse, <-chan error) {
	responseChan := make(chan *GenerateHumanSketchStyleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GenerateHumanSketchStyle(request)
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

// GenerateHumanSketchStyleWithCallback invokes the facebody.GenerateHumanSketchStyle API asynchronously
func (client *Client) GenerateHumanSketchStyleWithCallback(request *GenerateHumanSketchStyleRequest, callback func(response *GenerateHumanSketchStyleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GenerateHumanSketchStyleResponse
		var err error
		defer close(result)
		response, err = client.GenerateHumanSketchStyle(request)
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

// GenerateHumanSketchStyleRequest is the request struct for api GenerateHumanSketchStyle
type GenerateHumanSketchStyleRequest struct {
	*requests.RpcRequest
	ReturnType string `position:"Body" name:"ReturnType"`
	ImageURL   string `position:"Body" name:"ImageURL"`
}

// GenerateHumanSketchStyleResponse is the response struct for api GenerateHumanSketchStyle
type GenerateHumanSketchStyleResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateGenerateHumanSketchStyleRequest creates a request to invoke GenerateHumanSketchStyle API
func CreateGenerateHumanSketchStyleRequest() (request *GenerateHumanSketchStyleRequest) {
	request = &GenerateHumanSketchStyleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("facebody", "2019-12-30", "GenerateHumanSketchStyle", "", "")
	request.Method = requests.POST
	return
}

// CreateGenerateHumanSketchStyleResponse creates a response to parse from GenerateHumanSketchStyle response
func CreateGenerateHumanSketchStyleResponse() (response *GenerateHumanSketchStyleResponse) {
	response = &GenerateHumanSketchStyleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
