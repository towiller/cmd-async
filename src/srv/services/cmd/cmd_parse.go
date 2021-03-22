package asynccmd

import "fmt"

type CmdsContent struct {
	Topic      string                   `json:"topic"`
	FatalTopic string                   `json:"fatal_topic"`
	Cmds       map[string][]interface{} `json:"cmds"`
}

func ExeCmds(cnt CmdsContent) {
	for cmdName, cmdList := range cnt.Cmds {
		fmt.Println(cmdName, cmdList)
		//switch cmdName {
		//case "do_grpc":
		//}
	}
}

func ParseGetVal(retJson interface{}, ret interface{}) {
	var ok bool
	ret, ok = retJson.(*CmdsContent)
	if !ok {
		ret = nil
	}
}
