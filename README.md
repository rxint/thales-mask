# Thales Tokenization & Detokenization with Masking 

Masking the last two digits (or any other part) of sensitive information is common in scenarios where partial data exposure is required (for example, showing only the last 4 digits of a credit card for verification while keeping the rest hidden). This ensures compliance with security standards (such as PCI DSS) and protects sensitive information from unauthorized access.

This Go application interacts with Thales CipherTrust's `cte` command-line tool to perform tokenization and detokenization. After detokenizing, it masks the last two digits of the result for security purposes (e.g., replacing them with `**`).

## Features
- **Detokenization**: Convert tokens back to their original values using the `cte detokenize` command.
- **Masking**: Mask the last two digits of the detokenized value with `**`.

## Prerequisites
- Thales CipherTrust's `cte` tool installed and available in your system's `PATH`.
- Go installed.

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/thales-tokenize-detokenize-go.git
   cd thales-tokenize-detokenize-go

2. Build the Go application:
   ```bash
   go build
Usage
To detokenize a value and mask the last two digits:

1. Run the application with your token:
   ```bash
   ./thales-tokenize-detokenize-go

Configuration:
Replace the <TOKEN> in the main.go file with your actual token.
Adjust the tokengroup and format parameters if necessary.
