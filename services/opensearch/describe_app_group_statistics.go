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

// DescribeAppGroupStatistics invokes the opensearch.DescribeAppGroupStatistics API synchronously
func (client *Client) DescribeAppGroupStatistics(request *DescribeAppGroupStatisticsRequest) (response *DescribeAppGroupStatisticsResponse, err error) {
	response = CreateDescribeAppGroupStatisticsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAppGroupStatisticsWithChan invokes the opensearch.DescribeAppGroupStatistics API asynchronously
func (client *Client) DescribeAppGroupStatisticsWithChan(request *DescribeAppGroupStatisticsRequest) (<-chan *DescribeAppGroupStatisticsResponse, <-chan error) {
	responseChan := make(chan *DescribeAppGroupStatisticsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAppGroupStatistics(request)
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

// DescribeAppGroupStatisticsWithCallback invokes the opensearch.DescribeAppGroupStatistics API asynchronously
func (client *Client) DescribeAppGroupStatisticsWithCallback(request *DescribeAppGroupStatisticsRequest, callback func(response *DescribeAppGroupStatisticsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAppGroupStatisticsResponse
		var err error
		defer close(result)
		response, err = client.DescribeAppGroupStatistics(request)
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

// DescribeAppGroupStatisticsRequest is the request struct for api DescribeAppGroupStatistics
type DescribeAppGroupStatisticsRequest struct {
	*requests.RoaRequest
	AppGroupIdentity string `position:"Path" name:"appGroupIdentity"`
}

// DescribeAppGroupStatisticsResponse is the response struct for api DescribeAppGroupStatistics
type DescribeAppGroupStatisticsResponse struct {
	*responses.BaseResponse
	Result    map[string]interface{} `json:"result" xml:"result"`
	RequestId string                 `json:"requestId" xml:"requestId"`
}

// CreateDescribeAppGroupStatisticsRequest creates a request to invoke DescribeAppGroupStatistics API
func CreateDescribeAppGroupStatisticsRequest() (request *DescribeAppGroupStatisticsRequest) {
	request = &DescribeAppGroupStatisticsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("OpenSearch", "2017-12-25", "DescribeAppGroupStatistics", "/v4/openapi/app-groups/[appGroupIdentity]/statistics", "", "")
	request.Method = requests.GET
	return
}

// CreateDescribeAppGroupStatisticsResponse creates a response to parse from DescribeAppGroupStatistics response
func CreateDescribeAppGroupStatisticsResponse() (response *DescribeAppGroupStatisticsResponse) {
	response = &DescribeAppGroupStatisticsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
