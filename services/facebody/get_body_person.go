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

// GetBodyPerson invokes the facebody.GetBodyPerson API synchronously
func (client *Client) GetBodyPerson(request *GetBodyPersonRequest) (response *GetBodyPersonResponse, err error) {
	response = CreateGetBodyPersonResponse()
	err = client.DoAction(request, response)
	return
}

// GetBodyPersonWithChan invokes the facebody.GetBodyPerson API asynchronously
func (client *Client) GetBodyPersonWithChan(request *GetBodyPersonRequest) (<-chan *GetBodyPersonResponse, <-chan error) {
	responseChan := make(chan *GetBodyPersonResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetBodyPerson(request)
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

// GetBodyPersonWithCallback invokes the facebody.GetBodyPerson API asynchronously
func (client *Client) GetBodyPersonWithCallback(request *GetBodyPersonRequest, callback func(response *GetBodyPersonResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetBodyPersonResponse
		var err error
		defer close(result)
		response, err = client.GetBodyPerson(request)
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

// GetBodyPersonRequest is the request struct for api GetBodyPerson
type GetBodyPersonRequest struct {
	*requests.RpcRequest
	PersonId requests.Integer `position:"Query" name:"PersonId"`
	DbId     requests.Integer `position:"Query" name:"DbId"`
}

// GetBodyPersonResponse is the response struct for api GetBodyPerson
type GetBodyPersonResponse struct {
	*responses.BaseResponse
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateGetBodyPersonRequest creates a request to invoke GetBodyPerson API
func CreateGetBodyPersonRequest() (request *GetBodyPersonRequest) {
	request = &GetBodyPersonRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("facebody", "2019-12-30", "GetBodyPerson", "", "")
	request.Method = requests.GET
	return
}

// CreateGetBodyPersonResponse creates a response to parse from GetBodyPerson response
func CreateGetBodyPersonResponse() (response *GetBodyPersonResponse) {
	response = &GetBodyPersonResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
