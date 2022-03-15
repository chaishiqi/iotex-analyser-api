// Code generated by proroc-gen-graphql, DO NOT EDIT.
package api

import (
	"context"

	"github.com/graphql-go/graphql"
	pagination "github.com/iotexproject/iotex-analyser-api/api/pagination"
	"github.com/pkg/errors"
	"github.com/ysugimoto/grpc-graphql-gateway/runtime"
	"google.golang.org/grpc"
)

var (
	gql__type_XrcInfo          *graphql.Object      // message XrcInfo in api_action.proto
	gql__type_EvmTransferInfo  *graphql.Object      // message EvmTransferInfo in api_action.proto
	gql__type_ActionResponse   *graphql.Object      // message ActionResponse in api_action.proto
	gql__type_ActionRequest    *graphql.Object      // message ActionRequest in api_action.proto
	gql__type_ActionInfo       *graphql.Object      // message ActionInfo in api_action.proto
	gql__input_XrcInfo         *graphql.InputObject // message XrcInfo in api_action.proto
	gql__input_EvmTransferInfo *graphql.InputObject // message EvmTransferInfo in api_action.proto
	gql__input_ActionResponse  *graphql.InputObject // message ActionResponse in api_action.proto
	gql__input_ActionRequest   *graphql.InputObject // message ActionRequest in api_action.proto
	gql__input_ActionInfo      *graphql.InputObject // message ActionInfo in api_action.proto
)

func Gql__type_XrcInfo() *graphql.Object {
	if gql__type_XrcInfo == nil {
		gql__type_XrcInfo = graphql.NewObject(graphql.ObjectConfig{
			Name: "Api_Type_XrcInfo",
			Fields: graphql.Fields{
				"actHash": &graphql.Field{
					Type: graphql.String,
				},
				"from": &graphql.Field{
					Type: graphql.String,
				},
				"to": &graphql.Field{
					Type: graphql.String,
				},
				"quantity": &graphql.Field{
					Type: graphql.String,
				},
				"blkHeight": &graphql.Field{
					Type: graphql.Int,
				},
				"timestamp": &graphql.Field{
					Type: graphql.Int,
				},
				"contract": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__type_XrcInfo
}

func Gql__type_EvmTransferInfo() *graphql.Object {
	if gql__type_EvmTransferInfo == nil {
		gql__type_EvmTransferInfo = graphql.NewObject(graphql.ObjectConfig{
			Name: "Api_Type_EvmTransferInfo",
			Fields: graphql.Fields{
				"actHash": &graphql.Field{
					Type: graphql.String,
				},
				"blkHash": &graphql.Field{
					Type: graphql.String,
				},
				"from": &graphql.Field{
					Type: graphql.String,
				},
				"to": &graphql.Field{
					Type: graphql.String,
				},
				"quantity": &graphql.Field{
					Type: graphql.String,
				},
				"blkHeight": &graphql.Field{
					Type: graphql.Int,
				},
				"timestamp": &graphql.Field{
					Type: graphql.Int,
				},
			},
		})
	}
	return gql__type_EvmTransferInfo
}

func Gql__type_ActionResponse() *graphql.Object {
	if gql__type_ActionResponse == nil {
		gql__type_ActionResponse = graphql.NewObject(graphql.ObjectConfig{
			Name: "Api_Type_ActionResponse",
			Fields: graphql.Fields{
				"exist": &graphql.Field{
					Type: graphql.Boolean,
				},
				"count": &graphql.Field{
					Type: graphql.Int,
				},
				"actionList": &graphql.Field{
					Type: graphql.NewList(Gql__type_ActionInfo()),
				},
				"evmTransferList": &graphql.Field{
					Type: graphql.NewList(Gql__type_EvmTransferInfo()),
				},
				"xrcList": &graphql.Field{
					Type: graphql.NewList(Gql__type_XrcInfo()),
				},
			},
		})
	}
	return gql__type_ActionResponse
}

func Gql__type_ActionRequest() *graphql.Object {
	if gql__type_ActionRequest == nil {
		gql__type_ActionRequest = graphql.NewObject(graphql.ObjectConfig{
			Name: "Api_Type_ActionRequest",
			Fields: graphql.Fields{
				"address": &graphql.Field{
					Type: graphql.String,
				},
				"actHash": &graphql.Field{
					Type: graphql.String,
				},
				"pagination": &graphql.Field{
					Type: pagination.Gql__type_Pagination(),
				},
			},
		})
	}
	return gql__type_ActionRequest
}

func Gql__type_ActionInfo() *graphql.Object {
	if gql__type_ActionInfo == nil {
		gql__type_ActionInfo = graphql.NewObject(graphql.ObjectConfig{
			Name: "Api_Type_ActionInfo",
			Fields: graphql.Fields{
				"actHash": &graphql.Field{
					Type: graphql.String,
				},
				"blkHash": &graphql.Field{
					Type: graphql.String,
				},
				"actType": &graphql.Field{
					Type: graphql.String,
				},
				"sender": &graphql.Field{
					Type: graphql.String,
				},
				"recipient": &graphql.Field{
					Type: graphql.String,
				},
				"amount": &graphql.Field{
					Type: graphql.String,
				},
				"timestamp": &graphql.Field{
					Type: graphql.Int,
				},
				"gasFee": &graphql.Field{
					Type: graphql.String,
				},
				"blkHeight": &graphql.Field{
					Type: graphql.Int,
				},
			},
		})
	}
	return gql__type_ActionInfo
}

func Gql__input_XrcInfo() *graphql.InputObject {
	if gql__input_XrcInfo == nil {
		gql__input_XrcInfo = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Api_Input_XrcInfo",
			Fields: graphql.InputObjectConfigFieldMap{
				"actHash": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"from": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"to": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"quantity": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"blkHeight": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				"timestamp": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				"contract": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__input_XrcInfo
}

func Gql__input_EvmTransferInfo() *graphql.InputObject {
	if gql__input_EvmTransferInfo == nil {
		gql__input_EvmTransferInfo = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Api_Input_EvmTransferInfo",
			Fields: graphql.InputObjectConfigFieldMap{
				"actHash": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"blkHash": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"from": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"to": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"quantity": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"blkHeight": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				"timestamp": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
			},
		})
	}
	return gql__input_EvmTransferInfo
}

func Gql__input_ActionResponse() *graphql.InputObject {
	if gql__input_ActionResponse == nil {
		gql__input_ActionResponse = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Api_Input_ActionResponse",
			Fields: graphql.InputObjectConfigFieldMap{
				"exist": &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
				"count": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				"actionList": &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(Gql__input_ActionInfo()),
				},
				"evmTransferList": &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(Gql__input_EvmTransferInfo()),
				},
				"xrcList": &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(Gql__input_XrcInfo()),
				},
			},
		})
	}
	return gql__input_ActionResponse
}

