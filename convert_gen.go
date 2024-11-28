package cs2_inspect_gen

import (
	"encoding/binary"
	"fmt"
	"hash/crc32"
	"math"
	"strings"

	"github.com/lian-yang/cs2-inspect-gen/econ_pb2"
	"google.golang.org/protobuf/proto"
)

func GenerateInspect(protoMessage *econ_pb2.CEconItemPreviewDataBlock) (string, error) {
	// 1. 序列化 Protobuf 消息
	serialized, err := proto.Marshal(protoMessage)
	if err != nil {
		return "", err
	}

	// 2. 创建一个字节缓冲区并在开头加一个 null byte
	buffer := append([]byte{0}, serialized...)

	// 3. 计算 CRC32 校验和
	crc := crc32.ChecksumIEEE(buffer)

	// 4. 执行 XOR 运算
	xorCrc := (crc & 0xFFFF) ^ (uint32(len(serialized)) * crc)

	// 5. 将 xorCrc 转换为 4 字节，并追加到缓冲区
	xorCrcBytes := make([]byte, 4)
	for i := 0; i < 4; i++ {
		xorCrcBytes[3-i] = byte(xorCrc >> (8 * i))
	}

	// 6. 将 CRC32 字节追加到缓冲区
	buffer = append(buffer, xorCrcBytes...)

	// 7. 返回大写的十六进制字符串
	return strings.ToUpper(fmt.Sprintf("%x", buffer)), nil
}

func Uint32Ptr(v uint) *uint32 {
	u32 := uint32(v)
	return &u32
}

func Float32Ptr(v float64) *float32 {
	f32 := float32(v)
	return &f32
}

func Wear2Uint32Ptr(wear float64) *uint32 {
	// 将浮动值转换为字节数组（大端字节序）
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, math.Float32bits(float32(wear)))

	// 将字节数组转换为整数
	v := binary.BigEndian.Uint32(buf)
	return &v
}
