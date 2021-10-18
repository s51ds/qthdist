package api

import "testing"

func TestDistance(t *testing.T) {
	type args struct {
		locatorA string
		locatorB string
	}
	tests := []struct {
		name         string
		args         args
		wantDistance float64
		wantAzimuth  float64
		wantErr      bool
	}{
		{
			name: "zero",
			args: args{
				locatorA: "JN76TO",
				locatorB: "JN76TO",
			},
			wantDistance: 0,
			wantAzimuth:  0,
			wantErr:      false,
		},
		{
			name: "zero",
			args: args{
				locatorA: "JN76TO",
				locatorB: "jn76TO",
			},
			wantDistance: 0,
			wantAzimuth:  0,
			wantErr:      false,
		},
		{
			name: "zero",
			args: args{
				locatorA: "JN76",
				locatorB: "jn76",
			},
			wantDistance: 0,
			wantAzimuth:  0,
			wantErr:      false,
		},

		{
			name: "N",
			args: args{
				locatorA: "JN76TO",
				locatorB: "JN77TO",
			},
			wantDistance: 111.19508023353235,
			wantAzimuth:  0,
			wantErr:      false,
		},
		{
			name: "N",
			args: args{
				locatorA: "JN",
				locatorB: "JO",
			},
			wantDistance: 1111.950802335329,
			wantAzimuth:  0,
			wantErr:      false,
		},

		{
			name: "E",
			args: args{
				locatorA: "JN76TO",
				locatorB: "JN86TO",
			},
			wantDistance: 152.78565493005073,
			wantAzimuth:  89.27334053790945,
			wantErr:      false,
		},
		{
			name: "E",
			args: args{
				locatorA: "JN",
				locatorB: "KN",
			},
			wantDistance: 1568.5227233314433,
			wantAzimuth:  82.89292388955346,
			wantErr:      false,
		},

		{
			name: "S",
			args: args{
				locatorA: "JN76TO",
				locatorB: "JN75TO",
			},
			wantDistance: 111.19508023353306,
			wantAzimuth:  180,
			wantErr:      false,
		},
		{
			name: "W",
			args: args{
				locatorA: "JN76TO",
				locatorB: "JN66TO",
			},
			wantDistance: 152.78565493005073,
			wantAzimuth:  270.72665946209054,
			wantErr:      false,
		},
		{
			name: "err",
			args: args{
				locatorA: "JN76TO",
				locatorB: "JN76T",
			},
			wantDistance: 0,
			wantAzimuth:  0,
			wantErr:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDistance, gotAzimuth, err := Distance(tt.args.locatorA, tt.args.locatorB)
			if (err != nil) != tt.wantErr {
				t.Errorf("Distance() error = %v, wantErr %v", err, tt.wantErr)
				//				return
			}
			if gotDistance != tt.wantDistance {
				t.Errorf("Distance() gotDistance = %v, want %v", gotDistance, tt.wantDistance)
			}
			if gotAzimuth != tt.wantAzimuth {
				t.Errorf("Distance() gotAzimuth = %v, want %v", gotAzimuth, tt.wantAzimuth)
			}
		})
	}
}
