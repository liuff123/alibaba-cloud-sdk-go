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

// DescribeDomainBpsDataByTimeStamp invokes the cdn.DescribeDomainBpsDataByTimeStamp API synchronously
func (client *Client) DescribeDomainBpsDataByTimeStamp(request *DescribeDomainBpsDataByTimeStampRequest) (response *DescribeDomainBpsDataByTimeStampResponse, err error) {
	response = CreateDescribeDomainBpsDataByTimeStampResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDomainBpsDataByTimeStampWithChan invokes the cdn.DescribeDomainBpsDataByTimeStamp API asynchronously
func (client *Client) DescribeDomainBpsDataByTimeStampWithChan(request *DescribeDomainBpsDataByTimeStampRequest) (<-chan *DescribeDomainBpsDataByTimeStampResponse, <-chan error) {
	responseChan := make(chan *DescribeDomainBpsDataByTimeStampResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDomainBpsDataByTimeStamp(request)
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

// DescribeDomainBpsDataByTimeStampWithCallback invokes the cdn.DescribeDomainBpsDataByTimeStamp API asynchronously
func (client *Client) DescribeDomainBpsDataByTimeStampWithCallback(request *DescribeDomainBpsDataByTimeStampRequest, callback func(response *DescribeDomainBpsDataByTimeStampResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDomainBpsDataByTimeStampResponse
		var err error
		defer close(result)
		response, err = client.DescribeDomainBpsDataByTimeStamp(request)
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

// DescribeDomainBpsDataByTimeStampRequest is the request struct for api DescribeDomainBpsDataByTimeStamp
type DescribeDomainBpsDataByTimeStampRequest struct {
	*requests.RpcRequest
	LocationNames string           `position:"Query" name:"LocationNames"`
	IspNames      string           `position:"Query" name:"IspNames"`
	DomainName    string           `position:"Query" name:"DomainName"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	TimePoint     string           `position:"Query" name:"TimePoint"`
}

// DescribeDomainBpsDataByTimeStampResponse is the response struct for api DescribeDomainBpsDataByTimeStamp
type DescribeDomainBpsDataByTimeStampResponse struct {
	*responses.BaseResponse
	TimeStamp   string      `json:"TimeStamp" xml:"TimeStamp"`
	RequestId   string      `json:"RequestId" xml:"RequestId"`
	DomainName  string      `json:"DomainName" xml:"DomainName"`
	BpsDataList BpsDataList `json:"BpsDataList" xml:"BpsDataList"`
}

// CreateDescribeDomainBpsDataByTimeStampRequest creates a request to invoke DescribeDomainBpsDataByTimeStamp API
func CreateDescribeDomainBpsDataByTimeStampRequest() (request *DescribeDomainBpsDataByTimeStampRequest) {
	request = &DescribeDomainBpsDataByTimeStampRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "DescribeDomainBpsDataByTimeStamp", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeDomainBpsDataByTimeStampResponse creates a response to parse from DescribeDomainBpsDataByTimeStamp response
func CreateDescribeDomainBpsDataByTimeStampResponse() (response *DescribeDomainBpsDataByTimeStampResponse) {
	response = &DescribeDomainBpsDataByTimeStampResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
