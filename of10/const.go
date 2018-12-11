/*
 * Copyright (c) 2008 The Board of Trustees of The Leland Stanford Junior University
 * Copyright (c) 2011, 2012 Open Networking Foundation
 * Copyright 2013, Big Switch Networks, Inc. This library was generated by the LoxiGen Compiler.
 * Copyright 2018, Red Hat, Inc.
 */
// Automatically generated by LOXI from template const.go
// Do not modify

package of10

import (
	"fmt"
	"strings"
)

const (
	// Identifiers from group macro_definitions
	MaxTableNameLen    = 32         // OFP_MAX_TABLE_NAME_LEN
	MaxPortNameLen     = 16         // OFP_MAX_PORT_NAME_LEN
	TCPPort            = 6653       // OFP_TCP_PORT
	SSLPort            = 6653       // OFP_SSL_PORT
	EthAlen            = 6          // OFP_ETH_ALEN
	DefaultMissSendLen = 128        // OFP_DEFAULT_MISS_SEND_LEN
	VLANNone           = 65535      // OFP_VLAN_NONE
	OFPFWICMPType      = 64         // OFPFW_ICMP_TYPE
	OFPFWICMPCode      = 128        // OFPFW_ICMP_CODE
	DlTypeEth2Cutoff   = 1536       // OFP_DL_TYPE_ETH2_CUTOFF
	DlTypeNotEthType   = 1535       // OFP_DL_TYPE_NOT_ETH_TYPE
	FlowPermanent      = 0          // OFP_FLOW_PERMANENT
	DefaultPriority    = 32768      // OFP_DEFAULT_PRIORITY
	DescStrLen         = 256        // DESC_STR_LEN
	SerialNumLen       = 32         // SERIAL_NUM_LEN
	OFPQAll            = 4294967295 // OFPQ_ALL
	OFPQMinRateUncfg   = 65535      // OFPQ_MIN_RATE_UNCFG
)

const (
	// Identifiers from group nx_flow_monitor_flags
	NxfmfInitial = 1  // NXFMF_INITIAL
	NxfmfAdd     = 2  // NXFMF_ADD
	NxfmfDelete  = 4  // NXFMF_DELETE
	NxfmfModify  = 8  // NXFMF_MODIFY
	NxfmfActions = 16 // NXFMF_ACTIONS
	NxfmfOwn     = 32 // NXFMF_OWN
)

type NxFlowMonitorFlags uint16

func (self NxFlowMonitorFlags) MarshalJSON() ([]byte, error) {
	var flags []string
	if self&NxfmfInitial == NxfmfInitial {
		flags = append(flags, "\"NxfmfInitial\": true")
	}
	if self&NxfmfAdd == NxfmfAdd {
		flags = append(flags, "\"NxfmfAdd\": true")
	}
	if self&NxfmfDelete == NxfmfDelete {
		flags = append(flags, "\"NxfmfDelete\": true")
	}
	if self&NxfmfModify == NxfmfModify {
		flags = append(flags, "\"NxfmfModify\": true")
	}
	if self&NxfmfActions == NxfmfActions {
		flags = append(flags, "\"NxfmfActions\": true")
	}
	if self&NxfmfOwn == NxfmfOwn {
		flags = append(flags, "\"NxfmfOwn\": true")
	}
	return []byte("{" + strings.Join(flags, ", ") + "}"), nil
}

const (
	// Identifiers from group of_bsn_pdu_slot_num
	BsnPduSlotNumAny = 255 // BSN_PDU_SLOT_NUM_ANY
)

type BsnPduSlotNum uint8

func (self BsnPduSlotNum) MarshalJSON() ([]byte, error) {
	switch self {
	case BsnPduSlotNumAny:
		return []byte("\"BsnPduSlotNumAny\""), nil
	default:
		return nil, fmt.Errorf("Invalid value '%d' for BsnPduSlotNum", self)
	}
}

const (
	// Identifiers from group ofp_action_type
	OFPATOutput       = 0     // OFPAT_OUTPUT
	OFPATSetVLANVid   = 1     // OFPAT_SET_VLAN_VID
	OFPATSetVLANPCP   = 2     // OFPAT_SET_VLAN_PCP
	OFPATStripVLAN    = 3     // OFPAT_STRIP_VLAN
	OFPATSetDlSrc     = 4     // OFPAT_SET_DL_SRC
	OFPATSetDlDst     = 5     // OFPAT_SET_DL_DST
	OFPATSetNwSrc     = 6     // OFPAT_SET_NW_SRC
	OFPATSetNwDst     = 7     // OFPAT_SET_NW_DST
	OFPATSetNwTos     = 8     // OFPAT_SET_NW_TOS
	OFPATSetTpSrc     = 9     // OFPAT_SET_TP_SRC
	OFPATSetTpDst     = 10    // OFPAT_SET_TP_DST
	OFPATEnqueue      = 11    // OFPAT_ENQUEUE
	OFPATExperimenter = 65535 // OFPAT_EXPERIMENTER
)

type ActionType uint16

func (self ActionType) MarshalJSON() ([]byte, error) {
	switch self {
	case OFPATOutput:
		return []byte("\"OFPATOutput\""), nil
	case OFPATSetVLANVid:
		return []byte("\"OFPATSetVLANVid\""), nil
	case OFPATSetVLANPCP:
		return []byte("\"OFPATSetVLANPCP\""), nil
	case OFPATStripVLAN:
		return []byte("\"OFPATStripVLAN\""), nil
	case OFPATSetDlSrc:
		return []byte("\"OFPATSetDlSrc\""), nil
	case OFPATSetDlDst:
		return []byte("\"OFPATSetDlDst\""), nil
	case OFPATSetNwSrc:
		return []byte("\"OFPATSetNwSrc\""), nil
	case OFPATSetNwDst:
		return []byte("\"OFPATSetNwDst\""), nil
	case OFPATSetNwTos:
		return []byte("\"OFPATSetNwTos\""), nil
	case OFPATSetTpSrc:
		return []byte("\"OFPATSetTpSrc\""), nil
	case OFPATSetTpDst:
		return []byte("\"OFPATSetTpDst\""), nil
	case OFPATEnqueue:
		return []byte("\"OFPATEnqueue\""), nil
	case OFPATExperimenter:
		return []byte("\"OFPATExperimenter\""), nil
	default:
		return nil, fmt.Errorf("Invalid value '%d' for ActionType", self)
	}
}

