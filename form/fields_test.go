package form

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFields(t *testing.T) {

	tests := []struct {
		strct interface{}
		want  []field
	}{
		{
			strct: struct {
				Name string
			}{},
			want: []field{{
				Label:       "Name",
				Name:        "Name",
				Type:        "text",
				Placeholder: "Name",
			}},
		},
		{
			strct: struct {
				FullName string
			}{},
			want: []field{{
				Label:       "FullName",
				Name:        "FullName",
				Type:        "text",
				Placeholder: "FullName",
			},
			},
		},
		// NOTE
		{
			strct: struct {
				Name  string
				Email string
				Age   int
			}{},
			want: []field{
				{
					Label:       "Name",
					Name:        "Name",
					Type:        "text",
					Placeholder: "Name",
				},
				{
					Label:       "Email",
					Name:        "Email",
					Type:        "text",
					Placeholder: "Email",
				},
				{
					Label:       "Age",
					Name:        "Age",
					Type:        "text",
					Placeholder: "Age",
				},
			},
		},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%T", tc.strct), func(t *testing.T) {
			got := fields(tc.strct)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("fields() = %v; want %v", got, tc.want)
			}
		})

	}

}

// func Test_fields(t *testing.T) {
// 	type args struct {
// 		strct interface{}
// 	}
// 	tests := []struct {
// 		name         string
// 		args         args
// 		wantMyFields field
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if gotMyFields := fields(tt.args.strct); !reflect.DeepEqual(gotMyFields, tt.wantMyFields) {
// 				t.Errorf("fields() = %v, want %v", gotMyFields, tt.wantMyFields)
// 			}
// 		})
// 	}
// }
