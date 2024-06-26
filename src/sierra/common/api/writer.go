package api

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

type Writer struct {
	writer        http.ResponseWriter
	writtenHeader bool
}

func (w *Writer) Write(p []byte) (n int, err error) {
	return w.writer.Write(p)
}

func NewWriter(w http.ResponseWriter) *Writer {
	return &Writer{writer: w, writtenHeader: false}
}

func (w *Writer) WriteAt(p []byte, off int64) (int, error) {
	if off == 0 {
		w.writeHeader(http.StatusOK)
	}

	return w.writer.Write(p)
}
func (w *Writer) writeHeader(status int) {
	if w.writtenHeader {
		return
	}

	w.writer.WriteHeader(status)
	w.writtenHeader = true
}

func (w *Writer) writeError(aErr *Error) {
	log := logrus.WithError(aErr.Err).WithField("httpStatusCode", aErr.httpStatusCode).WithField("userFacingMessage", aErr.UserFacingMessage)
	if aErr.httpStatusCode == http.StatusInternalServerError {
		log.Error("internal api error")
	} else {
		log.Info("expected api error")
	}

	aErr = w.writeJSON(aErr.httpStatusCode, aErr)
	if aErr != nil {
		logrus.WithError(aErr.Err).Error("failed writing error JSON")
	}
}

func (w *Writer) writeJSON(httpStatus int, object any) *Error {
	w.writeHeader(httpStatus)

	if object != nil {
		b, err := Marshal(object)
		if err != nil {
			return NewInternalError(err, "failed marshaling response")
		}

		_, err = w.writer.Write(b)
		if err != nil {
			return NewInternalError(err, "failed writing marshalled response")
		}
	}

	return nil
}
