pdfopts=-t beamer -s --slide-level=2 --listings
htmlopts=-t slidy -s --slide-level=2 --listings -V theme=white
notesopts=-H notesheader.tex

presentations:=$(shell find . -name '*.md')

.PHONY=all serve pdf clean

all: $(presentations:.md=.all)

serve:
	npm run serve

pdf: $(presentations:.md=.pdf)

html: $(presentations:.md=.html)

$(presentations:.md=.all): %.all: %.pdf %.html %.notes.pdf

$(presentations:.md=.pdf): %.pdf: %.md
	pandoc $(pdfopts) $^ -o $@

$(presentations:.md=.html): %.html: %.md
	pandoc $(htmlopts) $^ -o $@

$(presentations:.md=.notes.pdf): %.notes.pdf: %.md
	pandoc $(pdfopts) $(notesopts) $^ -o $@

clean:
	rm -f $(presentations:.md=.pdf) $(presentations:.md=.notes.pdf) $(presentations:.md=.html)
