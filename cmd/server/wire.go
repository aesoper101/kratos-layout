//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/aeoper101/kratos-layout/internal/biz"
	"github.com/aeoper101/kratos-layout/internal/conf"
	"github.com/aeoper101/kratos-layout/internal/data"
	"github.com/aeoper101/kratos-layout/internal/i18n"
	"github.com/aeoper101/kratos-layout/internal/server"
	"github.com/aeoper101/kratos-layout/internal/service"
	"github.com/aesoper101/kratos-utils/protobuf/types/confpb"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, *confpb.Registry, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, i18n.ProviderSet, newApp))
}
