###### MONPROC

<code>go get -v github.com/rootVIII/monproc</code><br>

<pre>
    <code>
As a standalone exe:

./monproc &lt;max records&gt;


In your own code:

maxRecords := 50
for _, p := range GetProcesses(maxRecords) {
		fmt.Printf(p)
	}
    </code>
</pre>

<b>References used:</b><br>
http://man7.org/linux/man-pages/man5/proc.5.html


<img src="https://github.com/rootVIII/monproc/blob/master/terminal_screenshot.png" alt="stdout">

This was developed on Ubuntu 18.04.4 LTS.
<hr>
<b>Author: rootVIII 24OCT2019</b><br><br>

