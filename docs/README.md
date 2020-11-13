# PTerm 


<p align="center">

<a href="https://github.com/dops-cli/dops/releases">
<img src="https://img.shields.io/badge/platform-windows%20%7C%20macos%20%7C%20linux-informational?style=for-the-badge" alt="Downloads">
</a>

<br/>

<a style="text-decoration: none" href="https://github.com/pterm/pterm/releases">
<img src="https://img.shields.io/github/v/release/pterm/pterm?style=flat-square" alt="Latest Release">
</a>
&nbsp;
<a style="text-decoration: none" href="https://github.com/pterm/pterm/stargazers">
<img src="https://img.shields.io/github/stars/pterm/pterm.svg?style=flat-square" alt="Stars">
</a>
&nbsp;
<a style="text-decoration: none" href="https://github.com/pterm/pterm/fork">
<img src="https://img.shields.io/github/forks/pterm/pterm.svg?style=flat-square" alt="Forks">
</a>
&nbsp;
<a style="text-decoration: none" href="https://github.com/pterm/pterm/issues">
<img src="https://img.shields.io/github/issues/pterm/pterm.svg?style=flat-square" alt="Issues">
</a>
&nbsp;
<a style="text-decoration: none" href="https://opensource.org/licenses/MIT">
<img src="https://img.shields.io/badge/License-MIT-yellow.svg?style=flat-square" alt="License: MIT">
</a>

</p>

> ✨ A modern go module to beautify console output

![PTerm Demo](https://raw.githubusercontent.com/pterm/pterm/master/_examples/demo/animation.svg)

## 🥅 Goal of PTerm

PTerm is designed to create a **platform independent way to create beautiful terminal output**. Most modules that want to improve the terminal output do not guarantee platform independence - PTerm does. PTerm follows the **most common methods for displaying color in a terminal**. With PTerm, it is possible to create beautiful output **even in low-level environments**. 

### • 🪀 Easy to use

Our first priority is to keep PTerm as easy to use as possible. With many [examples](https://github.com/pterm/pterm/tree/master/_examples) for each individual component, getting started with PTerm is extremely easy. All components are similar in design and implement interfaces to simplify mixing individual components together.

### • 🤹‍♀️ Cross-Platform

We take special precautions to ensure that PTerm works on as many operating systems and terminals as possible. Whether it's `Windows CMD`, `macOS iTerm2` or in the backend (for example inside a `GitHub Action` or other CI systems), PTerm **guarantees** beautiful output!\
\
*PTerm is actively tested on `Windows`, `Linux (Debian & Ubuntu)` and `macOS`.*

### • 🧪 Well tested

> PTerm has a 100% test coverage, which means that every line of code inside PTerm gets tested automatically
We test PTerm continuously. However, since a human cannot test everything all the time, we have our own test system with which we currently run thousands of automated tests to ensure that PTerm has no bugs. 

### • ✨ Consistent Colors

PTerm uses the [ANSI color scheme](https://en.wikipedia.org/wiki/ANSI_escape_code#3/4_bit) which is widely used by terminals to ensure consistent colors in different terminal themes.
If that's not enough, PTerm can be used to access the full RGB color scheme (16 million colors) in terminals that support `TrueColor`.

![ANSI Colors](https://user-images.githubusercontent.com/31022056/96002009-f10c3a80-0e38-11eb-8d90-f3150150599c.png)

### • 📚 Component system

PTerm consists of many components, called `Printers`, which can be used individually or together to generate pretty console output.

### • 🛠 Configurable

PTerm can be used by without any configuration. However, you can easily configure each component with little code, so everyone has the freedom to design their own terminal output.
