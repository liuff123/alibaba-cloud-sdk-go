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

// DescribeDcdnDomainRegionData invokes the dcdn.DescribeDcdnDomainRegionData API synchronously
func (client *Client) DescribeDcdnDomainRegionData(request *DescribeDcdnDomainRegionDataRequest) (response *DescribeDcdnDomainRegionDataResponse, err error) {
	response = CreateDescribeDcdnDomainRegionDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDcdnDomainRegionDataWithChan invokes the dcdn.DescribeDcdnDomainRegionData API asynchronously
func (client *Client) DescribeDcdnDomainRegionDataWithChan(request *DescribeDcdnDomainRegionDataRequest) (<-chan *DescribeDcdnDomainRegionDataResponse, <-chan error) {
	responseChan := make(chan *DescribeDcdnDomainRegionDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDcdnDomainRegionData(request)
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

// DescribeDcdnDomainRegionDataWithCallback invokes the dcdn.DescribeDcdnDomainRegionData API asynchronously
func (client *Client) DescribeDcdnDomainRegionDataWithCallback(request *DescribeDcdnDomainRegionDataRequest, callback func(response *DescribeDcdnDomainRegionDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDcdnDomainRegionDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeDcdnDomainRegionData(request)
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

// DescribeDcdnDomainRegionDataRequest is the request struct for api DescribeDcdnDomainRegionData
type DescribeDcdnDomainRegionDataRequest struct {
	*requests.RpcRequest
	StartTime  string           `position:"Query" name:"StartTime"`
	DomainName string           `position:"Query" name:"DomainName"`
	EndTime    string           `position:"Query" name:"EndTime"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeDcdnDomainRegionDataResponse is the response struct for api DescribeDcdnDomainRegionData
type DescribeDcdnDomainRegionDataResponse struct {
	*responses.BaseResponse
	EndTime      string                              `json:"EndTime" xml:"EndTime"`
	StartTime    string                              `json:"StartTime" xml:"StartTime"`
	RequestId    string                              `json:"RequestId" xml:"RequestId"`
	DomainName   string                              `json:"DomainName" xml:"DomainName"`
	DataInterval string                              `json:"DataInterval" xml:"DataInterval"`
	Value        ValueInDescribeDcdnDomainRegionData `json:"Value" xml:"Value"`
}

// CreateDescribeDcdnDomainRegionDataRequest creates a request to invoke DescribeDcdnDomainRegionData API
func CreateDescribeDcdnDomainRegionDataRequest() (request *DescribeDcdnDomainRegionDataRequest) {
	request = &DescribeDcdnDomainRegionDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "DescribeDcdnDomainRegionData", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeDcdnDomainRegionDataResponse creates a response to parse from DescribeDcdnDomainRegionData response
func CreateDescribeDcdnDomainRegionDataResponse() (response *DescribeDcdnDomainRegionDataResponse) {
	response = &DescribeDcdnDomainRegionDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
