package rolexserver

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)
func TestGenToken(t *testing.T) {
	token, _ := generateToken("192.168.0.1", 24 * 365 * time.Hour, []byte("my-rolex-server"))
	fmt.Printf("token=%s\n", token)

	_, err := VerifyToken("192.168.0.1", 
	"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJyb2xleC1zZXJ2ZXIiLCJzdWIiOiIxOTIuMTY4LjAuMSIsImV4cCI6MTcwMDQ1Mjk5OSwiaWF0IjoxNzAwNDUyOTk5fQ.OgPHS49E-zSyjVoxcmybi44nQunTs2TpsYzPI2JkvMU", 
	[]byte("my-rolex-server"))
	assert.NoError(t, err)
}
