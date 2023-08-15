package transaction

import (
	"math/big"
	"strconv"
	"time"

	"github.com/LMF709268224/titan-vps/api/types"
	"github.com/LMF709268224/titan-vps/lib/trxbridge"
	"github.com/LMF709268224/titan-vps/lib/trxbridge/api"
	"github.com/LMF709268224/titan-vps/lib/trxbridge/core"
	"github.com/LMF709268224/titan-vps/lib/trxbridge/hexutil"
	"github.com/LMF709268224/titan-vps/node/db"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smirkcat/hdwallet"
	"golang.org/x/xerrors"
	"google.golang.org/protobuf/proto"
)

const checkBlockInterval = 3 * time.Second

// GetGrpcClient
func (m *Manager) getGrpcClient() (*trxbridge.GrpcClient, error) {
	node := trxbridge.NewGrpcClient(m.cfg.TrxHTTPSAddr)
	err := node.Start()
	if err != nil {
		return nil, err
	}

	return node, nil
}

func (m *Manager) watchTronTransactions() {
	ticker := time.NewTicker(checkBlockInterval)
	defer ticker.Stop()

	client, err := m.getGrpcClient()
	if err != nil {
		log.Errorln("getGrpcClient err :", err.Error())
		return
	}

	height := int64(39297600)
	limit := int64(50)
	heightStr := ""

	err = m.LoadConfigValue(db.ConfigTronHeight, &heightStr)
	if err == nil {
		i, err := strconv.ParseInt(heightStr, 10, 64)
		if err == nil {
			height = i
		}
	}

	for {
		<-ticker.C

		block, err := client.GetNowBlock()
		if err != nil {
			log.Errorf("GetNowBlock err:%s", err.Error())
			continue
		}

		nowHeight := block.BlockHeader.RawData.Number
		endHeight := height + limit
		if endHeight >= nowHeight {
			endHeight = nowHeight
		}
		if height >= endHeight {
			continue
		}

		log.Debugf(" handleBlock height :%d, endHeight:%d \n", height, endHeight)
		blocks, err := client.GetBlockByLimitNext(height, endHeight)
		if err == nil && len(blocks.Block) > 0 {
			m.handleBlocks(blocks)
		}

		height = endHeight
		str := strconv.FormatInt(height, 10)
		err = m.SaveConfigValue(db.ConfigTronHeight, str)
		if err != nil {
			log.Errorf("SaveConfigValue err:%s", err.Error())
		}
	}
}

func (m *Manager) handleBlocks(blocks *api.BlockListExtention) {
	for _, v := range blocks.Block {
		err := m.handleBlock(v)
		if err != nil {
			log.Errorln(" handleBlock err :", err.Error())
		}
	}
}

func (m *Manager) handleBlock(blockExtention *api.BlockExtention) error {
	if blockExtention == nil || blockExtention.BlockHeader == nil {
		return xerrors.New("block is nil")
	}

	bNum := blockExtention.BlockHeader.RawData.Number
	// log.Debugln(" handleBlock height :", bNum)

	bid := hexutil.Encode(blockExtention.Blockid)

	for _, te := range blockExtention.Transactions {
		if len(te.Transaction.GetRet()) == 0 {
			continue
		}

		state := te.Transaction.GetRet()[0].ContractRet
		txid := hexutil.Encode(te.Txid)

		// userAddr := string(te.Transaction.RawData.Data)

		for _, contract := range te.Transaction.RawData.Contract {
			m.filterTransaction(contract, txid, bid, bNum, state)
		}
	}

	return nil
}

func (m *Manager) filterTransaction(contract *core.Transaction_Contract, txid, bid string, bNum int64, state core.Transaction_ResultContractResult) {
	if contract.Type == core.Transaction_Contract_TriggerSmartContract {
		// trc20
		unObj := &core.TriggerSmartContract{}
		err := proto.Unmarshal(contract.Parameter.GetValue(), unObj)
		if err != nil {
			// log.Errorf("parse trc20 err: %s", err.Error())
			return
		}

		contractAddress := hdwallet.EncodeCheck(unObj.GetContractAddress())

		if contractAddress != m.cfg.TrxContractorAddr {
			// log.Errorf("contractAddress: %s", contractAddress)
			return
		}

		from := hdwallet.EncodeCheck(unObj.GetOwnerAddress())
		data := unObj.GetData()

		to, amount, isOk := m.decodeData(data)
		if !isOk {
			// log.Errorf("decodeData err: %s", txid)
			return
		}

		m.handleTransfer(txid, from, to, bid, bNum, amount, contractAddress, state)
	}
}

func (m *Manager) decodeData(trc20 []byte) (to string, amount string, flag bool) {
	if len(trc20) >= 68 {
		if hexutil.Encode(trc20[:4]) != "a9059cbb" {
			return
		}
		trc20[15] = 65 // 0x41

		bb := common.TrimLeftZeroes(trc20[36:68])
		bu := new(big.Int).SetBytes(bb)
		amount = bu.String()

		to = hdwallet.EncodeCheck(trc20[15:36])
		flag = true
	}
	return
}

func (m *Manager) handleTransfer(mCid, from, to, blockCid string, height int64, amount string, contract string, state core.Transaction_ResultContractResult) {
	log.Infof("Transfer :%s,%s,%s,%s,%s,%s", mCid, to, from, contract, amount, state)

	if userAddr, ok := m.tronAddrs[to]; ok {
		m.notify.Pub(&types.TronTransferWatch{
			TxHash:   mCid,
			From:     from,
			To:       to,
			Value:    amount,
			State:    state,
			Height:   height,
			UserAddr: userAddr,
		}, types.EventTronTransferWatch.String())
	}
}
