package ws

import "fmt"

type Dispatch struct {
}

func NewDisPatch() *Dispatch {
	return &Dispatch{}
}

func (d *Dispatch) Execute(cmd BaseCmd) ([]byte, error) {
	switch cmd.Uri {
	case BlockInfoCmd:
		return d.doBlockInfo()
	case ChainInfoCmd:
		return d.doChainInfo()
	default:
		return nil, fmt.Errorf("%s unsupport error  ", cmd.Uri)
	}
}

func (d *Dispatch) doBlockInfo() ([]byte, error) {
	// todo
	return NewRespErr(BaseCmd{})
}

func (d *Dispatch) doChainInfo() ([]byte, error) {
	// todo
	return NewRespErr(BaseCmd{})
}
