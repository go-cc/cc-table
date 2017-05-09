# 多音字字典处理过程记录

## 多音字表

### Prepare 多音字表 list text file

```sh
$ perl -ne 'print if /^\d/' duoyinzi-dict.md | wc -l 
108

perl -ne 'if (/^\d/) { s/^\d+(?|?)//; s/(?| *:).*$//; print }' duoyinzi-dict.md | tee ../../info/duoyinzi/duoyinzi-dict-F.txt | ccjf fj -i | tee ../../info/duoyinzi/duoyinzi-dict-J.txt
```

## 多音字词语表

### Prepare 多音字词语表 list text file

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

## 多音字字典

Goal: 

- Turn [pinyin.txt from mozillazg/phrase-pinyin-data](https://github.com/mozillazg/phrase-pinyin-data/) from text file into Go lookup table/code. 
- Show-off the Unix-Motto:
  * every tool just does one thing, but does it best
  * putting all such versatile tools together, it can be as powerful as you can imagine

### duoyinzi-dict-gen.go

```sh
duoyinzi-dict-gen.pl > duoyinzi-dict-gen_test.txt

cat duoyinzi-dict-gen_test.txt | go run duoyinzi-dict-gen.go  | tee /tmp/dict.json
{"A":{"B":{"C":{"D":{"G":"A B C D G ","K":"A B C D K "},"F":"A B C F ","G":{"H":{"I":{"J":{"K":"A B C G H I J K "}}},"L":{"L":"A B C G L L "}}}},"K":"A K ","L":{"M":{"N":{"O":{"P":{"Q":"A L M N O P Q "}}}}},"M":{"O":{"Q":"A M O Q "}}},"B":{"D":{"D":"B D D "}},"E":{"F":{"G":{"H":{"I":{"J":"E F G H I J "}}}}}}

cat /tmp/dict.json | jq .
```

```json
{
  "A": {
    "B": {
      "C": {
        "D": {
          "G": "A B C D G ",
          "K": "A B C D K "
        },
        "F": "A B C F ",
        "G": {
          "H": {
            "I": {
              "J": {
                "K": "A B C G H I J K "
              }
            }
          },
          "L": {
            "L": "A B C G L L "
          }
        }
      }
    },
    "K": "A K ",
    "L": {
      "M": {
        "N": {
          "O": {
            "P": {
              "Q": "A L M N O P Q "
            }
          }
        }
      }
    },
    "M": {
      "O": {
        "Q": "A M O Q "
      }
    }
  },
  "B": {
    "D": {
      "D": "B D D "
    }
  },
  "E": {
    "F": {
      "G": {
        "H": {
          "I": {
            "J": "E F G H I J "
          }
        }
      }
    }
  }
}
```


