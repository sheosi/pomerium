// Package main contains the pomerium integration tests
package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"testing"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/publicsuffix"
)

var IDP, ClusterType string

func TestMain(m *testing.M) {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	flag.Parse()
	if testing.Verbose() {
		log.Logger = log.Logger.Level(zerolog.DebugLevel)
	} else {
		log.Logger = log.Logger.Level(zerolog.InfoLevel)
	}

	logger := log.With().Logger()
	ctx := logger.WithContext(context.Background())

	if err := waitForHealthy(ctx); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "services not healthy")
		os.Exit(1)
		return
	}

	setIDPAndClusterType(ctx)

	status := m.Run()
	os.Exit(status)
}

type loggingRoundTripper struct {
	t         testing.TB
	transport http.RoundTripper
}

func (l loggingRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if l.t != nil {
		l.t.Logf("%s %s", req.Method, req.URL.String())
	}
	return l.transport.RoundTrip(req)
}

func getTransport(t testing.TB) http.RoundTripper {
	if t != nil {
		t.Helper()
	}

	rootCAs, err := x509.SystemCertPool()
	if err != nil {
		panic(err)
	}

	bs, err := os.ReadFile(filepath.Join(".", "tpl", "files", "ca.pem"))
	if err != nil {
		panic(err)
	}
	_ = rootCAs.AppendCertsFromPEM(bs)
	transport := &http.Transport{
		DisableKeepAlives: true,
		TLSClientConfig: &tls.Config{
			RootCAs: rootCAs,
		},
	}
	return loggingRoundTripper{t, transport}
}

func getClient(t testing.TB) *http.Client {
	if t != nil {
		t.Helper()
	}

	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		panic(err)
	}

	return &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Transport: getTransport(t),
		Jar:       jar,
	}
}

func waitForHealthy(ctx context.Context) error {
	client := getClient(nil)
	check := func(endpoint string) error {
		reqCtx, clearTimeout := context.WithTimeout(ctx, time.Second)
		defer clearTimeout()

		req, err := http.NewRequestWithContext(reqCtx, "GET", endpoint, nil)
		if err != nil {
			return err
		}

		res, err := client.Do(req)
		if err != nil {
			return err
		}
		defer res.Body.Close()

		if res.StatusCode/100 != 2 {
			return fmt.Errorf("%s unavailable: %s", endpoint, res.Status)
		}

		log.Info().Int("status", res.StatusCode).Msgf("%s healthy", endpoint)

		return nil
	}

	ticker := time.NewTicker(time.Second * 3)
	defer ticker.Stop()

	endpoints := []string{
		"https://authenticate.localhost.pomerium.io/.well-known/pomerium/jwks.json",
		"https://mock-idp.localhost.pomerium.io/.well-known/jwks.json",
	}

	for {
		var err error
		for _, endpoint := range endpoints {
			err = check(endpoint)
			if err != nil {
				break
			}
		}
		if err == nil {
			return nil
		}

		log.Ctx(ctx).Info().Err(err).Msg("waiting for healthy")

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
		}
	}
}

func setIDPAndClusterType(ctx context.Context) {
	IDP = "oidc"
	ClusterType = "single"

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Error().Err(err).Msg("failed to create docker client")
		return
	}

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		log.Error().Err(err).Msg("failed to retrieve docker containers")
	}
	for _, container := range containers {
		for _, name := range container.Names {
			parts := regexp.MustCompile(`^/(\w+?)[-_]pomerium.*$`).FindStringSubmatch(name)
			if len(parts) == 2 {
				ClusterType = parts[1]
			}
		}
	}

	log.Info().Str("idp", IDP).Str("cluster-type", ClusterType).Send()
}

func mustParseURL(str string) *url.URL {
	u, err := url.Parse(str)
	if err != nil {
		panic(err)
	}
	return u
}
