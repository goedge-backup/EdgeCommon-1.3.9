// Copyright 2021 GoEdge CDN goedge.cdn@gmail.com. All rights reserved.

package nodeconfigs

type NodeLogType = string

const (
	NodeLogTypeListenAddressFailed    NodeLogType = "listenAddressFailed"
	NodeLogTypeServerConfigInitFailed NodeLogType = "serverConfigInitFailed"
	NodeLogTypeRunScriptFailed        NodeLogType = "runScriptFailed"
	NodeLogTypeScriptConsoleLog       NodeLogType = "scriptConsoleLog"
)
