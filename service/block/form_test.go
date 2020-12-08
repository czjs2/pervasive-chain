package block

import (
	"encoding/json"
	"fmt"
	"pervasive-chain/utils"
	"testing"
	"time"
)

func TestFrom(t *testing.T){
	form := ReportBlockForm{}
	bytes, e := json.Marshal(form)
	fmt.Println(string(bytes),e)

}

type Param struct {
	Time time.Time
}




func Test01(t *testing.T){
	bytes, s := utils.Base58Encode([]byte("sdfsdfsdfsdfsdfsdfssdfsdf"))
	fmt.Println(bytes,s)
}
