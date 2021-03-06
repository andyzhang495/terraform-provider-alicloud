package cr

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

// CreateUserSourceAccount invokes the cr.CreateUserSourceAccount API synchronously
// api document: https://help.aliyun.com/api/cr/createusersourceaccount.html
func (client *Client) CreateUserSourceAccount(request *CreateUserSourceAccountRequest) (response *CreateUserSourceAccountResponse, err error) {
	response = CreateCreateUserSourceAccountResponse()
	err = client.DoAction(request, response)
	return
}

// CreateUserSourceAccountWithChan invokes the cr.CreateUserSourceAccount API asynchronously
// api document: https://help.aliyun.com/api/cr/createusersourceaccount.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateUserSourceAccountWithChan(request *CreateUserSourceAccountRequest) (<-chan *CreateUserSourceAccountResponse, <-chan error) {
	responseChan := make(chan *CreateUserSourceAccountResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateUserSourceAccount(request)
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

// CreateUserSourceAccountWithCallback invokes the cr.CreateUserSourceAccount API asynchronously
// api document: https://help.aliyun.com/api/cr/createusersourceaccount.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateUserSourceAccountWithCallback(request *CreateUserSourceAccountRequest, callback func(response *CreateUserSourceAccountResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateUserSourceAccountResponse
		var err error
		defer close(result)
		response, err = client.CreateUserSourceAccount(request)
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

// CreateUserSourceAccountRequest is the request struct for api CreateUserSourceAccount
type CreateUserSourceAccountRequest struct {
	*requests.RoaRequest
}

// CreateUserSourceAccountResponse is the response struct for api CreateUserSourceAccount
type CreateUserSourceAccountResponse struct {
	*responses.BaseResponse
}

// CreateCreateUserSourceAccountRequest creates a request to invoke CreateUserSourceAccount API
func CreateCreateUserSourceAccountRequest() (request *CreateUserSourceAccountRequest) {
	request = &CreateUserSourceAccountRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("cr", "2016-06-07", "CreateUserSourceAccount", "/users/sourceAccount", "acr", "openAPI")
	request.Method = requests.PUT
	return
}

// CreateCreateUserSourceAccountResponse creates a response to parse from CreateUserSourceAccount response
func CreateCreateUserSourceAccountResponse() (response *CreateUserSourceAccountResponse) {
	response = &CreateUserSourceAccountResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
