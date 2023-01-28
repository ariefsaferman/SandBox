## Why we need React Optimization?

In general, a fast website streamlines the browsing process and gives your users a better overall experience when viewing your site. Slow loading times have been shown as one of the factors that increase the user bounce rate of websites.

- Better Conversion Rate

A good user experience can encourage people to convert more easily. This may sound like something unbelievable, but consider this: would you bother buying something off an e-commerce site if it took you in exorable amount of time just to go from your shopping to the checkout page? Site performance and design does play a role in improving your conversion rate. Big sites like Amazon have constantly evolved over the years to provide a better shopping experience for its customers and you should consider doing the same.

- Speed affects the SEO rankings of a website

Google made it very clear that they operate by speed-obsession for every product on the web. Thus, for websites to rank well on Google, teams must ensure that web pages are optimized enough to load faster. Higher rankings lead to growth in organic traffic, which is very important for businesses.

## So how we benchmark our website?

- Lighthouse
- GTMetrix.com
- pagespeed.web.dev

Lighthouse

**[Lighthouse](https://chrome.google.com/webstore/detail/lighthouse/blipmdconlkpinefehnmjammfjpmpbjk?hl=en)** is an open-source tool that allows us to measure web performance with two main components:

- **A score** from 0 to 100 based on your page’s web performance
- **A report** with practical tips on how to improve your page’s performance
- [First Contentful Paint](https://developers.google.com/web/tools/lighthouse/audits/first-contentful-paint?utm_source=lighthouse&utm_medium=devtools): the amount of time it took for **any** bit of the first text or image to render.
- [First Meaningful Paint](https://developers.google.com/web/tools/lighthouse/audits/first-meaningful-paint?utm_source=lighthouse&utm_medium=devtools): the amount of time it took for a **complete** text or image to render.
- [Speed Index](https://developers.google.com/web/tools/lighthouse/audits/speed-index?utm_source=lighthouse&utm_medium=devtools): the amount of time it took for the last text or image to render, in seconds.
- [Time to Interactive](https://developers.google.com/web/tools/lighthouse/audits/time-to-interactive?utm_source=lighthouse&utm_medium=devtools): the amount of time it took for the page to become fully interactive.
- CPU Idle : First CPU Idle measures how long it takes a page to become *minimally* interactive. The strategies for improving First CPU Idle are largely the same as the strategies for improving TTI.
    - [https://developer.chrome.com/docs/lighthouse/performance/first-cpu-idle/](https://developer.chrome.com/docs/lighthouse/performance/first-cpu-idle/)
- Estimated Input Latency: is an estimate of how long your app takes to respond to user input during the busiest 5 second window of page load.

Tips you can do

1. Improve Bundle Size
    - Checking bundle size
        - [https://bundlephobia.com/](https://bundlephobia.com/)
        - VS Code extension import cost
            - ttps://marketplace.visualstudio.com/items?itemName=wix.vscode-import-cost
    - Analyze
        - Webpack bundle analyzer if using Create React App
            - [https://blog.jakoblind.no/webpack-bundle-analyzer/](https://blog.jakoblind.no/webpack-bundle-analyzer/)
        - Rollup if using vite
            - rollup plugin visualizer https://github.com/btd/rollup-plugin-visualizer
    - Remove unused CSS
        - Purge CSS
        - rollup-plugin-purgecss
            - https://www.npmjs.com/package/rollup-plugin-purgecss
    - Don’t use @import
    - Use “preload” with 3rd party font API / local
2. Code Splitting
    - Lazy load
    - react suspense
3. Image optimization
    - Use CDN
    - Use next-gen image format like webP, svg
    - Preload Image