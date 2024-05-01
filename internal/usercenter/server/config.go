package server

import genericoptions "github.com/costa92/k8s-krm-go/pkg/options"

type Config struct {
	HTTP genericoptions.HTTPOptions
	GRPC genericoptions.GRPCOptions
}
