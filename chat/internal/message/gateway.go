package message

type PubGateway interface {
	PublishMessage(data []byte) error
}
