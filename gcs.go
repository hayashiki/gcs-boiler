package gcsboiler

import (
	"cloud.google.com/go/storage"
	"context"
	"io"
)

type StorageClient interface {
	Read(ctx context.Context, name string) (io.ReadCloser, error)
	Write(ctx context.Context, object string, file io.ReadCloser) error
	Delete(ctx context.Context, object string) error
}

type gcsStorage struct {
	bucket string
}

func (s *gcsStorage) Read(ctx context.Context, name string) (io.ReadCloser, error) {
	client, err := s.new(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	return client.Bucket(s.bucket).Object(name).NewReader(ctx)
}

func New(bucket string) StorageClient {
	return &gcsStorage{
		bucket: bucket,
	}
}

func (s *gcsStorage) new(ctx context.Context) (*storage.Client, error) {
	return storage.NewClient(ctx)
}

func (s *gcsStorage) Write(ctx context.Context, object string, file io.ReadCloser) error {
	defer file.Close()
	client, err := s.new(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	w :=  client.Bucket(s.bucket).Object(object).NewWriter(ctx)
	// Make public
	//w.ACL = []storage.ACLRule{{Entity: storage.AllUsers, Role: storage.RoleReader}}
	defer w.Close()

	if _, err := io.Copy(w, file); err != nil {
		return err
	}
	return nil
}

func (s *gcsStorage) Delete(ctx context.Context, object string) error {
	client, err := s.new(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	return client.Bucket(s.bucket).Object(object).Delete(ctx)
}
