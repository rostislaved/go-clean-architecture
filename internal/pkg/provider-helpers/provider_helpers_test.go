package providerhelpers

import (
	"testing"
)

type TestStruct1 struct {
	Endpoint1 Endpoint
	Endpoint2 Endpoint
}

type TestStruct2 struct{}

func TestValidateFields(t *testing.T) {
	t.Parallel()

	type args struct {
		endpoints any
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "struct with endpoints, all ok",
			args: args{
				endpoints: TestStruct1{
					Endpoint{
						Method:  "GET",
						Path:    "/api/v1/document",
						Headers: nil,
					},
					Endpoint{
						Method:  "POST",
						Path:    "/api/v1/user",
						Headers: nil,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "struct with endpoints, one has wrong method",
			args: args{
				endpoints: TestStruct1{
					Endpoint{
						Method:  "",
						Path:    "/api/v1/document",
						Headers: nil,
					},
					Endpoint{
						Method:  "POST",
						Path:    "/api/v1/user",
						Headers: nil,
					},
				},
			},
			wantErr: true,
		},
		{
			name: "struct with endpoints, two has wrong methods",
			args: args{
				endpoints: TestStruct1{
					Endpoint{
						Method:  "",
						Path:    "/api/v1/document",
						Headers: nil,
					},
					Endpoint{
						Method:  "POSTasd",
						Path:    "/api/v1/user",
						Headers: nil,
					},
				},
			},
			wantErr: true,
		},
		{
			name: "struct with no endpoints, no error",
			args: args{
				endpoints: TestStruct2{},
			},
			wantErr: false,
		},
		{
			name: "ok endpoint struct itself, no error",
			args: args{
				endpoints: Endpoint{
					Method:  "POST",
					Path:    "/api/v1/user",
					Headers: nil,
				},
			},
			wantErr: false,
		},
		// TODO implement
		// {
		//	name: "wrong endpoint struct itself, error",
		//	args: args{
		//		endpoints: config.Endpoint{
		//			Method:  "POST123",
		//			Path:    "/api/v1/user",
		//			Headers: nil,
		//		},
		//	},
		//	wantErr: true,
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := ValidateEndpoints(tt.args.endpoints)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateEndpoints() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateEndpoint(t *testing.T) {
	t.Parallel()

	type args struct {
		endpoint any
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "ok method, ok path, no error 1",
			args: args{
				endpoint: Endpoint{
					Method:  "GET",
					Path:    "/api/v1",
					Headers: nil,
				},
			},
			wantErr: false,
		},
		{
			name: "ok method, ok path, no error 2",
			args: args{
				endpoint: Endpoint{
					Method:  "POST",
					Path:    "/api/v1",
					Headers: nil,
				},
			},
			wantErr: false,
		},
		{
			name: "wrong method, ok path, error",
			args: args{
				endpoint: Endpoint{
					Method:  "POST1",
					Path:    "/api/v1",
					Headers: nil,
				},
			},
			wantErr: true,
		},
		{
			name: "ok method, wrong path, error",
			args: args{
				endpoint: Endpoint{
					Method:  "POST",
					Path:    "",
					Headers: nil,
				},
			},
			wantErr: true,
		},
		{
			name: "wrong method, wrong path, error",
			args: args{
				endpoint: Endpoint{
					Method:  "paspf",
					Path:    "",
					Headers: nil,
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := ValidateEndpoint(tt.args.endpoint)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateEndpoint() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
