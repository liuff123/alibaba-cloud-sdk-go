package ocr

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

// TrimDocument invokes the ocr.TrimDocument API synchronously
func (client *Client) TrimDocument(request *TrimDocumentRequest) (response *TrimDocumentResponse, err error) {
	response = CreateTrimDocumentResponse()
	err = client.DoAction(request, response)
	return
}

// TrimDocumentWithChan invokes the ocr.TrimDocument API asynchronously
func (client *Client) TrimDocumentWithChan(request *TrimDocumentRequest) (<-chan *TrimDocumentResponse, <-chan error) {
	responseChan := make(chan *TrimDocumentResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.TrimDocument(request)
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

// TrimDocumentWithCallback invokes the ocr.TrimDocument API asynchronously
func (client *Client) TrimDocumentWithCallback(request *TrimDocumentRequest, callback func(response *TrimDocumentResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *TrimDocumentResponse
		var err error
		defer close(result)
		response, err = client.TrimDocument(request)
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

// TrimDocumentRequest is the request struct for api TrimDocument
type TrimDocumentRequest struct {
	*requests.RpcRequest
	FileType   string           `position:"Body" name:"FileType"`
	OutputType string           `position:"Body" name:"OutputType"`
	Async      requests.Boolean `position:"Body" name:"Async"`
	FileURL    string           `position:"Body" name:"FileURL"`
}

// TrimDocumentResponse is the response struct for api TrimDocument
type TrimDocumentResponse struct {
	*responses.BaseResponse
	RequestId string             `json:"RequestId" xml:"RequestId"`
	Code      string             `json:"Code" xml:"Code"`
	Message   string             `json:"Message" xml:"Message"`
	Data      DataInTrimDocument `json:"Data" xml:"Data"`
}

// CreateTrimDocumentRequest creates a request to invoke TrimDocument API
func CreateTrimDocumentRequest() (request *TrimDocumentRequest) {
	request = &TrimDocumentRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ocr", "2019-12-30", "TrimDocument", "ocr", "openAPI")
	request.Method = requests.POST
	return
}

// CreateTrimDocumentResponse creates a response to parse from TrimDocument response
func CreateTrimDocumentResponse() (response *TrimDocumentResponse) {
	response = &TrimDocumentResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
