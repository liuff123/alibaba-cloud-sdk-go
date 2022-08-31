package dcdn

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

// DescribeDcdnsecService invokes the dcdn.DescribeDcdnsecService API synchronously
func (client *Client) DescribeDcdnsecService(request *DescribeDcdnsecServiceRequest) (response *DescribeDcdnsecServiceResponse, err error) {
	response = CreateDescribeDcdnsecServiceResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDcdnsecServiceWithChan invokes the dcdn.DescribeDcdnsecService API asynchronously
func (client *Client) DescribeDcdnsecServiceWithChan(request *DescribeDcdnsecServiceRequest) (<-chan *DescribeDcdnsecServiceResponse, <-chan error) {
	responseChan := make(chan *DescribeDcdnsecServiceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDcdnsecService(request)
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

// DescribeDcdnsecServiceWithCallback invokes the dcdn.DescribeDcdnsecService API asynchronously
func (client *Client) DescribeDcdnsecServiceWithCallback(request *DescribeDcdnsecServiceRequest, callback func(response *DescribeDcdnsecServiceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDcdnsecServiceResponse
		var err error
		defer close(result)
		response, err = client.DescribeDcdnsecService(request)
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

// DescribeDcdnsecServiceRequest is the request struct for api DescribeDcdnsecService
type DescribeDcdnsecServiceRequest struct {
	*requests.RpcRequest
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
}

// DescribeDcdnsecServiceResponse is the response struct for api DescribeDcdnsecService
type DescribeDcdnsecServiceResponse struct {
	*responses.BaseResponse
	EndTime            string                                 `json:"EndTime" xml:"EndTime"`
	StartTime          string                                 `json:"StartTime" xml:"StartTime"`
	ChangingAffectTime string                                 `json:"ChangingAffectTime" xml:"ChangingAffectTime"`
	DomainNum          string                                 `json:"DomainNum" xml:"DomainNum"`
	ChangingChargeType string                                 `json:"ChangingChargeType" xml:"ChangingChargeType"`
	RequestId          string                                 `json:"RequestId" xml:"RequestId"`
	Version            string                                 `json:"Version" xml:"Version"`
	RequestType        string                                 `json:"RequestType" xml:"RequestType"`
	FlowType           string                                 `json:"FlowType" xml:"FlowType"`
	InternetChargeType string                                 `json:"InternetChargeType" xml:"InternetChargeType"`
	InstanceId         string                                 `json:"InstanceId" xml:"InstanceId"`
	OperationLocks     OperationLocksInDescribeDcdnsecService `json:"OperationLocks" xml:"OperationLocks"`
}

// CreateDescribeDcdnsecServiceRequest creates a request to invoke DescribeDcdnsecService API
func CreateDescribeDcdnsecServiceRequest() (request *DescribeDcdnsecServiceRequest) {
	request = &DescribeDcdnsecServiceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "DescribeDcdnsecService", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeDcdnsecServiceResponse creates a response to parse from DescribeDcdnsecService response
func CreateDescribeDcdnsecServiceResponse() (response *DescribeDcdnsecServiceResponse) {
	response = &DescribeDcdnsecServiceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
