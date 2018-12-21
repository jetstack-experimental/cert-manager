package alidns

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

// AddBatchDomainRecords invokes the alidns.AddBatchDomainRecords API synchronously
// api document: https://help.aliyun.com/api/alidns/addbatchdomainrecords.html
func (client *Client) AddBatchDomainRecords(request *AddBatchDomainRecordsRequest) (response *AddBatchDomainRecordsResponse, err error) {
	response = CreateAddBatchDomainRecordsResponse()
	err = client.DoAction(request, response)
	return
}

// AddBatchDomainRecordsWithChan invokes the alidns.AddBatchDomainRecords API asynchronously
// api document: https://help.aliyun.com/api/alidns/addbatchdomainrecords.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddBatchDomainRecordsWithChan(request *AddBatchDomainRecordsRequest) (<-chan *AddBatchDomainRecordsResponse, <-chan error) {
	responseChan := make(chan *AddBatchDomainRecordsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddBatchDomainRecords(request)
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

// AddBatchDomainRecordsWithCallback invokes the alidns.AddBatchDomainRecords API asynchronously
// api document: https://help.aliyun.com/api/alidns/addbatchdomainrecords.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddBatchDomainRecordsWithCallback(request *AddBatchDomainRecordsRequest, callback func(response *AddBatchDomainRecordsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddBatchDomainRecordsResponse
		var err error
		defer close(result)
		response, err = client.AddBatchDomainRecords(request)
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

// AddBatchDomainRecordsRequest is the request struct for api AddBatchDomainRecords
type AddBatchDomainRecordsRequest struct {
	*requests.RpcRequest
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Records      string `position:"Query" name:"Records"`
}

// AddBatchDomainRecordsResponse is the response struct for api AddBatchDomainRecords
type AddBatchDomainRecordsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	TraceId   string `json:"TraceId" xml:"TraceId"`
}

// CreateAddBatchDomainRecordsRequest creates a request to invoke AddBatchDomainRecords API
func CreateAddBatchDomainRecordsRequest() (request *AddBatchDomainRecordsRequest) {
	request = &AddBatchDomainRecordsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "AddBatchDomainRecords", "", "")
	return
}

// CreateAddBatchDomainRecordsResponse creates a response to parse from AddBatchDomainRecords response
func CreateAddBatchDomainRecordsResponse() (response *AddBatchDomainRecordsResponse) {
	response = &AddBatchDomainRecordsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
