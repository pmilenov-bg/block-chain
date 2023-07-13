package models

type BlockData struct {
	stringData string
}

func (blockData *BlockData) ExportToString() string {
	return blockData.stringData
}

func CreateBlockData(strData string) BlockData {
	blockData := BlockData{
		stringData: strData,
	}
	return blockData
}
