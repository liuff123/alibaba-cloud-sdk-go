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

// DeleteUsageDetailDataExportTask invokes the cdn.DeleteUsageDetailDataExportTask API synchronously
func (client *Client) DeleteUsageDetailDataExportTask(request *DeleteUsageDetailDataExportTaskRequest) (response *DeleteUsageDetailDataExportTaskResponse, err error) {
	response = CreateDeleteUsageDetailDataExportTaskResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteUsageDetailDataExportTaskWithChan invokes the cdn.DeleteUsageDetailDataExportTask API asynchronously
func (client *Client) DeleteUsageDetailDataExportTaskWithChan(request *DeleteUsageDetailDataExportTaskRequest) (<-chan *DeleteUsageDetailDataExportTaskResponse, <-chan error) {
	responseChan := make(chan *DeleteUsageDetailDataExportTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteUsageDetailDataExportTask(request)
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

// DeleteUsageDetailDataExportTaskWithCallback invokes the cdn.DeleteUsageDetailDataExportTask API asynchronously
func (client *Client) DeleteUsageDetailDataExportTaskWithCallback(request *DeleteUsageDetailDataExportTaskRequest, callback func(response *DeleteUsageDetailDataExportTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteUsageDetailDataExportTaskResponse
		var err error
		defer close(result)
		response, err = client.DeleteUsageDetailDataExportTask(request)
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

// DeleteUsageDetailDataExportTaskRequest is the request struct for api DeleteUsageDetailDataExportTask
type DeleteUsageDetailDataExportTaskRequest struct {
	*requests.RpcRequest
	TaskId  string           `position:"Query" name:"TaskId"`
	OwnerId requests.Integer `position:"Query" name:"OwnerId"`
}

// DeleteUsageDetailDataExportTaskResponse is the response struct for api DeleteUsageDetailDataExportTask
type DeleteUsageDetailDataExportTaskResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteUsageDetailDataExportTaskRequest creates a request to invoke DeleteUsageDetailDataExportTask API
func CreateDeleteUsageDetailDataExportTaskRequest() (request *DeleteUsageDetailDataExportTaskRequest) {
	request = &DeleteUsageDetailDataExportTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "DeleteUsageDetailDataExportTask", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteUsageDetailDataExportTaskResponse creates a response to parse from DeleteUsageDetailDataExportTask response
func CreateDeleteUsageDetailDataExportTaskResponse() (response *DeleteUsageDetailDataExportTaskResponse) {
	response = &DeleteUsageDetailDataExportTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
