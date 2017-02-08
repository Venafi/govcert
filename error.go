package govcert

import (
	"regexp"
	"strconv"

	"strings"

	"encoding/json"

	"github.com/davecgh/go-spew/spew"
)

// HTTPError extends error and represents an error with get back from the API
type HTTPError interface {
	error
	HTTPStatusCode() int
	HTTPStatusText() string
	ErrWithCode() (string, int)
	HasErrorCode() bool
	JsonError() ([]byte, error)
}

type HttpError struct {
	CmdError       string
	CmdErrorCode   int
	HttpStatusCode int
	HttpStatusText string
}

func (h *HttpError) Error() string {
	return h.CmdError
}

func (h *HttpError) ErrWithCode() (string, int) {
	return h.CmdError, h.CmdErrorCode
}

func (h *HttpError) JsonError() ([]byte, error) {

	strJSON := map[string]map[string]interface{}{
		"error": map[string]interface{}{
			"code":    h.CmdErrorCode,
			"message": h.CmdError,
		},
	}
	return json.Marshal(strJSON)

}

func (h *HttpError) HasErrorCode() bool {
	return h.CmdErrorCode > 0
}

func (h *HttpError) HTTPStatusCode() int {
	return h.HttpStatusCode
}

func (h *HttpError) HTTPStatusText() string {
	return h.HttpStatusText
}

func newHTTPError(out []byte) HTTPError {
	spew.Dump(out)
	hterr := &HttpError{}
	// hterr := new(httpError)
	re, _ := regexp.Compile("Status: ([0-9]+) (.+)")
	spew.Dump(hterr)
	matches := re.FindStringSubmatch(string(out))
	spew.Dump(matches)
	if len(matches) == 3 {
		hterr.HttpStatusCode, _ = strconv.Atoi(matches[1])
		hterr.HttpStatusText = matches[2]
		spew.Dump(hterr)
	} else {
		hterr.HttpStatusCode = 500
		spew.Dump(hterr)
	}

	spew.Dump(hterr.HttpStatusCode)

	re, _ = regexp.Compile("Error Code: ([0-9]+) Error: (.*)$")
	errMatches := re.FindStringSubmatch(string(out))
	if len(errMatches) < 3 {
		lineArr := strings.Split(string(out), "\n")
		lastLine := lineArr[len(lineArr)-1]
		hterr.CmdError = lastLine
		return hterr
	}
	hterr.CmdErrorCode, _ = strconv.Atoi(errMatches[1])
	hterr.CmdError = errMatches[2]
	return hterr
}
