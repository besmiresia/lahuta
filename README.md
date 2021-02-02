# Goals
## Ideas are power
> "Reading is the pathway from slavery to freedom." -Frederick Douglass

The internet can be a one-stop shop of mind-draining reality shows (if you let it,)
but it can also be a collection of the best thoughts and ideas from all of humanity.

Ultimately its what we make of it, so let's make our little corner of it awesome!
## Beautiful things don't just happen, you make them happen
This was inspired by Project Gutenburg. Unfortunately, the amount of volunteer
hours it would take for us to start contributing to them and translating books from 
scratch is too much at the moment, so we stand on the shoulders of giants.

A lot of great books have already been digitized, they are already in the public domain
and translation software exists, so there you go. 95% of the work has already been done
for us. The last 5% is up to us.

This website is totally customizable. You can edit any and every bit of text you see.

Something looks wrong? Fix it.

Don't know how? Ask for help.

Like I said, a lot of the work has been done so instead of volunteers spending weeks or months
on a single book, a few corrections from every reader and our library will grow and grow into
something we are proud of.

Learn while simultaneously helping others learn.
## The most important thing
Is always to have fun. I made this cuz:

1. I wanted to do an albanian book club with my friends but was disappointed that albanian resources are hard to find on the internet
2. Making things is fun
3. Making things and sharing them with others is the most fun.

I hope you don't just see text on screen but the opportunity to connect with others over what I think
are some truly inspirational ideas.

Included with every book:
1. An html file for reading on the web
2. An epub for reading on a kindle, tablet or smartphone
3. A text file (you got me technically its [markdown](https://www.markdownguide.org/basic-syntax/) but whatever) that can be easily edited
4. The english version that can be placed side-by-side as reference for translators or for those who want to learn/improve their albanian
5. A pdf file for uhh.. well if you want it its there!
6. All my love and support! (Which sure isn't worth much, but no refunds people!)

## How do I get involved??!!
#### Great I haven't scared you away 

1. Read a book (Enjoy yourself, have fun, learn something new)
2. Make a small correction (Reading and notice a mistranslation or akward phrasing? Help fix it so others don't have to see it too)
3. Share this project or a book with a friend
4. Start a book club/read with others
5. Find other ways to help promote albanian language or culture
    - We are stuck at home anyways so some potential ideas:
        - Sign up to help make albanian [duolingo](https://duolingo-incubator-web.duolingo.com/application) (If the Dothraki can do it so can we)


Comments, any questions on how to get involved, praise, suggestions, new ideas, criticizms, putdowns and general abuse can be left above in the issues tab or can be emailed to:

`besmiresia [AT] gmail.com`

Nice try robots not today.

Anyways I will leave you with some words of wisdom.

> "Divided we are weak. United we are unstoppable" - Some medieval albanian guy. probably.



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


