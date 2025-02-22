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

// DescribeDomainsUsageByDay invokes the cdn.DescribeDomainsUsageByDay API synchronously
func (client *Client) DescribeDomainsUsageByDay(request *DescribeDomainsUsageByDayRequest) (response *DescribeDomainsUsageByDayResponse, err error) {
	response = CreateDescribeDomainsUsageByDayResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDomainsUsageByDayWithChan invokes the cdn.DescribeDomainsUsageByDay API asynchronously
func (client *Client) DescribeDomainsUsageByDayWithChan(request *DescribeDomainsUsageByDayRequest) (<-chan *DescribeDomainsUsageByDayResponse, <-chan error) {
	responseChan := make(chan *DescribeDomainsUsageByDayResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDomainsUsageByDay(request)
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

// DescribeDomainsUsageByDayWithCallback invokes the cdn.DescribeDomainsUsageByDay API asynchronously
func (client *Client) DescribeDomainsUsageByDayWithCallback(request *DescribeDomainsUsageByDayRequest, callback func(response *DescribeDomainsUsageByDayResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDomainsUsageByDayResponse
		var err error
		defer close(result)
		response, err = client.DescribeDomainsUsageByDay(request)
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

// DescribeDomainsUsageByDayRequest is the request struct for api DescribeDomainsUsageByDay
type DescribeDomainsUsageByDayRequest struct {
	*requests.RpcRequest
	StartTime  string           `position:"Query" name:"StartTime"`
	DomainName string           `position:"Query" name:"DomainName"`
	EndTime    string           `position:"Query" name:"EndTime"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeDomainsUsageByDayResponse is the response struct for api DescribeDomainsUsageByDay
type DescribeDomainsUsageByDayResponse struct {
	*responses.BaseResponse
	EndTime      string      `json:"EndTime" xml:"EndTime"`
	StartTime    string      `json:"StartTime" xml:"StartTime"`
	RequestId    string      `json:"RequestId" xml:"RequestId"`
	DomainName   string      `json:"DomainName" xml:"DomainName"`
	DataInterval string      `json:"DataInterval" xml:"DataInterval"`
	UsageTotal   UsageTotal  `json:"UsageTotal" xml:"UsageTotal"`
	UsageByDays  UsageByDays `json:"UsageByDays" xml:"UsageByDays"`
}

// CreateDescribeDomainsUsageByDayRequest creates a request to invoke DescribeDomainsUsageByDay API
func CreateDescribeDomainsUsageByDayRequest() (request *DescribeDomainsUsageByDayRequest) {
	request = &DescribeDomainsUsageByDayRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "DescribeDomainsUsageByDay", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeDomainsUsageByDayResponse creates a response to parse from DescribeDomainsUsageByDay response
func CreateDescribeDomainsUsageByDayResponse() (response *DescribeDomainsUsageByDayResponse) {
	response = &DescribeDomainsUsageByDayResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
