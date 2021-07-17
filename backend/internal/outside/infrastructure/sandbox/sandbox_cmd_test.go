package sandbox

import (
	"testing"
)

func TestCompileAndRun(t *testing.T) {
	type args struct {
		body []byte
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 error
		want2 float64
	}{
		{
			name: "a",
			args: args{body: []byte(`package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
}`)},
			want:  "Hello, playground\n",
			want2: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := CompileAndRun(tt.args.body)
			if got != tt.want {
				t.Errorf("CompileAndRun() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("CompileAndRun() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("CompileAndRun() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

/*func TestFormatCode(t *testing.T) {
	type args struct {
		body []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Format code",
			args: args{body: []byte(`package main

import (
	"fmt"
)

func main() {
fmt.Println("Hello, playground")
}`)},
			want: `
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
}
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatCode(tt.args.body); got != tt.want {
				t.Errorf("FormatCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
*/
