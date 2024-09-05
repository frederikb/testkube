package artifacts

import (
	"context"
	"io"
	"path/filepath"
	"sync"
	"time"

	"github.com/kubeshop/testkube/cmd/testworkflow-toolkit/env"
)

type InternalArtifactStorage interface {
	FullPath(artifactPath string) string
	SaveStream(artifactPath string, stream io.Reader) error
	Wait() error
}

type withLength interface {
	Len() int
}

type internalArtifactStorage struct {
	prefix   string
	uploader Uploader
	startMu  sync.Mutex
	started  bool
}

func newArtifactUploader() Uploader {
	if env.CloudEnabled() {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		client, _ := env.Cloud(ctx)
		return NewCloudUploader(client, WithParallelismCloud(30), CloudDetectMimetype)
	}
	return NewDirectUploader(WithParallelism(30), DirectDetectMimetype)
}

func InternalStorage() InternalArtifactStorage {
	return &internalArtifactStorage{
		prefix:   filepath.Join(".testkube", env.Ref()),
		uploader: newArtifactUploader(),
	}
}

func (s *internalArtifactStorage) start() error {
	s.startMu.Lock()
	defer s.startMu.Unlock()
	if s.started {
		return nil
	}
	s.started = true
	return s.uploader.Start()
}

func (s *internalArtifactStorage) FullPath(filePath string) string {
	return filepath.Join(s.prefix, filePath)
}

func (s *internalArtifactStorage) SaveStream(artifactPath string, stream io.Reader) error {
	err := s.start()
	if err != nil {
		return err
	}

	size := -1
	if streamL, ok := stream.(withLength); ok {
		size = streamL.Len()
	}
	err = s.uploader.Add(filepath.Join(s.prefix, artifactPath), stream, int64(size))
	if err != nil {
		return err
	}
	return s.uploader.End()
}

func (s *internalArtifactStorage) Wait() error {
	s.startMu.Lock()
	defer s.startMu.Unlock()
	if !s.started {
		return nil
	}
	return s.uploader.End()
}
