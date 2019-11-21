package boldchat

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"net/url"
	"strings"
	"sync"
	"time"
)

// const definitions of the BoldChat API
// endpoints for the 2 regions available
//		US for North America
// 		EMEA for Europe, Middle East, or Africa
// baseURL is the non-formatted string to use
// for generating the final URL for each
// API Call
const (
	US      = "api.boldchat.com"
	EMEA    = "api-eu.boldchat.com"
	baseURL = "https://%s/aid/%s/data/rest/json/v2/%s?auth=%s%s"
)

// Client is the BoldChat object that is
// used for all operations
type Client struct {
	apiID      string
	apiSetting string
	apiKey     string
	apiHash    string
	authHash   string
	endpoint   string
	hashTime   time.Time
	mx         *sync.RWMutex
}

// New creates a new BoldChat Client object
func New(id, setting, key, end string) *Client {
	c := &Client{
		apiID:      id,
		apiSetting: setting,
		apiKey:     key,
		endpoint:   end,
		mx:         new(sync.RWMutex),
	}
	c.genHash()
	return c
}

// genHash creates a new hash code used for authentication
// locks the mutex for a write to block the authHash
func (c *Client) genHash() {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.hashTime = time.Now()
	tok := fmt.Sprintf("%s:%s:%d000", c.apiID, c.apiSetting, time.Now().Unix())
	h := sha512.Sum512([]byte(tok + c.apiKey))
	c.authHash = fmt.Sprintf("%s:%s", tok, hex.EncodeToString(h[:]))
}

// checkHash checks if the authHash is expired
// generates a new hash if needed
////
// BoldChat Auth tokens expire after 5 minutes
// Due to this, a new token is generated every
// 4 minutes to minimize failed API calls
// This also ensures that a key will be valid
// even during a lock operation on the mutex
func (c *Client) checkHash() {
	if time.Since(c.hashTime).Minutes() < 4.0 {
		return
	}
	c.genHash()
}

// url gives the final URL that is called based on
// the settings
// locks a read on the mutex to ensure authHash is valid
func (c *Client) url(op string, data ...string) string {
	c.mx.RLock()
	defer c.mx.RUnlock()
	a := strings.Join(data, "")
	return fmt.Sprintf(baseURL, c.endpoint, c.apiID, op, c.authHash, a)
}

// qs is a convenience function to generate a query string
func qs(n, v string) string {
	return fmt.Sprintf("&%s=%s", n, url.QueryEscape(v))
}
