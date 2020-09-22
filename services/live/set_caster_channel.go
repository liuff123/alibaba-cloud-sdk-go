package live

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

// SetCasterChannel invokes the live.SetCasterChannel API synchronously
func (client *Client) SetCasterChannel(request *SetCasterChannelRequest) (response *SetCasterChannelResponse, err error) {
	response = CreateSetCasterChannelResponse()
	err = client.DoAction(request, response)
	return
}

// SetCasterChannelWithChan invokes the live.SetCasterChannel API asynchronously
func (client *Client) SetCasterChannelWithChan(request *SetCasterChannelRequest) (<-chan *SetCasterChannelResponse, <-chan error) {
	responseChan := make(chan *SetCasterChannelResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetCasterChannel(request)
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

// SetCasterChannelWithCallback invokes the live.SetCasterChannel API asynchronously
func (client *Client) SetCasterChannelWithCallback(request *SetCasterChannelRequest, callback func(response *SetCasterChannelResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetCasterChannelResponse
		var err error
		defer close(result)
		response, err = client.SetCasterChannel(request)
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

// SetCasterChannelRequest is the request struct for api SetCasterChannel
type SetCasterChannelRequest struct {
	*requests.RpcRequest
	SeekOffset requests.Integer `position:"Query" name:"SeekOffset"`
	PlayStatus requests.Integer `position:"Query" name:"PlayStatus"`
	ResourceId string           `position:"Query" name:"ResourceId"`
	CasterId   string           `position:"Query" name:"CasterId"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
	ReloadFlag requests.Integer `position:"Query" name:"ReloadFlag"`
	ChannelId  string           `position:"Query" name:"ChannelId"`
}

// SetCasterChannelResponse is the response struct for api SetCasterChannel
type SetCasterChannelResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSetCasterChannelRequest creates a request to invoke SetCasterChannel API
func CreateSetCasterChannelRequest() (request *SetCasterChannelRequest) {
	request = &SetCasterChannelRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "SetCasterChannel", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSetCasterChannelResponse creates a response to parse from SetCasterChannel response
func CreateSetCasterChannelResponse() (response *SetCasterChannelResponse) {
	response = &SetCasterChannelResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
