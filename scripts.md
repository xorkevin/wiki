gpg export only single subkey

```bash
gpg --list-keys --keyid-format long KEYID
gpg --armor --export-secret-subkeys SUBKEYID!
```

profile program
```bash
perf stat
```

change all files to 644 and all dirs to 755:

```bash
find . -type d -exec chmod 0755 {} \;
find . -type f -exec chmod 0644 {} \;
```

reset xrandr:

**Note**: currently broken

```bash
xrandr --output DisplayPort-2 --off && xrandr --auto
```

reset alsa:

```bash
alsactl restore
```

optimize mp4:

```bash
ffmpeg -i animated.gif -movflags faststart -pix_fmt yuv420p -vf "scale=trunc(iw/2)*2:trunc(ih/2)*2" video.mp4
```

encode in h265:

```bash
ffmpeg -i input.mp4 -c:v libx265 -crf 28 output.mp4
```
quality setting 18 is high

encode in h264:

```bash
ffmpeg -i input.mp4 -c:v libx264 -crf 23 output.mp4
```

valgrind memcheck:

```bash
valgrind --tool=memcheck --leak-check=yes --show-reachable=yes --num-callers=20 --track-fds=yes ./test --arg1 --arg2
```

cutting video:

```bash
ffmpeg -i video.mp4 -ss 00:11:00 -t 00:01:00 -vcodec copy -acodec copy cutvid.mp4
```

optimize png:

```bash
optipng -o7 -strip all input.png
```

optimize jpg:

```bash
jpegtran -optimize -progressive -copy none -outfile output.jpg input.jpg
```

rsync:

```bash
rsync -azvP from_dir/ to_dir/
```

pacman remove group:

```bash
pacman -Qgq package_group | pacman -Rs -
```

pacman verify files:

```bash
sudo pacman -Qk | grep -vE '0 missing files$'
```

pacman list explicitly installed:

```bash
pacman -Qeq
```

pacman remove orphans:

```bash
pacman -Rns $(pacman -Qtdq)
```

clean pacman cache:

```bash
paccache -r
```

fix corrupt zsh history:

```bash
#!/usr/bin/env sh

mv ~/.zhistory ~/.zsh_history_bad
strings ~/.zsh_history_bad > ~/.zhistory
fc -R ~/.zhistory
rm ~/.zsh_history_bad
```

generate a file with random contents:

```bash
dd if=/dev/urandom of=file.txt bs=1048576 count=100
```

refresh keys

```bash
# pacman-key --init
# pacman-key --populate archlinux
# pacman-key --refresh keys
```

export gpg subkey

```bash
gpg --armor --export-secret-subkey keyid! | gpg --armor --symmetric --output filename.sec.asc
```

git line stats:

```bash
git log --author="Kevin Wang" --pretty=tformat: --numstat | awk '{add+=$1; subs+=$2} END {printf "added: %s, removed: %s\n", add, subs}'
```

Grub boot without monitor

/etc/default/grub:

```
GRUB_CMDLINE_LINUX="nomodeset"
```

Grub generate config with os prober

/etc/default/grub:

```
GRUB_DISABLE_OS_PROBER=false
```

```bash
grub-mkconfig -o /boot/grub/grub.cfg
```

change backlight

```bash
tee /sys/class/backlight/intel_backlight/brightness <<< 384
```

ssh match group

```
Match Group wheel
  PasswordAuthentication no
```

Makefile latex

```Makefile
NAME=srcfile
SRC=$(NAME).tex
OUT=$(NAME).pdf

.PHONY: clean

$(OUT): $(SRC)
	latexmk -pdf -bibtex -pdflatex="pdflatex -interaction=nonstopmode" $^

clean:
	latexmk -bibtex -CA
```

```tex
\usepackage[style=trad-abbrv]{biblatex}
\addbibresource{bibfile.bib}
```

