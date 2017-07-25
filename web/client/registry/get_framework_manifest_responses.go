// Code generated by go-swagger; DO NOT EDIT.

package registry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/rai-project/dlframework/web/models"
)

// GetFrameworkManifestReader is a Reader for the GetFrameworkManifest structure.
type GetFrameworkManifestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFrameworkManifestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetFrameworkManifestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetFrameworkManifestOK creates a GetFrameworkManifestOK with default headers values
func NewGetFrameworkManifestOK() *GetFrameworkManifestOK {
	return &GetFrameworkManifestOK{}
}

/*GetFrameworkManifestOK handles this case with default header values.

GetFrameworkManifestOK get framework manifest o k
*/
type GetFrameworkManifestOK struct {
	Payload *models.DlframeworkFrameworkManifest
}

func (o *GetFrameworkManifestOK) Error() string {
	return fmt.Sprintf("[GET /v1/framework/{framework_name}/info][%d] getFrameworkManifestOK  %+v", 200, o.Payload)
}

func (o *GetFrameworkManifestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DlframeworkFrameworkManifest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
