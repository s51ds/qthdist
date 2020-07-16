package geo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLatLon_Equal(t *testing.T) {
	a := LatLonDeg{}
	b := LatLonDeg{}
	if !a.Equal(b) {
		t.Fatal()
	}
	//
	a.Lon = 1.1
	a.Lat = 2.2
	if a.Equal(b) {
		t.Fatal()
	}
	//
	b.Lon = 1.1
	b.Lat = 2.2
	if !a.Equal(b) {
		t.Fatal()
	}
}

func TestLatLon_String(t *testing.T) {
	a := LatLonDeg{1.12345, -1.12345}
	if a.String() != "{1.123450 -1.123450}" {
		fmt.Println(a.String())
		t.Fatal()
	}
}

func TestLatLonChar_LatLonChar(t *testing.T) {
	a := latLonChar{}
	a.setLatChar("A")
	a.setLonChar("B")
	if a.String() != "BA" {
		t.Fatal()
	}
	a.setLatChar("a")
	a.setLonChar("b")
	if a.String() != "BA" {
		t.Fatal()
	}

	a.setLatChar("AA")
	a.setLonChar("BB")
	if a.String() != "  " {
		t.Fatal()
	}
	a.setLatChar("0")
	a.setLonChar("9")
	if a.String() != "90" {
		t.Fatal()
	}
	a.setLatChar("!")
	a.setLonChar("9")
	if a.String() != "9 " {
		t.Fatal()
	}
	fmt.Println(a.String())
}

func TestLatLonChar_Equal(t *testing.T) {
	a := latLonChar{}
	if !a.Equal(a) {
		t.Fatal()
	}
	b := latLonChar{}
	b.setLatChar("J")
	b.setLonChar("N")
	if a.Equal(b) {
		t.Fatal()
	}
	a.setLatChar("J")
	a.setLonChar("N")
	if !a.Equal(b) {
		t.Fatal()
	}

}

