import{b as T,G as he}from"./index.387a2346.js";import{q as R}from"./base.87fcf6e2.js";import{aN as fe,x as ye}from"./arco.17b1a46f.js";function be(p){return T.post("/api/v1/log/chat/page",p)}function Ce(p){return T.get("/api/v1/log/chat/detail",{params:p,paramsSerializer:B=>R.stringify(B)})}function ke(p){return T.post("/api/v1/log/image/page",p)}function we(p){return T.get("/api/v1/log/image/detail",{params:p,paramsSerializer:B=>R.stringify(B)})}function Ne(p){return T.post("/api/v1/log/audio/page",p)}function Se(p){return T.get("/api/v1/log/audio/detail",{params:p,paramsSerializer:B=>R.stringify(B)})}function je(p){return T.post("/api/v1/log/chat/export",p,{responseType:"blob"})}function Oe(p){return T.post("/api/v1/log/chat/batch/operate",p)}var le={exports:{}};(function(p,B){(function(_,M){p.exports=M(he)})(ye,function(_){return function(){var M={789:function(h){h.exports=_}},G={};function S(h){var k=G[h];if(k!==void 0)return k.exports;var v=G[h]={exports:{}};return M[h](v,v.exports,S),v.exports}S.d=function(h,k){for(var v in k)S.o(k,v)&&!S.o(h,v)&&Object.defineProperty(h,v,{enumerable:!0,get:k[v]})},S.o=function(h,k){return Object.prototype.hasOwnProperty.call(h,k)},S.r=function(h){typeof Symbol<"u"&&Symbol.toStringTag&&Object.defineProperty(h,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(h,"__esModule",{value:!0})};var z={};return function(){function h(e,t){(t==null||t>e.length)&&(t=e.length);for(var o=0,l=new Array(t);o<t;o++)l[o]=e[o];return l}function k(e,t){if(e){if(typeof e=="string")return h(e,t);var o=Object.prototype.toString.call(e).slice(8,-1);return o==="Object"&&e.constructor&&(o=e.constructor.name),o==="Map"||o==="Set"?Array.from(e):o==="Arguments"||/^(?:Ui|I)nt(?:8|16|32)(?:Clamped)?Array$/.test(o)?h(e,t):void 0}}function v(e){return function(t){if(Array.isArray(t))return h(t)}(e)||function(t){if(typeof Symbol<"u"&&t[Symbol.iterator]!=null||t["@@iterator"]!=null)return Array.from(t)}(e)||k(e)||function(){throw new TypeError(`Invalid attempt to spread non-iterable instance.
In order to be iterable, non-array objects must have a [Symbol.iterator]() method.`)}()}function K(e,t,o){return t in e?Object.defineProperty(e,t,{value:o,enumerable:!0,configurable:!0,writable:!0}):e[t]=o,e}S.r(z),S.d(z,{default:function(){return de}});var r=S(789),ie=(0,r.defineComponent)({props:{data:{required:!0,type:String},onClick:Function},render:function(){var e=this.data,t=this.onClick;return(0,r.createVNode)("span",{class:"vjs-tree-brackets",onClick:t},[e])}}),ce=(0,r.defineComponent)({emits:["change","update:modelValue"],props:{checked:{type:Boolean,default:!1},isMultiple:Boolean,onChange:Function},setup:function(e,t){var o=t.emit;return{uiType:(0,r.computed)(function(){return e.isMultiple?"checkbox":"radio"}),model:(0,r.computed)({get:function(){return e.checked},set:function(l){return o("update:modelValue",l)}})}},render:function(){var e=this.uiType,t=this.model,o=this.$emit;return(0,r.createVNode)("label",{class:["vjs-check-controller",t?"is-checked":""],onClick:function(l){return l.stopPropagation()}},[(0,r.createVNode)("span",{class:"vjs-check-controller-inner is-".concat(e)},null),(0,r.createVNode)("input",{checked:t,class:"vjs-check-controller-original is-".concat(e),type:e,onChange:function(){return o("change",t)}},null)])}}),ue=(0,r.defineComponent)({props:{nodeType:{required:!0,type:String},onClick:Function},render:function(){var e=this.nodeType,t=this.onClick,o=e==="objectStart"||e==="arrayStart";return o||e==="objectCollapsed"||e==="arrayCollapsed"?(0,r.createVNode)("span",{class:"vjs-carets vjs-carets-".concat(o?"open":"close"),onClick:t},[(0,r.createVNode)("svg",{viewBox:"0 0 1024 1024",focusable:"false","data-icon":"caret-down",width:"1em",height:"1em",fill:"currentColor","aria-hidden":"true"},[(0,r.createVNode)("path",{d:"M840.4 300H183.6c-19.7 0-30.7 20.8-18.5 35l328.4 380.8c9.4 10.9 27.5 10.9 37 0L858.9 335c12.2-14.2 1.2-35-18.5-35z"},null)])]):null}});function $(e){return $=typeof Symbol=="function"&&typeof Symbol.iterator=="symbol"?function(t){return typeof t}:function(t){return t&&typeof Symbol=="function"&&t.constructor===Symbol&&t!==Symbol.prototype?"symbol":typeof t},$(e)}function W(e){return Object.prototype.toString.call(e).slice(8,-1).toLowerCase()}function D(e){var t=arguments.length>1&&arguments[1]!==void 0?arguments[1]:"root",o=arguments.length>2&&arguments[2]!==void 0?arguments[2]:0,l=arguments.length>3?arguments[3]:void 0,f=l||{},b=f.key,j=f.index,d=f.type,N=d===void 0?"content":d,g=f.showComma,P=g!==void 0&&g,O=f.length,A=O===void 0?1:O,L=W(e);if(L==="array"){var F=U(e.map(function(C,u,n){return D(C,"".concat(t,"[").concat(u,"]"),o+1,{index:u,showComma:u!==n.length-1,length:A,type:N})}));return[D("[",t,o,{showComma:!1,key:b,length:e.length,type:"arrayStart"})[0]].concat(F,D("]",t,o,{showComma:P,length:e.length,type:"arrayEnd"})[0])}if(L==="object"){var x=Object.keys(e),I=U(x.map(function(C,u,n){return D(e[C],/^[a-zA-Z_]\w*$/.test(C)?"".concat(t,".").concat(C):"".concat(t,'["').concat(C,'"]'),o+1,{key:C,showComma:u!==n.length-1,length:A,type:N})}));return[D("{",t,o,{showComma:!1,key:b,index:j,length:x.length,type:"objectStart"})[0]].concat(I,D("}",t,o,{showComma:P,length:x.length,type:"objectEnd"})[0])}return[{content:e,level:o,key:b,index:j,path:t,showComma:P,length:A,type:N}]}function U(e){if(typeof Array.prototype.flat=="function")return e.flat();for(var t=v(e),o=[];t.length;){var l=t.shift();Array.isArray(l)?t.unshift.apply(t,v(l)):o.push(l)}return o}function Q(e){var t=arguments.length>1&&arguments[1]!==void 0?arguments[1]:new WeakMap;if(e==null)return e;if(e instanceof Date)return new Date(e);if(e instanceof RegExp)return new RegExp(e);if($(e)!=="object")return e;if(t.get(e))return t.get(e);if(Array.isArray(e)){var o=e.map(function(b){return Q(b,t)});return t.set(e,o),o}var l={};for(var f in e)l[f]=Q(e[f],t);return t.set(e,l),l}function Z(e,t){var o=Object.keys(e);if(Object.getOwnPropertySymbols){var l=Object.getOwnPropertySymbols(e);t&&(l=l.filter(function(f){return Object.getOwnPropertyDescriptor(e,f).enumerable})),o.push.apply(o,l)}return o}function X(e){for(var t=1;t<arguments.length;t++){var o=arguments[t]!=null?arguments[t]:{};t%2?Z(Object(o),!0).forEach(function(l){K(e,l,o[l])}):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(o)):Z(Object(o)).forEach(function(l){Object.defineProperty(e,l,Object.getOwnPropertyDescriptor(o,l))})}return e}var ee={showLength:{type:Boolean,default:!1},showDoubleQuotes:{type:Boolean,default:!0},renderNodeKey:Function,renderNodeValue:Function,selectableType:String,showSelectController:{type:Boolean,default:!1},showLine:{type:Boolean,default:!0},showLineNumber:{type:Boolean,default:!1},selectOnClickNode:{type:Boolean,default:!0},nodeSelectable:{type:Function,default:function(){return!0}},highlightSelectedNode:{type:Boolean,default:!0},showIcon:{type:Boolean,default:!1},showKeyValueSpace:{type:Boolean,default:!0},editable:{type:Boolean,default:!1},editableTrigger:{type:String,default:"click"},onNodeClick:{type:Function},onBracketsClick:{type:Function},onIconClick:{type:Function},onValueChange:{type:Function}},se=(0,r.defineComponent)({name:"TreeNode",props:X(X({},ee),{},{node:{type:Object,required:!0},collapsed:Boolean,checked:Boolean,style:Object,onSelectedChange:{type:Function}}),emits:["nodeClick","bracketsClick","iconClick","selectedChange","valueChange"],setup:function(e,t){var o=t.emit,l=(0,r.computed)(function(){return W(e.node.content)}),f=(0,r.computed)(function(){return"vjs-value vjs-value-".concat(l.value)}),b=(0,r.computed)(function(){return e.showDoubleQuotes?'"'.concat(e.node.key,'"'):e.node.key}),j=(0,r.computed)(function(){return e.selectableType==="multiple"}),d=(0,r.computed)(function(){return e.selectableType==="single"}),N=(0,r.computed)(function(){return e.nodeSelectable(e.node)&&(j.value||d.value)}),g=(0,r.reactive)({editing:!1}),P=function(u){var n,a,i=(a=(n=u.target)===null||n===void 0?void 0:n.value)==="null"?null:a==="undefined"?void 0:a==="true"||a!=="false"&&(a[0]+a[a.length-1]==='""'||a[0]+a[a.length-1]==="''"?a.slice(1,-1):typeof Number(a)=="number"&&!isNaN(Number(a))||a==="NaN"?Number(a):a);o("valueChange",i,e.node.path)},O=(0,r.computed)(function(){var u,n=(u=e.node)===null||u===void 0?void 0:u.content;return n===null?n="null":n===void 0&&(n="undefined"),l.value==="string"?'"'.concat(n,'"'):n+""}),A=function(){var u=e.renderNodeValue;return u?u({node:e.node,defaultValue:O.value}):O.value},L=function(){o("bracketsClick",!e.collapsed,e.node.path)},F=function(){o("iconClick",!e.collapsed,e.node.path)},x=function(){o("selectedChange",e.node)},I=function(){o("nodeClick",e.node),N.value&&e.selectOnClickNode&&o("selectedChange",e.node)},C=function(u){if(e.editable&&!g.editing){g.editing=!0;var n=function a(i){var c;i.target!==u.target&&((c=i.target)===null||c===void 0?void 0:c.parentElement)!==u.target&&(g.editing=!1,document.removeEventListener("click",a))};document.removeEventListener("click",n),document.addEventListener("click",n)}};return function(){var u,n=e.node;return(0,r.createVNode)("div",{class:{"vjs-tree-node":!0,"has-selector":e.showSelectController,"has-carets":e.showIcon,"is-highlight":e.highlightSelectedNode&&e.checked},onClick:I,style:e.style},[e.showLineNumber&&(0,r.createVNode)("span",{class:"vjs-node-index"},[n.id+1]),e.showSelectController&&N.value&&n.type!=="objectEnd"&&n.type!=="arrayEnd"&&(0,r.createVNode)(ce,{isMultiple:j.value,checked:e.checked,onChange:x},null),(0,r.createVNode)("div",{class:"vjs-indent"},[Array.from(Array(n.level)).map(function(a,i){return(0,r.createVNode)("div",{key:i,class:{"vjs-indent-unit":!0,"has-line":e.showLine}},null)}),e.showIcon&&(0,r.createVNode)(ue,{nodeType:n.type,onClick:F},null)]),n.key&&(0,r.createVNode)("span",{class:"vjs-key"},[(u=e.renderNodeKey,u?u({node:e.node,defaultKey:b.value||""}):b.value),(0,r.createVNode)("span",{class:"vjs-colon"},[":".concat(e.showKeyValueSpace?" ":"")])]),(0,r.createVNode)("span",null,[n.type!=="content"&&n.content?(0,r.createVNode)(ie,{data:n.content.toString(),onClick:L},null):(0,r.createVNode)("span",{class:f.value,onClick:!e.editable||e.editableTrigger&&e.editableTrigger!=="click"?void 0:C,onDblclick:e.editable&&e.editableTrigger==="dblclick"?C:void 0},[e.editable&&g.editing?(0,r.createVNode)("input",{value:O.value,onChange:P,style:{padding:"3px 8px",border:"1px solid #eee",boxShadow:"none",boxSizing:"border-box",borderRadius:5,fontFamily:"inherit"}},null):A()]),n.showComma&&(0,r.createVNode)("span",null,[","]),e.showLength&&e.collapsed&&(0,r.createVNode)("span",{class:"vjs-comment"},[(0,r.createTextVNode)(" // "),n.length,(0,r.createTextVNode)(" items ")])])])}}});function te(e,t){var o=Object.keys(e);if(Object.getOwnPropertySymbols){var l=Object.getOwnPropertySymbols(e);t&&(l=l.filter(function(f){return Object.getOwnPropertyDescriptor(e,f).enumerable})),o.push.apply(o,l)}return o}function m(e){for(var t=1;t<arguments.length;t++){var o=arguments[t]!=null?arguments[t]:{};t%2?te(Object(o),!0).forEach(function(l){K(e,l,o[l])}):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(o)):te(Object(o)).forEach(function(l){Object.defineProperty(e,l,Object.getOwnPropertyDescriptor(o,l))})}return e}var de=(0,r.defineComponent)({name:"Tree",props:m(m({},ee),{},{data:{type:[String,Number,Boolean,Array,Object],default:null},collapsedNodeLength:{type:Number,default:1/0},deep:{type:Number,default:1/0},pathCollapsible:{type:Function,default:function(){return!1}},rootPath:{type:String,default:"root"},virtual:{type:Boolean,default:!1},height:{type:Number,default:400},itemHeight:{type:Number,default:20},selectedValue:{type:[String,Array],default:function(){return""}},collapsedOnClickBrackets:{type:Boolean,default:!0},style:Object,onSelectedChange:{type:Function}}),slots:["renderNodeKey","renderNodeValue"],emits:["nodeClick","bracketsClick","iconClick","selectedChange","update:selectedValue","update:data"],setup:function(e,t){var o=t.emit,l=t.slots,f=(0,r.ref)(),b=(0,r.computed)(function(){return D(e.data,e.rootPath)}),j=function(n,a){return b.value.reduce(function(i,c){var s,y=c.level>=n||c.length>=a,w=(s=e.pathCollapsible)===null||s===void 0?void 0:s.call(e,c);return c.type!=="objectStart"&&c.type!=="arrayStart"||!y&&!w?i:m(m({},i),{},K({},c.path,1))},{})},d=(0,r.reactive)({translateY:0,visibleData:null,hiddenPaths:j(e.deep,e.collapsedNodeLength)}),N=(0,r.computed)(function(){for(var n=null,a=[],i=b.value.length,c=0;c<i;c++){var s=m(m({},b.value[c]),{},{id:c}),y=d.hiddenPaths[s.path];if(n&&n.path===s.path){var w=n.type==="objectStart",H=m(m(m({},s),n),{},{showComma:s.showComma,content:w?"{...}":"[...]",type:w?"objectCollapsed":"arrayCollapsed"});n=null,a.push(H)}else{if(y&&!n){n=s;continue}if(n)continue;a.push(s)}}return a}),g=(0,r.computed)(function(){var n=e.selectedValue;return n&&e.selectableType==="multiple"&&Array.isArray(n)?n:[n]}),P=(0,r.computed)(function(){return!e.selectableType||e.selectOnClickNode||e.showSelectController?"":"When selectableType is not null, selectOnClickNode and showSelectController cannot be false at the same time, because this will cause the selection to fail."}),O=function(){var n=N.value;if(e.virtual){var a,i=e.height/e.itemHeight,c=((a=f.value)===null||a===void 0?void 0:a.scrollTop)||0,s=Math.floor(c/e.itemHeight),y=s<0?0:s+i>n.length?n.length-i:s;y<0&&(y=0);var w=y+i;d.translateY=y*e.itemHeight,d.visibleData=n.filter(function(H,q){return q>=y&&q<w})}else d.visibleData=n},A=function(){O()},L=function(n){var a,i,c=n.path,s=e.selectableType;if(s==="multiple"){var y=g.value.findIndex(function(V){return V===c}),w=v(g.value);y!==-1?w.splice(y,1):w.push(c),o("update:selectedValue",w),o("selectedChange",w,v(g.value))}else if(s==="single"&&g.value[0]!==c){var H=(a=g.value,i=1,function(V){if(Array.isArray(V))return V}(a)||function(V,ne){var E=V==null?null:typeof Symbol<"u"&&V[Symbol.iterator]||V["@@iterator"];if(E!=null){var oe,re,Y=[],J=!0,ae=!1;try{for(E=E.call(V);!(J=(oe=E.next()).done)&&(Y.push(oe.value),!ne||Y.length!==ne);J=!0);}catch(pe){ae=!0,re=pe}finally{try{J||E.return==null||E.return()}finally{if(ae)throw re}}return Y}}(a,i)||k(a,i)||function(){throw new TypeError(`Invalid attempt to destructure non-iterable instance.
In order to be iterable, non-array objects must have a [Symbol.iterator]() method.`)}())[0],q=c;o("update:selectedValue",q),o("selectedChange",q,H)}},F=function(n){o("nodeClick",n)},x=function(n,a){if(n)d.hiddenPaths=m(m({},d.hiddenPaths),{},K({},a,1));else{var i=m({},d.hiddenPaths);delete i[a],d.hiddenPaths=i}},I=function(n,a){e.collapsedOnClickBrackets&&x(n,a),o("bracketsClick",n)},C=function(n,a){x(n,a),o("iconClick",n)},u=function(n,a){var i=Q(e.data),c=e.rootPath;new Function("data","val","data".concat(a.slice(c.length),"=val"))(i,n),o("update:data",i)};return(0,r.watchEffect)(function(){P.value&&function(n){throw new Error("[VueJSONPretty] ".concat(n))}(P.value)}),(0,r.watchEffect)(function(){N.value&&O()}),(0,r.watch)(function(){return e.deep},function(n){n&&(d.hiddenPaths=j(n,e.collapsedNodeLength))}),(0,r.watch)(function(){return e.collapsedNodeLength},function(n){n&&(d.hiddenPaths=j(e.deep,n))}),function(){var n,a,i=(n=e.renderNodeKey)!==null&&n!==void 0?n:l.renderNodeKey,c=(a=e.renderNodeValue)!==null&&a!==void 0?a:l.renderNodeValue,s=d.visibleData&&d.visibleData.map(function(y){return(0,r.createVNode)(se,{key:y.id,node:y,collapsed:!!d.hiddenPaths[y.path],showDoubleQuotes:e.showDoubleQuotes,showLength:e.showLength,checked:g.value.includes(y.path),selectableType:e.selectableType,showLine:e.showLine,showLineNumber:e.showLineNumber,showSelectController:e.showSelectController,selectOnClickNode:e.selectOnClickNode,nodeSelectable:e.nodeSelectable,highlightSelectedNode:e.highlightSelectedNode,editable:e.editable,editableTrigger:e.editableTrigger,showIcon:e.showIcon,showKeyValueSpace:e.showKeyValueSpace,renderNodeKey:i,renderNodeValue:c,onNodeClick:F,onBracketsClick:I,onIconClick:C,onSelectedChange:L,onValueChange:u,style:e.itemHeight&&e.itemHeight!==20?{lineHeight:"".concat(e.itemHeight,"px")}:{}},null)});return(0,r.createVNode)("div",{ref:f,class:{"vjs-tree":!0,"is-virtual":e.virtual},onScroll:e.virtual?A:void 0,style:e.showLineNumber?m({paddingLeft:"".concat(12*Number(b.value.length.toString().length),"px")},e.style):e.style},[e.virtual?(0,r.createVNode)("div",{class:"vjs-tree-list",style:{height:"".concat(e.height,"px")}},[(0,r.createVNode)("div",{class:"vjs-tree-list-holder",style:{height:"".concat(N.value.length*e.itemHeight,"px")}},[(0,r.createVNode)("div",{class:"vjs-tree-list-holder-inner",style:{transform:"translateY(".concat(d.translateY,"px)")}},[s])])]):s])}}})}(),z}()})})(le);const Ve=fe(le.exports);export{Ve as V,Ce as a,Oe as b,ke as c,we as d,Ne as e,Se as f,be as q,je as s};
