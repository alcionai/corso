"use strict";(self.webpackChunkdocs=self.webpackChunkdocs||[]).push([[671],{3905:(e,t,r)=>{r.d(t,{Zo:()=>u,kt:()=>f});var n=r(7294);function o(e,t,r){return t in e?Object.defineProperty(e,t,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[t]=r,e}function i(e,t){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),r.push.apply(r,n)}return r}function a(e){for(var t=1;t<arguments.length;t++){var r=null!=arguments[t]?arguments[t]:{};t%2?i(Object(r),!0).forEach((function(t){o(e,t,r[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):i(Object(r)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(r,t))}))}return e}function c(e,t){if(null==e)return{};var r,n,o=function(e,t){if(null==e)return{};var r,n,o={},i=Object.keys(e);for(n=0;n<i.length;n++)r=i[n],t.indexOf(r)>=0||(o[r]=e[r]);return o}(e,t);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);for(n=0;n<i.length;n++)r=i[n],t.indexOf(r)>=0||Object.prototype.propertyIsEnumerable.call(e,r)&&(o[r]=e[r])}return o}var s=n.createContext({}),l=function(e){var t=n.useContext(s),r=t;return e&&(r="function"==typeof e?e(t):a(a({},t),e)),r},u=function(e){var t=l(e.components);return n.createElement(s.Provider,{value:t},e.children)},p={inlineCode:"code",wrapper:function(e){var t=e.children;return n.createElement(n.Fragment,{},t)}},d=n.forwardRef((function(e,t){var r=e.components,o=e.mdxType,i=e.originalType,s=e.parentName,u=c(e,["components","mdxType","originalType","parentName"]),d=l(r),f=o,m=d["".concat(s,".").concat(f)]||d[f]||p[f]||i;return r?n.createElement(m,a(a({ref:t},u),{},{components:r})):n.createElement(m,a({ref:t},u))}));function f(e,t){var r=arguments,o=t&&t.mdxType;if("string"==typeof e||o){var i=r.length,a=new Array(i);a[0]=d;var c={};for(var s in t)hasOwnProperty.call(t,s)&&(c[s]=t[s]);c.originalType=e,c.mdxType="string"==typeof e?e:o,a[1]=c;for(var l=2;l<i;l++)a[l]=r[l];return n.createElement.apply(null,a)}return n.createElement.apply(null,r)}d.displayName="MDXCreateElement"},9881:(e,t,r)=>{r.r(t),r.d(t,{assets:()=>s,contentTitle:()=>a,default:()=>p,frontMatter:()=>i,metadata:()=>c,toc:()=>l});var n=r(7462),o=(r(7294),r(3905));r(8209);const i={},a="Introduction",c={unversionedId:"intro",id:"intro",title:"Introduction",description:"Overview",source:"@site/docs/intro.md",sourceDirName:".",slug:"/intro",permalink:"/docs/intro",draft:!1,editUrl:"https://github.com/alcionai/corso/tree/main/docs/docs/intro.md",tags:[],version:"current",frontMatter:{},sidebar:"docsSidebar",next:{title:"Quick start",permalink:"/docs/quickstart"}},s={},l=[{value:"Overview",id:"overview",level:2},{value:"Getting started",id:"getting-started",level:2}],u={toc:l};function p(e){let{components:t,...r}=e;return(0,o.kt)("wrapper",(0,n.Z)({},u,r,{components:t,mdxType:"MDXLayout"}),(0,o.kt)("h1",{id:"introduction"},"Introduction"),(0,o.kt)("h2",{id:"overview"},"Overview"),(0,o.kt)("p",null,"Corso is the first open-source tool that aims to assist IT admins with the critical task of protecting their\nMicrosoft 365 data. It provides a reliable, secure, and efficient data protection engine. Admins decide where to store\nthe backup data and have the flexibility to perform backups of their desired service through an intuitive interface.\nAs Corso evolves, it can become a great building block for more complex data protection workflows."),(0,o.kt)("p",null,"Corso supports M365 Exchange, OneDrive, SharePoint, and Teams. Coverage for more services, possibly\nbeyond M365, will expand based on the interest and needs of the community."),(0,o.kt)("h2",{id:"getting-started"},"Getting started"),(0,o.kt)("p",null,"You can follow the ",(0,o.kt)("a",{parentName:"p",href:"quickstart"},"Quick Start")," guide for an end-to-end Corso walk through. Alternatively, follow\nthe instructions in the ",(0,o.kt)("a",{parentName:"p",href:"setup/concepts"},"Corso Setup")," section to dive into the details on how to configure and\nrun Corso."))}p.isMDXComponent=!0},8209:(e,t,r)=>{r(7294)}}]);