package server

import "testing"

func Test_queryType(t *testing.T) {
	type args struct {
		query string
	}
	tests := []struct {
		name string
		args args
		want qt
	}{

		{
			name: "empty",
			args: args{
				query: "",
			},
			want: qtUnsupported,
		},

		{
			name: "xy",
			args: args{
				query: "xy",
			},
			want: qtQthLocator,
		},

		{
			name: "qtQthPosition",
			args: args{
				query: "lat=46.604&lon=15.625",
			},
			want: qtQthPosition,
		},
		{
			name: "qtQthLocator",
			args: args{
				query: "jn76to",
			},
			want: qtQthLocator,
		},
		{
			name: "qtDistLocator",
			args: args{
				query: "jn76to;jn76PO",
			},
			want: qtDistLocator,
		},
		{
			name: "qtDistPosition",
			args: args{
				query: "lat=46.604&lon=15.625;lat=46.604&lon=15.291",
			},
			want: qtDistPosition,
		},
		{
			name: "qtDistLocatorPosition",
			args: args{
				query: "jn76to;lat=46.604&lon=15.291",
			},
			want: qtDistLocatorPosition,
		},
		{
			name: "qtDistPositionLocator",
			args: args{
				query: "lat=46.604&lon=15.625;jn76PO",
			},
			want: qtDistPositionLocator,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := queryType(tt.args.query); got != tt.want {
				t.Errorf("queryType() = %v, want %v", got, tt.want)
			}
		})
	}
}
