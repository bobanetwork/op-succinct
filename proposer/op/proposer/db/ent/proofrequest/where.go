// Code generated by ent, DO NOT EDIT.

package proofrequest

import (
	"entgo.io/ent/dialect/sql"
	"github.com/succinctlabs/op-succinct-go/proposer/db/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldLTE(FieldID, id))
}

// StartBlock applies equality check predicate on the "start_block" field. It's identical to StartBlockEQ.
func StartBlock(v uint64) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldEQ(FieldStartBlock, v))
}

// EndBlock applies equality check predicate on the "end_block" field. It's identical to EndBlockEQ.
func EndBlock(v uint64) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldEQ(FieldEndBlock, v))
}

// RequestAddedTime applies equality check predicate on the "request_added_time" field. It's identical to RequestAddedTimeEQ.
func RequestAddedTime(v uint64) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldEQ(FieldRequestAddedTime, v))
}

// ProverRequestID applies equality check predicate on the "prover_request_id" field. It's identical to ProverRequestIDEQ.
func ProverRequestID(v string) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldEQ(FieldProverRequestID, v))
}

// ProofRequestTime applies equality check predicate on the "proof_request_time" field. It's identical to ProofRequestTimeEQ.
func ProofRequestTime(v uint64) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldEQ(FieldProofRequestTime, v))
}

// L1BlockNumber applies equality check predicate on the "l1_block_number" field. It's identical to L1BlockNumberEQ.
func L1BlockNumber(v uint64) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldEQ(FieldL1BlockNumber, v))
}

// L1BlockHash applies equality check predicate on the "l1_block_hash" field. It's identical to L1BlockHashEQ.
func L1BlockHash(v string) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldEQ(FieldL1BlockHash, v))
}

// Proof applies equality check predicate on the "proof" field. It's identical to ProofEQ.
func Proof(v []byte) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldEQ(FieldProof, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v Type) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v Type) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...Type) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...Type) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldNotIn(FieldType, vs...))
}

// StartBlockEQ applies the EQ predicate on the "start_block" field.
func StartBlockEQ(v uint64) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldEQ(FieldStartBlock, v))
}

// StartBlockNEQ applies the NEQ predicate on the "start_block" field.
func StartBlockNEQ(v uint64) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldNEQ(FieldStartBlock, v))
}

// StartBlockIn applies the In predicate on the "start_block" field.
func StartBlockIn(vs ...uint64) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldIn(FieldStartBlock, vs...))
}

// StartBlockNotIn applies the NotIn predicate on the "start_block" field.
func StartBlockNotIn(vs ...uint64) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldNotIn(FieldStartBlock, vs...))
}

// StartBlockGT applies the GT predicate on the "start_block" field.
func StartBlockGT(v uint64) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldGT(FieldStartBlock, v))
}

// StartBlockGTE applies the GTE predicate on the "start_block" field.
func StartBlockGTE(v uint64) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldGTE(FieldStartBlock, v))
}

// StartBlockLT applies the LT predicate on the "start_block" field.
func StartBlockLT(v uint64) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldLT(FieldStartBlock, v))
}

// StartBlockLTE applies the LTE predicate on the "start_block" field.
func StartBlockLTE(v uint64) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldLTE(FieldStartBlock, v))
}

// EndBlockEQ applies the EQ predicate on the "end_block" field.
func EndBlockEQ(v uint64) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldEQ(FieldEndBlock, v))
}

// EndBlockNEQ applies the NEQ predicate on the "end_block" field.
func EndBlockNEQ(v uint64) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldNEQ(FieldEndBlock, v))
}

// EndBlockIn applies the In predicate on the "end_block" field.
func EndBlockIn(vs ...uint64) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldIn(FieldEndBlock, vs...))
}

// EndBlockNotIn applies the NotIn predicate on the "end_block" field.
func EndBlockNotIn(vs ...uint64) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldNotIn(FieldEndBlock, vs...))
}

// EndBlockGT applies the GT predicate on the "end_block" field.
func EndBlockGT(v uint64) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldGT(FieldEndBlock, v))
}

// EndBlockGTE applies the GTE predicate on the "end_block" field.
func EndBlockGTE(v uint64) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldGTE(FieldEndBlock, v))
}

// EndBlockLT applies the LT predicate on the "end_block" field.
func EndBlockLT(v uint64) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldLT(FieldEndBlock, v))
}

