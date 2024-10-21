# TFTP Client

A simple TFTP client written in Go, capable of performing file transfers using the TFTP protocol. This client supports both downloading (GET) and uploading (PUT) files to and from a TFTP server.

## Features

- **GET**: Download files from a TFTP server.
- **PUT**: Upload files to a TFTP server.

## Prerequisites

- Go version 1.16 or higher.
- A running TFTP server.

## Installation

To build the TFTP client, make sure Go is installed on your system and then run:

```bash
go build -o tftp_client main.go
```

This will create an executable named tftp_client.

## Usage
The tftp_client accepts several parameters to define the TFTP operation, server address, and file paths.

### Downloading a File (GET)

To download a file from a TFTP server, use the following command:

```bash
./tftp_client -op=get -server=<server_ip>:<port> -remote=<remote_file_name> -local=<local_file_name>
```

> Example:

```bash
./tftp_client -op=get -server=10.0.0.1:69 -remote=remote_file.efi -local=local_file.efi
    -op=get: Specifies the download operation.
    -server=<server_ip>:<port>: TFTP server IP and port (port is optional, default is 69).
    -remote=<remote_file_name>: The name of the file on the server.
    -local=<local_file_name>: The name of the file to be saved locally.
```

### Uploading a File (PUT)

To upload a file to a TFTP server, use the following command:

```bash
./tftp_client -op=put -server=<server_ip>:<port> -remote=<remote_file_name> -local=<local_file_name>
```

> Example:

```bash
./tftp_client -op=put -server=10.0.0.1:69 -remote=remote_file.efi -local=local_file.efi
    -op=put: Specifies the upload operation.
    -server=<server_ip>:<port>: TFTP server IP and port (port is optional, default is 69).
    -remote=<remote_file_name>: The name of the file on the server to be saved.
    -local=<local_file_name>: The name of the local file to be uploaded.
```

### Error Handling
If the file names (either remote or local) are not provided, the program will terminate with an error message.
If there is a failure to connect to the TFTP server, or if the transfer fails, the error will be logged, and the program will exit.

## License

This project is licensed under the MIT License. See the LICENSE file for more details.

## Author
[Leonardo Magdanello Araujo]