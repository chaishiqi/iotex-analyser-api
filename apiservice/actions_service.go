package apiservice

import (
	"context"

	"github.com/iotexproject/iotex-address/address"
	"github.com/iotexproject/iotex-analyser-api/api"
	"github.com/iotexproject/iotex-analyser-api/db"
)

type XrcType int

const (
	Xrc20 XrcType = iota
	Xrc721
)

type ActionsService struct {
	api.UnimplementedActionsServiceServer
}

//curl -d '{"address": "io102s4660k3cynae2r8gde6scg74mf6f7k9dq955", "height":11900487 }' http://127.0.0.1:7778/api.ActionsService.GetActionsByAddress
func (s *ActionsService) GetActionsByAddress(ctx context.Context, req *api.ActionsRequest) (*api.ActionsByAddressResponse, error) {
	resp := &api.ActionsByAddressResponse{
		Count: 0,
	}
	db := db.DB()
	addr := req.GetAddress()
	if addr[:2] == "0x" || addr[:2] == "0X" {
		add, err := address.FromHex(addr)
		if err != nil {
			return nil, err
		}

		addr = add.String()
	}
	offset := req.GetOffset()
	size := req.GetSize()
	if size == 0 {
		size = 25
	}
	sort := req.GetSort()
	if sort != "asc" && sort != "desc" {
		sort = "asc"
	}

	var count int64
	err := db.Table("block_action a").Where("a.from=? or a.to=?", addr, addr).Count(&count).Error

	if err != nil {
		return nil, err
	}
	resp.Count = uint64(count)

	query := "SELECT a.action_hash,a.action_type,a.block_height,a.from,a.to,a.gas_price*r.gas_consumed,a.gas_limit,a.nonce,a.amount,r.status,b.block_hash,b.timestamp FROM block_action a inner join block b on b.block_height=a.block_height inner join block_receipt r on r.action_hash=a.action_hash where a.from=? or a.to=? order by a.id " + sort + " limit ? offset ?"
	rows, err := db.Raw(query, addr, addr, size, offset).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var actHash, actType, from, to, blkHash, gasFee, amount string
	var blkHeight, gasLimit, nonce, status, timestamp uint64
	for rows.Next() {
		if err := rows.Scan(&actHash, &actType, &blkHeight, &from, &to, &gasFee, &gasLimit, &nonce, &amount, &status, &blkHash, &timestamp); err != nil {
			return nil, err
		}
		resp.Results = append(resp.Results, &api.ActionsByAddressResult{
			ActHash:   actHash,
			ActType:   actType,
			TimeStamp: timestamp,
			BlkHash:   blkHash,
			Sender:    from,
			Recipient: to,
			Amount:    amount,
			GasFee:    gasFee,
		})
	}
	return resp, nil
}

func (s *ActionsService) GetXrc20ByAddress(ctx context.Context, req *api.ActionsRequest) (*api.Xrc20ByAddressResponse, error) {
	return s.getXrcByAddress(Xrc20, req)
}

func (s *ActionsService) GetXrc721ByAddress(ctx context.Context, req *api.ActionsRequest) (*api.Xrc20ByAddressResponse, error) {
	return s.getXrcByAddress(Xrc721, req)
}

