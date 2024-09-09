"use strict";(self.webpackChunkdocsd=self.webpackChunkdocsd||[]).push([[313],{98:(e,n,t)=>{t.r(n),t.d(n,{assets:()=>a,contentTitle:()=>s,default:()=>u,frontMatter:()=>o,metadata:()=>l,toc:()=>c});var r=t(4848),i=t(8453);const o={id:"gotry",title:"GoTry Documentation",sidebar_label:"GoTry",slug:"/"},s="GoTry Documentation",l={id:"gotry",title:"GoTry Documentation",description:"Welcome to the official documentation for GoTry. This library provides a flexible and customizable way to manage retries and errors in Go applications.",source:"@site/docs/markdown-features.mdx",sourceDirName:".",slug:"/",permalink:"/gotry/docs/",draft:!1,unlisted:!1,editUrl:"https://github.com/facebook/docusaurus/tree/main/packages/create-docusaurus/templates/shared/docs/markdown-features.mdx",tags:[],version:"current",frontMatter:{id:"gotry",title:"GoTry Documentation",sidebar_label:"GoTry",slug:"/"},sidebar:"tutorialSidebar"},a={},c=[{value:"Introduction",id:"introduction",level:2},{value:"Features",id:"features",level:2},{value:"Installation",id:"installation",level:2},{value:"Usage",id:"usage",level:2},{value:"Parameters",id:"parameters",level:2},{value:"Contribution",id:"contribution",level:2},{value:"License",id:"license",level:2},{value:"Contact",id:"contact",level:2}];function d(e){const n={admonition:"admonition",code:"code",em:"em",h1:"h1",h2:"h2",header:"header",li:"li",ol:"ol",p:"p",pre:"pre",strong:"strong",ul:"ul",...(0,i.R)(),...e.components};return(0,r.jsxs)(r.Fragment,{children:[(0,r.jsx)(n.header,{children:(0,r.jsx)(n.h1,{id:"gotry-documentation",children:"GoTry Documentation"})}),"\n",(0,r.jsxs)(n.p,{children:["Welcome to the official documentation for ",(0,r.jsx)(n.strong,{children:"GoTry"}),". This library provides a flexible and customizable way to manage retries and errors in Go applications."]}),"\n",(0,r.jsx)(n.h2,{id:"introduction",children:"Introduction"}),"\n",(0,r.jsx)(n.p,{children:"GoTry simplifies retry logic in Go by providing a reusable, configurable, and fault-tolerant mechanism to handle transient failures, with optional backoff strategies and other enhancements."}),"\n",(0,r.jsx)(n.h2,{id:"features",children:"Features"}),"\n",(0,r.jsxs)(n.ul,{children:["\n",(0,r.jsxs)(n.li,{children:[(0,r.jsx)(n.strong,{children:"Configurable retry logic"}),": Customize retries with backoff strategies."]}),"\n",(0,r.jsxs)(n.li,{children:[(0,r.jsx)(n.strong,{children:"Jitter support"}),": Adds randomness to avoid retry storms."]}),"\n",(0,r.jsxs)(n.li,{children:[(0,r.jsx)(n.strong,{children:"Adjustable multiplier"}),": Define exponential growth for retries."]}),"\n",(0,r.jsxs)(n.li,{children:[(0,r.jsx)(n.strong,{children:"Simple API"}),": Easy-to-use API for configuring retries."]}),"\n"]}),"\n",(0,r.jsx)(n.h2,{id:"installation",children:"Installation"}),"\n",(0,r.jsx)(n.p,{children:"To install GoTry in your Go project, run:"}),"\n",(0,r.jsx)(n.pre,{children:(0,r.jsx)(n.code,{className:"language-sh",children:"go get github.com/PabloSanchi/gotry\n"})}),"\n",(0,r.jsx)(n.h2,{id:"usage",children:"Usage"}),"\n",(0,r.jsx)(n.p,{children:"Here's an example of how to use GoTry to manage retries in your Go application:"}),"\n",(0,r.jsx)(n.pre,{children:(0,r.jsx)(n.code,{className:"language-go",children:'package main\n\nimport (\n\t  "encoding/json"\n\t  "log"\n\t  "net/http"\n\t  "time"\n\n\t  retry "github.com/PabloSanchi/gotry"\n)\n\nconst (\n\t  url = "https://official-joke-api.appspot.com/random_joke"\n)\n\ntype Joke struct {\n\t  Type      string `json:"type"`\n\t  Setup     string `json:"setup"`\n\t  Punchline string `json:"punchline"`\n\t  Id        uint   `json:"id"`\n}\n\nfunc main() {\n    resp, err := retry.Retry(\n        func() (*http.Response, error) {\n            return http.Get(url)\n        },\n        retry.WithRetries(2),\n        retry.WithBackoff(2*time.Second),\n        retry.WithExponentialBackoff(),\n        retry.WithOnRetry(func(n uint, err error) {\n                log.Printf("Retrying request after error: %v", err)\n        }),\n    )\n\n    if err != nil {\n        log.Fatalf("Request failed: %v", err)\n    }\n    defer resp.Body.Close()\n\n    var joke Joke\n    if err := json.NewDecoder(resp.Body).Decode(&joke); err != nil {\n        log.Fatalf("Failed to decode response: %v", err)\n    }\n\n    log.Printf("Joke: %s", joke.Setup)\n    log.Printf("Punchline: %s", joke.Punchline)\n}\n'})}),"\n",(0,r.jsx)(n.h2,{id:"parameters",children:"Parameters"}),"\n",(0,r.jsx)(n.p,{children:"GoTry allows configuring various parameters for the retry logic:"}),"\n",(0,r.jsxs)(n.ul,{children:["\n",(0,r.jsxs)(n.li,{children:[(0,r.jsx)(n.code,{children:"Retries"}),": Set the maximum number of retries."]}),"\n",(0,r.jsxs)(n.li,{children:[(0,r.jsx)(n.code,{children:"Backoff"}),": Choose between different backoff strategies (exponential, constant, etc.)."]}),"\n",(0,r.jsxs)(n.li,{children:[(0,r.jsx)(n.code,{children:"BackoffStrategies"}),":","\n",(0,r.jsxs)(n.ul,{children:["\n",(0,r.jsxs)(n.li,{children:[(0,r.jsx)(n.code,{children:"ExponentialBackoff"}),": Increases the wait time exponentially between each retry."]}),"\n",(0,r.jsxs)(n.li,{children:[(0,r.jsx)(n.code,{children:"LinearBackoff"}),": Uses a multiplier to increase the wait time between retries. (backoff * multiplier) and multiplier is the retry attempt."]}),"\n",(0,r.jsxs)(n.li,{children:[(0,r.jsx)(n.code,{children:"CustomBackoff"}),": Implement your own backoff logic."]}),"\n"]}),"\n"]}),"\n",(0,r.jsxs)(n.li,{children:[(0,r.jsx)(n.code,{children:"MaxJitter"}),": Add randomness to the backoff to stagger the retries. (MaxJitter is the maximum value allowed to add to the backoff)."]}),"\n",(0,r.jsxs)(n.li,{children:[(0,r.jsx)(n.code,{children:"OnRetry"}),": Callback function to execute on each retry."]}),"\n",(0,r.jsxs)(n.li,{children:[(0,r.jsx)(n.code,{children:"RetryIf"}),": Function to determine if a retry should be attempted based on the error."]}),"\n",(0,r.jsxs)(n.li,{children:[(0,r.jsx)(n.code,{children:"Context"}),": Context (provided or created) to control the retry loop."]}),"\n"]}),"\n",(0,r.jsx)(n.pre,{children:(0,r.jsx)(n.code,{className:"language-go",children:"type RetryConfig struct {\n    retries         uint\n    backoff         time.Duration\n    backoffStrategy func(base time.Duration, n uint) time.Duration\n    backoffLimit    time.Duration // maximum backoff time allowed\n    maxJitter       time.Duration\n    onRetry         OnRetryFunc\n    retryIf         RetryIfFunc\n    context         context.Context\n}\n"})}),"\n",(0,r.jsx)(n.h2,{id:"contribution",children:"Contribution"}),"\n",(0,r.jsx)(n.p,{children:"We welcome contributions! To get started:"}),"\n",(0,r.jsxs)(n.ol,{children:["\n",(0,r.jsx)(n.li,{children:"Fork the repository"}),"\n",(0,r.jsx)(n.li,{children:"Make your changes"}),"\n",(0,r.jsx)(n.li,{children:"Submit a pull request."}),"\n"]}),"\n",(0,r.jsx)(n.admonition,{type:"note",children:(0,r.jsx)(n.p,{children:(0,r.jsx)(n.em,{children:"Please ensure your changes are well-tested; No tested code will be reviewed nor accepted"})})}),"\n",(0,r.jsx)(n.h2,{id:"license",children:"License"}),"\n",(0,r.jsx)(n.p,{children:"GoTry is licensed under the Apache License. See the LICENSE file for more details."}),"\n",(0,r.jsx)(n.h2,{id:"contact",children:"Contact"}),"\n",(0,r.jsx)(n.p,{children:"For any questions or issues, feel free to open an issue in our GitHub repository."}),"\n",(0,r.jsx)(n.p,{children:"Happy coding with GoTry! \ud83c\udf89"})]})}function u(e={}){const{wrapper:n}={...(0,i.R)(),...e.components};return n?(0,r.jsx)(n,{...e,children:(0,r.jsx)(d,{...e})}):d(e)}}}]);