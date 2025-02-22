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

// DeleteCenInterRegionTrafficQosQueue invokes the cbn.DeleteCenInterRegionTrafficQosQueue API synchronously
func (client *Client) DeleteCenInterRegionTrafficQosQueue(request *DeleteCenInterRegionTrafficQosQueueRequest) (response *DeleteCenInterRegionTrafficQosQueueResponse, err error) {
	response = CreateDeleteCenInterRegionTrafficQosQueueResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteCenInterRegionTrafficQosQueueWithChan invokes the cbn.DeleteCenInterRegionTrafficQosQueue API asynchronously
func (client *Client) DeleteCenInterRegionTrafficQosQueueWithChan(request *DeleteCenInterRegionTrafficQosQueueRequest) (<-chan *DeleteCenInterRegionTrafficQosQueueResponse, <-chan error) {
	responseChan := make(chan *DeleteCenInterRegionTrafficQosQueueResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteCenInterRegionTrafficQosQueue(request)
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

// DeleteCenInterRegionTrafficQosQueueWithCallback invokes the cbn.DeleteCenInterRegionTrafficQosQueue API asynchronously
func (client *Client) DeleteCenInterRegionTrafficQosQueueWithCallback(request *DeleteCenInterRegionTrafficQosQueueRequest, callback func(response *DeleteCenInterRegionTrafficQosQueueResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteCenInterRegionTrafficQosQueueResponse
		var err error
		defer close(result)
		response, err = client.DeleteCenInterRegionTrafficQosQueue(request)
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

// DeleteCenInterRegionTrafficQosQueueRequest is the request struct for api DeleteCenInterRegionTrafficQosQueue
type DeleteCenInterRegionTrafficQosQueueRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	DryRun               requests.Boolean `position:"Query" name:"DryRun"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	QosQueueId           string           `position:"Query" name:"QosQueueId"`
}

// DeleteCenInterRegionTrafficQosQueueResponse is the response struct for api DeleteCenInterRegionTrafficQosQueue
type DeleteCenInterRegionTrafficQosQueueResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteCenInterRegionTrafficQosQueueRequest creates a request to invoke DeleteCenInterRegionTrafficQosQueue API
func CreateDeleteCenInterRegionTrafficQosQueueRequest() (request *DeleteCenInterRegionTrafficQosQueueRequest) {
	request = &DeleteCenInterRegionTrafficQosQueueRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cbn", "2017-09-12", "DeleteCenInterRegionTrafficQosQueue", "cbn", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteCenInterRegionTrafficQosQueueResponse creates a response to parse from DeleteCenInterRegionTrafficQosQueue response
func CreateDeleteCenInterRegionTrafficQosQueueResponse() (response *DeleteCenInterRegionTrafficQosQueueResponse) {
	response = &DeleteCenInterRegionTrafficQosQueueResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
