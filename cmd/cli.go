package main

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/eniac-x-labs/manta-relayer/common/cliapp"
	"github.com/eniac-x-labs/manta-relayer/config"
	"github.com/eniac-x-labs/manta-relayer/manager"
	"github.com/eniac-x-labs/manta-relayer/node"
	"github.com/eniac-x-labs/manta-relayer/node/conversion"
	"github.com/eniac-x-labs/manta-relayer/sign"
	"github.com/eniac-x-labs/manta-relayer/store"
	"github.com/eniac-x-labs/manta-relayer/ws/server"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/urfave/cli/v2"
	"os"
)

var (
	PrivateKeyFlagName     = "private-key"
	KeyPairFlagName        = "key-pair"
	DefaultPubKeyFilename  = "fn_bls.pub"
	DefaultPrivKeyFilename = "fn_bls.piv"
)

var (
	ConfigFlag = &cli.StringFlag{
		Name:    "config",
		Value:   "./manta-relayer.yaml",
		Aliases: []string{"c"},
		Usage:   "path to config file",
		EnvVars: []string{"MANTA_RELAYER_CONFIG"},
	}
	PrivateKeyFlag = &cli.StringFlag{
		Name:    PrivateKeyFlagName,
		Usage:   "Private Key corresponding to manta relayer",
		EnvVars: []string{"MANTA_RELAYER_PRIVATE_KEY"},
	}
	KeyPairFlag = &cli.StringFlag{
		Name:    KeyPairFlagName,
		Usage:   "key pair corresponding to manta relayer",
		EnvVars: []string{"MANTA_RELAYER_KEY_PAIR"},
	}
)

func newCli(GitCommit string, GitDate string) *cli.App {
	nodeFlags := []cli.Flag{ConfigFlag, PrivateKeyFlag, KeyPairFlag}
	managerFlags := []cli.Flag{ConfigFlag, PrivateKeyFlag}
	peerIDFlags := []cli.Flag{PrivateKeyFlag}
	return &cli.App{
		Version:              params.VersionWithCommit(GitCommit, GitDate),
		Description:          "A decentralized Relayer that synchronizes contract events from Babylon",
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			{
				Name:        "node",
				Flags:       nodeFlags,
				Description: "Runs the relayer node service",
				Action:      cliapp.LifecycleCmd(runNode),
			},
			{
				Name:        "manager",
				Flags:       managerFlags,
				Description: "Runs the relayer manager service",
				Action:      cliapp.LifecycleCmd(runManager),
			},
			{
				Name:        "parse-peer-id",
				Flags:       peerIDFlags,
				Description: "Parse peer id of the key",
				Action:      runParsePeerID,
			},
			{
				Name:        "version",
				Description: "print version",
				Action: func(ctx *cli.Context) error {
					cli.ShowVersion(ctx)
					return nil
				},
			},
		},
	}
}

func runNode(ctx *cli.Context, shutdown context.CancelCauseFunc) (cliapp.Lifecycle, error) {
	logger := log.NewLogger(log.NewTerminalHandlerWithLevel(os.Stdout, log.LevelDebug, true))
	log.SetDefault(logger)

	cfg, err := config.NewConfig(ctx.String(ConfigFlag.Name))
	if err != nil {
		log.Error("failed to load config", "err", err)
		return nil, err
	}

	var privKey *ecdsa.PrivateKey
	if ctx.IsSet(PrivateKeyFlagName) {
		privKey, err = crypto.HexToECDSA(ctx.String(PrivateKeyFlagName))
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("need to config private key")
	}

	var keyPairs *sign.KeyPair
	if ctx.IsSet(KeyPairFlagName) {
		keyPairs, err = sign.MakeKeyPairFromString(ctx.String(KeyPairFlagName))
		if err != nil {
			return nil, err
		}
	} else {
		keyPairs, err = sign.GenRandomBlsKeys()
		if err != nil {
			return nil, err
		}
		pubKeyPath := cfg.Node.KeyPath + "/" + DefaultPubKeyFilename
		privKeyPath := cfg.Node.KeyPath + "/" + DefaultPrivKeyFilename
		err = os.WriteFile(pubKeyPath, []byte(keyPairs.PubKey.String()), 0o600)
		if err != nil {
			return nil, err
		}
		err = os.WriteFile(privKeyPath, []byte((keyPairs.PrivKey.String())), 0o600)
		if err != nil {
			return nil, err
		}
	}

	db, err := store.NewStorage(cfg.Node.LevelDbFolder)
	if err != nil {
		return nil, err
	}

	return node.NewFinalityNode(ctx.Context, db, privKey, keyPairs, cfg, logger, shutdown)
}

func runManager(ctx *cli.Context, shutdown context.CancelCauseFunc) (cliapp.Lifecycle, error) {
	logger := log.NewLogger(log.NewTerminalHandlerWithLevel(os.Stdout, log.LevelDebug, true))
	log.SetDefault(logger)

	cfg, err := config.NewConfig(ctx.String(ConfigFlag.Name))
	if err != nil {
		log.Error("failed to load config", "err", err)
		return nil, err
	}

	var privKey *ecdsa.PrivateKey
	if ctx.IsSet(PrivateKeyFlagName) {
		privKey, err = crypto.HexToECDSA(ctx.String(PrivateKeyFlagName))
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("need to config private key")
	}

	db, err := store.NewStorage(cfg.Manager.LevelDbFolder)
	if err != nil {
		return nil, err
	}

	wsServer, err := server.NewWSServer(cfg.Manager.WsAddr)
	if err != nil {
		return nil, err
	}

	return manager.NewFinalityManager(ctx.Context, db, wsServer, cfg, shutdown, logger, privKey)
}

func runParsePeerID(ctx *cli.Context) error {
	privateKey := ctx.String(PrivateKeyFlag.Name)

	var publicBz []byte
	if len(privateKey) != 0 {
		privKey, err := crypto.HexToECDSA(privateKey)
		if err != nil {
			return err
		}
		pubkeybytes := crypto.CompressPubkey(&privKey.PublicKey)
		publicBz = pubkeybytes
	} else {
		return errors.New("pri-key needs to be specified")
	}

	peerId, err := conversion.GetPeerIDFromSecp256PubKey(publicBz)
	if err != nil {
		return err
	}
	fmt.Println(peerId)
	return nil
}
