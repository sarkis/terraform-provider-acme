// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package dns

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetSteeringPolicyRequest wrapper for the GetSteeringPolicy operation
type GetSteeringPolicyRequest struct {

	// The OCID of the target steering policy.
	SteeringPolicyId *string `mandatory:"true" contributesTo:"path" name:"steeringPolicyId"`

	// The `If-None-Match` header field makes the request method conditional on
	// the absence of any current representation of the target resource, when
	// the field-value is `*`, or having a selected representation with an
	// entity-tag that does not match any of those listed in the field-value.
	IfNoneMatch *string `mandatory:"false" contributesTo:"header" name:"If-None-Match"`

	// The `If-Modified-Since` header field makes a GET or HEAD request method
	// conditional on the selected representation's modification date being more
	// recent than the date provided in the field-value.  Transfer of the
	// selected representation's data is avoided if that data has not changed.
	IfModifiedSince *string `mandatory:"false" contributesTo:"header" name:"If-Modified-Since"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetSteeringPolicyRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetSteeringPolicyRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetSteeringPolicyRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetSteeringPolicyResponse wrapper for the GetSteeringPolicy operation
type GetSteeringPolicyResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The SteeringPolicy instance
	SteeringPolicy `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to
	// contact Oracle about a particular request, please provide the request
	// ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// The current version of the resource, ending with a
	// representation-specific suffix. This value may be used in If-Match
	// and If-None-Match headers for later requests of the same resource.
	ETag *string `presentIn:"header" name:"etag"`

	// Flag to indicate whether or not the object was modified.  If this is true,
	// the getter for the object itself will return null.  Callers should check this
	// if they specified one of the request params that might result in a conditional
	// response (like 'if-match'/'if-none-match').
	IsNotModified bool
}

func (response GetSteeringPolicyResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetSteeringPolicyResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
