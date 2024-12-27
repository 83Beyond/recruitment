package network

type Request struct {
	Name          string
	URL           string
	Method        string
	Params        map[string]string
	Headers       map[string]string
	Cookies       map[string]string
	PostData      map[string]any
	CurrentPage   int
	StartPage     int
	Offset        int
	ScrapePage    int
	Timestamp     string
	PageKey       string
	TimeKey       string
	CitySelectKey string
	TypeSelectKey string
	KeywordKey    string
	TypeSelectMap map[string]any
	Parse         func(response Response)
}
