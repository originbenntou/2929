package trend

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{}

func NewTrendResolver() *Resolver {
	return &Resolver{}
}

//func getGRPCConn(target string, interceptors ...grpc.UnaryClientInterceptor) *grpc.ClientConn {
//	chain := grpc_middleware.ChainUnaryClient(interceptors...)
//	conn, err := grpc.Dial(target, grpc.WithInsecure(), grpc.WithUnaryInterceptor(chain))
//	if err != nil {
//		log.Fatalf("failed to dial: %s", err)
//	}
//	return conn
//}
//
//func xTraceID(ctx context.Context,
//	method string,
//	req, reply interface{},
//	cc *grpc.ClientConn,
//	invoker grpc.UnaryInvoker,
//	opts ...grpc.CallOption) error {
//	//traceID := support.GetTraceIDFromContext(ctx)
//	//ctx = md.AddTraceIDToContext(ctx, traceID)
//	return invoker(ctx, method, req, reply, cc, opts...)
//}
