package chat

type ConnGateway interface {
	ReadConn() (interface{}, error)
}
