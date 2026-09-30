package dhcp

import "fmt"

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
	DHCPDiscover          MessageType = 1
	DHCPOffer             MessageType = 2
	DHCPRequest           MessageType = 3
	DHCPDecline           MessageType = 4
	DHCPACK               MessageType = 5
	DHCPNAK               MessageType = 6
	DHCPRelease           MessageType = 7
	DHCPInform            MessageType = 8
	DHCPMessageTypeOption             = 53
)

// GetMessageType получает тип DHCP-сообщения из Option 53.
func (message Message) GetMessageType() (MessageType, error) {
	options := message.Options

	for position := 0; position < len(options); {
		code := options[position]
		position++

		if code == 0 {
			continue
		}
		if code == 255 {
			break
		}

		if position >= len(options) {
			return 0, fmt.Errorf(
				"option %d has no length byte",
				code,
			)
		}

		length := int(options[position])
		position++

		// Проверяем, что значение целиком находится в массиве.
		if position+length > len(options) {
			return 0, fmt.Errorf(
				"option %d is truncated: need %d bytes, have %d",
				code,
				length,
				len(options)-position,
			)
		}

		if code == DHCPMessageTypeOption {
			if length != 1 {
				return 0, fmt.Errorf(
					"message type option must contain exactly 1 byte, got %d",
					length,
				)
			}

			messageType := MessageType(options[position])

			if !isValidMessageType(messageType) {
				return 0, fmt.Errorf(
					"unknown DHCP message type: %d",
					messageType,
				)
			}

			return messageType, nil
		}

		position += length
	}

	return 0, fmt.Errorf("message type option %d not found", DHCPMessageTypeOption)
}

func isValidMessageType(messageType MessageType) bool {
	switch messageType {
	case DHCPDiscover,
		DHCPOffer,
		DHCPRequest,
		DHCPDecline,
		DHCPACK,
		DHCPNAK,
		DHCPRelease,
		DHCPInform:
		return true
	default:
		return false
	}
}
