tseitin
===========

A tseitin encoder, and solve satisfactory problem.

Learning tseitin and yacc.



Installation
--------------

```sh
go get -d github.com/pocke/tseitin
cd $GOPATH/src/github.com/pocke/tseitin/
cd parser/
make
cd ../
go install
```

Usage
-------

```sh
$ tseitin '(a&b&!c)|(!a&b&c)'
# => (!D|!a)&(D|a)&(!C|b)&(!C|D)&(!b|!D|C)&(!B|c)&(!B|C)&(!c|!C|B)&(!F|!c)&(F|c)&(!G|b)&(!G|a)&(!b|!a|G)&(!E|F)&(!E|G)&(!F|!G|E)&(!B|A)&(!E|A)&(!A|B|E)
#    true map[b:true a:true c:false]
```

Refs
-----

- [SATソルバを使うためにCNFを作る - soutaroブログ](http://soutaro.hatenablog.com/entry/20100125/1264420334)
- [tseitin.pdf](http://www.cs.ox.ac.uk/people/thomas.wahl/Teaching/Software-Verification/Class-02-04/tseitin.pdf)