const (
	// Identifiers from group ofp_bad_action_code
	OFPBACBadType             = 0 // OFPBAC_BAD_TYPE
	OFPBACBadLen              = 1 // OFPBAC_BAD_LEN
	OFPBACBadExperimenter     = 2 // OFPBAC_BAD_EXPERIMENTER
	OFPBACBadExperimenterType = 3 // OFPBAC_BAD_EXPERIMENTER_TYPE
	OFPBACBadOutPort          = 4 // OFPBAC_BAD_OUT_PORT
	OFPBACBadArgument         = 5 // OFPBAC_BAD_ARGUMENT
	OFPBACEperm               = 6 // OFPBAC_EPERM
	OFPBACTooMany             = 7 // OFPBAC_TOO_MANY
	OFPBACBadQueue            = 8 // OFPBAC_BAD_QUEUE
)

type BadActionCode uint16

func (self BadActionCode) MarshalJSON() ([]byte, error) {
	switch self {
	case OFPBACBadType:
		return []byte("\"OFPBACBadType\""), nil
	case OFPBACBadLen:
		return []byte("\"OFPBACBadLen\""), nil
	case OFPBACBadExperimenter:
		return []byte("\"OFPBACBadExperimenter\""), nil
	case OFPBACBadExperimenterType:
		return []byte("\"OFPBACBadExperimenterType\""), nil
	case OFPBACBadOutPort:
		return []byte("\"OFPBACBadOutPort\""), nil
	case OFPBACBadArgument:
		return []byte("\"OFPBACBadArgument\""), nil
	case OFPBACEperm:
		return []byte("\"OFPBACEperm\""), nil
	case OFPBACTooMany:
		return []byte("\"OFPBACTooMany\""), nil
	case OFPBACBadQueue:
		return []byte("\"OFPBACBadQueue\""), nil
	default:
		return nil, fmt.Errorf("Invalid value '%d' for BadActionCode", self)
	}
}

const (
	// Identifiers from group ofp_bad_request_code
	OFPBRCBadVersion      = 0 // OFPBRC_BAD_VERSION
	OFPBRCBadType         = 1 // OFPBRC_BAD_TYPE
	OFPBRCBadStat         = 2 // OFPBRC_BAD_STAT
	OFPBRCBadExperimenter = 3 // OFPBRC_BAD_EXPERIMENTER
	OFPBRCBadSubtype      = 4 // OFPBRC_BAD_SUBTYPE
	OFPBRCEperm           = 5 // OFPBRC_EPERM
	OFPBRCBadLen          = 6 // OFPBRC_BAD_LEN
	OFPBRCBufferEmpty     = 7 // OFPBRC_BUFFER_EMPTY
	OFPBRCBufferUnknown   = 8 // OFPBRC_BUFFER_UNKNOWN
)

type BadRequestCode uint16

func (self BadRequestCode) MarshalJSON() ([]byte, error) {
	switch self {
	case OFPBRCBadVersion:
		return []byte("\"OFPBRCBadVersion\""), nil
	case OFPBRCBadType:
		return []byte("\"OFPBRCBadType\""), nil
	case OFPBRCBadStat:
		return []byte("\"OFPBRCBadStat\""), nil
	case OFPBRCBadExperimenter:
		return []byte("\"OFPBRCBadExperimenter\""), nil
	case OFPBRCBadSubtype:
		return []byte("\"OFPBRCBadSubtype\""), nil
	case OFPBRCEperm:
		return []byte("\"OFPBRCEperm\""), nil
	case OFPBRCBadLen:
		return []byte("\"OFPBRCBadLen\""), nil
	case OFPBRCBufferEmpty:
		return []byte("\"OFPBRCBufferEmpty\""), nil
	case OFPBRCBufferUnknown:
		return []byte("\"OFPBRCBufferUnknown\""), nil
	default:
		return nil, fmt.Errorf("Invalid value '%d' for BadRequestCode", self)
	}
}

const (
	// Identifiers from group ofp_bsn_vport_l2gre_flags
	OFBSNVportL2GreLocalMACIsValid  = 1  // OF_BSN_VPORT_L2GRE_LOCAL_MAC_IS_VALID
	OFBSNVportL2GreDSCPAssign       = 2  // OF_BSN_VPORT_L2GRE_DSCP_ASSIGN
	OFBSNVportL2GreDSCPCopy         = 4  // OF_BSN_VPORT_L2GRE_DSCP_COPY
	OFBSNVportL2GreLoopbackIsValid  = 8  // OF_BSN_VPORT_L2GRE_LOOPBACK_IS_VALID
	OFBSNVportL2GreRateLimitIsValid = 16 // OF_BSN_VPORT_L2GRE_RATE_LIMIT_IS_VALID
)

type BsnVportL2GreFlags uint32

func (self BsnVportL2GreFlags) MarshalJSON() ([]byte, error) {
	var flags []string
	if self&OFBSNVportL2GreLocalMACIsValid == OFBSNVportL2GreLocalMACIsValid {
		flags = append(flags, "\"OFBSNVportL2GreLocalMACIsValid\": true")
	}
	if self&OFBSNVportL2GreDSCPAssign == OFBSNVportL2GreDSCPAssign {
		flags = append(flags, "\"OFBSNVportL2GreDSCPAssign\": true")
	}
	if self&OFBSNVportL2GreDSCPCopy == OFBSNVportL2GreDSCPCopy {
		flags = append(flags, "\"OFBSNVportL2GreDSCPCopy\": true")
	}
	if self&OFBSNVportL2GreLoopbackIsValid == OFBSNVportL2GreLoopbackIsValid {
		flags = append(flags, "\"OFBSNVportL2GreLoopbackIsValid\": true")
	}
	if self&OFBSNVportL2GreRateLimitIsValid == OFBSNVportL2GreRateLimitIsValid {
		flags = append(flags, "\"OFBSNVportL2GreRateLimitIsValid\": true")
	}
	return []byte("{" + strings.Join(flags, ", ") + "}"), nil
}

