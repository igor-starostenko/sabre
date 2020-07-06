# sabre
Slice your files like a champ with **sabre**

## Setup

Install [go](https://golang.org/doc/install).

Clone the repo:
```
git clone https://github.com/igor-starostenko/sabre.git
```

Install:
```
go install
```
It compiles the module and makes **sabre** executable available under `~/go/bin`

## Usage

```
usage: sabre [options] SOURCE OUTPUT

  -e string
    	Output file extension (default "txt")
  -l int
    	Max lines sliced per file (default 100)
  -q	Supress informational output
  -v	Print version info about sabre and exit
```
where `SOURCE` is the source file to be sliced
and `OUTPUT` is a directory to save the slices to

## Contributing

Bug reports and pull requests are welcome on GitHub at https://github.com/igor-starostenko/sabre. This project is intended to be a safe, welcoming space for collaboration, and contributors are expected to adhere to the [Contributor Covenant](http://contributor-covenant.org) code of conduct.

## License

The application is available as open source under the terms of the [MIT License](https://opensource.org/licenses/MIT).
