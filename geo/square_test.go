package geo

import (
	"reflect"
	"testing"
)

//func TestSquare_Decode_01(t *testing.T) {
//	fc := latLonChar{}
//	fc.setLonChar("J")
//	fc.setLatChar("N")
//	sc := latLonChar{}
//	sc.setLonChar("7")
//	sc.setLatChar("6")
//	f := fieldDecode(fc)
//	s := squareDecode(f, sc)
//	fmt.Println(f.String())
//	fmt.Println("s")
//	fmt.Println(s.String())
//	//
//	//
//	lld := LatLonDeg{}
//	lld.Lon = f.decoded.Lon + s.decoded.Lon
//	lld.Lat = f.decoded.Lat + s.decoded.Lat
//
//	_, sa := squareEncode(lld)
//	fmt.Println("sa")
//	fmt.Println(sa.String())
//	if s.Equals(sa) {
//		t.Fatal()
//	}
//}

func TestSquare_String(t *testing.T) {
	type fields struct {
		decoded LatLonDeg
		encoded latLonChar
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
			name: "toString-decoded-zero-2",
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
				encoded: latLonChar{
					latChar: 0,
					lonChar: 0,
				},
			},
			want: "Decoded:{0.000000 0.000000}",
		},
		{
			name: "toString-decoded-1",
			fields: fields{
				decoded: LatLonDeg{
					Lat: 1,
					Lon: 2,
				},
			},
			want: "Decoded:{1.000000 2.000000}",
		},
		{
			name: "toString-decoded-2",
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
			a := &square{
				decoded: tt.fields.decoded,
				encoded: tt.fields.encoded,
			}
			if got := a.String(); got != tt.want {
				t.Errorf("square.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSquare_Equals(t *testing.T) {
	type fields struct {
		decoded LatLonDeg
		encoded latLonChar
	}
	type args struct {
		b square
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Equals-1",
			fields: fields{
				decoded: LatLonDeg{
					Lat: 0,
					Lon: 0,
				},
				encoded: latLonChar{
					latChar: 0,
					lonChar: 0,
				},
			},
			args: args{
				b: square{
					decoded: LatLonDeg{
						Lat: 0,
						Lon: 0,
					},
					encoded: latLonChar{
						latChar: 0,
						lonChar: 0,
					},
				},
			},
			want: true,
		},
		{
			name: "Equals-2",
			fields: fields{
				decoded: LatLonDeg{
					Lat: 1,
					Lon: 2,
				},
				encoded: latLonChar{
					latChar: 48,
					lonChar: 49,
				},
			},
			args: args{
				b: square{
					decoded: LatLonDeg{
						Lat: 0,
						Lon: 0,
					},
					encoded: latLonChar{
						latChar: 0,
						lonChar: 0,
					},
				},
			},
			want: false,
		},
		{
			name: "Equals-3",
			fields: fields{
				decoded: LatLonDeg{
					Lat: 1,
					Lon: 2,
				},
				encoded: latLonChar{
					latChar: 48,
					lonChar: 49,
				},
			},
			args: args{
				b: square{
					decoded: LatLonDeg{
						Lat: 1,
						Lon: 2,
					},
					encoded: latLonChar{
						latChar: 0,
						lonChar: 0,
					},
				},
			},
			want: false,
		},

		{
			name: "Equals-4",
			fields: fields{
				decoded: LatLonDeg{
					Lat: 1,
					Lon: 2,
				},
				encoded: latLonChar{
					latChar: 48,
					lonChar: 49,
				},
			},
			args: args{
				b: square{
					decoded: LatLonDeg{
						Lat: 1,
						Lon: 2,
					},
					encoded: latLonChar{
						latChar: 48,
						lonChar: 49,
					},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &square{
				decoded: tt.fields.decoded,
				encoded: tt.fields.encoded,
			}
			if got := a.Equals(tt.args.b); got != tt.want {
				t.Errorf("square.Equals() = %v, want %v", got, tt.want)
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
		want  field
		want1 square
	}{
		{
			name: "encode-JJ00-1",
			args: args{
				lld: LatLonDeg{
					Lat: 0,
					Lon: 0,
				},
			},
			want: field{
				decoded: LatLonDeg{
					Lat: 0,
					Lon: 0,
				},
				encoded: latLonChar{
					latChar: 74,
					lonChar: 74,
				},
			},
			want1: square{
				decoded: LatLonDeg{
					Lat: 0,
					Lon: 0,
				},
				encoded: latLonChar{
					latChar: 48,
					lonChar: 48,
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
			want: field{
				decoded: LatLonDeg{
					Lat: 0,
					Lon: 0,
				},
				encoded: latLonChar{
					latChar: 74,
					lonChar: 74,
				},
			},
			want1: square{
				decoded: LatLonDeg{
					Lat: 0,
					Lon: 0,
				},
				encoded: latLonChar{
					latChar: 48,
					lonChar: 48,
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
			want: field{
				decoded: LatLonDeg{
					Lat: 0,
					Lon: 0,
				},
				encoded: latLonChar{
					latChar: 74,
					lonChar: 74,
				},
			},
			want1: square{
				decoded: LatLonDeg{
					Lat: 0,
					Lon: 0,
				},
				encoded: latLonChar{
					latChar: 48,
					lonChar: 48,
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
			want: field{
				decoded: LatLonDeg{
					Lat: 0,
					Lon: 0,
				},
				encoded: latLonChar{
					latChar: 74,
					lonChar: 74,
				},
			},
			want1: square{
				decoded: LatLonDeg{
					Lat: 9,
					Lon: 18,
				},
				encoded: latLonChar{
					latChar: 57,
					lonChar: 57,
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
			want: field{
				decoded: LatLonDeg{
					Lat: 10,
					Lon: 20,
				},
				encoded: latLonChar{
					latChar: 75,
					lonChar: 75,
				},
			},
			want1: square{
				decoded: LatLonDeg{
					Lat: 0,
					Lon: 0,
				},
				encoded: latLonChar{
					latChar: 48,
					lonChar: 48,
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
			want: field{
				decoded: LatLonDeg{
					Lat: -90,
					Lon: -180,
				},
				encoded: latLonChar{
					latChar: 65,
					lonChar: 65,
				},
			},
			want1: square{
				decoded: LatLonDeg{
					Lat: 0,
					Lon: 0,
				},
				encoded: latLonChar{
					latChar: 48,
					lonChar: 48,
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
			want: field{
				decoded: LatLonDeg{
					Lat: -90,
					Lon: -180,
				},
				encoded: latLonChar{
					latChar: 65,
					lonChar: 65,
				},
			},
			want1: square{
				decoded: LatLonDeg{
					Lat: 0,
					Lon: 0,
				},
				encoded: latLonChar{
					latChar: 48,
					lonChar: 48,
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
			want: field{
				decoded: LatLonDeg{
					Lat: -90,
					Lon: -180,
				},
				encoded: latLonChar{
					latChar: 65,
					lonChar: 65,
				},
			},
			want1: square{
				decoded: LatLonDeg{
					Lat: 9,
					Lon: 18,
				},
				encoded: latLonChar{
					latChar: 57,
					lonChar: 57,
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
			want: field{
				decoded: LatLonDeg{
					Lat: -80,
					Lon: -160,
				},
				encoded: latLonChar{
					latChar: 66,
					lonChar: 66,
				},
			},
			want1: square{
				decoded: LatLonDeg{
					Lat: 0,
					Lon: 0,
				},
				encoded: latLonChar{
					latChar: 48,
					lonChar: 48,
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
			want: field{
				decoded: LatLonDeg{
					Lat: 40,
					Lon: 0,
				},
				encoded: latLonChar{
					latChar: 78,
					lonChar: 74,
				},
			},
			want1: square{
				decoded: LatLonDeg{
					Lat: 6,
					Lon: 14,
				},
				encoded: latLonChar{
					latChar: 54,
					lonChar: 55,
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
			want: field{
				decoded: LatLonDeg{
					Lat: 40,
					Lon: -80,
				},
				encoded: latLonChar{
					latChar: 78,
					lonChar: 70,
				},
			},
			want1: square{
				decoded: LatLonDeg{
					Lat: 2,
					Lon: 6,
				},
				encoded: latLonChar{
					latChar: 50,
					lonChar: 51,
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
			want: field{
				decoded: LatLonDeg{
					Lat: -30,
					Lon: -60,
				},
				encoded: latLonChar{
					latChar: 71,
					lonChar: 71,
				},
			},
			want1: square{
				decoded: LatLonDeg{
					Lat: 8,
					Lon: 10,
				},
				encoded: latLonChar{
					latChar: 56,
					lonChar: 53,
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
			want: field{
				decoded: LatLonDeg{
					Lat: -40,
					Lon: 160,
				},
				encoded: latLonChar{
					latChar: 70,
					lonChar: 82,
				},
			},
			want1: square{
				decoded: LatLonDeg{
					Lat: 0,
					Lon: 16,
				},
				encoded: latLonChar{
					latChar: 48,
					lonChar: 56,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := squareEncode(tt.args.lld)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FielsEncode() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("squareEncode() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSquareDecode(t *testing.T) {
	type args struct {
		f         field
		squareLLC latLonChar
	}
	tests := []struct {
		name string
		args args
		want square
	}{
		{
			name: "decode-zero",
			args: args{
				f: field{
					decoded: LatLonDeg{
						Lat: 0,
						Lon: 0,
					},
					encoded: latLonChar{
						latChar: 0,
						lonChar: 0,
					},
				},
				squareLLC: latLonChar{
					latChar: 0,
					lonChar: 0,
				},
			},
			want: square{
				decoded: LatLonDeg{
					Lat: 0,
					Lon: 0,
				},
				encoded: latLonChar{
					latChar: 0,
					lonChar: 0,
				},
			},
		},

		{
			name: "decode-JJ00",
			args: args{
				f: field{
					encoded: latLonChar{
						latChar: 74,
						lonChar: 74,
					},
				},
				squareLLC: latLonChar{
					latChar: 48,
					lonChar: 48,
				},
			},
			want: square{
				decoded: LatLonDeg{
					Lat: 0,
					Lon: 0,
				},
				encoded: latLonChar{
					latChar: 48,
					lonChar: 48,
				},
			},
		},

		{
			name: "decode-AA00",
			args: args{
				f: field{
					encoded: latLonChar{
						latChar: 65,
						lonChar: 65,
					},
				},
				squareLLC: latLonChar{
					latChar: 48,
					lonChar: 48,
				},
			},
			want: square{
				decoded: LatLonDeg{
					Lat: 0,
					Lon: 0,
				},
				encoded: latLonChar{
					latChar: 48,
					lonChar: 48,
				},
			},
		},

		{
			name: "decode-JJ99",
			args: args{
				f: field{
					encoded: latLonChar{
						latChar: 74,
						lonChar: 74,
					},
				},
				squareLLC: latLonChar{
					latChar: 57,
					lonChar: 57,
				},
			},
			want: square{
				decoded: LatLonDeg{
					Lat: 9,
					Lon: 18,
				},
				encoded: latLonChar{
					latChar: 57,
					lonChar: 57,
				},
			},
		},

		{
			name: "decode-AA99",
			args: args{
				f: field{
					encoded: latLonChar{
						latChar: 65,
						lonChar: 65,
					},
				},
				squareLLC: latLonChar{
					latChar: 57,
					lonChar: 57,
				},
			},
			want: square{
				decoded: LatLonDeg{
					Lat: 9,
					Lon: 18,
				},
				encoded: latLonChar{
					latChar: 57,
					lonChar: 57,
				},
			},
		},

		{
			name: "decode-S59ABC-JN76TO",
			args: args{
				f: field{
					encoded: latLonChar{
						latChar: 78,
						lonChar: 74,
					},
				},
				squareLLC: latLonChar{
					latChar: 54,
					lonChar: 55,
				},
			},
			want: square{
				decoded: LatLonDeg{
					Lat: 6,
					Lon: 14,
				},
				encoded: latLonChar{
					latChar: 54,
					lonChar: 55,
				},
			},
		},

		{
			name: "decode-K1TTT-FN32LL",
			args: args{
				f: field{
					encoded: latLonChar{
						latChar: 78, // N
						lonChar: 70, // F
					},
				},
				squareLLC: latLonChar{
					latChar: 50, // 2
					lonChar: 51, // 3
				},
			},
			want: square{
				decoded: LatLonDeg{
					Lat: 2,
					Lon: 6,
				},
				encoded: latLonChar{
					latChar: 50,
					lonChar: 51,
				},
			},
		},

		{
			name: "decode--JN32",
			args: args{
				f: field{
					encoded: latLonChar{
						latChar: 78, // N
						lonChar: 74, // J
					},
				},
				squareLLC: latLonChar{
					latChar: 50, // 2
					lonChar: 51, // 3
				},
			},
			want: square{
				decoded: LatLonDeg{
					Lat: 2,
					Lon: 6,
				},
				encoded: latLonChar{
					latChar: 50,
					lonChar: 51,
				},
			},
		},

		{
			name: "decode-PS2T-GG58WG",
			args: args{
				f: field{
					encoded: latLonChar{
						latChar: 71, // G
						lonChar: 71, // G
					},
				},
				squareLLC: latLonChar{
					latChar: 56, // 8
					lonChar: 53, // 5
				},
			},
			want: square{
				decoded: LatLonDeg{
					Lat: 8,
					Lon: 10,
				},
				encoded: latLonChar{
					latChar: 56,
					lonChar: 53,
				},
			},
		},

		{
			name: "decode-ZM4T-RF80LQ",
			args: args{
				f: field{
					encoded: latLonChar{
						latChar: 70, // F
						lonChar: 82, // R
					},
				},
				squareLLC: latLonChar{
					latChar: 48, // 0
					lonChar: 56, // 8
				},
			},
			want: square{
				decoded: LatLonDeg{
					Lat: 0,
					Lon: 16,
				},
				encoded: latLonChar{
					latChar: 48,
					lonChar: 56,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := squareDecode(tt.args.squareLLC); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("squareDecode() = %v, want %v", got, tt.want)
			}
		})
	}
}
