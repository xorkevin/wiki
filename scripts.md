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
