https://gist.github.com/suntong/a532d5136713a7f1c67a3f118cc0d51c

Ubuntu12.04版本中默认的是微米黑. Ubuntu从13.04开始取消了默认的微米黑，回退为之前的正黑。这我觉得只能用回退形容了。

在默认设置下，Ubuntu 13以后中文使用的是文泉驿正黑。我总觉得它的效果有些发虚，模糊，不满意。

于是决定替换掉回来。

Ubuntu 对中文字体的控制集中在一个文件，/etc/fonts/conf.d/69-language-selector-zh-cn.conf，from the language-selector-common pacakge.

修改如下，简单说，就是把系统默认的正黑替换为微米黑并提前。有这一个在最前面保证就足够了。

这样，英文默认Ubuntu，中文默认文泉驿微米黑，即使浏览器中带着CSS样式的文字，也有不错的显示效果。效果理想。

History:

- Include the 69-language-selector-zh-cn.conf from Ubuntu 15.04 Vivid.
- Update it to Ubuntu 16.04 LTS Xenial
- Done Changing 中文默认 to 文泉驿微米黑
- Add Antialias when pointsize <13
- Add URxvt Xresources file

Update 2016-05-23, 

Given the package size info as following, decided to remove the wqy-zenhei (文泉驿正黑) entirely, and use fonts-noto instead.

    16461 fonts-wqy-zenhei
    15581 fonts-noto-hinted
    5167 fonts-wqy-microhei

The file 69-language-selector-zh-cn.conf can be found in this directory. 


Ref:

http://www.cnblogs.com/daizhe11/p/3384391.html

