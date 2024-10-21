package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/pin/tftp"
)

func main() {
	// Defining flags for the parameters
	operation := flag.String("op", "get", "Operation (get or put)")
	server := flag.String("server", "127.0.0.1:69", "TFTP server address")
	remoteFile := flag.String("remote", "", "File name on the server")
	localFile := flag.String("local", "", "Local file name")

	flag.Parse()

	// Check if both remote and local file names are provided
	if *remoteFile == "" || *localFile == "" {
		log.Fatalf("Please provide both remote and local file names.")
	}

	// Connect to the TFTP server
	c, err := tftp.NewClient(*server)
	if err != nil {
		log.Fatalf("Failed to connect to TFTP server: %v", err)
	}

	// Determine the operation: GET or PUT
	if *operation == "get" {
		// GET operation (download file from server)
		downloadFile(c, *remoteFile, *localFile)
	} else if *operation == "put" {
		// PUT operation (upload file to server)
		uploadFile(c, *remoteFile, *localFile)
	} else {
		log.Fatalf("Invalid operation: use 'get' or 'put'.")
	}
}

// Function to download a file (GET)
func downloadFile(c *tftp.Client, remoteFile string, localFile string) {
	r, err := c.Receive(remoteFile, "octet")
	if err != nil {
		log.Fatalf("Failed to download the file: %v", err)
	}

	// Create the local file
	local, err := os.Create(localFile)
	if err != nil {
		log.Fatalf("Failed to create local file: %v", err)
	}
	defer local.Close()

	// Write data to the local file
	if _, err := r.WriteTo(local); err != nil {
		log.Fatalf("Error writing to local file: %v", err)
	}

	fmt.Printf("File '%s' successfully downloaded as '%s'\n", remoteFile, localFile)
}

// Function to upload a file (PUT)
func uploadFile(c *tftp.Client, remoteFile string, localFile string) {
	local, err := os.Open(localFile)
	if err != nil {
		log.Fatalf("Failed to open local file: %v", err)
	}
	defer local.Close()

	// Send the file to the server
	w, err := c.Send(remoteFile, "octet")
	if err != nil {
		log.Fatalf("Failed to upload the file: %v", err)
	}

	// Read the local file and send it to the server
	if _, err := w.ReadFrom(local); err != nil {
		log.Fatalf("Error uploading the file: %v", err)
	}

	fmt.Printf("File '%s' successfully uploaded as '%s'\n", localFile, remoteFile)
}
