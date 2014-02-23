package main
import (
    "net/http"
    "regexp"
    "io/ioutil"
    "encoding/json"
    "fmt"
)

type 里長 struct {
    Town string       `json:"town"`
    Village string    `json:"village"`
    Name string       `json:"name"`
}

func parseUrl(client http.Client, url string, refer string) (string, error) {
    req, err := http.NewRequest("GET", url, nil)
    req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.9; rv:27.0) Gecko/20100101 Firefox/27.0")

    if refer != ""{
	req.Header.Set("Referer", refer)
    }
    req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
    resp, err := client.Do(req)

    if err != nil {
	// handle error
	return "", err
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil { }
    return string(body), nil
}

func main() {
    var page = 1
    r := regexp.MustCompile(`(?m)<tr>\s*<td>([^<]+)</td>\s*<td>([^<]+)</td>\s*<td>([^<]+)</td>\s*</tr>`)
    var result = []里長{}
    client := http.Client{}
    lastUrl := ""
    for page < 25 {
	curUrl := fmt.Sprintf("http://axe-level-4.herokuapp.com/lv4/?page=%d",page)
	html, err := parseUrl(client, curUrl, lastUrl)
	if err != nil {break}
	lastUrl = curUrl
	allStrings := r.FindAllStringSubmatch(html, -1)
	for _, res := range allStrings[1:] {
	    record := 里長{res[1], res[2], res[3]}
	    result = append(result, record);
	}
	page += 1
    }
    js, _ := json.Marshal(result)
    fmt.Printf("%s", js)
}
