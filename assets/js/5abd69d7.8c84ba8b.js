"use strict";(self.webpackChunkdocs=self.webpackChunkdocs||[]).push([[6586],{8248:(e,n,t)=>{t.r(n),t.d(n,{assets:()=>l,contentTitle:()=>a,default:()=>u,frontMatter:()=>o,metadata:()=>r,toc:()=>c});var i=t(1527),s=t(4744);const o={sidebar_label:"Dependcy Injection",title:"Dependcy Injection"},a=void 0,r={id:"learn/fundamentals/di",title:"Dependcy Injection",description:"Dependency Injection (DI) is a technique used in software development to manage dependencies between different components or classes. It involves providing dependencies to a class from the outside, rather than the class creating them internally. This enhances flexibility, testability, and maintainability of the code.",source:"@site/docs/02-learn/01-fundamentals/01-di.mdx",sourceDirName:"02-learn/01-fundamentals",slug:"/learn/fundamentals/di",permalink:"/kit/docs/learn/fundamentals/di",draft:!1,unlisted:!1,editUrl:"https://github.com/go-saas/kit/tree/main/docs/docs/02-learn/01-fundamentals/01-di.mdx",tags:[],version:"current",sidebarPosition:1,frontMatter:{sidebar_label:"Dependcy Injection",title:"Dependcy Injection"},sidebar:"tutorialSidebar",previous:{title:"Fundamentals",permalink:"/kit/docs/category/fundamentals"},next:{title:"Gateway",permalink:"/kit/docs/learn/fundamentals/gateway"}},l={},c=[{value:"Wire",id:"wire",level:2},{value:"defval/di",id:"defvaldi",level:2}];function d(e){const n={a:"a",admonition:"admonition",blockquote:"blockquote",h2:"h2",li:"li",p:"p",ul:"ul",...(0,s.a)(),...e.components};return(0,i.jsxs)(i.Fragment,{children:[(0,i.jsx)(n.p,{children:"Dependency Injection (DI) is a technique used in software development to manage dependencies between different components or classes. It involves providing dependencies to a class from the outside, rather than the class creating them internally. This enhances flexibility, testability, and maintainability of the code."}),"\n",(0,i.jsx)(n.admonition,{type:"tip",children:(0,i.jsx)(n.p,{children:"Dependency Injection is a crucial design pattern in developing complex applications. It is essential for you to master it."})}),"\n",(0,i.jsx)(n.h2,{id:"wire",children:(0,i.jsx)(n.a,{href:"https://github.com/google/wire",children:"Wire"})}),"\n",(0,i.jsxs)(n.blockquote,{children:["\n",(0,i.jsx)(n.p,{children:"Wire is a code generation tool that automates connecting components using dependency injection. Dependencies between components are represented in Wire as function parameters, encouraging explicit initialization instead of global variables. Because Wire operates without runtime state or reflection, code written to be used with Wire is useful even for hand-written initialization."}),"\n"]}),"\n",(0,i.jsxs)(n.p,{children:["Wire is popular and recommended by ",(0,i.jsx)(n.a,{href:"https://github.com/go-kratos/kratos",children:"Kratos"}),", but there are some missing features."]}),"\n",(0,i.jsxs)(n.ul,{children:["\n",(0,i.jsxs)(n.li,{children:["wire.List : inject one type multiple and resolved as list (usefully for seed.Contributor). see ",(0,i.jsx)(n.a,{href:"https://github.com/google/wire/issues/207",children:"https://github.com/google/wire/issues/207"})]}),"\n",(0,i.jsxs)(n.li,{children:["wire.Bind : bind one type as multiple interfaces, (usefully for service proxy). see ",(0,i.jsx)(n.a,{href:"https://github.com/google/wire/issues/191",children:"https://github.com/google/wire/issues/191"})]}),"\n",(0,i.jsxs)(n.li,{children:["wire.Subtract: replace wire.Set type by another. see ",(0,i.jsx)(n.a,{href:"https://github.com/google/wire/issues/8",children:"https://github.com/google/wire/issues/8"})]}),"\n"]}),"\n",(0,i.jsx)(n.p,{children:"So we choose the next one:"}),"\n",(0,i.jsx)(n.h2,{id:"defvaldi",children:(0,i.jsx)(n.a,{href:"https://github.com/defval/di",children:"defval/di"})})]})}function u(e={}){const{wrapper:n}={...(0,s.a)(),...e.components};return n?(0,i.jsx)(n,{...e,children:(0,i.jsx)(d,{...e})}):d(e)}},4744:(e,n,t)=>{t.d(n,{Z:()=>r,a:()=>a});var i=t(959);const s={},o=i.createContext(s);function a(e){const n=i.useContext(o);return i.useMemo((function(){return"function"==typeof e?e(n):{...n,...e}}),[n,e])}function r(e){let n;return n=e.disableParentContext?"function"==typeof e.components?e.components(s):e.components||s:a(e.components),i.createElement(o.Provider,{value:n},e.children)}}}]);