package internal

import (
	"reflect"
	"testing"
)

func (a *Square) equals(b Square) bool {
	return a.Encoded.equal(b.Encoded) && a.Decoded.equal(b.Decoded)
}

//func TestSquare_Decode_01(t *testing.T) {
//	fc := LatLonChar{}
//	fc.setLonChar("J")
//	fc.setLatChar("N")
//	sc := LatLonChar{}
//	sc.setLonChar("7")
//	sc.setLatChar("6")
//	f := FieldDecode(fc)
//	s := SquareDecode(f, sc)
//	fmt.Println(f.String())
//	fmt.Println("s")
//	fmt.Println(s.String())
//	//
//	//
//	lld := LatLonDeg{}
//	lld.Lon = f.Decoded.Lon + s.Decoded.Lon
//	lld.Lat = f.Decoded.Lat + s.Decoded.Lat
//
//	_, sa := SquareEncode(lld)
//	fmt.Println("sa")
//	fmt.Println(sa.String())
//	if s.equals(sa) {
//		t.Fatal()
//	}
//}

func TestSquare_String(t *testing.T) {
	type fields struct {
		decoded LatLonDeg
		encoded LatLonChar
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "toString-zero-1",
			fields: fields{},
			want:   "Decoded:{0.000000 0.000000}",
		},
		{
			name: "toString-Decoded-zero-2",
			fields: fields{
				decoded: LatLonDeg{
					Lat: 0,
					Lon: 0,
				},
			},
			want: "Decoded:{0.000000 0.000000}",
		},
		{
			name: "toString-zero-3",
			fields: fields{
				decoded: LatLonDeg{
					Lat: 0,
					Lon: 0,
				},
				encoded: LatLonChar{
					LatChar: 0,
					LonChar: 0,
				},
			},
			want: "Decoded:{0.000000 0.000000}",
		},
		{
			name: "toString-Decoded-1",
			fields: fields{
				decoded: LatLonDeg{
					Lat: 1,
					Lon: 2,
				},
			},
			want: "Decoded:{1.000000 2.000000}",
		},
		{
			name: "toString-Decoded-2",
			fields: fields{
				decoded: LatLonDeg{
					Lat: 9,
					Lon: 18,
				},
			},
			want: "Decoded:{9.000000 18.000000}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Square{
				Decoded: tt.fields.decoded,
				Encoded: tt.fields.encoded,
			}
			if got := a.String(); got != tt.want {
				t.Errorf("Square.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSquare_Equals(t *testing.T) {
	type fields struct {
		decoded LatLonDeg
		encoded LatLonChar
	}
	type args struct {
		b Square
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "equals-1",
			fields: fields{
				decoded: LatLonDeg{
					Lat: 0,
					Lon: 0,
				},
				encoded: LatLonChar{
					LatChar: 0,
					LonChar: 0,
				},
			},
			args: args{
				b: Square{
					Decoded: LatLonDeg{
						Lat: 0,
						Lon: 0,
					},
					Encoded: LatLonChar{
						LatChar: 0,
						LonChar: 0,
					},
				},
			},
			want: true,
		},
		{
			name: "equals-2",
			fields: fields{
				decoded: LatLonDeg{
					Lat: 1,
					Lon: 2,
				},
				encoded: LatLonChar{
					LatChar: 48,
					LonChar: 49,
				},
			},
			args: args{
				b: Square{
					Decoded: LatLonDeg{
						Lat: 0,
						Lon: 0,
					},
					Encoded: LatLonChar{
						LatChar: 0,
						LonChar: 0,
					},
				},
			},
			want: false,
		},
		{
			name: "equals-3",
			fields: fields{
				decoded: LatLonDeg{
					Lat: 1,
					Lon: 2,
				},
				encoded: LatLonChar{
					LatChar: 48,
					LonChar: 49,
				},
			},
			args: args{
				b: Square{
					Decoded: LatLonDeg{
						Lat: 1,
						Lon: 2,
					},
					Encoded: LatLonChar{
						LatChar: 0,
						LonChar: 0,
					},
				},
			},
			want: false,
		},

		{
			name: "equals-4",
			fields: fields{
				decoded: LatLonDeg{
					Lat: 1,
					Lon: 2,
				},
				encoded: LatLonChar{
					LatChar: 48,
					LonChar: 49,
				},
			},
			args: args{
				b: Square{
					Decoded: LatLonDeg{
						Lat: 1,
						Lon: 2,
					},
					Encoded: LatLonChar{
						LatChar: 48,
						LonChar: 49,
					},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Square{
				Decoded: tt.fields.decoded,
				Encoded: tt.fields.encoded,
			}
			if got := a.equals(tt.args.b); got != tt.want {
				t.Errorf("Square.equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSquareEncode(t *testing.T) {
	type args struct {
		lld LatLonDeg
	}
	tests := []struct {
		name  string
		args  args
		want  Field
		want1 Square
	}{
		{
			name: "encode-JJ00-1",
			args: args{
				lld: LatLonDeg{
					Lat: 0,
					Lon: 0,
				},
			},
			want: Field{
				Decoded: LatLonDeg{
					Lat: 0,
					Lon: 0,
				},
				Encoded: LatLonChar{
					LatChar: 74,
					LonChar: 74,
				},
			},
			want1: Square{
				Decoded: LatLonDeg{
					Lat: 0,
					Lon: 0,
				},
				Encoded: LatLonChar{
					LatChar: 48,
					LonChar: 48,
				},
			},
		},

		{
			name: "encode-JJ00-2",
			args: args{
				lld: LatLonDeg{
					Lat: 0.0001,
					Lon: 0.0001,
				},
			},
			want: Field{
				Decoded: LatLonDeg{
					Lat: 0,
					Lon: 0,
				},
				Encoded: LatLonChar{
					LatChar: 74,
					LonChar: 74,
				},
			},
			want1: Square{
				Decoded: LatLonDeg{
					Lat: 0,
					Lon: 0,
				},
				Encoded: LatLonChar{
					LatChar: 48,
					LonChar: 48,
				},
			},
		},

		{
			name: "encode-JJ00-3",
			args: args{
				lld: LatLonDeg{
					Lat: 0.01,
					Lon: 0.01,
				},
			},
			want: Field{
				Decoded: LatLonDeg{
					Lat: 0,
					Lon: 0,
				},
				Encoded: LatLonChar{
					LatChar: 74,
					LonChar: 74,
				},
			},
			want1: Square{
				Decoded: LatLonDeg{
					Lat: 0,
					Lon: 0,
				},
				Encoded: LatLonChar{
					LatChar: 48,
					LonChar: 48,
				},
			},
		},

		{
			name: "encode-JJ99-4",
			args: args{
				lld: LatLonDeg{
					Lat: 9.99,
					Lon: 19.99,
				},
			},
			want: Field{
				Decoded: LatLonDeg{
					Lat: 0,
					Lon: 0,
				},
				Encoded: LatLonChar{
					LatChar: 74,
					LonChar: 74,
				},
			},
			want1: Square{
				Decoded: LatLonDeg{
					Lat: 9,
					Lon: 18,
				},
				Encoded: LatLonChar{
					LatChar: 57,
					LonChar: 57,
				},
			},
		},

		{
			name: "encode-KK00-5",
			args: args{
				lld: LatLonDeg{
					Lat: 10,
					Lon: 20,
				},
			},
			want: Field{
				Decoded: LatLonDeg{
					Lat: 10,
					Lon: 20,
				},
				Encoded: LatLonChar{
					LatChar: 75,
					LonChar: 75,
				},
			},
			want1: Square{
				Decoded: LatLonDeg{
					Lat: 0,
					Lon: 0,
				},
				Encoded: LatLonChar{
					LatChar: 48,
					LonChar: 48,
				},
			},
		},

		{
			name: "encode-AA00-5",
			args: args{
				lld: LatLonDeg{
					Lat: -90,
					Lon: -180,
				},
			},
			want: Field{
				Decoded: LatLonDeg{
					Lat: -90,
					Lon: -180,
				},
				Encoded: LatLonChar{
					LatChar: 65,
					LonChar: 65,
				},
			},
			want1: Square{
				Decoded: LatLonDeg{
					Lat: 0,
					Lon: 0,
				},
				Encoded: LatLonChar{
					LatChar: 48,
					LonChar: 48,
				},
			},
		},

		{
			name: "encode-AA00-6",
			args: args{
				lld: LatLonDeg{
					Lat: -89.99,
					Lon: -179.99,
				},
			},
			want: Field{
				Decoded: LatLonDeg{
					Lat: -90,
					Lon: -180,
				},
				Encoded: LatLonChar{
					LatChar: 65,
					LonChar: 65,
				},
			},
			want1: Square{
				Decoded: LatLonDeg{
					Lat: 0,
					Lon: 0,
				},
				Encoded: LatLonChar{
					LatChar: 48,
					LonChar: 48,
				},
			},
		},

		{
			name: "encode-AA99-7",
			args: args{
				lld: LatLonDeg{
					Lat: -80.01,
					Lon: -160.01,
				},
			},
			want: Field{
				Decoded: LatLonDeg{
					Lat: -90,
					Lon: -180,
				},
				Encoded: LatLonChar{
					LatChar: 65,
					LonChar: 65,
				},
			},
			want1: Square{
				Decoded: LatLonDeg{
					Lat: 9,
					Lon: 18,
				},
				Encoded: LatLonChar{
					LatChar: 57,
					LonChar: 57,
				},
			},
		},

		{
			name: "encode-BB00-7",
			args: args{
				lld: LatLonDeg{
					Lat: -80,
					Lon: -160,
				},
			},
			want: Field{
				Decoded: LatLonDeg{
					Lat: -80,
					Lon: -160,
				},
				Encoded: LatLonChar{
					LatChar: 66,
					LonChar: 66,
				},
			},
			want1: Square{
				Decoded: LatLonDeg{
					Lat: 0,
					Lon: 0,
				},
				Encoded: LatLonChar{
					LatChar: 48,
					LonChar: 48,
				},
			},
		},

		{
			name: "encode-S59ABC-JN76TO",
			args: args{
				lld: LatLonDeg{
					Lat: 46.3,
					Lon: 15.3,
				},
			},
			want: Field{
				Decoded: LatLonDeg{
					Lat: 40,
					Lon: 0,
				},
				Encoded: LatLonChar{
					LatChar: 78,
					LonChar: 74,
				},
			},
			want1: Square{
				Decoded: LatLonDeg{
					Lat: 6,
					Lon: 14,
				},
				Encoded: LatLonChar{
					LatChar: 54,
					LonChar: 55,
				},
			},
		},

		{
			name: "encode-K1TTT-FN32LL",
			args: args{
				lld: LatLonDeg{
					Lat: 42.4662,
					Lon: -73.0232,
				},
			},
			want: Field{
				Decoded: LatLonDeg{
					Lat: 40,
					Lon: -80,
				},
				Encoded: LatLonChar{
					LatChar: 78,
					LonChar: 70,
				},
			},
			want1: Square{
				Decoded: LatLonDeg{
					Lat: 2,
					Lon: 6,
				},
				Encoded: LatLonChar{
					LatChar: 50,
					LonChar: 51,
				},
			},
		},

		{
			name: "encode-PS2T-GG58WG",
			args: args{
				lld: LatLonDeg{
					Lat: -21.7487,
					Lon: -48.1268,
				},
			},
			want: Field{
				Decoded: LatLonDeg{
					Lat: -30,
					Lon: -60,
				},
				Encoded: LatLonChar{
					LatChar: 71,
					LonChar: 71,
				},
			},
			want1: Square{
				Decoded: LatLonDeg{
					Lat: 8,
					Lon: 10,
				},
				Encoded: LatLonChar{
					LatChar: 56,
					LonChar: 53,
				},
			},
		},

		{
			name: "encode-ZM4T-RF80LQ",
			args: args{
				lld: LatLonDeg{
					Lat: -39.3125,
					Lon: 176.9583333,
				},
			},
			want: Field{
				Decoded: LatLonDeg{
					Lat: -40,
					Lon: 160,
				},
				Encoded: LatLonChar{
					LatChar: 70,
					LonChar: 82,
				},
			},
			want1: Square{
				Decoded: LatLonDeg{
					Lat: 0,
					Lon: 16,
				},
				Encoded: LatLonChar{
					LatChar: 48,
					LonChar: 56,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := SquareEncode(tt.args.lld)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FielsEncode() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("SquareEncode() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSquareDecode(t *testing.T) {
	type args struct {
		f         Field
		squareLLC LatLonChar
	}
	tests := []struct {
		name string
		args args
		want Square
	}{
		{
			name: "decode-zero",
			args: args{
				f: Field{
					Decoded: LatLonDeg{
						Lat: 0,
						Lon: 0,
					},
					Encoded: LatLonChar{
						LatChar: 0,
						LonChar: 0,
					},
				},
				squareLLC: LatLonChar{
					LatChar: 0,
					LonChar: 0,
				},
			},
			want: Square{
				Decoded: LatLonDeg{
					Lat: 0,
					Lon: 0,
				},
				Encoded: LatLonChar{
					LatChar: 0,
					LonChar: 0,
				},
			},
		},

		{
			name: "decode-JJ00",
			args: args{
				f: Field{
					Encoded: LatLonChar{
						LatChar: 74,
						LonChar: 74,
					},
				},
				squareLLC: LatLonChar{
					LatChar: 48,
					LonChar: 48,
				},
			},
			want: Square{
				Decoded: LatLonDeg{
					Lat: 0,
					Lon: 0,
				},
				Encoded: LatLonChar{
					LatChar: 48,
					LonChar: 48,
				},
			},
		},

		{
			name: "decode-AA00",
			args: args{
				f: Field{
					Encoded: LatLonChar{
						LatChar: 65,
						LonChar: 65,
					},
				},
				squareLLC: LatLonChar{
					LatChar: 48,
					LonChar: 48,
				},
			},
			want: Square{
				Decoded: LatLonDeg{
					Lat: 0,
					Lon: 0,
				},
				Encoded: LatLonChar{
					LatChar: 48,
					LonChar: 48,
				},
			},
		},

		{
			name: "decode-JJ99",
			args: args{
				f: Field{
					Encoded: LatLonChar{
						LatChar: 74,
						LonChar: 74,
					},
				},
				squareLLC: LatLonChar{
					LatChar: 57,
					LonChar: 57,
				},
			},
			want: Square{
				Decoded: LatLonDeg{
					Lat: 9,
					Lon: 18,
				},
				Encoded: LatLonChar{
					LatChar: 57,
					LonChar: 57,
				},
			},
		},

		{
			name: "decode-AA99",
			args: args{
				f: Field{
					Encoded: LatLonChar{
						LatChar: 65,
						LonChar: 65,
					},
				},
				squareLLC: LatLonChar{
					LatChar: 57,
					LonChar: 57,
				},
			},
			want: Square{
				Decoded: LatLonDeg{
					Lat: 9,
					Lon: 18,
				},
				Encoded: LatLonChar{
					LatChar: 57,
					LonChar: 57,
				},
			},
		},

		{
			name: "decode-S59ABC-JN76TO",
			args: args{
				f: Field{
					Encoded: LatLonChar{
						LatChar: 78,
						LonChar: 74,
					},
				},
				squareLLC: LatLonChar{
					LatChar: 54,
					LonChar: 55,
				},
			},
			want: Square{
				Decoded: LatLonDeg{
					Lat: 6,
					Lon: 14,
				},
				Encoded: LatLonChar{
					LatChar: 54,
					LonChar: 55,
				},
			},
		},

		{
			name: "decode-K1TTT-FN32LL",
			args: args{
				f: Field{
					Encoded: LatLonChar{
						LatChar: 78, // N
						LonChar: 70, // F
					},
				},
				squareLLC: LatLonChar{
					LatChar: 50, // 2
					LonChar: 51, // 3
				},
			},
			want: Square{
				Decoded: LatLonDeg{
					Lat: 2,
					Lon: 6,
				},
				Encoded: LatLonChar{
					LatChar: 50,
					LonChar: 51,
				},
			},
		},

		{
			name: "decode--JN32",
			args: args{
				f: Field{
					Encoded: LatLonChar{
						LatChar: 78, // N
						LonChar: 74, // J
					},
				},
				squareLLC: LatLonChar{
					LatChar: 50, // 2
					LonChar: 51, // 3
				},
			},
			want: Square{
				Decoded: LatLonDeg{
					Lat: 2,
					Lon: 6,
				},
				Encoded: LatLonChar{
					LatChar: 50,
					LonChar: 51,
				},
			},
		},

		{
			name: "decode-PS2T-GG58WG",
			args: args{
				f: Field{
					Encoded: LatLonChar{
						LatChar: 71, // G
						LonChar: 71, // G
					},
				},
				squareLLC: LatLonChar{
					LatChar: 56, // 8
					LonChar: 53, // 5
				},
			},
			want: Square{
				Decoded: LatLonDeg{
					Lat: 8,
					Lon: 10,
				},
				Encoded: LatLonChar{
					LatChar: 56,
					LonChar: 53,
				},
			},
		},

		{
			name: "decode-ZM4T-RF80LQ",
			args: args{
				f: Field{
					Encoded: LatLonChar{
						LatChar: 70, // F
						LonChar: 82, // R
					},
				},
				squareLLC: LatLonChar{
					LatChar: 48, // 0
					LonChar: 56, // 8
				},
			},
			want: Square{
				Decoded: LatLonDeg{
					Lat: 0,
					Lon: 16,
				},
				Encoded: LatLonChar{
					LatChar: 48,
					LonChar: 56,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SquareDecode(tt.args.squareLLC); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SquareDecode() = %v, want %v", got, tt.want)
			}
		})
	}
}
