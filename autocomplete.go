package main

import (
	"bytes"
	"encoding/json"
	"os/exec"
	"strings"
)

type AutocompleteResponse struct {
	Candidates []*Candidate
}

type Candidate struct {
	Caption string `json:"caption"`
	Snippet string `json:"snippet"`
	Meta    string `json:"meta"`
}

func autoComplete(fileName string, content []byte, offset string) *AutocompleteResponse {
	cmd := exec.Command(gocodePath, "-f=json", "--in="+goPath+fileName, "autocomplete", goPath+fileName, "c"+offset)

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		logger.Fatal(err)
	}

	result := &AutocompleteResponse{}

	var v []interface{}
	json.Unmarshal(out.Bytes(), &v)

	if len(v) == 0 {
		return nil
	}

	candidates := v[1].([]interface{})

	for _, gc := range candidates {
		m := gc.(map[string]interface{})
		c := &Candidate{}

		c.Meta = m["class"].(string)
		c.Caption = m["name"].(string)
		c.Snippet = m["name"].(string)
		typ := m["type"].(string)

		if strings.HasPrefix(typ, c.Meta) {
			c.Caption = c.Snippet + strings.TrimPrefix(typ, c.Meta)
		}

		result.Candidates = append(result.Candidates, c)
	}

	return result
}
