package internal

import (
	"reflect"
	"testing"
)

func (a *Field) equals(b Field) bool {
	return a.Encoded.equal(b.Encoded) && a.Decoded.equal(b.Decoded)
}

func TestField_String(t *testing.T) {
	type fields struct {
		decoded LatLonDeg
		encoded LatLonChar
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"toString-Decoded-zero-1", fields{decoded: LatLonDeg{0, 0}}, "Decoded:{0.000000 0.000000}"},
		{"toString-Decoded-zero-2", fields{decoded: LatLonDeg{}}, "Decoded:{0.000000 0.000000}"},
		{"toString-Decoded-set", fields{decoded: LatLonDeg{0.0000001, 0.0000001}}, "Decoded:{0.000000 0.000000}"},
		{"toString-Encoded-zero", fields{encoded: LatLonChar{}}, "Decoded:{0.000000 0.000000}"},
		{"toString-Encoded-set", fields{encoded: LatLonChar{byte("A"[0]), byte("A"[0])}}, "Decoded:{0.000000 0.000000} Encoded:AA"},
		{"toString-Encoded-Decoded-set", fields{encoded: LatLonChar{byte("A"[0]), byte("A"[0])}, decoded: LatLonDeg{-90, -180}}, "Decoded:{-90.000000 -180.000000} Encoded:AA"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Field{
				Decoded: tt.fields.decoded,
				Encoded: tt.fields.encoded,
			}
			if got := a.String(); got != tt.want {
				t.Errorf("Field.String() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestField_Equals(t *testing.T) {
	type fields struct {
		decoded LatLonDeg
		encoded LatLonChar
	}
	type args struct {
		b Field
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{"equals-zero-1", fields{}, args{Field{}}, true},
		{"equals-zero-2", fields{decoded: LatLonDeg{}}, args{Field{}}, true},
		{"equals-zero-3", fields{decoded: LatLonDeg{}, encoded: LatLonChar{}}, args{Field{}}, true},

		{"equals-set-1", fields{decoded: LatLonDeg{10, 10}, encoded: LatLonChar{}}, args{Field{}}, false},
		{"equals-set-2", fields{decoded: LatLonDeg{10, 10}, encoded: LatLonChar{47, 47}}, args{Field{}}, false},
		{"equals-set-3", fields{decoded: LatLonDeg{10, 10}, encoded: LatLonChar{}}, args{Field{Decoded: LatLonDeg{10, 10}, Encoded: LatLonChar{47, 47}}}, false},

		{"equals-set-4",
			fields{
				decoded: LatLonDeg{10, 10},
				encoded: LatLonChar{47, 47}},
			args{Field{
				Decoded: LatLonDeg{10, 10},
				Encoded: LatLonChar{47, 47}}},
			true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Field{
				Decoded: tt.fields.decoded,
				Encoded: tt.fields.encoded,
			}
			if got := a.equals(tt.args.b); got != tt.want {
				t.Errorf("Field.equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFieldEncode(t *testing.T) {
	type args struct {
		lld LatLonDeg
	}
	tests := []struct {
		name string
		args args
		want Field
	}{
		{"encode-zero-JJ", args{}, Field{Encoded: LatLonChar{74, 74}}},

		{"encode-set-JJ-1", args{LatLonDeg{0.000001, 0.0000001}}, Field{Encoded: LatLonChar{74, 74}}},
		{"encode-set-JJ-2", args{LatLonDeg{0.01, 0.01}}, Field{Encoded: LatLonChar{74, 74}}},
		{"encode-set-JJ-3", args{LatLonDeg{9.99, 19.99}}, Field{Encoded: LatLonChar{74, 74}}},
		{"encode-set-KK-4", args{LatLonDeg{10, 20}}, Field{Encoded: LatLonChar{75, 75}, Decoded: LatLonDeg{10, 20}}},

		{"encode-set-AA-1", args{LatLonDeg{-90, -180}}, Field{Encoded: LatLonChar{65, 65}, Decoded: LatLonDeg{-90, -180}}},
		{"encode-set-AA-2", args{LatLonDeg{-89.99, -179.99}}, Field{Encoded: LatLonChar{65, 65}, Decoded: LatLonDeg{-90, -180}}},
		{"encode-set-AA-3", args{LatLonDeg{-80.01, -160.01}}, Field{Encoded: LatLonChar{65, 65}, Decoded: LatLonDeg{-90, -180}}},
		{"encode-set-BB-3", args{LatLonDeg{-80, -160}}, Field{Encoded: LatLonChar{66, 66}, Decoded: LatLonDeg{-80, -160}}},

		{"encode-set-RR-1", args{LatLonDeg{80, 160}}, Field{Encoded: LatLonChar{82, 82}, Decoded: LatLonDeg{80, 160}}},
		{"encode-set-RR-2", args{LatLonDeg{80.1, 160.1}}, Field{Encoded: LatLonChar{82, 82}, Decoded: LatLonDeg{80, 160}}},
		{"encode-set-RR-3", args{LatLonDeg{89.999, 170.99}}, Field{Encoded: LatLonChar{82, 82}, Decoded: LatLonDeg{80, 160}}},
		{"encode-set-PP-4", args{LatLonDeg{70.999, 159.99}}, Field{Encoded: LatLonChar{81, 81}, Decoded: LatLonDeg{70, 140}}},

		{"encode-set-AR-1", args{LatLonDeg{-90, 179.99}}, Field{Encoded: LatLonChar{65, 82}, Decoded: LatLonDeg{-90, 160}}},
		{"encode-set-AR-2", args{LatLonDeg{-89.999, 170.99}}, Field{Encoded: LatLonChar{65, 82}, Decoded: LatLonDeg{-90, 160}}},
		{"encode-set-AR-3", args{LatLonDeg{-80.0001, 160.001}}, Field{Encoded: LatLonChar{65, 82}, Decoded: LatLonDeg{-90, 160}}},

		{"encode-set-RA-1", args{LatLonDeg{89.99999, -180}}, Field{Encoded: LatLonChar{82, 65}, Decoded: LatLonDeg{80, -180}}},
		{"encode-set-RA-2", args{LatLonDeg{89.999, -170.99}}, Field{Encoded: LatLonChar{82, 65}, Decoded: LatLonDeg{80, -180}}},
		{"encode-set-RA-3", args{LatLonDeg{80.0001, -160.001}}, Field{Encoded: LatLonChar{82, 65}, Decoded: LatLonDeg{80, -180}}},

		{"encode-set-S59ABC-JN76TO", args{LatLonDeg{46.3, 15.3}}, Field{Encoded: LatLonChar{78, 74}, Decoded: LatLonDeg{40, 0}}},
		{"encode-set-K1TTT-FN32LL", args{LatLonDeg{42.4662, -73.0232}}, Field{Encoded: LatLonChar{78, 70}, Decoded: LatLonDeg{40, -80}}},
		{"encode-set-PS2T-GG58WG", args{LatLonDeg{-21.7487, -48.1268}}, Field{Encoded: LatLonChar{71, 71}, Decoded: LatLonDeg{-30, -60}}},
		{"encode-set-ZM4T-RF80LQ", args{LatLonDeg{-39.3125, 176.9583333}}, Field{Encoded: LatLonChar{70, 82}, Decoded: LatLonDeg{-40, 160}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FieldEncode(tt.args.lld); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FieldEncode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFieldDecode(t *testing.T) {
	type args struct {
		llc LatLonChar
	}
	tests := []struct {
		name string
		args args
		want Field
	}{
		{"decode-zero", args{}, Field{Encoded: LatLonChar{0, 0}, Decoded: LatLonDeg{0, 0}}},
		{"decode-set-JJ-1", args{LatLonChar{74, 74}}, Field{Encoded: LatLonChar{74, 74}, Decoded: LatLonDeg{0, 0}}},
		{"decode-set-AA-1", args{LatLonChar{65, 65}}, Field{Encoded: LatLonChar{65, 65}, Decoded: LatLonDeg{-90, -180}}},
		{"decode-set-BB-1", args{LatLonChar{66, 66}}, Field{Encoded: LatLonChar{66, 66}, Decoded: LatLonDeg{-80, -160}}},
		{"decode-set-PP-1", args{LatLonChar{81, 81}}, Field{Encoded: LatLonChar{81, 81}, Decoded: LatLonDeg{70, 140}}},
		{"decode-set-RR-1", args{LatLonChar{82, 82}}, Field{Encoded: LatLonChar{82, 82}, Decoded: LatLonDeg{80, 160}}},

		{"decode-set-AR-1", args{LatLonChar{65, 82}}, Field{Encoded: LatLonChar{65, 82}, Decoded: LatLonDeg{-90, 160}}},
		{"decode-set-RA-1", args{LatLonChar{82, 65}}, Field{Encoded: LatLonChar{82, 65}, Decoded: LatLonDeg{80, -180}}},

		{"encode-set-S59ABC-JN76TO", args{LatLonChar{78, 74}}, Field{Encoded: LatLonChar{78, 74}, Decoded: LatLonDeg{40, 0}}},
		{"encode-set-K1TTT-FN32LL", args{LatLonChar{78, 70}}, Field{Encoded: LatLonChar{78, 70}, Decoded: LatLonDeg{40, -80}}},
		{"encode-set-PS2T-GG58WG", args{LatLonChar{71, 71}}, Field{Encoded: LatLonChar{71, 71}, Decoded: LatLonDeg{-30, -60}}},
		{"encode-set-ZM4T-RF80LQ", args{LatLonChar{70, 82}}, Field{Encoded: LatLonChar{70, 82}, Decoded: LatLonDeg{-40, 160}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FieldDecode(tt.args.llc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FieldDecode() = %v, want %v", got, tt.want)
			}
		})
	}
}
