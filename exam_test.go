package main

import (
	"testing"
)

func TestExam(t *testing.T) {

	t.Run("StringReduction 'abcabc' output: 4", func(t *testing.T) {
		v := StringReduction("abcabc")

		if v != 4 {
			t.Error("StringReduction output not 4 but have", v)
		}
	})

	t.Run("StringReduction 'abca' output: 1", func(t *testing.T) {
		v := StringReduction("abca")

		if v != 1 {
			t.Error("StringReduction output not 1 but have", v)
		}
	})

	t.Run("StringReduction 'cccc' output: 4", func(t *testing.T) {
		v := StringReduction("cccc")

		if v != 4 {
			t.Error("StringReduction output not 4 but have", v)
		}
	})
}
