package libs

import (
	"io/ioutil"
	"os/exec"
	"time"

	"gopkg.in/yaml.v2"
)

// ExecCmd 执行shell命令
func ExecCmd(c string) ([]byte, error) {
	cmd := exec.Command("sh", "-c", c)
	out, err := cmd.Output()
	return out, err
}

//Load load info by json
func Load(filePath string, v interface{}) error {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

//TodayTime today time
func TodayTime() time.Time {
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	return today
}

//LoadYaml load info by yaml
func LoadYaml(filePath string, v interface{}) error {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(data, v)
}
