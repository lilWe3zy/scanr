# Scanr
![](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)

This is just a small CLI tool I wrote to speed up some tedious CSV editing that we have to do every now and again.
This is for filtering connections from a MikroTik CPE scan list.

## File breakdown

Within the file, we are given a list of entries in the form of:

```text
0F:0F:0F:0F:0F:0F,'SSID',5180/20/an,-48,nv2,privacy
```

**A breakdown of the elements are as follows:**

| Index | Name                       | Description                             | Example             |
|-------|----------------------------|-----------------------------------------|---------------------|
| 1     | MAC Address                | Unique access point identifier (BSSID)  | `0F:0F:0F:0F:0F:0F` |
| 2     | SSID                       | Network name (might be empty if hidden) | `Some&Where`        |
| 3     | Frequency/Channel/Standard | Frequency in MHz and channel width      | `5180/20/an`        |
| 4     | Signal Strength            | Signal strength in dBm                  | `-48`               |
| 5     | Protocol                   | Wireless protocol                       | `nv2`               |
| 6     | Security                   | Security type                           | `privacy`           |

**An index of channel flags present in index 3:**

| Flag       | Description                                                  |
|------------|--------------------------------------------------------------|
| `Ce`       | Control channel is lower, extension channel is above         |
| `eCee`     | Indicates 80 MHz channel width with specific bonding pattern |
| `Ceee`     | Control channel is at the start of a 160 MHz block           |
| `-` (None) | Standard 20 MHz channel, no bonding                          |

**An index for possible standards as found in index 3:**

| Code   | Standard  | Band        |
|--------|-----------|-------------|
| `a`    | 802.11a   | 5 GHz       |
| `n`    | 802.11n   | 2,4/5 GHz   |
| `ac`   | 802.11ac  | 5 Ghz       |
| `ax`   | 802.11ax  | 2.4/5/6 GHz |
| `b/g`  | 802.11b/g | 2.4 GHz     |


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