const (
	// Identifiers from group ofp_bsn_vport_q_in_q_untagged
	OFBSNVportQInQUntagged = 65535 // OF_BSN_VPORT_Q_IN_Q_UNTAGGED
)

type BsnVportQInQUntagged uint16

func (self BsnVportQInQUntagged) MarshalJSON() ([]byte, error) {
	switch self {
	case OFBSNVportQInQUntagged:
		return []byte("\"OFBSNVportQInQUntagged\""), nil
	default:
		return nil, fmt.Errorf("Invalid value '%d' for BsnVportQInQUntagged", self)
	}
}

const (
	// Identifiers from group ofp_bsn_vport_status
	OFBSNVportStatusOk     = 0 // OF_BSN_VPORT_STATUS_OK
	OFBSNVportStatusFailed = 1 // OF_BSN_VPORT_STATUS_FAILED
)

const (
	// Identifiers from group ofp_capabilities
	OFPCFlowStats  = 1   // OFPC_FLOW_STATS
	OFPCTableStats = 2   // OFPC_TABLE_STATS
	OFPCPortStats  = 4   // OFPC_PORT_STATS
	OFPCSTP        = 8   // OFPC_STP
	OFPCReserved   = 16  // OFPC_RESERVED
	OFPCIpReasm    = 32  // OFPC_IP_REASM
	OFPCQueueStats = 64  // OFPC_QUEUE_STATS
	OFPCARPMatchIp = 128 // OFPC_ARP_MATCH_IP
)

type Capabilities uint32

func (self Capabilities) MarshalJSON() ([]byte, error) {
	var flags []string
	if self&OFPCFlowStats == OFPCFlowStats {
		flags = append(flags, "\"OFPCFlowStats\": true")
	}
	if self&OFPCTableStats == OFPCTableStats {
		flags = append(flags, "\"OFPCTableStats\": true")
	}
	if self&OFPCPortStats == OFPCPortStats {
		flags = append(flags, "\"OFPCPortStats\": true")
	}
	if self&OFPCSTP == OFPCSTP {
		flags = append(flags, "\"OFPCSTP\": true")
	}
	if self&OFPCReserved == OFPCReserved {
		flags = append(flags, "\"OFPCReserved\": true")
	}
	if self&OFPCIpReasm == OFPCIpReasm {
		flags = append(flags, "\"OFPCIpReasm\": true")
	}
	if self&OFPCQueueStats == OFPCQueueStats {
		flags = append(flags, "\"OFPCQueueStats\": true")
	}
	if self&OFPCARPMatchIp == OFPCARPMatchIp {
		flags = append(flags, "\"OFPCARPMatchIp\": true")
	}
	return []byte("{" + strings.Join(flags, ", ") + "}"), nil
}

const (
	// Identifiers from group ofp_config_flags
	OFPCFragNormal = 0 // OFPC_FRAG_NORMAL
	OFPCFragDrop   = 1 // OFPC_FRAG_DROP
	OFPCFragReasm  = 2 // OFPC_FRAG_REASM
	OFPCFragMask   = 3 // OFPC_FRAG_MASK
)

type ConfigFlags uint16

func (self ConfigFlags) MarshalJSON() ([]byte, error) {
	var flags []string
	if self&OFPCFragNormal == OFPCFragNormal {
		flags = append(flags, "\"OFPCFragNormal\": true")
	}
	if self&OFPCFragDrop == OFPCFragDrop {
		flags = append(flags, "\"OFPCFragDrop\": true")
	}
	if self&OFPCFragReasm == OFPCFragReasm {
		flags = append(flags, "\"OFPCFragReasm\": true")
	}
	if self&OFPCFragMask == OFPCFragMask {
		flags = append(flags, "\"OFPCFragMask\": true")
	}
	return []byte("{" + strings.Join(flags, ", ") + "}"), nil
}

const (
	// Identifiers from group ofp_ed_nsh_prop_type
	OFPPPTPropNshNone   = 0 // OFPPPT_PROP_NSH_NONE
	OFPPPTPropNshMdtype = 1 // OFPPPT_PROP_NSH_MDTYPE
	OFPPPTPropNshTlv    = 2 // OFPPPT_PROP_NSH_TLV
)

type EdNshPropType uint8

func (self EdNshPropType) MarshalJSON() ([]byte, error) {
	switch self {
	case OFPPPTPropNshNone:
		return []byte("\"OFPPPTPropNshNone\""), nil
	case OFPPPTPropNshMdtype:
		return []byte("\"OFPPPTPropNshMdtype\""), nil
	case OFPPPTPropNshTlv:
		return []byte("\"OFPPPTPropNshTlv\""), nil
	default:
		return nil, fmt.Errorf("Invalid value '%d' for EdNshPropType", self)
	}
}

const (
	// Identifiers from group ofp_ed_prop_class
	OFPPPCBasic        = 0     // OFPPPC_BASIC
	OFPPPCMpls         = 1     // OFPPPC_MPLS
	OFPPPCGRE          = 2     // OFPPPC_GRE
	OFPPPCGtp          = 3     // OFPPPC_GTP
	OFPPPCNsh          = 4     // OFPPPC_NSH
	OFPPPCExperimenter = 65535 // OFPPPC_EXPERIMENTER
)

type EdPropClass uint16

