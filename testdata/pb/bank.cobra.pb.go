// Code generated by tetratelabs/protoc-gen-cobra.
// source: pb/bank.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	pb/bank.proto

It has these top-level commands:
	BankClientCommand
*/

package pb

import (
	proto "github.com/golang/protobuf/proto"
	cobra "github.com/spf13/cobra"
	context "golang.org/x/net/context"
	credentials "google.golang.org/grpc/credentials"
	filepath "path/filepath"
	fmt1 "fmt"
	grpc "google.golang.org/grpc"
	io "io"
	iocodec "github.com/dio/protoc-gen-cobra/iocodec"
	ioutil "io/ioutil"
	json "encoding/json"
	log "log"
	net "net"
	oauth "google.golang.org/grpc/credentials/oauth"
	oauth2 "golang.org/x/oauth2"
	os "os"
	pflag "github.com/spf13/pflag"
	template "text/template"
	time "time"
	tls "crypto/tls"
	x509 "crypto/x509"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

var _DefaultBankClientCommandConfig = _NewBankClientCommandConfig()

type _BankClientCommandConfig struct {
	ServerAddr         string
	RequestFile        string
	Stdin              bool
	PrintSampleRequest bool
	ResponseFormat     string
	Timeout            time.Duration
	TLS                bool
	ServerName         string
	InsecureSkipVerify bool
	CACertFile         string
	CertFile           string
	KeyFile            string
	AuthToken          string
	AuthTokenType      string
	JWTKey             string
	JWTKeyFile         string
}

func _NewBankClientCommandConfig() *_BankClientCommandConfig {
	c := &_BankClientCommandConfig{
		ServerAddr:     "localhost:8080",
		ResponseFormat: "json",
		Timeout:        10 * time.Second,
		AuthTokenType:  "Bearer",
	}
	return c
}

func (o *_BankClientCommandConfig) AddFlags(fs *pflag.FlagSet) {
	fs.StringVarP(&o.ServerAddr, "server-addr", "s", o.ServerAddr, "server address in form of host:port")
	fs.StringVarP(&o.RequestFile, "request-file", "f", o.RequestFile, "client request file (must be json, yaml, or xml); use \"-\" for stdin + json")
	fs.BoolVar(&o.Stdin, "stdin", o.Stdin, "read client request from STDIN; alternative for '-f -'")
	fs.BoolVarP(&o.PrintSampleRequest, "print-sample-request", "p", o.PrintSampleRequest, "print sample request file and exit")
	fs.StringVarP(&o.ResponseFormat, "response-format", "o", o.ResponseFormat, "response format (json, prettyjson, yaml, or xml)")
	fs.DurationVar(&o.Timeout, "timeout", o.Timeout, "client connection timeout")
	fs.BoolVar(&o.TLS, "tls", o.TLS, "enable tls")
	fs.StringVar(&o.ServerName, "tls-server-name", o.ServerName, "tls server name override")
	fs.BoolVar(&o.InsecureSkipVerify, "tls-insecure-skip-verify", o.InsecureSkipVerify, "INSECURE: skip tls checks")
	fs.StringVar(&o.CACertFile, "tls-ca-cert-file", o.CACertFile, "ca certificate file")
	fs.StringVar(&o.CertFile, "tls-cert-file", o.CertFile, "client certificate file")
	fs.StringVar(&o.KeyFile, "tls-key-file", o.KeyFile, "client key file")
	fs.StringVar(&o.AuthToken, "auth-token", o.AuthToken, "authorization token")
	fs.StringVar(&o.AuthTokenType, "auth-token-type", o.AuthTokenType, "authorization token type")
	fs.StringVar(&o.JWTKey, "jwt-key", o.JWTKey, "jwt key")
	fs.StringVar(&o.JWTKeyFile, "jwt-key-file", o.JWTKeyFile, "jwt key file")
}

func BankClientCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "bank",
	}
	_DefaultBankClientCommandConfig.AddFlags(cmd.PersistentFlags())

	for _, s := range _BankClientSubCommands {
		cmd.AddCommand(s())
	}
	return cmd
}

