
# Unicode Character Ranges for CJK

http://jrgraphix.net/research/unicode.php

```
$ lynx -dump -width=512 "http://jrgraphix.net/research/unicode.php" | grep CJK | cut -c60-
       2E80 — 2EFF     [31]CJK Radicals Supplement
       3000 — 303F     [37]CJK Symbols and Punctuation
       3200 — 32FF     [53]Enclosed CJK Letters and Months
       3300 — 33FF     [55]CJK Compatibility
       3400 — 4DBF     [57]CJK Unified Ideographs Extension A
       4E00 — 9FFF     [61]CJK Unified Ideographs
       F900 — FAFF     [77]CJK Compatibility Ideographs
       FE30 — FE4F     [87]CJK Compatibility Forms
       20000 — 2A6DF   [125]CJK Unified Ideographs Extension B
       2F800 — 2FA1F   [127]CJK Compatibility Ideographs Supplement
```

# CJK Unified Ideographs Wiki

https://en.wikipedia.org/wiki/CJK_Unified_Ideographs_(Unicode_block)

https://en.wikipedia.org/wiki/CJK_Unified_Ideographs


# The complete range for Chinese characters in Unicode

https://stackoverflow.com/a/1366113/2125837

May be you would find a complete list through the [CJK Unicode FAQ][1] (which does include "Chinese, Japanese, and Korean" characters)

The "[East Asian Script][2]" document does mention:

> Blocks Containing Han Ideographs

> Han ideographic characters are found in five main blocks of the Unicode Standard, as
shown in Table 12-2

Table 12-2. Blocks Containing Han Ideographs

    Block                                   Range       Comment
    CJK Unified Ideographs                  4E00-9FFF   Common
    CJK Unified Ideographs Extension A      3400-4DBF   Rare
    CJK Unified Ideographs Extension B      20000-2A6DF Rare, historic
    CJK Unified Ideographs Extension C      2A700–2B73F Rare, historic
    CJK Unified Ideographs Extension D      2B740–2B81F Uncommon, some in current use
    CJK Unified Ideographs Extension E      2B820–2CEAF Rare, historic
    CJK Compatibility Ideographs            F900-FAFF   Duplicates, unifiable variants, corporate characters
    CJK Compatibility Ideographs Supplement 2F800-2FA1F Unifiable variants

Note: the block ranges can evolve over time: latest is in [CJK Unified Ideographs][3].

See also Wikipedia:

- [CJK Unified Ideographs Extension A][4]
- [CJK Unified Ideographs Extension B][5]
- [CJK Unified Ideographs Extension C][6]
- [CJK Unified Ideographs Extension D][7]
- [CJK Unified Ideographs Extension E][8]


  [1]: http://www.unicode.org/faq/han_cjk.html
  [2]: http://www.unicode.org/versions/Unicode5.0.0/ch12.pdf#G12159
  [3]: https://en.wikipedia.org/wiki/CJK_Unified_Ideographs
  [4]: https://en.wikipedia.org/wiki/CJK_Unified_Ideographs_Extension_A
  [5]: https://en.wikipedia.org/wiki/CJK_Unified_Ideographs_Extension_B
  [6]: https://en.wikipedia.org/wiki/CJK_Unified_Ideographs_Extension_C
  [7]: https://en.wikipedia.org/wiki/CJK_Unified_Ideographs_Extension_D
  [8]: https://en.wikipedia.org/wiki/CJK_Unified_Ideographs_Extension_E
 
Answered Sep 2 '09 at 6:27

VonC

## The complete range for CJK characters in Unicode

https://stackoverflow.com/a/11415841/2125837

Unicode currently has 74605 CJK characters. CJK characters not only includes characters used by Chinese, but also Japanese Kanji, Korean Hanja, and Vietnamese [Chu Nom][1]. Some CJK characters are **not** Chinese characters.

### 1) 20941 characters from the [CJK Unified Ideographs block][2].

Code points U+4E00 to U+9FCC.

