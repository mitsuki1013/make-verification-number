# make-verification-number
法別番号、都道府県番号及び保険者別番号から「検証番号」を求める。

[ロジック](https://www.mhlw.go.jp/topics/2008/03/dl/tp0305-1az_0004.pdf#page=2)
## in
```
▶ go run main.go
1528400
1528400
1528401
1528401
1528402
1528402
1528403
1528403
1528405
1528405
1528406
1528406
1528407
1528407
1528408
1528408
1528409
1528409
1528410
1528410
1528411
1528411
```
## out
```
15284003
15284003
15284011
15284011
15284029
15284029
15284037
15284037
15284052
15284052
15284060
15284060
15284078
15284078
15284086
15284086
15284094
15284094
15284102
15284102
15284110
15284110
```
