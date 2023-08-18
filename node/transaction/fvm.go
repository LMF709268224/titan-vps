package transaction

import (
	"time"

	"github.com/LMF709268224/titan-vps/api/types"
	"github.com/LMF709268224/titan-vps/lib/filecoinbridge"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/xerrors"
)

func (m *Manager) watchFvmTransactions() error {
	client, err := ethclient.Dial(m.cfg.LotusWsAddr)
	if err != nil {
		return xerrors.Errorf("Dial err:%s", err.Error())
	}

	tokenAddress := common.HexToAddress(m.cfg.TitanContractorAddr)

	fAbi, err := filecoinbridge.NewAbi(tokenAddress, client)
	if err != nil {
		return xerrors.Errorf("NewAbi err:%s", err.Error())
	}

	sink := make(chan *filecoinbridge.AbiTransfer)

	sub, err := fAbi.WatchTransfer(nil, sink, nil, nil)
	if err != nil {
		return xerrors.Errorf("WatchTransfer err:%s", err.Error())
	}

	for {
		select {
		case err := <-sub.Err():
			if err != nil {
				log.Debugln(time.Now().Format("2006-01-02 15:04:05"), " err:", err)
				sub, err = fAbi.WatchTransfer(nil, sink, nil, nil)
				if err != nil {
					return xerrors.Errorf("WatchTransfer err:%s", err.Error())
				}
			}
		case tr := <-sink:
			log.Debugf("from:%s,to:%s,value:%d, RawTxHash:%s,RawBlockNumber:%d, Removed:%v \n", tr.From.String(), tr.To.Hex(), tr.Value, tr.Raw.TxHash.String(), tr.Raw.BlockNumber, tr.Raw.Removed)
			if !tr.Raw.Removed {
				m.notification.Pub(&types.FvmTransferWatch{
					TxHash: tr.Raw.TxHash.Hex(),
					From:   tr.From.Hex(),
					To:     tr.To.Hex(),
					Value:  tr.Value.String(),
				}, types.EventFvmTransferWatch.String())
			}
		}
	}
}
