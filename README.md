# shachk
A simple sha256 compare tool. Checks and compares two sha256 sums and tells you if they're equal

## How to use it

Clone the repository and run `go install`

The command looks like this:

`shachk <path/to/downloaded/file/to/check> <url or shasum>`

As of now, the url can only go to either http or https urls. 
Can read all formats of text files where the shasum is the first thing there. Split by spaces, e.g.: <shasum256 *file.extension>
