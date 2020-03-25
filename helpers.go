package bitcoind

// handleError handle error returned by client.call
func handleError(err error, r *rpcResponse) error {
	if err != nil {
		return err
	}
	return nil
}
