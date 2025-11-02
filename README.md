# ByteRep

This is a simple command-line tool for encoding and decoding JPEG and PNG images to and from bytes, with an option to display images on the command-line using Unicode characters.

## Installation

Clone the repository and build the executable:

```bash
git clone https://github.com/ccianos/byterep.git
cd byterep
go build -o byterep
```

## Usage

### Encode an Image

Encode an image to bytes. The original format (JPEG or PNG) is preserved in the byte output.

```bash
# Encode a JPEG
./byterep -encode -input input_image.jpg > encoded_image.txt

# Encode a PNG
./byterep -encode -input input_image.png > encoded_image.txt
```

### Decode an Image (Smart Output)

The `-decode` command is now TTY-aware. It changes behavior based on whether you're viewing it in a terminal or redirecting the output to a file.

1. Display as Unicode (in Terminal)
When you run `-decode` and the output is your terminal, it will display the image using Unicode characters.

```bash
# Display an encoded file as Unicode
./byterep -decode -input encoded_image.txt

# You can also display an original image file directly as Unicode
./byterep -decode -input input_image.jpg

./byterep -decode -input input_image.png
```

2. Reconstruct Image File (Redirected Output)
When you redirect the output (using `>`), the command writes the raw image bytes, reconstructing the original file.

**Note:** Ensure your output file extension (.jpg or .png) matches the format of the original image you encoded.

```bash
# Reconstruct a JPEG image
./byterep -decode -input encoded_image.txt > decoded_image.jpg

# Reconstruct a PNG image
./byterep -decode -input encoded_image.txt > decoded_image.png
```

## Options

- `-encode`: Encode the input **JPEG** or **PNG** image to bytes
- `-decode`: Decode the encoded data. **Displays as Unicode to a terminal, or writes raw image bytes if redirected.**
- `-input`: Input file path
- `-help`: Show help message

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE.md) file for details.

