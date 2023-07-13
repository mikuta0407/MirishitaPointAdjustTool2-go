mltd-pat: **M**illion **L**ive **T**heater **D**ays **P**oint **A**djust **T**ool

https://github.com/mhrb-minase/MirishitaPointAdjustTool2 に触発されて、ミリシタイベのポイント調整ツールをGolangで作ろうとしています。まだ動きません。

## 使い方(予定)
- `go install github.com/mikuta0407/mltd-pat@latest`
- 周年イベなら `mltd-pat aniv <現在ポイント> <目標ポイント>
- pointlist.tomlにライブ定義を書けます。
- ライブ総数を最小化したときの各ライブの必要回数が出力されます(予定)

## どこまでで出来ている?
- tomlからの定義パース
- `go run mltd-pat.go 280000 300000`等をすると計算関数が呼ばれる(が、panicする)

## ToDo
- そもそもの計算ロジックの移植(まだ動いていない)

## 実装したいこと
- 厳密な調整ではなく、目標ptラインを超えるには何回仕事と7倍ライブとイベントライブが必要か?を計算する(目処が知りたいやつ)
  - 周年とシアター形式で対応したい。ツアーは運ゲーなので対象外
- 必要元気も出せたら良い

## その他
- このソースコードを実行したことで起こるいかなる損害も補償はいたしかねます。自己責任でご利用ください。
- 改変・再配布は自由です。
