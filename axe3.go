package main
import (
    "net/http"
    "regexp"
    "io/ioutil"
    "encoding/json"
    "net/http/cookiejar"
    "fmt"
)

type 里長 struct {
    Town string       `json:"town"`
    Village string    `json:"village"`
    Name string       `json:"name"`
}

func parseUrl(client http.Client, url string) (string, error) {
    resp, err := client.Get(url)

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
    options := cookiejar.Options{
	PublicSuffixList: nil,
    }
    jar, _ := cookiejar.New(&options)
    client := http.Client{Jar: jar}
    var page = 1
    url := "http://axe-level-1.herokuapp.com/lv3"
    r := regexp.MustCompile(`(?m)<tr>\s*<td>([^<]+)</td>\s*<td>([^<]+)</td>\s*<td>([^<]+)</td>\s*</tr>`)
    var result = []里長{}
    for page < 77 {
	html, err := parseUrl(client, url)
	url = "http://axe-level-1.herokuapp.com/lv3?page=next"
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
