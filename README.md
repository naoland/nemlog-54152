# 簡単プログラミング！LINEに自分で好きなメッセージを送ってみよう（exeファイルあり）

## はじめに

以前、[簡単プログラミング！XEMの現在価格をLINEに通知しよう（Go編）](https://nemlog.nem.social/blog/53471)という記事を書きましたが、今回はご自分の好きなメッセージを、ご自分のWindowsやMacのターミナルから送信できるようにしました。

Gitbotの使い方がよくわからないという方や、皆さんにも手軽に使ってみていただけるように、クロスコンパイルしてWindowsとMacのバイナリを作成しました。
ダウンロードしてお使いいただけます。

メッセージを送信するには、LINEのアクセストークンというものが必要になります。
[LINE Notify](https://notify-bot.line.me/ja/)の登録と設定が終わってない方は、[こちらの記事](https://nemlog.nem.social/blog/53471)を最初にご覧ください。

今回はソースコードの説明はしません。興味のある方は、[notif.go](https://github.com/naoland/nemlog-54152/blob/main/notif.go)をご覧ください。

## 使い方

- バイナリを[ここから](https://github.com/naoland/nemlog-54152/releases/tag/v1.0)ダウンロードして、任意の場所に保存します。
- ターミナルアプリを起動して、CDコマンドなどで、ダウンロードした場所に移動します。
- ターミナルアプリで次のコマンドを入力します。

`-token`の後ろがLINEのトークン（文字列）で、その後ろが送信したいメッセージです。

```
notif.exe -token KDJiqR84TRksLEwctIHgxrNrmEvgpV  テストメッセージです

注）`KDJiqR84TRksLEwctIHgxrNrmEvgpV`は適当な文字列です、ご自身のトークンに置き換えて入力してください。
```
メッセージの送信が成功すると次の結果がターミナルに表示されます。
スマホにもすぐにメッセージが届きます。

```
 {"status":200,"message":"ok"}
```

## まとめ

お好きなメッセージを送信できましたか？

今回はGo言語のクロスコンパイル機能を使って、Windows10やmacOS用のバイナリを簡単に作ることができました。素晴らしい（クロスコンパイルはUbuntu上で行っています）。

BATファイルを作ってより簡単に送信したり、時間を決めて送信したり、メッセージを動的に変えてみたりして、遊んでみてください。

動画で使い方などご紹介したかったのですが、機密情報にぼかしをかける方法を知らないので、ご勘弁くださいｗ。 そのうち勉強したいと思います。

ところで、Windows 10をご利用の方は、コマンドプロンプトではなく、Windows Terminalを使った方がよいです。下のリンク欄を参考にしてください。


## 関連情報へのリンク

- [Windows Terminal](https://www.microsoft.com/ja-jp/p/windows-terminal/9n0dx20hk701?rtc=1&activetab=pivot:overviewtab)
- [Windows ターミナルをインストールしてセットアップする](https://docs.microsoft.com/ja-jp/windows/terminal/get-started)
- [簡単プログラミング！LINEに自分で好きなメッセージを送ってみよう（exeファイルあり）](https://nemlog.nem.social/blog/54152) nemlog
- [簡単プログラミング！LINEに自分で好きなメッセージを送ってみよう（exeファイルあり）](https://github.com/naoland/nemlog-54152) github
- [naoland/nemlog-posts: nemlog投稿記事 メイン](https://github.com/naoland/nemlog-posts)
