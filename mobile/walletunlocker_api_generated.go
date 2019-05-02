

// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: rpc.proto
package lndmobile

import (
	"context"
	"net"
	"time"

	"google.golang.org/grpc"

	"github.com/golang/protobuf/proto"

	"github.com/lightningnetwork/lnd/lnrpc"
)

func getWalletUnlockerConn() (*grpc.ClientConn, func(), error) {
	conn, err := walletUnlockerLis.Dial()
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

// getWalletUnlockerClient returns a client connection to the server listening
// on lis.
func getWalletUnlockerClient() (lnrpc.WalletUnlockerClient, func(), error) {
	clientConn, close, err := getWalletUnlockerConn()
	if err != nil {
		return nil, nil, err
	}
	client := lnrpc.NewWalletUnlockerClient(clientConn)
	return client, close, nil
}

// GenSeed is the first method that should be used to instantiate a new lnd
// instance. This method allows a caller to generate a new aezeed cipher seed
// given an optional passphrase. If provided, the passphrase will be necessary
// to decrypt the cipherseed to expose the internal wallet seed.
// 
// Once the cipherseed is obtained and verified by the user, the InitWallet
// method should be used to commit the newly generated seed, and create the
// wallet.
//
// NOTE: This method produces a single result or error, and the callback will
// be called only once.
func GenSeed(msg []byte, callback Callback) {
	s := &syncHandler{
		newProto: func() proto.Message {
			return &lnrpc.GenSeedRequest{}
		},
		getSync: func(ctx context.Context,
			req proto.Message) (proto.Message, error) {

			// Get the gRPC client.
			client, close, err := getWalletUnlockerClient()
			if err != nil {
				return nil, err
			}
			defer close()

			r := req.(*lnrpc.GenSeedRequest)
			return client.GenSeed(ctx, r)
		},
	}
	s.start(msg, callback)
}

// InitWallet is used when lnd is starting up for the first time to fully
// initialize the daemon and its internal wallet. At the very least a wallet
// password must be provided. This will be used to encrypt sensitive material
// on disk.
// 
// In the case of a recovery scenario, the user can also specify their aezeed
// mnemonic and passphrase. If set, then the daemon will use this prior state
// to initialize its internal wallet.
// 
// Alternatively, this can be used along with the GenSeed RPC to obtain a
// seed, then present it to the user. Once it has been verified by the user,
// the seed can be fed into this RPC in order to commit the new wallet.
//
// NOTE: This method produces a single result or error, and the callback will
// be called only once.
func InitWallet(msg []byte, callback Callback) {
	s := &syncHandler{
		newProto: func() proto.Message {
			return &lnrpc.InitWalletRequest{}
		},
		getSync: func(ctx context.Context,
			req proto.Message) (proto.Message, error) {

			// Get the gRPC client.
			client, close, err := getWalletUnlockerClient()
			if err != nil {
				return nil, err
			}
			defer close()

			r := req.(*lnrpc.InitWalletRequest)
			return client.InitWallet(ctx, r)
		},
	}
	s.start(msg, callback)
}

// UnlockWallet is used at startup of lnd to provide a password to unlock
// the wallet database.
//
// NOTE: This method produces a single result or error, and the callback will
// be called only once.
func UnlockWallet(msg []byte, callback Callback) {
	s := &syncHandler{
		newProto: func() proto.Message {
			return &lnrpc.UnlockWalletRequest{}
		},
		getSync: func(ctx context.Context,
			req proto.Message) (proto.Message, error) {

			// Get the gRPC client.
			client, close, err := getWalletUnlockerClient()
			if err != nil {
				return nil, err
			}
			defer close()

			r := req.(*lnrpc.UnlockWalletRequest)
			return client.UnlockWallet(ctx, r)
		},
	}
	s.start(msg, callback)
}

// ChangePassword changes the password of the encrypted wallet. This will
// automatically unlock the wallet database if successful.
//
// NOTE: This method produces a single result or error, and the callback will
// be called only once.
func ChangePassword(msg []byte, callback Callback) {
	s := &syncHandler{
		newProto: func() proto.Message {
			return &lnrpc.ChangePasswordRequest{}
		},
		getSync: func(ctx context.Context,
			req proto.Message) (proto.Message, error) {

			// Get the gRPC client.
			client, close, err := getWalletUnlockerClient()
			if err != nil {
				return nil, err
			}
			defer close()

			r := req.(*lnrpc.ChangePasswordRequest)
			return client.ChangePassword(ctx, r)
		},
	}
	s.start(msg, callback)
}
