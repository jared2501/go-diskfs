package diskfs

const (
	logicalSectorSizeSyscall  = 0x1268     // blksszGet
	physicalSectorSizeSyscall = 0x80081270 // blkbszGet
)
