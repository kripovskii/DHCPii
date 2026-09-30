package dhcp

import "errors"

// DHCPOption представляет одну DHCP-опцию в формате код + значение.
type DHCPOption struct {
	Code  byte
	Value []byte
}

var ErrNilOptions = errors.New("options map is nil")

// Options хранит DHCP-опции по их кодам.
type Options map[byte]DHCPOption

func NewOptions() Options {
	return make(Options)
}

// Get возвращает опцию по коду.
func (opt Options) Get(code byte) (DHCPOption, bool) {
	option, ok := opt[code]
	return option, ok
}

// Set добавляет новую опцию или заменяет существующую.
func (options Options) Set(option DHCPOption) error {
	if options == nil {
		return ErrNilOptions
	}

	options[option.Code] = option
	return nil
}
