# Scanr
![](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)

This is just a small CLI tool I wrote to speed up some tedious CSV editing that we have to do every now and again.
This is for filtering connections from a MikroTik CPE scan list.

## Installation

Prebuilt binaries can be found in the releases tab

_Currently only have a build for windows, will get to multiplatform support as soon as I can_
## Usage

```shell
scanr -f<filePath> -s<number> -o<filepath> -n<name>
```

_By default, the program will look in the root directory for a file named `t.csv`, filter out all connections with a name starts with `RN`, and output all towers below `-55 dBm` into `scanlist.csv`_

Where:

 - ### `-f (file)`:
   - `string` Filepath to input csv file
   - **Default** `t.csv`
   

 - ### `-s (signal)`:
   - `int` Signed integer value indicating the max dBm floor to cut off searching
   - **Default** `-55`
   - **Note** This value should explicitly be called with a negative sign
 
- ### `-o (outdir)`:
  - `string` Filepath to the desired output directory
  - **Default** `scanlist.csv`
  - **Note** Will create file if it does not exist
  - **Note** Will overwrite file if it does exist

- ### `-n (name)`:
  - `string` Substring value to filter
  - **Default** `RN`

>[!NOTE]
> This tool is hardcoded to deal with RapidNetworks Tower naming, but I will add support for custom names as soon as I can
