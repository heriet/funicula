package rdb

import (
	"fmt"

	"github.com/heriet/funicula/nifcloud/request"
)

// NiftyGetMetricStatisticsInput is NiftyGetMetricStatistics request parameters
type NiftyGetMetricStatisticsInput struct {
	Dimensions []Dimension
	MetricName string
	StartTime  string
	EndTime    string
}

// Dimension is a part of NiftyGetMetricStatistics request parameters
type Dimension struct {
	Name  string
	Value string
}

// NiftyGetMetricStatisticsOutput is NiftyGetMetricStatistics response
type NiftyGetMetricStatisticsOutput struct {
	Datapoints []Datapoint `xml:"NiftyGetMetricStatisticsResult>Datapoints>member"`
}

// Datapoint is a part of NiftyGetMetricStatistics response
type Datapoint struct {
	NiftyTargetName string  `xml:"NiftyTargetName"`
	Timestamp       string  `xml:"Timestamp"`
	Sum             float64 `xml:"Sum"`
	SampleCount     int64   `xml:"SampleCount"`
}

// NiftyGetMetricStatisticsRequest generates request of NiftyGetMetricStatistics
func (svc *RDB) NiftyGetMetricStatisticsRequest(input *NiftyGetMetricStatisticsInput) (req *request.Request, output *NiftyGetMetricStatisticsOutput) {
	op := &request.Operation{
		Name:       "NiftyGetMetricStatistics",
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &NiftyGetMetricStatisticsInput{}
	}

	output = &NiftyGetMetricStatisticsOutput{}

	// TODO: req = svc.Client.NewRequest(op, input, output)
	params := make(map[string]string)

	for i, dimension := range input.Dimensions {
		params[fmt.Sprintf("Dimensions.member.%d.Name", i+1)] = dimension.Name
		params[fmt.Sprintf("Dimensions.member.%d.Value", i+1)] = dimension.Value
	}

	params["MetricName"] = input.MetricName
	params["StartTime"] = input.StartTime
	params["EndTime"] = input.EndTime

	req = svc.Client.NewRequest(op, params, output)

	return
}

// NiftyGetMetricStatistics sends NiftyGetMetricStatistics
func (svc *RDB) NiftyGetMetricStatistics(input *NiftyGetMetricStatisticsInput) (*NiftyGetMetricStatisticsOutput, error) {
	req, output := svc.NiftyGetMetricStatisticsRequest(input)
	return output, req.Send()
}
