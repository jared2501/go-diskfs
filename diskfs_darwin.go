package diskfs

import (
	"fmt"
	"golang.org/x/sys/unix"
	"os"
)

// to get the logical and physical sector sizes
func getBlockSizes(f *os.File) (int64, int64, int64, error) {
	fd := f.Fd()
	logicalSectorSize, err := unix.IoctlGetInt(int(fd), 0x40046418) // DKIOCGETBLOCKSIZE
	if err != nil {
		return 0, 0, 0, fmt.Errorf("Unable to get device logical sector size: %v", err)
	}
	physicalSectorSize, err := unix.IoctlGetInt(int(fd), 0x4004644D) // DKIOCGETPHYSICALBLOCKSIZE
	if err != nil {
		return 0, 0, 0, fmt.Errorf("Unable to get device physical sector size: %v", err)
	}
	blockCount, err := unix.IoctlGetInt(int(fd), 0x40086419) // DKIOCGETBLOCKCOUNT
	if err != nil {
		return 0, 0, 0, fmt.Errorf("Unable to get block count: %v", err)
	}
	return int64(physicalSectorSize * blockCount), int64(logicalSectorSize), int64(physicalSectorSize), nil
}
