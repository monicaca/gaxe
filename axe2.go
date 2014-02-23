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

func parseUrl(url string) (string, error) {
    resp, err := http.Get(url)

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
    for page < 13 {
	html, err := parseUrl(fmt.Sprintf("http://axe-level-1.herokuapp.com/lv2?page=%d",page))
	if err != nil {break}
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
