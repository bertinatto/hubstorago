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

func newClientWithResponse(code int, js string) *Client {
	c := Client{}
	c.HTTPClient.Transport = &staticResp{
		Resp: &http.Response{
			StatusCode: code,
			Body:       ioutil.NopCloser(strings.NewReader(js)),
		},
	}
	return &c

}

func TestBadStatus(t *testing.T) {
	c := newClientWithResponse(404, `[{"id":"1234"}]`)
	_, err := c.Items("1111111/1/1")
	if !strings.Contains(err.Error(), "Bad status") {
		t.Errorf("Bad status error not caught")
	}
}

func TestBadJsonReturn(t *testing.T) {
	c := newClientWithResponse(200, `[[{"id":"1234"}]`)
	_, err := c.Items("1111111/1/1")
	if !strings.Contains(err.Error(), "Bad JSON") {
		t.Errorf("Bad JSON response not caught")
	}
}

func TestItems(t *testing.T) {
	c := newClientWithResponse(200, `[{"id":"1234"}]`)
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

func TestJobQ(t *testing.T) {
	c := newClientWithResponse(200, `[
									{
										"close_reason": "finished",
										"elapsed": 1747455900,
										"finished_time": 1469729888547,
										"items": 66,
										"key": "1111111/26/355",
										"logs": 142,
										"pages": 52,
										"pending_time": 1469729637028,
										"running_time": 1469729639812,
										"spider": "spider_test",
										"state": "finished",
										"ts": 1469729650804,
										"version": "test"
									}
								]`)
	jobs, _ := c.JobQ("1111111/26/355")
	j := (*jobs)[0]
	if j.CloseReason != "finished" {
		t.Errorf("Wrong close_reason: %s", j.CloseReason)
	}
}

func TestRequests(t *testing.T) {
	c := newClientWithResponse(200, `[
									{
										"duration": 251,
										"fp": "f1f733a6cefcc6212b7692b62145b0cdcc802d8f",
										"method": "GET",
										"rs": 37,
										"status": 200,
										"time": 1469729650376
									}
								]`)
	requests, _ := c.Requests("1111111/8/2")
	r := (*requests)[0]
	if r.Duration != 251 {
		t.Errorf("Wrong duration: %s", r.Duration)
	}
}

func TestLogs(t *testing.T) {
	c := newClientWithResponse(200, `[{"time":1469729643312,"level":20,"message":"Log opened."}]`)
	logs, _ := c.Logs("1111111/8/2")
	l := (*logs)[0]
	if l.Time != 1469729643312 {
		t.Errorf("Wrong time: %s", l.Time)
	}

}

func TestUrlJoin(t *testing.T) {
	c := Client{}
	if url := c.urlJoin([]string{"a", "b"}); url != "https://storage.scrapinghub.com/a/b" {
		t.Errorf("Wrong url: %s", url)
	}
	if url := c.urlJoin([]string{"/a/", "/b/"}); url != "https://storage.scrapinghub.com/a/b" {
		t.Errorf("Wrong url: %s", url)
	}
	if url := c.urlJoin([]string{"/a/", "/", "/b/"}); url != "https://storage.scrapinghub.com/a/b" {
		t.Errorf("Wrong url: %s", url)
	}

}
