package hexii

import "testing"

func TestDump(t *testing.T) {
	var data []byte
	for i := 0; i < 28; i++ {
		data = append(data, byte(i))
	}
	for i := 0; i < 28; i++ {
		data = append(data, 0)
	}
	for i := 28; i < 130; i++ {
		data = append(data, byte(i))
	}
	for i := 0; i < 48; i++ {
		data = append(data, 0)
	}
	for i := 130; i < 266; i++ {
		data = append(data, byte(i))
	}
	for i := 0; i < 48; i++ {
		data = append(data, 0)
	}

	expected := "" +
		"0000:    01 02 03 04 05 06 07 08 09 0a 0b 0c 0d 0e 0f\n" +
		"  10: 10 11 12 13 14 15 16 17 18 19 1a 1b            \n" +
		"0030:                         1c 1d 1e 1f .  .! .\" .#\n" +
		"  40: .$ .% .& .' .( .) .* .+ ., .- .. ./ .0 .1 .2 .3\n" +
		"  50: .4 .5 .6 .7 .8 .9 .: .; .< .= .> .? .@ .A .B .C\n" +
		"  60: .D .E .F .G .H .I .J .K .L .M .N .O .P .Q .R .S\n" +
		"  70: .T .U .V .W .X .Y .Z .[ .\\ .] .^ ._ .` .a .b .c\n" +
		"  80: .d .e .f .g .h .i .j .k .l .m .n .o .p .q .r .s\n" +
		"  90: .t .u .v .w .x .y .z .{ .| .} .~ 7f 80 81      \n" +
		"00c0:                                           82 83\n" +
		"  d0: 84 85 86 87 88 89 8a 8b 8c 8d 8e 8f 90 91 92 93\n" +
		"  e0: 94 95 96 97 98 99 9a 9b 9c 9d 9e 9f a0 a1 a2 a3\n" +
		"  f0: a4 a5 a6 a7 a8 a9 aa ab ac ad ae af b0 b1 b2 b3\n" +
		" 100: b4 b5 b6 b7 b8 b9 ba bb bc bd be bf c0 c1 c2 c3\n" +
		"  10: c4 c5 c6 c7 c8 c9 ca cb cc cd ce cf d0 d1 d2 d3\n" +
		"  20: d4 d5 d6 d7 d8 d9 da db dc dd de df e0 e1 e2 e3\n" +
		"  30: e4 e5 e6 e7 e8 e9 ea eb ec ed ee ef f0 f1 f2 f3\n" +
		"  40: f4 f5 f6 f7 f8 f9 fa fb fc fd fe ##    01 02 03\n" +
		"  50: 04 05 06 07 08 09                              \n" +
		"0180:                   ]\n"
	res := Dump(data)
	if res != expected {
		t.Errorf("Result:\n%s\n\nExpected:\n%s", res, expected)
	}
}

func TestAppendLine(t *testing.T) {
	repr := string(appendLine(nil, []byte{0, 1, 31, 32, 126, 127, 254, 255}))
	expected := "    01 1f .  .~ 7f fe ##"
	if repr != expected {
		t.Errorf("%q = %q", repr, expected)
	}

}

func TestAppendRepr(t *testing.T) {
	cases := []struct {
		char byte
		repr string
	}{
		{0, "  "},
		{1, "01"},
		{31, "1f"},
		{32, ". "},
		{126, ".~"},
		{127, "7f"},
		{254, "fe"},
		{255, "##"},
	}

	for i, tc := range cases {
		repr := string(appendRepr(nil, tc.char))
		if repr != tc.repr {
			t.Errorf("%d: %q != %q for %d", i, repr, tc.repr, tc.char)
		}
	}
}
