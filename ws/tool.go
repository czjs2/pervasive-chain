package ws

import "pervasive-chain/model"

var cmdList = []string{TransferCmd}

func ValidCmd(cmd model.Cmd) bool {
	for i := 0; i < len(cmdList); i++ {
		if cmd.Body.Cmd.Key == cmdList[i] {
			return true
		}
	}
	return false
}
