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

// DescribeDcdnWafPolicyDomains invokes the dcdn.DescribeDcdnWafPolicyDomains API synchronously
func (client *Client) DescribeDcdnWafPolicyDomains(request *DescribeDcdnWafPolicyDomainsRequest) (response *DescribeDcdnWafPolicyDomainsResponse, err error) {
	response = CreateDescribeDcdnWafPolicyDomainsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDcdnWafPolicyDomainsWithChan invokes the dcdn.DescribeDcdnWafPolicyDomains API asynchronously
func (client *Client) DescribeDcdnWafPolicyDomainsWithChan(request *DescribeDcdnWafPolicyDomainsRequest) (<-chan *DescribeDcdnWafPolicyDomainsResponse, <-chan error) {
	responseChan := make(chan *DescribeDcdnWafPolicyDomainsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDcdnWafPolicyDomains(request)
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

// DescribeDcdnWafPolicyDomainsWithCallback invokes the dcdn.DescribeDcdnWafPolicyDomains API asynchronously
func (client *Client) DescribeDcdnWafPolicyDomainsWithCallback(request *DescribeDcdnWafPolicyDomainsRequest, callback func(response *DescribeDcdnWafPolicyDomainsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDcdnWafPolicyDomainsResponse
		var err error
		defer close(result)
		response, err = client.DescribeDcdnWafPolicyDomains(request)
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

// DescribeDcdnWafPolicyDomainsRequest is the request struct for api DescribeDcdnWafPolicyDomains
type DescribeDcdnWafPolicyDomainsRequest struct {
	*requests.RpcRequest
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	PolicyId   requests.Integer `position:"Query" name:"PolicyId"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeDcdnWafPolicyDomainsResponse is the response struct for api DescribeDcdnWafPolicyDomains
type DescribeDcdnWafPolicyDomainsResponse struct {
	*responses.BaseResponse
	PageSize   int          `json:"PageSize" xml:"PageSize"`
	RequestId  string       `json:"RequestId" xml:"RequestId"`
	PageNumber int          `json:"PageNumber" xml:"PageNumber"`
	TotalCount int          `json:"TotalCount" xml:"TotalCount"`
	Domains    []DomainItem `json:"Domains" xml:"Domains"`
}

// CreateDescribeDcdnWafPolicyDomainsRequest creates a request to invoke DescribeDcdnWafPolicyDomains API
func CreateDescribeDcdnWafPolicyDomainsRequest() (request *DescribeDcdnWafPolicyDomainsRequest) {
	request = &DescribeDcdnWafPolicyDomainsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "DescribeDcdnWafPolicyDomains", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeDcdnWafPolicyDomainsResponse creates a response to parse from DescribeDcdnWafPolicyDomains response
func CreateDescribeDcdnWafPolicyDomainsResponse() (response *DescribeDcdnWafPolicyDomainsResponse) {
	response = &DescribeDcdnWafPolicyDomainsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
