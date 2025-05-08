package middleware

type ContextKey struct {
	Name string
}

func (k *ContextKey) String() string {
	return "go-blog context value " + k.Name
}

type MiddlewareManager struct{}
