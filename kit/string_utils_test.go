package kit

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Strings_Distinct(t *testing.T) {
	s := Strings{"aaa", "bbb", "aaa"}
	assert.Equal(t, s.Distinct(), Strings{"aaa", "bbb"})
	s = Strings{"aaa", "bbb"}
	assert.Equal(t, s.Distinct(), Strings{"aaa", "bbb"})
}

func Test_Strings_Contains(t *testing.T) {
	s := Strings{"aaa", "bbb", "aaa"}
	assert.True(t, s.Contains("aaa"))
	assert.True(t, s.Contains("bbb"))
	assert.False(t, s.Contains("ccc"))
	assert.False(t, s.Contains(""))
	assert.False(t, s.Contains("aa"))
}

func Test_Strings_Intersect(t *testing.T) {
	assert.ElementsMatch(t, Strings{"aaa", "bbb"}, Strings{"aaa", "bbb", "aaa"}.Intersect(Strings{"aaa", "bbb"}))
	assert.ElementsMatch(t, Strings{"aaa", "bbb"}, Strings{"bbb", "aaa"}.Intersect(Strings{"aaa", "bbb"}))
	assert.ElementsMatch(t, Strings{}, Strings{}.Intersect(Strings{}))
	assert.ElementsMatch(t, Strings{}, Strings(nil).Intersect(nil))
	assert.ElementsMatch(t, Strings{}, Strings{"aa"}.Intersect(Strings{}))
	assert.ElementsMatch(t, Strings{}, Strings{}.Intersect(Strings{"bb"}))
	assert.ElementsMatch(t, Strings{}, Strings{"bb"}.Intersect(Strings{"aa"}))
	assert.ElementsMatch(t, Strings{"bb"}, Strings{"bb", "bb"}.Intersect(Strings{"bb", "bb", "bb", "bb"}))
}

func Test_StrToInt64(t *testing.T) {
	for _, s := range []struct {
		In  string
		Out int64
		Err bool
	}{
		{
			In:  "",
			Out: 0,
			Err: true,
		},
		{
			In:  "qqq",
			Out: 0,
			Err: true,
		},
		{
			In:  "0.23123",
			Out: 0,
			Err: true,
		},
		{
			In:  "-1",
			Out: -1,
			Err: false,
		},
		{
			In:  "1576663112362381",
			Out: 1576663112362381,
			Err: false,
		},
	} {
		out, err := StrToInt64(s.In)
		if s.Err {
			assert.Error(t, err)
		} else {
			assert.Nil(t, err)
			assert.Equal(t, s.Out, out)
		}
	}
}

func Test_RemoveNonAlfaDigital(t *testing.T) {
	tests := []struct {
		name string
		in   string
		out  string
	}{
		{
			name: "empty string",
			in:   "",
			out:  "",
		},
		{
			name: "complex case",
			in: "  A++B%%C///--	 %:*%*abc \t@#$%123   &&& АБВ *)__^^ абв",
			out: "ABCabc123АБВабв",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.out, RemoveNonAlfaDigital(tt.in))
		})
	}
}

func Test_Digits(t *testing.T) {
	for _, s := range []struct {
		In  string
		Out bool
	}{
		{
			In:  "",
			Out: false,
		},
		{
			In:  "qweqe",
			Out: false,
		},
		{
			In:  "123q231",
			Out: false,
		},
		{
			In:  "0",
			Out: true,
		},
		{
			In:  "0.5",
			Out: false,
		},
		{
			In:  "-5",
			Out: false,
		},
		{
			In:  "0214124214",
			Out: true,
		},
	} {
		assert.Equal(t, s.Out, Digits(s.In))
	}
}

func Test_YNtoBool(t *testing.T) {
	fl := false
	tr := true
	for _, s := range []struct {
		In  string
		Out *bool
	}{
		{
			In:  "",
			Out: nil,
		},
		{
			In:  "123",
			Out: nil,
		},
		{
			In:  "true",
			Out: nil,
		},
		{
			In:  "Yes",
			Out: &tr,
		},
		{
			In:  " дА ",
			Out: &tr,
		},
		{
			In:  "NO",
			Out: &fl,
		},
		{
			In:  " нет   ",
			Out: &fl,
		},
	} {
		assert.Equal(t, s.Out, YesNoToBool(s.In))
	}
}
