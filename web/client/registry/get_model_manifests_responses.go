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

// GetModelManifestsReader is a Reader for the GetModelManifests structure.
type GetModelManifestsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetModelManifestsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetModelManifestsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetModelManifestsOK creates a GetModelManifestsOK with default headers values
func NewGetModelManifestsOK() *GetModelManifestsOK {
	return &GetModelManifestsOK{}
}

/*GetModelManifestsOK handles this case with default header values.

GetModelManifestsOK get model manifests o k
*/
type GetModelManifestsOK struct {
	Payload *models.DlframeworkGetModelManifestsResponse
}

func (o *GetModelManifestsOK) Error() string {
	return fmt.Sprintf("[GET /v1/models][%d] getModelManifestsOK  %+v", 200, o.Payload)
}

func (o *GetModelManifestsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DlframeworkGetModelManifestsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}