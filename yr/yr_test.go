package yr

import (
	"testing"
)

func TestNewLines(t *testing.T) {
	type test struct {
		input string
		want  string
	}
	tests := []test{
		{input: "Kjevik;SN39040;18.03.2022 01:50;6",
			want: "Kjevik;SN39040;18.03.2022 01:50;42.8"},
		{input: "Kjevik;SN39040;18.03.2022 01:50",
			want: ""},
		{input: "Kjevik;SN39040;07.03.2023 18:20;0",
			want: "Kjevik;SN39040;07.03.2023 18:20;32.0"},
		{input: "Kjevik;SN39040;08.03.2023 02:20;-11",
			want: "Kjevik;SN39040;08.03.2023 02:20;12.2"},
		{input: "Data er gyldig per 18.03.2023 (CC BY 4.0), Meteorologisk institutt (MET);;;",
			want: "Data er basert på gyldig data (per 18.03.2023) (CC BY 4.0) fra Meteorologisk institutt (MET);endringen er gjort av Samet Demirezen"},
	}

	for _, tc := range tests {
		got, _ := NewLines(tc.input)
		if !(tc.want == got) {
			t.Errorf("expected %s, got: %s", tc.want, got)
		}
	}

}

func TestTotalLines(t *testing.T) {
	want := 16756

	got := TotalLines()
	if !(want == got) {
		t.Errorf("expected %v, got: %v", want, got)
	}

}

func TestAverageTempratur(t *testing.T) {
	want := 8.56

	got := AverageTempratureCelsius()
	if !(want == got) {
		t.Errorf("expected %v, got: %v", want, got)
	}

}
