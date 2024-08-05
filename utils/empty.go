package utils

// IsEmpty 判断是否为空
func IsEmpty(v any) bool {
	switch v.(type) {
	case []string:
		return len(v.([]string)) == 0
	case []int:
		return len(v.([]int)) == 0
	case []int32:
		return len(v.([]int32)) == 0
	case []int64:
		return len(v.([]int64)) == 0
	case []uint:
		return len(v.([]uint)) == 0
	case []uint32:
		return len(v.([]uint32)) == 0
	case []uint64:
		return len(v.([]uint64)) == 0
	case []float32:
		return len(v.([]float32)) == 0
	case []float64:
		return len(v.([]float64)) == 0
	case string:
		return v == ""
	case int:
		return v == 0
	case int32:
		return v == 0
	case int64:
		return v == 0
	case uint:
		return v == 0
	case uint32:
		return v == 0
	case uint64:
		return v == 0
	case float32:
		return v == 0
	case float64:
		return v == 0
	case nil:
		return true
	default:
		return false
	}
}
