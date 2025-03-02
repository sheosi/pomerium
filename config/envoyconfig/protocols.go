package envoyconfig

import (
	"context"

	envoy_config_core_v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	envoy_extensions_http_header_formatters_preserve_case_v3 "github.com/envoyproxy/go-control-plane/envoy/extensions/http/header_formatters/preserve_case/v3"
	envoy_extensions_upstreams_http_v3 "github.com/envoyproxy/go-control-plane/envoy/extensions/upstreams/http/v3"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/pomerium/pomerium/config"
	"github.com/pomerium/pomerium/internal/log"
)

type upstreamProtocolConfig byte

const (
	upstreamProtocolAuto upstreamProtocolConfig = iota
	upstreamProtocolHTTP2
	upstreamProtocolHTTP1
)

// recommended defaults: https://www.envoyproxy.io/docs/envoy/latest/configuration/best_practices/edge
const (
	connectionBufferLimit            uint32 = 32 * 1024
	maxConcurrentStreams             uint32 = 100
	initialStreamWindowSizeLimit     uint32 = 64 * 1024
	initialConnectionWindowSizeLimit uint32 = 1 * 1024 * 1024
)

var http1ProtocolOptions = &envoy_config_core_v3.Http1ProtocolOptions{
	// fix for #3935, preserve case of HTTP headers for applications that are case-sensitive
	HeaderKeyFormat: &envoy_config_core_v3.Http1ProtocolOptions_HeaderKeyFormat{
		HeaderFormat: &envoy_config_core_v3.Http1ProtocolOptions_HeaderKeyFormat_StatefulFormatter{
			StatefulFormatter: &envoy_config_core_v3.TypedExtensionConfig{
				Name:        "preserve_case",
				TypedConfig: marshalAny(&envoy_extensions_http_header_formatters_preserve_case_v3.PreserveCaseFormatterConfig{}),
			},
		},
	},
}

var http2ProtocolOptions = &envoy_config_core_v3.Http2ProtocolOptions{
	AllowConnect:                true,
	MaxConcurrentStreams:        wrapperspb.UInt32(maxConcurrentStreams),
	InitialStreamWindowSize:     wrapperspb.UInt32(initialStreamWindowSizeLimit),
	InitialConnectionWindowSize: wrapperspb.UInt32(initialConnectionWindowSizeLimit),
}

func buildTypedExtensionProtocolOptions(
	endpoints []Endpoint,
	upstreamProtocol upstreamProtocolConfig,
) map[string]*anypb.Any {
	return map[string]*anypb.Any{
		"envoy.extensions.upstreams.http.v3.HttpProtocolOptions": marshalAny(buildUpstreamProtocolOptions(endpoints, upstreamProtocol)),
	}
}

func buildUpstreamProtocolOptions(
	endpoints []Endpoint,
	upstreamProtocol upstreamProtocolConfig,
) *envoy_extensions_upstreams_http_v3.HttpProtocolOptions {
	switch upstreamProtocol {
	case upstreamProtocolHTTP2:
		// when explicitly configured, force HTTP/2
		return &envoy_extensions_upstreams_http_v3.HttpProtocolOptions{
			UpstreamProtocolOptions: &envoy_extensions_upstreams_http_v3.HttpProtocolOptions_ExplicitHttpConfig_{
				ExplicitHttpConfig: &envoy_extensions_upstreams_http_v3.HttpProtocolOptions_ExplicitHttpConfig{
					ProtocolConfig: &envoy_extensions_upstreams_http_v3.HttpProtocolOptions_ExplicitHttpConfig_Http2ProtocolOptions{
						Http2ProtocolOptions: http2ProtocolOptions,
					},
				},
			},
		}
	case upstreamProtocolAuto:
		// when using TLS use ALPN auto config
		tlsCount := 0
		for _, e := range endpoints {
			if e.transportSocket != nil {
				tlsCount++
			}
		}
		if tlsCount > 0 && tlsCount == len(endpoints) {
			return &envoy_extensions_upstreams_http_v3.HttpProtocolOptions{
				UpstreamProtocolOptions: &envoy_extensions_upstreams_http_v3.HttpProtocolOptions_AutoConfig{
					AutoConfig: &envoy_extensions_upstreams_http_v3.HttpProtocolOptions_AutoHttpConfig{
						HttpProtocolOptions:  http1ProtocolOptions,
						Http2ProtocolOptions: http2ProtocolOptions,
					},
				},
			}
		}
	}

	// otherwise only use http/1.1
	return &envoy_extensions_upstreams_http_v3.HttpProtocolOptions{
		UpstreamProtocolOptions: &envoy_extensions_upstreams_http_v3.HttpProtocolOptions_ExplicitHttpConfig_{
			ExplicitHttpConfig: &envoy_extensions_upstreams_http_v3.HttpProtocolOptions_ExplicitHttpConfig{
				ProtocolConfig: &envoy_extensions_upstreams_http_v3.HttpProtocolOptions_ExplicitHttpConfig_HttpProtocolOptions{
					HttpProtocolOptions: http1ProtocolOptions,
				},
			},
		},
	}
}

func buildUpstreamALPN(upstreamProtocol upstreamProtocolConfig) []string {
	switch upstreamProtocol {
	case upstreamProtocolAuto:
		return []string{"h2", "http/1.1"}
	case upstreamProtocolHTTP2:
		return []string{"h2"}
	default:
		return []string{"http/1.1"}
	}
}

func getUpstreamProtocolForPolicy(ctx context.Context, policy *config.Policy) upstreamProtocolConfig {
	upstreamProtocol := upstreamProtocolAuto
	if policy.AllowWebsockets {
		// #2388, force http/1 when using web sockets
		log.WarnWebSocketHTTP1_1(getClusterID(policy))
		upstreamProtocol = upstreamProtocolHTTP1
	}
	return upstreamProtocol
}
