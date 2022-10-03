# Boarder
##### A minimal 4chan board hoarder

## WIP
This is not fully baked, but still useable. If you experience any issues, 
see some improvement you think would be amazing, or just have some feedback 
for boarder, make an issue!

## Usage
Add `threads.txt` on the folder where you will be running boarder. The 
contents of the file would be the list of threads you want to subscribe to.
The format to add a thread is `board_thread`.
```
// Anime Wallpaper
w_2185924
w_2223911

// Wallpaper General
wg_7934675
wg_7920373
```

#### API used to:
 - Get threads of board : https://a.4cdn.org/{board}/threads.json
 - Get posts of thread : https://a.4cdn.org/{board}/thread/{thread}.json
 - Get media of post : https://i.4cdn.org/{board}/{tim.ext}
