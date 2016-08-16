package hubstorago

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

type staticResp struct {
	Resp *http.Response
}

func (t *staticResp) RoundTrip(req *http.Request) (*http.Response, error) {
	return t.Resp, nil
}

func TestItems(t *testing.T) {
	c := Client{}
	c.HTTPClient.Transport = &staticResp{
		Resp: &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`[{"id":"1234"}]`)),
		},
	}

	// Get the item
	raw_items, _ := c.Items("1111111/1/1")
	items := (*raw_items).([]interface{})

	// Just one, right?
	if len(items) != 1 {
		t.Errorf("Items lenght: %d", len(items))
	}

	// Does it have the correct value?
	item := items[0].(map[string]interface{})
	if item["id"] != "1234" {
		t.Errorf("Wrong value: %s", item["id"])
	}
}
