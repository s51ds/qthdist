package server

import "testing"

func Test_parseReqLatLon(t *testing.T) {
	type args struct {
		rawRequest string
	}
	tests := []struct {
		name     string
		args     args
		wantLatF float64
		wantLonF float64
		wantErr  bool
	}{
		{
			name: "ok-1",
			args: args{
				rawRequest: "lat=46.604&lon=15.625",
			},
			wantLatF: 46.604,
			wantLonF: 15.625,
			wantErr:  false,
		},
		{
			name: "ok-2",
			args: args{
				rawRequest: "lon=15.625&lat=46.604",
			},
			wantLatF: 46.604,
			wantLonF: 15.625,
			wantErr:  false,
		},
		{
			name: "nok-1",
			args: args{
				rawRequest: "lat=46.604&&lon=15.625",
			},
			wantLatF: 0,
			wantLonF: 0,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLatF, gotLonF, err := parseReqLatLon(tt.args.rawRequest)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseReqLatLon() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotLatF != tt.wantLatF {
				t.Errorf("parseReqLatLon() gotLatF = %v, want %v", gotLatF, tt.wantLatF)
			}
			if gotLonF != tt.wantLonF {
				t.Errorf("parseReqLatLon() gotLonF = %v, want %v", gotLonF, tt.wantLonF)
			}
		})
	}
}
