package disk

// Darwin re-reads the partition table without any extra ioctls.
func (d *Disk) ReReadPartitionTable() error {
	return nil
}
