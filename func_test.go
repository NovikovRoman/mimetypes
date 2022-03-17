package mimetypes

import (
	"reflect"
	"testing"
)

func TestExtensionByType(t *testing.T) {
	tests := []struct {
		name        string
		contentType string
		want        string
	}{
		{
			name:        "pdf",
			contentType: "application/pdf",
			want:        "pdf",
		},
		{
			name:        "unknown",
			contentType: "application/unknown",
			want:        "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExtensionByType(tt.contentType); got != tt.want {
				t.Errorf("ExtensionByType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_te(t *testing.T) {
	tests := []struct {
		name        string
		contentType string
		want        []string
		wantErr     bool
	}{
		{
			name:        "pdf",
			contentType: "application/pdf",
			want:        []string{".pdf"},
			wantErr:     false,
		},
		{
			name:        "communicator",
			contentType: "application/x-excel",
			want:        nil,
			wantErr:     false,
		},
		{
			name:        "unknown",
			contentType: "application/unknown",
			want:        nil,
			wantErr:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := te(tt.contentType)
			if (err != nil) != tt.wantErr {
				t.Errorf("te() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("te() got = %v, want %v", got, tt.want)
			}
		})
	}
}
