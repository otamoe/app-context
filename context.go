package context

import "context"

var defaultContext = context.Context)

func init() {
	defaultContext = context.Background()
}


func SetContext(ctx context.Context) {
	defaultContext = ctx
}
func GetContext() context.Context {
	return defaultContext
}