----
marp: true
theme: vinay
paginate: true
class: normal-code
footer: slides @ github.com/vinay-koho/lunch-n-learn
header: presentations with marp
----
![bg left:20%](https://iso.500px.com/wp-content/uploads/2016/09/stock-photo-174103407.jpg)

# Presentations with Marp
## Using markdown

----

## Metadata
- First line has to be `----` followed by `marp: true` for marp to parse the markdown properly
- Use themes with `theme: ` directive.
  - `gaia`, `uncover` and `default` are the default themes
  - Create custom themes with ` /* @theme vinay */` comment at the top of the css file
- you can use directives `header:` and `footer:` to have headers and footers
- you can use `class:` directive to attach global class to all the slides
- you can also use  `paginate: true`tag for slide numbers

----
<!-- _class: normal-code text-small -->

![bg vertical](https://fakeimg.pl/800x10/272020/fff/?text=%20)
![bg](https://fakeimg.pl/800x10/202720/fff/?text=%20)
![bg](https://fakeimg.pl/800x10/202027/fff/?text=%20)


## Slides and images
- Each slides starts with `----` in the markdown.
- Attach specific classes for each slide by adding a comment at the start with `_class:` directive like this,
  ```<!-- _class: text-small -->```
- Use image tags with multiple filters
```
![logo w:150](../assets/logo-inkscape.svg)
![logo w:150 invert:100% saturate:100%](../assets/logo-inkscape.svg)
![logo w:150 hue-rotate:90deg sepia:50%](../assets/logo-inkscape.svg)
![logo w:150 grayscale:0.5 opacity:.5](../assets/logo-inkscape.svg)
```
![logo w:150](../assets/logo-inkscape.svg)![logo w:150 invert:100% saturate:100%](../assets/logo-inkscape.svg)![logo w:150 hue-rotate:90deg sepia:50%](../assets/logo-inkscape.svg)![logo w:150 grayscale:0.5 opacity:.5](../assets/logo-inkscape.svg)

----
<!-- _class: leftbg normal-code -->
![bg 120%](https://iso.500px.com/wp-content/uploads/2016/09/stock-photo-174103407.jpg)

## Backgrounds

- You can have multiple background images with md img tags (as in earlier slide)
  ```md
  ![bg vertical](https://fakeimg.pl/800x10/272020/fff/?text=%20)
  ![bg](https://fakeimg.pl/800x10/202720/fff/?text=%20)
  ![bg](https://fakeimg.pl/800x10/202027/fff/?text=%20)
  ```
- You can have any image as background and scale to the size and few other props
```
![bg 120%](https://iso.500px.com/wp-content/uploads/2016/09/stock-photo-174103407.jpg)
```
- you can also have multiple side backgrounds using `left` or `right`directive
```
![bg left:20%](https://iso.500px.com/wp-content/uploads/2016/09/stock-photo-174103407.jpg)
```

----
## CSS tags and classes

- Each slide is wrapped with `section` tag.
- All text with-in  ``` ` ``` (backquotes) is wrapped with `code` tag.
- All text with-in ` ``` ` is wrapped with `pre > code` tag
- Code highlighting is done with `highlight.js` package. So you can use below classes to target sections of code
  - `.hljs-comment`, `.hljs-keyword`, `.hljs-title`, `.hljs-number`, `.hljs-string`, `.hljs-params`, `.hljs-built_in`,
- sample go code block
```go
  func add(x, y uint32) (z uint32) {
    z := x + y
    return z
  }

```

----

## VS Code usage

- Install Marp and Markdown extensions
- Add you custom css path to `markdown.marp.themes` settings file,
  - path is relative to the workspace folder.
```json
  markdown.marp.themes: [
      "./assets/default.css"
  ]
```
- Type `Cmd + Shift + p` and choose `Marp: Export slide deck...` option, to export the slides to desired format
----

# More information

[docs](https://marpit.marp.app/) - [site](https://marp.app) - [source](https://github.com/marp-team/marp)