# エラー
エラーの処理を共通化するパッケージ
## 使い方
新規に定義するエラーは`define_xxx.go`に定義する
```
errorName = newXXX(code, message)
```
## 呼び出す方法
### ラップしたい場合
```
if err := Update(); err != nil {
    return errors.Wrap(err, "UpdateUserData")
}
// 独自で定義したエラーを使う
if err != nil {
    return errors.UpdateUserData.Wrap(err, "UpdateUserData")
}
```
### エラーを発生させたい場合
```
if device != `` {
    // pattern1
    return errors.New("unknown device")
    // pattern2
    return errors.Errorf("AuthService device != ``")
    // pattern3
    return errors.InvalidDevice.New("AuthService device != ``")
}
```