func (self EdPropClass) MarshalJSON() ([]byte, error) {
	switch self {
	case OFPPPCBasic:
		return []byte("\"OFPPPCBasic\""), nil
	case OFPPPCMpls:
		return []byte("\"OFPPPCMpls\""), nil
	case OFPPPCGRE:
		return []byte("\"OFPPPCGRE\""), nil
	case OFPPPCGtp:
		return []byte("\"OFPPPCGtp\""), nil
	case OFPPPCNsh:
		return []byte("\"OFPPPCNsh\""), nil
	case OFPPPCExperimenter:
		return []byte("\"OFPPPCExperimenter\""), nil
	default:
		return nil, fmt.Errorf("Invalid value '%d' for EdPropClass", self)
	}
}

const (
	// Identifiers from group ofp_error_type
	OFPETHelloFailed   = 0 // OFPET_HELLO_FAILED
	OFPETBadRequest    = 1 // OFPET_BAD_REQUEST
	OFPETBadAction     = 2 // OFPET_BAD_ACTION
	OFPETFlowModFailed = 3 // OFPET_FLOW_MOD_FAILED
	OFPETPortModFailed = 4 // OFPET_PORT_MOD_FAILED
	OFPETQueueOpFailed = 5 // OFPET_QUEUE_OP_FAILED
)

type ErrorType uint16

func (self ErrorType) MarshalJSON() ([]byte, error) {
	switch self {
	case OFPETHelloFailed:
		return []byte("\"OFPETHelloFailed\""), nil
	case OFPETBadRequest:
		return []byte("\"OFPETBadRequest\""), nil
	case OFPETBadAction:
		return []byte("\"OFPETBadAction\""), nil
	case OFPETFlowModFailed:
		return []byte("\"OFPETFlowModFailed\""), nil
	case OFPETPortModFailed:
		return []byte("\"OFPETPortModFailed\""), nil
	case OFPETQueueOpFailed:
		return []byte("\"OFPETQueueOpFailed\""), nil
	default:
		return nil, fmt.Errorf("Invalid value '%d' for ErrorType", self)
	}
}

const (
	// Identifiers from group ofp_flow_mod_command
	OFPFCAdd          = 0 // OFPFC_ADD
	OFPFCModify       = 1 // OFPFC_MODIFY
	OFPFCModifyStrict = 2 // OFPFC_MODIFY_STRICT
	OFPFCDelete       = 3 // OFPFC_DELETE
	OFPFCDeleteStrict = 4 // OFPFC_DELETE_STRICT
)

type FlowModCommand uint16

func (self FlowModCommand) MarshalJSON() ([]byte, error) {
	switch self {
	case OFPFCAdd:
		return []byte("\"OFPFCAdd\""), nil
	case OFPFCModify:
		return []byte("\"OFPFCModify\""), nil
	case OFPFCModifyStrict:
		return []byte("\"OFPFCModifyStrict\""), nil
	case OFPFCDelete:
		return []byte("\"OFPFCDelete\""), nil
	case OFPFCDeleteStrict:
		return []byte("\"OFPFCDeleteStrict\""), nil
	default:
		return nil, fmt.Errorf("Invalid value '%d' for FlowModCommand", self)
	}
}

const (
	// Identifiers from group ofp_flow_mod_failed_code
	OFPFMFCAllTablesFull   = 0 // OFPFMFC_ALL_TABLES_FULL
	OFPFMFCOverlap         = 1 // OFPFMFC_OVERLAP
	OFPFMFCEperm           = 2 // OFPFMFC_EPERM
	OFPFMFCBadEmergTimeout = 3 // OFPFMFC_BAD_EMERG_TIMEOUT
	OFPFMFCBadCommand      = 4 // OFPFMFC_BAD_COMMAND
	OFPFMFCUnsupported     = 5 // OFPFMFC_UNSUPPORTED
)

type FlowModFailedCode uint16

func (self FlowModFailedCode) MarshalJSON() ([]byte, error) {
	switch self {
	case OFPFMFCAllTablesFull:
		return []byte("\"OFPFMFCAllTablesFull\""), nil
	case OFPFMFCOverlap:
		return []byte("\"OFPFMFCOverlap\""), nil
	case OFPFMFCEperm:
		return []byte("\"OFPFMFCEperm\""), nil
	case OFPFMFCBadEmergTimeout:
		return []byte("\"OFPFMFCBadEmergTimeout\""), nil
	case OFPFMFCBadCommand:
		return []byte("\"OFPFMFCBadCommand\""), nil
	case OFPFMFCUnsupported:
		return []byte("\"OFPFMFCUnsupported\""), nil
	default:
		return nil, fmt.Errorf("Invalid value '%d' for FlowModFailedCode", self)
	}
}

const (
	// Identifiers from group ofp_flow_mod_flags
	OFPFFSendFlowRem  = 1 // OFPFF_SEND_FLOW_REM
	OFPFFCheckOverlap = 2 // OFPFF_CHECK_OVERLAP
	OFPFFEmerg        = 4 // OFPFF_EMERG
)

type FlowModFlags uint16

func (self FlowModFlags) MarshalJSON() ([]byte, error) {
	var flags []string
	if self&OFPFFSendFlowRem == OFPFFSendFlowRem {
		flags = append(flags, "\"OFPFFSendFlowRem\": true")
	}
	if self&OFPFFCheckOverlap == OFPFFCheckOverlap {
		flags = append(flags, "\"OFPFFCheckOverlap\": true")
	}
	if self&OFPFFEmerg == OFPFFEmerg {
		flags = append(flags, "\"OFPFFEmerg\": true")
	}
	return []byte("{" + strings.Join(flags, ", ") + "}"), nil
}

const (
	// Identifiers from group ofp_flow_removed_reason
	OFPRRIdleTimeout = 0 // OFPRR_IDLE_TIMEOUT
	OFPRRHardTimeout = 1 // OFPRR_HARD_TIMEOUT
	OFPRRDelete      = 2 // OFPRR_DELETE
)

type FlowRemovedReason uint8

