# Tools to extract data form webdevdata.org

Small set of command line tools written in Go to help extracting data
from webdevdata.org.

## Using tools

You can find cross-compiled executables in the ```release``` directory.

```bash
$ ls release
webdevdata-tools-darwin-386.tgz
webdevdata-tools-darwin-amd64.tgz
webdevdata-tools-linux-386.tgz
webdevdata-tools-linux-amd64.tgz
webdevdata-tools-windows-386.tgz
webdevdata-tools-windows-amd64.tgz
```

## ```wdd_meta_names [file]```

Checks HTML meta tags from ```file``` and prints a CSV with
```file,meta_name``` to ```STDOUT```.

example:

```bash
$ wdd_meta_names
webdevdata.org-2013-10-30-231036/19/jimsmarketingblog.com_19932518c6d628a198247a3f2a1322e6.html.txt
./webdevdata.org-2013-10-30-231036/19/jimsmarketingblog.com_19932518c6d628a198247a3f2a1322e6.html.txt,description
./webdevdata.org-2013-10-30-231036/19/jimsmarketingblog.com_19932518c6d628a198247a3f2a1322e6.html.txt,google-site-verification
./webdevdata.org-2013-10-30-231036/19/jimsmarketingblog.com_19932518c6d628a198247a3f2a1322e6.html.txt,google-site-verification
./webdevdata.org-2013-10-30-231036/19/jimsmarketingblog.com_19932518c6d628a198247a3f2a1322e6.html.txt,y_key
```

Generating CSV with all meta tag names from webdevdata.org crawl (using
GNU/Parallel instead of ```xargs``` to parallelize work):

```bash
$ find webdevdata.org-2013-10-30-231036 -name "*tml.txt" | parallel "wdd_meta_names >> meta_names.csv"
```
## ```wdd_html_manifest [file]```

Checks for html tag with manifest attribute from ```file``` and prints a CSV
with ```file,manifest_value``` to ```STDOUT```.

example:

```bash
$ wdd_html_manifest webdevdata.org-2013-10-30-231036/49/forecast.io_49bd380f592ae37fc74709838d2ace13.html.txt
webdevdata.org-2013-10-30-231036/49/forecast.io_49bd380f592ae37fc74709838d2ace13.html.txt,cache.desktop.manifest
```

## ```wdd_tag_count [file]```

Counts all HTML tags from ```file``` and prints a CSV with
```tag,count``` to ```STDOUT```.

example:

```bash
$ wdd_tag_count webdevdata.org-2013-10-30-231036/19/jimsmarketingblog.com_19932518c6d628a198247a3f2a1322e6.html.txt
link,16
title,1
span,29
label,1
script,15
li,40
strong,1
[...]
```

## Building tools

 1. ```go get github.com/webdevdata/webdevdata-tools```
 2. ```cd $GOPATH/src/github.com/webdevdata/webdevdata-tools```
 3. ```make all```
 4. Tools are in the build directory

You can use ```make release``` to generate cross-compiled binaries for Linux,
Windows and Mac.
