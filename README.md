Snippets for transpilers.

---

Every group of snippets starts with `primary` file
marked with `transpilery:primary` comment. This is
the original code that is transpiled to other
languages.

---

Most sources are directly executable, except for
languages where you have to manually setup file
associations. Assuming "!.!" means filename with
extension:

  * .go - `go run !.!`
  * .hx - `haxe --run !.!`

`haxe --run` still needs a fix to ignore filename
https://github.com/HaxeFoundation/haxe/issues/6182