func (s *ActionsService) getXrcByAddress(xrcType XrcType, req *api.ActionsRequest) (*api.Xrc20ByAddressResponse, error) {
	var xrcTable string
	switch xrcType {
	case Xrc20:
		xrcTable = "token_erc20 a"
	case Xrc721:
		xrcTable = "token_erc721 a"
	}
	resp := &api.Xrc20ByAddressResponse{
		Count: 0,
	}
	db := db.DB()
	addr := req.GetAddress()
	if addr[:2] == "0x" || addr[:2] == "0X" {
		add, err := address.FromHex(addr)
		if err != nil {
			return nil, err
		}

		addr = add.String()
	}
	offset := req.GetOffset()
	size := req.GetSize()
	if size == 0 {
		size = 25
	}
	sort := req.GetSort()
	if sort != "asc" && sort != "desc" {
		sort = "asc"
	}

	var count int64
	err := db.Table(xrcTable).Where("a.from=? or a.to=?", addr, addr).Count(&count).Error

	if err != nil {
		return nil, err
	}
	resp.Count = uint64(count)

	query := "select t.*,(select timestamp from block where block_height=t.block_height) from (select a.block_height,a.action_hash,a.contract_address,a.amount,a.from,a.to from " + xrcTable + " where a.from=? or a.to=? order by a.id " + sort + " limit ? offset ?) t"
	rows, err := db.Raw(query, addr, addr, size, offset).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var actHash, contractAddress, from, to, amount string
	var blkHeight, timestamp uint64
	for rows.Next() {
		if err := rows.Scan(&blkHeight, &actHash, &contractAddress, &amount, &from, &to, &timestamp); err != nil {
			return nil, err
		}
		resp.Results = append(resp.Results, &api.Xrc20ByAddressResult{
			ActHash:         actHash,
			BlkHeight:       blkHeight,
			TimeStamp:       timestamp,
			ContractAddress: contractAddress,
			From:            from,
			To:              to,
			Amount:          amount,
		})
	}
	return resp, nil
}

func (s *ActionsService) GetEvmTransferDetailListByAddress(ctx context.Context, req *api.ActionsRequest) (*api.EvmTransferDetailListByAddressResponse, error) {
	resp := &api.EvmTransferDetailListByAddressResponse{
		Count: 0,
	}
	db := db.DB()
	addr := req.GetAddress()
	if addr[:2] == "0x" || addr[:2] == "0X" {
		add, err := address.FromHex(addr)
		if err != nil {
			return nil, err
		}

		addr = add.String()
	}
	offset := req.GetOffset()
	size := req.GetSize()
	if size == 0 {
		size = 25
	}
	sort := req.GetSort()
	if sort != "asc" && sort != "desc" {
		sort = "asc"
	}

	var count int64
	err := db.Table("block_receipt_transaction a").Where("a.sender=? or a.recipient=?", addr, addr).Count(&count).Error

	if err != nil {
		return nil, err
	}
	resp.Count = uint64(count)

	query := "SELECT a.action_hash,a.block_height,a.sender,a.recipient,a.amount,b.block_hash,b.timestamp FROM block_receipt_transaction a inner join block b on b.block_height=a.block_height where a.sender=? or a.recipient=? order by a.id " + sort + " limit ? offset ?"
	rows, err := db.Raw(query, addr, addr, size, offset).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var actHash, sender, recipient, blkHash, amount string
	var blkHeight, timestamp uint64
	for rows.Next() {
		if err := rows.Scan(&actHash, &blkHeight, &sender, &recipient, &amount, &blkHash, &timestamp); err != nil {
			return nil, err
		}
		resp.Results = append(resp.Results, &api.EvmTransferDetailResult{
			ActHash:   actHash,
			TimeStamp: timestamp,
			BlkHeight: blkHeight,
			BlkHash:   blkHash,
			Sender:    sender,
			Recipient: recipient,
			Amount:    amount,
		})
	}
	return resp, nil
}

