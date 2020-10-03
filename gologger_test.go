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

func TestGoLogger_Warn(t *testing.T) {
	logger := NewGoLogger(Config{LogLevel: 3})
	logger.SetScope("TestGoLogger_Warn")

	s := make([]map[string]interface{}, 0)

	s = append(s, H{"s": false})
	logger.Warn(s)
	logger.Warn("testing info log")
	logger.Warn(1)
}

func TestGoLogger_Warn2(t *testing.T) {
	logger := NewGoLogger(Config{LogLevel: 3}).SetScope("TestGoLogger_Warn2")

	s := make([]map[string]interface{}, 0)

	s = append(s, H{"s": false})
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		logger.Warn(s)
		wg.Done()
	}()

	go func() {
		logger.Warn(1)
		wg.Done()
	}()

	go func() {
		logger.Warn("testing info log")
		wg.Done()
	}()

	wg.Wait()
}

func TestGoLogger_Error(t *testing.T) {
	logger := NewGoLogger(Config{LogLevel: 3})
	logger.SetScope("TestGoLogger_Error")

	s := make([]map[string]interface{}, 0)

	s = append(s, H{"s": false})
	logger.Error(s)
	logger.Error("testing info log")
	logger.Error(1)
}

func TestGoLogger_Error2(t *testing.T) {
	logger := NewGoLogger(Config{LogLevel: 3}).SetScope("TestGoLogger_Warn2")

	s := make([]map[string]interface{}, 0)

	s = append(s, H{"s": false})
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		logger.Error(s)
		wg.Done()
	}()

	go func() {
		logger.Error(1)
		wg.Done()
	}()

	go func() {
		logger.Error("testing info log")
		wg.Done()
	}()

	wg.Wait()
}

func BenchmarkGoLogger_Info(b *testing.B) {
	logger := NewGoLogger(Config{LogLevel: 1}).SetScope("TestGoLogger_Info")
	s := make([]map[string]interface{}, 0)

	s = append(s, H{"s": false})
	for i:=0; i< b.N; i++ {
		logger.Info(s)
		logger.Info("testing info log")
		logger.Info(1)
	}
}
