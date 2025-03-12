# Phinances

Just a stupid little program to play with double entry accounting and graphing. 

## Directory structure

cmd/phinance/

    Executable, calls into main package

internal/

    Code that shouldn't be used by other projects

## Usage

To start, create an encrypted disk image

Steps:
```
	1.	Open Disk Utility (Press Cmd + Space, type “Disk Utility”, and open it).
	2.	Click File → New Image → Blank Image.
	3.	In the “Save As” field, choose a location and name for your file.
	4.	Set:
        •	Size: 1GB
        •	Format: APFS (or Mac OS Extended (Journaled) if you need compatibility with older macOS versions).
        •	Encryption: AES-256-bit (Recommended) (or AES-128-bit for slightly faster performance).
        •	Partitions: Single Partition - GUID Partition Map
        •	Image Format: Sparse Bundle disk image (allows flexible resizing) or Read/Write disk image (.dmg).
	5.	Click Save, then enter a strong password.
```

Then, attach the disk, and point phinaces at the encrypted disk image volume
