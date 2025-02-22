package config

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

// CreateCompliancePack invokes the config.CreateCompliancePack API synchronously
func (client *Client) CreateCompliancePack(request *CreateCompliancePackRequest) (response *CreateCompliancePackResponse, err error) {
	response = CreateCreateCompliancePackResponse()
	err = client.DoAction(request, response)
	return
}

// CreateCompliancePackWithChan invokes the config.CreateCompliancePack API asynchronously
func (client *Client) CreateCompliancePackWithChan(request *CreateCompliancePackRequest) (<-chan *CreateCompliancePackResponse, <-chan error) {
	responseChan := make(chan *CreateCompliancePackResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateCompliancePack(request)
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

// CreateCompliancePackWithCallback invokes the config.CreateCompliancePack API asynchronously
func (client *Client) CreateCompliancePackWithCallback(request *CreateCompliancePackRequest, callback func(response *CreateCompliancePackResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateCompliancePackResponse
		var err error
		defer close(result)
		response, err = client.CreateCompliancePack(request)
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

// CreateCompliancePackRequest is the request struct for api CreateCompliancePack
type CreateCompliancePackRequest struct {
	*requests.RpcRequest
	CompliancePackName       string                             `position:"Body" name:"CompliancePackName"`
	ClientToken              string                             `position:"Body" name:"ClientToken"`
	CompliancePackTemplateId string                             `position:"Body" name:"CompliancePackTemplateId"`
	Description              string                             `position:"Body" name:"Description"`
	DefaultEnable            requests.Boolean                   `position:"Body" name:"DefaultEnable"`
	ConfigRules              *[]CreateCompliancePackConfigRules `position:"Body" name:"ConfigRules"  type:"Json"`
	RiskLevel                requests.Integer                   `position:"Body" name:"RiskLevel"`
}

// CreateCompliancePackConfigRules is a repeated param struct in CreateCompliancePackRequest
type CreateCompliancePackConfigRules struct {
	ManagedRuleIdentifier string                                                     `name:"ManagedRuleIdentifier"`
	ConfigRuleParameters  *[]CreateCompliancePackConfigRulesConfigRuleParametersItem `name:"ConfigRuleParameters" type:"Repeated"`
	ConfigRuleId          string                                                     `name:"ConfigRuleId"`
	ConfigRuleName        string                                                     `name:"ConfigRuleName"`
	Description           string                                                     `name:"Description"`
	RiskLevel             string                                                     `name:"RiskLevel"`
}

// CreateCompliancePackConfigRulesConfigRuleParametersItem is a repeated param struct in CreateCompliancePackRequest
type CreateCompliancePackConfigRulesConfigRuleParametersItem struct {
	ParameterValue string `name:"ParameterValue"`
	ParameterName  string `name:"ParameterName"`
}

// CreateCompliancePackResponse is the response struct for api CreateCompliancePack
type CreateCompliancePackResponse struct {
	*responses.BaseResponse
	CompliancePackId string `json:"CompliancePackId" xml:"CompliancePackId"`
	RequestId        string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateCompliancePackRequest creates a request to invoke CreateCompliancePack API
func CreateCreateCompliancePackRequest() (request *CreateCompliancePackRequest) {
	request = &CreateCompliancePackRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Config", "2020-09-07", "CreateCompliancePack", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateCompliancePackResponse creates a response to parse from CreateCompliancePack response
func CreateCreateCompliancePackResponse() (response *CreateCompliancePackResponse) {
	response = &CreateCompliancePackResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
