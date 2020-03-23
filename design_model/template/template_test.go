package template

import "testing"

func TestTemplate(t *testing.T) {
	ftpDownloader := NewFTPDownloader()
	httpDownloader := NewHTTPDownloader()

	ftpDownloader.Download("github.com")
	httpDownloader.Download("google.com")
}
