# OUI

_Another OIU lookup tool_

## Overview

The OUI Lookup Tool is a command-line application written in Go, mainly for myself to experiment with writing a Go application.

It fetches the OUI database directly from IEEE<sup>[[1]](https://standards-oui.ieee.org/oui/oui.csv)</sup> and caches it locally for offline lookups.

There are probably a million tools that do something similar, but I made this one.

## Features

- **Update OUI Database**: Downloads the latest OUI data from IEEE and stores it locally.
- **Lookup Manufacturer**: Queries the local OUI database to find the manufacturer associated with a given MAC address prefix.

## Installation

1. Clone the repository:

```sh
git clone https://github.com/aliask/oui.git
cd oui
```

2. Build the application:

```sh
go build
sudo cp oui /usr/local/bin
```

## Usage

### Update the OUI Database

Before looking up any MAC addresses, you need to acquire the OUI database:

```sh
./oui update
```

This command downloads the latest OUI data from IEEE and stores it locally in your user configuration directory.

### Lookup a Manufacturer

To lookup the manufacturer for a given MAC address, use the `lookup` command followed by the MAC address:

```sh
./oui [lookup] ${MAC}
```

If `lookup` is omitted, the argument is assumed to be a MAC address.

Examples:

```sh
oui lookup 00:1A:2B:3C:4D:5E
Manufacturer for OUI 00:1A:2B:3C:4D:5E is: Ayecom Technology Co., Ltd.

oiu 00-24-E8
Manufacturer for OUI 00-24-E8 is: Dell Inc.

oiu lookup 00307E
Manufacturer for OUI 00307E is: Redflex Communication Systems
```

This command will output the manufacturer associated with the given MAC address prefix.

## Implementation Details

- Database Storage: The OUI database is stored in the user's configuration directory (typically `~/.config/oui_data.csv`).
- Data Format: The database is stored in CSV format, which is downloaded directly from the IEEE website.
- Input Validation: The tool ensures the provided MAC address is valid and conforms to one of the expected formats.
