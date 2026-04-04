package classifier

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func Classify(ctx context.Context, u url.URL) (ContentType, error) {

	// 1. First approach: try matching against the url
	if val, exists := testKnownURLS(u); exists {
		return val, nil
	}

	// 2. Fetch the headers only (HEAD url), and find matches
	head_req, err := http.NewRequestWithContext(ctx, "HEAD", u.String(), nil)

	if err != nil {
		return "", err
	}

	resp, err := http.DefaultClient.Do(head_req)
	if err != nil {
		return TypeUnknown, fmt.Errorf("network error: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("bad status: %d", resp.StatusCode)
	}

	if val, exists := testHeaders(resp.Header); exists {
		return val, nil
	}

	// TODO: Continue from here
	return TypeUnknown, nil
}

var knownURLS = map[string]ContentType{
	"youtube.com": TypeVideo,
	"youtu.be":    TypeVideo,
	"vimeo.com":   TypeVideo,
	"medium.com":  TypeArticle,
	"nytimes.com": TypeArticle,
	"twitter.com": TypeTwitterPost,
	"x.com":       TypeTwitterPost,
}

func testKnownURLS(u url.URL) (ContentType, bool) {
	host := strings.ToLower(u.Hostname())
	val, exists := knownURLS[host]
	return val, exists
}

func testHeaders(headers http.Header) (ContentType, bool) {
	ctype := headers.Get("Content-Type")
	switch {
	case strings.Contains(ctype, "application/pdf"):
		return TypeBook, true
	case strings.Contains(ctype, "video/"):
		return TypeVideo, true
	default:
		return TypeUnknown, false
	}
}
