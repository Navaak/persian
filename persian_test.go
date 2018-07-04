package persian

import "testing"

var (
	f2IntCases = map[string]int{
		"۲۳۵":    235,
		"۸۷۳۴۵۶": 873456,
		"۹۰۱":    901,
	}
)

func TestF2Int(t *testing.T) {
	for key, value := range f2IntCases {
		res := F2Int(key)
		if value != res {
			t.Error("converting ", key, " = > ", value, " != ", res)
		}
	}
}
