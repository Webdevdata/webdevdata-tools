# Tools to extract data form webdevdata.org

Small set of command line tools written in Go to help extracting data
from webdevdata.org.

## Installation

 1. Clone the repository
 2. Run ```make all``` in the project directory
 3. Tools are in the build directory

## wdd_meta_names

```wdd_meta_names [file]```

Checks HTML meta tags from ```file``` and prints a CSV with
```file,meta_name``` to ```STDOUT```.

example:

```bash
$ wdd_meta_names webdevdata.org-2013-10-30-231036/19/jimsmarketingblog.com_19932518c6d628a198247a3f2a1322e6.html.txt
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

## wdd_tag_count

```wdd_tag_count [file]```

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

