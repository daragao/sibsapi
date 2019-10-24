package client

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// Client struct that has the sibs functions
// TODO: make this struct private (with a public interface and factory)
type Client struct {
	Client     *http.Client
	ClientID   string
	Host       string
	PathPrefix string
}

func (c *Client) newRequest(method, relativePath string, payload io.Reader) (*http.Request, error) {
	return http.NewRequest(method, c.Host+c.PathPrefix+relativePath, payload)
}

func generateUUID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid
}

func (c *Client) executeRequest(req *http.Request) ([]byte, error) {
	//req.Header.Add("accept", "application/json")
	req.Header.Add("signature", "REPLACE_THIS_VALUE")
	req.Header.Add("tpp-certificate", "REPLACE_THIS_VALUE")
	req.Header.Add("tpp-request-id", generateUUID())
	req.Header.Add("tpp-transaction-id", generateUUID())
	req.Header.Add("x-ibm-client-id", c.ClientID)

	//client := http.DefaultClient
	resp, err := c.Client.Do(req)
	if err != nil {
		log.Printf("Failed to do http request: %s", err)
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Failed to parse response: %s", err)
		return nil, err
	}

	if resp.StatusCode != 200 {

		//log.Printf("Bad response code from %s: %d\n%s", req.URL, resp.StatusCode, body)
		errorStr := fmt.Sprintf("Bad response code from %s: %d\nBody\n\t%s\nHeader\n", req.URL, resp.StatusCode, body)
		for name, headers := range resp.Header {
			for _, h := range headers {
				//fmt.Printf("\t%v:\t%v\n", name, h)
				errorStr += fmt.Sprintf("\t%v:\t%v\n", name, h)
			}
		}
		return body, errors.New(errorStr)
	}

	return body, nil
}

// ListAvailableASPSP method
func (c *Client) ListAvailableASPSP() ([]byte, error) {
	req, err := c.newRequest("GET", "v1/available-aspsp", nil)
	if err != nil {
		return nil, err
	}

	return c.executeRequest(req)
}

// ListAccounts method
func (c *Client) ListAccounts(aspspCde string) ([]byte, error) {
	req, err := c.newRequest("GET", aspspCde+"/v1-0-2/accounts", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("consent-id", "REPLACE_THIS_VALUE")
	req.Header.Add("date", "REPLACE_THIS_VALUE")

	return c.executeRequest(req)
}

// GetAccount method
func (c *Client) GetAccount(aspspCde, accountID string) ([]byte, error) {
	req, err := c.newRequest("GET", aspspCde+"/v1-0-2/accounts/"+accountID, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("consent-id", "REPLACE_THIS_VALUE")
	req.Header.Add("date", "REPLACE_THIS_VALUE")

	return c.executeRequest(req)
}

// GetBalances method
func (c *Client) GetBalances(aspspCde, accountID string) ([]byte, error) {
	req, err := c.newRequest("GET", aspspCde+"/v1-0-2/accounts/"+accountID+"/balances", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("consent-id", "REPLACE_THIS_VALUE")
	req.Header.Add("date", "REPLACE_THIS_VALUE")

	return c.executeRequest(req)
}

// GetTransactions method
func (c *Client) GetTransactions(aspspCde, accountID string) ([]byte, error) {
	req, err := c.newRequest("GET", aspspCde+"/v1-0-2/accounts/"+accountID+"/transactions", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("consent-id", "REPLACE_THIS_VALUE")
	req.Header.Add("date", "REPLACE_THIS_VALUE")

	return c.executeRequest(req)
}

// GetConsent method
func (c *Client) GetConsent(aspspCde, consentID string) ([]byte, error) {
	req, err := c.newRequest("GET", aspspCde+"/v1-0-2/consents/"+consentID, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("psu-id", "REPLACE_THIS_VALUE")
	req.Header.Add("psu-id-type", "REPLACE_THIS_VALUE")
	req.Header.Add("psu-corporate-id", "REPLACE_THIS_VALUE")
	req.Header.Add("psu-corporate-id-type", "REPLACE_THIS_VALUE")

	req.Header.Add("date", "REPLACE_THIS_VALUE")

	return c.executeRequest(req)
}

// NewConsent method
func (c *Client) NewConsent(aspspCde string, payloadStruct ConsentPayload) ([]byte, error) {
	// Prepare Query Parameters
	// params := url.Values{}
	// params.Add("tppRedirectPreferred", "http://localhost")
	// params.Add("withBalance", 0)

	payload, err := json.Marshal(payloadStruct)
	if err != nil {
		return nil, err
	}

	req, err := c.newRequest("POST", aspspCde+"/v1-0-2/consents", bytes.NewReader([]byte(payload)))
	if err != nil {
		return nil, err
	}

	req.Header.Add("psu-id", "REPLACE_THIS_VALUE")
	req.Header.Add("psu-id-type", "REPLACE_THIS_VALUE")
	req.Header.Add("psu-corporate-id", "REPLACE_THIS_VALUE")
	req.Header.Add("psu-corporate-id-type", "REPLACE_THIS_VALUE")
	req.Header.Add("tpp-redirect-uri", "REPLACE_THIS_VALUE")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("date", "REPLACE_THIS_VALUE")
	req.Header.Add("digest", "REPLACE_THIS_VALUE") // FIX make digest

	return c.executeRequest(req)
}
