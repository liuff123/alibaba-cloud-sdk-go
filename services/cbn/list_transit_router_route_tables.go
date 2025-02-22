package cbn

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

// ListTransitRouterRouteTables invokes the cbn.ListTransitRouterRouteTables API synchronously
func (client *Client) ListTransitRouterRouteTables(request *ListTransitRouterRouteTablesRequest) (response *ListTransitRouterRouteTablesResponse, err error) {
	response = CreateListTransitRouterRouteTablesResponse()
	err = client.DoAction(request, response)
	return
}

// ListTransitRouterRouteTablesWithChan invokes the cbn.ListTransitRouterRouteTables API asynchronously
func (client *Client) ListTransitRouterRouteTablesWithChan(request *ListTransitRouterRouteTablesRequest) (<-chan *ListTransitRouterRouteTablesResponse, <-chan error) {
	responseChan := make(chan *ListTransitRouterRouteTablesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListTransitRouterRouteTables(request)
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

// ListTransitRouterRouteTablesWithCallback invokes the cbn.ListTransitRouterRouteTables API asynchronously
func (client *Client) ListTransitRouterRouteTablesWithCallback(request *ListTransitRouterRouteTablesRequest, callback func(response *ListTransitRouterRouteTablesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListTransitRouterRouteTablesResponse
		var err error
		defer close(result)
		response, err = client.ListTransitRouterRouteTables(request)
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

// ListTransitRouterRouteTablesRequest is the request struct for api ListTransitRouterRouteTables
type ListTransitRouterRouteTablesRequest struct {
	*requests.RpcRequest
	ResourceOwnerId               requests.Integer                   `position:"Query" name:"ResourceOwnerId"`
	TransitRouterRouteTableNames  *[]string                          `position:"Query" name:"TransitRouterRouteTableNames"  type:"Repeated"`
	TransitRouterRouteTableType   string                             `position:"Query" name:"TransitRouterRouteTableType"`
	TransitRouterRouteTableStatus string                             `position:"Query" name:"TransitRouterRouteTableStatus"`
	TransitRouterRouteTableIds    *[]string                          `position:"Query" name:"TransitRouterRouteTableIds"  type:"Repeated"`
	NextToken                     string                             `position:"Query" name:"NextToken"`
	Tag                           *[]ListTransitRouterRouteTablesTag `position:"Query" name:"Tag"  type:"Repeated"`
	ResourceOwnerAccount          string                             `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                  string                             `position:"Query" name:"OwnerAccount"`
	OwnerId                       requests.Integer                   `position:"Query" name:"OwnerId"`
	TransitRouterId               string                             `position:"Query" name:"TransitRouterId"`
	MaxResults                    requests.Integer                   `position:"Query" name:"MaxResults"`
}

// ListTransitRouterRouteTablesTag is a repeated param struct in ListTransitRouterRouteTablesRequest
type ListTransitRouterRouteTablesTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// ListTransitRouterRouteTablesResponse is the response struct for api ListTransitRouterRouteTables
type ListTransitRouterRouteTablesResponse struct {
	*responses.BaseResponse
	NextToken                string                    `json:"NextToken" xml:"NextToken"`
	RequestId                string                    `json:"RequestId" xml:"RequestId"`
	TotalCount               int                       `json:"TotalCount" xml:"TotalCount"`
	MaxResults               int                       `json:"MaxResults" xml:"MaxResults"`
	TransitRouterRouteTables []TransitRouterRouteTable `json:"TransitRouterRouteTables" xml:"TransitRouterRouteTables"`
}

// CreateListTransitRouterRouteTablesRequest creates a request to invoke ListTransitRouterRouteTables API
func CreateListTransitRouterRouteTablesRequest() (request *ListTransitRouterRouteTablesRequest) {
	request = &ListTransitRouterRouteTablesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cbn", "2017-09-12", "ListTransitRouterRouteTables", "cbn", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListTransitRouterRouteTablesResponse creates a response to parse from ListTransitRouterRouteTables response
func CreateListTransitRouterRouteTablesResponse() (response *ListTransitRouterRouteTablesResponse) {
	response = &ListTransitRouterRouteTablesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
