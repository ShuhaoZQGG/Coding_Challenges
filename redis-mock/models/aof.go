package models

import (
	"bufio"
	"os"
	"sync"
	"time"

	"github.com/redis-mock/configs"
)

type AOF struct {
	file *os.File
	rd   *bufio.Reader
	mu   sync.Mutex
}

func NewAOF() (*AOF, error) {
	config, _ := configs.GetConfig("config.yaml")
	aofLocation := config.Aof.Location
	syncPolicy := config.Aof.Appendfsync
	file, err := os.OpenFile(aofLocation, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	aof := &AOF{
		file: file,
		rd:   bufio.NewReader(file),
	}

	if syncPolicy == "everysec" {
		go func() {
			for {
				aof.mu.Lock()
				aof.file.Sync()
				aof.mu.Unlock()
				time.Sleep(time.Second)
			}
		}()
	}

	if syncPolicy == "always" {
		go func() {
			for {
				aof.mu.Lock()
				aof.file.Sync()
				aof.mu.Unlock()
			}
		}()
	}

	return aof, nil
}

func (aof *AOF) Close() error {
	aof.mu.Lock()
	defer aof.mu.Unlock()
	return aof.file.Close()
}

func (aof *AOF) Write(message string) error {
	aof.mu.Lock()
	defer aof.mu.Unlock()

	_, err := aof.file.Write([]byte(message))
	if err != nil {
		return err
	}
	return nil
}
