package auth

import (
	"net/http"
	"testing"
)

func TestApiKey(t *testing.T) {
	testcases := []struct {
		name       string
		key        string
		returnKey  string
		shouldFail bool
	}{
		{
			name:       "CI Failing Test",
			key:        "ApiKey ",
			returnKey:  "",
			shouldFail: true,
		},
		{
			name:       "Correct key",
			key:        "ApiKey testingkeyiscorrect",
			returnKey:  "testingkeyiscorrect",
			shouldFail: false,
		},
		{
			name:       "Empty Api Key",
			key:        "ApiKey ",
			returnKey:  "",
			shouldFail: false,
		},
		{
			name:       "Malformed Header splitted Array too short",
			key:        "ApiKey",
			shouldFail: true,
		},
		{
			name:       "Malformed Header wrong Name",
			key:        "Apikey testingforApiKey",
			shouldFail: true,
		},
		{
			name:       "No Header included",
			key:        "",
			shouldFail: true,
		},
		{
			name:       "Too many parts",
			key:        "ApiKey key extra",
			returnKey:  "key",
			shouldFail: false,
		},
		{
			name:       "Leading space",
			key:        " ApiKey key",
			shouldFail: true,
		},
		{
			name:       "Wrong capitalization",
			key:        "APIKEY key",
			shouldFail: true,
		},
	}

	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {
			header := http.Header{}
			if test.key != "" {
				header.Set("Authorization", test.key)
			}
			got, err := GetAPIKey(header)
			if !test.shouldFail && err != nil {
				t.Errorf("expected no error, got: %v", err)
			}
			if test.shouldFail && err == nil {
				t.Errorf("expected error, got nil")
			}

			if !test.shouldFail && got != test.returnKey {
				t.Errorf("got wrong key back - expected: %v - got: %v", test.returnKey, got)
			}

		})
	}

}
