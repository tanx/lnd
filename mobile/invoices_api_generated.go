// +build invoicesrpc

// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: invoicesrpc/invoices.proto
package lndmobile

import (
	"context"
	"net"
	"time"

	"google.golang.org/grpc"

	"github.com/golang/protobuf/proto"

	"github.com/lightningnetwork/lnd/lnrpc/invoicesrpc"
)

func getInvoicesConn() (*grpc.ClientConn, func(), error) {
	conn, err := lightningLis.Dial()
	if err != nil {
		return nil, nil, err
	}

	clientConn, err := grpc.Dial("",
		grpc.WithDialer(func(target string,
			timeout time.Duration) (net.Conn, error) {
			return conn, nil
		}),
		grpc.WithInsecure(),
		grpc.WithBackoffMaxDelay(10*time.Second),
	)
	if err != nil {
		conn.Close()
		return nil, nil, err
	}

	close := func() {
		conn.Close()
	}

	return clientConn, close, nil
}

// getInvoicesClient returns a client connection to the server listening
// on lis.
func getInvoicesClient() (invoicesrpc.InvoicesClient, func(), error) {
	clientConn, close, err := getInvoicesConn()
	if err != nil {
		return nil, nil, err
	}
	client := invoicesrpc.NewInvoicesClient(clientConn)
	return client, close, nil
}

// SubscribeSingleInvoice returns a uni-directional stream (server -> client)
// to notify the client of state transitions of the specified invoice.
// Initially the current invoice state is always sent out.
//
// NOTE: This method produces a stream of responses, and the receive stream can
// be called zero or more times. After EOF error is returned, no more responses
// will be produced.
func SubscribeSingleInvoice(msg []byte, rStream RecvStream) {
	s := &readStreamHandler{
		newProto: func() proto.Message {
			return &invoicesrpc.SubscribeSingleInvoiceRequest{}
		},
		recvStream: func(ctx context.Context,
			req proto.Message) (*receiver, func(), error) {

			// Get the gRPC client.
			client, close, err := getInvoicesClient()
			if err != nil {
				return nil, nil, err
			}

			r := req.(*invoicesrpc.SubscribeSingleInvoiceRequest)
			stream, err := client.SubscribeSingleInvoice(ctx, r)
			if err != nil {
				close()
				return nil, nil, err
			}
			return &receiver{
				recv: func() (proto.Message, error) {
					return stream.Recv()
				},
			}, close, nil
		},
	}
	s.start(msg, rStream)
}

// CancelInvoice cancels a currently open invoice. If the invoice is already 
// canceled, this call will succeed. If the invoice is already settled, it will
// fail.
//
// NOTE: This method produces a single result or error, and the callback will
// be called only once.
func CancelInvoice(msg []byte, callback Callback) {
	s := &syncHandler{
		newProto: func() proto.Message {
			return &invoicesrpc.CancelInvoiceMsg{}
		},
		getSync: func(ctx context.Context,
			req proto.Message) (proto.Message, error) {

			// Get the gRPC client.
			client, close, err := getInvoicesClient()
			if err != nil {
				return nil, err
			}
			defer close()

			r := req.(*invoicesrpc.CancelInvoiceMsg)
			return client.CancelInvoice(ctx, r)
		},
	}
	s.start(msg, callback)
}

// AddHoldInvoice creates a hold invoice. It ties the invoice to the hash
// supplied in the request.
//
// NOTE: This method produces a single result or error, and the callback will
// be called only once.
func AddHoldInvoice(msg []byte, callback Callback) {
	s := &syncHandler{
		newProto: func() proto.Message {
			return &invoicesrpc.AddHoldInvoiceRequest{}
		},
		getSync: func(ctx context.Context,
			req proto.Message) (proto.Message, error) {

			// Get the gRPC client.
			client, close, err := getInvoicesClient()
			if err != nil {
				return nil, err
			}
			defer close()

			r := req.(*invoicesrpc.AddHoldInvoiceRequest)
			return client.AddHoldInvoice(ctx, r)
		},
	}
	s.start(msg, callback)
}

// SettleInvoice settles an accepted invoice. If the invoice is already
// settled, this call will succeed.
//
// NOTE: This method produces a single result or error, and the callback will
// be called only once.
func SettleInvoice(msg []byte, callback Callback) {
	s := &syncHandler{
		newProto: func() proto.Message {
			return &invoicesrpc.SettleInvoiceMsg{}
		},
		getSync: func(ctx context.Context,
			req proto.Message) (proto.Message, error) {

			// Get the gRPC client.
			client, close, err := getInvoicesClient()
			if err != nil {
				return nil, err
			}
			defer close()

			r := req.(*invoicesrpc.SettleInvoiceMsg)
			return client.SettleInvoice(ctx, r)
		},
	}
	s.start(msg, callback)
}