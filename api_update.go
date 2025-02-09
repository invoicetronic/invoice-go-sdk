/*
Italian eInvoice API

The Italian eInvoice API is a RESTful API that allows you to send and receive invoices through the Italian [Servizio di Interscambio (SDI)][1], or Interchange Service. The API is designed by Invoicetronic to be simple and easy to use, abstracting away SDI complexity while providing complete control over the invoice send/receive process. The API also provides advanced features as encryption at rest, invoice validation, multiple upload formats, webhooks, event logging, client SDKs for commonly used languages, and CLI tools.  For more information, see  [Invoicetronic website][2]  [1]: https://www.fatturapa.gov.it/it/sistemainterscambio/cose-il-sdi/ [2]: https://invoicetronic.com/

API version: 1.0.0
Contact: support@invoicetronic.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package invoicesdk

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)


// UpdateAPIService UpdateAPI service
type UpdateAPIService service

type ApiInvoiceV1UpdateGetRequest struct {
	ctx context.Context
	ApiService *UpdateAPIService
	companyId *int32
	identifier *string
	unread *bool
	sendId *int32
	state *string
	lastUpdateFrom *time.Time
	lastUpdateTo *time.Time
	dateSentFrom *time.Time
	dateSentTo *time.Time
	page *int32
	pageSize *int32
}

// Company id
func (r ApiInvoiceV1UpdateGetRequest) CompanyId(companyId int32) ApiInvoiceV1UpdateGetRequest {
	r.companyId = &companyId
	return r
}

// SDI identifier.
func (r ApiInvoiceV1UpdateGetRequest) Identifier(identifier string) ApiInvoiceV1UpdateGetRequest {
	r.identifier = &identifier
	return r
}

// Unread items only.
func (r ApiInvoiceV1UpdateGetRequest) Unread(unread bool) ApiInvoiceV1UpdateGetRequest {
	r.unread = &unread
	return r
}

// Send item&#39;s id.
func (r ApiInvoiceV1UpdateGetRequest) SendId(sendId int32) ApiInvoiceV1UpdateGetRequest {
	r.sendId = &sendId
	return r
}

// SDI state
func (r ApiInvoiceV1UpdateGetRequest) State(state string) ApiInvoiceV1UpdateGetRequest {
	r.state = &state
	return r
}

// UTC ISO 8601 (2024-11-29T12:34:56Z)
func (r ApiInvoiceV1UpdateGetRequest) LastUpdateFrom(lastUpdateFrom time.Time) ApiInvoiceV1UpdateGetRequest {
	r.lastUpdateFrom = &lastUpdateFrom
	return r
}

// UTC ISO 8601 (2024-11-29T12:34:56Z)
func (r ApiInvoiceV1UpdateGetRequest) LastUpdateTo(lastUpdateTo time.Time) ApiInvoiceV1UpdateGetRequest {
	r.lastUpdateTo = &lastUpdateTo
	return r
}

// UTC ISO 8601 (2024-11-29T12:34:56Z)
func (r ApiInvoiceV1UpdateGetRequest) DateSentFrom(dateSentFrom time.Time) ApiInvoiceV1UpdateGetRequest {
	r.dateSentFrom = &dateSentFrom
	return r
}

// UTC ISO 8601 (2024-11-29T12:34:56Z)
func (r ApiInvoiceV1UpdateGetRequest) DateSentTo(dateSentTo time.Time) ApiInvoiceV1UpdateGetRequest {
	r.dateSentTo = &dateSentTo
	return r
}

// Page number. Defaults to 1.
func (r ApiInvoiceV1UpdateGetRequest) Page(page int32) ApiInvoiceV1UpdateGetRequest {
	r.page = &page
	return r
}

// Items per page. Defaults to 50. Cannot be greater than 200.
func (r ApiInvoiceV1UpdateGetRequest) PageSize(pageSize int32) ApiInvoiceV1UpdateGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiInvoiceV1UpdateGetRequest) Execute() ([]Update, *http.Response, error) {
	return r.ApiService.InvoiceV1UpdateGetExecute(r)
}

/*
InvoiceV1UpdateGet List updates

Updates are notifications that are sent by the SDI about the status of sent invoices.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiInvoiceV1UpdateGetRequest
*/
func (a *UpdateAPIService) InvoiceV1UpdateGet(ctx context.Context) ApiInvoiceV1UpdateGetRequest {
	return ApiInvoiceV1UpdateGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []Update
func (a *UpdateAPIService) InvoiceV1UpdateGetExecute(r ApiInvoiceV1UpdateGetRequest) ([]Update, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Update
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UpdateAPIService.InvoiceV1UpdateGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/invoice/v1/update"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.companyId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "company_id", r.companyId, "form", "")
	}
	if r.identifier != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "identifier", r.identifier, "form", "")
	}
	if r.unread != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "unread", r.unread, "form", "")
	}
	if r.sendId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "send_id", r.sendId, "form", "")
	}
	if r.state != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "state", r.state, "form", "")
	}
	if r.lastUpdateFrom != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "last_update_from", r.lastUpdateFrom, "form", "")
	}
	if r.lastUpdateTo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "last_update_to", r.lastUpdateTo, "form", "")
	}
	if r.dateSentFrom != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "date_sent_from", r.dateSentFrom, "form", "")
	}
	if r.dateSentTo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "date_sent_to", r.dateSentTo, "form", "")
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	} else {
		var defaultValue int32 = 1
		r.page = &defaultValue
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize, "form", "")
	} else {
		var defaultValue int32 = 100
		r.pageSize = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ProblemHttpResult
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiInvoiceV1UpdateIdGetRequest struct {
	ctx context.Context
	ApiService *UpdateAPIService
	id int32
}

func (r ApiInvoiceV1UpdateIdGetRequest) Execute() (*Update, *http.Response, error) {
	return r.ApiService.InvoiceV1UpdateIdGetExecute(r)
}

/*
InvoiceV1UpdateIdGet Get an update by id

Updates are notifications that are sent by the SDI about the status of sent invoices.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Item id
 @return ApiInvoiceV1UpdateIdGetRequest
*/
func (a *UpdateAPIService) InvoiceV1UpdateIdGet(ctx context.Context, id int32) ApiInvoiceV1UpdateIdGetRequest {
	return ApiInvoiceV1UpdateIdGetRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return Update
func (a *UpdateAPIService) InvoiceV1UpdateIdGetExecute(r ApiInvoiceV1UpdateIdGetRequest) (*Update, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Update
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UpdateAPIService.InvoiceV1UpdateIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/invoice/v1/update/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
