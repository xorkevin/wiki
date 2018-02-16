revealjsversion=3.6.0
revealjsurl=https://github.com/hakimel/reveal.js/archive/$(revealjsversion).tar.gz

pdfopts=-t beamer -s --slide-level=2 --listings
htmlopts=-t revealjs -s --slide-level=2 --listings -V theme=white
notesopts=-H notesheader.tex

presentations:=$(wildcard *.md)

.PHONY=all serve pdf revealjs clean

all: $(presentations:.md=.all)

serve:
	npm run serve

install-revealjs:
	wget https://github.com/hakimel/reveal.js/archive/3.6.0.tar.gz -O revealjs.tar.gz
	tar -xzvf revealjs.tar.gz
	rm revealjs.tar.gz
	rm -rf reveal.js
	mv reveal.js-$(revealjsversion) reveal.js

pdf: $(presentations:.md=.pdf)

revealjs: $(presentations:.md=.html)

$(presentations:.md=.all): %.all: %.pdf %.html %.notes.pdf

$(presentations:.md=.pdf): %.pdf: %.md
	pandoc $(pdfopts) $^ -o $@

$(presentations:.md=.html): %.html: %.md
	pandoc $(htmlopts) $^ -o $@

$(presentations:.md=.notes.pdf): %.notes.pdf: %.md
	pandoc $(pdfopts) $(notesopts) $^ -o $@

clean:
	rm -f $(presentations:.md=.pdf) $(presentations:.md=.html)
