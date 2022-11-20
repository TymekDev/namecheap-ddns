package main

import (
	"encoding/xml"
	"io"
	"strings"
)

func parse(body io.Reader) (error, error) {
	dec := xml.NewDecoder(body)
	dec.CharsetReader = func(charset string, r io.Reader) (io.Reader, error) {
		return r, nil
	}

	var result apiResponse
	if err := dec.Decode(&result); err != nil {
		return nil, err
	}

	if result.ErrCount > 0 {
		return result, nil
	}

	return nil, nil
}

type apiResponse struct {
	ErrCount int `xml:"ErrCount"`
	Errors   struct {
		Err1  string `xml:"Err1,omitempty"`
		Err2  string `xml:"Err2,omitempty"`
		Err3  string `xml:"Err3,omitempty"`
		Err4  string `xml:"Err4,omitempty"`
		Err5  string `xml:"Err5,omitempty"`
		Err6  string `xml:"Err6,omitempty"`
		Err7  string `xml:"Err7,omitempty"`
		Err8  string `xml:"Err8,omitempty"`
		Err9  string `xml:"Err9,omitempty"`
		Err10 string `xml:"Err10,omitempty"`
	} `xml:"errors"`
}

var _ error = apiResponse{}

func (r apiResponse) Error() string {
	return strings.Join([]string{
		r.Errors.Err1,
		r.Errors.Err2,
		r.Errors.Err3,
		r.Errors.Err4,
		r.Errors.Err5,
		r.Errors.Err6,
		r.Errors.Err7,
		r.Errors.Err8,
		r.Errors.Err9,
		r.Errors.Err10,
	}[:r.ErrCount], " & ")
}
