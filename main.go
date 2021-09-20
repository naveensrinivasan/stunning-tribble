package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	fail := false
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	l := log.New(os.Stderr, "", 0)
	d := log.New(os.Stdout, "", 0)
	for _, dep := range getGoDeps(string(bytes)) {
		vul, err := getOSV(dep)
		if err != nil {
			l.Println(err)
		}
		if vul != "" {
			d.Println(vul)
			fail = true
		}
	}
	if fail {
		os.Exit(1)
	}
}

type Dep struct {
	Name    string
	Version string
}

// GetGoDeps returns go repo dependencies.
func getGoDeps(out string) []Dep {
	deps := []Dep{}
	for _, l := range strings.Split(out, "\n") {
		l = strings.Trim(l, "'")
		replace := strings.Split(l, "_")
		if len(replace) > 1 {
			if replace[1] == "<nil>" {
				data := strings.Split(replace[0], "@")
				deps = append(deps, Dep{data[0], strings.TrimPrefix(data[1], "v")})
			} else {
				data := strings.Split(replace[1], " ")
				if len(data) > 1 {
					deps = append(deps, Dep{data[0], strings.TrimPrefix(data[1], "v")})
				}
			}
		} else {
			data := strings.Split(replace[0], "@")
			if len(data) > 1 {
				deps = append(deps, Dep{data[0], strings.TrimPrefix(data[1], "v")})
			}
		}
	}
	return deps
}

func getOSV(d Dep) (string, error) {
	// curl -X POST -H "Content-Type: application/json"  -d '{"version":"3.2.1+incompatible", "package": {"name": "github.com/golang-jwt/jwt", "ecosystem": "Go"}}' "https://api.osv.dev/v1/query"

	type Package struct {
		Name      string `json:"name"`
		Ecosystem string `json:"ecosystem"`
	}
	type Payload struct {
		Version string  `json:"version"`
		Package Package `json:"package"`
	}

	data := Payload{
		Version: d.Version,
		Package: Package{Name: d.Name, Ecosystem: "Go"},
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "https://api.osv.dev/v1/query", body)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	output, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if len(output) > 2 {
		return string(output), nil
	}
	return "", nil
}
