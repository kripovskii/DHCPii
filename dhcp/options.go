package dhcp

// DHCPOption представляет одну DHCP-опцию в формате код + значение.
type DHCPOption struct {
	Code  byte
	Value []byte
}

// Option хранит DHCP-опции по их кодам.
type Option map[byte]DHCPOption

// Get возвращает опцию по коду.
func (opt Option) Get(code byte) (DHCPOption, bool) {
	option, ok := opt[code]
	return option, ok
}

// Set добавляет или заменяет опцию.
func (opt Option) Set(option DHCPOption) {
	opt[option.Code] = option
}
