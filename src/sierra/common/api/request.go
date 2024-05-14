package api

import (
	"bytes"
	"fmt"
	"github.com/luminanceaudio/sierra/src/sierra/common/config"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/sirupsen/logrus"
)

func SendRequest[T any, R any](client *http.Client, authHeader, method, endpoint string, request *T) (*R, *Error) {
	var requestBody []byte
	var err error
	if request != nil {
		requestBody, err = Marshal(*request)
		if err != nil {
			return nil, NewInternalError(err, "failed to unmarshal request for http send")
		}
	}

	return sendRawRequestForJSON[R](client, authHeader, method, endpoint, "application/json", bytes.NewBuffer(requestBody))
}

func SendDownloadRequest[T any](client *http.Client, authHeader, method, endpoint string, request *T) ([]byte, *Error) {
	var requestBody []byte
	var err error
	if request != nil {
		requestBody, err = Marshal(*request)
		if err != nil {
			return nil, NewInternalError(err, "failed to unmarshal request for http send")
		}
	}

	return sendRawRequest(client, authHeader, method, endpoint, "application/json", bytes.NewBuffer(requestBody))
}

func SendMultipartFile[T any](client *http.Client, authHeader string, method string, endpoint string, data []byte) (*T, *Error) {
	buf := bytes.NewBuffer(data)

	// Create a new multipart writer
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Add the file to the request
	part, err := writer.CreateFormFile("file", "data.bin")
	if err != nil {
		return nil, NewInternalError(err, "failed creating form file")
	}
	_, err = io.Copy(part, buf)
	if err != nil {
		return nil, NewInternalError(err, "failed copying form file")
	}

	// Close the multipart writer
	err = writer.Close()
	if err != nil {
		return nil, NewInternalError(err, "failed closing writer for multipart file")
	}

	return sendRawRequestForJSON[T](client, authHeader, method, endpoint, writer.FormDataContentType(), body)
}

func sendRawRequestForJSON[T any](client *http.Client, authHeader, method, endpoint string, contentType string, buffer *bytes.Buffer) (*T, *Error) {
	body, aErr := sendRawRequest(client, authHeader, method, endpoint, contentType, buffer)

	var response T
	if len(body) != 0 {
		err := Unmarshal(body, &response)
		if err != nil {
			return nil, NewInternalError(err, "failed getting response body")
		}
	}

	return &response, aErr
}

func sendRawRequest(client *http.Client, authHeader, method, endpoint string, contentType string, buffer *bytes.Buffer) ([]byte, *Error) {
	logrus := logrus.WithField("method", method).WithField("endpoint", endpoint)

	req, err := http.NewRequest(method, fmt.Sprintf("%s/api/v1/%s", config.InternalIngressUrl, endpoint), buffer)
	if err != nil {
		return nil, NewInternalError(err, "failed creating new request")
	}
	req.Header.Set("Content-Type", contentType)

	if authHeader != "" {
		req.Header.Set("Authorization", authHeader)
	}

	logrus.Info("Sending request")
	resp, err := client.Do(req)
	if err != nil {
		return nil, NewInternalError(err, "failed sending http request")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, NewInternalError(err, "failed reading response body")
	}

	if resp.StatusCode != http.StatusOK {
		var aErr Error
		if len(body) != 0 {
			err = Unmarshal(body, &aErr)
			if err != nil {
				logrus.WithField("method", method).WithField("endpoint", endpoint).Warn("couldn't unmarshal error body, continuing anyway")
			} else {
				return nil, &aErr
			}
		}
		return nil, newError(fmt.Errorf("unexpected status code"), "unexpected status code", resp.StatusCode)
	}

	logrus.Info("Got response")
	return body, nil
}
