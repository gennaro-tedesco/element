<h1 align="center">
  <br>
  <img width="220" height="200" src="logo.png">
  <br>
  element
</h1>

<h2 align="center">
  <a href="#" onclick="return false;">
    <img alt="PR" src="https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat"/>
  </a>
  <a href="https://golang.org/">
    <img alt="Go" src="https://img.shields.io/badge/go-%2300ADD8.svg?&style=flat&logo=go&logoColor=white"/>
  </a>
  <a href="https://github.com/gennaro-tedesco/element/releases">
    <img alt="releases" src="https://img.shields.io/github/release/gennaro-tedesco/element"/>
  </a>
</h2>

<h4 align="center">The periodic table on the command line</h4>
<h3 align="center">
  <a href="#Installation">Installation</a> •
  <a href="#Usage">Usage</a>
</h3>

## Installation
Go get it!
```
go install github.com/gennaro-tedesco/element@latest
```

## Usage
![demo](https://user-images.githubusercontent.com/15387611/115306527-46b6fc00-a168-11eb-9b20-9fd437a7870b.gif)

Using `element` is element-ary! Run `element` and use the autocompletion menu to select the item to display properties of! Or, if you already know, run `element <name>` directly.

Check `element -h` for help and examples.


### Extras
<details>
  <summary>Other commands</summary>

  ```
  element random
  ```
  to show properties of a random element. Use it at shell start-up or as fortune cookie in `cowsay`!

  ```
  element table
  ```
  to display the periodic table in ascii format!

  <img alt="table" src="https://user-images.githubusercontent.com/15387611/115287360-7a862780-a150-11eb-9d6d-99a792d22610.png">
</details>

| command         | description                                | note
|:--------------- |:------------------------------------------ |:--------------------------------------
| element         | prompt autocompletion menu to select from  | `<C-d>` to quit, `<Up/Down>` to scroll
| element iron    | directly display properties of iron        |
| element random  | display properties of a random element     |
| element table   | display the periodic table in ascii format |
| element version | print current program version              |


## Feedback
If you find this application useful consider awarding it a ⭐, it is a great way to give feedback! Otherwise, any additional suggestions or merge request is warmly welcome!

