// Code generated by go-swagger; DO NOT EDIT.

package api_catalogs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/tsrob/l7portal-go/models"
)

// GetAPICatalogSwaggerWithContentUsingGETReader is a Reader for the GetAPICatalogSwaggerWithContentUsingGET structure.
type GetAPICatalogSwaggerWithContentUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPICatalogSwaggerWithContentUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPICatalogSwaggerWithContentUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAPICatalogSwaggerWithContentUsingGETOK creates a GetAPICatalogSwaggerWithContentUsingGETOK with default headers values
func NewGetAPICatalogSwaggerWithContentUsingGETOK() *GetAPICatalogSwaggerWithContentUsingGETOK {
	return &GetAPICatalogSwaggerWithContentUsingGETOK{}
}

/*
GetAPICatalogSwaggerWithContentUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type GetAPICatalogSwaggerWithContentUsingGETOK struct {
	Payload *models.APIAsset
}

// IsSuccess returns true when this get Api catalog swagger with content using g e t o k response has a 2xx status code
func (o *GetAPICatalogSwaggerWithContentUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get Api catalog swagger with content using g e t o k response has a 3xx status code
func (o *GetAPICatalogSwaggerWithContentUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Api catalog swagger with content using g e t o k response has a 4xx status code
func (o *GetAPICatalogSwaggerWithContentUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Api catalog swagger with content using g e t o k response has a 5xx status code
func (o *GetAPICatalogSwaggerWithContentUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get Api catalog swagger with content using g e t o k response a status code equal to that given
func (o *GetAPICatalogSwaggerWithContentUsingGETOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get Api catalog swagger with content using g e t o k response
func (o *GetAPICatalogSwaggerWithContentUsingGETOK) Code() int {
	return 200
}

func (o *GetAPICatalogSwaggerWithContentUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /api-management/1.0/api-catalogs/{apiUuid}/assets/swagger][%d] getApiCatalogSwaggerWithContentUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetAPICatalogSwaggerWithContentUsingGETOK) String() string {
	return fmt.Sprintf("[GET /api-management/1.0/api-catalogs/{apiUuid}/assets/swagger][%d] getApiCatalogSwaggerWithContentUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetAPICatalogSwaggerWithContentUsingGETOK) GetPayload() *models.APIAsset {
	return o.Payload
}

func (o *GetAPICatalogSwaggerWithContentUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIAsset)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
