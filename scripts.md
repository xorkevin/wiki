change all files to 644 and all dirs to 755:
```bash
find . -type d -exec chmod 0755 {} \;
find . -type f -exec chmod 0644 {} \;
```

reset xrandr:
```bash
xrandr --output DisplayPort-2 --off && xrandr --auto
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

rsync:
```bash
rsync -azvP from_dir/ to_dir/
```
