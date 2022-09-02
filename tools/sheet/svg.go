package main

import "fmt"
import "io/ioutil"
import "os"
import b64 "encoding/base64"

func png2svg(in string, out string, bank int) {
    payload, err := ioutil.ReadFile(in) 
    if err != nil {
      panic(err)
    }

	f, err := os.Create(out)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer f.Close()

    f.WriteString(svg_top)
    f.WriteString(b64.StdEncoding.EncodeToString(payload))
    f.WriteString(svg_middle)
    f.WriteString(fmt.Sprintf("%04x", bank << 8))
    f.WriteString(svg_end)
    
}

const svg_top = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<svg
   width="133.98253pt"
   height="136.40675pt"
   viewBox="0 0 47.266059 48.121267"
   version="1.1"
   id="svg5"
   sodipodi:docname="0x0000c.svg"
   inkscape:version="1.1.1 (3bf5ae0d25, 2021-09-20)"
   xmlns:inkscape="http://www.inkscape.org/namespaces/inkscape"
   xmlns:sodipodi="http://sodipodi.sourceforge.net/DTD/sodipodi-0.dtd"
   xmlns:xlink="http://www.w3.org/1999/xlink"
   xmlns="http://www.w3.org/2000/svg"
   xmlns:svg="http://www.w3.org/2000/svg">
  <sodipodi:namedview
     id="namedview7"
     pagecolor="#ffffff"
     bordercolor="#666666"
     borderopacity="1.0"
     inkscape:pageshadow="2"
     inkscape:pageopacity="0.0"
     inkscape:pagecheckerboard="0"
     inkscape:document-units="mm"
     showgrid="false"
     inkscape:zoom="3.4745541"
     inkscape:cx="103.32261"
     inkscape:cy="90.803018"
     inkscape:window-width="1907"
     inkscape:window-height="1045"
     inkscape:window-x="0"
     inkscape:window-y="0"
     inkscape:window-maximized="0"
     inkscape:current-layer="layer1"
     fit-margin-top="0"
     fit-margin-left="0"
     fit-margin-right="0"
     fit-margin-bottom="0"
     units="pt" />
  <defs
     id="defs2" />
  <g
     inkscape:label="Layer 1"
     inkscape:groupmode="layer"
     id="layer1"
     transform="translate(-47.252364,-100.20041)">
    <image
       width="35.241745"
       height="45.139061"
       preserveAspectRatio="none"
       image-rendering="optimizeSpeed"
       style="image-rendering:pixelated"
       xlink:href="data:image/png;base64,`

