package MiniCache

// ByteView 保存字节的不可变视图
type ByteView struct {
	b []byte
}

// Len 返回视图的长度
func (v ByteView) Len() int {
	return len(v.b)
}

// ByteSlice 以直接切片方式返回一个数据复制体
func (v ByteView) ByteSlice() []byte {
	return cloneBytes(v.b)
}

// string 将数据以字符串方式保存
func (v ByteView) String() string {
	return string(v.b)
}

func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
