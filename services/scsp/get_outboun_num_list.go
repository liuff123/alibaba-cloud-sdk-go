package scsp

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

// GetOutbounNumList invokes the scsp.GetOutbounNumList API synchronously
func (client *Client) GetOutbounNumList(request *GetOutbounNumListRequest) (response *GetOutbounNumListResponse, err error) {
	response = CreateGetOutbounNumListResponse()
	err = client.DoAction(request, response)
	return
}

// GetOutbounNumListWithChan invokes the scsp.GetOutbounNumList API asynchronously
func (client *Client) GetOutbounNumListWithChan(request *GetOutbounNumListRequest) (<-chan *GetOutbounNumListResponse, <-chan error) {
	responseChan := make(chan *GetOutbounNumListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetOutbounNumList(request)
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

// GetOutbounNumListWithCallback invokes the scsp.GetOutbounNumList API asynchronously
func (client *Client) GetOutbounNumListWithCallback(request *GetOutbounNumListRequest, callback func(response *GetOutbounNumListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetOutbounNumListResponse
		var err error
		defer close(result)
		response, err = client.GetOutbounNumList(request)
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

// GetOutbounNumListRequest is the request struct for api GetOutbounNumList
type GetOutbounNumListRequest struct {
	*requests.RpcRequest
	ClientToken string `position:"Body"`
	InstanceId  string `position:"Body"`
	AccountName string `position:"Body"`
}

// GetOutbounNumListResponse is the response struct for api GetOutbounNumList
type GetOutbounNumListResponse struct {
	*responses.BaseResponse
	Message        string `json:"Message" xml:"Message"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Code           string `json:"Code" xml:"Code"`
	Success        bool   `json:"Success" xml:"Success"`
	HttpStatusCode int64  `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Data           Data   `json:"Data" xml:"Data"`
}

// CreateGetOutbounNumListRequest creates a request to invoke GetOutbounNumList API
func CreateGetOutbounNumListRequest() (request *GetOutbounNumListRequest) {
	request = &GetOutbounNumListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("scsp", "2020-07-02", "GetOutbounNumList", "", "")
	request.Method = requests.POST
	return
}

// CreateGetOutbounNumListResponse creates a response to parse from GetOutbounNumList response
func CreateGetOutbounNumListResponse() (response *GetOutbounNumListResponse) {
	response = &GetOutbounNumListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
