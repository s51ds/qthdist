package geo

import (
	"testing"

	"github.com/golang/geo/s2"
)

func TestNewQthFromLOC_01(t *testing.T) {
	type args struct {
		qthLocator string
	}
	tests := []struct {
		name    string
		args    args
		want    QTH
		wantErr bool
	}{
		{
			name: "Empty",
			args: args{
				qthLocator: "",
			},
			want: QTH{
				Loc: "",
				LatLon: LatLonDeg{
					Lat: 0,
					Lon: 0,
				},
				LatLng: s2.LatLng{
					Lat: 0,
					Lng: 0,
				},
			},
			wantErr: true,
		},
		{
			name: "1",
			args: args{
				qthLocator: "1",
			},
			want: QTH{
				Loc: "",
				LatLon: LatLonDeg{
					Lat: 0,
					Lon: 0,
				},
				LatLng: s2.LatLng{
					Lat: 0,
					Lng: 0,
				},
			},
			wantErr: true,
		},
		{
			name: "123",
			args: args{
				qthLocator: "123",
			},
			want: QTH{
				Loc: "",
				LatLon: LatLonDeg{
					Lat: 0,
					Lon: 0,
				},
				LatLng: s2.LatLng{
					Lat: 0,
					Lng: 0,
				},
			},
			wantErr: true,
		},
		{
			name: "12345",
			args: args{
				qthLocator: "12345",
			},
			want: QTH{
				Loc: "",
				LatLon: LatLonDeg{
					Lat: 0,
					Lon: 0,
				},
				LatLng: s2.LatLng{
					Lat: 0,
					Lng: 0,
				},
			},
			wantErr: true,
		},

		{
			name: "1234567",
			args: args{
				qthLocator: "1234567",
			},
			want: QTH{
				Loc: "",
				LatLon: LatLonDeg{
					Lat: 0,
					Lon: 0,
				},
				LatLng: s2.LatLng{
					Lat: 0,
					Lng: 0,
				},
			},
			wantErr: true,
		},

		{
			name: "JN",
			args: args{
				qthLocator: "jN",
			},
			want: QTH{
				Loc: "JN",
				LatLon: LatLonDeg{
					Lat: 45,
					Lon: 10,
				},
				LatLng: s2.LatLng{
					Lat: 0.7853981633974483,
					Lng: 0.17453292519943295,
				},
			},
			wantErr: false,
		},

		{
			name: "76",
			args: args{
				qthLocator: "76",
			},
			want: QTH{
				Loc: "",
				LatLon: LatLonDeg{
					Lat: 0,
					Lon: 0,
				},
				LatLng: s2.LatLng{
					Lat: 0,
					Lng: 0,
				},
			},
			wantErr: true,
		},

		{
			name: "JN76",
			args: args{
				qthLocator: "JN76",
			},
			want: QTH{
				Loc: "JN76",
				LatLon: LatLonDeg{
					Lat: 46.5,
					Lon: 15,
				},
				LatLng: s2.LatLng{
					Lat: 0.8115781021773633,
					Lng: 0.2617993877991494,
				},
			},
			wantErr: false,
		},

		{
			name: "76JN",
			args: args{
				qthLocator: "",
			},
			want: QTH{
				Loc: "",
				LatLon: LatLonDeg{
					Lat: 0,
					Lon: 0,
				},
				LatLng: s2.LatLng{
					Lat: 0,
					Lng: 0,
				},
			},
			wantErr: true,
		},

		{
			name: "JN76to",
			args: args{
				qthLocator: "JN76to",
			},
			want: QTH{
				Loc: "JN76TO",
				LatLon: LatLonDeg{
					Lat: 46.60416666333334,
					Lon: 15.625000003333334,
				},
				LatLng: s2.LatLng{
					Lat: 0.8133961534233465,
					Lng: 0.27270769568229164,
				},
			},
			wantErr: false,
		},

		{
			name: "K1TTT-FN32LL",
			args: args{
				qthLocator: "FN32LL",
			},
			want: QTH{
				Loc: "FN32LL",
				LatLon: LatLonDeg{
					Lat: 42.47916666333334,
					Lon: -73.04166666333333,
				},
				LatLng: s2.LatLng{
					Lat: 0.7414013217785803,
					Lng: -1.2748175744193473,
				},
			},
			wantErr: false,
		},

		{
			name: "PS2T-GG58WG",
			args: args{
				qthLocator: "GG58WG",
			},
			want: QTH{
				Loc: "GG58WG",
				LatLon: LatLonDeg{
					Lat: -21.72916667,
					Lon: -48.124999996666666,
				},
				LatLng: s2.LatLng{
					Lat: -0.3792455021061122,
					Lng: -0.8399397024640934,
				},
			},
			wantErr: false,
		},

		{
			name: "ZM4T-RF80LQ",
			args: args{
				qthLocator: "RF80LQ",
			},
			want: QTH{
				Loc: "RF80LQ",
				LatLon: LatLonDeg{
					Lat: -39.312500003333334,
					Lon: 176.95833333666667,
				},
				LatLng: s2.LatLng{
					Lat: -0.6861325622484484,
					Lng: 3.088505555566477,
				},
			},
			wantErr: false,
		},
		{
			name: "KM3T",
			args: args{
				qthLocator: "FN42ET",
			},
			want: QTH{
				Loc: "FN42ET",
				LatLon: LatLonDeg{
					Lat: 42.812499996666666,
					Lon: -71.62499999666667,
				},
				LatLng: s2.LatLng{
					Lat: 0.7472190859518947,
					Lng: -1.250092076682761,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewQthFromLOC(tt.args.qthLocator)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewQthFromLOC() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !QthEqual(got, tt.want) {
				t.Errorf("NewQthFromLOC() = %s, want %s", got.String(), tt.want.String())
			}
		})
	}
}

func TestNewQthFromLatLon(t *testing.T) {
	type args struct {
		latitude  float64
		longitude float64
	}
	tests := []struct {
		name    string
		args    args
		want    QTH
		wantErr bool
	}{
		{
			name: "S50ABC-JN76TO",
			args: args{
				latitude:  46.60416666333334,
				longitude: 15.625000003333334,
			},
			want: QTH{
				Loc: "JN76TO",
				LatLon: LatLonDeg{
					Lat: 46.60416666333334,
					Lon: 15.625000003333334,
				},
				LatLng: s2.LatLng{
					Lat: 0.8133961534233465,
					Lng: 0.27270769568229164,
				},
			},
			wantErr: false,
		},
		{
			name: "K1TTT-FN32LL",
			args: args{
				latitude:  42.47916666333334,
				longitude: -73.04166666333333,
			},
			want: QTH{
				Loc: "FN32LL",
				LatLon: LatLonDeg{
					Lat: 42.47916666333334,
					Lon: -73.04166666333333,
				},
				LatLng: s2.LatLng{
					Lat: 0.7414013217785803,
					Lng: -1.2748175744193473,
				},
			},
			wantErr: false,
		},

		{
			name: "PS2T-GG58WG",
			args: args{
				latitude:  -21.72916667,
				longitude: -48.124999996666666,
			},
			want: QTH{
				Loc: "GG58WG",
				LatLon: LatLonDeg{
					Lat: -21.72916667,
					Lon: -48.124999996666666,
				},
				LatLng: s2.LatLng{
					Lat: -0.3792455021061122,
					Lng: -0.8399397024640934,
				},
			},
			wantErr: false,
		},

		{
			name: "ZM4T-RF80LQ",
			args: args{
				latitude:  -39.312500003333334,
				longitude: 176.95833333666667,
			},
			want: QTH{
				Loc: "RF80LQ",
				LatLon: LatLonDeg{
					Lat: -39.312500003333334,
					Lon: 176.95833333666667,
				},
				LatLng: s2.LatLng{
					Lat: -0.6861325622484484,
					Lng: 3.088505555566477,
				},
			},
			wantErr: false,
		},

		{
			name: "wrong-arg-1",
			args: args{
				latitude:  -90.0001,
				longitude: 180.001,
			},
			want: QTH{
				Loc: "",
				LatLon: LatLonDeg{
					Lat: 0,
					Lon: 0,
				},
				LatLng: s2.LatLng{
					Lat: 0,
					Lng: 0,
				},
			},
			wantErr: true,
		},

		{
			name: "North-pole",
			args: args{
				latitude:  89.99999999999999,
				longitude: -180,
			},
			want: QTH{
				Loc: "AR09AX",
				LatLon: LatLonDeg{
					Lat: 89.99999999999999,
					Lon: -180,
				},
				LatLng: s2.LatLng{
					Lat: 1.5707963267948963,
					Lng: -3.141592653589793,
				},
			},
			wantErr: false,
		},

		{
			name: "South-pole",
			args: args{
				latitude:  -90,
				longitude: -180,
			},
			want: QTH{
				Loc: "AA00AA",
				LatLon: LatLonDeg{
					Lat: -90,
					Lon: -180,
				},
				LatLng: s2.LatLng{
					Lat: -1.5707963267948966,
					Lng: -3.141592653589793,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewQthFromLatLon(tt.args.latitude, tt.args.longitude)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewQthFromLatLon() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !QthEqual(got, tt.want) {
				t.Errorf("NewQthFromLOC() = %s, want %s", got.String(), tt.want.String())
				//				t.Errorf("NewQthFromLatLon() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistance(t *testing.T) {
	type args struct {
		locatorA string
		locatorB string
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "S59ABC-ZM4T",
			args: args{
				locatorA: "JN76TO",
				locatorB: "RF80LQ",
			},
			want:    18299.250785366803,
			wantErr: false,
		},
		{
			name: "North-South-Pole",
			args: args{
				locatorA: "AR09AX",
				locatorB: "AA00AA",
			},
			want:    20010.481313695705,
			wantErr: false,
		},
		{
			name: "S59ABC-S57M",
			args: args{
				locatorA: "jn76to",
				locatorB: "jn76po",
			},
			want:    25.464939494933233,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DistanceLocator(tt.args.locatorA, tt.args.locatorB)
			if (err != nil) != tt.wantErr {
				t.Errorf("DistanceLocator() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DistanceLocator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewQthFromLOC(t *testing.T) {
	type args struct {
		qthLocator string
	}
	tests := []struct {
		name    string
		args    args
		want    QTH
		wantErr bool
	}{
		{
			name: "KM3T",
			args: args{
				qthLocator: "FN42ET",
			},
			want: QTH{
				Loc: "FN42ET",
				LatLon: LatLonDeg{
					Lat: 42.812499996666666,
					Lon: -71.62499999666667,
				},
				LatLng: s2.LatLng{
					Lat: 0.7472190859518947,
					Lng: -1.250092076682761,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewQthFromLOC(tt.args.qthLocator)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewQthFromLOC() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !QthEqual(got, tt.want) {
				t.Errorf("NewQthFromLOC() = %s, want %s", got.String(), tt.want.String())
				//				t.Errorf("NewQthFromLOC() = %v, want %v", got, tt.want)
			}
		})
	}
}
