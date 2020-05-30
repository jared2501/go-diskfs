package diskfs

const (
	logicalSectorSizeSyscall  = 0x40046418 // DKIOCGETBLOCKSIZE
	physicalSectorSizeSyscall = 0x4004644D // DKIOCGETPHYSICALBLOCKSIZE
)
