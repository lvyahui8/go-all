package unittest

import (
	"errors"
	"fmt"
	"sync"
	"testing"
	"time"

	. "github.com/bytedance/mockey"
	. "github.com/smartystreets/goconvey/convey"
)

type UserService struct {
	lock sync.Mutex
}

func SayHello() {
	fmt.Println("hello world")
}

var lockFailed = errors.New("get lock failed")

func (s *UserService) Hello() error {
	locked := s.lock.TryLock()
	if locked {
		defer s.lock.Unlock()
	} else {
		return lockFailed
	}
	SayHello()
	return nil
}

func TestConcurrentControl(t *testing.T) {
	PatchConvey("TestConcurrentControl", t, func() {
		Mock(SayHello).To(func() { time.Sleep(3 * time.Second) }).Build()
		s := &UserService{}
		w := &sync.WaitGroup{}
		n := 3
		w.Add(n)
		errCnt := 0
		for i := 0; i < n; i++ {
			go func() {
				defer w.Done()
				if errors.Is(s.Hello(), lockFailed) {
					errCnt++
				}
			}()
		}
		w.Wait()
		So(errCnt, ShouldEqual, n-1)
	})
}
