package packet

import (
	"github.com/RadiatedMonkey/gophertunnel/minecraft/protocol"
)

const (
	InventoryTransactionTypeNormal = iota
	InventoryTransactionTypeMismatch
	InventoryTransactionTypeUseItem
	InventoryTransactionTypeUseItemOnEntity
	InventoryTransactionTypeReleaseItem
)

// InventoryTransaction is a packet sent by the client. It essentially exists out of multiple sub-packets,
// each of which have something to do with the inventory in one way or another. Some of these sub-packets
// directly relate to the inventory, others relate to interaction with the world, that could potentially
// result in a change in the inventory.
type InventoryTransaction struct {
	// LegacyRequestID is an ID that is only non-zero at times when sent by the client. The server should
	// always send 0 for this. When this field is not 0, the LegacySetItemSlots slice below will have values
	// in it.
	// LegacyRequestID ties in with the ItemStackResponse packet. If this field is non-0, the server should
	// respond with an ItemStackResponse packet. Some inventory actions such as dropping an item out of the
	// hotbar are still one using this packet, and the ItemStackResponse packet needs to tie in with it.
	LegacyRequestID int32
	// LegacySetItemSlots are only present if the LegacyRequestID is non-zero. These item slots inform the
	// server of the slots that were changed during the inventory transaction, and the server should send
	// back an ItemStackResponse packet with these slots present in it. (Or false with no slots, if rejected.)
	LegacySetItemSlots []protocol.LegacySetItemSlot
	// Actions is a list of actions that took place, that form the inventory transaction together. Each of
	// these actions hold one slot in which one item was changed to another. In general, the combination of
	// all of these actions results in a balanced inventory transaction. This should be checked to ensure that
	// no items are cheated into the inventory.
	Actions []protocol.InventoryAction
	// TransactionData is a data object that holds data specific to the type of transaction that the
	// TransactionPacket held. Its concrete type must be one of NormalTransactionData, MismatchTransactionData
	// UseItemTransactionData, UseItemOnEntityTransactionData or ReleaseItemTransactionData. If nil is set,
	// the transaction will be assumed to of type InventoryTransactionTypeNormal.
	TransactionData protocol.InventoryTransactionData
}

// ID ...
func (*InventoryTransaction) ID() uint32 {
	return IDInventoryTransaction
}

// Marshal ...
func (pk *InventoryTransaction) Marshal(w *protocol.Writer) {
	w.Varint32(&pk.LegacyRequestID)
	if pk.LegacyRequestID != 0 {
		protocol.Slice(w, &pk.LegacySetItemSlots)
	}
	var id uint32
	switch pk.TransactionData.(type) {
	case nil, *protocol.NormalTransactionData:
		id = InventoryTransactionTypeNormal
	case *protocol.MismatchTransactionData:
		id = InventoryTransactionTypeMismatch
	case *protocol.UseItemTransactionData:
		id = InventoryTransactionTypeUseItem
	case *protocol.UseItemOnEntityTransactionData:
		id = InventoryTransactionTypeUseItemOnEntity
	case *protocol.ReleaseItemTransactionData:
		id = InventoryTransactionTypeReleaseItem
	}
	w.Varuint32(&id)
	protocol.Slice(w, &pk.Actions)
	if pk.TransactionData != nil {
		pk.TransactionData.Marshal(w)
	}
}

// Unmarshal ...
func (pk *InventoryTransaction) Unmarshal(r *protocol.Reader) {
	r.Varint32(&pk.LegacyRequestID)
	if pk.LegacyRequestID != 0 {
		protocol.Slice(r, &pk.LegacySetItemSlots)
	}
	var transactionType uint32
	r.Varuint32(&transactionType)
	protocol.Slice(r, &pk.Actions)
	switch transactionType {
	case InventoryTransactionTypeNormal:
		pk.TransactionData = &protocol.NormalTransactionData{}
	case InventoryTransactionTypeMismatch:
		pk.TransactionData = &protocol.MismatchTransactionData{}
	case InventoryTransactionTypeUseItem:
		pk.TransactionData = &protocol.UseItemTransactionData{}
	case InventoryTransactionTypeUseItemOnEntity:
		pk.TransactionData = &protocol.UseItemOnEntityTransactionData{}
	case InventoryTransactionTypeReleaseItem:
		pk.TransactionData = &protocol.ReleaseItemTransactionData{}
	default:
		r.UnknownEnumOption(transactionType, "inventory transaction type")
	}
	pk.TransactionData.Marshal(r)
}
