package ess

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

// ModifyLifecycleHook invokes the ess.ModifyLifecycleHook API synchronously
// api document: https://help.aliyun.com/api/ess/modifylifecyclehook.html
func (client *Client) ModifyLifecycleHook(request *ModifyLifecycleHookRequest) (response *ModifyLifecycleHookResponse, err error) {
	response = CreateModifyLifecycleHookResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyLifecycleHookWithChan invokes the ess.ModifyLifecycleHook API asynchronously
// api document: https://help.aliyun.com/api/ess/modifylifecyclehook.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyLifecycleHookWithChan(request *ModifyLifecycleHookRequest) (<-chan *ModifyLifecycleHookResponse, <-chan error) {
	responseChan := make(chan *ModifyLifecycleHookResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyLifecycleHook(request)
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

// ModifyLifecycleHookWithCallback invokes the ess.ModifyLifecycleHook API asynchronously
// api document: https://help.aliyun.com/api/ess/modifylifecyclehook.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyLifecycleHookWithCallback(request *ModifyLifecycleHookRequest, callback func(response *ModifyLifecycleHookResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyLifecycleHookResponse
		var err error
		defer close(result)
		response, err = client.ModifyLifecycleHook(request)
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

// ModifyLifecycleHookRequest is the request struct for api ModifyLifecycleHook
type ModifyLifecycleHookRequest struct {
	*requests.RpcRequest
	DefaultResult        string           `position:"Query" name:"DefaultResult"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	HeartbeatTimeout     requests.Integer `position:"Query" name:"HeartbeatTimeout"`
	LifecycleHookId      string           `position:"Query" name:"LifecycleHookId"`
	ScalingGroupId       string           `position:"Query" name:"ScalingGroupId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	NotificationMetadata string           `position:"Query" name:"NotificationMetadata"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	LifecycleTransition  string           `position:"Query" name:"LifecycleTransition"`
	LifecycleHookName    string           `position:"Query" name:"LifecycleHookName"`
	NotificationArn      string           `position:"Query" name:"NotificationArn"`
}

// ModifyLifecycleHookResponse is the response struct for api ModifyLifecycleHook
type ModifyLifecycleHookResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyLifecycleHookRequest creates a request to invoke ModifyLifecycleHook API
func CreateModifyLifecycleHookRequest() (request *ModifyLifecycleHookRequest) {
	request = &ModifyLifecycleHookRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ess", "2014-08-28", "ModifyLifecycleHook", "", "")
	return
}

// CreateModifyLifecycleHookResponse creates a response to parse from ModifyLifecycleHook response
func CreateModifyLifecycleHookResponse() (response *ModifyLifecycleHookResponse) {
	response = &ModifyLifecycleHookResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
