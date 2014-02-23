package main
import (
    "net/http"
    "regexp"
    "io/ioutil"
    "encoding/json"
    "strconv"
    "fmt"
)

type All struct {
    Name string  `json:"name"`
    Grades Grade `json:"grades"`
}

type Grade struct {
    F1 int `json:"國語"`
    F2 int `json:"數學"`
    F3 int `json:"自然"`
    F4 int `json:"社會"`
    F5 int `json:"健康教育"`
}

func main() {
    resp, err := http.Get("http://axe-level-1.herokuapp.com/")

    if err != nil {
	// handle error
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)

    r := regexp.MustCompile(`(?m)<tr>\s*<td>([^<]+)</td>\s*<td>([^<]+)</td>\s*<td>([^<]+)</td>\s*<td>([^<]+)</td>\s*<td>([^<]+)</td>\s*<td>([^<]+)</td>\s*</tr>`)
    allStrings := r.FindAllStringSubmatch(string(body), -1)
    var result = make([]All, len(allStrings) - 1)
    for i, res := range allStrings[1:] {
	var ints = make([]int, len(res))
	for j, e := range res[2:] {
	    ints[j], _ = strconv.Atoi(e)
	}
	grade := &Grade{ints[0], ints[1], ints[2], ints[3], ints[4]}
	record := All{res[1], *grade}
	result[i] = record;
    }
    js, _ := json.Marshal(result)
    fmt.Printf("%s", js)
}
