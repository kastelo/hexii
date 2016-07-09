hexii
=====

Package hexii formats data in the HexII format.

HexII format: https://github.com/gunmetalbackupgooglecode/corkami/tree/master/src/HexII

HexII is a more compact hex dump format, based on the principle that in most
cases ASCII output is preferred for ASCII characters and hex output for
other data.

A common hex dump shows both hex and ASCII:

```
00000000  cf fa ed fe 07 00 00 01  03 00 00 80 02 00 00 00  |................|
00000010  12 00 00 00 08 07 00 00  85 00 20 00 00 00 00 00  |.......... .....|
00000020  19 00 00 00 48 00 00 00  5f 5f 50 41 47 45 5a 45  |....H...__PAGEZE|
00000030  52 4f 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |RO..............|
00000040  00 00 00 00 01 00 00 00  00 00 00 00 00 00 00 00  |................|
00000050  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
00000060  00 00 00 00 00 00 00 00  19 00 00 00 28 02 00 00  |............(...|
00000070  5f 5f 54 45 58 54 00 00  00 00 00 00 00 00 00 00  |__TEXT..........|
00000080  00 00 00 00 01 00 00 00  00 50 00 00 00 00 00 00  |.........P......|
00000090  00 00 00 00 00 00 00 00  00 50 00 00 00 00 00 00  |.........P......|
000000a0  07 00 00 00 05 00 00 00  06 00 00 00 00 00 00 00  |................|
000000b0  5f 5f 74 65 78 74 00 00  00 00 00 00 00 00 00 00  |__text..........|
000000c0  5f 5f 54 45 58 54 00 00  00 00 00 00 00 00 00 00  |__TEXT..........|
000000d0  94 0e 00 00 01 00 00 00  99 35 00 00 00 00 00 00  |.........5......|
000000e0  94 0e 00 00 02 00 00 00  00 00 00 00 00 00 00 00  |................|
000000f0  00 04 00 80 00 00 00 00  00 00 00 00 00 00 00 00  |................|
00000100  5f 5f 73 74 75 62 73 00  00 00 00 00 00 00 00 00  |__stubs.........|
00000110  5f 5f 54 45 58 54 00 00  00 00 00 00 00 00 00 00  |__TEXT..........|
00000120  2e 44 00 00 01 00 00 00  c8 01 00 00 00 00 00 00  |.D..............|
00000130  2e 44 00 00 01 00 00 00  00 00 00 00 00 00 00 00  |.D..............|
00000140  08 04 00 80 00 00 00 00  06 00 00 00 00 00 00 00  |................|
00000150  5f 5f 73 74 75 62 5f 68  65 6c 70 65 72 00 00 00  |__stub_helper...|
00000160  5f 5f 54 45 58 54 00 00  00 00 00 00 00 00 00 00  |__TEXT..........|
00000170  f8 45 00 00 01 00 00 00  08 03 00 00 00 00 00 00  |.E..............|
```

HexII instead compresses zeroes, highlights the nonzero data more clearly,
and shows only ASCII when appropriate:

```
       0  1  2  3  4  5  6  7  8  9  a  b  c  d  e  f

0000: cf fa ed fe 07       01 03       80 02
  10: 12          08 07       85    .
  20: 19          .H          ._ ._ .P .A .G .E .Z .E
  30: .R .O
  40:             01
0060:                         19          .( 02
  70: ._ ._ .T .E .X .T
  80:             01             .P
  90:                            .P
  a0: 07          05          06
  b0: ._ ._ .t .e .x .t
  c0: ._ ._ .T .E .X .T
  d0: 94 0e       01          99 .5
  e0: 94 0e       02
  f0:    04    80
 100: ._ ._ .s .t .u .b .s
  10: ._ ._ .T .E .X .T
  20: .. .D       01          c8 01
  30: .. .D       01
  40: 08 04    80             06
  50: ._ ._ .s .t .u .b ._ .h .e .l .p .e .r
  60: ._ ._ .T .E .X .T
  70: f8 .E       01          08 03
  80: ]
```

## Rules

### Hex

 - ASCII chars are displayed as ".char"
 - 00 is shown as "  "
 - FF is shown as "##"
 - other bytes are displayed in hex

### Output

 - a hex ruler is shown at the top of the display
 - lines consisting solely of zeroes are skipped
 - no leading zeroes in offsets
 - offsets have leading digits repeated from the previous line removed
 - offsets after a skip are written out fully
 - last offset + 1 is marked with "]" to indicate the EOF position

## License

MIT