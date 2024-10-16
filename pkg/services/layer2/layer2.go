package layer2

import (
	"crypto/ed25519"
	"crypto/x509"
	"digital-bank/internal/system/infrastructure/service_credentials"
	"digital-bank/pkg/cache"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type (
	Layer2 struct {
		httpClient  *resty.Client
		endPointURL string
		payload     interface{}
		method      string
		credentials credentials.Layer2Credentials
	}

	TokenResponse struct {
		AccessToken string `json:"access_token"`
	}
)

func NewLayer2(credentials credentials.Layer2Credentials) *Layer2 {
	client := resty.New()

	return &Layer2{
		httpClient:  client,
		endPointURL: "",
		credentials: credentials,
	}
}

func (l *Layer2) getHeader() map[string]string {
	token, _ := l.getToken()

	header := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + token,
	}

	if l.credentials.Signature != "" {
		signature, timestamp := l.genSignature()
		fmt.Println("signature", signature)

		header["x-signature"] = signature
		header["x-timestamp"] = fmt.Sprintf("%d", timestamp)
	}

	return header
}

func (l *Layer2) getToken() (string, error) {
	layer2Token, err := cache.RecoverData("layer2-token")

	if err != nil && err.Error() != "mongo: no documents in result" {
		return "", err
	}

	if layer2Token["token"].(string) == "" {
		log.Println("Token not found in cache")
		return l.Auth()
	}

	return layer2Token["token"].(string), nil
}

func (l *Layer2) Auth() (string, error) {
	log.Println("Authenticate against layer2")

	url := "https://layer2financial.okta.com/oauth2/ausj0isa571aIN3mL696/v1/token?grant_type=client_credentials&scope="
	if os.Getenv("GO_ENV") != "prod" {
		url = "https://auth.layer2financial.com/oauth2/ausbdqlx69rH6OjWd696/v1/token?grant_type=client_credentials&scope="
	}

	scope := "subscriptions:write+subscriptions:read+accounts:read+settlements:read+customers:read+customers:write+accounts:write+withdrawals:read+withdrawals:write+adjustments:read+adjustments:write+exchanges:read+exchanges:write+transfers:read+transfers:write+deposits:read+deposits:write+applications:read+applications:write"

	var header = map[string]string{
		"Content-Type":  "application/x-www-form-urlencoded",
		"Authorization": "Basic " + l.credentials.AuthToken,
	}

	r, err := l.httpClient.R().SetHeaders(header).Post(url + scope)

	if err != nil {
		log.Println("Auth Layer2 Error", err)
		return "", err
	}

	if r.StatusCode() != 200 {
		log.Println("Error in request", r.String())
		return "", err

	}

	var resp TokenResponse
	if err := json.Unmarshal(r.Body(), &resp); err != nil {
		log.Println(err)
		return "", err
	}

	_ = cache.SaveData("layer2-token", map[string]interface{}{"token": resp.AccessToken}, 0)

	return resp.AccessToken, nil
}

func (l *Layer2) genSignature() (string, int64) {
	timestamp := time.Now().Unix()

	payloadBytes, _ := json.Marshal(l.payload)

	var message string
	if l.payload == nil {
		message = fmt.Sprintf("%d%s/api/%s", timestamp, l.method, l.endPointURL)
	} else {
		message = fmt.Sprintf("%d%s/api/%s%s", timestamp, l.method, l.endPointURL, string(payloadBytes))
	}

	fmt.Println("message", message)

	privateKeyPem := hexToPem(l.credentials.Signature, "PRIVATE")
	return signMessage(privateKeyPem, message), timestamp
}

func (l *Layer2) Post() ([]byte, error) {
	l.method = "POST"

	r, err := l.httpClient.R().
		SetHeaders(l.getHeader()).
		SetBody(l.payload).
		Post(l.credentials.URL + l.endPointURL)

	if err != nil {
		return nil, err
	}

	if r.IsError() && r.StatusCode() != http.StatusUnauthorized {
		return nil, errors.New(r.String())
	}

	if r.StatusCode() == http.StatusUnauthorized {
		_, err := l.Auth()
		if err != nil {
			return nil, err
		}

		return l.Post()
	}

	return r.Body(), nil
}

func (l *Layer2) Get() ([]byte, error) {
	l.payload = nil
	l.method = "GET"

	r, err := l.httpClient.R().SetHeaders(l.getHeader()).Get(l.credentials.URL + l.endPointURL)
	if err != nil {
		return nil, err
	}

	if r.IsError() && r.StatusCode() != http.StatusUnauthorized {
		return nil, errors.New(r.String())
	}

	if r.StatusCode() == http.StatusUnauthorized {
		_, err := l.Auth()
		if err != nil {
			return nil, err
		}

		return l.Get()
	}

	return r.Body(), nil
}

func (l *Layer2) SendDocument(file []byte, docID string) error {
	token, _ := l.getToken()

	header := map[string]string{
		"Content-Type":  "application/octet-stream",
		"Authorization": "Bearer " + token,
	}

	r, err := l.httpClient.R().SetHeaders(header).SetBody(file).Post(l.credentials.URL + "v1/documents/" + docID)

	if err != nil {
		return err
	}

	if r.IsError() && r.StatusCode() != http.StatusUnauthorized {
		return errors.New(r.String())
	}

	if r.StatusCode() == http.StatusUnauthorized {
		_, err := l.Auth()
		if err != nil {
			return err
		}

		return l.SendDocument(file, docID)
	}

	return nil
}

func hexToPem(hexString string, keyType string) string {
	hexBytes, _ := hex.DecodeString(hexString)
	base64String := base64.StdEncoding.EncodeToString(hexBytes)

	// Split the base64 string every 64 characters
	splitBase64 := splitSubN(base64String, 64)

	header := fmt.Sprintf("-----BEGIN %s KEY-----\n", keyType)
	footer := fmt.Sprintf("\n-----END %s KEY-----", keyType)

	return header + strings.Join(splitBase64, "\n") + footer
}

// splitSubN splits a string every n characters
func splitSubN(s string, n int) []string {
	split := make([]string, 0, len(s)/n+1)
	for len(s) > n {
		split = append(split, s[:n])
		s = s[n:]
	}
	split = append(split, s)
	return split
}

func signMessage(privateKeyPem string, message string) string {
	block, _ := pem.Decode([]byte(privateKeyPem))
	if block == nil {
		log.Fatal("failed to decode PEM block containing private key")
	}

	priv, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		log.Fatal("failed to parse private key: " + err.Error())
	}

	ed25519Priv, ok := priv.(ed25519.PrivateKey)
	if !ok {
		log.Fatal("not Ed25519 private key")
	}

	messageBytes := []byte(message)

	signature := ed25519.Sign(ed25519Priv, messageBytes)

	return hex.EncodeToString(signature)
}
