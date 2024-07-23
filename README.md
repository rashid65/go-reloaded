Go-Reoladed
===========

Go-Reoladed is a versatile Go application designed for efficient and flexible text manipulation. This tool offers a range of functions to process and transform text, making it ideal for various data processing tasks.

Features
--------

-   Hexadecimal to Decimal Conversion
-   Binary to Decimal Conversion
-   Uppercase Conversion
-   Lowercase Conversion
-   Capitalized Conversion
-   Batch Case Conversion
-   Punctuation Adjustment
-   Special Punctuation Groups
-   Quotation Marks Adjustment
-   Article Adjustment

Text Manipulation Modifications
-------------------------------

The tool you are about to build will receive as arguments the name of a file containing a text that needs some modifications (the input) and the name of the file the modified text should be placed in (the output). Below is a list of possible modifications that your program should execute:

1.  **Hexadecimal to Decimal Conversion**

    -   Replace every instance of `(hex)` with the decimal version of the preceding hexadecimal word.
    -   **Example:** `"1E (hex) files were added"` -> `"30 files were added"`
2.  **Binary to Decimal Conversion**

    -   Replace every instance of `(bin)` with the decimal version of the preceding binary word.
    -   **Example:** `"It has been 10 (bin) years"` -> `"It has been 2 years"`
3.  **Uppercase Conversion**

    -   Replace every instance of `(up)` with the uppercase version of the preceding word.
    -   **Example:** `"Ready, set, go (up)!"` -> `"Ready, set, GO!"`
4.  **Lowercase Conversion**

    -   Replace every instance of `(low)` with the lowercase version of the preceding word.
    -   **Example:** `"I should stop SHOUTING (low)"` -> `"I should stop shouting"`
5.  **Capitalized Conversion**

    -   Replace every instance of `(cap)` with the capitalized version of the preceding word.
    -   **Example:** `"Welcome to the Brooklyn bridge (cap)"` -> `"Welcome to the Brooklyn Bridge"`
6.  **Batch Case Conversion**

    -   For `(low, <number>)`, `(up, <number>)`, and `(cap, <number>)`, convert the specified number of preceding words to lowercase, uppercase, or capitalized, respectively.
    -   **Example:** `"This is so exciting (up, 2)"` -> `"This is SO EXCITING"`
7.  **Punctuation Adjustment**

    -   Ensure punctuation marks (., ,, !, ?, :, ;) are close to the previous word and spaced apart from the next one.
    -   **Example:** `"I was sitting over there ,and then BAMM !!"` -> `"I was sitting over there, and then BAMM!!"`
8.  **Special Punctuation Groups**

    -   Handle groups of punctuation like `...` or `!?` by ensuring proper formatting.
    -   **Example:** `"I was thinking ... You were right"` -> `"I was thinking... You were right"`
9.  **Quotation Marks Adjustment**

    -   Ensure pairs of single quotation marks (`'`) are placed around the word(s) without spaces.
    -   **Example:** `"I am exactly how they describe me: ' awesome '"` -> `"I am exactly how they describe me: 'awesome'"`
    -   For multiple words, place the marks next to the corresponding words.
    -   **Example:** `"As Elton John said: ' I am the most well-known homosexual in the world '"` -> `"As Elton John said: 'I am the most well-known homosexual in the world'"`
10. **Article Adjustment**

    -   Replace "a" with "an" if the next word begins with a vowel (a, e, i, o, u) or an h.
    -   **Example:** `"There it was. A amazing rock!"` -> `"There it was. An amazing rock!"`

Technologies Used
-----------------

-   Go

Getting Started
---------------

1.  Clone the repository.
2.  Install Go on your system.
3.  Obtain the input text file and specify the output file.
4.  Run the application to start manipulating text.
