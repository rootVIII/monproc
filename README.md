###### MONPROC

###### Get the project:
<pre>
  <code>
# Clone project:
git clone https://github.com/rootVIII/monproc.git

# Build and run (show top 50 processes):
cd &lt;project root&gt;
go build -o bin/monproc
./bin/monproc 50

# Build binary in ~/go/bin (available in path) and run (show top 50 processes):
cd &lt;project root&gt;
go install .
monproc 50
  </code>
</pre>


<b>References used:</b><br>
http://man7.org/linux/man-pages/man5/proc.5.html


<img src="https://github.com/rootVIII/monproc/blob/master/terminal_screenshot.png" alt="stdout">

This was developed/tested on Ubuntu 18.04.4 LTS and MacOS Big Sur.
<hr>
<b>Author: rootVIII 24OCT2019</b><br><br>