func Gql__input_ActionRequest() *graphql.InputObject {
	if gql__input_ActionRequest == nil {
		gql__input_ActionRequest = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Api_Input_ActionRequest",
			Fields: graphql.InputObjectConfigFieldMap{
				"address": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"actHash": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"pagination": &graphql.InputObjectFieldConfig{
					Type: pagination.Gql__input_Pagination(),
				},
			},
		})
	}
	return gql__input_ActionRequest
}

func Gql__input_ActionInfo() *graphql.InputObject {
	if gql__input_ActionInfo == nil {
		gql__input_ActionInfo = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Api_Input_ActionInfo",
			Fields: graphql.InputObjectConfigFieldMap{
				"actHash": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"blkHash": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"actType": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"sender": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"recipient": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"amount": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"timestamp": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				"gasFee": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"blkHeight": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
			},
		})
	}
	return gql__input_ActionInfo
}

// graphql__resolver_ActionService is a struct for making query, mutation and resolve fields.
// This struct must be implemented runtime.SchemaBuilder interface.
type graphql__resolver_ActionService struct {

	// Automatic connection host
	host string

	// grpc dial options
	dialOptions []grpc.DialOption

	// grpc client connection.
	// this connection may be provided by user
	conn *grpc.ClientConn
}

// new_graphql_resolver_ActionService creates pointer of service struct
func new_graphql_resolver_ActionService(conn *grpc.ClientConn) *graphql__resolver_ActionService {
	return &graphql__resolver_ActionService{
		conn:        conn,
		host:        "localhost:50051",
		dialOptions: []grpc.DialOption{},
	}
}

