package opensearch

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

// ListAppGroupErrors invokes the opensearch.ListAppGroupErrors API synchronously
func (client *Client) ListAppGroupErrors(request *ListAppGroupErrorsRequest) (response *ListAppGroupErrorsResponse, err error) {
	response = CreateListAppGroupErrorsResponse()
	err = client.DoAction(request, response)
	return
}

// ListAppGroupErrorsWithChan invokes the opensearch.ListAppGroupErrors API asynchronously
func (client *Client) ListAppGroupErrorsWithChan(request *ListAppGroupErrorsRequest) (<-chan *ListAppGroupErrorsResponse, <-chan error) {
	responseChan := make(chan *ListAppGroupErrorsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListAppGroupErrors(request)
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

// ListAppGroupErrorsWithCallback invokes the opensearch.ListAppGroupErrors API asynchronously
func (client *Client) ListAppGroupErrorsWithCallback(request *ListAppGroupErrorsRequest, callback func(response *ListAppGroupErrorsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListAppGroupErrorsResponse
		var err error
		defer close(result)
		response, err = client.ListAppGroupErrors(request)
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

// ListAppGroupErrorsRequest is the request struct for api ListAppGroupErrors
type ListAppGroupErrorsRequest struct {
	*requests.RoaRequest
	AppId            string           `position:"Query" name:"appId"`
	PageSize         requests.Integer `position:"Query" name:"pageSize"`
	StartTime        requests.Integer `position:"Query" name:"startTime"`
	StopTime         requests.Integer `position:"Query" name:"stopTime"`
	AppGroupIdentity string           `position:"Path" name:"appGroupIdentity"`
	PageNumber       requests.Integer `position:"Query" name:"pageNumber"`
}

// ListAppGroupErrorsResponse is the response struct for api ListAppGroupErrors
type ListAppGroupErrorsResponse struct {
	*responses.BaseResponse
	TotalCount int64                    `json:"totalCount" xml:"totalCount"`
	RequestId  string                   `json:"requestId" xml:"requestId"`
	Result     []map[string]interface{} `json:"result" xml:"result"`
}

// CreateListAppGroupErrorsRequest creates a request to invoke ListAppGroupErrors API
func CreateListAppGroupErrorsRequest() (request *ListAppGroupErrorsRequest) {
	request = &ListAppGroupErrorsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("OpenSearch", "2017-12-25", "ListAppGroupErrors", "/v4/openapi/app-groups/[appGroupIdentity]/errors", "", "")
	request.Method = requests.GET
	return
}

// CreateListAppGroupErrorsResponse creates a response to parse from ListAppGroupErrors response
func CreateListAppGroupErrorsResponse() (response *ListAppGroupErrorsResponse) {
	response = &ListAppGroupErrorsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
