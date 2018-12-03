package ess

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

// CreateScalingConfiguration invokes the ess.CreateScalingConfiguration API synchronously
// api document: https://help.aliyun.com/api/ess/createscalingconfiguration.html
func (client *Client) CreateScalingConfiguration(request *CreateScalingConfigurationRequest) (response *CreateScalingConfigurationResponse, err error) {
	response = CreateCreateScalingConfigurationResponse()
	err = client.DoAction(request, response)
	return
}

// CreateScalingConfigurationWithChan invokes the ess.CreateScalingConfiguration API asynchronously
// api document: https://help.aliyun.com/api/ess/createscalingconfiguration.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateScalingConfigurationWithChan(request *CreateScalingConfigurationRequest) (<-chan *CreateScalingConfigurationResponse, <-chan error) {
	responseChan := make(chan *CreateScalingConfigurationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateScalingConfiguration(request)
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

// CreateScalingConfigurationWithCallback invokes the ess.CreateScalingConfiguration API asynchronously
// api document: https://help.aliyun.com/api/ess/createscalingconfiguration.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateScalingConfigurationWithCallback(request *CreateScalingConfigurationRequest, callback func(response *CreateScalingConfigurationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateScalingConfigurationResponse
		var err error
		defer close(result)
		response, err = client.CreateScalingConfiguration(request)
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

// CreateScalingConfigurationRequest is the request struct for api CreateScalingConfiguration
type CreateScalingConfigurationRequest struct {
	*requests.RpcRequest
	ImageId                     string                                      `position:"Query" name:"ImageId"`
	Memory                      requests.Integer                            `position:"Query" name:"Memory"`
	ScalingGroupId              string                                      `position:"Query" name:"ScalingGroupId"`
	InstanceTypes               *[]string                                   `position:"Query" name:"InstanceTypes"  type:"Repeated"`
	IoOptimized                 string                                      `position:"Query" name:"IoOptimized"`
	SecurityGroupId             string                                      `position:"Query" name:"SecurityGroupId"`
	InternetMaxBandwidthOut     requests.Integer                            `position:"Query" name:"InternetMaxBandwidthOut"`
	SecurityEnhancementStrategy string                                      `position:"Query" name:"SecurityEnhancementStrategy"`
	KeyPairName                 string                                      `position:"Query" name:"KeyPairName"`
	SpotPriceLimit              *[]CreateScalingConfigurationSpotPriceLimit `position:"Query" name:"SpotPriceLimit"  type:"Repeated"`
	SystemDiskCategory          string                                      `position:"Query" name:"SystemDisk.Category"`
	UserData                    string                                      `position:"Query" name:"UserData"`
	HostName                    string                                      `position:"Query" name:"HostName"`
	Password                    string                                      `position:"Query" name:"Password"`
	PasswordInherit             requests.Boolean                            `position:"Query" name:"PasswordInherit"`
	ImageName                   string                                      `position:"Query" name:"ImageName"`
	InstanceType                string                                      `position:"Query" name:"InstanceType"`
	DeploymentSetId             string                                      `position:"Query" name:"DeploymentSetId"`
	ResourceOwnerAccount        string                                      `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                string                                      `position:"Query" name:"OwnerAccount"`
	Cpu                         requests.Integer                            `position:"Query" name:"Cpu"`
	RamRoleName                 string                                      `position:"Query" name:"RamRoleName"`
	OwnerId                     requests.Integer                            `position:"Query" name:"OwnerId"`
	DataDisk                    *[]CreateScalingConfigurationDataDisk       `position:"Query" name:"DataDisk"  type:"Repeated"`
	ScalingConfigurationName    string                                      `position:"Query" name:"ScalingConfigurationName"`
	Tags                        string                                      `position:"Query" name:"Tags"`
	SpotStrategy                string                                      `position:"Query" name:"SpotStrategy"`
	LoadBalancerWeight          requests.Integer                            `position:"Query" name:"LoadBalancerWeight"`
	InstanceName                string                                      `position:"Query" name:"InstanceName"`
	SystemDiskSize              requests.Integer                            `position:"Query" name:"SystemDisk.Size"`
	InternetChargeType          string                                      `position:"Query" name:"InternetChargeType"`
	InternetMaxBandwidthIn      requests.Integer                            `position:"Query" name:"InternetMaxBandwidthIn"`
}

// CreateScalingConfigurationSpotPriceLimit is a repeated param struct in CreateScalingConfigurationRequest
type CreateScalingConfigurationSpotPriceLimit struct {
	InstanceType string `name:"InstanceType"`
	PriceLimit   string `name:"PriceLimit"`
}

// CreateScalingConfigurationDataDisk is a repeated param struct in CreateScalingConfigurationRequest
type CreateScalingConfigurationDataDisk struct {
	SnapshotId         string `name:"SnapshotId"`
	Size               string `name:"Size"`
	Category           string `name:"Category"`
	Device             string `name:"Device"`
	DeleteWithInstance string `name:"DeleteWithInstance"`
}

// CreateScalingConfigurationResponse is the response struct for api CreateScalingConfiguration
type CreateScalingConfigurationResponse struct {
	*responses.BaseResponse
	ScalingConfigurationId string `json:"ScalingConfigurationId" xml:"ScalingConfigurationId"`
	RequestId              string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateScalingConfigurationRequest creates a request to invoke CreateScalingConfiguration API
func CreateCreateScalingConfigurationRequest() (request *CreateScalingConfigurationRequest) {
	request = &CreateScalingConfigurationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ess", "2014-08-28", "CreateScalingConfiguration", "", "")
	return
}

// CreateCreateScalingConfigurationResponse creates a response to parse from CreateScalingConfiguration response
func CreateCreateScalingConfigurationResponse() (response *CreateScalingConfigurationResponse) {
	response = &CreateScalingConfigurationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
