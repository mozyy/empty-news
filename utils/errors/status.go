package errors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrMissingMetadata  = status.Errorf(codes.InvalidArgument, "missing metadata")
	ErrInvalidToken     = status.Errorf(codes.Unauthenticated, "invalid token")
	ErrInvalidArgument  = status.Errorf(codes.InvalidArgument, "参数错误")
	ErrPermissionDenied = status.Errorf(codes.PermissionDenied, "permission denied")
)
