package sonyflake

import "testing"

func Test_ID(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{name: "id", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ID()
			if (err != nil) != tt.wantErr {
				t.Errorf("ID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == 0 {
				t.Errorf("ID() got = %v, want 0", got)
			} else {
				t.Logf("ID() got = %v", got)
			}
		})
	}
}

func Benchmark_ID(b *testing.B) {
	for i := 0; i < b.N; i++ { //use b.N for looping
		_, err := ID()
		if err != nil {
			b.Errorf("ID() error = %v", err)
		}
	}
}