const svg_middle = `"
       id="image998"
       x="59.201832"
       y="103.10542" />
    <path
       style="fill:none;stroke:#000000;stroke-width:0.0499005;stroke-miterlimit:4;stroke-dasharray:none;stroke-opacity:1"
       d="m 61.346892,100.28408 v 47.96275"
       inkscape:label="MajorXDiv2"
       id="path832" />
    <path
       style="fill:none;stroke:#000000;stroke-width:0.0499005;stroke-miterlimit:4;stroke-dasharray:none;stroke-opacity:1"
       d="m 63.553338,100.28408 v 47.96275"
       inkscape:label="MajorXDiv3"
       id="path834" />
    <path
       style="fill:none;stroke:#000000;stroke-width:0.0499005;stroke-miterlimit:4;stroke-dasharray:none;stroke-opacity:1"
       d="m 65.759783,100.28408 v 47.96275"
       inkscape:label="MajorXDiv4"
       id="path836" />
    <path
       style="fill:none;stroke:#000000;stroke-width:0.0499005;stroke-miterlimit:4;stroke-dasharray:none;stroke-opacity:1"
       d="m 67.966231,100.28408 v 47.96275"
       inkscape:label="MajorXDiv5"
       id="path838" />
    <path
       style="fill:none;stroke:#000000;stroke-width:0.0499005;stroke-miterlimit:4;stroke-dasharray:none;stroke-opacity:1"
       d="m 70.172675,100.28408 v 47.96275"
       inkscape:label="MajorXDiv6"
       id="path840" />
    <path
       style="fill:none;stroke:#000000;stroke-width:0.0499005;stroke-miterlimit:4;stroke-dasharray:none;stroke-opacity:1"
       d="m 72.37912,100.28408 v 47.96275"
       inkscape:label="MajorXDiv7"
       id="path842" />
    <path
       style="fill:none;stroke:#000000;stroke-width:0.0499005;stroke-miterlimit:4;stroke-dasharray:none;stroke-opacity:1"
       d="m 74.585566,100.28408 v 47.96275"
       inkscape:label="MajorXDiv8"
       id="path844" />
    <path
       style="fill:none;stroke:#000000;stroke-width:0.0499005;stroke-miterlimit:4;stroke-dasharray:none;stroke-opacity:1"
       d="m 76.792011,100.28408 v 47.96275"
       inkscape:label="MajorXDiv9"
       id="path846" />
    <path
       style="fill:none;stroke:#000000;stroke-width:0.0499005;stroke-miterlimit:4;stroke-dasharray:none;stroke-opacity:1"
       d="m 78.998458,100.28408 v 47.96275"
       inkscape:label="MajorXDiv10"
       id="path848" />
    <path
       style="fill:none;stroke:#000000;stroke-width:0.0499005;stroke-miterlimit:4;stroke-dasharray:none;stroke-opacity:1"
       d="m 81.204902,100.28408 v 47.96275"
       inkscape:label="MajorXDiv11"
       id="path850" />
    <path
       style="fill:none;stroke:#000000;stroke-width:0.0499005;stroke-miterlimit:4;stroke-dasharray:none;stroke-opacity:1"
       d="m 83.411354,100.28408 v 47.96275"
       inkscape:label="MajorXDiv12"
       id="path852" />
    <path
       style="fill:none;stroke:#000000;stroke-width:0.0499005;stroke-miterlimit:4;stroke-dasharray:none;stroke-opacity:1"
       d="m 85.617793,100.28408 v 47.96275"
       inkscape:label="MajorXDiv13"
       id="path854" />
    <path
       style="fill:none;stroke:#000000;stroke-width:0.0499005;stroke-miterlimit:4;stroke-dasharray:none;stroke-opacity:1"
       d="m 87.824239,100.28408 v 47.96275"
       inkscape:label="MajorXDiv14"
       id="path856" />
    <path
       style="fill:none;stroke:#000000;stroke-width:0.0499005;stroke-miterlimit:4;stroke-dasharray:none;stroke-opacity:1"
       d="m 90.030689,100.28408 v 47.96275"
       inkscape:label="MajorXDiv15"
       id="path858" />
    <path
       style="fill:none;stroke:#000000;stroke-width:0.0499005;stroke-miterlimit:4;stroke-dasharray:none;stroke-opacity:1"
       d="m 92.237125,100.28408 v 47.96275"
       inkscape:label="MajorXDiv16"
       id="path860" />
    <g
       id="g27601"
       style="stroke:#000000;stroke-width:0.3;stroke-miterlimit:4;stroke-dasharray:none;stroke-opacity:1"
       transform="matrix(0.44128915,0,0,0.56426752,33.113119,45.343226)">
      <path
         style="fill:none;stroke:#000000;stroke-width:0.3;stroke-miterlimit:4;stroke-dasharray:none;stroke-opacity:1"
         d="M 58.980209,97.366669 V 182.36667"
         inkscape:label="MajorXDiv1"
         id="path830" />
      <path
         style="fill:none;stroke:#000000;stroke-width:0.3;stroke-miterlimit:4;stroke-dasharray:none;stroke-opacity:1"
         d="M 53.980209,102.36667 H 138.98021"
         inkscape:label="MajorYDiv1"
         id="path864" />
    </g>
    <path
       style="fill:none;stroke:#000000;stroke-width:0.0499005;stroke-miterlimit:4;stroke-dasharray:none;stroke-opacity:1"
       d="M 56.934,105.92676 H 94.443574"
       inkscape:label="MajorYDiv2"
       id="path866" />
    <path
       style="fill:none;stroke:#000000;stroke-width:0.0499005;stroke-miterlimit:4;stroke-dasharray:none;stroke-opacity:1"
       d="M 56.934,108.7481 H 94.443574"
       inkscape:label="MajorYDiv3"
       id="path868" />
    <path
       style="fill:none;stroke:#000000;stroke-width:0.0499005;stroke-miterlimit:4;stroke-dasharray:none;stroke-opacity:1"
       d="M 56.934,111.56943 H 94.443574"
       inkscape:label="MajorYDiv4"
       id="path870" />
    <path
       style="fill:none;stroke:#000000;stroke-width:0.0499005;stroke-miterlimit:4;stroke-dasharray:none;stroke-opacity:1"
       d="M 56.934,114.39076 H 94.443574"
       inkscape:label="MajorYDiv5"
       id="path872" />
    <path
       style="fill:none;stroke:#000000;stroke-width:0.0499005;stroke-miterlimit:4;stroke-dasharray:none;stroke-opacity:1"
       d="M 56.934,117.2121 H 94.443574"
       inkscape:label="MajorYDiv6"
       id="path874" />
    <path
       style="fill:none;stroke:#000000;stroke-width:0.0499005;stroke-miterlimit:4;stroke-dasharray:none;stroke-opacity:1"
       d="M 56.934,120.03345 H 94.443574"
       inkscape:label="MajorYDiv7"
       id="path876" />
    <path
       style="fill:none;stroke:#000000;stroke-width:0.0499005;stroke-miterlimit:4;stroke-dasharray:none;stroke-opacity:1"
       d="M 56.934,122.85479 H 94.443574"
       inkscape:label="MajorYDiv8"
       id="path878" />
    <path
       style="fill:none;stroke:#000000;stroke-width:0.0499005;stroke-miterlimit:4;stroke-dasharray:none;stroke-opacity:1"
       d="M 56.934,125.67612 H 94.443574"
       inkscape:label="MajorYDiv9"
       id="path880" />
    <path
       style="fill:none;stroke:#000000;stroke-width:0.0499005;stroke-miterlimit:4;stroke-dasharray:none;stroke-opacity:1"
       d="M 56.934,128.49746 H 94.443574"
       inkscape:label="MajorYDiv10"
       id="path882" />
    <path
       style="fill:none;stroke:#000000;stroke-width:0.0499005;stroke-miterlimit:4;stroke-dasharray:none;stroke-opacity:1"
       d="M 56.934,131.31879 H 94.443574"
       inkscape:label="MajorYDiv11"
       id="path884" />
    <path
       style="fill:none;stroke:#000000;stroke-width:0.0499005;stroke-miterlimit:4;stroke-dasharray:none;stroke-opacity:1"
       d="M 56.934,134.14013 H 94.443574"
       inkscape:label="MajorYDiv12"
       id="path886" />
    <path
       style="fill:none;stroke:#000000;stroke-width:0.0499005;stroke-miterlimit:4;stroke-dasharray:none;stroke-opacity:1"
       d="M 56.934,136.96148 H 94.443574"
       inkscape:label="MajorYDiv13"
       id="path888" />
    <path
       style="fill:none;stroke:#000000;stroke-width:0.0499005;stroke-miterlimit:4;stroke-dasharray:none;stroke-opacity:1"
       d="M 56.934,139.78282 H 94.443574"
       inkscape:label="MajorYDiv14"
       id="path890" />
    <path
       style="fill:none;stroke:#000000;stroke-width:0.0499005;stroke-miterlimit:4;stroke-dasharray:none;stroke-opacity:1"
       d="M 56.934,142.60414 H 94.443574"
       inkscape:label="MajorYDiv15"
       id="path892" />
    <path
       style="fill:none;stroke:#000000;stroke-width:0.0499005;stroke-miterlimit:4;stroke-dasharray:none;stroke-opacity:1"
       d="M 56.934,145.42549 H 94.443574"
       inkscape:label="MajorYDiv16"
       id="path894" />
    <path
       id="rect898"
       style="fill:none;stroke:#000000;stroke-width:0.0499005"
       inkscape:label="Border"
       d="m 56.934,100.28408 h 37.509574 v 47.96275 H 56.934 Z" />
    <path
       id="rect53037"
       style="fill:none;stroke:#000000;stroke-width:0.15;stroke-miterlimit:4;stroke-dasharray:none"
       d="M 47.327364,100.27541 H 56.934 v 8.47269 h -9.606636 z" />
    <path
       style="fill:none;stroke:#000000;stroke-width:0.149701;stroke-linecap:butt;stroke-linejoin:miter;stroke-miterlimit:4;stroke-dasharray:none;stroke-opacity:1"
       d="M 56.934,103.10542 H 47.34652"
       id="path54934" />
    <path
       id="rect57414"
       style="fill:none;stroke:#000000;stroke-width:0.149701"
       d="m 56.934,100.28408 h 37.509574 v 47.96275 H 56.934 Z" />
    <text
       xml:space="preserve"
       style="font-style:normal;font-weight:normal;font-size:1.76389px;line-height:1.25;font-family:sans-serif;fill:#000000;fill-opacity:1;stroke:none;stroke-width:0.0640422"
       x="48.649235"
       y="106.51819"
       id="text8779"><tspan
         sodipodi:role="line"
         id="tspan8777"
         style="font-style:normal;font-variant:normal;font-weight:normal;font-stretch:normal;font-size:1.76389px;font-family:'Roboto Mono';-inkscape-font-specification:'Roboto Mono';stroke-width:0.0640422"
         x="48.649235"
         y="106.51819">0x`








