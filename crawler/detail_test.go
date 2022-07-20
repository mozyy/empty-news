package crawler

import "testing"

func TestDetail(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
	}{
		{"蔚来汽车", args{url: "https://m.cnbeta.com/view/1076513.htm"}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Detail(tt.args.url)
		})
	}
}
