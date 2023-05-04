// Code generated by go-swagger; DO NOT EDIT.

package apis

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetAPIBundleUsingGETReader is a Reader for the GetAPIBundleUsingGET structure.
type GetAPIBundleUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIBundleUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIBundleUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAPIBundleUsingGETOK creates a GetAPIBundleUsingGETOK with default headers values
func NewGetAPIBundleUsingGETOK() *GetAPIBundleUsingGETOK {
	return &GetAPIBundleUsingGETOK{}
}

/*
GetAPIBundleUsingGETOK describes a response with status code 200, with default header values.

Success
*/
type GetAPIBundleUsingGETOK struct {
	Payload string
}

// IsSuccess returns true when this get Api bundle using g e t o k response has a 2xx status code
func (o *GetAPIBundleUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get Api bundle using g e t o k response has a 3xx status code
func (o *GetAPIBundleUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Api bundle using g e t o k response has a 4xx status code
func (o *GetAPIBundleUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Api bundle using g e t o k response has a 5xx status code
func (o *GetAPIBundleUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get Api bundle using g e t o k response a status code equal to that given
func (o *GetAPIBundleUsingGETOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get Api bundle using g e t o k response
func (o *GetAPIBundleUsingGETOK) Code() int {
	return 200
}

func (o *GetAPIBundleUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /api-management/1.0/apis/{apiUuid}/bundle][%d] getApiBundleUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetAPIBundleUsingGETOK) String() string {
	return fmt.Sprintf("[GET /api-management/1.0/apis/{apiUuid}/bundle][%d] getApiBundleUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetAPIBundleUsingGETOK) GetPayload() string {
	return o.Payload
}

func (o *GetAPIBundleUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
