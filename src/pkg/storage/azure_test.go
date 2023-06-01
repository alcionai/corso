package storage

import (
	"reflect"
	"testing"
)

func TestAzureConfig_Normalize(t *testing.T) {
	type fields struct {
		Container string
		Prefix    string
	}
	tests := []struct {
		name   string
		fields fields
		want   AzureConfig
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := AzureConfig{
				Container: tt.fields.Container,
				Prefix:    tt.fields.Prefix,
			}
			if got := c.Normalize(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AzureConfig.Normalize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAzureConfig_StringConfig(t *testing.T) {
	type fields struct {
		Container string
		Prefix    string
	}
	tests := []struct {
		name    string
		fields  fields
		want    map[string]string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := AzureConfig{
				Container: tt.fields.Container,
				Prefix:    tt.fields.Prefix,
			}
			got, err := c.StringConfig()
			if (err != nil) != tt.wantErr {
				t.Errorf("AzureConfig.StringConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AzureConfig.StringConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_AzureConfig(t *testing.T) {
	type fields struct {
		Provider storageProvider
		Config   map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		want    AzureConfig
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Storage{
				Provider: tt.fields.Provider,
				Config:   tt.fields.Config,
			}
			got, err := s.AzureConfig()
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.AzureConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Storage.AzureConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
