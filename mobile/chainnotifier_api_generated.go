// Code generated by falafel 0.3. DO NOT EDIT.
// source: chainrpc/chainnotifier.proto

// +build chainrpc

package lndmobile

import (
	"context"
	"net"
	"time"

	"google.golang.org/grpc"

	"github.com/golang/protobuf/proto"

	"github.com/lightningnetwork/lnd/lnrpc/chainrpc"
)

func getChainNotifierConn() (*grpc.ClientConn, func(), error) {
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

// getChainNotifierClient returns a client connection to the server listening
// on lis.
func getChainNotifierClient() (chainrpc.ChainNotifierClient, func(), error) {
	clientConn, close, err := getChainNotifierConn()
	if err != nil {
		return nil, nil, err
	}
	client := chainrpc.NewChainNotifierClient(clientConn)
	return client, close, nil
}

// RegisterConfirmationsNtfn is a synchronous response-streaming RPC that
// registers an intent for a client to be notified once a confirmation request
// has reached its required number of confirmations on-chain.
//
// A client can specify whether the confirmation request should be for a
// particular transaction by its hash or for an output script by specifying a
// zero hash.
//
// NOTE: This method produces a stream of responses, and the receive stream can
// be called zero or more times. After EOF error is returned, no more responses
// will be produced.
func RegisterConfirmationsNtfn(msg []byte, rStream RecvStream) {
	s := &readStreamHandler{
		newProto: func() proto.Message {
			return &chainrpc.ConfRequest{}
		},
		recvStream: func(ctx context.Context,
			req proto.Message) (*receiver, func(), error) {

			// Get the gRPC client.
			client, close, err := getChainNotifierClient()
			if err != nil {
				return nil, nil, err
			}

			r := req.(*chainrpc.ConfRequest)
			stream, err := client.RegisterConfirmationsNtfn(ctx, r)
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

// RegisterSpendNtfn is a synchronous response-streaming RPC that registers an
// intent for a client to be notification once a spend request has been spent
// by a transaction that has confirmed on-chain.
//
// A client can specify whether the spend request should be for a particular
// outpoint  or for an output script by specifying a zero outpoint.
//
// NOTE: This method produces a stream of responses, and the receive stream can
// be called zero or more times. After EOF error is returned, no more responses
// will be produced.
func RegisterSpendNtfn(msg []byte, rStream RecvStream) {
	s := &readStreamHandler{
		newProto: func() proto.Message {
			return &chainrpc.SpendRequest{}
		},
		recvStream: func(ctx context.Context,
			req proto.Message) (*receiver, func(), error) {

			// Get the gRPC client.
			client, close, err := getChainNotifierClient()
			if err != nil {
				return nil, nil, err
			}

			r := req.(*chainrpc.SpendRequest)
			stream, err := client.RegisterSpendNtfn(ctx, r)
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

// RegisterBlockEpochNtfn is a synchronous response-streaming RPC that
// registers an intent for a client to be notified of blocks in the chain. The
// stream will return a hash and height tuple of a block for each new/stale
// block in the chain. It is the client's responsibility to determine whether
// the tuple returned is for a new or stale block in the chain.
//
// A client can also request a historical backlog of blocks from a particular
// point. This allows clients to be idempotent by ensuring that they do not
// missing processing a single block within the chain.
//
// NOTE: This method produces a stream of responses, and the receive stream can
// be called zero or more times. After EOF error is returned, no more responses
// will be produced.
func RegisterBlockEpochNtfn(msg []byte, rStream RecvStream) {
	s := &readStreamHandler{
		newProto: func() proto.Message {
			return &chainrpc.BlockEpoch{}
		},
		recvStream: func(ctx context.Context,
			req proto.Message) (*receiver, func(), error) {

			// Get the gRPC client.
			client, close, err := getChainNotifierClient()
			if err != nil {
				return nil, nil, err
			}

			r := req.(*chainrpc.BlockEpoch)
			stream, err := client.RegisterBlockEpochNtfn(ctx, r)
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
