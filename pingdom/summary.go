package pingdom

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
)

// SummaryService provides an interface to Pingdom summary
type SummaryService struct {
	client *Client
}

// Summary is an interface representing a pingdom summary.
// Specific summary types should implement the methods of this interface
type Summary interface {
	Valid() error
}

// Return a list of performance summary from Pingdom.
// This returns type SummaryResponse rather than Summary since the
// pingdom API has different respresentation of a summary.
func (cs *SummaryService) Performance(id int, params ...map[string]string) (*SummaryResponse, error) {
	param := map[string]string{}
	if len(params) == 1 {
		param = params[0]
	}
	req, err := cs.client.NewRequest("GET", "/summary.performance/"+strconv.Itoa(id), param)
	if err != nil {
		return nil, err
	}

	resp, err := cs.client.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := validateResponse(resp); err != nil {
		return nil, err
	}

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)
	m := &summaryDetailsJsonResponse{}
	err = json.Unmarshal([]byte(bodyString), &m)

	return m.Summary, err
}
