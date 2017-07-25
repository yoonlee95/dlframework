// Code generated by go-swagger; DO NOT EDIT.

package predictor

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"
)

// PredictURL generates an URL for the predict operation
type PredictURL struct {
	FrameworkName string
	ModelName     string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *PredictURL) WithBasePath(bp string) *PredictURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *PredictURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *PredictURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/v1/framework/{framework_name}/model/{model_name}/predict"

	frameworkName := o.FrameworkName
	if frameworkName != "" {
		_path = strings.Replace(_path, "{framework_name}", frameworkName, -1)
	} else {
		return nil, errors.New("FrameworkName is required on PredictURL")
	}
	modelName := o.ModelName
	if modelName != "" {
		_path = strings.Replace(_path, "{model_name}", modelName, -1)
	} else {
		return nil, errors.New("ModelName is required on PredictURL")
	}
	_basePath := o._basePath
	result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *PredictURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *PredictURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *PredictURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on PredictURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on PredictURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *PredictURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