func (self FlowRemovedReason) MarshalJSON() ([]byte, error) {
	switch self {
	case OFPRRIdleTimeout:
		return []byte("\"OFPRRIdleTimeout\""), nil
	case OFPRRHardTimeout:
		return []byte("\"OFPRRHardTimeout\""), nil
	case OFPRRDelete:
		return []byte("\"OFPRRDelete\""), nil
	default:
		return nil, fmt.Errorf("Invalid value '%d' for FlowRemovedReason", self)
	}
}

const (
	// Identifiers from group ofp_flow_wildcards
	OFPFWInPort     = 1       // OFPFW_IN_PORT
	OFPFWDlVLAN     = 2       // OFPFW_DL_VLAN
	OFPFWDlSrc      = 4       // OFPFW_DL_SRC
	OFPFWNwDstBits  = 6       // OFPFW_NW_DST_BITS
	OFPFWNwSrcBits  = 6       // OFPFW_NW_SRC_BITS
	OFPFWNwSrcShift = 8       // OFPFW_NW_SRC_SHIFT
	OFPFWDlDst      = 8       // OFPFW_DL_DST
	OFPFWNwDstShift = 14      // OFPFW_NW_DST_SHIFT
	OFPFWDlType     = 16      // OFPFW_DL_TYPE
	OFPFWNwProto    = 32      // OFPFW_NW_PROTO
	OFPFWTpSrc      = 64      // OFPFW_TP_SRC
	OFPFWTpDst      = 128     // OFPFW_TP_DST
	OFPFWNwSrcAll   = 8192    // OFPFW_NW_SRC_ALL
	OFPFWNwSrcMask  = 16128   // OFPFW_NW_SRC_MASK
	OFPFWNwDstAll   = 524288  // OFPFW_NW_DST_ALL
	OFPFWNwDstMask  = 1032192 // OFPFW_NW_DST_MASK
	OFPFWDlVLANPCP  = 1048576 // OFPFW_DL_VLAN_PCP
	OFPFWNwTos      = 2097152 // OFPFW_NW_TOS
	OFPFWAll        = 4194303 // OFPFW_ALL
)

type FlowWildcards uint32

func (self FlowWildcards) MarshalJSON() ([]byte, error) {
	var flags []string
	if self&OFPFWInPort == OFPFWInPort {
		flags = append(flags, "\"OFPFWInPort\": true")
	}
	if self&OFPFWDlVLAN == OFPFWDlVLAN {
		flags = append(flags, "\"OFPFWDlVLAN\": true")
	}
	if self&OFPFWDlSrc == OFPFWDlSrc {
		flags = append(flags, "\"OFPFWDlSrc\": true")
	}
	if self&OFPFWNwDstBits == OFPFWNwDstBits {
		flags = append(flags, "\"OFPFWNwDstBits\": true")
	}
	if self&OFPFWNwSrcBits == OFPFWNwSrcBits {
		flags = append(flags, "\"OFPFWNwSrcBits\": true")
	}
	if self&OFPFWNwSrcShift == OFPFWNwSrcShift {
		flags = append(flags, "\"OFPFWNwSrcShift\": true")
	}
	if self&OFPFWDlDst == OFPFWDlDst {
		flags = append(flags, "\"OFPFWDlDst\": true")
	}
	if self&OFPFWNwDstShift == OFPFWNwDstShift {
		flags = append(flags, "\"OFPFWNwDstShift\": true")
	}
	if self&OFPFWDlType == OFPFWDlType {
		flags = append(flags, "\"OFPFWDlType\": true")
	}
	if self&OFPFWNwProto == OFPFWNwProto {
		flags = append(flags, "\"OFPFWNwProto\": true")
	}
	if self&OFPFWTpSrc == OFPFWTpSrc {
		flags = append(flags, "\"OFPFWTpSrc\": true")
	}
	if self&OFPFWTpDst == OFPFWTpDst {
		flags = append(flags, "\"OFPFWTpDst\": true")
	}
	if self&OFPFWNwSrcAll == OFPFWNwSrcAll {
		flags = append(flags, "\"OFPFWNwSrcAll\": true")
	}
	if self&OFPFWNwSrcMask == OFPFWNwSrcMask {
		flags = append(flags, "\"OFPFWNwSrcMask\": true")
	}
	if self&OFPFWNwDstAll == OFPFWNwDstAll {
		flags = append(flags, "\"OFPFWNwDstAll\": true")
	}
	if self&OFPFWNwDstMask == OFPFWNwDstMask {
		flags = append(flags, "\"OFPFWNwDstMask\": true")
	}
	if self&OFPFWDlVLANPCP == OFPFWDlVLANPCP {
		flags = append(flags, "\"OFPFWDlVLANPCP\": true")
	}
	if self&OFPFWNwTos == OFPFWNwTos {
		flags = append(flags, "\"OFPFWNwTos\": true")
	}
	if self&OFPFWAll == OFPFWAll {
		flags = append(flags, "\"OFPFWAll\": true")
	}
	return []byte("{" + strings.Join(flags, ", ") + "}"), nil
}

const (
	// Identifiers from group ofp_hello_failed_code
	OFPHFCIncompatible = 0 // OFPHFC_INCOMPATIBLE
	OFPHFCEperm        = 1 // OFPHFC_EPERM
)

type HelloFailedCode uint16

func (self HelloFailedCode) MarshalJSON() ([]byte, error) {
	switch self {
	case OFPHFCIncompatible:
		return []byte("\"OFPHFCIncompatible\""), nil
	case OFPHFCEperm:
		return []byte("\"OFPHFCEperm\""), nil
	default:
		return nil, fmt.Errorf("Invalid value '%d' for HelloFailedCode", self)
	}
}

const (
	// Identifiers from group ofp_nicira_controller_role
	NxRoleOther  = 0 // NX_ROLE_OTHER
	NxRoleMaster = 1 // NX_ROLE_MASTER
	NxRoleSlave  = 2 // NX_ROLE_SLAVE
)

type NiciraControllerRole uint32

