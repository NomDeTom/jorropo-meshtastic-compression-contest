package training
package training

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// StoreTrainingPacket writes a single packet to the training packets directory.
// It creates the directory if it doesn't exist.
func StoreTrainingPacket(trainPacketIndex uint, packetData []byte) error {
	filename := "train/packets/" + strconv.FormatUint(uint64(trainPacketIndex), 36)

storePacket:
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0622)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			err = os.MkdirAll("train/packets", 0755)
			if err != nil {
				return fmt.Errorf("creating train/packets directory: %w", err)
			}
			goto storePacket
		}
		return fmt.Errorf("creating individual packet file: %w", err)
	}
	defer f.Close()

	_, err = f.Write(packetData)
	if err != nil {
		return fmt.Errorf("writing individual packet file: %w", err)
	}

	return nil
}

// StorePortnumTrainingPayload writes a payload to a portnum-specific training directory.
// It creates the directory structure if it doesn't exist.
func StorePortnumTrainingPayload(portnum uint64, payloadIndex uint, payloadData []byte) error {
	dirname := "train/" + strconv.FormatUint(portnum, 10)
	filename := dirname + "/" + strconv.FormatUint(uint64(payloadIndex), 36)

storePayload:
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0622)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			err = os.MkdirAll(dirname, 0755)
			if err != nil {
				return fmt.Errorf("creating train/%d directory: %w", portnum, err)
			}
			goto storePayload
		}
		return fmt.Errorf("creating individual payload file: %w", err)
	}
	defer f.Close()

	_, err = f.Write(payloadData)
	if err != nil {
		return fmt.Errorf("writing individual payload file: %w", err)
	}

	return nil
}

// CleanupTrainingData removes all training data directories and files.
func CleanupTrainingData() error {
	return os.RemoveAll("train")
}
