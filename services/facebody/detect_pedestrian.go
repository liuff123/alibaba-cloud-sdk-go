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

// DetectPedestrian invokes the facebody.DetectPedestrian API synchronously
func (client *Client) DetectPedestrian(request *DetectPedestrianRequest) (response *DetectPedestrianResponse, err error) {
	response = CreateDetectPedestrianResponse()
	err = client.DoAction(request, response)
	return
}

// DetectPedestrianWithChan invokes the facebody.DetectPedestrian API asynchronously
func (client *Client) DetectPedestrianWithChan(request *DetectPedestrianRequest) (<-chan *DetectPedestrianResponse, <-chan error) {
	responseChan := make(chan *DetectPedestrianResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DetectPedestrian(request)
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

// DetectPedestrianWithCallback invokes the facebody.DetectPedestrian API asynchronously
func (client *Client) DetectPedestrianWithCallback(request *DetectPedestrianRequest, callback func(response *DetectPedestrianResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DetectPedestrianResponse
		var err error
		defer close(result)
		response, err = client.DetectPedestrian(request)
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

// DetectPedestrianRequest is the request struct for api DetectPedestrian
type DetectPedestrianRequest struct {
	*requests.RpcRequest
	ImageURL string `position:"Body" name:"ImageURL"`
}

// DetectPedestrianResponse is the response struct for api DetectPedestrian
type DetectPedestrianResponse struct {
	*responses.BaseResponse
	RequestId string                 `json:"RequestId" xml:"RequestId"`
	Data      DataInDetectPedestrian `json:"Data" xml:"Data"`
}

// CreateDetectPedestrianRequest creates a request to invoke DetectPedestrian API
func CreateDetectPedestrianRequest() (request *DetectPedestrianRequest) {
	request = &DetectPedestrianRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("facebody", "2019-12-30", "DetectPedestrian", "", "")
	request.Method = requests.POST
	return
}

// CreateDetectPedestrianResponse creates a response to parse from DetectPedestrian response
func CreateDetectPedestrianResponse() (response *DetectPedestrianResponse) {
	response = &DetectPedestrianResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
