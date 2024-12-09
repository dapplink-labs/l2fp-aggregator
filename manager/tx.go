package manager

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	types2 "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"time"
)

func (m *Manager) craftTx(ctx context.Context, data []byte, to common.Address) (*types2.Transaction, error) {
	if m.privateKey == nil {
		m.log.Error("finality manager create signer error")
		return nil, errors.New("finality manager create signer error")
	}

	nonce, err := m.l1Client.NonceAt(ctx, m.from, nil)
	if err != nil {
		m.log.Error("failed to get account nonce", "err", err)
		return nil, err
	}

	tip, err := m.l1Client.SuggestGasTipCap(ctx)
	if err != nil {
		m.log.Error(fmt.Errorf("failed to fetch the suggested gas tip cap: %w", err).Error())
		return nil, err
	}

	header, err := m.l1Client.HeaderByNumber(ctx, nil)
	if err != nil {
		m.log.Error(fmt.Errorf("failed to fetch the suggested base fee: %w", err).Error())
		return nil, err
	}
	baseFee := header.BaseFee
	gasFeeCap := calcGasFeeCap(baseFee, tip)

	gasLimit, err := m.l1Client.EstimateGas(ctx, ethereum.CallMsg{
		From: m.from,
		//To:        &m.msmContractAddr,
		GasFeeCap: gasFeeCap,
		GasTipCap: tip,
		Data:      data,
	})

	rawTx := &types2.DynamicFeeTx{
		ChainID:   big.NewInt(int64(m.l1ChainID)),
		Nonce:     nonce,
		To:        &to,
		Gas:       gasLimit,
		GasTipCap: tip,
		GasFeeCap: gasFeeCap,
		Data:      data,
	}

	tx, err := types2.SignNewTx(m.privateKey, types2.LatestSignerForChainID(big.NewInt(int64(m.l1ChainID))), rawTx)
	if err != nil {
		m.log.Error("failed to sign transaction", "err", err)
		return nil, err
	}

	return tx, nil
}

func getTransactionReceipt(ctx context.Context, client *ethclient.Client, txHash common.Hash) (*types2.Receipt, error) {
	var receipt *types2.Receipt
	var err error

	ticker := time.NewTicker(10 * time.Second)
	for {
		<-ticker.C
		receipt, err = client.TransactionReceipt(ctx, txHash)
		if err != nil && !errors.Is(err, ethereum.NotFound) {
			return nil, err
		}

		if errors.Is(err, ethereum.NotFound) {
			continue
		}
		return receipt, nil
	}
}

func calcGasFeeCap(baseFee, gasTipCap *big.Int) *big.Int {
	return new(big.Int).Add(
		gasTipCap,
		new(big.Int).Mul(baseFee, big.NewInt(2)),
	)
}
