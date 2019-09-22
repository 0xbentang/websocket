// Code generated by "stringer -type=StatusCode"; DO NOT EDIT.

package websocket

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[StatusNormalClosure-1000]
	_ = x[StatusGoingAway-1001]
	_ = x[StatusProtocolError-1002]
	_ = x[StatusUnsupportedData-1003]
	_ = x[StatusNoStatusRcvd-1005]
	_ = x[StatusAbnormalClosure-1006]
	_ = x[StatusInvalidFramePayloadData-1007]
	_ = x[StatusPolicyViolation-1008]
	_ = x[StatusMessageTooBig-1009]
	_ = x[StatusMandatoryExtension-1010]
	_ = x[StatusInternalError-1011]
	_ = x[StatusServiceRestart-1012]
	_ = x[StatusTryAgainLater-1013]
	_ = x[StatusBadGateway-1014]
	_ = x[StatusTLSHandshake-1015]
}

const (
	_StatusCode_name_0 = "StatusNormalClosureStatusGoingAwayStatusProtocolErrorStatusUnsupportedData"
	_StatusCode_name_1 = "StatusNoStatusRcvdStatusAbnormalClosureStatusInvalidFramePayloadDataStatusPolicyViolationStatusMessageTooBigStatusMandatoryExtensionStatusInternalErrorStatusServiceRestartStatusTryAgainLaterStatusBadGatewayStatusTLSHandshake"
)

var (
	_StatusCode_index_0 = [...]uint8{0, 19, 34, 53, 74}
	_StatusCode_index_1 = [...]uint8{0, 18, 39, 68, 89, 108, 132, 151, 171, 190, 206, 224}
)

func (i StatusCode) String() string {
	switch {
	case 1000 <= i && i <= 1003:
		i -= 1000
		return _StatusCode_name_0[_StatusCode_index_0[i]:_StatusCode_index_0[i+1]]
	case 1005 <= i && i <= 1015:
		i -= 1005
		return _StatusCode_name_1[_StatusCode_index_1[i]:_StatusCode_index_1[i+1]]
	default:
		return "StatusCode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
