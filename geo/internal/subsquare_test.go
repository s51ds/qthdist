package internal

import (
	"reflect"
	"testing"
)

func TestSubsquareEncode(t *testing.T) {
	type args struct {
		lld LatLonDeg
	}
	tests := []struct {
		name  string
		args  args
		want  Field
		want1 Square
		want2 Subsquare
	}{
		{
			name: "encode-S59ABC-JN76TO",
			args: args{
				lld: LatLonDeg{
					Lat: 46.60333,
					Lon: 15.62333,
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
			want2: Subsquare{
				Decoded: LatLonDeg{
					Lat: 35, //minutes
					Lon: 95, //minutes
				},
				Encoded: LatLonChar{
					LatChar: 79, // O
					LonChar: 84, // T
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
					LatChar: 78, //N
					LonChar: 70, //F
				},
			},
			want1: Square{
				Decoded: LatLonDeg{
					Lat: 2,
					Lon: 6,
				},
				Encoded: LatLonChar{
					LatChar: 50, //2
					LonChar: 51, //3
				},
			},
			want2: Subsquare{
				Decoded: LatLonDeg{
					Lat: 27.5, //minutes
					Lon: 55,   //minutes
				},
				Encoded: LatLonChar{
					LatChar: 76, // L
					LonChar: 76, // L
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
			want2: Subsquare{
				Decoded: LatLonDeg{
					Lat: 15,
					Lon: 110,
				},
				Encoded: LatLonChar{
					LatChar: 71,
					LonChar: 87,
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
					LatChar: 70, // F
					LonChar: 82, // R
				},
			},
			want1: Square{
				Decoded: LatLonDeg{
					Lat: 0,
					Lon: 16,
				},
				Encoded: LatLonChar{
					LatChar: 48, // 0
					LonChar: 56, // 8
				},
			},
			want2: Subsquare{
				Decoded: LatLonDeg{
					Lat: 40,
					Lon: 55,
				},
				Encoded: LatLonChar{
					LatChar: 81, // L
					LonChar: 76, // Q
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := SubsquareEncode(tt.args.lld)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FieldEncode() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("SquareEncode() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("SubsquareEncode() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestSubsquareDecode(t *testing.T) {
	type args struct {
		llc LatLonChar
	}
	tests := []struct {
		name string
		args args
		want Subsquare
	}{
		{
			name: "decode-zero",
			args: args{
				llc: LatLonChar{
					LatChar: 0,
					LonChar: 0,
				},
			},
			want: Subsquare{
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
			name: "decode-AA",
			args: args{
				llc: LatLonChar{
					LatChar: 65,
					LonChar: 65,
				},
			},
			want: Subsquare{
				Decoded: LatLonDeg{
					Lat: 0,
					Lon: 0,
				},
				Encoded: LatLonChar{
					LatChar: 65,
					LonChar: 65,
				},
			},
		},

		{
			name: "decode-XX",
			args: args{
				llc: LatLonChar{
					LatChar: 88,
					LonChar: 88,
				},
			},
			want: Subsquare{
				Decoded: LatLonDeg{
					Lat: 57.5,
					Lon: 115,
				},
				Encoded: LatLonChar{
					LatChar: 88,
					LonChar: 88,
				},
			},
		},

		{
			name: "decode-S59ABC-JN76TO",
			args: args{
				llc: LatLonChar{
					LatChar: 79, // O
					LonChar: 84, // T
				},
			},
			want: Subsquare{
				Decoded: LatLonDeg{
					Lat: 35, //minutes
					Lon: 95, //minutes
				},
				Encoded: LatLonChar{
					LatChar: 79, // O
					LonChar: 84, // T
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SubsquareDecode(tt.args.llc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SubsquareDecode() = %v, want %v", got, tt.want)
			}
		})
	}
}
