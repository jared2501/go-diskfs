package diskfs

import (
	"fmt"
	"golang.org/x/sys/unix"
	"os"
)

// to get the logical and physical sector sizes
func getBlockSizes(f *os.File) (int64, int64, int64, error) {
	fd := f.Fd()
	logicalSectorSize, err := unix.IoctlGetInt(int(fd), 0x1268) // blksszGet
	if err != nil {
		return 0, 0, 0, fmt.Errorf("Unable to get device logical sector size: %v", err)
	}
	physicalSectorSize, err := unix.IoctlGetInt(int(fd), 0x80081270) // blkbszGet
	if err != nil {
		return 0, 0, 0, fmt.Errorf("Unable to get device physical sector size: %v", err)
	}
	totalSize, err := unix.IoctlGetInt(int(fd), 0x80081272) // BLKGETSIZE64
	if err != nil {
		return 0, 0, 0, fmt.Errorf("Unable to get block count: %v", err)
	}
	return int64(totalSize), int64(logicalSectorSize), int64(physicalSectorSize), nil
}
