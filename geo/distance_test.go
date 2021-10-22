package geo

import (
	"fmt"
	"github.com/golang/geo/s2"
	"github.com/s51ds/qthgeo/geo/internal"
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
		a QTH
		b QTH
	}
	tests := []struct {
		name string
		args args
		want float64
	}{

		{
			name: "zero qth",
			args: args{
				a: QTH{
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
				b: QTH{
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
		a    QTH
		b    QTH
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

func TestQTH_AzimuthTo(t *testing.T) {
	S59ABC, _ := NewQthFromLocator("JN76TO")
	aa00aa, _ := NewQthFromLocator("AA00AA")
	E74G, _ := NewQthFromLocator("JN94FQ")
	LZ9X, _ := NewQthFromLocator("KN22XS")
	S52ME, _ := NewQthFromLocator("JN76TM")
	S57M, _ := NewQthFromLocator("JN76PO")
	ZM4T, _ := NewQthFromLocator("RF80LQ")
	K1TTT, _ := NewQthFromLocator("FN32LL")

	tests := []struct {
		name string
		a    QTH
		b    QTH
		want float64
	}{
		{
			name: "S59ABC-S59ABC",
			a:    S59ABC,
			b:    S59ABC,
			want: 0,
		},
		{
			name: "aa00aa",
			a:    aa00aa,
			b:    aa00aa,
			want: 0,
		},
		{
			name: "S59ABC-E74G",
			a:    S59ABC,
			b:    E74G,
			want: 133.03748572823045,
		},
		{
			name: "E74G-S59ABC",
			a:    E74G,
			b:    S59ABC,
			want: 315.0638953021112,
		},

		{
			name: "S59ABC-LZ9X",
			a:    S59ABC,
			b:    LZ9X,
			want: 113.84651969013157,
		},
		{
			name: "LZ9X-S59ABC",
			a:    LZ9X,
			b:    S59ABC,
			want: 301.12735573508303,
		},

		{
			name: "S59ABC-S52ME",
			a:    S59ABC,
			b:    S52ME,
			want: 180,
		},
		{
			name: "S52ME-S59ABC",
			a:    S52ME,
			b:    S59ABC,
			want: 0,
		},

		{
			name: "S59ABC-S57M",
			a:    S59ABC,
			b:    S57M,
			want: 270.1211042671341,
		},
		{
			name: "S57M-S59ABC",
			a:    S57M,
			b:    S59ABC,
			want: 89.87889573286589,
		},

		{
			name: "S59ABC-ZM4T",
			a:    S59ABC,
			b:    ZM4T,
			want: 68.53963523265749,
		},
		{
			name: "ZM4T-S59ABC",
			a:    ZM4T,
			b:    S59ABC,
			want: 304.2672317587961,
		},

		{
			name: "S59ABC-K1TTT",
			a:    S59ABC,
			b:    K1TTT,
			want: 301.48121380922015,
		},
		{
			name: "K1TTT-S59ABC",
			a:    K1TTT,
			b:    S59ABC,
			want: 52.60153719955525,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.AzimuthTo(tt.b); got != tt.want {
				t.Errorf("AzimuthTo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQTH_DistanceAndAzimuth(t *testing.T) {
	S59ABC, _ := NewQthFromLocator("JN76TO")
	DL7ULM, _ := NewQthFromLocator("JO62MS")

	dist, azim := S59ABC.DistanceAndAzimuth(DL7ULM)
	//710.3165135631285 345.8128442875058
	if dist != 710.3165135631285 {
		fmt.Println("dist", dist)
		t.Error("wtf-1")
	}
	if azim != 345.8128442875058 {
		fmt.Println("azim", azim)
		t.Error("wtf-2")
	}

	dist, azim = DL7ULM.DistanceAndAzimuth(S59ABC)
	//710.3165135631285 345.8128442875058
	if dist != 710.3165135631285 {
		fmt.Println("dist", dist)
		t.Error("wtf-3")
	}
	if azim != 163.83998723945137 {
		fmt.Println("azim", azim)
		t.Error("wtf-4")
	}

}
