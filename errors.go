package bitcoind

import "fmt"

//! Standard JSON-RPC 2.0 errors
// RPC_INVALID_REQUEST is internally mapped to HTTP_BAD_REQUEST (400).
// It should not be used for application-layer errors.
var ErrInvalidRequest = fmt.Errorf("RPC_INVALID_REQUEST") // -32600,
// RPC_METHOD_NOT_FOUND is internally mapped to HTTP_NOT_FOUND (404).
// It should not be used for application-layer errors.
var ErrMethodNotFound = fmt.Errorf("RPC_METHOD_NOT_FOUND") // -32601,
var ErrInvalidParams = fmt.Errorf("RPC_INVALID_PARAMS")    // -32602,
// RPC_INTERNAL_ERROR should only be used for genuine errors in bitcoind
// (for example datadir corruption).
var ErrInternalError = fmt.Errorf("RPC_INTERNAL_ERROR") // -32603,
var ErrParseError = fmt.Errorf("RPC_PARSE_ERROR")       // -32700,

//! General application defined errors
var ErrMiscError = fmt.Errorf("RPC_MISC_ERROR")                         // -1,  //!< std::exception thrown in command handling
var ErrTypeError = fmt.Errorf("RPC_TYPE_ERROR")                         // -3,  //!< Unexpected type was passed as parameter
var ErrInvalidAddressOrKey = fmt.Errorf("RPC_INVALID_ADDRESS_OR_KEY")   // -5,  //!< Invalid address or key
var ErrOutOfMemory = fmt.Errorf("RPC_OUT_OF_MEMORY")                    // -7,  //!< Ran out of memory during operation
var ErrInvalidParameter = fmt.Errorf("RPC_INVALID_PARAMETER")           // -8,  //!< Invalid, missing or duplicate parameter
var ErrDatabaseError = fmt.Errorf("RPC_DATABASE_ERROR")                 // -20, //!< Database error
var ErrDeserialzationError = fmt.Errorf("RPC_DESERIALIZATION_ERROR")    // -22, //!< Error parsing or validating structure in raw format
var ErrVerifyError = fmt.Errorf("RPC_VERIFY_ERROR")                     // -25, //!< General error during transaction or block submission
var ErrVerifyRejected = fmt.Errorf("RPC_VERIFY_REJECTED")               // -26, //!< Transaction or block was rejected by network rules
var ErrVerifyAlreadyInChain = fmt.Errorf("RPC_VERIFY_ALREADY_IN_CHAIN") // -27, //!< Transaction already in chain
var ErrInWarmup = fmt.Errorf("RPC_IN_WARMUP")                           // -28, //!< Client still warming up
var ErrMethodDepricated = fmt.Errorf("RPC_METHOD_DEPRECATED")           // -32, //!< RPC method is deprecated

//! P2P client errors
var ErrClientNotConnected = fmt.Errorf("RPC_CLIENT_NOT_CONNECTED")             // -9,  //!< Bitcoin is not connected
var ErrClientInInitialDownload = fmt.Errorf("RPC_CLIENT_IN_INITIAL_DOWNLOAD")  // -10, //!< Still downloading initial blocks
var ErrClientNodeAlreadyAdded = fmt.Errorf("RPC_CLIENT_NODE_ALREADY_ADDED")    // -23, //!< Node is already added
var ErrClientNodeNotAdded = fmt.Errorf("RPC_CLIENT_NODE_NOT_ADDED")            // -24, //!< Node has not been added before
var ErrClientNodeNotConnected = fmt.Errorf("RPC_CLIENT_NODE_NOT_CONNECTED")    // -29, //!< Node to disconnect not found in connected nodes
var ErrClientInvalidIPOrSubnet = fmt.Errorf("RPC_CLIENT_INVALID_IP_OR_SUBNET") // -30, //!< Invalid IP/Subnet
var ErrClientP2PDisabled = fmt.Errorf("RPC_CLIENT_P2P_DISABLED")               // -31, //!< No valid connection manager instance found

//! Wallet errors
var ErrWalletError = fmt.Errorf("RPC_WALLET_ERROR")                              // -4,  //!< Unspecified problem with wallet (key not found etc.)
var ErrInsufficientFunds = fmt.Errorf("RPC_WALLET_INSUFFICIENT_FUNDS")           // -6,  //!< Not enough funds in wallet or account
var ErrWalletInvalidLabelName = fmt.Errorf("RPC_WALLET_INVALID_LABEL_NAME")      // -11, //!< Invalid label name
var ErrWalletKeypoolRanOut = fmt.Errorf("RPC_WALLET_KEYPOOL_RAN_OUT")            // -12, //!< Keypool ran out, call keypoolrefill first
var ErrWalletUnlockNeeded = fmt.Errorf("RPC_WALLET_UNLOCK_NEEDED")               // -13, //!< Enter the wallet passphrase with walletpassphrase first
var ErrWalletPassphraseIncorrect = fmt.Errorf("RPC_WALLET_PASSPHRASE_INCORRECT") // -14, //!< The wallet passphrase entered was incorrect
var ErrWalletWrongEncState = fmt.Errorf("RPC_WALLET_WRONG_ENC_STATE")            // -15, //!< Command given in wrong wallet encryption state (encrypting an encrypted wallet etc.)
var ErrWalletEncryptionFailed = fmt.Errorf("RPC_WALLET_ENCRYPTION_FAILED")       // -16, //!< Failed to encrypt the wallet
var ErrWalletAlreadyUnlocked = fmt.Errorf("RPC_WALLET_ALREADY_UNLOCKED")         // -17, //!< Wallet is already unlocked
var ErrWalletNotFound = fmt.Errorf("RPC_WALLET_NOT_FOUND")                       // -18, //!< Invalid wallet specified
var ErrWalletNotSpecified = fmt.Errorf("RPC_WALLET_NOT_SPECIFIED")               // -19, //!< No wallet specified (error when there are multiple wallets loaded)

var ErrMap = map[int16]error{
	-32600: ErrInvalidRequest,
	-32601: ErrMethodNotFound,
	-32602: ErrInvalidParams,
	-32603: ErrInternalError,
	-32700: ErrParseError,
	-1:     ErrMiscError,
	-3:     ErrTypeError,
	-5:     ErrInvalidAddressOrKey,
	-7:     ErrOutOfMemory,
	-8:     ErrInvalidParameter,
	-20:    ErrDatabaseError,
	-22:    ErrDeserialzationError,
	-25:    ErrVerifyError,
	-26:    ErrVerifyRejected,
	-27:    ErrVerifyAlreadyInChain,
	-28:    ErrInWarmup,
	-32:    ErrMethodDepricated,
	-9:     ErrClientNotConnected,
	-10:    ErrClientInInitialDownload,
	-23:    ErrClientNodeAlreadyAdded,
	-24:    ErrClientNodeNotAdded,
	-29:    ErrClientNodeNotConnected,
	-30:    ErrClientInvalidIPOrSubnet,
	-31:    ErrClientP2PDisabled,
	-4:     ErrWalletError,
	-6:     ErrInsufficientFunds,
	-11:    ErrWalletInvalidLabelName,
	-12:    ErrWalletKeypoolRanOut,
	-13:    ErrWalletUnlockNeeded,
	-14:    ErrWalletPassphraseIncorrect,
	-15:    ErrWalletWrongEncState,
	-16:    ErrWalletEncryptionFailed,
	-17:    ErrWalletAlreadyUnlocked,
	-18:    ErrWalletNotFound,
	-19:    ErrWalletNotSpecified,
}
