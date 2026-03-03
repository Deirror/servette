package transport

type TransportType = string

const (
	TCPKey="TCP"	
	UDSKey="UDS"
)

func IsValidType(t TransportType) bool {
	return t == TCPKey || t == UDSKey
}
