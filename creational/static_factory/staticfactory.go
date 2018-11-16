// 需要借助静态方法来实现
package static_factory

type FormatType int

const (
	FORMAT_NUMBER_TYPE FormatType = iota
	FORMAT_STRING_TYPE
)

func Factory(formatterType FormatType) Formatter {
	if formatterType == FORMAT_NUMBER_TYPE {
		return &FormatNumber{}
	}

	if formatterType == FORMAT_STRING_TYPE {
		return &FormatString{}
	}
	return nil
}
