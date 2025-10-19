package logicaltest

import "testing"

type testStruct struct {
	Name string
	Text string
	Want string
}

func TestReverse(t *testing.T) {
	testTable := []testStruct{
		{"test melati", "italem irad irigayaj", "melati dari jayagiri"},
		{"test badai", "iadab itsap ulalreb", "badai pasti berlalu"},
		{"test bulan", "nalub kusutret gnalali", "bulan tertusuk ilalang"},
	}

	for _, tt := range testTable {
		t.Run(tt.Name, func(t *testing.T) {
			result := Reverse(tt.Text)
			if result != tt.Want {
				t.Errorf("expected %s, found %s", tt.Want, result)
			}
		})
	}
	// text := "italem irad irigayaj"
	// want := "melati dari jayagiri"
	// result := Reverse(text)
	// if result != want {
	// 	t.Errorf("expected %s, found %s", want, result)
	// }
}
