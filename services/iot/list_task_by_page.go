package iot

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

// ListTaskByPage invokes the iot.ListTaskByPage API synchronously
func (client *Client) ListTaskByPage(request *ListTaskByPageRequest) (response *ListTaskByPageResponse, err error) {
	response = CreateListTaskByPageResponse()
	err = client.DoAction(request, response)
	return
}

// ListTaskByPageWithChan invokes the iot.ListTaskByPage API asynchronously
func (client *Client) ListTaskByPageWithChan(request *ListTaskByPageRequest) (<-chan *ListTaskByPageResponse, <-chan error) {
	responseChan := make(chan *ListTaskByPageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListTaskByPage(request)
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

// ListTaskByPageWithCallback invokes the iot.ListTaskByPage API asynchronously
func (client *Client) ListTaskByPageWithCallback(request *ListTaskByPageRequest, callback func(response *ListTaskByPageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListTaskByPageResponse
		var err error
		defer close(result)
		response, err = client.ListTaskByPage(request)
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

// ListTaskByPageRequest is the request struct for api ListTaskByPage
type ListTaskByPageRequest struct {
	*requests.RpcRequest
	JobId         string `position:"Query" name:"JobId"`
	IotInstanceId string `position:"Query" name:"IotInstanceId"`
	PageSize      string `position:"Query" name:"PageSize"`
	JobName       string `position:"Query" name:"JobName"`
	PageNo        string `position:"Query" name:"PageNo"`
	ApiProduct    string `position:"Body" name:"ApiProduct"`
	ApiRevision   string `position:"Body" name:"ApiRevision"`
	DeviceName    string `position:"Query" name:"DeviceName"`
	Device        string `position:"Query" name:"Device"`
	Status        string `position:"Query" name:"Status"`
}

// ListTaskByPageResponse is the response struct for api ListTaskByPage
type ListTaskByPageResponse struct {
	*responses.BaseResponse
	RequestId    string               `json:"RequestId" xml:"RequestId"`
	Success      bool                 `json:"Success" xml:"Success"`
	Code         string               `json:"Code" xml:"Code"`
	ErrorMessage string               `json:"ErrorMessage" xml:"ErrorMessage"`
	Total        int                  `json:"Total" xml:"Total"`
	PageSize     int                  `json:"PageSize" xml:"PageSize"`
	PageCount    int                  `json:"PageCount" xml:"PageCount"`
	Page         int                  `json:"Page" xml:"Page"`
	Data         DataInListTaskByPage `json:"Data" xml:"Data"`
}

// CreateListTaskByPageRequest creates a request to invoke ListTaskByPage API
func CreateListTaskByPageRequest() (request *ListTaskByPageRequest) {
	request = &ListTaskByPageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "ListTaskByPage", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListTaskByPageResponse creates a response to parse from ListTaskByPage response
func CreateListTaskByPageResponse() (response *ListTaskByPageResponse) {
	response = &ListTaskByPageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
