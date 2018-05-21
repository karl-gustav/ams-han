package decode_mbus

import (
	"testing"
	"time"
)

func TestInt(t *testing.T) {
	v := struct {
		Data int
	}{}

	const expect = 16843009
	err := MarshalItems([]byte{0x06, 0x01, 0x01, 0x01, 0x01}, &v)
	if err != nil {
		t.Errorf("Didn't expect error... %v", err)
	}
	if v.Data != expect {
		t.Errorf("Expcted %v, was %v", expect, v.Data)
	}
}

func TestString(t *testing.T) {
	v := struct {
		Data string
	}{}

	const expect = "Help me!"
	err := MarshalItems(append([]byte{0x09, 0x08}, []byte(expect)...), &v)
	if err != nil {
		t.Errorf("Didn't expect error... %v", err)
	}
	if v.Data != expect {
		t.Errorf("Data was `%v`, expected `%v`", v.Data, expect)
	}
}

func TestTime(t *testing.T) {
	v := struct {
		Data time.Time
	}{}

	const layout = "2006-01-02T15:04:05"
	ts := time.Now()

	yearPart1 := byte(ts.Year() >> 8)
	yearPart2 := byte(ts.Year() & 0xFF)
	month := byte(ts.Month())
	day := byte(ts.Day())
	hour := byte(ts.Hour())
	minute := byte(ts.Minute())
	second := byte(ts.Second())

	err := MarshalItems([]byte{0x09, 0x07, yearPart1, yearPart2, month, day, hour, minute, second}, &v)
	if err != nil {
		t.Errorf("Didn't expect error... %v", err)
	}
	if v.Data.Format(layout) != ts.Format(layout) {
		t.Errorf("Data wasn't what we expcted, was %v vs %v", v.Data.Format(layout), ts.Format(layout))
	}
}

func TestTwoInts(t *testing.T) {
	v := struct {
		Data1 int
		Data2 int
	}{}

	const expect = 16843009
	err := MarshalItems([]byte{0x06, 0x01, 0x01, 0x01, 0x01, 0x06, 0x01, 0x01, 0x01, 0x01}, &v)
	if err != nil {
		t.Errorf("Didn't expect error... %v", err)
	}
	if v.Data1 != expect {
		t.Errorf("#1 Expcted %v, was %v", expect, v.Data1)
	}
	if v.Data2 != expect {
		t.Errorf("#2 Expcted %v, was %v", expect, v.Data2)
	}
}

func TestTwoStrings(t *testing.T) {
	v := struct {
		Data1 string
		Data2 string
	}{}

	const expect = "Help me!"
	err := MarshalItems(append(
		append(
			[]byte{0x09, 0x08},
			[]byte(expect)...,
		),
		append(
			[]byte{0x09, 0x08},
			[]byte(expect)...,
		)...,
	), &v)
	if err != nil {
		t.Errorf("Didn't expect error... %v", err)
	}
	if v.Data1 != expect {
		t.Errorf("#1 Data was `%v`, expected `%v`", v.Data1, expect)
	}
	if v.Data2 != expect {
		t.Errorf("#2 Data was `%v`, expected `%v`", v.Data2, expect)
	}
}

func TestTooShortString(t *testing.T) {
	v := struct {
		Data1 string
		Data2 string
	}{}

	const expect = "Help m"
	err := MarshalItems(
		append(
			[]byte{0x09, 0x08},
			[]byte(expect)...,
		),
		&v,
	)
	if err != nil {
		t.Errorf("Didn't expect error... %v", err)
	}
	if v.Data1 != expect {
		t.Errorf("#1 Data was `%v`, expected `%v`", v.Data1, expect)
	}
	if v.Data2 != expect {
		t.Errorf("#2 Data was `%v`, expected `%v`", v.Data2, expect)
	}
}
