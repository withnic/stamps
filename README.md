# stamps

CSVのデータをtemplateに埋め込むプログラムです。

## Install

```bash
go get github.com/withnic/stamps
```

## Usage

```
$stamps replacementmessage csvfile.csv
```

置換するためのキーは ":" です。

ex.

data.csv

```
A,B,C
```

result

```bash
$./stamps "hoge:1:2:3:1" data.csv

# hogeABCA
```

":"を他で使っているときはflagで変更できます

ex.

```bash
$./stamps -k=@ "hoge@1@2@3" data.csv 

# hogeABCA
```

第一引き数にファイル名を取ることも出来ます。
その場合flagでfをわたします。

ex.

template.txt

```
hoge:1:2:3
```

```bash
$./stamps -f template.txt data.csv

# hogeABCA
```

フラグのeを渡すとそのまま実行します

ex.

exected.txt

```
-f
-1
```

```bash
$./stamps -e "ls :1" exected.txt
```

result

```bash
.
..
.git
.gitignore
data.csv
exected.txt
LICENSE
main.go
README.md
template.txt

LICENSE
README.md
data.csv
exected.txt
main.go
template.txt
```

フラグのgを一緒に渡すと非同期実行できます。


ex

```bash
$./stamps -e -g "echo :1" async.text
```

result

```bash
Hello

awsdefaw

HelloHOHGOEHOGEHOGH

Hessodas

asdfasdfasfdasdfasdfasdfasdfasdcfas

Hellosdfaswdfasdfsdf

asdfasdfsfasdsdfd

asdfasdfasdfas
```