const svg_end = `</tspan></text>
    <text
       xml:space="preserve"
       style="font-style:normal;font-weight:normal;font-size:1.76389px;line-height:1.25;font-family:sans-serif;fill:#000000;fill-opacity:1;stroke:none;stroke-width:0.0640422"
       x="47.727791"
       y="102.28716"
       id="text8779-4"><tspan
         sodipodi:role="line"
         id="tspan8777-2"
         style="font-style:normal;font-variant:normal;font-weight:bold;font-stretch:normal;font-size:1.76389px;font-family:'Roboto Mono';-inkscape-font-specification:'Roboto Mono Bold';stroke-width:0.0640422"
         x="47.727791"
         y="102.28716">BANK</tspan></text>
    <text
       xml:space="preserve"
       style="font-style:normal;font-weight:normal;font-size:1.41111px;line-height:1.25;font-family:sans-serif;fill:#000000;fill-opacity:1;stroke:none;stroke-width:0.264583"
       x="59.461987"
       y="102.28972"
       id="text13476"><tspan
         sodipodi:role="line"
         id="tspan13474"
         style="font-style:normal;font-variant:normal;font-weight:normal;font-stretch:normal;font-size:1.41111px;font-family:'Roboto Mono';-inkscape-font-specification:'Roboto Mono';stroke-width:0.264583"
         x="59.461987"
         y="102.28972">00</tspan></text>
    <text
       xml:space="preserve"
       style="font-style:normal;font-weight:normal;font-size:1.41111px;line-height:1.25;font-family:sans-serif;fill:#000000;fill-opacity:1;stroke:none;stroke-width:0.264583"
       x="57.163712"
       y="105.19036"
       id="text13476-9"><tspan
         sodipodi:role="line"
         id="tspan13474-85"
         style="font-style:normal;font-variant:normal;font-weight:normal;font-stretch:normal;font-size:1.41111px;font-family:'Roboto Mono';-inkscape-font-specification:'Roboto Mono';stroke-width:0.264583"
         x="57.163712"
         y="105.19036">00</tspan></text>
    <text
       xml:space="preserve"
       style="font-style:normal;font-weight:normal;font-size:1.41111px;line-height:1.25;font-family:sans-serif;fill:#000000;fill-opacity:1;stroke:none;stroke-width:0.264583"
       x="61.675983"
       y="102.28972"
       id="text13476-7"><tspan
         sodipodi:role="line"
         id="tspan13474-7"
         style="font-style:normal;font-variant:normal;font-weight:normal;font-stretch:normal;font-size:1.41111px;font-family:'Roboto Mono';-inkscape-font-specification:'Roboto Mono';stroke-width:0.264583"
         x="61.675983"
         y="102.28972">01</tspan></text>
    <text
       xml:space="preserve"
       style="font-style:normal;font-weight:normal;font-size:1.41111px;line-height:1.25;font-family:sans-serif;fill:#000000;fill-opacity:1;stroke:none;stroke-width:0.264583"
       x="63.673626"
       y="102.28972"
       id="text13476-75"><tspan
         sodipodi:role="line"
         id="tspan13474-8"
         style="font-style:normal;font-variant:normal;font-weight:normal;font-stretch:normal;font-size:1.41111px;font-family:'Roboto Mono';-inkscape-font-specification:'Roboto Mono';stroke-width:0.264583"
         x="63.673626"
         y="102.28972">02</tspan></text>
    <text
       xml:space="preserve"
       style="font-style:normal;font-weight:normal;font-size:1.41111px;line-height:1.25;font-family:sans-serif;fill:#000000;fill-opacity:1;stroke:none;stroke-width:0.264583"
       x="65.873154"
       y="102.28972"
       id="text13476-6"><tspan
         sodipodi:role="line"
         id="tspan13474-4"
         style="font-style:normal;font-variant:normal;font-weight:normal;font-stretch:normal;font-size:1.41111px;font-family:'Roboto Mono';-inkscape-font-specification:'Roboto Mono';stroke-width:0.264583"
         x="65.873154"
         y="102.28972">03</tspan></text>
    <text
       xml:space="preserve"
       style="font-style:normal;font-weight:normal;font-size:1.41111px;line-height:1.25;font-family:sans-serif;fill:#000000;fill-opacity:1;stroke:none;stroke-width:0.264583"
       x="68.038231"
       y="102.28972"
       id="text13476-2"><tspan
         sodipodi:role="line"
         id="tspan13474-3"
         style="font-style:normal;font-variant:normal;font-weight:normal;font-stretch:normal;font-size:1.41111px;font-family:'Roboto Mono';-inkscape-font-specification:'Roboto Mono';stroke-width:0.264583"
         x="68.038231"
         y="102.28972">04</tspan></text>
    <text
       xml:space="preserve"
       style="font-style:normal;font-weight:normal;font-size:1.41111px;line-height:1.25;font-family:sans-serif;fill:#000000;fill-opacity:1;stroke:none;stroke-width:0.264583"
       x="70.279099"
       y="102.28972"
       id="text13476-26"><tspan
         sodipodi:role="line"
         id="tspan13474-1"
         style="font-style:normal;font-variant:normal;font-weight:normal;font-stretch:normal;font-size:1.41111px;font-family:'Roboto Mono';-inkscape-font-specification:'Roboto Mono';stroke-width:0.264583"
         x="70.279099"
         y="102.28972">05</tspan></text>
    <text
       xml:space="preserve"
       style="font-style:normal;font-weight:normal;font-size:1.41111px;line-height:1.25;font-family:sans-serif;fill:#000000;fill-opacity:1;stroke:none;stroke-width:0.264583"
       x="72.503433"
       y="102.28972"
       id="text13476-8"><tspan
         sodipodi:role="line"
         id="tspan13474-9"
         style="font-style:normal;font-variant:normal;font-weight:normal;font-stretch:normal;font-size:1.41111px;font-family:'Roboto Mono';-inkscape-font-specification:'Roboto Mono';stroke-width:0.264583"
         x="72.503433"
         y="102.28972">06</tspan></text>
    <text
       xml:space="preserve"
       style="font-style:normal;font-weight:normal;font-size:1.41111px;line-height:1.25;font-family:sans-serif;fill:#000000;fill-opacity:1;stroke:none;stroke-width:0.264583"
       x="74.698822"
       y="102.28972"
       id="text13476-0"><tspan
         sodipodi:role="line"
         id="tspan13474-91"
         style="font-style:normal;font-variant:normal;font-weight:normal;font-stretch:normal;font-size:1.41111px;font-family:'Roboto Mono';-inkscape-font-specification:'Roboto Mono';stroke-width:0.264583"
         x="74.698822"
         y="102.28972">07</tspan></text>
    <text
       xml:space="preserve"
       style="font-style:normal;font-weight:normal;font-size:1.41111px;line-height:1.25;font-family:sans-serif;fill:#000000;fill-opacity:1;stroke:none;stroke-width:0.264583"
       x="76.918335"
       y="102.28972"
       id="text13476-23"><tspan
         sodipodi:role="line"
         id="tspan13474-6"
         style="font-style:normal;font-variant:normal;font-weight:normal;font-stretch:normal;font-size:1.41111px;font-family:'Roboto Mono';-inkscape-font-specification:'Roboto Mono';stroke-width:0.264583"
         x="76.918335"
         y="102.28972">08</tspan></text>
    <text
       xml:space="preserve"
       style="font-style:normal;font-weight:normal;font-size:1.41111px;line-height:1.25;font-family:sans-serif;fill:#000000;fill-opacity:1;stroke:none;stroke-width:0.264583"
       x="79.142662"
       y="102.28972"
       id="text13476-1"><tspan
         sodipodi:role="line"
         id="tspan13474-60"
         style="font-style:normal;font-variant:normal;font-weight:normal;font-stretch:normal;font-size:1.41111px;font-family:'Roboto Mono';-inkscape-font-specification:'Roboto Mono';stroke-width:0.264583"
         x="79.142662"
         y="102.28972">09</tspan></text>
    <text
       xml:space="preserve"
       style="font-style:normal;font-weight:normal;font-size:1.41111px;line-height:1.25;font-family:sans-serif;fill:#000000;fill-opacity:1;stroke:none;stroke-width:0.264583"
       x="81.340813"
       y="102.28972"
       id="text13476-86"><tspan
         sodipodi:role="line"
         id="tspan13474-70"
         style="font-style:normal;font-variant:normal;font-weight:normal;font-stretch:normal;font-size:1.41111px;font-family:'Roboto Mono';-inkscape-font-specification:'Roboto Mono';stroke-width:0.264583"
         x="81.340813"
         y="102.28972">0A</tspan></text>
    <text
       xml:space="preserve"
       style="font-style:normal;font-weight:normal;font-size:1.41111px;line-height:1.25;font-family:sans-serif;fill:#000000;fill-opacity:1;stroke:none;stroke-width:0.264583"
       x="83.609932"
       y="102.28972"
       id="text13476-5"><tspan
         sodipodi:role="line"
         id="tspan13474-93"
         style="font-style:normal;font-variant:normal;font-weight:normal;font-stretch:normal;font-size:1.41111px;font-family:'Roboto Mono';-inkscape-font-specification:'Roboto Mono';stroke-width:0.264583"
         x="83.609932"
         y="102.28972">0B</tspan></text>
    <text
       xml:space="preserve"
       style="font-style:normal;font-weight:normal;font-size:1.41111px;line-height:1.25;font-family:sans-serif;fill:#000000;fill-opacity:1;stroke:none;stroke-width:0.264583"
       x="85.845306"
       y="102.28972"
       id="text13476-4"><tspan
         sodipodi:role="line"
         id="tspan13474-0"
         style="font-style:normal;font-variant:normal;font-weight:normal;font-stretch:normal;font-size:1.41111px;font-family:'Roboto Mono';-inkscape-font-specification:'Roboto Mono';stroke-width:0.264583"
         x="85.845306"
         y="102.28972">0C</tspan></text>
    <text
       xml:space="preserve"
       style="font-style:normal;font-weight:normal;font-size:1.41111px;line-height:1.25;font-family:sans-serif;fill:#000000;fill-opacity:1;stroke:none;stroke-width:0.264583"
       x="88.079285"
       y="102.28972"
       id="text13476-01"><tspan
         sodipodi:role="line"
         id="tspan13474-33"
         style="font-style:normal;font-variant:normal;font-weight:normal;font-stretch:normal;font-size:1.41111px;font-family:'Roboto Mono';-inkscape-font-specification:'Roboto Mono';stroke-width:0.264583"
         x="88.079285"
         y="102.28972">0D</tspan></text>
    <text
       xml:space="preserve"
       style="font-style:normal;font-weight:normal;font-size:1.41111px;line-height:1.25;font-family:sans-serif;fill:#000000;fill-opacity:1;stroke:none;stroke-width:0.264583"
       x="90.326355"
       y="102.28972"
       id="text13476-68"><tspan
         sodipodi:role="line"
         id="tspan13474-5"
         style="font-style:normal;font-variant:normal;font-weight:normal;font-stretch:normal;font-size:1.41111px;font-family:'Roboto Mono';-inkscape-font-specification:'Roboto Mono';stroke-width:0.264583"
         x="90.326355"
         y="102.28972">0E</tspan></text>
    <text
       xml:space="preserve"
       style="font-style:normal;font-weight:normal;font-size:1.41111px;line-height:1.25;font-family:sans-serif;fill:#000000;fill-opacity:1;stroke:none;stroke-width:0.264583"
       x="92.532082"
       y="102.28972"
       id="text13476-21"><tspan
         sodipodi:role="line"
         id="tspan13474-99"
         style="font-style:normal;font-variant:normal;font-weight:normal;font-stretch:normal;font-size:1.41111px;font-family:'Roboto Mono';-inkscape-font-specification:'Roboto Mono';stroke-width:0.264583"
         x="92.532082"
         y="102.28972">0F</tspan></text>
    <text
       xml:space="preserve"
       style="font-style:normal;font-weight:normal;font-size:1.41111px;line-height:1.25;font-family:sans-serif;fill:#000000;fill-opacity:1;stroke:none;stroke-width:0.264583"
       x="57.1203"
       y="107.99264"
       id="text13476-76"><tspan
         sodipodi:role="line"
         id="tspan13474-92"
         style="font-style:normal;font-variant:normal;font-weight:normal;font-stretch:normal;font-size:1.41111px;font-family:'Roboto Mono';-inkscape-font-specification:'Roboto Mono';stroke-width:0.264583"
         x="57.1203"
         y="107.99264">10</tspan></text>
    <text
       xml:space="preserve"
       style="font-style:normal;font-weight:normal;font-size:1.41111px;line-height:1.25;font-family:sans-serif;fill:#000000;fill-opacity:1;stroke:none;stroke-width:0.264583"
       x="57.205051"
       y="110.79491"
       id="text13476-53"><tspan
         sodipodi:role="line"
         id="tspan13474-2"
         style="font-style:normal;font-variant:normal;font-weight:normal;font-stretch:normal;font-size:1.41111px;font-family:'Roboto Mono';-inkscape-font-specification:'Roboto Mono';stroke-width:0.264583"
         x="57.205051"
         y="110.79491">20</tspan></text>
    <text
       xml:space="preserve"
       style="font-style:normal;font-weight:normal;font-size:1.41111px;line-height:1.25;font-family:sans-serif;fill:#000000;fill-opacity:1;stroke:none;stroke-width:0.264583"
       x="57.198849"
       y="113.5972"
       id="text13476-79"><tspan
         sodipodi:role="line"
         id="tspan13474-59"
         style="font-style:normal;font-variant:normal;font-weight:normal;font-stretch:normal;font-size:1.41111px;font-family:'Roboto Mono';-inkscape-font-specification:'Roboto Mono';stroke-width:0.264583"
         x="57.198849"
         y="113.5972">30</tspan></text>
    <text
       xml:space="preserve"
       style="font-style:normal;font-weight:normal;font-size:1.41111px;line-height:1.25;font-family:sans-serif;fill:#000000;fill-opacity:1;stroke:none;stroke-width:0.264583"
       x="57.211941"
       y="116.39948"
       id="text13476-3"><tspan
         sodipodi:role="line"
         id="tspan13474-81"
         style="font-style:normal;font-variant:normal;font-weight:normal;font-stretch:normal;font-size:1.41111px;font-family:'Roboto Mono';-inkscape-font-specification:'Roboto Mono';stroke-width:0.264583"
         x="57.211941"
         y="116.39948">40</tspan></text>
    <text
       xml:space="preserve"
       style="font-style:normal;font-weight:normal;font-size:1.41111px;line-height:1.25;font-family:sans-serif;fill:#000000;fill-opacity:1;stroke:none;stroke-width:0.264583"
       x="57.134773"
       y="119.20175"
       id="text13476-05"><tspan
         sodipodi:role="line"
         id="tspan13474-52"
         style="font-style:normal;font-variant:normal;font-weight:normal;font-stretch:normal;font-size:1.41111px;font-family:'Roboto Mono';-inkscape-font-specification:'Roboto Mono';stroke-width:0.264583"
         x="57.134773"
         y="119.20175">50</tspan></text>
    <text
       xml:space="preserve"
       style="font-style:normal;font-weight:normal;font-size:1.41111px;line-height:1.25;font-family:sans-serif;fill:#000000;fill-opacity:1;stroke:none;stroke-width:0.264583"
       x="57.166466"
       y="122.00403"
       id="text13476-09"><tspan
         sodipodi:role="line"
         id="tspan13474-94"
         style="font-style:normal;font-variant:normal;font-weight:normal;font-stretch:normal;font-size:1.41111px;font-family:'Roboto Mono';-inkscape-font-specification:'Roboto Mono';stroke-width:0.264583"
         x="57.166466"
         y="122.00403">60</tspan></text>
    <text
       xml:space="preserve"
       style="font-style:normal;font-weight:normal;font-size:1.41111px;line-height:1.25;font-family:sans-serif;fill:#000000;fill-opacity:1;stroke:none;stroke-width:0.264583"
       x="57.186447"
       y="124.80631"
       id="text13476-73"><tspan
         sodipodi:role="line"
         id="tspan13474-58"
         style="font-style:normal;font-variant:normal;font-weight:normal;font-stretch:normal;font-size:1.41111px;font-family:'Roboto Mono';-inkscape-font-specification:'Roboto Mono';stroke-width:0.264583"
         x="57.186447"
         y="124.80631">70</tspan></text>
    <text
       xml:space="preserve"
       style="font-style:normal;font-weight:normal;font-size:1.41111px;line-height:1.25;font-family:sans-serif;fill:#000000;fill-opacity:1;stroke:none;stroke-width:0.264583"
       x="57.141663"
       y="127.60859"
       id="text13476-43"><tspan
         sodipodi:role="line"
         id="tspan13474-40"
         style="font-style:normal;font-variant:normal;font-weight:normal;font-stretch:normal;font-size:1.41111px;font-family:'Roboto Mono';-inkscape-font-specification:'Roboto Mono';stroke-width:0.264583"
         x="57.141663"
         y="127.60859">80</tspan></text>
    <text
       xml:space="preserve"
       style="font-style:normal;font-weight:normal;font-size:1.41111px;line-height:1.25;font-family:sans-serif;fill:#000000;fill-opacity:1;stroke:none;stroke-width:0.264583"
       x="57.160954"
       y="130.41087"
       id="text13476-32"><tspan
         sodipodi:role="line"
         id="tspan13474-64"
         style="font-style:normal;font-variant:normal;font-weight:normal;font-stretch:normal;font-size:1.41111px;font-family:'Roboto Mono';-inkscape-font-specification:'Roboto Mono';stroke-width:0.264583"
         x="57.160954"
         y="130.41087">90</tspan></text>
    <text
       xml:space="preserve"
       style="font-style:normal;font-weight:normal;font-size:1.41111px;line-height:1.25;font-family:sans-serif;fill:#000000;fill-opacity:1;stroke:none;stroke-width:0.264583"
       x="57.207806"
       y="133.21315"
       id="text13476-090"><tspan
         sodipodi:role="line"
         id="tspan13474-942"
         style="font-style:normal;font-variant:normal;font-weight:normal;font-stretch:normal;font-size:1.41111px;font-family:'Roboto Mono';-inkscape-font-specification:'Roboto Mono';stroke-width:0.264583"
         x="57.207806"
         y="133.21315">A0</tspan></text>
    <text
       xml:space="preserve"
       style="font-style:normal;font-weight:normal;font-size:1.41111px;line-height:1.25;font-family:sans-serif;fill:#000000;fill-opacity:1;stroke:none;stroke-width:0.264583"
       x="57.145107"
       y="136.01543"
       id="text13476-64"><tspan
         sodipodi:role="line"
         id="tspan13474-32"
         style="font-style:normal;font-variant:normal;font-weight:normal;font-stretch:normal;font-size:1.41111px;font-family:'Roboto Mono';-inkscape-font-specification:'Roboto Mono';stroke-width:0.264583"
         x="57.145107"
         y="136.01543">B0</tspan></text>
    <text
       xml:space="preserve"
       style="font-style:normal;font-weight:normal;font-size:1.41111px;line-height:1.25;font-family:sans-serif;fill:#000000;fill-opacity:1;stroke:none;stroke-width:0.264583"
       x="57.189892"
       y="138.8177"
       id="text13476-42"><tspan
         sodipodi:role="line"
         id="tspan13474-400"
         style="font-style:normal;font-variant:normal;font-weight:normal;font-stretch:normal;font-size:1.41111px;font-family:'Roboto Mono';-inkscape-font-specification:'Roboto Mono';stroke-width:0.264583"
         x="57.189892"
         y="138.8177">C0</tspan></text>
    <text
       xml:space="preserve"
       style="font-style:normal;font-weight:normal;font-size:1.41111px;line-height:1.25;font-family:sans-serif;fill:#000000;fill-opacity:1;stroke:none;stroke-width:0.264583"
       x="57.156818"
       y="141.61998"
       id="text13476-67"><tspan
         sodipodi:role="line"
         id="tspan13474-14"
         style="font-style:normal;font-variant:normal;font-weight:normal;font-stretch:normal;font-size:1.41111px;font-family:'Roboto Mono';-inkscape-font-specification:'Roboto Mono';stroke-width:0.264583"
         x="57.156818"
         y="141.61998">D0</tspan></text>
    <text
       xml:space="preserve"
       style="font-style:normal;font-weight:normal;font-size:1.41111px;line-height:1.25;font-family:sans-serif;fill:#000000;fill-opacity:1;stroke:none;stroke-width:0.264583"
       x="57.138218"
       y="144.42227"
       id="text13476-66"><tspan
         sodipodi:role="line"
         id="tspan13474-333"
         style="font-style:normal;font-variant:normal;font-weight:normal;font-stretch:normal;font-size:1.41111px;font-family:'Roboto Mono';-inkscape-font-specification:'Roboto Mono';stroke-width:0.264583"
         x="57.138218"
         y="144.42227">E0</tspan></text>
    <text
       xml:space="preserve"
       style="font-style:normal;font-weight:normal;font-size:1.41111px;line-height:1.25;font-family:sans-serif;fill:#000000;fill-opacity:1;stroke:none;stroke-width:0.264583"
       x="57.132015"
       y="147.22455"
       id="text13476-71"><tspan
         sodipodi:role="line"
         id="tspan13474-07"
         style="font-style:normal;font-variant:normal;font-weight:normal;font-stretch:normal;font-size:1.41111px;font-family:'Roboto Mono';-inkscape-font-specification:'Roboto Mono';stroke-width:0.264583"
         x="57.132015"
         y="147.22455">F0</tspan></text>
  </g>
</svg>`