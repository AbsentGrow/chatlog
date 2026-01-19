// Package grpcserver provides the gRPC server implementation.
package grpcserver

import (
	"context"
	"fmt"
	"net"
	"runtime/debug"

	"github.com/sjzar/chatlog/internal/chatlog"
	iwechat "github.com/sjzar/chatlog/internal/wechat"
	"github.com/sjzar/chatlog/pkg/logger"
	"github.com/sjzar/chatlog/pkg/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

// Server implements the gRPC ManagerService.
var _ pb.ManagerServiceServer = (*Server)(nil)

type Server struct {
	pb.UnimplementedManagerServiceServer
	manager chatlog.Manager
	server  *grpc.Server
}

func (s *Server) Backup(ctx context.Context, request *pb.BackupRequest) (*pb.BackupResponse, error) {
	//TODO implement me
	panic("implement me")
}

// New creates a new gRPC server.
func New(mgr chatlog.Manager) *Server {
	return &Server{
		manager: mgr,
	}
}

// Start starts the gRPC server on the given address.
func (s *Server) Start(addr string) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	// Create gRPC server with recovery interceptor
	s.server = grpc.NewServer(
		grpc.UnaryInterceptor(s.recoveryInterceptor),
	)

	// Register our service
	pb.RegisterManagerServiceServer(s.server, s)

	// Enable reflection for debugging (grpcurl, grpc-ui)
	reflection.Register(s.server)

	logger.Info().Str("addr", addr).Msg("gRPC server starting...")

	return s.server.Serve(lis)
}

// Stop gracefully stops the gRPC server.
func (s *Server) Stop() {
	if s.server != nil {
		s.server.GracefulStop()
	}
}

// recoveryInterceptor catches panics and returns gRPC errors.
func (s *Server) recoveryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			stack := string(debug.Stack())
			logger.Error().
				Interface("panic", r).
				Str("method", info.FullMethod).
				Str("stack", stack).
				Msg("Panic recovered in gRPC handler")
			err = status.Errorf(codes.Internal, "internal panic: %v", r)
		}
	}()
	return handler(ctx, req)
}

// SetLogLevel changes the log level.
func (s *Server) SetLogLevel(ctx context.Context, req *pb.SetLogLevelRequest) (*pb.SetLogLevelResponse, error) {
	logger.SetLevel(req.Level)
	logger.Info().Str("level", req.Level).Msg("Log level changed")
	return &pb.SetLogLevelResponse{}, nil
}

// Run executes the Run command.
func (s *Server) Run(ctx context.Context, req *pb.RunRequest) (*pb.RunResponse, error) {
	if err := s.manager.Run(req.ConfigPath); err != nil {
		return nil, status.Errorf(codes.Internal, "Run failed: %v", err)
	}
	return &pb.RunResponse{}, nil
}

// Switch switches to a different WeChat account.
func (s *Server) Switch(ctx context.Context, req *pb.SwitchRequest) (*pb.SwitchResponse, error) {
	var info *iwechat.Account
	if req.Info != nil {
		info = &iwechat.Account{
			Name:        req.Info.Name,
			Platform:    req.Info.Platform,
			Version:     int(req.Info.Version),
			FullVersion: req.Info.FullVersion,
			DataDir:     req.Info.DataDir,
			Key:         req.Info.Key,
			ImgKey:      req.Info.ImgKey,
			PID:         req.Info.Pid,
			ExePath:     req.Info.ExePath,
			Status:      req.Info.Status,
		}
	}

	if err := s.manager.Switch(info, req.History); err != nil {
		return nil, status.Errorf(codes.Internal, "Switch failed: %v", err)
	}
	return &pb.SwitchResponse{}, nil
}

// StartService starts the service.
func (s *Server) StartService(ctx context.Context, req *pb.StartServiceRequest) (*pb.StartServiceResponse, error) {
	if err := s.manager.StartService(); err != nil {
		return nil, status.Errorf(codes.Internal, "StartService failed: %v", err)
	}
	return &pb.StartServiceResponse{}, nil
}

// StopService stops the service.
func (s *Server) StopService(ctx context.Context, req *pb.StopServiceRequest) (*pb.StopServiceResponse, error) {
	if err := s.manager.StopService(); err != nil {
		return nil, status.Errorf(codes.Internal, "StopService failed: %v", err)
	}
	return &pb.StopServiceResponse{}, nil
}

// SetHTTPAddr sets the HTTP address.
func (s *Server) SetHTTPAddr(ctx context.Context, req *pb.SetHTTPAddrRequest) (*pb.SetHTTPAddrResponse, error) {
	if err := s.manager.SetHTTPAddr(req.Text); err != nil {
		return nil, status.Errorf(codes.Internal, "SetHTTPAddr failed: %v", err)
	}
	return &pb.SetHTTPAddrResponse{}, nil
}

// GetDataKey gets the data key.
func (s *Server) GetDataKey(ctx context.Context, req *pb.GetDataKeyRequest) (*pb.GetDataKeyResponse, error) {
	if err := s.manager.GetDataKey(); err != nil {
		return nil, status.Errorf(codes.Internal, "GetDataKey failed: %v", err)
	}
	return &pb.GetDataKeyResponse{}, nil
}

