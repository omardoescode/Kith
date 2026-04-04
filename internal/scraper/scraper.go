package scraper

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/JohannesKaufmann/html-to-markdown/v2"
	"github.com/microcosm-cc/bluemonday"
)

func Scrape(ctx context.Context, url *url.URL) (string, error) {
	client := &http.Client{
		Timeout:       time.Millisecond * 5000,
		CheckRedirect: checkRedirect,
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url.String(), nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("User-Agent", "KithBot/1.0 (+https://kith.io)")

	resp, err := client.Do(req)

	if err != nil {
		return "", fmt.Errorf("network error: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("bad status: %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	limitReader := io.LimitReader(resp.Body, 2*1024*1024)
	body, err := io.ReadAll(limitReader)
	if err != nil {
		return "", fmt.Errorf("read error: %w", err)
	}

	p := bluemonday.UGCPolicy()
	cleanHTML := p.SanitizeBytes(body)

	md, err := htmltomarkdown.ConvertString(string(cleanHTML))
	if err != nil {
		return "", fmt.Errorf("markdown conversion failed: %w", err)
	}

	return md, nil
}

func checkRedirect(req *http.Request, via []*http.Request) error {
	if len(via) > 5 {
		return fmt.Errorf("stopped after 5 redirects")
	}

	host := req.URL.Hostname()

	ips, err := net.LookupIP(host)

	if err != nil {
		return fmt.Errorf("redirect DNS failed: %w", err)
	}

	for _, ip := range ips {
		if isRestrictedIP(ip) {
			return fmt.Errorf("redirect URL points to unsafe IP: %s", ip)
		}
	}

	return nil
}