/*
select t.*,(select action_type from block_action where action_hash=t.action_hash),(select timestamp from block where block_height=t.block_height) from ((select a.block_height,a.action_hash,a.amount,a.sender,a.recipient, 'xrc20' as rtype from token_erc20 a where a.sender='io15qr5fzpxsnp7garl4m7k355rafzqn8grrm0grz' or a.recipient='io15qr5fzpxsnp7garl4m7k355rafzqn8grrm0grz') union
        (SELECT a.block_height,a.action_hash,a.amount,a.sender,a.recipient, 'native' as rtype FROM block_action a where a.sender='io15qr5fzpxsnp7garl4m7k355rafzqn8grrm0grz' or a.recipient='io15qr5fzpxsnp7garl4m7k355rafzqn8grrm0grz') union
        (SELECT a.block_height,a.action_hash,a.amount,a.sender,a.recipient,'evmtransfer' as rtype FROM block_receipt_transaction a where a.sender='io15qr5fzpxsnp7garl4m7k355rafzqn8grrm0grz' or a.recipient='io15qr5fzpxsnp7garl4m7k355rafzqn8grrm0grz'))t order by block_height asc limit 25 offset 0;

select t.* from ((select a.block_height,a.action_hash,a.amount,a.sender,a.recipient, 'xrc20' as rtype from token_erc20 a where a.sender='io1hp6y4eqr90j7tmul4w2wa8pm7wx462hq0mg4tw' or a.recipient='io1hp6y4eqr90j7tmul4w2wa8pm7wx462hq0mg4tw' order by block_height desc) union all
        (SELECT a.block_height,a.action_hash,a.amount,a.sender,a.recipient, 'native' as rtype FROM block_action a where a.sender='io1hp6y4eqr90j7tmul4w2wa8pm7wx462hq0mg4tw' or a.recipient='io1hp6y4eqr90j7tmul4w2wa8pm7wx462hq0mg4tw' order by block_height desc) union all
        (SELECT a.block_height,a.action_hash,a.amount,a.sender,a.recipient,'evmtransfer' as rtype FROM block_receipt_transaction a where a.sender='io1hp6y4eqr90j7tmul4w2wa8pm7wx462hq0mg4tw' or a.recipient='io1hp6y4eqr90j7tmul4w2wa8pm7wx462hq0mg4tw' order by block_height desc))t  order by block_height desc limit 25 offset 30;
*/
func (s *ActionsService) GetAllActionsByAddress(ctx context.Context, req *api.ActionsRequest) (*api.AllActionsByAddressResponse, error) {
	resp := &api.AllActionsByAddressResponse{
		Count: 0,
	}
	db := db.DB()
	addr := req.GetAddress()
	if addr[:2] == "0x" || addr[:2] == "0X" {
		add, err := address.FromHex(addr)
		if err != nil {
			return nil, err
		}

		addr = add.String()
	}
	offset := req.GetOffset()
	size := req.GetSize()
	if size == 0 {
		size = 25
	}
	sort := req.GetSort()
	if sort != "asc" && sort != "desc" {
		sort = "asc"
	}

	query := `(select a.block_height,a.action_hash,a.amount,a.sender,a.recipient, 'xrc20' as rtype from token_erc20 a where a.sender=? or a.recipient=?) union all (SELECT a.block_height,a.action_hash,a.amount,a.sender,a.recipient, 'native' as rtype FROM block_action a where a.sender=? or a.recipient=?) union all (SELECT a.block_height,a.action_hash,a.amount,a.sender,a.recipient,'evmtransfer' as rtype FROM block_receipt_transaction a where a.sender=? or a.recipient=?)`
	coutQuery := `select count(*) from(` + query + `)c`
	var count int64
	err := db.Raw(coutQuery, addr, addr, addr, addr, addr, addr).Scan(&count).Error

	if err != nil {
		return nil, err
	}
	resp.Count = uint64(count)
	resQuery := `select t.*,(select action_type from block_action where action_hash=t.action_hash),(select timestamp from block where block_height=t.block_height) from (` + query + `)t order by block_height ` + sort + ` limit ? offset ?`
	rows, err := db.Raw(resQuery, addr, addr, addr, addr, addr, addr, size, offset).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var actHash, sender, recipient, amount, actType, rtype string
	var blkHeight, timestamp uint64
	for rows.Next() {
		if err := rows.Scan(&blkHeight, &actHash, &amount, &sender, &recipient, &rtype, &actType, &timestamp); err != nil {
			return nil, err
		}
		var rType api.AllActionsByAddressResult_RecordType
		switch rtype {
		case "native":
			rType = api.AllActionsByAddressResult_NATIVE
		case "xrc20":
			rType = api.AllActionsByAddressResult_XRC20
		case "evmtransfer":
			rType = api.AllActionsByAddressResult_EVMTRANSFER
		}
		resp.Results = append(resp.Results, &api.AllActionsByAddressResult{
			ActHash:    actHash,
			TimeStamp:  timestamp,
			ActType:    actType,
			BlkHeight:  blkHeight,
			Sender:     sender,
			Recipient:  recipient,
			Amount:     amount,
			RecordType: rType,
		})
	}
	return resp, nil
}
