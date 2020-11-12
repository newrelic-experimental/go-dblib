// SPDX-FileCopyrightText: 2020 SAP SE
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by "stringer -type=TDSMsgId"; DO NOT EDIT.

package tds

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[TDS_MSG_SEC_ENCRYPT-1]
	_ = x[TDS_MSG_SEC_LOGPWD-2]
	_ = x[TDS_MSG_SEC_REMPWD-3]
	_ = x[TDS_MSG_SEC_CHALLENGE-4]
	_ = x[TDS_MSG_SEC_RESPONSE-5]
	_ = x[TDS_MSG_SEC_GETLABEL-6]
	_ = x[TDS_MSG_SEC_LABEL-7]
	_ = x[TDS_MSG_SQL_TBLNAME-8]
	_ = x[TDS_MSG_GW_RESERVED-9]
	_ = x[TDS_MSG_OMNI_CAPABILITIES-10]
	_ = x[TDS_MSG_SEC_OPAQUE-11]
	_ = x[TDS_MSG_HAFAILOVER-12]
	_ = x[TDS_MSG_EMPTY-13]
	_ = x[TDS_MSG_SEC_ENCRYPT2-14]
	_ = x[TDS_MSG_SEC_LOGPWD2-15]
	_ = x[TDS_MSG_SEC_SUP_CIPHER2-16]
	_ = x[TDS_MSG_MIG_REQ-17]
	_ = x[TDS_MSG_MIG_SYNC-18]
	_ = x[TDS_MSG_MIG_CONT-19]
	_ = x[TDS_MSG_MIG_IGN-20]
	_ = x[TDS_MSG_MIG_FAIL-21]
	_ = x[TDS_MSG_SEC_REMPWD2-22]
	_ = x[TDS_MSG_MIG_RESUME-23]
	_ = x[TDS_MSG_HELLO-24]
	_ = x[TDS_MSG_LOGINPARAMS-25]
	_ = x[TDS_MSG_GRID_MIGREQ-26]
	_ = x[TDS_MSG_GRID_QUIESCE-27]
	_ = x[TDS_MSG_GRID_UNQUIESCE-28]
	_ = x[TDS_MSG_GRID_EVENT-29]
	_ = x[TDS_MSG_SEC_ENCRYPT3-30]
	_ = x[TDS_MSG_SEC_LOGPWD3-31]
	_ = x[TDS_MSG_SEC_REMPWD3-32]
	_ = x[TDS_MSG_DR_MAP-33]
	_ = x[TDS_MSG_SEC_SYMKEY-34]
	_ = x[TDS_MSG_SEC_ENCRYPT4-35]
}

const _TDSMsgId_name = "TDS_MSG_SEC_ENCRYPTTDS_MSG_SEC_LOGPWDTDS_MSG_SEC_REMPWDTDS_MSG_SEC_CHALLENGETDS_MSG_SEC_RESPONSETDS_MSG_SEC_GETLABELTDS_MSG_SEC_LABELTDS_MSG_SQL_TBLNAMETDS_MSG_GW_RESERVEDTDS_MSG_OMNI_CAPABILITIESTDS_MSG_SEC_OPAQUETDS_MSG_HAFAILOVERTDS_MSG_EMPTYTDS_MSG_SEC_ENCRYPT2TDS_MSG_SEC_LOGPWD2TDS_MSG_SEC_SUP_CIPHER2TDS_MSG_MIG_REQTDS_MSG_MIG_SYNCTDS_MSG_MIG_CONTTDS_MSG_MIG_IGNTDS_MSG_MIG_FAILTDS_MSG_SEC_REMPWD2TDS_MSG_MIG_RESUMETDS_MSG_HELLOTDS_MSG_LOGINPARAMSTDS_MSG_GRID_MIGREQTDS_MSG_GRID_QUIESCETDS_MSG_GRID_UNQUIESCETDS_MSG_GRID_EVENTTDS_MSG_SEC_ENCRYPT3TDS_MSG_SEC_LOGPWD3TDS_MSG_SEC_REMPWD3TDS_MSG_DR_MAPTDS_MSG_SEC_SYMKEYTDS_MSG_SEC_ENCRYPT4"

var _TDSMsgId_index = [...]uint16{0, 19, 37, 55, 76, 96, 116, 133, 152, 171, 196, 214, 232, 245, 265, 284, 307, 322, 338, 354, 369, 385, 404, 422, 435, 454, 473, 493, 515, 533, 553, 572, 591, 605, 623, 643}

func (i TDSMsgId) String() string {
	i -= 1
	if i >= TDSMsgId(len(_TDSMsgId_index)-1) {
		return "TDSMsgId(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _TDSMsgId_name[_TDSMsgId_index[i]:_TDSMsgId_index[i+1]]
}