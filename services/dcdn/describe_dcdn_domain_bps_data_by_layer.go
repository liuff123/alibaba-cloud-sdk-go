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

// DescribeDcdnDomainBpsDataByLayer invokes the dcdn.DescribeDcdnDomainBpsDataByLayer API synchronously
func (client *Client) DescribeDcdnDomainBpsDataByLayer(request *DescribeDcdnDomainBpsDataByLayerRequest) (response *DescribeDcdnDomainBpsDataByLayerResponse, err error) {
	response = CreateDescribeDcdnDomainBpsDataByLayerResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDcdnDomainBpsDataByLayerWithChan invokes the dcdn.DescribeDcdnDomainBpsDataByLayer API asynchronously
func (client *Client) DescribeDcdnDomainBpsDataByLayerWithChan(request *DescribeDcdnDomainBpsDataByLayerRequest) (<-chan *DescribeDcdnDomainBpsDataByLayerResponse, <-chan error) {
	responseChan := make(chan *DescribeDcdnDomainBpsDataByLayerResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDcdnDomainBpsDataByLayer(request)
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

// DescribeDcdnDomainBpsDataByLayerWithCallback invokes the dcdn.DescribeDcdnDomainBpsDataByLayer API asynchronously
func (client *Client) DescribeDcdnDomainBpsDataByLayerWithCallback(request *DescribeDcdnDomainBpsDataByLayerRequest, callback func(response *DescribeDcdnDomainBpsDataByLayerResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDcdnDomainBpsDataByLayerResponse
		var err error
		defer close(result)
		response, err = client.DescribeDcdnDomainBpsDataByLayer(request)
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

// DescribeDcdnDomainBpsDataByLayerRequest is the request struct for api DescribeDcdnDomainBpsDataByLayer
type DescribeDcdnDomainBpsDataByLayerRequest struct {
	*requests.RpcRequest
	LocationNameEn string           `position:"Query" name:"LocationNameEn"`
	StartTime      string           `position:"Query" name:"StartTime"`
	IspNameEn      string           `position:"Query" name:"IspNameEn"`
	Layer          string           `position:"Query" name:"Layer"`
	DomainName     string           `position:"Query" name:"DomainName"`
	EndTime        string           `position:"Query" name:"EndTime"`
	OwnerId        requests.Integer `position:"Query" name:"OwnerId"`
	Interval       string           `position:"Query" name:"Interval"`
}

// DescribeDcdnDomainBpsDataByLayerResponse is the response struct for api DescribeDcdnDomainBpsDataByLayer
type DescribeDcdnDomainBpsDataByLayerResponse struct {
	*responses.BaseResponse
	DataInterval    string          `json:"DataInterval" xml:"DataInterval"`
	RequestId       string          `json:"RequestId" xml:"RequestId"`
	BpsDataInterval BpsDataInterval `json:"BpsDataInterval" xml:"BpsDataInterval"`
}

// CreateDescribeDcdnDomainBpsDataByLayerRequest creates a request to invoke DescribeDcdnDomainBpsDataByLayer API
func CreateDescribeDcdnDomainBpsDataByLayerRequest() (request *DescribeDcdnDomainBpsDataByLayerRequest) {
	request = &DescribeDcdnDomainBpsDataByLayerRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "DescribeDcdnDomainBpsDataByLayer", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeDcdnDomainBpsDataByLayerResponse creates a response to parse from DescribeDcdnDomainBpsDataByLayer response
func CreateDescribeDcdnDomainBpsDataByLayerResponse() (response *DescribeDcdnDomainBpsDataByLayerResponse) {
	response = &DescribeDcdnDomainBpsDataByLayerResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