// EndBlockLTE applies the LTE predicate on the "end_block" field.
func EndBlockLTE(v uint64) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldLTE(FieldEndBlock, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldNotIn(FieldStatus, vs...))
}

// RequestAddedTimeEQ applies the EQ predicate on the "request_added_time" field.
func RequestAddedTimeEQ(v uint64) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldEQ(FieldRequestAddedTime, v))
}

// RequestAddedTimeNEQ applies the NEQ predicate on the "request_added_time" field.
func RequestAddedTimeNEQ(v uint64) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldNEQ(FieldRequestAddedTime, v))
}

// RequestAddedTimeIn applies the In predicate on the "request_added_time" field.
func RequestAddedTimeIn(vs ...uint64) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldIn(FieldRequestAddedTime, vs...))
}

// RequestAddedTimeNotIn applies the NotIn predicate on the "request_added_time" field.
func RequestAddedTimeNotIn(vs ...uint64) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldNotIn(FieldRequestAddedTime, vs...))
}

// RequestAddedTimeGT applies the GT predicate on the "request_added_time" field.
func RequestAddedTimeGT(v uint64) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldGT(FieldRequestAddedTime, v))
}

// RequestAddedTimeGTE applies the GTE predicate on the "request_added_time" field.
func RequestAddedTimeGTE(v uint64) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldGTE(FieldRequestAddedTime, v))
}

// RequestAddedTimeLT applies the LT predicate on the "request_added_time" field.
func RequestAddedTimeLT(v uint64) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldLT(FieldRequestAddedTime, v))
}

// RequestAddedTimeLTE applies the LTE predicate on the "request_added_time" field.
func RequestAddedTimeLTE(v uint64) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldLTE(FieldRequestAddedTime, v))
}

// ProverRequestIDEQ applies the EQ predicate on the "prover_request_id" field.
func ProverRequestIDEQ(v string) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldEQ(FieldProverRequestID, v))
}

// ProverRequestIDNEQ applies the NEQ predicate on the "prover_request_id" field.
func ProverRequestIDNEQ(v string) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldNEQ(FieldProverRequestID, v))
}

// ProverRequestIDIn applies the In predicate on the "prover_request_id" field.
func ProverRequestIDIn(vs ...string) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldIn(FieldProverRequestID, vs...))
}

// ProverRequestIDNotIn applies the NotIn predicate on the "prover_request_id" field.
func ProverRequestIDNotIn(vs ...string) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldNotIn(FieldProverRequestID, vs...))
}

// ProverRequestIDGT applies the GT predicate on the "prover_request_id" field.
func ProverRequestIDGT(v string) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldGT(FieldProverRequestID, v))
}

// ProverRequestIDGTE applies the GTE predicate on the "prover_request_id" field.
func ProverRequestIDGTE(v string) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldGTE(FieldProverRequestID, v))
}

// ProverRequestIDLT applies the LT predicate on the "prover_request_id" field.
func ProverRequestIDLT(v string) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldLT(FieldProverRequestID, v))
}

// ProverRequestIDLTE applies the LTE predicate on the "prover_request_id" field.
func ProverRequestIDLTE(v string) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldLTE(FieldProverRequestID, v))
}

// ProverRequestIDContains applies the Contains predicate on the "prover_request_id" field.
func ProverRequestIDContains(v string) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldContains(FieldProverRequestID, v))
}

// ProverRequestIDHasPrefix applies the HasPrefix predicate on the "prover_request_id" field.
func ProverRequestIDHasPrefix(v string) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldHasPrefix(FieldProverRequestID, v))
}

// ProverRequestIDHasSuffix applies the HasSuffix predicate on the "prover_request_id" field.
func ProverRequestIDHasSuffix(v string) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldHasSuffix(FieldProverRequestID, v))
}

// ProverRequestIDIsNil applies the IsNil predicate on the "prover_request_id" field.
func ProverRequestIDIsNil() predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldIsNull(FieldProverRequestID))
}

// ProverRequestIDNotNil applies the NotNil predicate on the "prover_request_id" field.
func ProverRequestIDNotNil() predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldNotNull(FieldProverRequestID))
}

// ProverRequestIDEqualFold applies the EqualFold predicate on the "prover_request_id" field.
func ProverRequestIDEqualFold(v string) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldEqualFold(FieldProverRequestID, v))
}