func (self NiciraControllerRole) MarshalJSON() ([]byte, error) {
	switch self {
	case NxRoleOther:
		return []byte("\"NxRoleOther\""), nil
	case NxRoleMaster:
		return []byte("\"NxRoleMaster\""), nil
	case NxRoleSlave:
		return []byte("\"NxRoleSlave\""), nil
	default:
		return nil, fmt.Errorf("Invalid value '%d' for NiciraControllerRole", self)
	}
}

const (
	// Identifiers from group ofp_packet_in_reason
	OFPRNoMatch = 0 // OFPR_NO_MATCH
	OFPRAction  = 1 // OFPR_ACTION
)

type PacketInReason uint8

func (self PacketInReason) MarshalJSON() ([]byte, error) {
	switch self {
	case OFPRNoMatch:
		return []byte("\"OFPRNoMatch\""), nil
	case OFPRAction:
		return []byte("\"OFPRAction\""), nil
	default:
		return nil, fmt.Errorf("Invalid value '%d' for PacketInReason", self)
	}
}

const (
	// Identifiers from group ofp_port
	OFPPMax        = 0xff00 // OFPP_MAX
	OFPPInPort     = 0xfff8 // OFPP_IN_PORT
	OFPPTable      = 0xfff9 // OFPP_TABLE
	OFPPNormal     = 0xfffa // OFPP_NORMAL
	OFPPFlood      = 0xfffb // OFPP_FLOOD
	OFPPAll        = 0xfffc // OFPP_ALL
	OFPPController = 0xfffd // OFPP_CONTROLLER
	OFPPLocal      = 0xfffe // OFPP_LOCAL
	OFPPNone       = 0xffff // OFPP_NONE
)

type Port uint16

func (self Port) MarshalJSON() ([]byte, error) {
	switch self {
	case OFPPMax:
		return []byte("\"OFPPMax\""), nil
	case OFPPInPort:
		return []byte("\"OFPPInPort\""), nil
	case OFPPTable:
		return []byte("\"OFPPTable\""), nil
	case OFPPNormal:
		return []byte("\"OFPPNormal\""), nil
	case OFPPFlood:
		return []byte("\"OFPPFlood\""), nil
	case OFPPAll:
		return []byte("\"OFPPAll\""), nil
	case OFPPController:
		return []byte("\"OFPPController\""), nil
	case OFPPLocal:
		return []byte("\"OFPPLocal\""), nil
	case OFPPNone:
		return []byte("\"OFPPNone\""), nil
	default:
		return nil, fmt.Errorf("Invalid value '%d' for Port", self)
	}
}

const (
	// Identifiers from group ofp_port_config
	OFPPCPortDown      = 1          // OFPPC_PORT_DOWN
	OFPPCNoSTP         = 2          // OFPPC_NO_STP
	OFPPCNoRecv        = 4          // OFPPC_NO_RECV
	OFPPCNoRecvSTP     = 8          // OFPPC_NO_RECV_STP
	OFPPCNoFlood       = 16         // OFPPC_NO_FLOOD
	OFPPCNoFwd         = 32         // OFPPC_NO_FWD
	OFPPCNoPacketIn    = 64         // OFPPC_NO_PACKET_IN
	OFPPCBSNMirrorDest = 2147483648 // OFPPC_BSN_MIRROR_DEST
)

type PortConfig uint32

func (self PortConfig) MarshalJSON() ([]byte, error) {
	var flags []string
	if self&OFPPCPortDown == OFPPCPortDown {
		flags = append(flags, "\"OFPPCPortDown\": true")
	}
	if self&OFPPCNoSTP == OFPPCNoSTP {
		flags = append(flags, "\"OFPPCNoSTP\": true")
	}
	if self&OFPPCNoRecv == OFPPCNoRecv {
		flags = append(flags, "\"OFPPCNoRecv\": true")
	}
	if self&OFPPCNoRecvSTP == OFPPCNoRecvSTP {
		flags = append(flags, "\"OFPPCNoRecvSTP\": true")
	}
	if self&OFPPCNoFlood == OFPPCNoFlood {
		flags = append(flags, "\"OFPPCNoFlood\": true")
	}
	if self&OFPPCNoFwd == OFPPCNoFwd {
		flags = append(flags, "\"OFPPCNoFwd\": true")
	}
	if self&OFPPCNoPacketIn == OFPPCNoPacketIn {
		flags = append(flags, "\"OFPPCNoPacketIn\": true")
	}
	if self&OFPPCBSNMirrorDest == OFPPCBSNMirrorDest {
		flags = append(flags, "\"OFPPCBSNMirrorDest\": true")
	}
	return []byte("{" + strings.Join(flags, ", ") + "}"), nil
}

const (
	// Identifiers from group ofp_port_features
	OFPPF10MbHd    = 1    // OFPPF_10MB_HD
	OFPPF10MbFd    = 2    // OFPPF_10MB_FD
	OFPPF100MbHd   = 4    // OFPPF_100MB_HD
	OFPPF100MbFd   = 8    // OFPPF_100MB_FD
	OFPPF1GbHd     = 16   // OFPPF_1GB_HD
	OFPPF1GbFd     = 32   // OFPPF_1GB_FD
	OFPPF10GbFd    = 64   // OFPPF_10GB_FD
	OFPPFCopper    = 128  // OFPPF_COPPER
	OFPPFFiber     = 256  // OFPPF_FIBER
	OFPPFAutoneg   = 512  // OFPPF_AUTONEG
	OFPPFPause     = 1024 // OFPPF_PAUSE
	OFPPFPauseAsym = 2048 // OFPPF_PAUSE_ASYM
)

type PortFeatures uint32

