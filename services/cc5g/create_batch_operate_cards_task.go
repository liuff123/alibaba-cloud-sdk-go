package cc5g

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

// CreateBatchOperateCardsTask invokes the cc5g.CreateBatchOperateCardsTask API synchronously
func (client *Client) CreateBatchOperateCardsTask(request *CreateBatchOperateCardsTaskRequest) (response *CreateBatchOperateCardsTaskResponse, err error) {
	response = CreateCreateBatchOperateCardsTaskResponse()
	err = client.DoAction(request, response)
	return
}

// CreateBatchOperateCardsTaskWithChan invokes the cc5g.CreateBatchOperateCardsTask API asynchronously
func (client *Client) CreateBatchOperateCardsTaskWithChan(request *CreateBatchOperateCardsTaskRequest) (<-chan *CreateBatchOperateCardsTaskResponse, <-chan error) {
	responseChan := make(chan *CreateBatchOperateCardsTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateBatchOperateCardsTask(request)
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

// CreateBatchOperateCardsTaskWithCallback invokes the cc5g.CreateBatchOperateCardsTask API asynchronously
func (client *Client) CreateBatchOperateCardsTaskWithCallback(request *CreateBatchOperateCardsTaskRequest, callback func(response *CreateBatchOperateCardsTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateBatchOperateCardsTaskResponse
		var err error
		defer close(result)
		response, err = client.CreateBatchOperateCardsTask(request)
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

// CreateBatchOperateCardsTaskRequest is the request struct for api CreateBatchOperateCardsTask
type CreateBatchOperateCardsTaskRequest struct {
	*requests.RpcRequest
	Iccids                    *[]string        `position:"Query" name:"Iccids"  type:"Repeated"`
	ClientToken               string           `position:"Query" name:"ClientToken"`
	Description               string           `position:"Query" name:"Description"`
	Threshold                 requests.Integer `position:"Query" name:"Threshold"`
	EffectType                string           `position:"Query" name:"EffectType"`
	WirelessCloudConnectorIds *[]string        `position:"Query" name:"WirelessCloudConnectorIds"  type:"Repeated"`
	DryRun                    requests.Boolean `position:"Query" name:"DryRun"`
	Name                      string           `position:"Query" name:"Name"`
	OperateType               string           `position:"Query" name:"OperateType"`
	IccidsOssFilePath         string           `position:"Query" name:"IccidsOssFilePath"`
}

// CreateBatchOperateCardsTaskResponse is the response struct for api CreateBatchOperateCardsTask
type CreateBatchOperateCardsTaskResponse struct {
	*responses.BaseResponse
	RequestId                string `json:"RequestId" xml:"RequestId"`
	BatchOperateCardsTaskId  string `json:"BatchOperateCardsTaskId" xml:"BatchOperateCardsTaskId"`
	OperateResultOssFilePath string `json:"OperateResultOssFilePath" xml:"OperateResultOssFilePath"`
}

// CreateCreateBatchOperateCardsTaskRequest creates a request to invoke CreateBatchOperateCardsTask API
func CreateCreateBatchOperateCardsTaskRequest() (request *CreateBatchOperateCardsTaskRequest) {
	request = &CreateBatchOperateCardsTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CC5G", "2022-03-14", "CreateBatchOperateCardsTask", "fivegcc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateBatchOperateCardsTaskResponse creates a response to parse from CreateBatchOperateCardsTask response
func CreateCreateBatchOperateCardsTaskResponse() (response *CreateBatchOperateCardsTaskResponse) {
	response = &CreateBatchOperateCardsTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
