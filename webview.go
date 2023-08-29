package main

const html = `<html>
<body>
<style>
  .thing {
    background-color: #000;
	width: 300px;
    height: 300px;
  }
</style>
<button id="increment">Tap me</button>
<div>You tapped <span id="count">0</span> time(s).</div>
<div class="thing"><svg>%s</svg></div>
<script>
  const [incrementElement, countElement] =
    document.querySelectorAll("#increment, #count");
  document.addEventListener("DOMContentLoaded", () => {
    incrementElement.addEventListener("click", () => {
      window.increment().then(result => {
        countElement.textContent = result.count;
      });
    });
  });
</script></body></html>`

type Webview struct {
	Count int `json:"count"`
}

func (v Webview) Run(inputFilename string, data []byte) ([]byte, error) {
	//count := 0
	//w := webview.New(false)
	//defer w.Destroy()
	//w.SetTitle("Bind Example")
	//w.SetSize(1024, 1024, webview.HintNone)
	//
	//// A binding that increments a value and immediately returns the new value.
	//w.Bind("increment", func() Webview {
	//	count++
	//	return Webview{Count: count}
	//})
	//
	//w.SetHtml(fmt.Sprintf(html, string(data)))
	//w.Run()
	return nil, nil
}
