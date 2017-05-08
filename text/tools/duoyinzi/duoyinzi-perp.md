# 多音字表

## Prepare 多音字表 list text file

```sh
$ perl -ne 'print if /^\d/' duoyinzi-dict.md | wc -l 
108

perl -ne 'if (/^\d/) { s/^\d+(?|?)//; s/(?| *:).*$//; print }' duoyinzi-dict.md | tee ../../info/duoyinzi/duoyinzi-dict-F.txt | ccjf fj -i | tee ../../info/duoyinzi/duoyinzi-dict-J.txt
```

## Prepare 多音字词语表 list text file

Generate:

    grep -E "`cat ../../info/duoyinzi/duoyinzi-dict-J.txt | tr '\n' '|' | sed 's/|*$//'`" ../../../external/pinyin-phrase-mozillazg/pinyin.txt | ccjf fj -i | sort -u | tee ../../info/duoyinzi/duoyinzi-phrase-J.txt  | ccjf jf -i | tee ../../info/duoyinzi/duoyinzi-phrase-F.txt

Check:

```sh
$ wc -l ../../info/duoyinzi/duoyinzi-phrase-?.txt
  21347 ../../info/duoyinzi/duoyinzi-phrase-F.txt
  21346 ../../info/duoyinzi/duoyinzi-phrase-J.txt
  42693 total

$ cat ../../info/duoyinzi/duoyinzi-phrase-?.txt | sort -u | tee ../../info/duoyinzi/duoyinzi-phrase.txt | wc -l
36146

cat ../../info/duoyinzi/duoyinzi-phrase-J.txt | cc2py -t 3 -p -i

grep -E "`cat ../../info/duoyinzi/duoyinzi-dict-J.txt | tr '\n' '|' | sed 's/|*$//'`" ../../info/duoyinzi/duoyinzi-phrase-J.txt
```



