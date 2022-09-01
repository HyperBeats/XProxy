<h1 align="center">XProxy</h1>

`Powerfull proxy scraper and checker.`

### Features:
- scrape proxies from url/public proxy list.
- Support `http - socks4 - socks5`.
- check for proxy anonymity level.
- ultra fast.

### Speed:
- 100K+ scrape under 5s.
- (800 goroutines) Checked 45973 proxies in 329.308784s.

### Requirements:
- golang 1.18+

### Download:
- download the lasted compiled version [here](https://github.com/Its-Vichy/XProxy/releases/tag/lasted).

### Todo:
    [ ] Crawl html page with regex

### Known issue:
- [FIXED] Checker crash if there is invalid format into proxies file, this can happen if you are using scraper.
- [FIXED] `fixed: "bufio.Scanner: token too long"`, this error happen when you are loading large proxies file (100K+)

---

<p align="center">
    <img alt="GitHub Repo stars" src="https://img.shields.io/github/stars/Its-Vichy/XProxy?style=for-the-badge&logo=stylelint&color=black">
    <img alt="GitHub top language" src="https://img.shields.io/github/languages/top/Its-Vichy/XProxy?style=for-the-badge&logo=stylelint&color=black">
    <img alt="GitHub last commit" src="https://img.shields.io/github/last-commit/Its-Vichy/XProxy?style=for-the-badge&logo=stylelint&color=black">
</p>