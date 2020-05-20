package annual1

import (
	"fmt"
	"github.com/kataras/iris/httptest"
	"sync"
	"testing"
)

func TestMVC(t *testing.T) {
	e := httptest.New(t, newApp())
	var wg sync.WaitGroup
	e.GET("/").Expect().Status(httptest.StatusOK).Body().Equal("当前共参与抽奖的用户数： 0\n")
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func(i int) {
			defer wg.Done()
			e.POST("/import").WithFormField("users", fmt.Sprintf("test_u%d", i)).Expect().
				Status(httptest.StatusOK)
		}(i)
	}
	wg.Wait()

	e.GET("/").Expect().Status(httptest.StatusOK).Body().Equal("当前共参与抽奖的用户数： 100\n")
	e.GET("/lucky").Expect().Status(httptest.StatusOK)
	e.GET("/").Expect().Status(httptest.StatusOK).Body().Equal("当前共参与抽奖的用户数： 99\n")
}
