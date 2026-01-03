# Shell implementation in Go
A small, private, Codecrafters-assisted side project — a minimal interactive shell (REPL) implemented in Go that supports builtins (`exit`, `echo`, `type`) and running external commands found on `PATH`. This repository is not open source; it’s for learning and practice with Codecrafters.

---

## Table of contents
- **Why this project?** 💡
- **Features** ✨
- **Quick start** 🔧
- **Usage examples** 🧪
- **How it works** 🔍
- **Development notes** 🛠️
- **Status & License** ⚠️
- **Author / Contact** 📫

---


## Features
- [x] **Builtin commands:**
    - [x] `exit` — exit the shell
    - [x] `echo` — print the arguments
    - [x] `type` — show whether a command is a builtin or where an executable lives
- [x] **Execute external programs** available in the `PATH`
---

## Quick start 

Clone/use your private workspace and then:

- Run:
```bash
./your_program.sh 
```

---

## Usage examples 🧪

Example REPL session:
```
$ echo Hello, Codecrafters!
Hello, Codecrafters!
$ type echo
echo is a shell builtin
$ type ls
ls is /bin/ls
$ ls -la
# (output from /bin/ls)
$ exit
```

---

