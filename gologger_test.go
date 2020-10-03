package gologger

import (
	"sync"
	"testing"
)

type H map[string]interface{}

func TestGoLogger_Info(t *testing.T) {
	logger := NewGoLogger(Config{LogLevel: 1}).SetScope("TestGoLogger_Info")
	s := make([]map[string]interface{}, 0)

	s = append(s, H{"s": false})
	logger.Info(s)
	logger.Info("testing info log")
	logger.Info(1)
}

func TestGoLogger_Info2(t *testing.T) {
	logger := NewGoLogger(Config{LogLevel: 1}).SetScope("TestGoLogger_Info")

	s := make([]map[string]interface{}, 0)

	s = append(s, H{"s": false})
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		logger.Info(s)
		wg.Done()
	}()

	go func() {
		logger.Info(1)
		wg.Done()
	}()

	go func() {
		logger.Info("testing info log")
		wg.Done()
	}()

	wg.Wait()

}

func TestGoLogger_Debug(t *testing.T) {
	logger := NewGoLogger(Config{LogLevel: 3})
	logger.SetScope("TESTING")

	s := make([]map[string]interface{}, 0)

	s = append(s, H{"s": false})
	logger.Debug(s)
	logger.Debug("testing info log")
	logger.Debug(1)
}

func TestGoLogger_Debug2(t *testing.T) {
	logger := NewGoLogger(Config{LogLevel: 3}).SetScope("TestGoLogger")

	s := make([]map[string]interface{}, 0)

	s = append(s, H{"s": false})
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		logger.Debug(s)
		wg.Done()
	}()

	go func() {
		logger.Debug(1)
		wg.Done()
	}()

	go func() {
		logger.Debug("testing info log")
		wg.Done()
	}()

	wg.Wait()
}