// ProverRequestIDContainsFold applies the ContainsFold predicate on the "prover_request_id" field.
func ProverRequestIDContainsFold(v string) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldContainsFold(FieldProverRequestID, v))
}

// ProofRequestTimeEQ applies the EQ predicate on the "proof_request_time" field.
func ProofRequestTimeEQ(v uint64) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldEQ(FieldProofRequestTime, v))
}

// ProofRequestTimeNEQ applies the NEQ predicate on the "proof_request_time" field.
func ProofRequestTimeNEQ(v uint64) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldNEQ(FieldProofRequestTime, v))
}

// ProofRequestTimeIn applies the In predicate on the "proof_request_time" field.
func ProofRequestTimeIn(vs ...uint64) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldIn(FieldProofRequestTime, vs...))
}

// ProofRequestTimeNotIn applies the NotIn predicate on the "proof_request_time" field.
func ProofRequestTimeNotIn(vs ...uint64) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldNotIn(FieldProofRequestTime, vs...))
}

// ProofRequestTimeGT applies the GT predicate on the "proof_request_time" field.
func ProofRequestTimeGT(v uint64) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldGT(FieldProofRequestTime, v))
}

// ProofRequestTimeGTE applies the GTE predicate on the "proof_request_time" field.
func ProofRequestTimeGTE(v uint64) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldGTE(FieldProofRequestTime, v))
}

// ProofRequestTimeLT applies the LT predicate on the "proof_request_time" field.
func ProofRequestTimeLT(v uint64) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldLT(FieldProofRequestTime, v))
}

// ProofRequestTimeLTE applies the LTE predicate on the "proof_request_time" field.
func ProofRequestTimeLTE(v uint64) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldLTE(FieldProofRequestTime, v))
}

// ProofRequestTimeIsNil applies the IsNil predicate on the "proof_request_time" field.
func ProofRequestTimeIsNil() predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldIsNull(FieldProofRequestTime))
}

// ProofRequestTimeNotNil applies the NotNil predicate on the "proof_request_time" field.
func ProofRequestTimeNotNil() predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldNotNull(FieldProofRequestTime))
}

// L1BlockNumberEQ applies the EQ predicate on the "l1_block_number" field.
func L1BlockNumberEQ(v uint64) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldEQ(FieldL1BlockNumber, v))
}

// L1BlockNumberNEQ applies the NEQ predicate on the "l1_block_number" field.
func L1BlockNumberNEQ(v uint64) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldNEQ(FieldL1BlockNumber, v))
}

// L1BlockNumberIn applies the In predicate on the "l1_block_number" field.
func L1BlockNumberIn(vs ...uint64) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldIn(FieldL1BlockNumber, vs...))
}

// L1BlockNumberNotIn applies the NotIn predicate on the "l1_block_number" field.
func L1BlockNumberNotIn(vs ...uint64) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldNotIn(FieldL1BlockNumber, vs...))
}

// L1BlockNumberGT applies the GT predicate on the "l1_block_number" field.
func L1BlockNumberGT(v uint64) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldGT(FieldL1BlockNumber, v))
}

// L1BlockNumberGTE applies the GTE predicate on the "l1_block_number" field.
func L1BlockNumberGTE(v uint64) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldGTE(FieldL1BlockNumber, v))
}

// L1BlockNumberLT applies the LT predicate on the "l1_block_number" field.
func L1BlockNumberLT(v uint64) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldLT(FieldL1BlockNumber, v))
}

// L1BlockNumberLTE applies the LTE predicate on the "l1_block_number" field.
func L1BlockNumberLTE(v uint64) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldLTE(FieldL1BlockNumber, v))
}

// L1BlockNumberIsNil applies the IsNil predicate on the "l1_block_number" field.
func L1BlockNumberIsNil() predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldIsNull(FieldL1BlockNumber))
}

// L1BlockNumberNotNil applies the NotNil predicate on the "l1_block_number" field.
func L1BlockNumberNotNil() predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldNotNull(FieldL1BlockNumber))
}

// L1BlockHashEQ applies the EQ predicate on the "l1_block_hash" field.
func L1BlockHashEQ(v string) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldEQ(FieldL1BlockHash, v))
}

