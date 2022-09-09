package polardb

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

// ModifyPendingMaintenanceAction invokes the polardb.ModifyPendingMaintenanceAction API synchronously
func (client *Client) ModifyPendingMaintenanceAction(request *ModifyPendingMaintenanceActionRequest) (response *ModifyPendingMaintenanceActionResponse, err error) {
	response = CreateModifyPendingMaintenanceActionResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyPendingMaintenanceActionWithChan invokes the polardb.ModifyPendingMaintenanceAction API asynchronously
func (client *Client) ModifyPendingMaintenanceActionWithChan(request *ModifyPendingMaintenanceActionRequest) (<-chan *ModifyPendingMaintenanceActionResponse, <-chan error) {
	responseChan := make(chan *ModifyPendingMaintenanceActionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyPendingMaintenanceAction(request)
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

// ModifyPendingMaintenanceActionWithCallback invokes the polardb.ModifyPendingMaintenanceAction API asynchronously
func (client *Client) ModifyPendingMaintenanceActionWithCallback(request *ModifyPendingMaintenanceActionRequest, callback func(response *ModifyPendingMaintenanceActionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyPendingMaintenanceActionResponse
		var err error
		defer close(result)
		response, err = client.ModifyPendingMaintenanceAction(request)
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

// ModifyPendingMaintenanceActionRequest is the request struct for api ModifyPendingMaintenanceAction
type ModifyPendingMaintenanceActionRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceGroupId      string           `position:"Query" name:"ResourceGroupId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	SwitchTime           string           `position:"Query" name:"SwitchTime"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Ids                  string           `position:"Query" name:"Ids"`
}

// ModifyPendingMaintenanceActionResponse is the response struct for api ModifyPendingMaintenanceAction
type ModifyPendingMaintenanceActionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Ids       string `json:"Ids" xml:"Ids"`
}

// CreateModifyPendingMaintenanceActionRequest creates a request to invoke ModifyPendingMaintenanceAction API
func CreateModifyPendingMaintenanceActionRequest() (request *ModifyPendingMaintenanceActionRequest) {
	request = &ModifyPendingMaintenanceActionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("polardb", "2017-08-01", "ModifyPendingMaintenanceAction", "", "")
	request.Method = requests.POST
	return
}

// CreateModifyPendingMaintenanceActionResponse creates a response to parse from ModifyPendingMaintenanceAction response
func CreateModifyPendingMaintenanceActionResponse() (response *ModifyPendingMaintenanceActionResponse) {
	response = &ModifyPendingMaintenanceActionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
