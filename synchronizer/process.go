package synchronizer

import (
	types2 "github.com/babylonlabs-io/babylon/x/btcstaking/types"
	"github.com/babylonlabs-io/babylon/x/finality/types"
	"github.com/dapplink-labs/babylonfp-relayer-ethl2/common"
	"github.com/dapplink-labs/babylonfp-relayer-ethl2/store"
)

func (syncer *Synchronizer) ProcessNewFinalityProvider(txMessage store.TxMessage) error {
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
	}

	return nil
}

func (syncer *Synchronizer) ProcessCreateBTCDelegation(txMessage store.TxMessage) error {
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
	}
	return nil
}

func (syncer *Synchronizer) ProcessCommitPubRandList(txMessage store.TxMessage) error {
	var err error
	var mCPR types.MsgCommitPubRandList

	if txMessage.Type == common.MsgCreateBTCDelegation {
		mCPR.Unmarshal(txMessage.Data)
		err = syncer.db.SetCommitPubRandListMsg(store.CommitPubRandList{
			CPR:    mCPR,
			TxHash: txMessage.TransactionHash,
		})
		if err != nil {
			syncer.log.Error("failed to store CommitPubRandList message", "err", err)
			return err
		}
	}

	return nil
}
