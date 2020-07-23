package logsinput

import (
	"testing"

	"github.com/valyala/fasthttp"
	"github.com/walkline/ipcounter"
)

func BenchmarkHandleLogsMsg(b *testing.B) {
	server := NewServer(ipcounter.NewIPv4BucketIndex(), nil)

	ctx := &fasthttp.RequestCtx{}
	ctx.Request.SetBody([]byte(`{"timestamp":"2020-06-24T15:27:00.123456Z","ip":"83.150.59.250"}`))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		server.handleLogMsg(ctx)
	}
}
