package chat

type ConnGateway interface {
	ReadConn() ([]byte, error)
	HandleError(error)
}
