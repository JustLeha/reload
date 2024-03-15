
## Text Modifier Tool

This tool is designed to modify text files based on specific rules provided as arguments. It performs various modifications to the text, such as converting hexadecimal and binary numbers to decimal, changing letter case, handling punctuation, and more.



### Modifications

The tool executes the following modifications:

- **Hexadecimal Conversion:** Replaces instances of `(hex)` with the decimal version of the preceding hexadecimal number.
- **Binary Conversion:** Replaces instances of `(bin)` with the decimal version of the preceding binary number.
- **Uppercase Conversion:** Converts the word preceding `(up)` to uppercase.
- **Lowercase Conversion:** Converts the word preceding `(low)` to lowercase.
- **Capitalization:** Converts the word preceding `(cap)` to capitalized form.
- **Custom Case Conversion:** Converts the specified number of words preceding `(up)`, `(low)`, or `(cap)` to uppercase, lowercase, or capitalized form, respectively.
- **Punctuation Formatting:** Ensures proper spacing around common punctuation marks.
- **Ellipsis and Interrogation/Exclamation Combination:** Formats ellipsis and combined punctuation marks appropriately.
- **Single Quotes Placement:** Places single quotes correctly around words or phrases.

### Example

For example, given the input text:

```
"1E (hex) files were added. It has been 10 (bin) years. Ready, set, go (up) ! I should stop SHOUTING (low). Welcome to the Brooklyn bridge (cap). This is so exciting (up, 2). I was sitting over there ,and then BAMM !! As Elton John said: ' I am the most well-known homosexual in the world '. There it was. A amazing rock!"
```

The output will be:

```
"30 files were added. It has been 2 years. Ready, set, GO ! I should stop shouting. Welcome to the Brooklyn Bridge. THIS IS so EXCITING. I was sitting over there, and then BAMM!! As Elton John said: 'I am the most well-known homosexual in the world'. There it was. An amazing rock!"
```

### Note

- Ensure that your input file is properly formatted and contains the text you want to modify.
- The output file will contain the modified text based on the specified rules.

Enjoy using the Text Modifier Tool for your text modification needs!
