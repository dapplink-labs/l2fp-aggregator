package synchronizer

import (
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/dapplink-labs/l2fp-aggregator/common"
	"github.com/dapplink-labs/l2fp-aggregator/store"

	types2 "github.com/babylonlabs-io/babylon/x/btcstaking/types"
	"github.com/babylonlabs-io/babylon/x/finality/types"
)

func (syncer *BabylonSynchronizer) ProcessNewFinalityProvider(txMessage store.TxMessage) error {
	var err error
	var mCFP types2.MsgCreateFinalityProvider

	if txMessage.Type == common.MsgCreateFinalityProvider {
		err = mCFP.Unmarshal(txMessage.Data)
		if err != nil {
			syncer.log.Error("failed to unmarshal NewFinalityProvider message value", "err", err)
			return err
		}
		err = syncer.db.SetCreateFinalityProviderMsg(store.CreateFinalityProvider{
			FP:     mCFP,
			TxHash: txMessage.TransactionHash,
		})
		if err != nil {
			syncer.log.Error("failed to store NewFinalityProvider message", "err", err)
			return err
		}
		syncer.log.Info("success to store NewFinalityProvider message", "tx_hash", hexutil.Encode(txMessage.TransactionHash))
	}

	return nil
}

func (syncer *BabylonSynchronizer) ProcessCreateBTCDelegation(txMessage store.TxMessage) error {
	var err error
	var mCBD types2.MsgCreateBTCDelegation

	if txMessage.Type == common.MsgCreateBTCDelegation {
		mCBD.Unmarshal(txMessage.Data)
		err = syncer.db.SetCreateBTCDelegationMsg(store.CreateBTCDelegation{
			CBD:    mCBD,
			TxHash: txMessage.TransactionHash,
		})
		if err != nil {
			syncer.log.Error("failed to store CreateBTCDelegation message", "err", err)
			return err
		}
		syncer.log.Info("success to store CreateBTCDelegation message", "tx_hash", hexutil.Encode(txMessage.TransactionHash))
	}
	return nil
}

func (syncer *BabylonSynchronizer) ProcessCommitPubRandList(txMessage store.TxMessage) error {
	var err error
	var mCPR types.MsgCommitPubRandList

	if txMessage.Type == common.MsgCommitPubRandList {
		mCPR.Unmarshal(txMessage.Data)
		err = syncer.db.SetCommitPubRandListMsg(store.CommitPubRandList{
			CPR:    mCPR,
			TxHash: txMessage.TransactionHash,
		})
		if err != nil {
			syncer.log.Error("failed to store CommitPubRandList message", "err", err)
			return err
		}
		syncer.log.Info("success to store CommitPubRandList message", "tx_hash", hexutil.Encode(txMessage.TransactionHash))
	}

	return nil
}
