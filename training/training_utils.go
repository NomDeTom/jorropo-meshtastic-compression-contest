package training
// Training utilities and helper functions
// This file contains functions extracted from main.go for organizing training-related code
// Place this file in the main package alongside main.go

package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// storeTrainingPacket writes a raw packet to the training packets directory.
// Used when generateTrainingDataset is enabled.
func storeTrainingPacket(trainPacketIndex uint, packetData []byte) error {
	filename := "train/packets/" + strconv.FormatUint(uint64(trainPacketIndex), 36)

retry:
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0622)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			err = os.MkdirAll("train/packets", 0755)
			if err != nil {
				return fmt.Errorf("creating train/packets directory: %w", err)
			}
			goto retry
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

// storePortnumTrainingPayload writes a portnum-specific payload to the training directory.
// Used when generateTrainingDataset is enabled.
func storePortnumTrainingPayload(portnum uint64, payloadIndex uint, payloadData []byte) error {
	dirname := "train/" + strconv.FormatUint(portnum, 10)
	filename := dirname + "/" + strconv.FormatUint(uint64(payloadIndex), 36)

retry:
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0622)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			err = os.MkdirAll(dirname, 0755)
			if err != nil {
				return fmt.Errorf("creating train/%d directory: %w", portnum, err)
			}
			goto retry
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

// cleanupTrainingData removes all training data directories and files.
// Useful for cleaning up between training runs.
func cleanupTrainingData() error {
	return os.RemoveAll("train")
}
