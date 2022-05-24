package main

import "testing"

func TestRectangle_Print(t *testing.T) {
	type fields struct {
		width  int
		height int
	}
	type args struct {
		test  int
		test2 int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Rectangle{
				width:  tt.fields.width,
				height: tt.fields.height,
			}
			r.Print(tt.args.test, tt.args.test2)
		})
	}
}

func Test(t *testing.T) {
	type fields struct {
		width  int
		height int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Rectangle{
				width:  tt.fields.width,
				height: tt.fields.height,
			}
			r.Test()
		})
	}
}

func Test_test(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_test2(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test2()
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
