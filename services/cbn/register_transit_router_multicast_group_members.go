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

// RegisterTransitRouterMulticastGroupMembers invokes the cbn.RegisterTransitRouterMulticastGroupMembers API synchronously
func (client *Client) RegisterTransitRouterMulticastGroupMembers(request *RegisterTransitRouterMulticastGroupMembersRequest) (response *RegisterTransitRouterMulticastGroupMembersResponse, err error) {
	response = CreateRegisterTransitRouterMulticastGroupMembersResponse()
	err = client.DoAction(request, response)
	return
}

// RegisterTransitRouterMulticastGroupMembersWithChan invokes the cbn.RegisterTransitRouterMulticastGroupMembers API asynchronously
func (client *Client) RegisterTransitRouterMulticastGroupMembersWithChan(request *RegisterTransitRouterMulticastGroupMembersRequest) (<-chan *RegisterTransitRouterMulticastGroupMembersResponse, <-chan error) {
	responseChan := make(chan *RegisterTransitRouterMulticastGroupMembersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RegisterTransitRouterMulticastGroupMembers(request)
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

// RegisterTransitRouterMulticastGroupMembersWithCallback invokes the cbn.RegisterTransitRouterMulticastGroupMembers API asynchronously
func (client *Client) RegisterTransitRouterMulticastGroupMembersWithCallback(request *RegisterTransitRouterMulticastGroupMembersRequest, callback func(response *RegisterTransitRouterMulticastGroupMembersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RegisterTransitRouterMulticastGroupMembersResponse
		var err error
		defer close(result)
		response, err = client.RegisterTransitRouterMulticastGroupMembers(request)
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

// RegisterTransitRouterMulticastGroupMembersRequest is the request struct for api RegisterTransitRouterMulticastGroupMembers
type RegisterTransitRouterMulticastGroupMembersRequest struct {
	*requests.RpcRequest
	ResourceOwnerId                   requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken                       string           `position:"Query" name:"ClientToken"`
	NetworkInterfaceIds               *[]string        `position:"Query" name:"NetworkInterfaceIds"  type:"Repeated"`
	TransitRouterMulticastDomainId    string           `position:"Query" name:"TransitRouterMulticastDomainId"`
	ConnectPeerIds                    *[]string        `position:"Query" name:"ConnectPeerIds"  type:"Repeated"`
	GroupIpAddress                    string           `position:"Query" name:"GroupIpAddress"`
	DryRun                            requests.Boolean `position:"Query" name:"DryRun"`
	ResourceOwnerAccount              string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                      string           `position:"Query" name:"OwnerAccount"`
	PeerTransitRouterMulticastDomains *[]string        `position:"Query" name:"PeerTransitRouterMulticastDomains"  type:"Repeated"`
	OwnerId                           requests.Integer `position:"Query" name:"OwnerId"`
	VpcId                             string           `position:"Query" name:"VpcId"`
}

// RegisterTransitRouterMulticastGroupMembersResponse is the response struct for api RegisterTransitRouterMulticastGroupMembers
type RegisterTransitRouterMulticastGroupMembersResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateRegisterTransitRouterMulticastGroupMembersRequest creates a request to invoke RegisterTransitRouterMulticastGroupMembers API
func CreateRegisterTransitRouterMulticastGroupMembersRequest() (request *RegisterTransitRouterMulticastGroupMembersRequest) {
	request = &RegisterTransitRouterMulticastGroupMembersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cbn", "2017-09-12", "RegisterTransitRouterMulticastGroupMembers", "cbn", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRegisterTransitRouterMulticastGroupMembersResponse creates a response to parse from RegisterTransitRouterMulticastGroupMembers response
func CreateRegisterTransitRouterMulticastGroupMembersResponse() (response *RegisterTransitRouterMulticastGroupMembersResponse) {
	response = &RegisterTransitRouterMulticastGroupMembersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
