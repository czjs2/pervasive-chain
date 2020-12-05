package block

import (
	"pervasive-chain/form"
	"pervasive-chain/utils"
)

func ReportBlockValidate(req string) (form.IFormValidateInterface, error) {
	var blockFrom form.ReportBlockForm
	return &blockFrom, utils.Unmarshal(req, &blockFrom)
}
