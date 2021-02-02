# Bëni një korrigjim
## Metoda e lehtë (nëse doni të bëni korrigjime të shpejta nga faqja në internet)
1. Bëni llogari në github
2. Zgjidhni një [libër](https://github.com/besmiresia/lahuta/tree/master/books)
3. Klikoni te libri që dëshironi të redaktoni
4. Klikoni në {emri-i-librit}.txt
5. Klikoni lapsin për ta redaktuar
6. Bëni redaktimin tuaj

# Contributing

## Easy Method (if you want to make quick corrections from webpage)
1. Make github account
2. Choose a [book](https://github.com/besmiresia/lahuta/tree/master/books)
3. Click the folder of the book you want to edit
4. Click the file that is {name-of-book}.txt
5. Click the pencil in the top right corner to edit
6. Make your edit

## Contributing on your computer
This project uses crowbook, a tool to generate text into the PDF and EPUB formats.

If you want to download this tool:

### Linux
Everything is already for you so no work is needed

### Mac
Everything is already for you so no work is needed except when whenever you are 
supposed to type `bin/crowbook` type instead `bin/crowbook-mac`

## Make a change
Let's say you want to update a sentence.

1. Open the file that ends in `.txt` of the book you want with any text editor
 - (So if you want to edit "Gatsbi I Madh", `books/gatsbi_i_madh/gatsbi.txt`)
 - `vim books/gatsbi_i_madh/gatsbi.txt`)
2. Find the line you want to edit
3. Make your change
4. `bin/crowbook -s books/gatsbi_i_madh/gatsbi.txt`
6. Thank you so much for your contribution!


## Add a book
Let's say you want to add a whole new book

1. Create a new folder with the name of the book
 - (So if you want to add a new book called "Gatsbi I Madh", `gatsbi_i_madh`)
 - `mkdir books/gatsbi_i_madh`)
2. Create a txt file inside this new folder called {name-of-book.txt}
 - `touch books/gatsbi_i_madh/gatsbi.txt`)
3. Add the appropriate metadata
 - See the first 6 lines of an existing book.
 - All books must start with something like this but edit the Author And title fields to be correct for your book
4. Generate your book 
 - `bin/crowbook -s {name-of-book.txt}`
 - So in our example something like: `bin/crowbook -s books/gatsbi_i_madh/gatsbi.txt`
6. Thank you so much for your contribution!


