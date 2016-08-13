package hubstorago

// logEntry is a single "line" of the log.
type logEntry struct {
	Level   int    `json:"level"`
	Message string `json:"message"`
	Time    int    `json:"time"`
}

// LogsData is the complete set of log lines.
type LogsData []logEntry

// RequestEntry is a single request made by a spider.
type RequestEntry struct {
	Duration int    `json:"duration"`
	Fp       string `json:"fp"`
	Method   string `json:"method"`
	Rs       int    `json:"rs"`
	Status   int    `json:"status"`
	Time     int    `json:"time"`
	Url      string `json:"url"`
}

// RequestsData represents all requests made by a spider in a specific job.
type RequestsData []RequestEntry

// jobqEntry is a single job from a queue.
type jobqEntry struct {
	CloseReason  string `json:"close_reason"`
	Elapsed      int    `json:"elapsed"`
	FinishedTime int    `json:"finished_time"`
	Key          string `json:"key"`
	Logs         int    `json:"logs"`
	Pages        int    `json:"pages"`
	PendingTime  int    `json:"pending_time"`
	RunningTime  int    `json:"running_time"`
	Spider       string `json:"spider"`
	State        string `json:"state"`
	Ts           int    `json:"ts"`
	Version      string `json:"version"`
}

// JobQData is the list of all jobs returned by Hubstorage.
type JobQData []jobqEntry

// collectionEntry is a single Collection record.
type collectionEntry struct {
	Key   string `json:"_key"`
	Value string `json:"value"`
}

// CollectionsData represents the data fetched from the Collections endpoint,
// which can be more than one entry.
type CollectionsData []collectionEntry
