(self.webpackChunkdocs=self.webpackChunkdocs||[]).push([[3085],{8084:(e,n,s)=>{"use strict";s.r(n),s.d(n,{default:()=>j});var t=s(59496),a=s(45924),l=s(27567),r=s(89693),c=s(30750),i=s(70458),o=s(90217);const m={mdxPageWrapper:"mdxPageWrapper_yBV7"};function j(e){const{content:n}=e,{metadata:{title:s,description:j,frontMatter:d}}=n,{wrapperClassName:u,hide_table_of_contents:f}=d;return t.createElement(l.FG,{className:(0,a.Z)(u??r.k.wrapper.mdxPages,r.k.page.mdxPage)},t.createElement(l.d,{title:s,description:j}),t.createElement(c.Z,null,t.createElement("main",{className:"container container--fluid margin-vert--lg"},t.createElement("div",{className:(0,a.Z)("row",m.mdxPageWrapper)},t.createElement("div",{className:(0,a.Z)("col",!f&&"col--8")},t.createElement("article",null,t.createElement(i.Z,null,t.createElement(n,null)))),!f&&n.toc.length>0&&t.createElement("div",{className:"col col--2"},t.createElement(o.Z,{toc:n.toc,minHeadingLevel:d.toc_min_heading_level,maxHeadingLevel:d.toc_max_heading_level}))))))}},90217:(e,n,s)=>{"use strict";s.d(n,{Z:()=>m});var t=s(64778),a=s(59496),l=s(45924),r=s(96086);const c={tableOfContents:"tableOfContents_a4zv",docItemContainer:"docItemContainer_Nnrp"},i="table-of-contents__link toc-highlight",o="table-of-contents__link--active";function m(e){let{className:n,...s}=e;return a.createElement("div",{className:(0,l.Z)(c.tableOfContents,"thin-scrollbar",n)},a.createElement(r.Z,(0,t.Z)({},s,{linkClassName:i,linkActiveClassName:o})))}},96086:(e,n,s)=>{"use strict";s.d(n,{Z:()=>f});var t=s(64778),a=s(59496),l=s(91848);function r(e){const n=e.map((e=>({...e,parentIndex:-1,children:[]}))),s=Array(7).fill(-1);n.forEach(((e,n)=>{const t=s.slice(2,e.level);e.parentIndex=Math.max(...t),s[e.level]=n}));const t=[];return n.forEach((e=>{const{parentIndex:s,...a}=e;s>=0?n[s].children.push(a):t.push(a)})),t}function c(e){let{toc:n,minHeadingLevel:s,maxHeadingLevel:t}=e;return n.flatMap((e=>{const n=c({toc:e.children,minHeadingLevel:s,maxHeadingLevel:t});return function(e){return e.level>=s&&e.level<=t}(e)?[{...e,children:n}]:n}))}function i(e){const n=e.getBoundingClientRect();return n.top===n.bottom?i(e.parentNode):n}function o(e,n){let{anchorTopOffset:s}=n;const t=e.find((e=>i(e).top>=s));if(t){return function(e){return e.top>0&&e.bottom<window.innerHeight/2}(i(t))?t:e[e.indexOf(t)-1]??null}return e[e.length-1]??null}function m(){const e=(0,a.useRef)(0),{navbar:{hideOnScroll:n}}=(0,l.L)();return(0,a.useEffect)((()=>{e.current=n?0:document.querySelector(".navbar").clientHeight}),[n]),e}function j(e){const n=(0,a.useRef)(void 0),s=m();(0,a.useEffect)((()=>{if(!e)return()=>{};const{linkClassName:t,linkActiveClassName:a,minHeadingLevel:l,maxHeadingLevel:r}=e;function c(){const e=function(e){return Array.from(document.getElementsByClassName(e))}(t),c=function(e){let{minHeadingLevel:n,maxHeadingLevel:s}=e;const t=[];for(let a=n;a<=s;a+=1)t.push(`h${a}.anchor`);return Array.from(document.querySelectorAll(t.join()))}({minHeadingLevel:l,maxHeadingLevel:r}),i=o(c,{anchorTopOffset:s.current}),m=e.find((e=>i&&i.id===function(e){return decodeURIComponent(e.href.substring(e.href.indexOf("#")+1))}(e)));e.forEach((e=>{!function(e,s){s?(n.current&&n.current!==e&&n.current.classList.remove(a),e.classList.add(a),n.current=e):e.classList.remove(a)}(e,e===m)}))}return document.addEventListener("scroll",c),document.addEventListener("resize",c),c(),()=>{document.removeEventListener("scroll",c),document.removeEventListener("resize",c)}}),[e,s])}function d(e){let{toc:n,className:s,linkClassName:t,isChild:l}=e;return n.length?a.createElement("ul",{className:l?void 0:s},n.map((e=>a.createElement("li",{key:e.id},a.createElement("a",{href:`#${e.id}`,className:t??void 0,dangerouslySetInnerHTML:{__html:e.value}}),a.createElement(d,{isChild:!0,toc:e.children,className:s,linkClassName:t}))))):null}const u=a.memo(d);function f(e){let{toc:n,className:s="table-of-contents table-of-contents__left-border",linkClassName:i="table-of-contents__link",linkActiveClassName:o,minHeadingLevel:m,maxHeadingLevel:d,...f}=e;const v=(0,l.L)(),h=m??v.tableOfContents.minHeadingLevel,g=d??v.tableOfContents.maxHeadingLevel,k=function(e){let{toc:n,minHeadingLevel:s,maxHeadingLevel:t}=e;return(0,a.useMemo)((()=>c({toc:r(n),minHeadingLevel:s,maxHeadingLevel:t})),[n,s,t])}({toc:n,minHeadingLevel:h,maxHeadingLevel:g});return j((0,a.useMemo)((()=>{if(i&&o)return{linkClassName:i,linkActiveClassName:o,minHeadingLevel:h,maxHeadingLevel:g}}),[i,o,h,g])),a.createElement(u,(0,t.Z)({toc:k,className:s,linkClassName:i},f))}},38660:(e,n,s)=>{var t={"./af":28163,"./af.js":28163,"./ar":70096,"./ar-dz":50005,"./ar-dz.js":50005,"./ar-kw":16042,"./ar-kw.js":16042,"./ar-ly":43183,"./ar-ly.js":43183,"./ar-ma":35656,"./ar-ma.js":35656,"./ar-sa":32343,"./ar-sa.js":32343,"./ar-tn":23815,"./ar-tn.js":23815,"./ar.js":70096,"./az":62320,"./az.js":62320,"./be":32204,"./be.js":32204,"./bg":94994,"./bg.js":94994,"./bm":93556,"./bm.js":93556,"./bn":85853,"./bn-bd":48735,"./bn-bd.js":48735,"./bn.js":85853,"./bo":44547,"./bo.js":44547,"./br":29491,"./br.js":29491,"./bs":33046,"./bs.js":33046,"./ca":20568,"./ca.js":20568,"./cs":46889,"./cs.js":46889,"./cv":7267,"./cv.js":7267,"./cy":24362,"./cy.js":24362,"./da":34133,"./da.js":34133,"./de":44268,"./de-at":28288,"./de-at.js":28288,"./de-ch":42916,"./de-ch.js":42916,"./de.js":44268,"./dv":39938,"./dv.js":39938,"./el":51336,"./el.js":51336,"./en-au":62902,"./en-au.js":62902,"./en-ca":70270,"./en-ca.js":70270,"./en-gb":48554,"./en-gb.js":48554,"./en-ie":95733,"./en-ie.js":95733,"./en-il":71910,"./en-il.js":71910,"./en-in":1697,"./en-in.js":1697,"./en-nz":48158,"./en-nz.js":48158,"./en-sg":88683,"./en-sg.js":88683,"./eo":61190,"./eo.js":61190,"./es":19327,"./es-do":46495,"./es-do.js":46495,"./es-mx":97130,"./es-mx.js":97130,"./es-us":63738,"./es-us.js":63738,"./es.js":19327,"./et":19477,"./et.js":19477,"./eu":1817,"./eu.js":1817,"./fa":77570,"./fa.js":77570,"./fi":51753,"./fi.js":51753,"./fil":39975,"./fil.js":39975,"./fo":10756,"./fo.js":10756,"./fr":79876,"./fr-ca":32951,"./fr-ca.js":32951,"./fr-ch":43803,"./fr-ch.js":43803,"./fr.js":79876,"./fy":53470,"./fy.js":53470,"./ga":31347,"./ga.js":31347,"./gd":72624,"./gd.js":72624,"./gl":93758,"./gl.js":93758,"./gom-deva":22327,"./gom-deva.js":22327,"./gom-latn":97089,"./gom-latn.js":97089,"./gu":60929,"./gu.js":60929,"./he":4438,"./he.js":4438,"./hi":57746,"./hi.js":57746,"./hr":85893,"./hr.js":85893,"./hu":55169,"./hu.js":55169,"./hy-am":60394,"./hy-am.js":60394,"./id":34171,"./id.js":34171,"./is":28435,"./is.js":28435,"./it":22412,"./it-ch":33640,"./it-ch.js":33640,"./it.js":22412,"./ja":76261,"./ja.js":76261,"./jv":2826,"./jv.js":2826,"./ka":9887,"./ka.js":9887,"./kk":16262,"./kk.js":16262,"./km":95730,"./km.js":95730,"./kn":55773,"./kn.js":55773,"./ko":91206,"./ko.js":91206,"./ku":20571,"./ku.js":20571,"./ky":67971,"./ky.js":67971,"./lb":75817,"./lb.js":75817,"./lo":93863,"./lo.js":93863,"./lt":20610,"./lt.js":20610,"./lv":55147,"./lv.js":55147,"./me":53827,"./me.js":53827,"./mi":90358,"./mi.js":90358,"./mk":88489,"./mk.js":88489,"./ml":72930,"./ml.js":72930,"./mn":97248,"./mn.js":97248,"./mr":98619,"./mr.js":98619,"./ms":98583,"./ms-my":30810,"./ms-my.js":30810,"./ms.js":98583,"./mt":53425,"./mt.js":53425,"./my":72319,"./my.js":72319,"./nb":37866,"./nb.js":37866,"./ne":57630,"./ne.js":57630,"./nl":96940,"./nl-be":34037,"./nl-be.js":34037,"./nl.js":96940,"./nn":34411,"./nn.js":34411,"./oc-lnc":21584,"./oc-lnc.js":21584,"./pa-in":24633,"./pa-in.js":24633,"./pl":54390,"./pl.js":54390,"./pt":95040,"./pt-br":54013,"./pt-br.js":54013,"./pt.js":95040,"./ro":12338,"./ro.js":12338,"./ru":9874,"./ru.js":9874,"./sd":34992,"./sd.js":34992,"./se":28076,"./se.js":28076,"./si":22737,"./si.js":22737,"./sk":83699,"./sk.js":83699,"./sl":44565,"./sl.js":44565,"./sq":35255,"./sq.js":35255,"./sr":95098,"./sr-cyrl":3830,"./sr-cyrl.js":3830,"./sr.js":95098,"./ss":19587,"./ss.js":19587,"./sv":34285,"./sv.js":34285,"./sw":13065,"./sw.js":13065,"./ta":52703,"./ta.js":52703,"./te":90618,"./te.js":90618,"./tet":19035,"./tet.js":19035,"./tg":45334,"./tg.js":45334,"./th":41853,"./th.js":41853,"./tk":37250,"./tk.js":37250,"./tl-ph":50704,"./tl-ph.js":50704,"./tlh":1866,"./tlh.js":1866,"./tr":81299,"./tr.js":81299,"./tzl":35885,"./tzl.js":35885,"./tzm":21370,"./tzm-latn":11878,"./tzm-latn.js":11878,"./tzm.js":21370,"./ug-cn":95521,"./ug-cn.js":95521,"./uk":63646,"./uk.js":63646,"./ur":16237,"./ur.js":16237,"./uz":90519,"./uz-latn":23225,"./uz-latn.js":23225,"./uz.js":90519,"./vi":31153,"./vi.js":31153,"./x-pseudo":77210,"./x-pseudo.js":77210,"./yo":13086,"./yo.js":13086,"./zh-cn":53933,"./zh-cn.js":53933,"./zh-hk":45749,"./zh-hk.js":45749,"./zh-mo":51044,"./zh-mo.js":51044,"./zh-tw":13759,"./zh-tw.js":13759};function a(e){var n=l(e);return s(n)}function l(e){if(!s.o(t,e)){var n=new Error("Cannot find module '"+e+"'");throw n.code="MODULE_NOT_FOUND",n}return t[e]}a.keys=function(){return Object.keys(t)},a.resolve=l,e.exports=a,a.id=38660}}]);