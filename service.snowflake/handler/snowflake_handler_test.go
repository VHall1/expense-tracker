package handler

// import (
// 	"bytes"
// 	"encoding/base64"
// 	"encoding/binary"
// 	"net/http"
// 	"net/http/httptest"
// 	"strings"
// 	"testing"
// )

// type MockIDGenerator struct {
// 	IDToReturn int64
// }

// func (m *MockIDGenerator) GenerateID() int64 {
// 	return m.IDToReturn
// }

// func TestSnowflakeHandler_GenerateID(t *testing.T) {
// 	mockIDGen := &MockIDGenerator{IDToReturn: 1234567890}
// 	h := NewSnowflakeHandler(mockIDGen)

// 	req, err := http.NewRequest("GET", "/snowflake", nil)
// 	if err != nil {
// 		t.Fatalf("Could not create request: %v", err)
// 	}

// 	rr := httptest.NewRecorder()
// 	h.GenerateID().ServeHTTP(rr, req)

// 	if status := rr.Code; status != http.StatusOK {
// 		t.Errorf("Handler returned wrong status code: got %v, want %v", status, http.StatusOK)
// 	}

// 	// Create a byte buffer to store the binary representation
// 	var buf bytes.Buffer
// 	// Write the int64 into the buffer in big-endian byte order
// 	err = binary.Write(&buf, binary.BigEndian, int64(1234567890))
// 	if err != nil {
// 		t.Errorf("Error writing int64 to buffer: %v", err)
// 	}

// 	expectedID := base64.StdEncoding.EncodeToString(buf.Bytes())
// 	if strings.TrimSpace(rr.Body.String()) != expectedID {
// 		t.Errorf("Unexpected response body: got %q, want %q", rr.Body.String(), expectedID)
// 	}
// }
