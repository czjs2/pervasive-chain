package dao

type IBlockDao interface {
	Insert() (interface{}, error)
	Query() (interface{}, error)
}

type ITransGroupDao interface {
	Insert() (interface{}, error)
	Query() (interface{}, error)
}

type ITransDao interface {
	Insert() (interface{}, error)
	Query() (interface{}, error)
}

type INodeBandDao interface {

}

type ITotalBandDao interface {

}

type INodeDao interface {

}