func (self PortFeatures) MarshalJSON() ([]byte, error) {
	var flags []string
	if self&OFPPF10MbHd == OFPPF10MbHd {
		flags = append(flags, "\"OFPPF10MbHd\": true")
	}
	if self&OFPPF10MbFd == OFPPF10MbFd {
		flags = append(flags, "\"OFPPF10MbFd\": true")
	}
	if self&OFPPF100MbHd == OFPPF100MbHd {
		flags = append(flags, "\"OFPPF100MbHd\": true")
	}
	if self&OFPPF100MbFd == OFPPF100MbFd {
		flags = append(flags, "\"OFPPF100MbFd\": true")
	}
	if self&OFPPF1GbHd == OFPPF1GbHd {
		flags = append(flags, "\"OFPPF1GbHd\": true")
	}
	if self&OFPPF1GbFd == OFPPF1GbFd {
		flags = append(flags, "\"OFPPF1GbFd\": true")
	}
	if self&OFPPF10GbFd == OFPPF10GbFd {
		flags = append(flags, "\"OFPPF10GbFd\": true")
	}
	if self&OFPPFCopper == OFPPFCopper {
		flags = append(flags, "\"OFPPFCopper\": true")
	}
	if self&OFPPFFiber == OFPPFFiber {
		flags = append(flags, "\"OFPPFFiber\": true")
	}
	if self&OFPPFAutoneg == OFPPFAutoneg {
		flags = append(flags, "\"OFPPFAutoneg\": true")
	}
	if self&OFPPFPause == OFPPFPause {
		flags = append(flags, "\"OFPPFPause\": true")
	}
	if self&OFPPFPauseAsym == OFPPFPauseAsym {
		flags = append(flags, "\"OFPPFPauseAsym\": true")
	}
	return []byte("{" + strings.Join(flags, ", ") + "}"), nil
}

const (
	// Identifiers from group ofp_port_mod_failed_code
	OFPPMFCBadPort   = 0 // OFPPMFC_BAD_PORT
	OFPPMFCBadHwAddr = 1 // OFPPMFC_BAD_HW_ADDR
)

type PortModFailedCode uint16

func (self PortModFailedCode) MarshalJSON() ([]byte, error) {
	switch self {
	case OFPPMFCBadPort:
		return []byte("\"OFPPMFCBadPort\""), nil
	case OFPPMFCBadHwAddr:
		return []byte("\"OFPPMFCBadHwAddr\""), nil
	default:
		return nil, fmt.Errorf("Invalid value '%d' for PortModFailedCode", self)
	}
}

const (
	// Identifiers from group ofp_port_reason
	OFPPRAdd    = 0 // OFPPR_ADD
	OFPPRDelete = 1 // OFPPR_DELETE
	OFPPRModify = 2 // OFPPR_MODIFY
)

type PortReason uint8

func (self PortReason) MarshalJSON() ([]byte, error) {
	switch self {
	case OFPPRAdd:
		return []byte("\"OFPPRAdd\""), nil
	case OFPPRDelete:
		return []byte("\"OFPPRDelete\""), nil
	case OFPPRModify:
		return []byte("\"OFPPRModify\""), nil
	default:
		return nil, fmt.Errorf("Invalid value '%d' for PortReason", self)
	}
}

const (
	// Identifiers from group ofp_port_state
	OFPPSLinkDown   = 1   // OFPPS_LINK_DOWN
	OFPPSSTPListen  = 0   // OFPPS_STP_LISTEN
	OFPPSSTPLearn   = 256 // OFPPS_STP_LEARN
	OFPPSSTPForward = 512 // OFPPS_STP_FORWARD
	OFPPSSTPBlock   = 768 // OFPPS_STP_BLOCK
	OFPPSSTPMask    = 768 // OFPPS_STP_MASK
)

type PortState uint32

func (self PortState) MarshalJSON() ([]byte, error) {
	var flags []string
	if self&OFPPSLinkDown == OFPPSLinkDown {
		flags = append(flags, "\"OFPPSLinkDown\": true")
	}
	if self&OFPPSSTPListen == OFPPSSTPListen {
		flags = append(flags, "\"OFPPSSTPListen\": true")
	}
	if self&OFPPSSTPLearn == OFPPSSTPLearn {
		flags = append(flags, "\"OFPPSSTPLearn\": true")
	}
	if self&OFPPSSTPForward == OFPPSSTPForward {
		flags = append(flags, "\"OFPPSSTPForward\": true")
	}
	if self&OFPPSSTPBlock == OFPPSSTPBlock {
		flags = append(flags, "\"OFPPSSTPBlock\": true")
	}
	if self&OFPPSSTPMask == OFPPSSTPMask {
		flags = append(flags, "\"OFPPSSTPMask\": true")
	}
	return []byte("{" + strings.Join(flags, ", ") + "}"), nil
}

const (
	// Identifiers from group ofp_queue_op_failed_code
	OFPQOFCBadPort  = 0 // OFPQOFC_BAD_PORT
	OFPQOFCBadQueue = 1 // OFPQOFC_BAD_QUEUE
	OFPQOFCEperm    = 2 // OFPQOFC_EPERM
)

type QueueOpFailedCode uint16

func (self QueueOpFailedCode) MarshalJSON() ([]byte, error) {
	switch self {
	case OFPQOFCBadPort:
		return []byte("\"OFPQOFCBadPort\""), nil
	case OFPQOFCBadQueue:
		return []byte("\"OFPQOFCBadQueue\""), nil
	case OFPQOFCEperm:
		return []byte("\"OFPQOFCEperm\""), nil
	default:
		return nil, fmt.Errorf("Invalid value '%d' for QueueOpFailedCode", self)
	}
}

const (
	// Identifiers from group ofp_queue_properties
	OFPQTNone    = 0 // OFPQT_NONE
	OFPQTMinRate = 1 // OFPQT_MIN_RATE
)

type QueueProperties uint32

func (self QueueProperties) MarshalJSON() ([]byte, error) {
	switch self {
	case OFPQTNone:
		return []byte("\"OFPQTNone\""), nil
	case OFPQTMinRate:
		return []byte("\"OFPQTMinRate\""), nil
	default:
		return nil, fmt.Errorf("Invalid value '%d' for QueueProperties", self)
	}
}

const (
	// Identifiers from group ofp_stats_reply_flags
	OFPSFReplyMore = 1 // OFPSF_REPLY_MORE
)

type StatsReplyFlags uint16

