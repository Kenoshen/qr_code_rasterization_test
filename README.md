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
name | input | expected | cons2p | chrome | oksvg | imagemagick | resvg
-----------------------------------------------------------------------
frame.svg | ![img](input/frame.svg) | ![img](compare/frame.png) | ![oksvg](output/frame_oksvg.png) | ![imagemagick](output/frame_imagemagick.png) | ![resvg](output/frame_resvg.png) | ![cons2p](output/frame_cons2p.png) | ![chrome](output/frame_chrome.png)
gradient.svg | ![img](input/gradient.svg) | ![img](compare/gradient.png) | ![chrome](output/gradient_chrome.png) | ![oksvg](output/gradient_oksvg.png) | ![imagemagick](output/gradient_imagemagick.png) | ![resvg](output/gradient_resvg.png) | ![cons2p](output/gradient_cons2p.png)
text.svg | ![img](input/text.svg) | ![img](compare/text.png) | ![oksvg](output/text_oksvg.png) | ![imagemagick](output/text_imagemagick.png) | ![resvg](output/text_resvg.png) | ![cons2p](output/text_cons2p.png) | ![chrome](output/text_chrome.png)
-----------------------------------------------------------------------