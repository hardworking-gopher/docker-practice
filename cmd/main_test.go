package main

import (
	"github.com/go-redis/redis"
	"github.com/hardworking-gopher/docker-kubernetes/mocks"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func Test_VisitsEndpoint(t *testing.T) {

	t.Run("first successful visit", func(t *testing.T) {
		rdbMock := mocks.NewMockRedisClientInterface(t)

		rdb = rdbMock

		rdbMock.On("Get", redisVisitsKey).Return(redis.NewStringResult("", redis.Nil))
		rdbMock.On("Set", redisVisitsKey, 1, time.Duration(0)).Return(redis.NewStatusCmd("1", nil))

		var (
			w = httptest.NewRecorder()
			r = httptest.NewRequest(http.MethodGet, "/visits", nil)
		)

		visitsEndpoint(w, r)

		rdbMock.AssertCalled(t, "Get", redisVisitsKey)
		rdbMock.AssertCalled(t, "Set", redisVisitsKey, 1, time.Duration(0))

		assert.Equal(t, w.Code, 200)
	})

	t.Run("subsequent successful visit", func(t *testing.T) {
		rdbMock := mocks.NewMockRedisClientInterface(t)

		rdb = rdbMock

		rdbMock.On("Get", redisVisitsKey).Return(redis.NewStringResult("1", nil))
		rdbMock.On("Set", redisVisitsKey, 2, time.Duration(0)).Return(redis.NewStatusCmd("2", nil))

		var (
			w = httptest.NewRecorder()
			r = httptest.NewRequest(http.MethodGet, "/visits", nil)
		)

		visitsEndpoint(w, r)

		rdbMock.AssertCalled(t, "Get", redisVisitsKey)
		rdbMock.AssertCalled(t, "Set", redisVisitsKey, 2, time.Duration(0))

		assert.Equal(t, w.Code, 200)
	})

}
