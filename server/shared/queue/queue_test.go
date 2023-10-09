package queue

import (
	"reflect"
	"testing"
)

func TestNewQueue(t *testing.T) {
	type args struct {
		size uint
	}
	type testCase[T any] struct {
		name string
		args args
		want Queue[T]
	}
	tests := []testCase[int]{
		{
			name: "initializes queue correctly",
			args: args{size: 10},
			want: Queue[int]{items: make([]int, 0, 10), count: 0, size: 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewQueue[int](tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_Pop(t *testing.T) {
	type testCase[T any] struct {
		name    string
		q       Queue[T]
		want    T
		wantErr bool
	}
	tests := []testCase[int]{
		{
			name:    "pops correctly the first item",
			q:       Queue[int]{items: []int{420, 69}, count: 2, size: 10},
			want:    420,
			wantErr: false,
		},
		{
			name:    "returns an error when there's nothing to return",
			q:       Queue[int]{items: []int{}, count: 0, size: 10},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.q.Pop()
			if (err != nil) != tt.wantErr {
				t.Errorf("Pop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pop() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_Push(t *testing.T) {
	type args[T any] struct {
		item T
	}
	type testCase[T any] struct {
		name    string
		q       Queue[T]
		args    args[T]
		wantErr bool
	}
	tests := []testCase[int]{
		{
			name:    "pushes correctly the item",
			q:       Queue[int]{items: []int{}, count: 0, size: 2},
			wantErr: false,
		},
		{
			name:    "returns an error when is full",
			q:       Queue[int]{items: []int{420, 69}, count: 2, size: 2},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.q.Push(tt.args.item); (err != nil) != tt.wantErr {
				t.Errorf("Push() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestQueue_Peek(t *testing.T) {
	type testCase[T any] struct {
		name    string
		q       Queue[T]
		want    T
		wantErr bool
	}
	tests := []testCase[int]{
		{
			name:    "has elements",
			q:       Queue[int]{items: []int{420, 69}, count: 2, size: 2},
			want:    420,
			wantErr: false,
		},
		{
			name:    "is empty",
			q:       Queue[int]{items: []int{}, count: 0, size: 2},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.q.Peek()
			if (err != nil) != tt.wantErr {
				t.Errorf("Peek() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Peek() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_IsFull(t *testing.T) {
	type testCase[T any] struct {
		name string
		q    Queue[T]
		want bool
	}
	tests := []testCase[int]{
		{
			name: "is full",
			q:    Queue[int]{items: []int{420, 69}, count: 2, size: 2},
			want: true,
		},
		{
			name: "has space",
			q:    Queue[int]{items: []int{69}, count: 1, size: 2},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.q.IsFull(); got != tt.want {
				t.Errorf("IsFull() = %v, want %v", got, tt.want)
			}
		})
	}
}
