# Phinances

Just a stupid little program to play with double entry accounting. I want it to
have pretty graphs, and to use sqlcipher so that the "save file" is not in
plaintext.

## Directory structure

cmd/phinance/

    Executable, calls into main package

internal/

    Code that shouldn't be used by other projects

## Usage

Set the following environment variables:

- PHINANCES_DATABASE_PASSWORD: The password to use for the database
