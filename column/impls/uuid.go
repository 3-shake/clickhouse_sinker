package impls

type UUIDColumn struct {
}

const zeroValue = "00000000-0000-0000-0000-000000000000"

func (c *UUIDColumn) Name() string {
	return "UUID"
}

func (c *UUIDColumn) DefaultValue() interface{} {
	return zeroValue
}

func NewUUIDColumn() *UUIDColumn {
	return &UUIDColumn{}
}

func (c *UUIDColumn) GetValue(val interface{}) interface{} {
	if v, ok := val.(string); ok && v != "" {
		return v
	}
	return zeroValue
}
