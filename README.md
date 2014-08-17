# done-list (in golang)

Golang implementation of [fukayatsu/done-list](https://github.com/fukayatsu/done-list).

## Example

```
$ done-list
## yurufuwa/yurufuwa/issues
- [x] [認証部分を別ライブラリに切り分けました by satococoa](https://github.com/yurufuwa/yurufuwa/pull/5)

## satococoa/hoge/issues
- [ ] [Hogeの検討スレ by satococoa](https://github.com/satococoa/hoge/issues/27)
- [ ] [Fuga足そう by satococoa](https://github.com/satococoa/hoge/issues/28)
- [ ] [デザイン修正点 by foo](https://github.com/satococoa/hoge/issues/21)
```

## How to use

### Install

You need golang. (e.g. `$ brew install go`)

```
$ go get github.com/satococoa/done-list/cmd/done-list
```

### Update

```
$ go get -u github.com/satococoa/done-list/cmd/done-list
```
