package security

import "testing"

func Test_jwtManager_GenerateToken(t *testing.T) {
	tests := []struct {
		name    string
		want    bool
		wantErr bool
	}{
		{
			name:    "test jwt token generation",
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// given
			j := jwtManager{}

			// when
			got, err := j.GenerateToken("noah-id")

			// then
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (len(got) != 0) != tt.want {
				t.Errorf("GenerateToken() got = %v, want %v", got, tt.want)
			}
			t.Logf("token: %s", got)
		})
	}
}

func Test_jwtManager_ValidateToken(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{
			name:    "test jwt token validation",
			want:    "noah-id",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// given
			jw := NewJwtManager()
			token, _ := jw.GenerateToken("noah-id")

			// when
			got, err := jw.ValidateToken(token)

			// then
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ValidateToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}
