package cdn

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

// PushObjectCache invokes the cdn.PushObjectCache API synchronously
func (client *Client) PushObjectCache(request *PushObjectCacheRequest) (response *PushObjectCacheResponse, err error) {
	response = CreatePushObjectCacheResponse()
	err = client.DoAction(request, response)
	return
}

// PushObjectCacheWithChan invokes the cdn.PushObjectCache API asynchronously
func (client *Client) PushObjectCacheWithChan(request *PushObjectCacheRequest) (<-chan *PushObjectCacheResponse, <-chan error) {
	responseChan := make(chan *PushObjectCacheResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PushObjectCache(request)
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

// PushObjectCacheWithCallback invokes the cdn.PushObjectCache API asynchronously
func (client *Client) PushObjectCacheWithCallback(request *PushObjectCacheRequest, callback func(response *PushObjectCacheResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PushObjectCacheResponse
		var err error
		defer close(result)
		response, err = client.PushObjectCache(request)
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

// PushObjectCacheRequest is the request struct for api PushObjectCache
type PushObjectCacheRequest struct {
	*requests.RpcRequest
	ObjectPath    string           `position:"Query" name:"ObjectPath"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	L2Preload     requests.Boolean `position:"Query" name:"L2Preload"`
	Area          string           `position:"Query" name:"Area"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
}

// PushObjectCacheResponse is the response struct for api PushObjectCache
type PushObjectCacheResponse struct {
	*responses.BaseResponse
	PushTaskId string `json:"PushTaskId" xml:"PushTaskId"`
	RequestId  string `json:"RequestId" xml:"RequestId"`
}

// CreatePushObjectCacheRequest creates a request to invoke PushObjectCache API
func CreatePushObjectCacheRequest() (request *PushObjectCacheRequest) {
	request = &PushObjectCacheRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "PushObjectCache", "", "")
	request.Method = requests.POST
	return
}

// CreatePushObjectCacheResponse creates a response to parse from PushObjectCache response
func CreatePushObjectCacheResponse() (response *PushObjectCacheResponse) {
	response = &PushObjectCacheResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
