package geo

import (
	"reflect"
	"testing"
)

func TestField_String(t *testing.T) {
	type fields struct {
		decoded LatLonDeg
		encoded latLonChar
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"toString-decoded-zero-1", fields{decoded: LatLonDeg{0, 0}}, "Decoded:{0.000000 0.000000}"},
		{"toString-decoded-zero-2", fields{decoded: LatLonDeg{}}, "Decoded:{0.000000 0.000000}"},
		{"toString-decoded-set", fields{decoded: LatLonDeg{0.0000001, 0.0000001}}, "Decoded:{0.000000 0.000000}"},
		{"toString-encoded-zero", fields{encoded: latLonChar{}}, "Decoded:{0.000000 0.000000}"},
		{"toString-encoded-set", fields{encoded: latLonChar{byte("A"[0]), byte("A"[0])}}, "Decoded:{0.000000 0.000000} Encoded:AA"},
		{"toString-encoded-decoded-set", fields{encoded: latLonChar{byte("A"[0]), byte("A"[0])}, decoded: LatLonDeg{-90, -180}}, "Decoded:{-90.000000 -180.000000} Encoded:AA"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &field{
				decoded: tt.fields.decoded,
				encoded: tt.fields.encoded,
			}
			if got := a.String(); got != tt.want {
				t.Errorf("field.String() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestField_Equals(t *testing.T) {
	type fields struct {
		decoded LatLonDeg
		encoded latLonChar
	}
	type args struct {
		b field
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{"equals-zero-1", fields{}, args{field{}}, true},
		{"equals-zero-2", fields{decoded: LatLonDeg{}}, args{field{}}, true},
		{"equals-zero-3", fields{decoded: LatLonDeg{}, encoded: latLonChar{}}, args{field{}}, true},

		{"equals-set-1", fields{decoded: LatLonDeg{10, 10}, encoded: latLonChar{}}, args{field{}}, false},
		{"equals-set-2", fields{decoded: LatLonDeg{10, 10}, encoded: latLonChar{47, 47}}, args{field{}}, false},
		{"equals-set-3", fields{decoded: LatLonDeg{10, 10}, encoded: latLonChar{}}, args{field{decoded: LatLonDeg{10, 10}, encoded: latLonChar{47, 47}}}, false},

		{"equals-set-4",
			fields{
				decoded: LatLonDeg{10, 10},
				encoded: latLonChar{47, 47}},
			args{field{
				decoded: LatLonDeg{10, 10},
				encoded: latLonChar{47, 47}}},
			true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &field{
				decoded: tt.fields.decoded,
				encoded: tt.fields.encoded,
			}
			if got := a.Equals(tt.args.b); got != tt.want {
				t.Errorf("field.Equals() = %v, want %v", got, tt.want)
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
		want field
	}{
		{"encode-zero-JJ", args{}, field{encoded: latLonChar{74, 74}}},

		{"encode-set-JJ-1", args{LatLonDeg{0.000001, 0.0000001}}, field{encoded: latLonChar{74, 74}}},
		{"encode-set-JJ-2", args{LatLonDeg{0.01, 0.01}}, field{encoded: latLonChar{74, 74}}},
		{"encode-set-JJ-3", args{LatLonDeg{9.99, 19.99}}, field{encoded: latLonChar{74, 74}}},
		{"encode-set-KK-4", args{LatLonDeg{10, 20}}, field{encoded: latLonChar{75, 75}, decoded: LatLonDeg{10, 20}}},

		{"encode-set-AA-1", args{LatLonDeg{-90, -180}}, field{encoded: latLonChar{65, 65}, decoded: LatLonDeg{-90, -180}}},
		{"encode-set-AA-2", args{LatLonDeg{-89.99, -179.99}}, field{encoded: latLonChar{65, 65}, decoded: LatLonDeg{-90, -180}}},
		{"encode-set-AA-3", args{LatLonDeg{-80.01, -160.01}}, field{encoded: latLonChar{65, 65}, decoded: LatLonDeg{-90, -180}}},
		{"encode-set-BB-3", args{LatLonDeg{-80, -160}}, field{encoded: latLonChar{66, 66}, decoded: LatLonDeg{-80, -160}}},

		{"encode-set-RR-1", args{LatLonDeg{80, 160}}, field{encoded: latLonChar{82, 82}, decoded: LatLonDeg{80, 160}}},
		{"encode-set-RR-2", args{LatLonDeg{80.1, 160.1}}, field{encoded: latLonChar{82, 82}, decoded: LatLonDeg{80, 160}}},
		{"encode-set-RR-3", args{LatLonDeg{89.999, 170.99}}, field{encoded: latLonChar{82, 82}, decoded: LatLonDeg{80, 160}}},
		{"encode-set-PP-4", args{LatLonDeg{70.999, 159.99}}, field{encoded: latLonChar{81, 81}, decoded: LatLonDeg{70, 140}}},

		{"encode-set-AR-1", args{LatLonDeg{-90, 179.99}}, field{encoded: latLonChar{65, 82}, decoded: LatLonDeg{-90, 160}}},
		{"encode-set-AR-2", args{LatLonDeg{-89.999, 170.99}}, field{encoded: latLonChar{65, 82}, decoded: LatLonDeg{-90, 160}}},
		{"encode-set-AR-3", args{LatLonDeg{-80.0001, 160.001}}, field{encoded: latLonChar{65, 82}, decoded: LatLonDeg{-90, 160}}},

		{"encode-set-RA-1", args{LatLonDeg{89.99999, -180}}, field{encoded: latLonChar{82, 65}, decoded: LatLonDeg{80, -180}}},
		{"encode-set-RA-2", args{LatLonDeg{89.999, -170.99}}, field{encoded: latLonChar{82, 65}, decoded: LatLonDeg{80, -180}}},
		{"encode-set-RA-3", args{LatLonDeg{80.0001, -160.001}}, field{encoded: latLonChar{82, 65}, decoded: LatLonDeg{80, -180}}},

		{"encode-set-S59ABC-JN76TO", args{LatLonDeg{46.3, 15.3}}, field{encoded: latLonChar{78, 74}, decoded: LatLonDeg{40, 0}}},
		{"encode-set-K1TTT-FN32LL", args{LatLonDeg{42.4662, -73.0232}}, field{encoded: latLonChar{78, 70}, decoded: LatLonDeg{40, -80}}},
		{"encode-set-PS2T-GG58WG", args{LatLonDeg{-21.7487, -48.1268}}, field{encoded: latLonChar{71, 71}, decoded: LatLonDeg{-30, -60}}},
		{"encode-set-ZM4T-RF80LQ", args{LatLonDeg{-39.3125, 176.9583333}}, field{encoded: latLonChar{70, 82}, decoded: LatLonDeg{-40, 160}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fieldEncode(tt.args.lld); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fieldEncode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFieldDecode(t *testing.T) {
	type args struct {
		llc latLonChar
	}
	tests := []struct {
		name string
		args args
		want field
	}{
		{"decode-zero", args{}, field{encoded: latLonChar{0, 0}, decoded: LatLonDeg{0, 0}}},
		{"decode-set-JJ-1", args{latLonChar{74, 74}}, field{encoded: latLonChar{74, 74}, decoded: LatLonDeg{0, 0}}},
		{"decode-set-AA-1", args{latLonChar{65, 65}}, field{encoded: latLonChar{65, 65}, decoded: LatLonDeg{-90, -180}}},
		{"decode-set-BB-1", args{latLonChar{66, 66}}, field{encoded: latLonChar{66, 66}, decoded: LatLonDeg{-80, -160}}},
		{"decode-set-PP-1", args{latLonChar{81, 81}}, field{encoded: latLonChar{81, 81}, decoded: LatLonDeg{70, 140}}},
		{"decode-set-RR-1", args{latLonChar{82, 82}}, field{encoded: latLonChar{82, 82}, decoded: LatLonDeg{80, 160}}},

		{"decode-set-AR-1", args{latLonChar{65, 82}}, field{encoded: latLonChar{65, 82}, decoded: LatLonDeg{-90, 160}}},
		{"decode-set-RA-1", args{latLonChar{82, 65}}, field{encoded: latLonChar{82, 65}, decoded: LatLonDeg{80, -180}}},

		{"encode-set-S59ABC-JN76TO", args{latLonChar{78, 74}}, field{encoded: latLonChar{78, 74}, decoded: LatLonDeg{40, 0}}},
		{"encode-set-K1TTT-FN32LL", args{latLonChar{78, 70}}, field{encoded: latLonChar{78, 70}, decoded: LatLonDeg{40, -80}}},
		{"encode-set-PS2T-GG58WG", args{latLonChar{71, 71}}, field{encoded: latLonChar{71, 71}, decoded: LatLonDeg{-30, -60}}},
		{"encode-set-ZM4T-RF80LQ", args{latLonChar{70, 82}}, field{encoded: latLonChar{70, 82}, decoded: LatLonDeg{-40, 160}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fieldDecode(tt.args.llc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fieldDecode() = %v, want %v", got, tt.want)
			}
		})
	}
}
