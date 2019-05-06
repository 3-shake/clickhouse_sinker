package util

import (
	"github.com/3-shake/clickhouse_sinker/column"
	"github.com/3-shake/clickhouse_sinker/model"
)

//这里对metric的value类型，只有三种情况， （float64，string，map[string]interface{})
func GetValueByType(metric model.Metric, cwt *model.ColumnWithType) interface{} {
	swType := switchType(cwt.Type)
	switch swType {
	case "int":
		return metric.GetInt(cwt.Name)
	case "float":
		return metric.GetFloat(cwt.Name)
	case "string":
		return metric.GetString(cwt.Name)
	default:
		val := metric.Get(cwt.Name)
		col := column.GetColumnByName(cwt.Type)
		if val == nil {
			return col.DefaultValue()
		}
		return col.GetValue(val)
	}
}

func switchType(typ string) string {
	switch typ {
	case "UInt8", "UInt16", "UInt32", "UInt64", "Int8", "Int16", "Int32", "Int64":
		return "int"
	case "String", "FixString":
		return "string"
	case "UUID":
		return "uuid"
	case "Float32", "Float64":
		return "float"
	default:
		panic("unsupport type " + typ)
	}
}
