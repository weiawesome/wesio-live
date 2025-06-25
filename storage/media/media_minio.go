package media

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/url"
	"strconv"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinIOMedia struct {
	minioClient *minio.Client
	cdnDomain   string
	cdnSignKey  string // CDN signing key for generating signed URLs
	domain      string
	secure      bool // true for https, false for http
}

func CreateMinIOMedia(ctx context.Context, cdnDomain string, cdnSignKey string, domain string, accessKey string, secretKey string, secure bool) *MinIOMedia {
	minioClient, err := minio.New(domain, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: secure,
	})
	if err != nil {
		panic(err)
	}

	return &MinIOMedia{
		minioClient: minioClient,
		cdnDomain:   cdnDomain,
		cdnSignKey:  cdnSignKey,
		domain:      domain,
		secure:      secure,
	}
}

// getBucketName returns the bucket name based on file type
func (m *MinIOMedia) getBucketName(fileType FileType) string {
	switch fileType {
	case Image:
		return "images"
	case Video:
		return "videos"
	default:
		return "files"
	}
}

// ensureBucketExists creates bucket if it doesn't exist and sets public read policy
func (m *MinIOMedia) ensureBucketExists(ctx context.Context, bucketName string) error {
	exists, err := m.minioClient.BucketExists(ctx, bucketName)
	if err != nil {
		return fmt.Errorf("failed to check bucket existence: %w", err)
	}

	if !exists {
		return fmt.Errorf("bucket %s does not exist", bucketName)
	}

	// if !exists {
	// 	err = m.minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
	// 	if err != nil {
	// 		return fmt.Errorf("failed to create bucket: %w", err)
	// 	}
	// }

	return nil
}

func (m *MinIOMedia) Upload(ctx context.Context, fileType FileType, filename string, data io.Reader, opts *UploadOptions) (string, error) {
	bucketName := m.getBucketName(fileType)

	// Ensure bucket exists
	if err := m.ensureBucketExists(ctx, bucketName); err != nil {
		return "", err
	}

	// Set up put object options
	putOpts := minio.PutObjectOptions{}

	if opts != nil {
		if opts.ContentType != "" {
			putOpts.ContentType = opts.ContentType
		}
		if opts.Metadata != nil {
			putOpts.UserMetadata = opts.Metadata
		}
	}

	// Upload the object
	_, err := m.minioClient.PutObject(ctx, bucketName, filename, data, -1, putOpts)
	if err != nil {
		return "", fmt.Errorf("failed to upload file: %w", err)
	}

	// Return the file path/key
	return fmt.Sprintf("%s/%s", bucketName, filename), nil
}

func (m *MinIOMedia) Download(ctx context.Context, fileType FileType, filename string) (io.ReadCloser, error) {
	bucketName := m.getBucketName(fileType)

	object, err := m.minioClient.GetObject(ctx, bucketName, filename, minio.GetObjectOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to get object: %w", err)
	}

	return object, nil
}

func (m *MinIOMedia) GetURL(ctx context.Context, fileType FileType, filename string, expiration time.Duration) (string, error) {
	bucketName := m.getBucketName(fileType)

	// Always generate a presigned URL with the specified expiration
	presignedURL, err := m.minioClient.PresignedGetObject(ctx, bucketName, filename, expiration, nil)
	if err != nil {
		return "", fmt.Errorf("failed to generate presigned URL: %w", err)
	}
	return presignedURL.String(), nil
}

func (m *MinIOMedia) GetCDNURL(ctx context.Context, fileType FileType, filename string, expiration time.Duration) (string, error) {
	if m.cdnDomain == "" {
		return "", fmt.Errorf("CDN URL not configured")
	}

	if m.cdnSignKey == "" {
		return "", fmt.Errorf("CDN signing key not configured")
	}

	bucketName := m.getBucketName(fileType)

	// Parse CDN URL and append the file path
	cdnURL, err := url.Parse(m.cdnDomain)
	if err != nil {
		return "", fmt.Errorf("invalid CDN URL: %w", err)
	}

	// Construct the full CDN URL path
	cdnURL.Path = fmt.Sprintf("/%s/%s", bucketName, filename)

	// Generate signed CDN URL
	signedURL, err := m.generateCDNSignedURL(cdnURL.String(), expiration)
	if err != nil {
		return "", fmt.Errorf("failed to generate CDN signed URL: %w", err)
	}

	return signedURL, nil
}

// generateCDNSignedURL generates a signed URL for CDN access
// This is a generic implementation - you may need to adapt it for specific CDN providers
func (m *MinIOMedia) generateCDNSignedURL(baseURL string, expiration time.Duration) (string, error) {
	// Calculate expiration timestamp
	expirationTime := time.Now().Add(expiration).Unix()

	// Parse the URL to add parameters
	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		return "", err
	}

	// Create the string to sign (URL path + expiration)
	stringToSign := parsedURL.Path + strconv.FormatInt(expirationTime, 10)

	// Generate HMAC signature
	h := hmac.New(sha256.New, []byte(m.cdnSignKey))
	h.Write([]byte(stringToSign))
	signature := hex.EncodeToString(h.Sum(nil))

	// Add signature and expiration as query parameters
	query := parsedURL.Query()
	query.Set("expires", strconv.FormatInt(expirationTime, 10))
	query.Set("signature", signature)
	parsedURL.RawQuery = query.Encode()

	return parsedURL.String(), nil
}

func (m *MinIOMedia) Delete(ctx context.Context, fileType FileType, filename string) error {
	bucketName := m.getBucketName(fileType)

	err := m.minioClient.RemoveObject(ctx, bucketName, filename, minio.RemoveObjectOptions{})
	if err != nil {
		return fmt.Errorf("failed to delete object: %w", err)
	}

	return nil
}