func TestLatLonDeg_ToLatLonDMS(t *testing.T) {
	type fields struct {
		Lat float64
		Lon float64
	}
	tests := []struct {
		name   string
		fields fields
		want   LatLonDMS
	}{
		{
			name: "DMS_0_0",
			fields: fields{
				Lat: 0,
				Lon: 0,
			},
			want: LatLonDMS{
				latDMS: dms{
					degrees: 0,
					minutes: 0,
					seconds: 0,
				},
				lonDMS: dms{
					degrees: 0,
					minutes: 0,
					seconds: 0,
				},
			},
		},

		{
			name: "DMS_180_90",
			fields: fields{
				Lat: 180,
				Lon: 90,
			},
			want: LatLonDMS{
				latDMS: dms{
					degrees: 180,
					minutes: 0,
					seconds: 0,
				},
				lonDMS: dms{
					degrees: 90,
					minutes: 0,
					seconds: 0,
				},
			},
		},

		{
			name: "DMS_-180_-90",
			fields: fields{
				Lat: -180,
				Lon: -90,
			},
			want: LatLonDMS{
				latDMS: dms{
					degrees: -180,
					minutes: 0,
					seconds: 0,
				},
				lonDMS: dms{
					degrees: -90,
					minutes: 0,
					seconds: 0,
				},
			},
		},

		{
			name: "DMS_1_0",
			fields: fields{
				Lat: 1,
				Lon: 0,
			},
			want: LatLonDMS{
				latDMS: dms{
					degrees: 1,
					minutes: 0,
					seconds: 0,
				},
				lonDMS: dms{
					degrees: 0,
					minutes: 0,
					seconds: 0,
				},
			},
		},

		{
			name: "DMS_1.1_0",
			fields: fields{
				Lat: 1.1,
				Lon: 0,
			},
			want: LatLonDMS{
				latDMS: dms{
					degrees: 1,
					minutes: 6,
					seconds: 0,
				},
				lonDMS: dms{
					degrees: 0,
					minutes: 0,
					seconds: 0,
				},
			},
		},

		{
			name: "DMS_1.1_1.1",
			fields: fields{
				Lat: 1.1,
				Lon: 1.1,
			},
			want: LatLonDMS{
				latDMS: dms{
					degrees: 1,
					minutes: 6,
					seconds: 0,
				},
				lonDMS: dms{
					degrees: 1,
					minutes: 6,
					seconds: 0,
				},
			},
		},

		{
			name: "DMS_-1.1_1.1",
			fields: fields{
				Lat: -1.1,
				Lon: 1.1,
			},
			want: LatLonDMS{
				latDMS: dms{
					degrees: -1,
					minutes: 6,
					seconds: 0,
				},
				lonDMS: dms{
					degrees: 1,
					minutes: 6,
					seconds: 0,
				},
			},
		},

		{
			name: "DMS_-1.11_1.1",
			fields: fields{
				Lat: -1.11,
				Lon: 1.1,
			},
			want: LatLonDMS{
				latDMS: dms{
					degrees: -1,
					minutes: 6,
					seconds: 36,
				},
				lonDMS: dms{
					degrees: 1,
					minutes: 6,
					seconds: 0,
				},
			},
		},

		{
			name: "DMS_-1.11_-1.11",
			fields: fields{
				Lat: -1.11,
				Lon: -1.11,
			},
			want: LatLonDMS{
				latDMS: dms{
					degrees: -1,
					minutes: 6,
					seconds: 36,
				},
				lonDMS: dms{
					degrees: -1,
					minutes: 6,
					seconds: 36,
				},
			},
		},

		{
			name: "DMS_S59ABC-JN76TO",
			fields: fields{
				Lat: 46.60333,
				Lon: 15.62333,
			},
			want: LatLonDMS{
				latDMS: dms{
					degrees: 46,
					minutes: 36,
					seconds: 11,
				},
				lonDMS: dms{
					degrees: 15,
					minutes: 37,
					seconds: 23,
				},
			},
		},

		{
			name: "DMS_K1TTT-FN32II",
			fields: fields{
				Lat: 42.4662,
				Lon: -73.0232,
			},
			want: LatLonDMS{
				latDMS: dms{
					degrees: 42,
					minutes: 27,
					seconds: 58,
				},
				lonDMS: dms{
					degrees: -73,
					minutes: 1,
					seconds: 23,
				},
			},
		},

		{
			name: "DMS_PS2T-GG58WG",
			fields: fields{
				Lat: -21.7487,
				Lon: -48.1268,
			},
			want: LatLonDMS{
				latDMS: dms{
					degrees: -21,
					minutes: 44,
					seconds: 55,
				},
				lonDMS: dms{
					degrees: -48,
					minutes: 7,
					seconds: 36,
				},
			},
		},

		{
			name: "DMS_ZM4T-RF80IV",
			fields: fields{
				Lat: -39.109,
				Lon: 176.742,
			},
			want: LatLonDMS{
				latDMS: dms{
					degrees: -39,
					minutes: 6,
					seconds: 32,
				},
				lonDMS: dms{
					degrees: 176,
					minutes: 44,
					seconds: 31,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &LatLonDeg{
				Lat: tt.fields.Lat,
				Lon: tt.fields.Lon,
			}
			if got := a.ToLatLonDMS(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LatLonDeg.ToLatLonDMS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLatLonDMS_String(t *testing.T) {
	type fields struct {
		latDMS dms
		lonDMS dms
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "LatLonDMS_String_0",
			fields: fields{
				latDMS: dms{
					degrees: 0,
					minutes: 0,
					seconds: 0,
				},
				lonDMS: dms{
					degrees: 0,
					minutes: 0,
					seconds: 0,
				},
			},
			want: `Lat=0°0'0", Lon=0°0'0"`,
		},

		{
			name: "LatLonDMS_String_1",
			fields: fields{
				latDMS: dms{
					degrees: 1,
					minutes: 2,
					seconds: 3,
				},
				lonDMS: dms{
					degrees: 4,
					minutes: 5,
					seconds: 6,
				},
			},
			want: `Lat=1°2'3", Lon=4°5'6"`,
		},

		{
			name: "LatLonDMS_String_2",
			fields: fields{
				latDMS: dms{
					degrees: -1,
					minutes: 2,
					seconds: 3,
				},
				lonDMS: dms{
					degrees: -4,
					minutes: 5,
					seconds: 6,
				},
			},
			want: `Lat=-1°2'3", Lon=-4°5'6"`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &LatLonDMS{
				latDMS: tt.fields.latDMS,
				lonDMS: tt.fields.lonDMS,
			}
			if got := a.String(); got != tt.want {
				t.Errorf("LatLonDMS.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
