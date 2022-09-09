package polardb

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

// SwitchOverGlobalDatabaseNetwork invokes the polardb.SwitchOverGlobalDatabaseNetwork API synchronously
func (client *Client) SwitchOverGlobalDatabaseNetwork(request *SwitchOverGlobalDatabaseNetworkRequest) (response *SwitchOverGlobalDatabaseNetworkResponse, err error) {
	response = CreateSwitchOverGlobalDatabaseNetworkResponse()
	err = client.DoAction(request, response)
	return
}

// SwitchOverGlobalDatabaseNetworkWithChan invokes the polardb.SwitchOverGlobalDatabaseNetwork API asynchronously
func (client *Client) SwitchOverGlobalDatabaseNetworkWithChan(request *SwitchOverGlobalDatabaseNetworkRequest) (<-chan *SwitchOverGlobalDatabaseNetworkResponse, <-chan error) {
	responseChan := make(chan *SwitchOverGlobalDatabaseNetworkResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SwitchOverGlobalDatabaseNetwork(request)
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

// SwitchOverGlobalDatabaseNetworkWithCallback invokes the polardb.SwitchOverGlobalDatabaseNetwork API asynchronously
func (client *Client) SwitchOverGlobalDatabaseNetworkWithCallback(request *SwitchOverGlobalDatabaseNetworkRequest, callback func(response *SwitchOverGlobalDatabaseNetworkResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SwitchOverGlobalDatabaseNetworkResponse
		var err error
		defer close(result)
		response, err = client.SwitchOverGlobalDatabaseNetwork(request)
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

// SwitchOverGlobalDatabaseNetworkRequest is the request struct for api SwitchOverGlobalDatabaseNetwork
type SwitchOverGlobalDatabaseNetworkRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceGroupId      string           `position:"Query" name:"ResourceGroupId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	GDNId                string           `position:"Query" name:"GDNId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string           `position:"Query" name:"DBClusterId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// SwitchOverGlobalDatabaseNetworkResponse is the response struct for api SwitchOverGlobalDatabaseNetwork
type SwitchOverGlobalDatabaseNetworkResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSwitchOverGlobalDatabaseNetworkRequest creates a request to invoke SwitchOverGlobalDatabaseNetwork API
func CreateSwitchOverGlobalDatabaseNetworkRequest() (request *SwitchOverGlobalDatabaseNetworkRequest) {
	request = &SwitchOverGlobalDatabaseNetworkRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("polardb", "2017-08-01", "SwitchOverGlobalDatabaseNetwork", "", "")
	request.Method = requests.POST
	return
}

// CreateSwitchOverGlobalDatabaseNetworkResponse creates a response to parse from SwitchOverGlobalDatabaseNetwork response
func CreateSwitchOverGlobalDatabaseNetworkResponse() (response *SwitchOverGlobalDatabaseNetworkResponse) {
	response = &SwitchOverGlobalDatabaseNetworkResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
