package xapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const baseURL = "https://api.x.com/2"

type Client struct {
	httpClient  *http.Client
	bearerToken string
}

func NewClient() (*Client, error) {
	token := strings.TrimSpace(os.Getenv("X_BEARER_TOKEN"))
	if token == "" {
		return nil, fmt.Errorf(
			"X_BEARER_TOKEN environment variable is not set",
		)
	}

	return &Client{
		httpClient: &http.Client{
			Timeout: 15 * time.Second,
		},
		bearerToken: token,
	}, nil
}

func (c *Client) GetUser(
	ctx context.Context,
	username string,
) (*User, error) {
	username = strings.TrimPrefix(
		strings.TrimSpace(username), "@",
	)

	if username == "" || strings.ContainsAny(username, "/?# ") {
		return nil, fmt.Errorf("invalid X username")
	}

	endpoint := baseURL + "/users/by/username/" +
		url.PathEscape(username)

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		endpoint,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	req.Header.Set(
		"Authorization",
		"Bearer "+c.bearerToken,
	)

	query := url.Values{}
	query.Set(
		"user.fields",
		"description,created_at,public_metrics,verified,protected",
	)
	req.URL.RawQuery = query.Encode()

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("X API request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var body struct {
			Title  string `json:"title"`
			Detail string `json:"detail"`
			Type   string `json:"type"`
		}

		if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
			return nil, fmt.Errorf(
				"X API returned HTTP %d",
				resp.StatusCode,
			)
		}

		return nil, fmt.Errorf(
			"X API returned HTTP %d: %s — %s",
			resp.StatusCode,
			body.Title,
			body.Detail,
		)
	}

	var result UserResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("decode X API response: %w", err)
	}

	if result.Data.ID == "" {
		return nil, fmt.Errorf("X API returned no user data")
	}

	return &result.Data, nil
}
