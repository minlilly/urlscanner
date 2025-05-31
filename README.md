# URLScanner

```
 ██╗   ██╗██████╗ ██╗     ███████╗ ██████╗ █████╗ ███╗   ██╗███╗   ██╗███████╗██████╗ 
 ██║   ██║██╔══██╗██║     ██╔════╝██╔════╝██╔══██╗████╗  ██║████╗  ██║██╔════╝██╔══██╗
 ██║   ██║██████╔╝██║     ███████╗██║     ███████║██╔██╗ ██║██╔██╗ ██║█████╗  ██████╔╝
 ██║   ██║██╔══██╗██║     ╚════██║██║     ██╔══██║██║╚██╗██║██║╚██╗██║██╔══╝  ██╔══██╗
 ╚██████╔╝██║  ██║███████╗███████║╚██████╗██║  ██║██║ ╚████║██║ ╚████║███████╗██║  ██║
  ╚═════╝ ╚═╝  ╚═╝╚══════╝╚══════╝ ╚═════╝╚═╝  ╚═╝╚═╝  ╚═══╝╚═╝  ╚═══╝╚══════╝╚═╝  ╚═╝

           URLScan.io Powered Subdomain Discovery Tool
                           by: Psikoz
```

A fast and efficient subdomain discovery tool powered by URLScan.io API. URLScanner helps security researchers and penetration testers discover subdomains for reconnaissance purposes.

## Features

🔍 **Fast Subdomain Discovery** - Leverages URLScan.io's massive database  
📁 **Bulk Processing** - Process multiple domains from a file  
💾 **Output Options** - Save results to file or display in terminal  
🤫 **Silent Mode** - Perfect for automation and scripting  
🔑 **API Key Management** - Secure local storage of API credentials  
🎯 **Deduplication** - Automatic removal of duplicate results  

## Installation

### Method 1: Go Install (Recommended)
```bash
go install github.com/Psikoz-coder/urlscanner@latest
```

### Method 2: Manual Build
```bash
git clone https://github.com/Psikoz-coder/urlscanner.git
cd urlscanner
go build -o urlscanner
```

## Prerequisites

- Go 1.24
- URLScan.io API key (Get yours at [urlscan.io](https://urlscan.io/))

## Quick Start

1. **Set your URLScan.io API key:**
```bash
urlscanner -set YOUR_API_KEY_HERE
```

2. **Discover subdomains:**
```bash
urlscanner -d example.com
```

## Usage

### Basic Commands

```bash
# Set API key (required first time)
urlscanner -set <API_KEY>

# Discover subdomains for a single domain
urlscanner -d <domain>

# Process multiple domains from file
urlscanner -dL <file>

# Save results to file
urlscanner -d <domain> -o <output_file>

# Silent mode (no banner/verbose output)
urlscanner -d <domain> -silent
```

### Examples

```bash
# Basic subdomain discovery
urlscanner -d example.com

# Save results to file
urlscanner -d example.com -o subdomains.txt

# Process multiple domains and save results
urlscanner -dL domains.txt -o all_subdomains.txt

# Silent mode for automation
urlscanner -d example.com -silent

# Combined options
urlscanner -dL domains.txt -o results.txt -silent
```

### Input File Format

For bulk processing (`-dL` option), create a text file with one domain per line:

```
example.com
test.com
target.org
```

## Command Line Options

| Option | Description |
|--------|-------------|
| `-set <key>` | Set URLScan.io API key |
| `-d <domain>` | Target domain for subdomain discovery |
| `-dL <file>` | Process domains from file (one per line) |
| `-o <file>` | Save results to output file |
| `-silent` | Silent mode (no banner or verbose output) |

## Configuration

URLScanner stores your API key securely in `~/.urlscanner/config.json`. The configuration directory is created automatically when you first set your API key.

## Output

URLScanner outputs discovered subdomains in the following format:

```
[*] Scanning subdomains for: example.com
[+] Found 25 unique subdomains

api.example.com
blog.example.com
dev.example.com
mail.example.com
www.example.com
...

[+] Total unique subdomains found: 25
```

In silent mode, only the subdomain list is displayed:
```
api.example.com
blog.example.com
dev.example.com
```

## API Rate Limits

URLScan.io has rate limits for their API. URLScanner respects these limits, but for large-scale scanning, consider:

- Using the free tier limits responsibly
- Upgrading to a paid URLScan.io plan for higher limits
- Adding delays between requests for very large domain lists

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## Disclaimer

This tool is for educational and ethical security testing purposes only. Users are responsible for complying with applicable laws and obtaining proper authorization before scanning any domains they do not own.


## Author

**Psikoz** - *Initial work*

## Acknowledgments

- [URLScan.io](https://urlscan.io/) for providing the awesome API
- The Go community for excellent documentation and libraries

---

⭐ **Star this repository if you find it useful!**

🐛 **Found a bug?** [Open an issue](https://github.com/Psikoz-coder/urlscanner/issues)

💡 **Have a feature request?** [Start a discussion](https://github.com/Psikoz-coder/urlscanner/discussions)
