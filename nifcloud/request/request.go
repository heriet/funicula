package request

import (
	"encoding/xml"
	"errors"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/heriet/funicula/nifcloud"
	"github.com/heriet/funicula/nifcloud/signature/v2"
)

// Request is request of service api
type Request struct {
	Config       *nifcloud.Config
	Operation    *Operation
	HTTPRequest  *http.Request
	HTTPResponse *http.Response

	Params interface{}
	Body   interface{}
	Data   interface{}

	Time       time.Time
	ExpireTime time.Duration

	Error error
}

// Operation is operation of service api
type Operation struct {
	Name       string
	HTTPMethod string
	HTTPPath   string
}

// New creates Request
func New(cfg *nifcloud.Config, operation *Operation, params map[string]string, data interface{}) *Request {
	method := operation.HTTPMethod
	if method == "" {
		method = "POST"
	}

	// TODO only v2
	params["Action"] = operation.Name
	params["AccessKeyId"] = cfg.Credential.AccessKeyId
	params["SignatureMethod"] = "HmacSHA256"
	params["SignatureVersion"] = "2"

	host := generateEndpointHostname(cfg.Endpoint)

	values := url.Values{}
	for key, value := range params {
		if value != "" {
			values.Add(key, value)
		}
	}

	// nifcloud sigunature needs spaces convert to %20
	encodedValues := strings.Replace(values.Encode(), "+", "%20", -1)

	signature := v2.CalcSignature(cfg.Credential, method, operation.HTTPPath, encodedValues, host)
	values.Add("Signature", signature)

	httpReq, _ := http.NewRequest(method, "", strings.NewReader(values.Encode()))

	var err error
	httpReq.URL, err = url.Parse(cfg.Endpoint + operation.HTTPPath)
	if err != nil {
		httpReq.URL = &url.URL{}
		err = errors.New("InvalidEndpointURL")
	}

	req := &Request{
		Config:      cfg,
		Time:        time.Now(),
		ExpireTime:  0,
		Operation:   operation,
		HTTPRequest: httpReq,
		Body:        nil,
		Params:      params,
		Data:        data,
		Error:       err,
	}

	return req
}

func generateEndpointHostname(endpoint string) string {
	rep := regexp.MustCompile(`^https?://`)
	return rep.ReplaceAllString(endpoint, "")
}

// Send will request api
func (req *Request) Send() error {

	client := &http.Client{
		Timeout: time.Duration(10) * time.Second,
	}

	resp, err := client.Do(req.HTTPRequest)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	req.HTTPResponse = resp

	// TODO: need handler by any services
	if resp.StatusCode != 200 {
		return errors.New("RequestError")
	}

	err = xml.NewDecoder(resp.Body).Decode(&req.Data)
	if err != nil {
		return err
	}

	return nil
}
