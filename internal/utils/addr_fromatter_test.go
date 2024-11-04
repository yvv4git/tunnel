package utils

import "testing"

func TestFormatAddrString(t *testing.T) {
	type args struct {
		host string
		port uint16
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Valid host and port",
			args: args{
				host: "localhost",
				port: 8080,
			},
			want:    "localhost:8080",
			wantErr: false,
		},
		{
			name: "Empty host",
			args: args{
				host: "",
				port: 8080,
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Zero port",
			args: args{
				host: "localhost",
				port: 0,
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Valid IPv4 address",
			args: args{
				host: "192.168.1.1",
				port: 80,
			},
			want:    "192.168.1.1:80",
			wantErr: false,
		},
		{
			name: "Valid host with leading and trailing spaces",
			args: args{
				host: "  example.com  ",
				port: 8080,
			},
			want:    "  example.com  :8080",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FormatAddrString(tt.args.host, tt.args.port)
			if (err != nil) != tt.wantErr {
				t.Errorf("FormatAddrString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FormatAddrString() got = %v, want %v", got, tt.want)
			}
		})
	}
}
