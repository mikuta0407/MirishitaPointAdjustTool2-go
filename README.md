mltd-pat: **M**illion **L**ive **T**heater **D**ays **P**oint **A**djust **T**ool

https://github.com/mhrb-minase/MirishitaPointAdjustTool2 に触発されて、ミリシタイベのポイント調整ツールをGolangで作ろうとしています。とりあえず動きます。(元プロジェクト内のイベントアイテム数に応じた計算は未実装です)

## 使い方(予定)
- `go install github.com/mikuta0407/mltd-pat@latest`
- 周年イベなら `mltd-pat aniv <現在ポイント> <目標ポイント>`
- pointlist.tomlにライブ定義を書けます。
- ライブ総数を最小化したときの各ライブの必要回数が出力されます

## どこまでで出来ている?
- tomlからの定義パース
- `go run mltd-pat.go 280000 300000`等をすると、最小回数が表示される
  - 例:
    ```console
    $ mltd-pat aniv 280000 300000
    aniv called
    live count is below.
    normal 4M 10.5x 2
    normal MM 10.5x 2
    event 4x 8
    ```

## ToDo
- イベントアイテム数に応じた対応に追従する
- 10倍消費を考えてみる
- 通常のシアター形式への対応
- 調整ではなく、目標を達成できる(指定ボーダーをクリアできる)最小回数目処を表示する機能
  - 周年とシアター形式で対応
- 仕事を何回すればよいのか、イベントアイテム数はいくつ必要なのか、を計算する機能
- 必要元気数を出せたら良いな

## その他
- このソースコードを実行したことで起こるいかなる損害も補償はいたしかねます。自己責任でご利用ください。
- 改変・再配布は自由です。