1. [U+4E00 - U+62FF][3] 
1. [U+6300 - U+77FF][4]
1. [U+7800 - U+8CFF][5]
1. [U+8D00 - U+9FCC][6]


### 2) 6582 characters from the [CJKUI Ext A block][7].

Code points [U+3400 to U+4DB5][8]. Unicode 3.0 (1999). 

### 3) 42711 characters from the [CJKUI Ext B block][9].

Code points U+20000 to U+2A6D6. Unicode 3.1 (2001). 

1. [U+20000 - U+215FF][10]
1. [U+21600 - U+230FF][11]
1. [U+23100 - U+245FF][12]
1. [U+24600 - U+260FF][13]
1. [U+26100 - U+275FF][14]
1. [U+27600 - U+290FF][15]
1. [U+29100 - U+2A6DF][16]

### 3) 4149 characters from the [CJKUI Ext C block][17].

Code points [U+2A700 to U+2B734][18]. Unicode 5.2 (2009).

### 4) 222 characters from the [CJKUI Ext D block][19].

Code points [U+2B740 to U+2B81D][20]. Unicode 6.0 (2010).

### 5) CJKUI Ext E block.

[Coming soon][21]

If the above is not spaghetti enough, take a look at [known issues][22]. Have fun =)


  [1]: http://en.wikipedia.org/wiki/Ch%E1%BB%AF_N%C3%B4m
  [2]: http://www.fileformat.info/info/unicode/block/cjk_unified_ideographs/index.htm
  [3]: http://en.wikipedia.org/wiki/List_of_CJK_Unified_Ideographs,_part_1_of_4
  [4]: http://en.wikipedia.org/wiki/List_of_CJK_Unified_Ideographs,_part_2_of_4
  [5]: http://en.wikipedia.org/wiki/List_of_CJK_Unified_Ideographs,_part_3_of_4
  [6]: http://en.wikipedia.org/wiki/List_of_CJK_Unified_Ideographs,_part_4_of_4
  [7]: http://www.fileformat.info/info/unicode/block/cjk_unified_ideographs_extension_a/index.htm
  [8]: http://en.wikipedia.org/wiki/List_of_CJK_Unified_Ideographs,_Extension_A
  [9]: http://www.fileformat.info/info/unicode/block/cjk_unified_ideographs_extension_b/index.htm
  [10]: http://en.wikipedia.org/wiki/List_of_CJK_Unified_Ideographs,_Extension_B_%2820000-215FF%29
  [11]: http://en.wikipedia.org/wiki/List_of_CJK_Unified_Ideographs,_Extension_B_%2821600-230FF%29
  [12]: http://en.wikipedia.org/wiki/List_of_CJK_Unified_Ideographs,_Extension_B_%2823100-245FF%29
  [13]: http://en.wikipedia.org/wiki/List_of_CJK_Unified_Ideographs,_Extension_B_%2824600-260FF%29
  [14]: http://en.wikipedia.org/wiki/List_of_CJK_Unified_Ideographs,_Extension_B_%2826100-275FF%29
  [15]: http://en.wikipedia.org/wiki/List_of_CJK_Unified_Ideographs,_Extension_B_%2827600-290FF%29
  [16]: http://en.wikipedia.org/wiki/List_of_CJK_Unified_Ideographs,_Extension_B_%2829100-2A6DF%29
  [17]: http://www.fileformat.info/info/unicode/block/cjk_unified_ideographs_extension_c/index.htm
  [18]: http://en.wikipedia.org/wiki/List_of_CJK_Unified_Ideographs,_Extension_C
  [19]: http://www.fileformat.info/info/unicode/block/cjk_unified_ideographs_extension_d/index.htm
  [20]: http://en.wikipedia.org/wiki/List_of_CJK_Unified_Ideographs,_Extension_D
  [21]: http://unicode.org/roadmaps/sip/
  [22]: http://en.wikipedia.org/wiki/CJK_Unified_Ideographs#Known_issues
