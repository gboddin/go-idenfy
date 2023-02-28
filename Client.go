package idenfy

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Client struct {
	accessKey       string
	secretKey       string
	endpoint        string
	httpClient      *http.Client
	callbackSignKey []byte
}

// NewClient Creates a new Idenfy API client
func NewClient(options ...ClientOption) (*Client, error) {
	client := &Client{
		httpClient: http.DefaultClient,
		endpoint:   "https://ivs.idenfy.com",
	}
	for _, option := range options {
		err := option(client)
		if err != nil {
			return nil, err
		}
	}
	return client, nil
}

// CreateIdentityVerificationSession Creates a new identity verification session and returns the response
func (c *Client) CreateIdentityVerificationSession(ctx context.Context, request TokenRequest) (*TokenResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}
	payload, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.apiEndpoint("/api/v2/token"), bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}
	c.setRequestParameters(req)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, ErrServerError
	}
	var response TokenResponse
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// DecodeReaderIdentityCallback Decodes an identity callback response from an io.Reader and verify its signature
func (c *Client) DecodeReaderIdentityCallback(reader io.Reader, sigHeader string) (*IdentityCallbackResp, error) {
	if len(c.callbackSignKey) < 1 {
		return nil, errors.New("callback was received but no signature key was provided")
	}
	sig, err := hex.DecodeString(sigHeader)
	if err != nil {
		return nil, err
	}
	mac := hmac.New(sha256.New, c.callbackSignKey)
	tee := io.TeeReader(reader, mac)
	decoder := json.NewDecoder(tee)
	var callbackResp IdentityCallbackResp
	err = decoder.Decode(&callbackResp)
	if err != nil {
		return nil, err
	}
	if !hmac.Equal(sig, mac.Sum(nil)) {
		return &callbackResp, errors.New("signature verification failed")
	}
	return &callbackResp, nil
}

// DecodeHttpRequestIdentityCallback Decodes an identity callback response from an *http.Request and verify its signature
func (c *Client) DecodeHttpRequestIdentityCallback(request *http.Request) (*IdentityCallbackResp, error) {
	if request.Method != http.MethodPost {
		return nil, errors.New("expected POST request")
	}
	sigHeader := request.Header.Get("Idenfy-Signature")
	if len(sigHeader) < 1 {
		return nil, errors.New("no signature provided")
	}
	return c.DecodeReaderIdentityCallback(request.Body, sigHeader)
}

func (c *Client) apiEndpoint(path string) string {
	return fmt.Sprintf("%s%s", strings.TrimPrefix(c.endpoint, "/"), path)
}

func (c *Client) setRequestParameters(req *http.Request) {
	req.SetBasicAuth(c.accessKey, c.secretKey)
	req.Header.Set("User-Agent", "go-idenfy/1.0 (+https://github.com/gboddin/go-idenfy)")
	req.Header.Set("Content-Type", "application/json")
}

type ClientOption func(client *Client) error

// WithCustomEndpoint Option is used to specify an alternative API endpoint
func WithCustomEndpoint(endpoint string) ClientOption {
	return func(client *Client) error {
		client.endpoint = endpoint
		return nil
	}
}

// WithCallbackSignatureKey Option is used to specify a callback signature key to verify requests signatures
func WithCallbackSignatureKey(key string) ClientOption {
	return func(client *Client) (err error) {
		client.callbackSignKey, err = hex.DecodeString(key)
		return
	}
}

// WithCustomHttpClient Option is used to specify a custom *http.Client to make requests with
func WithCustomHttpClient(httpClient *http.Client) ClientOption {
	return func(client *Client) error {
		client.httpClient = httpClient
		return nil
	}
}

// WithApiCredentials Option is used to specify Idenfy API credentials
func WithApiCredentials(accessKey, secretKey string) ClientOption {
	return func(client *Client) error {
		client.accessKey = accessKey
		client.secretKey = secretKey
		return nil
	}
}

var ErrServerError = errors.New("received server error")
