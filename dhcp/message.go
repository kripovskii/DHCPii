package dhcp

const (
	// Фиксированные размеры полей в формате BOOTP/DHCPv4.
	HardwareAddressSize = 16
	ServerNameSize      = 64
	BootFileNameSize    = 128
)

// Message содержит фиксированную часть BOOTP-заголовка DHCPv4-сообщения.
//
// Типы полей соответствуют их размерам в сетевом формате. Options хранятся
// отдельно, поскольку имеют переменную длину и следуют после заголовка.
type Message struct {
	Op     uint8
	Htype  uint8
	Hlen   uint8
	Hops   uint8
	Xid    uint32
	Secs   uint16
	Flags  uint16
	Ciaddr [4]byte
	Yiaddr [4]byte
	Siaddr [4]byte
	Giaddr [4]byte
	Chaddr [HardwareAddressSize]byte
	Sname  [ServerNameSize]byte
	File   [BootFileNameSize]byte

	Options []byte
}

type MessageType byte

// MessageType определяет тип DHCP-сообщения.
// Значение передаётся в DHCP option 53, а не в фиксированной части BOOTP-заголовка.
const (
	DHCPDiscover MessageType = 1
	DHCPOffer    MessageType = 2
	DHCPRequest  MessageType = 3
	DHCPDecline  MessageType = 4
	DHCPACK      MessageType = 5
	DHCPNAK      MessageType = 6
	DHCPRelease  MessageType = 7
	DHCPInform   MessageType = 8
)
