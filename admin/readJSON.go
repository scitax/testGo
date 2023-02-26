package admin

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var defaultJsonPath string = "admin/admin files/environment.json"

type Rules struct {
	Vms      []Vm      `json:"vms"`
	Fv_rules []Fv_rule `json:"fv_rule"`
}
type Vm struct {
	Vm_id string `json:"vm_id"`
	Name  string `json:"name"`
	Tags  string `json:"tags"`
}
type Fv_rule struct {
	Fv_id      string `json:"fv_id"`
	Source_tag string `json:"source_tag"`
	Dest_tag   string `json:"dest_tag"`
}

func ReadJSON(path string) (*Rules, error) {
	if path == "" {
		path = defaultJsonPath
	}

	jsonFile, err := os.Open(path)
	defer jsonFile.Close()

	if err != nil {
		return nil, err
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var rulesData Rules

	json.Unmarshal(byteValue, &rulesData)

	return &rulesData, nil
}