// DecryptDBFiles decrypts database files.
func (s *Server) DecryptDBFiles(ctx context.Context, req *pb.DecryptDBFilesRequest) (*pb.DecryptDBFilesResponse, error) {
	if err := s.manager.DecryptDBFiles(); err != nil {
		return nil, status.Errorf(codes.Internal, "DecryptDBFiles failed: %v", err)
	}
	return &pb.DecryptDBFilesResponse{}, nil
}

// StartAutoDecrypt starts auto decryption.
func (s *Server) StartAutoDecrypt(ctx context.Context, req *pb.StartAutoDecryptRequest) (*pb.StartAutoDecryptResponse, error) {
	if err := s.manager.StartAutoDecrypt(); err != nil {
		return nil, status.Errorf(codes.Internal, "StartAutoDecrypt failed: %v", err)
	}
	return &pb.StartAutoDecryptResponse{}, nil
}

// StopAutoDecrypt stops auto decryption.
func (s *Server) StopAutoDecrypt(ctx context.Context, req *pb.StopAutoDecryptRequest) (*pb.StopAutoDecryptResponse, error) {
	if err := s.manager.StopAutoDecrypt(); err != nil {
		return nil, status.Errorf(codes.Internal, "StopAutoDecrypt failed: %v", err)
	}
	return &pb.StopAutoDecryptResponse{}, nil
}

// RefreshSession refreshes the session.
func (s *Server) RefreshSession(ctx context.Context, req *pb.RefreshSessionRequest) (*pb.RefreshSessionResponse, error) {
	if err := s.manager.RefreshSession(); err != nil {
		return nil, status.Errorf(codes.Internal, "RefreshSession failed: %v", err)
	}
	return &pb.RefreshSessionResponse{}, nil
}

// CommandKey executes the key command.
func (s *Server) CommandKey(ctx context.Context, req *pb.CommandKeyRequest) (*pb.CommandKeyResponse, error) {
	result, err := s.manager.CommandKey(req.ConfigPath, int(req.Pid), req.Force, req.ShowXorKey)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "CommandKey failed: %v", err)
	}
	return &pb.CommandKeyResponse{Result: result}, nil
}

// CommandDecrypt executes the decrypt command.
func (s *Server) CommandDecrypt(ctx context.Context, req *pb.CommandDecryptRequest) (*pb.CommandDecryptResponse, error) {
	cmdConf := make(map[string]any, len(req.CmdConf))
	for k, v := range req.CmdConf {
		cmdConf[k] = v
	}
	if err := s.manager.CommandDecrypt(req.ConfigPath, cmdConf); err != nil {
		return nil, status.Errorf(codes.Internal, "CommandDecrypt failed: %v", err)
	}
	return &pb.CommandDecryptResponse{}, nil
}

// CommandHTTPServer executes the HTTP server command.
func (s *Server) CommandHTTPServer(ctx context.Context, req *pb.CommandHTTPServerRequest) (*pb.CommandHTTPServerResponse, error) {
	cmdConf := make(map[string]any, len(req.CmdConf))
	for k, v := range req.CmdConf {
		cmdConf[k] = v
	}
	if err := s.manager.CommandHTTPServer(req.ConfigPath, cmdConf); err != nil {
		return nil, status.Errorf(codes.Internal, "CommandHTTPServer failed: %v", err)
	}
	return &pb.CommandHTTPServerResponse{}, nil
}

// GetWeChatInstances gets all WeChat instances.
func (s *Server) GetWeChatInstances(ctx context.Context, req *pb.GetWeChatInstancesRequest) (*pb.GetWeChatInstancesResponse, error) {
	accounts := s.manager.GetWeChatInstances()

	pbAccounts := make([]*pb.Account, len(accounts))
	for i, acc := range accounts {
		pbAccounts[i] = &pb.Account{
			Name:        acc.Name,
			Platform:    acc.Platform,
			Version:     int32(acc.Version),
			FullVersion: acc.FullVersion,
			DataDir:     acc.DataDir,
			Key:         acc.Key,
			ImgKey:      acc.ImgKey,
			Pid:         acc.PID,
			ExePath:     acc.ExePath,
			Status:      acc.Status,
		}
	}

	return &pb.GetWeChatInstancesResponse{Accounts: pbAccounts}, nil
}

// GetKey gets the encryption key.
func (s *Server) GetKey(ctx context.Context, req *pb.GetKeyRequest) (*pb.GetKeyResponse, error) {
	keyData, err := s.manager.GetKey(req.ConfigPath, int(req.Pid), req.Force, req.ShowXorKey)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "GetKey failed: %v", err)
	}

	return &pb.GetKeyResponse{
		Data: &pb.KeyData{
			Key:    keyData.DataKey,
			ImgKey: keyData.ImageKey,
		},
	}, nil
}

// Decrypt decrypts files.
func (s *Server) Decrypt(ctx context.Context, req *pb.DecryptRequest) (*pb.DecryptResponse, error) {
	cmdConf := make(map[string]any, len(req.CmdConf))
	for k, v := range req.CmdConf {
		cmdConf[k] = v
	}
	if err := s.manager.Decrypt(req.ConfigPath, cmdConf); err != nil {
		return nil, status.Errorf(codes.Internal, "Decrypt failed: %v", err)
	}
	return &pb.DecryptResponse{}, nil
}
