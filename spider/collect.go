package spider

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/83Beyond/recruitment/models"
	"github.com/83Beyond/recruitment/network"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type Fetcher struct {
	Name       string
	Url        string
	Method     string
	Params     map[string]string
	Headers    map[string]string
	Cookies    map[string]string
	IsJson     bool
	PostData   map[string]string
	ScrapePage int
	Timestamp  string
	PageKey    string
	TimeKey    string
}

func StartRequest(r network.Request) network.Response {

	var req *http.Request
	var err error
	if r.Method == "GET" {
		if r.Params != nil {
			params := url.Values{}
			for k, v := range r.Params {
				params.Add(k, v)
			}
			if r.TimeKey != "" {
				timestamp := time.Now().UnixNano() / int64(time.Millisecond)
				params.Add(r.TimeKey, strconv.Itoa(int(timestamp)))
				params.Add("_", strconv.Itoa(int(timestamp)))
			}
			params.Add(r.PageKey, strconv.Itoa(r.CurrentPage))

			req, err = http.NewRequest(r.Method, r.URL+"?"+params.Encode(), nil)
		} else {
			req, err = http.NewRequest(r.Method, r.URL, nil)
		}

	} else {
		if models.SelectedType != "" {
			if strings.Index(r.TypeSelectKey, ".") != -1 {
				k := strings.Split(r.TypeSelectKey, ".")
				index, _ := strconv.Atoi(k[1])
				r.PostData[k[0]].([]string)[index] = models.SelectedType
			}
		}
		if models.Keyword != "" {
			r.PostData[r.KeywordKey] = models.Keyword
		}
		r.PostData[r.PageKey] = r.CurrentPage
		jsonData, _ := json.Marshal(r.PostData)
		fmt.Println(string(jsonData))
		req, err = http.NewRequest(r.Method, r.URL, bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
	}

	if r.Headers != nil {
		for k, v := range r.Headers {
			req.Header.Set(k, v)
		}
	}
	if r.Cookies != nil {
		for key, value := range r.Cookies {
			cookie := &http.Cookie{
				Name:  key,
				Value: value,
			}
			req.AddCookie(cookie)
		}
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("fetch failed %v", err)
	}
	defer resp.Body.Close()

	bodyReader := bufio.NewReader(resp.Body)
	e := DetermineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())

	body, err := ioutil.ReadAll(utf8Reader)

	if err != nil {
		fmt.Printf("read body failed %v", err)
	}
	return network.Response{
		//Website:    r.Name,
		//Text:       string(body),
		//StatusCode: resp.StatusCode,
		Source:  r.Name,
		Content: body,
	}
}

func DetermineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)

	if err != nil {
		fmt.Printf("fetch error: %v", err)
		return unicode.UTF8
	}

	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
