package test

import (
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"bwind.com/config-management-service/app"
)

func TestConfig(t *testing.T) {

	app := app.NewApp()

	// Create Config
	payload := `{"config_name":"payment_config","data":{"max_limit":1000,"enabled":true}}`
	req := httptest.NewRequest("POST", "/api/v1/config", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)
	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}

	if !strings.Contains(string(body), "Success") {
		t.Errorf("expected body to contain 'Success', got %s", string(body))
	}

	// Update Config
	payload = `{"max_limit":2000,"enabled":true}`
	req = httptest.NewRequest("POST", "/api/v1/config/payment_config", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	resp, _ = app.Test(req)
	body, _ = io.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}

	if !strings.Contains(string(body), "Success") {
		t.Errorf("expected body to contain 'Success', got %s", string(body))
	}

	// Rollback Config
	req = httptest.NewRequest("POST", "/api/v1/config/payment_config/rollback/1", nil)
	resp, _ = app.Test(req)
	body, _ = io.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}

	if !strings.Contains(string(body), "Success") {
		t.Errorf("expected body to contain 'Success', got %s", string(body))
	}

	// Fetch Config
	req = httptest.NewRequest("GET", "/api/v1/config/payment_config", nil)
	resp, _ = app.Test(req)
	body, _ = io.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}

	if !strings.Contains(string(body), "Success") {
		t.Errorf("expected body to contain 'Success', got %s", string(body))
	}

	// List Versions
	req = httptest.NewRequest("GET", "/api/v1/config/payment_config/versions", nil)
	resp, _ = app.Test(req)
	body, _ = io.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}

	if !strings.Contains(string(body), "Success") {
		t.Errorf("expected body to contain 'Success', got %s", string(body))
	}

}
