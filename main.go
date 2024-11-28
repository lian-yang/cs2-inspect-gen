package main

import (
	"flag"
	"fmt"

	"github.com/lian-yang/cs2-inspect-gen/econ_pb2"
)

var (
	rarity     uint
	defIndex   uint
	paintIndex uint
	paintSeed  uint
	paintWear  float64

	sticker1Id   uint
	sticker1Wear float64
	sticker2Id   uint
	sticker2Wear float64
	sticker3Id   uint
	sticker3Wear float64
	sticker4Id   uint
	sticker4Wear float64
	sticker5Id   uint
	sticker5Wear float64

	customName string

	stattrak      uint
	stattrakCount uint
)

func init() {
	// 品质
	flag.UintVar(&rarity, "r", 0, "rarity")
	// 物品编号
	flag.UintVar(&defIndex, "di", 1, "defindex")
	// 图案编号
	flag.UintVar(&paintIndex, "pi", 0, "paintindex")
	// 图案模板
	flag.UintVar(&paintSeed, "ps", 0, "paintseed")
	// 磨损度
	flag.Float64Var(&paintWear, "pw", 0, "paintwear")

	// 印花 + 磨损
	flag.UintVar(&sticker1Id, "si1", 0, "sticker 1 id")
	flag.Float64Var(&sticker1Wear, "sw1", 0, "sticker 1 wear")
	flag.UintVar(&sticker2Id, "si2", 0, "sticker 2 id")
	flag.Float64Var(&sticker2Wear, "sw2", 0, "sticker 2 wear")
	flag.UintVar(&sticker3Id, "si3", 0, "sticker 3 id")
	flag.Float64Var(&sticker3Wear, "sw3", 0, "sticker 3 wear")
	flag.UintVar(&sticker4Id, "si4", 0, "sticker 4 id")
	flag.Float64Var(&sticker4Wear, "sw4", 0, "sticker 4 wear")
	flag.UintVar(&sticker5Id, "si5", 0, "sticker 5 id")
	flag.Float64Var(&sticker5Wear, "sw5", 0, "sticker 5 wear")

	// 是否改名
	flag.StringVar(&customName, "cn", "", "customname")

	// 是否是stattrak
	flag.UintVar(&stattrak, "st", 1, "stattrak")
	flag.UintVar(&stattrakCount, "stc", 123, "stattrak count")

	flag.Parse()
}

func main() {
	protoMessage := &econ_pb2.CEconItemPreviewDataBlock{}
	protoMessage.Rarity = Uint32Ptr(rarity)
	protoMessage.Defindex = Uint32Ptr(defIndex)
	protoMessage.Paintindex = Uint32Ptr(paintIndex)
	protoMessage.Paintseed = Uint32Ptr(paintSeed)
	protoMessage.Paintwear = Wear2Uint32Ptr(paintWear)

	// 是否改名
	if customName != "" {
		protoMessage.Customname = &customName
	}

	if stattrak == 0 {
		protoMessage.Killeaterscoretype = Uint32Ptr(0)
		protoMessage.Killeatervalue = Uint32Ptr(stattrakCount)
	}

	var stickers string
	if sticker1Id != 0 {
		protoMessage.Stickers = append(protoMessage.Stickers, &econ_pb2.CEconItemPreviewDataBlock_Sticker{
			Slot:      Uint32Ptr(0),
			StickerId: Uint32Ptr(sticker1Id),
			Wear:      Float32Ptr(sticker1Wear),
		})

		stickers += fmt.Sprintf("%d %f ", sticker1Id, sticker1Wear)
	}

	if sticker2Id != 0 {
		protoMessage.Stickers = append(protoMessage.Stickers, &econ_pb2.CEconItemPreviewDataBlock_Sticker{
			Slot:      Uint32Ptr(1),
			StickerId: Uint32Ptr(sticker2Id),
			Wear:      Float32Ptr(sticker2Wear),
		})

		stickers += fmt.Sprintf("%d %f ", sticker2Id, sticker2Wear)
	}

	if sticker3Id != 0 {
		protoMessage.Stickers = append(protoMessage.Stickers, &econ_pb2.CEconItemPreviewDataBlock_Sticker{
			Slot:      Uint32Ptr(2),
			StickerId: Uint32Ptr(sticker3Id),
			Wear:      Float32Ptr(sticker3Wear),
		})

		stickers += fmt.Sprintf("%d %f ", sticker3Id, sticker3Wear)
	}

	if sticker4Id != 0 {
		protoMessage.Stickers = append(protoMessage.Stickers, &econ_pb2.CEconItemPreviewDataBlock_Sticker{
			Slot:      Uint32Ptr(3),
			StickerId: Uint32Ptr(sticker4Id),
			Wear:      Float32Ptr(sticker4Wear),
		})

		stickers += fmt.Sprintf("%d %f ", sticker4Id, sticker4Wear)
	}

	if sticker5Id != 0 {
		protoMessage.Stickers = append(protoMessage.Stickers, &econ_pb2.CEconItemPreviewDataBlock_Sticker{
			Slot:      Uint32Ptr(4),
			StickerId: Uint32Ptr(sticker5Id),
			Wear:      Float32Ptr(sticker5Wear),
		})

		stickers += fmt.Sprintf("%d %f ", sticker5Id, sticker5Wear)
	}

	generatedPayload, err := GenerateInspect(protoMessage)
	if err != nil {
		panic(err)
	}

	fmt.Printf("csgo_econ_action_preview %s\nsteam://run/730/en/+csgo_econ_action_preview%%20%s\n!gen %d %d %d %v %s\n", generatedPayload, generatedPayload, defIndex, paintIndex, paintSeed, paintWear, stickers)

}
