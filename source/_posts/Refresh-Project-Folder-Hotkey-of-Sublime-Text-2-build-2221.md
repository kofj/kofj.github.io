title: "Refresh Project Folders Hotkey of Sublime Text 2,build 2221"
date: 2015-02-03 13:25:39
tags: 
- 编辑器 
- hotkey
categories: 术业
---
&emsp;&emsp;Sublime Text is a sophisticated text editor for code.I use it on Windows and OS X.But,the the problem I have with it of Version 2.0.2(Build 2221,OS X Yosemite),is that the project folder does not auto refresh when I add new files or folders that were created.
<!-- more -->
&emsp;&emsp;You can manually refresh the folders,but it's not efficient.We can set up a hotkey to do the work.

1. Open Sublime Text.
2. Select **Preferences** form the the top menu and then open **Key Bindings – Default** 
3. Go to the end of the file add add a comma to the last entry.
4. Add the following entry:
```JavaScript
// Refresh folder list with super+shift+r
{ "keys": ["super+shift+r"], "command": "refresh_folder_list" }
```
&emsp;&emsp;Now all you have to to is save the file and enjoy with this hotkey(![](http://km.support.apple.com/library/APPLE/APPLECARE_ALLGEOS/HT1343/ks_command.gif) + ![](http://km.support.apple.com/library/APPLE/APPLECARE_ALLGEOS/HT1343/ks_shift.gif) + R) to refresh your project!
&emsp;&emsp;Okay.In fact,this post is copy from [Michael's blog](http://michael1e.com/sublime-text-refresh-folder-hotkey/),but I don't like use **F5** on the Mac Book Pro.So, I changed the hotkey value to **Super + shift + R**. If you want to change the hotkey, simply switch [“super+shift+r″] for another value.