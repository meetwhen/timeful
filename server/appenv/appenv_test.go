package appenv

import "testing"

func TestParse(t *testing.T) {
	t.Parallel()

	testCases := map[string]Environment{
		"":             Development,
		"   ":          Development,
		"development":  Development,
		" DEVELOPMENT": Development,
		"test":         Test,
		" TEST ":       Test,
		"staging":      Staging,
		" StAgInG ":    Staging,
		"production":   Production,
		" PRODUCTION ": Production,
		"invalid":      Development,
	}

	for input, expected := range testCases {
		input := input
		expected := expected

		t.Run(input, func(t *testing.T) {
			t.Parallel()

			if actual := Parse(input); actual != expected {
				t.Fatalf("Parse(%q) = %q, want %q", input, actual, expected)
			}
		})
	}
}

func TestShouldUseReleaseMode(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		ginMode  string
		env      Environment
		expected bool
	}{
		{name: "development defaults to debug", env: Development, expected: false},
		{name: "test defaults to debug", env: Test, expected: false},
		{name: "staging defaults to release", env: Staging, expected: true},
		{name: "production defaults to release", env: Production, expected: true},
		{name: "release override wins in development", ginMode: "release", env: Development, expected: true},
		{name: "production string wins in development", ginMode: " production ", env: Development, expected: true},
		{name: "debug override wins in staging", ginMode: "debug", env: Staging, expected: false},
		{name: "development override wins in production", ginMode: " DEVELOPMENT ", env: Production, expected: false},
	}

	for _, testCase := range testCases {
		testCase := testCase

		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			if actual := ShouldUseReleaseMode(testCase.ginMode, testCase.env); actual != testCase.expected {
				t.Fatalf(
					"ShouldUseReleaseMode(%q, %q) = %t, want %t",
					testCase.ginMode,
					testCase.env,
					actual,
					testCase.expected,
				)
			}
		})
	}
}

func TestPort(t *testing.T) {
	t.Parallel()

	if actual := Port(Development); actual != "3002" {
		t.Fatalf("Port(development) = %q, want %q", actual, "3002")
	}

	if actual := Port(Test); actual != "3003" {
		t.Fatalf("Port(test) = %q, want %q", actual, "3003")
	}

	if actual := Port(Staging); actual != "3004" {
		t.Fatalf("Port(staging) = %q, want %q", actual, "3004")
	}

	if actual := Port(Production); actual != "3005" {
		t.Fatalf("Port(production) = %q, want %q", actual, "3005")
	}
}

func TestResolvePort(t *testing.T) {
	testCases := []struct {
		name     string
		env      Environment
		override string
		want     string
		wantErr  bool
	}{
		{name: "uses environment default when unset", env: Test, want: "3003"},
		{name: "uses environment default when blank", env: Staging, override: "  ", want: "3004"},
		{name: "uses valid override", env: Test, override: "4300", want: "4300"},
		{name: "rejects non numeric override", env: Test, override: "test", wantErr: true},
		{name: "rejects zero", env: Test, override: "0", wantErr: true},
		{name: "rejects port above range", env: Test, override: "65536", wantErr: true},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual, err := ResolvePort(testCase.env, testCase.override)
			if testCase.wantErr {
				if err == nil {
					t.Fatalf("ResolvePort(%q, %q) returned nil error", testCase.env, testCase.override)
				}
				return
			}

			if err != nil {
				t.Fatalf("ResolvePort(%q, %q) returned error: %v", testCase.env, testCase.override, err)
			}
			if actual != testCase.want {
				t.Fatalf("ResolvePort(%q, %q) = %q, want %q", testCase.env, testCase.override, actual, testCase.want)
			}
		})
	}
}
