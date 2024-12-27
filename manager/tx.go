package manager

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (m *Manager) craftTx(ctx context.Context, data []byte, to common.Address) (*types.Transaction, error) {
	if m.privateKey == nil {
		m.log.Error("finality manager create signer error")
		return nil, errors.New("finality manager create signer error")
	}

	nonce, err := m.ethClient.NonceAt(ctx, m.from, nil)
	if err != nil {
		m.log.Error("failed to get account nonce", "err", err)
		return nil, err
	}

	tip, err := m.ethClient.SuggestGasTipCap(ctx)
	if err != nil {
		m.log.Error(fmt.Errorf("failed to fetch the suggested gas tip cap: %w", err).Error())
		return nil, err
	}

	header, err := m.ethClient.HeaderByNumber(ctx, nil)
	if err != nil {
		m.log.Error(fmt.Errorf("failed to fetch the suggested base fee: %w", err).Error())
		return nil, err
	}
	baseFee := header.BaseFee
	gasFeeCap := calcGasFeeCap(baseFee, tip)

	gasLimit, err := m.ethClient.EstimateGas(ctx, ethereum.CallMsg{
		From: m.from,
		//To:        &m.msmContractAddr,
		GasFeeCap: gasFeeCap,
		GasTipCap: tip,
		Data:      data,
	})

	rawTx := &types.DynamicFeeTx{
		ChainID:   big.NewInt(int64(m.ethChainID)),
		Nonce:     nonce,
		To:        &to,
		Gas:       gasLimit,
		GasTipCap: tip,
		GasFeeCap: gasFeeCap,
		Data:      data,
	}

	tx, err := types.SignNewTx(m.privateKey, types.LatestSignerForChainID(big.NewInt(int64(m.ethChainID))), rawTx)
	if err != nil {
		m.log.Error("failed to sign transaction", "err", err)
		return nil, err
	}

	return tx, nil
}

func calcGasFeeCap(baseFee, gasTipCap *big.Int) *big.Int {
	return new(big.Int).Add(
		gasTipCap,
		new(big.Int).Mul(baseFee, big.NewInt(2)),
	)
}