// L1BlockHashNEQ applies the NEQ predicate on the "l1_block_hash" field.
func L1BlockHashNEQ(v string) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldNEQ(FieldL1BlockHash, v))
}

// L1BlockHashIn applies the In predicate on the "l1_block_hash" field.
func L1BlockHashIn(vs ...string) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldIn(FieldL1BlockHash, vs...))
}

// L1BlockHashNotIn applies the NotIn predicate on the "l1_block_hash" field.
func L1BlockHashNotIn(vs ...string) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldNotIn(FieldL1BlockHash, vs...))
}

// L1BlockHashGT applies the GT predicate on the "l1_block_hash" field.
func L1BlockHashGT(v string) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldGT(FieldL1BlockHash, v))
}

// L1BlockHashGTE applies the GTE predicate on the "l1_block_hash" field.
func L1BlockHashGTE(v string) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldGTE(FieldL1BlockHash, v))
}

// L1BlockHashLT applies the LT predicate on the "l1_block_hash" field.
func L1BlockHashLT(v string) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldLT(FieldL1BlockHash, v))
}

// L1BlockHashLTE applies the LTE predicate on the "l1_block_hash" field.
func L1BlockHashLTE(v string) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldLTE(FieldL1BlockHash, v))
}

// L1BlockHashContains applies the Contains predicate on the "l1_block_hash" field.
func L1BlockHashContains(v string) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldContains(FieldL1BlockHash, v))
}

// L1BlockHashHasPrefix applies the HasPrefix predicate on the "l1_block_hash" field.
func L1BlockHashHasPrefix(v string) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldHasPrefix(FieldL1BlockHash, v))
}

// L1BlockHashHasSuffix applies the HasSuffix predicate on the "l1_block_hash" field.
func L1BlockHashHasSuffix(v string) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldHasSuffix(FieldL1BlockHash, v))
}

// L1BlockHashIsNil applies the IsNil predicate on the "l1_block_hash" field.
func L1BlockHashIsNil() predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldIsNull(FieldL1BlockHash))
}

// L1BlockHashNotNil applies the NotNil predicate on the "l1_block_hash" field.
func L1BlockHashNotNil() predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldNotNull(FieldL1BlockHash))
}

// L1BlockHashEqualFold applies the EqualFold predicate on the "l1_block_hash" field.
func L1BlockHashEqualFold(v string) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldEqualFold(FieldL1BlockHash, v))
}

// L1BlockHashContainsFold applies the ContainsFold predicate on the "l1_block_hash" field.
func L1BlockHashContainsFold(v string) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldContainsFold(FieldL1BlockHash, v))
}

// ProofEQ applies the EQ predicate on the "proof" field.
func ProofEQ(v []byte) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldEQ(FieldProof, v))
}

// ProofNEQ applies the NEQ predicate on the "proof" field.
func ProofNEQ(v []byte) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldNEQ(FieldProof, v))
}

// ProofIn applies the In predicate on the "proof" field.
func ProofIn(vs ...[]byte) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldIn(FieldProof, vs...))
}

// ProofNotIn applies the NotIn predicate on the "proof" field.
func ProofNotIn(vs ...[]byte) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldNotIn(FieldProof, vs...))
}

// ProofGT applies the GT predicate on the "proof" field.
func ProofGT(v []byte) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldGT(FieldProof, v))
}

// ProofGTE applies the GTE predicate on the "proof" field.
func ProofGTE(v []byte) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldGTE(FieldProof, v))
}

// ProofLT applies the LT predicate on the "proof" field.
func ProofLT(v []byte) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldLT(FieldProof, v))
}

// ProofLTE applies the LTE predicate on the "proof" field.
func ProofLTE(v []byte) predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldLTE(FieldProof, v))
}

// ProofIsNil applies the IsNil predicate on the "proof" field.
func ProofIsNil() predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldIsNull(FieldProof))
}

// ProofNotNil applies the NotNil predicate on the "proof" field.
func ProofNotNil() predicate.ProofRequest {
	return predicate.ProofRequest(sql.FieldNotNull(FieldProof))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ProofRequest) predicate.ProofRequest {
	return predicate.ProofRequest(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ProofRequest) predicate.ProofRequest {
	return predicate.ProofRequest(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ProofRequest) predicate.ProofRequest {
	return predicate.ProofRequest(sql.NotPredicates(p))
}