package utils

import "os"

func SaveImageToFile(image []byte, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(image)
	if err != nil {
		return err
	}

	// Flush the buffer to ensure the data is written to disk.
	err = file.Sync()
	if err != nil {
		return err
	}

	return nil
}
