package domain

import (
	"reflect"
	"testing"
)

func TestNewNodeInteger(t *testing.T) {
	type args[T any] struct {
		data T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want *Node[T]
	}
	tests := []testCase[any]{
		{
			name: "Test with int",
			args: args[any]{data: 42},
			want: &Node[any]{Data: 42, Next: nil},
		},
		{
			name: "Test with string",
			args: args[any]{data: "Hello"},
			want: &Node[any]{Data: "Hello", Next: nil},
		},
		{
			name: "Test with double",
			args: args[any]{data: 35.8},
			want: &Node[any]{Data: 35.8, Next: nil},
		},
		{
			name: "Test with boolean",
			args: args[any]{data: true},
			want: &Node[any]{Data: true, Next: nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNode(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
