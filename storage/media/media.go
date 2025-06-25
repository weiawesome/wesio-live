package media

import (
	"context"
	"io"
	"time"
)

type FileType string

const (
	Image FileType = "image"
	Video FileType = "video"
)

type UploadOptions struct {
	ContentType string
	Metadata    map[string]string
}

type Media interface {
	Upload(ctx context.Context, fileType FileType, filename string, data io.Reader, opts *UploadOptions) (string, error)

	Download(ctx context.Context, fileType FileType, filename string) (io.ReadCloser, error)

	GetURL(ctx context.Context, fileType FileType, filename string, expiration time.Duration) (string, error)

	GetCDNURL(ctx context.Context, fileType FileType, filename string, expiration time.Duration) (string, error)

	Delete(ctx context.Context, fileType FileType, filename string) error
}
