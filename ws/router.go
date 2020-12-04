package ws

func WebSocketRouterV1(c *Client) {
	c.Dispatch.Register("chainInfo", WsChainInfoHandler) // 链整体情况
	c.Dispatch.Register("blockInfo", WsChainInfoHandler) // 单区块信息
	c.Dispatch.Register("ssInfo",  WsChainInfoHandler)    // 跨分片交易组
	c.Dispatch.Register("tranInfo", WsChainInfoHandler)  // 交易信息
}
