package facebody

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

// Result is a nested struct in facebody response
type Result struct {
	Rate        float64     `json:"Rate" xml:"Rate"`
	Suggestion  string      `json:"Suggestion" xml:"Suggestion"`
	Label       string      `json:"Label" xml:"Label"`
	MessageTips string      `json:"MessageTips" xml:"MessageTips"`
	Rect        Rect        `json:"Rect" xml:"Rect"`
	Hands       Hands       `json:"Hands" xml:"Hands"`
	Box         Box         `json:"Box" xml:"Box"`
	Bodies      []Body      `json:"Bodies" xml:"Bodies"`
	Frames      []Frame     `json:"Frames" xml:"Frames"`
	SubResults  []SubResult `json:"SubResults" xml:"SubResults"`
}