// CreateConnection() returns grpc connection which user specified or newly connected and closing function
func (x *graphql__resolver_ActionService) CreateConnection(ctx context.Context) (*grpc.ClientConn, func(), error) {
	// If x.conn is not nil, user injected their own connection
	if x.conn != nil {
		return x.conn, func() {}, nil
	}

	// Otherwise, this handler opens connection with specified host
	conn, err := grpc.DialContext(ctx, x.host, x.dialOptions...)
	if err != nil {
		return nil, nil, err
	}
	return conn, func() { conn.Close() }, nil
}

// GetQueries returns acceptable graphql.Fields for Query.
func (x *graphql__resolver_ActionService) GetQueries(conn *grpc.ClientConn) graphql.Fields {
	return graphql.Fields{
		"GetActionByVoter": &graphql.Field{
			Type: Gql__type_ActionResponse(),
			Args: graphql.FieldConfigArgument{
				"address": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"actHash": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"pagination": &graphql.ArgumentConfig{
					Type: pagination.Gql__input_Pagination(),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var req ActionRequest
				if err := runtime.MarshalRequest(p.Args, &req, false); err != nil {
					return nil, errors.Wrap(err, "Failed to marshal request for GetActionByVoter")
				}
				client := NewActionServiceClient(conn)
				resp, err := client.GetActionByVoter(p.Context, &req)
				if err != nil {
					return nil, errors.Wrap(err, "Failed to call RPC GetActionByVoter")
				}
				return resp, nil
			},
		},
		"GetEvmTransfersByAddress": &graphql.Field{
			Type: Gql__type_ActionResponse(),
			Args: graphql.FieldConfigArgument{
				"address": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"actHash": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"pagination": &graphql.ArgumentConfig{
					Type: pagination.Gql__input_Pagination(),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var req ActionRequest
				if err := runtime.MarshalRequest(p.Args, &req, false); err != nil {
					return nil, errors.Wrap(err, "Failed to marshal request for GetEvmTransfersByAddress")
				}
				client := NewActionServiceClient(conn)
				resp, err := client.GetEvmTransfersByAddress(p.Context, &req)
				if err != nil {
					return nil, errors.Wrap(err, "Failed to call RPC GetEvmTransfersByAddress")
				}
				return resp, nil
			},
		},
		"GetXrc20ByAddress": &graphql.Field{
			Type: Gql__type_ActionResponse(),
			Args: graphql.FieldConfigArgument{
				"address": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"actHash": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"pagination": &graphql.ArgumentConfig{
					Type: pagination.Gql__input_Pagination(),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var req ActionRequest
				if err := runtime.MarshalRequest(p.Args, &req, false); err != nil {
					return nil, errors.Wrap(err, "Failed to marshal request for GetXrc20ByAddress")
				}
				client := NewActionServiceClient(conn)
				resp, err := client.GetXrc20ByAddress(p.Context, &req)
				if err != nil {
					return nil, errors.Wrap(err, "Failed to call RPC GetXrc20ByAddress")
				}
				return resp, nil
			},
		},
	}
}

// GetMutations returns acceptable graphql.Fields for Mutation.
func (x *graphql__resolver_ActionService) GetMutations(conn *grpc.ClientConn) graphql.Fields {
	return graphql.Fields{}
}

// Register package divided graphql handler "without" *grpc.ClientConn,
// therefore gRPC connection will be opened and closed automatically.
// Occasionally you may worry about open/close performance for each handling graphql request,
// then you can call RegisterActionServiceGraphqlHandler with *grpc.ClientConn manually.
func RegisterActionServiceGraphql(mux *runtime.ServeMux) error {
	return RegisterActionServiceGraphqlHandler(mux, nil)
}

// Register package divided graphql handler "with" *grpc.ClientConn.
// this function accepts your defined grpc connection, so that we reuse that and never close connection inside.
// You need to close it maunally when application will terminate.
// Otherwise, you can specify automatic opening connection with ServiceOption directive:
//
// service ActionService {
//    option (graphql.service) = {
//        host: "host:port"
//        insecure: true or false
//    };
//
//    ...with RPC definitions
// }
func RegisterActionServiceGraphqlHandler(mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return mux.AddHandler(new_graphql_resolver_ActionService(conn))
}
