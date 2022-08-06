package config

import (
	"strings"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSetupDefault(t *testing.T) {
	testCases := []struct {
		name    string
		sample  *viper.Viper
		expect  string
		wantErr bool
	}{
		{
			name:    "Should error without token",
			sample:  setupWithoutToken(),
			expect:  "token is required",
			wantErr: true,
		},
		{
			name:    "Should error without repo",
			sample:  setupWithoutRepo(),
			expect:  "repo is required",
			wantErr: true,
		},
		{
			name:   "Should not error and other fields has default value as expected",
			sample: setupComplete(),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := SetupDefault(tc.sample)

			switch tc.wantErr {
			case true:
				require.Error(t, err)
				assert.Equal(t, tc.expect, err.Error())
			case false:
				require.NoError(t, err)
				assert.Equal(t, "amd64", tc.sample.Get("arch"))
				assert.Equal(t, "latest", tc.sample.Get("tag"))
			}
		})
	}
}

func setupWithoutToken() *viper.Viper {
	v := viper.New()
	v.SetConfigType("json")
	jsonTest := `{"repo":"org/hello-world"}`
	v.ReadConfig(strings.NewReader(jsonTest))
	return v
}

func setupWithoutRepo() *viper.Viper {
	v := viper.New()
	v.SetConfigType("json")
	jsonTest := `{"token":"abc"}`
	v.ReadConfig(strings.NewReader(jsonTest))
	return v
}

func setupComplete() *viper.Viper {
	v := viper.New()
	v.SetConfigType("json")
	jsonTest := `{"token":"abc","repo":"org/hello-world"}`
	v.ReadConfig(strings.NewReader(jsonTest))
	return v
}