```tex
\documentclass{article}
\usepackage[left=4cm,right=4cm,top=2cm,bottom=2cm]{geometry}
\usepackage[T1]{fontenc}
\usepackage{mathtools}
\usepackage{enumitem}
\usepackage{amsthm}
\usepackage{amssymb}
\usepackage{algorithm}
\usepackage{algpseudocode}
\usepackage{listings}
\usepackage{float}
\usepackage{array}
\usepackage{booktabs}
\usepackage{graphicx}
\usepackage{tikz}
\usepackage{tikz-uml}
\usepackage{xcolor}
\usepackage[hypertexnames=false]{hyperref}

\newtheorem{theorem}{Theorem}[section]
\newtheorem{corollary}{Corollary}[theorem]
\newtheorem{lemma}[theorem]{Lemma}
\theoremstyle{definition}
\newtheorem{definition}[theorem]{Definition}
\theoremstyle{remark}
\newtheorem*{remark}{Remark}
\renewcommand\qedsymbol{$\blacksquare$}

\lstdefinestyle{codestyle}{
  backgroundcolor=\color[HTML]{FAFAFA},
  keywordstyle=\color[HTML]{8959a8},
  commentstyle=\color[HTML]{8e908c},
  stringstyle=\color[HTML]{718c00},
  numberstyle=\tiny\color[HTML]{f5871f},
  basicstyle=\footnotesize,
  breakatwhitespace=true,
  breaklines=true,
  captionpos=b,
  keepspaces=true,
  numbers=left,
  numbersep=4pt,
  showspaces=false,
  showstringspaces=false,
  showtabs=false,
  tabsize=2
}

\lstset{style=codestyle}
```

```Makefile
NAME=srcfile
SRC=$(NAME).tex
OUT=$(NAME).pdf
NOTESNAME=$(NAME).notes
NOTESSRC=$(NOTESNAME).tex
NOTES=$(NOTESNAME).pdf

IMGS=$(shell find . -name '*.dot')

.PHONY: all build notes

all: build notes

build: $(OUT)

$(OUT): $(SRC) images
	latexmk -pdfxe -pdfxelatex="xelatex -interaction=nonstopmode" $<

notes: $(NOTES)

$(NOTES): $(NOTESSRC) images
	latexmk -pdfxe -pdfxelatex="xelatex -interaction=nonstopmode" $<

.PHONY: clean init images clean-images

clean:
	latexmk -CA

images: $(IMGS:.dot=_gen.png)

$(IMGS:.dot=_gen.png): %_gen.png: %.dot
	dot -Tpng $^ > $@

clean-images:
	rm *_gen.png
```

```tex
\documentclass{beamer}
\usetheme{metropolis}
\usepackage{pgfpages}

\ifdefined\NOTES
  \setbeameroption{show only notes}
\fi

\title[Short title]{Longer title}
\author{Kevin Wang}
\date[Short org year]{Long org, year}
\logo{\includegraphics[height=0.5cm]{assets/logo.png}}

\begin{document}

\maketitle

\section{Section Title}

\subsection{Subsection}

\begin{frame}{Frame title}{subtitle}
  Frame text
  \note{notes text}
\end{frame}

\end{document}
```

```tex
\documentclass{beamer}
\usetheme{metropolis}
\usepackage{pgfpages}
\setbeameroption{show notes on second screen}

\title{Title}
\subtitle{Subtitle}
\author{Kevin Wang}
\year=2020\relax
\month=01\relax
\day=02\relax
\date{\today}
\institute{Org}
\logo{\includegraphics[height=0.5cm]{assets/logo.png}}

% fix text color issue for second screen notes
\makeatletter
\def\beamer@framenotesbegin{
  % at beginning of slide
  \usebeamercolor[fg]{normal text}
  \gdef\beamer@noteitems{}
  \gdef\beamer@notes{}
}
\makeatother

\begin{document}

\maketitle

\section{Section Title}

\subsection{Subsection}

\begin{frame}{Frame title}{subtitle}
  Frame text
  \note{notes text}
\end{frame}
\note[itemize]{
  \item notes text
}

\end{document}
```
