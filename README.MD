# Tony's Image converter

I needed a simple enough BMP to PNG converter, and while there are a bunch of those, even free, on the web, I really needed to preserve the original modification date of the files.

Since it was a rather unusual ask, and I could not find a decent enough solution, I decided to do it myself.

Nothing fancy here, the software asks for a source and target directories, then just repacks BMP into PNG with no compression whatsoever, while preserving the original modification date.

## [Version history](CHANGELOG.MD)

## *Notes*

**github.com/jdeng/goheif** package included here requires GCC installed. I took mine from here https://sourceforge.net/projects/mingw-w64/files/, and it unfortunatelly throws some errors.