func (self StatsReplyFlags) MarshalJSON() ([]byte, error) {
	var flags []string
	if self&OFPSFReplyMore == OFPSFReplyMore {
		flags = append(flags, "\"OFPSFReplyMore\": true")
	}
	return []byte("{" + strings.Join(flags, ", ") + "}"), nil
}

const (
// Identifiers from group ofp_stats_request_flags
)

type StatsRequestFlags uint16

func (self StatsRequestFlags) MarshalJSON() ([]byte, error) {
	var flags []string
	return []byte("{" + strings.Join(flags, ", ") + "}"), nil
}

const (
	// Identifiers from group ofp_stats_type
	OFPSTDesc         = 0     // OFPST_DESC
	OFPSTFlow         = 1     // OFPST_FLOW
	OFPSTAggregate    = 2     // OFPST_AGGREGATE
	OFPSTTable        = 3     // OFPST_TABLE
	OFPSTPort         = 4     // OFPST_PORT
	OFPSTQueue        = 5     // OFPST_QUEUE
	OFPSTExperimenter = 65535 // OFPST_EXPERIMENTER
)

type StatsType uint16

func (self StatsType) MarshalJSON() ([]byte, error) {
	switch self {
	case OFPSTDesc:
		return []byte("\"OFPSTDesc\""), nil
	case OFPSTFlow:
		return []byte("\"OFPSTFlow\""), nil
	case OFPSTAggregate:
		return []byte("\"OFPSTAggregate\""), nil
	case OFPSTTable:
		return []byte("\"OFPSTTable\""), nil
	case OFPSTPort:
		return []byte("\"OFPSTPort\""), nil
	case OFPSTQueue:
		return []byte("\"OFPSTQueue\""), nil
	case OFPSTExperimenter:
		return []byte("\"OFPSTExperimenter\""), nil
	default:
		return nil, fmt.Errorf("Invalid value '%d' for StatsType", self)
	}
}

const (
	// Identifiers from group ofp_type
	OFPTHello                 = 0  // OFPT_HELLO
	OFPTError                 = 1  // OFPT_ERROR
	OFPTEchoRequest           = 2  // OFPT_ECHO_REQUEST
	OFPTEchoReply             = 3  // OFPT_ECHO_REPLY
	OFPTExperimenter          = 4  // OFPT_EXPERIMENTER
	OFPTFeaturesRequest       = 5  // OFPT_FEATURES_REQUEST
	OFPTFeaturesReply         = 6  // OFPT_FEATURES_REPLY
	OFPTGetConfigRequest      = 7  // OFPT_GET_CONFIG_REQUEST
	OFPTGetConfigReply        = 8  // OFPT_GET_CONFIG_REPLY
	OFPTSetConfig             = 9  // OFPT_SET_CONFIG
	OFPTPacketIn              = 10 // OFPT_PACKET_IN
	OFPTFlowRemoved           = 11 // OFPT_FLOW_REMOVED
	OFPTPortStatus            = 12 // OFPT_PORT_STATUS
	OFPTPacketOut             = 13 // OFPT_PACKET_OUT
	OFPTFlowMod               = 14 // OFPT_FLOW_MOD
	OFPTPortMod               = 15 // OFPT_PORT_MOD
	OFPTStatsRequest          = 16 // OFPT_STATS_REQUEST
	OFPTStatsReply            = 17 // OFPT_STATS_REPLY
	OFPTBarrierRequest        = 18 // OFPT_BARRIER_REQUEST
	OFPTBarrierReply          = 19 // OFPT_BARRIER_REPLY
	OFPTQueueGetConfigRequest = 20 // OFPT_QUEUE_GET_CONFIG_REQUEST
	OFPTQueueGetConfigReply   = 21 // OFPT_QUEUE_GET_CONFIG_REPLY
)

type Type uint8

func (self Type) MarshalJSON() ([]byte, error) {
	switch self {
	case OFPTHello:
		return []byte("\"OFPTHello\""), nil
	case OFPTError:
		return []byte("\"OFPTError\""), nil
	case OFPTEchoRequest:
		return []byte("\"OFPTEchoRequest\""), nil
	case OFPTEchoReply:
		return []byte("\"OFPTEchoReply\""), nil
	case OFPTExperimenter:
		return []byte("\"OFPTExperimenter\""), nil
	case OFPTFeaturesRequest:
		return []byte("\"OFPTFeaturesRequest\""), nil
	case OFPTFeaturesReply:
		return []byte("\"OFPTFeaturesReply\""), nil
	case OFPTGetConfigRequest:
		return []byte("\"OFPTGetConfigRequest\""), nil
	case OFPTGetConfigReply:
		return []byte("\"OFPTGetConfigReply\""), nil
	case OFPTSetConfig:
		return []byte("\"OFPTSetConfig\""), nil
	case OFPTPacketIn:
		return []byte("\"OFPTPacketIn\""), nil
	case OFPTFlowRemoved:
		return []byte("\"OFPTFlowRemoved\""), nil
	case OFPTPortStatus:
		return []byte("\"OFPTPortStatus\""), nil
	case OFPTPacketOut:
		return []byte("\"OFPTPacketOut\""), nil
	case OFPTFlowMod:
		return []byte("\"OFPTFlowMod\""), nil
	case OFPTPortMod:
		return []byte("\"OFPTPortMod\""), nil
	case OFPTStatsRequest:
		return []byte("\"OFPTStatsRequest\""), nil
	case OFPTStatsReply:
		return []byte("\"OFPTStatsReply\""), nil
	case OFPTBarrierRequest:
		return []byte("\"OFPTBarrierRequest\""), nil
	case OFPTBarrierReply:
		return []byte("\"OFPTBarrierReply\""), nil
	case OFPTQueueGetConfigRequest:
		return []byte("\"OFPTQueueGetConfigRequest\""), nil
	case OFPTQueueGetConfigReply:
		return []byte("\"OFPTQueueGetConfigReply\""), nil
	default:
		return nil, fmt.Errorf("Invalid value '%d' for Type", self)
	}
}