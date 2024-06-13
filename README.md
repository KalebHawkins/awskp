# awskp - AWS Keypair Retriever
This is a command-line tool written in Go to retrieve the private key associated with an AWS EC2 keypair.
This is useful when you want to generate your keypairs using AWS CDK.

## Features
* Retrieves both private keys from an AWS keypair.
* Outputs the private key to the console or a specified file.

## Prerequisites
* Go programming language installed (https://go.dev/doc/install)
* AWS CLI configured with appropriate permissions (https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-welcome.html)

## Installation

Download or clone this repository.

```
git clone https//github.com/KalebHawkins/awskp.git
cd awskp/
```

Build the tool using 

```
go build . -o awskp
```

## Usage

```
awskp -r <region> -k <key-id> [-o <outfile>]

Arguments

-r, --region: The AWS region where the keypair resides (required)
-k, --key-id: The name of the keypair (required)
-o, --outfile: (Optional) The file to write the private key to. If not provided, the key will be printed to the console.
```

## Example

Retrieve the private key for a keypair named "my-keypair" in the us-east-1 region and write it to a file named "key.pem":

```
awskp -r us-east-1 -k my-keypair -o key.pem
```

## License

This project is licensed under the Apache License, Version 2.0. See the LICENSE file for details.