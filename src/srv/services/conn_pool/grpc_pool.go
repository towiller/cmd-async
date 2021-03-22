package connpool

import (
	"context"
	"github.com/0x5010/grpcp"
	//"github.com/silenceper/pool"
	"google.golang.org/grpc"
	//"google.golang.org/grpc/grpclog"
	"log"
	common "srv/proto/common"
	"strconv"
)

type GrpcPool struct {
	Pool        *grpcp.ConnectionTracker
	ServiceHost string
	ServicePort int
	Addr        string
}

func NewGrpcPool(serviceHost string, servicePort int, opts ...grpc.DialOption) *GrpcPool {
	addr := serviceHost + ":" + strconv.Itoa(servicePort)

	pool := grpcp.New(func(addr string) (*grpc.ClientConn, error) {
		return grpc.Dial(
			addr,
			grpc.WithInsecure(),
		)
	})

	/*
		PoolFactory := func() (interface{}, error) {
			var conn *grpc.ClientConn
			var err error

			if len(opts) < 1 {
				conn, err = grpc.Dial(addr, grpc.WithInsecure())
			} else {
				conn, err = grpc.Dial(addr, opts...)
			}

			defer func() {
				if err != nil {
					if cerr := conn.Close(); cerr != nil {
						grpclog.Printf("Failed to close conn to %s: %v", addr, cerr)
					}
					return
				}
				go func() {
					if cerr := conn.Close(); cerr != nil {
						grpclog.Printf("Failed to close conn to %s: %v", addr, cerr)
					}
				}()
			}()

			//log.Println("grpc conn:", conn)
			//client := common.NewCallClient(conn)

			return conn, nil
		}

		PoolClose := func(conn interface{}) error {
			return conn.(*grpc.ClientConn).Close()
		}

		PoolPing := func(conn interface{}) error {
			//return conn.(*grpc.ClientConn)
			return nil
		}

		pConfig := &pool.Config{
			InitialCap: 5,
			MaxCap:     10,
			Factory:    PoolFactory,
			Close:      PoolClose,
			Ping:       PoolPing,
			//连接最大空闲时间，超过该时间的连接 将会关闭，可避免空闲时连接EOF，自动失效的问题
			IdleTimeout: 60 * time.Second,
		}

		pool, _ := pool.NewChannelPool(pConfig)
	*/
	return &GrpcPool{
		Pool:        pool,
		ServiceHost: serviceHost,
		ServicePort: servicePort,
		Addr:        addr,
	}
}

func (g *GrpcPool) Action(ctx context.Context, in *common.CallRequest, opts ...grpc.CallOption) (*common.CallResponse, error) {
	conn, _ := g.Pool.GetConn(g.Addr)
	client := common.NewCallClient(conn)

	return client.Action(ctx, in, opts...)
}

func (g *GrpcPool) Request(ctx context.Context, method string, args common.CallRequest, opts ...grpc.CallOption) (interface{}, error) {
	conn, err := g.Pool.GetConn(g.Addr)
	//fmt.Println(conn)
	if err != nil {
		log.Println("get grpc conn error:", err)
	}
	var reply interface{}
	log.Println("ctx:", ctx, "method:", method, "args:", args)
	ierr := conn.Invoke(ctx, method, args, &reply, opts...)
	//log.Println(ierr)
	return reply, ierr
}
