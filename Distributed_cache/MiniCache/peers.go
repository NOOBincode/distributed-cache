package MiniCache

// PeerPicker 是被定位必须实现的接口
// 每一个peer 都拥有其特有的key
type PeerPicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool)
}

// PeerGetter 是必须被peer实现的接口
type PeerGetter interface {
	Get(group string, key string) ([]byte, error)
}
