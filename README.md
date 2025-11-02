# ByteRep

This is a simple command-line tool for encoding and decoding images to and from bytes, with an option to display images on the command-line using Unicode characters.

## Installation

Clone the repository and build the executable:

```bash
git clone https://github.com/ccianos/byterep.git
cd byterep
go build -o byterep
```

## Usage

### Encode an Image

Encode an image to bytes:

```bash
./byterep -encode -input input_image.jpg > encoded_image.txt
```

### Decode an Image

Decode bytes to an image:

```bash
./byterep -decode -input encoded_image.txt > decoded_image.jpg
```

### Display an Image using Unicode

Display an image using Unicode characters:

```bash
./byterep -decode -input encoded_image.txt
```

## Options

- `-encode`: Encode the input image to bytes
- `-decode`: Decode the encoded data and display the image
- `-input`: Input file path
- `-help`: Show help message

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE.md) file for details.

