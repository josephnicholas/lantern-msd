# lantern-msd

Multi Source File Downloader

## Description

This is a simple Go program that downloads files from multiple sources concurrently.

## Dependencies

| Name           | Version |
|----------------|---------|
| Go             | 1.22    |
| urfave/cli/v2  | 2.27.1  |
| nbio/st        | 0.0.0   |
| vbauerster/mpb | 3.4.0   |

## Building
```shell
go build lantern-msd
```

## Testing
```shell
go test -v lantern-msd/test
```

## Usage
1. View the help message for usage instructions.
    ```shell
    ./lantern-msd help

    NAME:
       lantern-msd - A simple multi-source downloader

    USAGE:
       lantern-msd [global options] command [command options]

    COMMANDS:
       help, h  Shows a list of commands or help for one command

    GLOBAL OPTIONS:
       --url value, -u value [ --url value, -u value ]  URLs to download
       --split value, -s value                          Download a file using N number of chunks. This will also determine the number of concurrent workers. (default: 5)
       --min-split-size value, -k value                 Minimum size of a single chunk in bytes. If the size of the file is less than this value, it will be downloaded in a single chunk. (default: 20MB)
       --help, -h                                       show help
    ```
2. Download a file from multiple resources.
    ```shell
    ./lantern-msd --url https://mirror.sitsa.com.ar/ubuntu-releases/noble/ubuntu-24.04-desktop-amd64.iso --url https://releases.ubuntu.com/noble/ubuntu-24.04-desktop-amd64.iso 
   
      100.00 MiB / 100.00 MiB [mirror.sitsa.com.ar] [==============================================================================] 00:00 ] 481.02 KiB/s
      100.00 MiB / 100.00 MiB [mirror.sitsa.com.ar] [==============================================================================] 00:00 ] 481.02 KiB/s
      100.00 MiB / 100.00 MiB [mirror.sitsa.com.ar] [==============================================================================] 00:00 ] 481.02 KiB/s
      100.00 MiB / 100.00 MiB [releases.ubuntu.com] [==============================================================================] 00:00 ] 1.07 MiB/s
      100.00 MiB / 100.00 MiB [mirror.sitsa.com.ar] [==============================================================================] 00:00 ] 298.68 KiB/s
      100.00 MiB / 100.00 MiB [releases.ubuntu.com] [==============================================================================] 00:00 ] 1000.39 KiB/s
      100.00 MiB / 100.00 MiB [mirror.sitsa.com.ar] [==============================================================================] 00:00 ] 286.83 KiB/s
      100.00 MiB / 100.00 MiB [releases.ubuntu.com] [==============================================================================] 00:00 ] 1.04 MiB/s
      100.00 MiB / 100.00 MiB [mirror.sitsa.com.ar] [==============================================================================] 00:00 ] 454.87 KiB/s
      100.00 MiB / 100.00 MiB [releases.ubuntu.com] [==============================================================================] 00:00 ] 1.12 MiB/s
      100.00 MiB / 100.00 MiB [mirror.sitsa.com.ar] [==============================================================================] 00:00 ] 319.36 KiB/s
      100.00 MiB / 100.00 MiB [releases.ubuntu.com] [==============================================================================] 00:00 ] 1.10 MiB/s
      100.00 MiB / 100.00 MiB [mirror.sitsa.com.ar] [==============================================================================] 00:00 ] 290.74 KiB/s
      100.00 MiB / 100.00 MiB [releases.ubuntu.com] [==============================================================================] 00:00 ] 1.05 MiB/s
      100.00 MiB / 100.00 MiB [mirror.sitsa.com.ar] [==============================================================================] 00:00 ] 301.85 KiB/s
      100.00 MiB / 100.00 MiB [releases.ubuntu.com] [==============================================================================] 00:00 ] 1.04 MiB/s
      100.00 MiB / 100.00 MiB [mirror.sitsa.com.ar] [==============================================================================] 00:00 ] 291.64 KiB/s
      100.00 MiB / 100.00 MiB [releases.ubuntu.com] [==============================================================================] 00:00 ] 1.17 MiB/s
      100.00 MiB / 100.00 MiB [mirror.sitsa.com.ar] [==============================================================================] 00:00 ] 332.04 KiB/s
      100.00 MiB / 100.00 MiB [releases.ubuntu.com] [==============================================================================] 00:00 ] 1.05 MiB/s
      100.00 MiB / 100.00 MiB [mirror.sitsa.com.ar] [==============================================================================] 00:00 ] 274.91 KiB/s
      100.00 MiB / 100.00 MiB [releases.ubuntu.com] [==============================================================================] 00:00 ] 1003.92 KiB/s
      100.00 MiB / 100.00 MiB [mirror.sitsa.com.ar] [==============================================================================] 00:00 ] 281.44 KiB/s
      100.00 MiB / 100.00 MiB [releases.ubuntu.com] [==============================================================================] 00:00 ] 1.04 MiB/s
      100.00 MiB / 100.00 MiB [mirror.sitsa.com.ar] [==============================================================================] 00:00 ] 297.43 KiB/s
      100.00 MiB / 100.00 MiB [releases.ubuntu.com] [==============================================================================] 00:00 ] 1.11 MiB/s
      100.00 MiB / 100.00 MiB [mirror.sitsa.com.ar] [==============================================================================] 00:00 ] 291.94 KiB/s
      100.00 MiB / 100.00 MiB [releases.ubuntu.com] [==============================================================================] 00:00 ] 1.03 MiB/s
      100.00 MiB / 100.00 MiB [mirror.sitsa.com.ar] [==============================================================================] 00:00 ] 300.47 KiB/s
      100.00 MiB / 100.00 MiB [releases.ubuntu.com] [==============================================================================] 00:00 ] 991.09 KiB/s
      100.00 MiB / 100.00 MiB [mirror.sitsa.com.ar] [==============================================================================] 00:00 ] 247.99 KiB/s
      2.86 GiB / 2.86 GiB [releases.ubuntu.com] [==============================================================================] 00:00 ] 4.15 MiB/s 
    ```
3. To specify the number of chunks to download a file with, use the `--split` flag.
    ```shell
    ./lantern-msd --url https://mirror.sitsa.com.ar/ubuntu-releases/noble/ubuntu-24.04-desktop-amd64.iso --split 10
    ```
4. To specify the minimum size of a single chunk in bytes, use the `--min-split-size` flag.
    ```shell
    ./lantern-msd --url https://mirror.sitsa.com.ar/ubuntu-releases/noble/ubuntu-24.04-desktop-amd64.iso --min-split-size 10MB
    ```
   
## TODO(Nice to haves)
1. Add support for downloading files from different sources FTP, BitTorrent, etc.
2. Much more stable and reliable download and UI progress bar representation.
3. Add support for resuming downloads.
4. Add re-download support for failed downloads.
5. Add support for verifying the integrity of downloaded files.
6. Add support for reading configuration and urls from a file.
7. Add support for using single progress through CLI flag.
   




