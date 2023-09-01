# qr_code_rasterization_test

## Prerequisites

You will need to download these external CLI programs for this to run successfully.

- `brew install imagemagick`
- `brew install resvg`
- `npm install -g convert-svg-to-png`

## Run

- `go build rasterization`
- `./rasterization`

## Output

Look in the `output/` folder to find the pngs created and compare them to the originals from the `compare/` directory (values from the compare directory have been copied into the output directory for convenience)


## comparison


### frame.svg
`input`
![img](input/frame.svg)
`expected output`
![img](compare/frame.png)
`oksvg`
![oksvg](output/frame_oksvg.png)
`imagemagick`
![imagemagick](output/frame_imagemagick.png)
`resvg`
![resvg](output/frame_resvg.png)
`cons2p`
![cons2p](output/frame_cons2p.png)
`chrome`
![chrome](output/frame_chrome.png)


### gradient.svg
`input`
![img](input/gradient.svg)
`expected output`
![img](compare/gradient.png)
`resvg`
![resvg](output/gradient_resvg.png)
`cons2p`
![cons2p](output/gradient_cons2p.png)
`chrome`
![chrome](output/gradient_chrome.png)
`oksvg`
![oksvg](output/gradient_oksvg.png)
`imagemagick`
![imagemagick](output/gradient_imagemagick.png)


### text.svg
`input`
![img](input/text.svg)
`expected output`
![img](compare/text.png)
`imagemagick`
![imagemagick](output/text_imagemagick.png)
`resvg`
![resvg](output/text_resvg.png)
`cons2p`
![cons2p](output/text_cons2p.png)
`chrome`
![chrome](output/text_chrome.png)
`oksvg`
![oksvg](output/text_oksvg.png)