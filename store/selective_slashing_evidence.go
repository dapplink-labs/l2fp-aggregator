package store

import (
	"encoding/json"

	"github.com/babylonlabs-io/babylon/x/btcstaking/types"
)

type SelectiveSlashingEvidence struct {
	SSE    types.MsgSelectiveSlashingEvidence
	TxHash []byte `json:"tx_hash"`
}

func (s *Storage) SetSelectiveSlashingEvidenceMsg(msg SelectiveSlashingEvidence) error {
	bz, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	//todo check the byte is right
	exist, cBD := s.GetCreateBTCDelegationMsg([]byte(msg.SSE.StakingTxHash))
	if exist {
		err = s.setBTCUnDelegateAmount([]byte(cBD.CBD.StakerAddr), uint64(cBD.CBD.StakingValue))
		if err != nil {
			return err
		}
	}

	return s.db.Put(getSelectiveSlashingEvidenceKey(msg.TxHash), bz, nil)
}

func (s *Storage) GetSelectiveSlashingEvidenceMsg(txHash []byte) (bool, SelectiveSlashingEvidence) {
	sSEB, err := s.db.Get(getSelectiveSlashingEvidenceKey(txHash), nil)
	if err != nil {
		return handleError2(SelectiveSlashingEvidence{}, err)
	}
	var sSE SelectiveSlashingEvidence
	if err = json.Unmarshal(sSEB, &sSE); err != nil {
		return false, SelectiveSlashingEvidence{}
	}
	return true, sSE
}
