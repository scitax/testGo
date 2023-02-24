package admin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Rules struct {
	Vms      []Vm      `json:"vms"`
	Fv_rules []Fv_rule `json:"fv_rule"`
}
type Vm struct {
	Vm_id string   `json:"vm_id"`
	Name  string   `json:"name"`
	Tags  []string `json:"tags"`
}
type Fv_rule struct {
	Fv_id      string `json:"fv_id"`
	Source_tag string `json:"source_tag"`
	Dest_tag   string `json:"dest_tag"`
}

func ReadJSON() (*Rules, error) {

	jsonFile, err := os.Open("admin/admin files/environment.json")
	defer jsonFile.Close()

	if err != nil {
		return nil, err
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var users Rules

	json.Unmarshal(byteValue, &users)

	return &users, nil
}

func ReadRules() error {
	rules, err := ReadJSON()
	if err != nil {
		return err
	}
	fmt.Println(rules.Vms[0].Tags[0])
	return nil
}
