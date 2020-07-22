package geo

import (
	"github.com/golang/geo/s2"
	"qth/geo/internal"
	"testing"
)

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
			got, err := LocatorDistance(tt.args.locatorA, tt.args.locatorB)
			if (err != nil) != tt.wantErr {
				t.Errorf("LocatorDistance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("LocatorDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistanceQTH(t *testing.T) {
	jn76to, _ := NewQthFromLocator("JN76TO")
	aa00aa, _ := NewQthFromLocator("AA00AA")

	type args struct {
		a *QTH
		b *QTH
	}
	tests := []struct {
		name string
		args args
		want float64
	}{

		{
			name: "zero qth",
			args: args{
				a: &QTH{
					Loc: "",
					LatLon: internal.LatLonDeg{
						Lat: 0,
						Lon: 0,
					},
					LatLng: s2.LatLng{
						Lat: 0,
						Lng: 0,
					},
				},
				b: &QTH{
					Loc: "",
					LatLon: internal.LatLonDeg{
						Lat: 0,
						Lon: 0,
					},
					LatLng: s2.LatLng{
						Lat: 0,
						Lng: 0,
					},
				},
			},
			want: 0,
		},
		{
			name: "jn76to",
			args: args{jn76to, jn76to},
			want: 0,
		},
		{
			name: "aa00aa",
			args: args{aa00aa, aa00aa},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.a.Distance(tt.args.b); got != tt.want {

				//			if got := DistanceQTH(*tt.args.a, *tt.args.b); got != tt.want {
				t.Errorf("DistanceQTH() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQTH_Distance(t *testing.T) {
	jn76to, _ := NewQthFromLocator("JN76TO")
	aa00aa, _ := NewQthFromLocator("AA00AA")
	JN94FQ, _ := NewQthFromLocator("JN94FQ")
	KN22XS, _ := NewQthFromLocator("KN22XS")
	JN76TM, _ := NewQthFromLocator("JN76TM")

	tests := []struct {
		name string
		a    *QTH
		b    *QTH
		want float64
	}{
		{
			name: "jn76to",
			a:    jn76to,
			b:    jn76to,
			want: 0,
		},
		{
			name: "aa00aa",
			a:    aa00aa,
			b:    aa00aa,
			want: 0,
		},
		{
			name: "E74G",
			a:    jn76to,
			b:    JN94FQ,
			want: 306.4447689992149,
		},
		{
			name: "LZ9X",
			a:    jn76to,
			b:    KN22XS,
			want: 920.3974068704536,
		},
		{
			name: "S52ME",
			a:    jn76to,
			b:    JN76TM,
			want: 9.266256686128344,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := tt.a
			if got := a.Distance(tt.b); got != tt.want {
				t.Errorf("Distance() = %v, want %v", got, tt.want)
			}
		})
	}
}
