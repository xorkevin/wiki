change all files to 644 and all dirs to 755:
```bash
find . -type d -exec chmod 0755 {} \;
find . -type f -exec chmod 0644 {} \;
```
