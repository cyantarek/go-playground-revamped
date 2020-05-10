package endpoints

import (
	"backend/api/playground"
	playgroundsvc "backend/internal/services/playground"
	"context"
	"reflect"
	"testing"
)

func TestPlaygroundEndpoint_Ping(t *testing.T) {
	type fields struct {
		pgService *playgroundsvc.Service
	}
	type args struct {
		ctx context.Context
		in1 *playground.EmptyRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *playground.PingResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PlaygroundEndpoint{
				pgService: tt.fields.pgService,
			}
			got, err := p.Ping(tt.args.ctx, tt.args.in1)
			if (err != nil) != tt.wantErr {
				t.Errorf("Ping() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ping() got = %v, want %v", got, tt.want)
			}
		})
	}
}