func _DialBank() (*grpc.ClientConn, BankClient, error) {
	cfg := _DefaultBankClientCommandConfig
	opts := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithTimeout(cfg.Timeout),
	}
	if cfg.TLS {
		tlsConfig := &tls.Config{}
		if cfg.InsecureSkipVerify {
			tlsConfig.InsecureSkipVerify = true
		}
		if cfg.CACertFile != "" {
			cacert, err := ioutil.ReadFile(cfg.CACertFile)
			if err != nil {
				return nil, nil, fmt.Errorf("ca cert: %v", err)
			}
			certpool := x509.NewCertPool()
			certpool.AppendCertsFromPEM(cacert)
			tlsConfig.RootCAs = certpool
		}
		if cfg.CertFile != "" {
			if cfg.KeyFile == "" {
				return nil, nil, fmt.Errorf("missing key file")
			}
			pair, err := tls.LoadX509KeyPair(cfg.CertFile, cfg.KeyFile)
			if err != nil {
				return nil, nil, fmt.Errorf("cert/key: %v", err)
			}
			tlsConfig.Certificates = []tls.Certificate{pair}
		}
		if cfg.ServerName != "" {
			tlsConfig.ServerName = cfg.ServerName
		} else {
			addr, _, _ := net.SplitHostPort(cfg.ServerAddr)
			tlsConfig.ServerName = addr
		}
		//tlsConfig.BuildNameToCertificate()
		cred := credentials.NewTLS(tlsConfig)
		opts = append(opts, grpc.WithTransportCredentials(cred))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}
	if cfg.AuthToken != "" {
		cred := oauth.NewOauthAccess(&oauth2.Token{
			AccessToken: cfg.AuthToken,
			TokenType:   cfg.AuthTokenType,
		})
		opts = append(opts, grpc.WithPerRPCCredentials(cred))
	}
	if cfg.JWTKey != "" {
		cred, err := oauth.NewJWTAccessFromKey([]byte(cfg.JWTKey))
		if err != nil {
			return nil, nil, fmt.Errorf("jwt key: %v", err)
		}
		opts = append(opts, grpc.WithPerRPCCredentials(cred))
	}
	if cfg.JWTKeyFile != "" {
		cred, err := oauth.NewJWTAccessFromFile(cfg.JWTKeyFile)
		if err != nil {
			return nil, nil, fmt.Errorf("jwt key file: %v", err)
		}
		opts = append(opts, grpc.WithPerRPCCredentials(cred))
	}
	conn, err := grpc.Dial(cfg.ServerAddr, opts...)
	if err != nil {
		return nil, nil, err
	}
	return conn, NewBankClient(conn), nil
}

type _BankRoundTripFunc func(cli BankClient, in iocodec.Decoder, out iocodec.Encoder) error

func _BankRoundTrip(sample interface{}, fn _BankRoundTripFunc) error {
	cfg := _DefaultBankClientCommandConfig
	var em iocodec.EncoderMaker
	var ok bool
	if cfg.ResponseFormat == "" {
		em = iocodec.DefaultEncoders["json"]
	} else {
		em, ok = iocodec.DefaultEncoders[cfg.ResponseFormat]
		if !ok {
			return fmt.Errorf("invalid response format: %q", cfg.ResponseFormat)
		}
	}
	if cfg.PrintSampleRequest {
		return em.NewEncoder(os.Stdout).Encode(sample)
	}
	// read the input request, first from stdin, then from a file, otherwise from args only
	var d iocodec.Decoder
	if cfg.Stdin || cfg.RequestFile == "-" {
		d = iocodec.DefaultDecoders["json"].NewDecoder(os.Stdin)
	} else if cfg.RequestFile != "" {
		f, err := os.Open(cfg.RequestFile)
		if err != nil {
			return fmt.Errorf("request file: %v", err)
		}
		defer f.Close()
		ext := filepath.Ext(cfg.RequestFile)
		if len(ext) > 0 && ext[0] == '.' {
			ext = ext[1:]
		}
		dm, ok := iocodec.DefaultDecoders[ext]
		if !ok {
			return fmt.Errorf("invalid request file format: %q", ext)
		}
		d = dm.NewDecoder(f)
	} else {
		d = iocodec.DefaultDecoders["noop"].NewDecoder(os.Stdin)
	}
	conn, client, err := _DialBank()
	if err != nil {
		return err
	}
	defer conn.Close()
	return fn(client, d, em.NewEncoder(os.Stdout))
}

func _BankDepositClientCommand() *cobra.Command {
	reqArgs := &DepositRequest{
		ClusterWithNamespaces: &DepositRequest_ClusterWithNamespaces{
			Cluster:    &Cluster{},
			Namespaces: []*DepositRequest_NamespaceWithDeployments{},
		},
		Clusters: []*DepositRequest_ClusterWithNamespaces{},
	}

	cmd := &cobra.Command{
		Use:     "deposit",
		Long:    "Deposit client; call by piping a request in to stdin (--stdin), reading a file (--file), or via flags per field",
		Example: "TODO: print protobuf method comments here",
		Run: func(cmd *cobra.Command, args []string) {
			var v DepositRequest
			err := _BankRoundTrip(v, func(cli BankClient, in iocodec.Decoder, out iocodec.Encoder) error {

				err := in.Decode(&v)
				if err != nil {
					return err
				}

				proto.Merge(&v, reqArgs)
				resp, err := cli.Deposit(context.Background(), &v)

				if err != nil {
					return err
				}

				return out.Encode(resp)

			})
			if err != nil {
				log.Fatal(err)
			}
		},
	}

	cmd.PersistentFlags().StringVar(&reqArgs.Parent, "parent", "", "get-comment-from-proto")
	cmd.PersistentFlags().StringVar(&reqArgs.Tenant, "tenant", "", "get-comment-from-proto")
	cmd.PersistentFlags().StringVar(&reqArgs.Environment, "environment", "", "get-comment-from-proto")
	cmd.PersistentFlags() // Warning: list flags are not yet supported (field "Clusters")
	cmd.PersistentFlags() // Warning: list flags are not yet supported (field "Namespaces")
	cmd.PersistentFlags().StringVar(&reqArgs.Id, "id", "", "get-comment-from-proto")
	cmd.PersistentFlags().StringVar(&reqArgs.Name, "name", "", "get-comment-from-proto")

	return cmd
}

var _BankClientSubCommands = []func() *cobra.Command{
	_BankDepositClientCommand,
}
