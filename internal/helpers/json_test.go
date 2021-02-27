package helpers

import (
	"testing"
)

func TestMustJsonString(t *testing.T) {
	type args struct {
		obj interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"ShouldReturnNullString_WhenNilGiven",
			args{
				obj: nil,
			},
			"null",
		},
		{
			"ShouldReturnEmptyString_WhenGivenInvalidJsonInputType",
			args{
				obj: make(chan int),
			},
			"",
		},
		{
			"ShouldReturnJsonString_WhenMapGiven",
			args{
				obj: map[string]interface{}{
					"test": "test",
					"test2": map[string]string{
						"test21": "test",
					},
				},
			},
			`{"test":"test","test2":{"test21":"test"}}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MustJsonString(tt.args.obj)
			if got != tt.want {
				t.Errorf("MustJsonString() = %v, want %v", got, tt.want)
			}
		})
	}
}
