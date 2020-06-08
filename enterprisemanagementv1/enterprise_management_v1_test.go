/**
 * (C) Copyright IBM Corp. 2020.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package enterprisemanagementv1_test

import (
	"bytes"
	"fmt"
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/IBM/platform-services-go-sdk/enterprisemanagementv1"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"time"
)

var _ = Describe(`EnterpriseManagementV1`, func() {
	var testServer *httptest.Server
    Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			testService, testServiceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(testService).ToNot(BeNil())
			Expect(testServiceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			testService, testServiceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			testService, testServiceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
				URL: "https://enterprisemanagementv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"ENTERPRISE_MANAGEMENT_URL": "https://enterprisemanagementv1/api",
				"ENTERPRISE_MANAGEMENT_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := enterprisemanagementv1.NewEnterpriseManagementV1UsingExternalConfig(&enterprisemanagementv1.EnterpriseManagementV1Options{
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := enterprisemanagementv1.NewEnterpriseManagementV1UsingExternalConfig(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL: "https://testService/api",
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := enterprisemanagementv1.NewEnterpriseManagementV1UsingExternalConfig(&enterprisemanagementv1.EnterpriseManagementV1Options{
				})
				err := testService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"ENTERPRISE_MANAGEMENT_URL": "https://enterprisemanagementv1/api",
				"ENTERPRISE_MANAGEMENT_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := enterprisemanagementv1.NewEnterpriseManagementV1UsingExternalConfig(&enterprisemanagementv1.EnterpriseManagementV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"ENTERPRISE_MANAGEMENT_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := enterprisemanagementv1.NewEnterpriseManagementV1UsingExternalConfig(&enterprisemanagementv1.EnterpriseManagementV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`CreateAccountGroup(createAccountGroupOptions *CreateAccountGroupOptions) - Operation response error`, func() {
		createAccountGroupPath := "/account-groups"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			 		defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createAccountGroupPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateAccountGroup with error: Operation response processing error`, func() {
				testService, testServiceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the CreateAccountGroupOptions model
				createAccountGroupOptionsModel := new(enterprisemanagementv1.CreateAccountGroupOptions)
				createAccountGroupOptionsModel.Parent = core.StringPtr("testString")
				createAccountGroupOptionsModel.Name = core.StringPtr("testString")
				createAccountGroupOptionsModel.PrimaryContactIamID = core.StringPtr("testString")
				createAccountGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.CreateAccountGroup(createAccountGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateAccountGroup(createAccountGroupOptions *CreateAccountGroupOptions)`, func() {
		createAccountGroupPath := "/account-groups"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createAccountGroupPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `{"account_group_id": "AccountGroupID"}`)
				}))
			})
			It(`Invoke CreateAccountGroup successfully`, func() {
				testService, testServiceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreateAccountGroup(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateAccountGroupOptions model
				createAccountGroupOptionsModel := new(enterprisemanagementv1.CreateAccountGroupOptions)
				createAccountGroupOptionsModel.Parent = core.StringPtr("testString")
				createAccountGroupOptionsModel.Name = core.StringPtr("testString")
				createAccountGroupOptionsModel.PrimaryContactIamID = core.StringPtr("testString")
 				createAccountGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateAccountGroup(createAccountGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke CreateAccountGroup with error: Operation validation and request error`, func() {
				testService, testServiceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the CreateAccountGroupOptions model
				createAccountGroupOptionsModel := new(enterprisemanagementv1.CreateAccountGroupOptions)
				createAccountGroupOptionsModel.Parent = core.StringPtr("testString")
				createAccountGroupOptionsModel.Name = core.StringPtr("testString")
				createAccountGroupOptionsModel.PrimaryContactIamID = core.StringPtr("testString")
				createAccountGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.CreateAccountGroup(createAccountGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateAccountGroupOptions model with no property values
				createAccountGroupOptionsModelNew := new(enterprisemanagementv1.CreateAccountGroupOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.CreateAccountGroup(createAccountGroupOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListAccountGroups(listAccountGroupsOptions *ListAccountGroupsOptions) - Operation response error`, func() {
		listAccountGroupsPath := "/account-groups"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			 		defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listAccountGroupsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["enterprise_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["parent_account_group_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["parent"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListAccountGroups with error: Operation response processing error`, func() {
				testService, testServiceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListAccountGroupsOptions model
				listAccountGroupsOptionsModel := new(enterprisemanagementv1.ListAccountGroupsOptions)
				listAccountGroupsOptionsModel.EnterpriseID = core.StringPtr("testString")
				listAccountGroupsOptionsModel.ParentAccountGroupID = core.StringPtr("testString")
				listAccountGroupsOptionsModel.Parent = core.StringPtr("testString")
				listAccountGroupsOptionsModel.Limit = core.Int64Ptr(int64(38))
				listAccountGroupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.ListAccountGroups(listAccountGroupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListAccountGroups(listAccountGroupsOptions *ListAccountGroupsOptions)`, func() {
		listAccountGroupsPath := "/account-groups"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listAccountGroupsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["enterprise_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["parent_account_group_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["parent"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"rows_count": 9, "next_url": "NextURL", "resources": [{"url": "URL", "id": "ID", "crn": "Crn", "parent": "Parent", "enterprise_account_id": "EnterpriseAccountID", "enterprise_id": "EnterpriseID", "enterprise_path": "EnterprisePath", "name": "Name", "state": "State", "primary_contact_iam_id": "PrimaryContactIamID", "primary_contact_email": "PrimaryContactEmail", "created_at": "2019-01-01T12:00:00", "created_by": "CreatedBy", "updated_at": "2019-01-01T12:00:00", "updated_by": "UpdatedBy"}]}`)
				}))
			})
			It(`Invoke ListAccountGroups successfully`, func() {
				testService, testServiceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListAccountGroups(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListAccountGroupsOptions model
				listAccountGroupsOptionsModel := new(enterprisemanagementv1.ListAccountGroupsOptions)
				listAccountGroupsOptionsModel.EnterpriseID = core.StringPtr("testString")
				listAccountGroupsOptionsModel.ParentAccountGroupID = core.StringPtr("testString")
				listAccountGroupsOptionsModel.Parent = core.StringPtr("testString")
				listAccountGroupsOptionsModel.Limit = core.Int64Ptr(int64(38))
 				listAccountGroupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListAccountGroups(listAccountGroupsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke ListAccountGroups with error: Operation request error`, func() {
				testService, testServiceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListAccountGroupsOptions model
				listAccountGroupsOptionsModel := new(enterprisemanagementv1.ListAccountGroupsOptions)
				listAccountGroupsOptionsModel.EnterpriseID = core.StringPtr("testString")
				listAccountGroupsOptionsModel.ParentAccountGroupID = core.StringPtr("testString")
				listAccountGroupsOptionsModel.Parent = core.StringPtr("testString")
				listAccountGroupsOptionsModel.Limit = core.Int64Ptr(int64(38))
				listAccountGroupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.ListAccountGroups(listAccountGroupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetAccountGroup(getAccountGroupOptions *GetAccountGroupOptions) - Operation response error`, func() {
		getAccountGroupPath := "/account-groups/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			 		defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getAccountGroupPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetAccountGroup with error: Operation response processing error`, func() {
				testService, testServiceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetAccountGroupOptions model
				getAccountGroupOptionsModel := new(enterprisemanagementv1.GetAccountGroupOptions)
				getAccountGroupOptionsModel.AccountGroupID = core.StringPtr("testString")
				getAccountGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetAccountGroup(getAccountGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetAccountGroup(getAccountGroupOptions *GetAccountGroupOptions)`, func() {
		getAccountGroupPath := "/account-groups/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getAccountGroupPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"url": "URL", "id": "ID", "crn": "Crn", "parent": "Parent", "enterprise_account_id": "EnterpriseAccountID", "enterprise_id": "EnterpriseID", "enterprise_path": "EnterprisePath", "name": "Name", "state": "State", "primary_contact_iam_id": "PrimaryContactIamID", "primary_contact_email": "PrimaryContactEmail", "created_at": "2019-01-01T12:00:00", "created_by": "CreatedBy", "updated_at": "2019-01-01T12:00:00", "updated_by": "UpdatedBy"}`)
				}))
			})
			It(`Invoke GetAccountGroup successfully`, func() {
				testService, testServiceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetAccountGroup(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetAccountGroupOptions model
				getAccountGroupOptionsModel := new(enterprisemanagementv1.GetAccountGroupOptions)
				getAccountGroupOptionsModel.AccountGroupID = core.StringPtr("testString")
 				getAccountGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetAccountGroup(getAccountGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetAccountGroup with error: Operation validation and request error`, func() {
				testService, testServiceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetAccountGroupOptions model
				getAccountGroupOptionsModel := new(enterprisemanagementv1.GetAccountGroupOptions)
				getAccountGroupOptionsModel.AccountGroupID = core.StringPtr("testString")
				getAccountGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetAccountGroup(getAccountGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetAccountGroupOptions model with no property values
				getAccountGroupOptionsModelNew := new(enterprisemanagementv1.GetAccountGroupOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.GetAccountGroup(getAccountGroupOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateAccountGroup(updateAccountGroupOptions *UpdateAccountGroupOptions)`, func() {
		updateAccountGroupPath := "/account-groups/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateAccountGroupPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke UpdateAccountGroup successfully`, func() {
				testService, testServiceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.UpdateAccountGroup(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the UpdateAccountGroupOptions model
				updateAccountGroupOptionsModel := new(enterprisemanagementv1.UpdateAccountGroupOptions)
				updateAccountGroupOptionsModel.AccountGroupID = core.StringPtr("testString")
				updateAccountGroupOptionsModel.Name = core.StringPtr("testString")
				updateAccountGroupOptionsModel.PrimaryContactIamID = core.StringPtr("testString")
 				updateAccountGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.UpdateAccountGroup(updateAccountGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke UpdateAccountGroup with error: Operation validation and request error`, func() {
				testService, testServiceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the UpdateAccountGroupOptions model
				updateAccountGroupOptionsModel := new(enterprisemanagementv1.UpdateAccountGroupOptions)
				updateAccountGroupOptionsModel.AccountGroupID = core.StringPtr("testString")
				updateAccountGroupOptionsModel.Name = core.StringPtr("testString")
				updateAccountGroupOptionsModel.PrimaryContactIamID = core.StringPtr("testString")
				updateAccountGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := testService.UpdateAccountGroup(updateAccountGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the UpdateAccountGroupOptions model with no property values
				updateAccountGroupOptionsModelNew := new(enterprisemanagementv1.UpdateAccountGroupOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = testService.UpdateAccountGroup(updateAccountGroupOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
    Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			testService, testServiceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(testService).ToNot(BeNil())
			Expect(testServiceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			testService, testServiceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			testService, testServiceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
				URL: "https://enterprisemanagementv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"ENTERPRISE_MANAGEMENT_URL": "https://enterprisemanagementv1/api",
				"ENTERPRISE_MANAGEMENT_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := enterprisemanagementv1.NewEnterpriseManagementV1UsingExternalConfig(&enterprisemanagementv1.EnterpriseManagementV1Options{
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := enterprisemanagementv1.NewEnterpriseManagementV1UsingExternalConfig(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL: "https://testService/api",
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := enterprisemanagementv1.NewEnterpriseManagementV1UsingExternalConfig(&enterprisemanagementv1.EnterpriseManagementV1Options{
				})
				err := testService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"ENTERPRISE_MANAGEMENT_URL": "https://enterprisemanagementv1/api",
				"ENTERPRISE_MANAGEMENT_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := enterprisemanagementv1.NewEnterpriseManagementV1UsingExternalConfig(&enterprisemanagementv1.EnterpriseManagementV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"ENTERPRISE_MANAGEMENT_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := enterprisemanagementv1.NewEnterpriseManagementV1UsingExternalConfig(&enterprisemanagementv1.EnterpriseManagementV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})

	Describe(`ImportAccountToEnterprise(importAccountToEnterpriseOptions *ImportAccountToEnterpriseOptions)`, func() {
		importAccountToEnterprisePath := "/enterprises/testString/import/accounts/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(importAccountToEnterprisePath))
					Expect(req.Method).To(Equal("PUT"))
					res.WriteHeader(202)
				}))
			})
			It(`Invoke ImportAccountToEnterprise successfully`, func() {
				testService, testServiceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.ImportAccountToEnterprise(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the ImportAccountToEnterpriseOptions model
				importAccountToEnterpriseOptionsModel := new(enterprisemanagementv1.ImportAccountToEnterpriseOptions)
				importAccountToEnterpriseOptionsModel.EnterpriseID = core.StringPtr("testString")
				importAccountToEnterpriseOptionsModel.AccountID = core.StringPtr("testString")
				importAccountToEnterpriseOptionsModel.Parent = core.StringPtr("testString")
				importAccountToEnterpriseOptionsModel.BillingUnitID = core.StringPtr("testString")
 				importAccountToEnterpriseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.ImportAccountToEnterprise(importAccountToEnterpriseOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke ImportAccountToEnterprise with error: Operation validation and request error`, func() {
				testService, testServiceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ImportAccountToEnterpriseOptions model
				importAccountToEnterpriseOptionsModel := new(enterprisemanagementv1.ImportAccountToEnterpriseOptions)
				importAccountToEnterpriseOptionsModel.EnterpriseID = core.StringPtr("testString")
				importAccountToEnterpriseOptionsModel.AccountID = core.StringPtr("testString")
				importAccountToEnterpriseOptionsModel.Parent = core.StringPtr("testString")
				importAccountToEnterpriseOptionsModel.BillingUnitID = core.StringPtr("testString")
				importAccountToEnterpriseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := testService.ImportAccountToEnterprise(importAccountToEnterpriseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the ImportAccountToEnterpriseOptions model with no property values
				importAccountToEnterpriseOptionsModelNew := new(enterprisemanagementv1.ImportAccountToEnterpriseOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = testService.ImportAccountToEnterprise(importAccountToEnterpriseOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateAccount(createAccountOptions *CreateAccountOptions) - Operation response error`, func() {
		createAccountPath := "/accounts"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			 		defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createAccountPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateAccount with error: Operation response processing error`, func() {
				testService, testServiceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the CreateAccountOptions model
				createAccountOptionsModel := new(enterprisemanagementv1.CreateAccountOptions)
				createAccountOptionsModel.Parent = core.StringPtr("testString")
				createAccountOptionsModel.Name = core.StringPtr("testString")
				createAccountOptionsModel.OwnerIamID = core.StringPtr("testString")
				createAccountOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.CreateAccount(createAccountOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateAccount(createAccountOptions *CreateAccountOptions)`, func() {
		createAccountPath := "/accounts"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createAccountPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `{"account_id": "AccountID"}`)
				}))
			})
			It(`Invoke CreateAccount successfully`, func() {
				testService, testServiceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreateAccount(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateAccountOptions model
				createAccountOptionsModel := new(enterprisemanagementv1.CreateAccountOptions)
				createAccountOptionsModel.Parent = core.StringPtr("testString")
				createAccountOptionsModel.Name = core.StringPtr("testString")
				createAccountOptionsModel.OwnerIamID = core.StringPtr("testString")
 				createAccountOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateAccount(createAccountOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke CreateAccount with error: Operation validation and request error`, func() {
				testService, testServiceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the CreateAccountOptions model
				createAccountOptionsModel := new(enterprisemanagementv1.CreateAccountOptions)
				createAccountOptionsModel.Parent = core.StringPtr("testString")
				createAccountOptionsModel.Name = core.StringPtr("testString")
				createAccountOptionsModel.OwnerIamID = core.StringPtr("testString")
				createAccountOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.CreateAccount(createAccountOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateAccountOptions model with no property values
				createAccountOptionsModelNew := new(enterprisemanagementv1.CreateAccountOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.CreateAccount(createAccountOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListAccounts(listAccountsOptions *ListAccountsOptions) - Operation response error`, func() {
		listAccountsPath := "/accounts"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			 		defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listAccountsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["enterprise_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["account_group_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["parent"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListAccounts with error: Operation response processing error`, func() {
				testService, testServiceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListAccountsOptions model
				listAccountsOptionsModel := new(enterprisemanagementv1.ListAccountsOptions)
				listAccountsOptionsModel.EnterpriseID = core.StringPtr("testString")
				listAccountsOptionsModel.AccountGroupID = core.StringPtr("testString")
				listAccountsOptionsModel.Parent = core.StringPtr("testString")
				listAccountsOptionsModel.Limit = core.Int64Ptr(int64(38))
				listAccountsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.ListAccounts(listAccountsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListAccounts(listAccountsOptions *ListAccountsOptions)`, func() {
		listAccountsPath := "/accounts"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listAccountsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["enterprise_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["account_group_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["parent"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"rows_count": 9, "next_url": "NextURL", "resources": [{"url": "URL", "id": "ID", "crn": "Crn", "parent": "Parent", "enterprise_account_id": "EnterpriseAccountID", "enterprise_id": "EnterpriseID", "enterprise_path": "EnterprisePath", "name": "Name", "state": "State", "owner_iam_id": "OwnerIamID", "paid": true, "owner_email": "OwnerEmail", "is_enterprise_account": false, "created_at": "2019-01-01T12:00:00", "created_by": "CreatedBy", "updated_at": "2019-01-01T12:00:00", "updated_by": "UpdatedBy"}]}`)
				}))
			})
			It(`Invoke ListAccounts successfully`, func() {
				testService, testServiceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListAccounts(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListAccountsOptions model
				listAccountsOptionsModel := new(enterprisemanagementv1.ListAccountsOptions)
				listAccountsOptionsModel.EnterpriseID = core.StringPtr("testString")
				listAccountsOptionsModel.AccountGroupID = core.StringPtr("testString")
				listAccountsOptionsModel.Parent = core.StringPtr("testString")
				listAccountsOptionsModel.Limit = core.Int64Ptr(int64(38))
 				listAccountsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListAccounts(listAccountsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke ListAccounts with error: Operation request error`, func() {
				testService, testServiceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListAccountsOptions model
				listAccountsOptionsModel := new(enterprisemanagementv1.ListAccountsOptions)
				listAccountsOptionsModel.EnterpriseID = core.StringPtr("testString")
				listAccountsOptionsModel.AccountGroupID = core.StringPtr("testString")
				listAccountsOptionsModel.Parent = core.StringPtr("testString")
				listAccountsOptionsModel.Limit = core.Int64Ptr(int64(38))
				listAccountsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.ListAccounts(listAccountsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetAccount(getAccountOptions *GetAccountOptions) - Operation response error`, func() {
		getAccountPath := "/accounts/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			 		defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getAccountPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetAccount with error: Operation response processing error`, func() {
				testService, testServiceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetAccountOptions model
				getAccountOptionsModel := new(enterprisemanagementv1.GetAccountOptions)
				getAccountOptionsModel.AccountID = core.StringPtr("testString")
				getAccountOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetAccount(getAccountOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetAccount(getAccountOptions *GetAccountOptions)`, func() {
		getAccountPath := "/accounts/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getAccountPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"url": "URL", "id": "ID", "crn": "Crn", "parent": "Parent", "enterprise_account_id": "EnterpriseAccountID", "enterprise_id": "EnterpriseID", "enterprise_path": "EnterprisePath", "name": "Name", "state": "State", "owner_iam_id": "OwnerIamID", "paid": true, "owner_email": "OwnerEmail", "is_enterprise_account": false, "created_at": "2019-01-01T12:00:00", "created_by": "CreatedBy", "updated_at": "2019-01-01T12:00:00", "updated_by": "UpdatedBy"}`)
				}))
			})
			It(`Invoke GetAccount successfully`, func() {
				testService, testServiceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetAccount(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetAccountOptions model
				getAccountOptionsModel := new(enterprisemanagementv1.GetAccountOptions)
				getAccountOptionsModel.AccountID = core.StringPtr("testString")
 				getAccountOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetAccount(getAccountOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetAccount with error: Operation validation and request error`, func() {
				testService, testServiceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetAccountOptions model
				getAccountOptionsModel := new(enterprisemanagementv1.GetAccountOptions)
				getAccountOptionsModel.AccountID = core.StringPtr("testString")
				getAccountOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetAccount(getAccountOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetAccountOptions model with no property values
				getAccountOptionsModelNew := new(enterprisemanagementv1.GetAccountOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.GetAccount(getAccountOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateAccount(updateAccountOptions *UpdateAccountOptions)`, func() {
		updateAccountPath := "/accounts/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateAccountPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke UpdateAccount successfully`, func() {
				testService, testServiceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.UpdateAccount(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the UpdateAccountOptions model
				updateAccountOptionsModel := new(enterprisemanagementv1.UpdateAccountOptions)
				updateAccountOptionsModel.AccountID = core.StringPtr("testString")
				updateAccountOptionsModel.Parent = core.StringPtr("testString")
 				updateAccountOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.UpdateAccount(updateAccountOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke UpdateAccount with error: Operation validation and request error`, func() {
				testService, testServiceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the UpdateAccountOptions model
				updateAccountOptionsModel := new(enterprisemanagementv1.UpdateAccountOptions)
				updateAccountOptionsModel.AccountID = core.StringPtr("testString")
				updateAccountOptionsModel.Parent = core.StringPtr("testString")
				updateAccountOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := testService.UpdateAccount(updateAccountOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the UpdateAccountOptions model with no property values
				updateAccountOptionsModelNew := new(enterprisemanagementv1.UpdateAccountOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = testService.UpdateAccount(updateAccountOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
    Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			testService, testServiceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(testService).ToNot(BeNil())
			Expect(testServiceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			testService, testServiceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			testService, testServiceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
				URL: "https://enterprisemanagementv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"ENTERPRISE_MANAGEMENT_URL": "https://enterprisemanagementv1/api",
				"ENTERPRISE_MANAGEMENT_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := enterprisemanagementv1.NewEnterpriseManagementV1UsingExternalConfig(&enterprisemanagementv1.EnterpriseManagementV1Options{
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := enterprisemanagementv1.NewEnterpriseManagementV1UsingExternalConfig(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL: "https://testService/api",
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := enterprisemanagementv1.NewEnterpriseManagementV1UsingExternalConfig(&enterprisemanagementv1.EnterpriseManagementV1Options{
				})
				err := testService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"ENTERPRISE_MANAGEMENT_URL": "https://enterprisemanagementv1/api",
				"ENTERPRISE_MANAGEMENT_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := enterprisemanagementv1.NewEnterpriseManagementV1UsingExternalConfig(&enterprisemanagementv1.EnterpriseManagementV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"ENTERPRISE_MANAGEMENT_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := enterprisemanagementv1.NewEnterpriseManagementV1UsingExternalConfig(&enterprisemanagementv1.EnterpriseManagementV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`CreateEnterprise(createEnterpriseOptions *CreateEnterpriseOptions) - Operation response error`, func() {
		createEnterprisePath := "/enterprises"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			 		defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createEnterprisePath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateEnterprise with error: Operation response processing error`, func() {
				testService, testServiceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the CreateEnterpriseOptions model
				createEnterpriseOptionsModel := new(enterprisemanagementv1.CreateEnterpriseOptions)
				createEnterpriseOptionsModel.SourceAccountID = core.StringPtr("testString")
				createEnterpriseOptionsModel.Name = core.StringPtr("testString")
				createEnterpriseOptionsModel.PrimaryContactIamID = core.StringPtr("testString")
				createEnterpriseOptionsModel.Domain = core.StringPtr("testString")
				createEnterpriseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.CreateEnterprise(createEnterpriseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateEnterprise(createEnterpriseOptions *CreateEnterpriseOptions)`, func() {
		createEnterprisePath := "/enterprises"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createEnterprisePath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, `{"enterprise_id": "EnterpriseID", "enterprise_account_id": "EnterpriseAccountID"}`)
				}))
			})
			It(`Invoke CreateEnterprise successfully`, func() {
				testService, testServiceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreateEnterprise(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateEnterpriseOptions model
				createEnterpriseOptionsModel := new(enterprisemanagementv1.CreateEnterpriseOptions)
				createEnterpriseOptionsModel.SourceAccountID = core.StringPtr("testString")
				createEnterpriseOptionsModel.Name = core.StringPtr("testString")
				createEnterpriseOptionsModel.PrimaryContactIamID = core.StringPtr("testString")
				createEnterpriseOptionsModel.Domain = core.StringPtr("testString")
 				createEnterpriseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateEnterprise(createEnterpriseOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke CreateEnterprise with error: Operation validation and request error`, func() {
				testService, testServiceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the CreateEnterpriseOptions model
				createEnterpriseOptionsModel := new(enterprisemanagementv1.CreateEnterpriseOptions)
				createEnterpriseOptionsModel.SourceAccountID = core.StringPtr("testString")
				createEnterpriseOptionsModel.Name = core.StringPtr("testString")
				createEnterpriseOptionsModel.PrimaryContactIamID = core.StringPtr("testString")
				createEnterpriseOptionsModel.Domain = core.StringPtr("testString")
				createEnterpriseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.CreateEnterprise(createEnterpriseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateEnterpriseOptions model with no property values
				createEnterpriseOptionsModelNew := new(enterprisemanagementv1.CreateEnterpriseOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.CreateEnterprise(createEnterpriseOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListEnterprises(listEnterprisesOptions *ListEnterprisesOptions) - Operation response error`, func() {
		listEnterprisesPath := "/enterprises"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			 		defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listEnterprisesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["enterprise_account_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["account_group_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListEnterprises with error: Operation response processing error`, func() {
				testService, testServiceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListEnterprisesOptions model
				listEnterprisesOptionsModel := new(enterprisemanagementv1.ListEnterprisesOptions)
				listEnterprisesOptionsModel.EnterpriseAccountID = core.StringPtr("testString")
				listEnterprisesOptionsModel.AccountGroupID = core.StringPtr("testString")
				listEnterprisesOptionsModel.AccountID = core.StringPtr("testString")
				listEnterprisesOptionsModel.Limit = core.Int64Ptr(int64(38))
				listEnterprisesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.ListEnterprises(listEnterprisesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListEnterprises(listEnterprisesOptions *ListEnterprisesOptions)`, func() {
		listEnterprisesPath := "/enterprises"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listEnterprisesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["enterprise_account_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["account_group_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"rows_count": 9, "next_url": "NextURL", "resources": [{"url": "URL", "id": "ID", "enterprise_account_id": "EnterpriseAccountID", "crn": "Crn", "name": "Name", "domain": "Domain", "state": "State", "primary_contact_iam_id": "PrimaryContactIamID", "primary_contact_email": "PrimaryContactEmail", "created_at": "2019-01-01T12:00:00", "created_by": "CreatedBy", "updated_at": "2019-01-01T12:00:00", "updated_by": "UpdatedBy"}]}`)
				}))
			})
			It(`Invoke ListEnterprises successfully`, func() {
				testService, testServiceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListEnterprises(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListEnterprisesOptions model
				listEnterprisesOptionsModel := new(enterprisemanagementv1.ListEnterprisesOptions)
				listEnterprisesOptionsModel.EnterpriseAccountID = core.StringPtr("testString")
				listEnterprisesOptionsModel.AccountGroupID = core.StringPtr("testString")
				listEnterprisesOptionsModel.AccountID = core.StringPtr("testString")
				listEnterprisesOptionsModel.Limit = core.Int64Ptr(int64(38))
 				listEnterprisesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListEnterprises(listEnterprisesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke ListEnterprises with error: Operation request error`, func() {
				testService, testServiceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListEnterprisesOptions model
				listEnterprisesOptionsModel := new(enterprisemanagementv1.ListEnterprisesOptions)
				listEnterprisesOptionsModel.EnterpriseAccountID = core.StringPtr("testString")
				listEnterprisesOptionsModel.AccountGroupID = core.StringPtr("testString")
				listEnterprisesOptionsModel.AccountID = core.StringPtr("testString")
				listEnterprisesOptionsModel.Limit = core.Int64Ptr(int64(38))
				listEnterprisesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.ListEnterprises(listEnterprisesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetEnterprise(getEnterpriseOptions *GetEnterpriseOptions) - Operation response error`, func() {
		getEnterprisePath := "/enterprises/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			 		defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getEnterprisePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetEnterprise with error: Operation response processing error`, func() {
				testService, testServiceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetEnterpriseOptions model
				getEnterpriseOptionsModel := new(enterprisemanagementv1.GetEnterpriseOptions)
				getEnterpriseOptionsModel.EnterpriseID = core.StringPtr("testString")
				getEnterpriseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetEnterprise(getEnterpriseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetEnterprise(getEnterpriseOptions *GetEnterpriseOptions)`, func() {
		getEnterprisePath := "/enterprises/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getEnterprisePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"url": "URL", "id": "ID", "enterprise_account_id": "EnterpriseAccountID", "crn": "Crn", "name": "Name", "domain": "Domain", "state": "State", "primary_contact_iam_id": "PrimaryContactIamID", "primary_contact_email": "PrimaryContactEmail", "created_at": "2019-01-01T12:00:00", "created_by": "CreatedBy", "updated_at": "2019-01-01T12:00:00", "updated_by": "UpdatedBy"}`)
				}))
			})
			It(`Invoke GetEnterprise successfully`, func() {
				testService, testServiceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetEnterprise(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetEnterpriseOptions model
				getEnterpriseOptionsModel := new(enterprisemanagementv1.GetEnterpriseOptions)
				getEnterpriseOptionsModel.EnterpriseID = core.StringPtr("testString")
 				getEnterpriseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetEnterprise(getEnterpriseOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetEnterprise with error: Operation validation and request error`, func() {
				testService, testServiceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetEnterpriseOptions model
				getEnterpriseOptionsModel := new(enterprisemanagementv1.GetEnterpriseOptions)
				getEnterpriseOptionsModel.EnterpriseID = core.StringPtr("testString")
				getEnterpriseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetEnterprise(getEnterpriseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetEnterpriseOptions model with no property values
				getEnterpriseOptionsModelNew := new(enterprisemanagementv1.GetEnterpriseOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.GetEnterprise(getEnterpriseOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateEnterprise(updateEnterpriseOptions *UpdateEnterpriseOptions)`, func() {
		updateEnterprisePath := "/enterprises/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateEnterprisePath))
					Expect(req.Method).To(Equal("PATCH"))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke UpdateEnterprise successfully`, func() {
				testService, testServiceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.UpdateEnterprise(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the UpdateEnterpriseOptions model
				updateEnterpriseOptionsModel := new(enterprisemanagementv1.UpdateEnterpriseOptions)
				updateEnterpriseOptionsModel.EnterpriseID = core.StringPtr("testString")
				updateEnterpriseOptionsModel.Name = core.StringPtr("testString")
				updateEnterpriseOptionsModel.Domain = core.StringPtr("testString")
				updateEnterpriseOptionsModel.PrimaryContactIamID = core.StringPtr("testString")
 				updateEnterpriseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.UpdateEnterprise(updateEnterpriseOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke UpdateEnterprise with error: Operation validation and request error`, func() {
				testService, testServiceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the UpdateEnterpriseOptions model
				updateEnterpriseOptionsModel := new(enterprisemanagementv1.UpdateEnterpriseOptions)
				updateEnterpriseOptionsModel.EnterpriseID = core.StringPtr("testString")
				updateEnterpriseOptionsModel.Name = core.StringPtr("testString")
				updateEnterpriseOptionsModel.Domain = core.StringPtr("testString")
				updateEnterpriseOptionsModel.PrimaryContactIamID = core.StringPtr("testString")
				updateEnterpriseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := testService.UpdateEnterprise(updateEnterpriseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the UpdateEnterpriseOptions model with no property values
				updateEnterpriseOptionsModelNew := new(enterprisemanagementv1.UpdateEnterpriseOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = testService.UpdateEnterprise(updateEnterpriseOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			testService, _ := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
				URL:           "http://enterprisemanagementv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewCreateAccountGroupOptions successfully`, func() {
				// Construct an instance of the CreateAccountGroupOptions model
				createAccountGroupOptionsParent := "testString"
				createAccountGroupOptionsName := "testString"
				createAccountGroupOptionsPrimaryContactIamID := "testString"
				createAccountGroupOptionsModel := testService.NewCreateAccountGroupOptions(createAccountGroupOptionsParent, createAccountGroupOptionsName, createAccountGroupOptionsPrimaryContactIamID)
				createAccountGroupOptionsModel.SetParent("testString")
				createAccountGroupOptionsModel.SetName("testString")
				createAccountGroupOptionsModel.SetPrimaryContactIamID("testString")
				createAccountGroupOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createAccountGroupOptionsModel).ToNot(BeNil())
				Expect(createAccountGroupOptionsModel.Parent).To(Equal(core.StringPtr("testString")))
				Expect(createAccountGroupOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createAccountGroupOptionsModel.PrimaryContactIamID).To(Equal(core.StringPtr("testString")))
				Expect(createAccountGroupOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateAccountOptions successfully`, func() {
				// Construct an instance of the CreateAccountOptions model
				createAccountOptionsParent := "testString"
				createAccountOptionsName := "testString"
				createAccountOptionsOwnerIamID := "testString"
				createAccountOptionsModel := testService.NewCreateAccountOptions(createAccountOptionsParent, createAccountOptionsName, createAccountOptionsOwnerIamID)
				createAccountOptionsModel.SetParent("testString")
				createAccountOptionsModel.SetName("testString")
				createAccountOptionsModel.SetOwnerIamID("testString")
				createAccountOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createAccountOptionsModel).ToNot(BeNil())
				Expect(createAccountOptionsModel.Parent).To(Equal(core.StringPtr("testString")))
				Expect(createAccountOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createAccountOptionsModel.OwnerIamID).To(Equal(core.StringPtr("testString")))
				Expect(createAccountOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateEnterpriseOptions successfully`, func() {
				// Construct an instance of the CreateEnterpriseOptions model
				createEnterpriseOptionsSourceAccountID := "testString"
				createEnterpriseOptionsName := "testString"
				createEnterpriseOptionsPrimaryContactIamID := "testString"
				createEnterpriseOptionsModel := testService.NewCreateEnterpriseOptions(createEnterpriseOptionsSourceAccountID, createEnterpriseOptionsName, createEnterpriseOptionsPrimaryContactIamID)
				createEnterpriseOptionsModel.SetSourceAccountID("testString")
				createEnterpriseOptionsModel.SetName("testString")
				createEnterpriseOptionsModel.SetPrimaryContactIamID("testString")
				createEnterpriseOptionsModel.SetDomain("testString")
				createEnterpriseOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createEnterpriseOptionsModel).ToNot(BeNil())
				Expect(createEnterpriseOptionsModel.SourceAccountID).To(Equal(core.StringPtr("testString")))
				Expect(createEnterpriseOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createEnterpriseOptionsModel.PrimaryContactIamID).To(Equal(core.StringPtr("testString")))
				Expect(createEnterpriseOptionsModel.Domain).To(Equal(core.StringPtr("testString")))
				Expect(createEnterpriseOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetAccountGroupOptions successfully`, func() {
				// Construct an instance of the GetAccountGroupOptions model
				accountGroupID := "testString"
				getAccountGroupOptionsModel := testService.NewGetAccountGroupOptions(accountGroupID)
				getAccountGroupOptionsModel.SetAccountGroupID("testString")
				getAccountGroupOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getAccountGroupOptionsModel).ToNot(BeNil())
				Expect(getAccountGroupOptionsModel.AccountGroupID).To(Equal(core.StringPtr("testString")))
				Expect(getAccountGroupOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetAccountOptions successfully`, func() {
				// Construct an instance of the GetAccountOptions model
				accountID := "testString"
				getAccountOptionsModel := testService.NewGetAccountOptions(accountID)
				getAccountOptionsModel.SetAccountID("testString")
				getAccountOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getAccountOptionsModel).ToNot(BeNil())
				Expect(getAccountOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(getAccountOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetEnterpriseOptions successfully`, func() {
				// Construct an instance of the GetEnterpriseOptions model
				enterpriseID := "testString"
				getEnterpriseOptionsModel := testService.NewGetEnterpriseOptions(enterpriseID)
				getEnterpriseOptionsModel.SetEnterpriseID("testString")
				getEnterpriseOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getEnterpriseOptionsModel).ToNot(BeNil())
				Expect(getEnterpriseOptionsModel.EnterpriseID).To(Equal(core.StringPtr("testString")))
				Expect(getEnterpriseOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewImportAccountToEnterpriseOptions successfully`, func() {
				// Construct an instance of the ImportAccountToEnterpriseOptions model
				enterpriseID := "testString"
				accountID := "testString"
				importAccountToEnterpriseOptionsModel := testService.NewImportAccountToEnterpriseOptions(enterpriseID, accountID)
				importAccountToEnterpriseOptionsModel.SetEnterpriseID("testString")
				importAccountToEnterpriseOptionsModel.SetAccountID("testString")
				importAccountToEnterpriseOptionsModel.SetParent("testString")
				importAccountToEnterpriseOptionsModel.SetBillingUnitID("testString")
				importAccountToEnterpriseOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(importAccountToEnterpriseOptionsModel).ToNot(BeNil())
				Expect(importAccountToEnterpriseOptionsModel.EnterpriseID).To(Equal(core.StringPtr("testString")))
				Expect(importAccountToEnterpriseOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(importAccountToEnterpriseOptionsModel.Parent).To(Equal(core.StringPtr("testString")))
				Expect(importAccountToEnterpriseOptionsModel.BillingUnitID).To(Equal(core.StringPtr("testString")))
				Expect(importAccountToEnterpriseOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListAccountGroupsOptions successfully`, func() {
				// Construct an instance of the ListAccountGroupsOptions model
				listAccountGroupsOptionsModel := testService.NewListAccountGroupsOptions()
				listAccountGroupsOptionsModel.SetEnterpriseID("testString")
				listAccountGroupsOptionsModel.SetParentAccountGroupID("testString")
				listAccountGroupsOptionsModel.SetParent("testString")
				listAccountGroupsOptionsModel.SetLimit(int64(38))
				listAccountGroupsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listAccountGroupsOptionsModel).ToNot(BeNil())
				Expect(listAccountGroupsOptionsModel.EnterpriseID).To(Equal(core.StringPtr("testString")))
				Expect(listAccountGroupsOptionsModel.ParentAccountGroupID).To(Equal(core.StringPtr("testString")))
				Expect(listAccountGroupsOptionsModel.Parent).To(Equal(core.StringPtr("testString")))
				Expect(listAccountGroupsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(38))))
				Expect(listAccountGroupsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListAccountsOptions successfully`, func() {
				// Construct an instance of the ListAccountsOptions model
				listAccountsOptionsModel := testService.NewListAccountsOptions()
				listAccountsOptionsModel.SetEnterpriseID("testString")
				listAccountsOptionsModel.SetAccountGroupID("testString")
				listAccountsOptionsModel.SetParent("testString")
				listAccountsOptionsModel.SetLimit(int64(38))
				listAccountsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listAccountsOptionsModel).ToNot(BeNil())
				Expect(listAccountsOptionsModel.EnterpriseID).To(Equal(core.StringPtr("testString")))
				Expect(listAccountsOptionsModel.AccountGroupID).To(Equal(core.StringPtr("testString")))
				Expect(listAccountsOptionsModel.Parent).To(Equal(core.StringPtr("testString")))
				Expect(listAccountsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(38))))
				Expect(listAccountsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListEnterprisesOptions successfully`, func() {
				// Construct an instance of the ListEnterprisesOptions model
				listEnterprisesOptionsModel := testService.NewListEnterprisesOptions()
				listEnterprisesOptionsModel.SetEnterpriseAccountID("testString")
				listEnterprisesOptionsModel.SetAccountGroupID("testString")
				listEnterprisesOptionsModel.SetAccountID("testString")
				listEnterprisesOptionsModel.SetLimit(int64(38))
				listEnterprisesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listEnterprisesOptionsModel).ToNot(BeNil())
				Expect(listEnterprisesOptionsModel.EnterpriseAccountID).To(Equal(core.StringPtr("testString")))
				Expect(listEnterprisesOptionsModel.AccountGroupID).To(Equal(core.StringPtr("testString")))
				Expect(listEnterprisesOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(listEnterprisesOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(38))))
				Expect(listEnterprisesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateAccountGroupOptions successfully`, func() {
				// Construct an instance of the UpdateAccountGroupOptions model
				accountGroupID := "testString"
				updateAccountGroupOptionsModel := testService.NewUpdateAccountGroupOptions(accountGroupID)
				updateAccountGroupOptionsModel.SetAccountGroupID("testString")
				updateAccountGroupOptionsModel.SetName("testString")
				updateAccountGroupOptionsModel.SetPrimaryContactIamID("testString")
				updateAccountGroupOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateAccountGroupOptionsModel).ToNot(BeNil())
				Expect(updateAccountGroupOptionsModel.AccountGroupID).To(Equal(core.StringPtr("testString")))
				Expect(updateAccountGroupOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(updateAccountGroupOptionsModel.PrimaryContactIamID).To(Equal(core.StringPtr("testString")))
				Expect(updateAccountGroupOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateAccountOptions successfully`, func() {
				// Construct an instance of the UpdateAccountOptions model
				accountID := "testString"
				updateAccountOptionsParent := "testString"
				updateAccountOptionsModel := testService.NewUpdateAccountOptions(accountID, updateAccountOptionsParent)
				updateAccountOptionsModel.SetAccountID("testString")
				updateAccountOptionsModel.SetParent("testString")
				updateAccountOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateAccountOptionsModel).ToNot(BeNil())
				Expect(updateAccountOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(updateAccountOptionsModel.Parent).To(Equal(core.StringPtr("testString")))
				Expect(updateAccountOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateEnterpriseOptions successfully`, func() {
				// Construct an instance of the UpdateEnterpriseOptions model
				enterpriseID := "testString"
				updateEnterpriseOptionsModel := testService.NewUpdateEnterpriseOptions(enterpriseID)
				updateEnterpriseOptionsModel.SetEnterpriseID("testString")
				updateEnterpriseOptionsModel.SetName("testString")
				updateEnterpriseOptionsModel.SetDomain("testString")
				updateEnterpriseOptionsModel.SetPrimaryContactIamID("testString")
				updateEnterpriseOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateEnterpriseOptionsModel).ToNot(BeNil())
				Expect(updateEnterpriseOptionsModel.EnterpriseID).To(Equal(core.StringPtr("testString")))
				Expect(updateEnterpriseOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(updateEnterpriseOptionsModel.Domain).To(Equal(core.StringPtr("testString")))
				Expect(updateEnterpriseOptionsModel.PrimaryContactIamID).To(Equal(core.StringPtr("testString")))
				Expect(updateEnterpriseOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
		})
	})
	Describe(`Utility function tests`, func() {
		It(`Invoke CreateMockByteArray() successfully`, func() {
			mockByteArray := CreateMockByteArray("This is a test")
			Expect(mockByteArray).ToNot(BeNil())
		})
		It(`Invoke CreateMockUUID() successfully`, func() {
			mockUUID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
			Expect(mockUUID).ToNot(BeNil())
		})
		It(`Invoke CreateMockReader() successfully`, func() {
			mockReader := CreateMockReader("This is a test.")
			Expect(mockReader).ToNot(BeNil())
		})
		It(`Invoke CreateMockDate() successfully`, func() {
			mockDate := CreateMockDate()
			Expect(mockDate).ToNot(BeNil())
		})
		It(`Invoke CreateMockDateTime() successfully`, func() {
			mockDateTime := CreateMockDateTime()
			Expect(mockDateTime).ToNot(BeNil())
		})
	})
})

//
// Utility functions used by the generated test code
//

func CreateMockByteArray(mockData string) *[]byte {
	ba := make([]byte, 0)
	ba = append(ba, mockData...)
	return &ba
}

func CreateMockUUID(mockData string) *strfmt.UUID {
	uuid := strfmt.UUID(mockData)
	return &uuid
}

func CreateMockReader(mockData string) io.ReadCloser {
	return ioutil.NopCloser(bytes.NewReader([]byte(mockData)))
}

func CreateMockDate() *strfmt.Date {
	d := strfmt.Date(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))
	return &d
}

func CreateMockDateTime() *strfmt.DateTime {
	d := strfmt.DateTime(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))
	return &d
}

func SetTestEnvironment(testEnvironment map[string]string) {
	for key, value := range testEnvironment {
		os.Setenv(key, value)
	}
}

func ClearTestEnvironment(testEnvironment map[string]string) {
	for key := range testEnvironment {
		os.Unsetenv(key)
	}
}