package playground

import (
	"golang.org/x/net/context"
	"reflect"
	"testing"
)

func Test_playgroundService_RunCode(t *testing.T) {
	type args struct {
		ctx context.Context
		req *CommonCodeRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *CodeRunResponse
		wantErr bool
	}{
		{
			name: "Run Good Code",
			args: args{
				ctx: context.Background(),
				req: &CommonCodeRequest{
					Body: `package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
}`,
				},
			},
			want: &CodeRunResponse{
				Status:  "ok",
				Output:  "Hello, playground\n",
				RunTime: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Service{}
			got, err := p.RunCode(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("RunCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RunCode() got = %v, want %v", got, tt.want)
			}
		})
	}
}
