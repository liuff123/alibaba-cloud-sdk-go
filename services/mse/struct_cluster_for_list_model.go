package mse

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

// ClusterForListModel is a nested struct in mse response
type ClusterForListModel struct {
	EndDate          string `json:"EndDate" xml:"EndDate"`
	IntranetDomain   string `json:"IntranetDomain" xml:"IntranetDomain"`
	InternetDomain   string `json:"InternetDomain" xml:"InternetDomain"`
	CreateTime       string `json:"CreateTime" xml:"CreateTime"`
	ChargeType       string `json:"ChargeType" xml:"ChargeType"`
	IntranetAddress  string `json:"IntranetAddress" xml:"IntranetAddress"`
	InstanceId       string `json:"InstanceId" xml:"InstanceId"`
	InternetAddress  string `json:"InternetAddress" xml:"InternetAddress"`
	ClusterAliasName string `json:"ClusterAliasName" xml:"ClusterAliasName"`
	ClusterType      string `json:"ClusterType" xml:"ClusterType"`
	InitStatus       string `json:"InitStatus" xml:"InitStatus"`
	AppVersion       string `json:"AppVersion" xml:"AppVersion"`
	ClusterId        string `json:"ClusterId" xml:"ClusterId"`
	CanUpdate        bool   `json:"CanUpdate" xml:"CanUpdate"`
	VersionCode      string `json:"VersionCode" xml:"VersionCode"`
	InstanceCount    int64  `json:"InstanceCount" xml:"InstanceCount"`
	ClusterName      string `json:"ClusterName" xml:"ClusterName"`
	MseVersion       string `json:"MseVersion" xml:"MseVersion"`
}
