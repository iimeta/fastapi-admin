import{b as H,G as he}from"./index.b8904415.js";import{q as ae}from"./base.87fcf6e2.js";import{aN as fe,x as ye}from"./arco.17b1a46f.js";function be(N){return H.post("/api/v1/log/chat/page",N)}function Ce(N){return H.get("/api/v1/log/chat/detail",{params:N,paramsSerializer:I=>ae.stringify(I)})}function ke(N){return H.post("/api/v1/log/image/page",N)}function we(N){return H.get("/api/v1/log/image/detail",{params:N,paramsSerializer:I=>ae.stringify(I)})}var le={exports:{}};(function(N,I){(function(R,M){N.exports=M(he)})(ye,function(R){return function(){var M={789:function(p){p.exports=R}},_={};function j(p){var C=_[p];if(C!==void 0)return C.exports;var g=_[p]={exports:{}};return M[p](g,g.exports,j),g.exports}j.d=function(p,C){for(var g in C)j.o(C,g)&&!j.o(p,g)&&Object.defineProperty(p,g,{enumerable:!0,get:C[g]})},j.o=function(p,C){return Object.prototype.hasOwnProperty.call(p,C)},j.r=function(p){typeof Symbol<"u"&&Symbol.toStringTag&&Object.defineProperty(p,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(p,"__esModule",{value:!0})};var z={};return function(){function p(e,t){(t==null||t>e.length)&&(t=e.length);for(var o=0,l=new Array(t);o<t;o++)l[o]=e[o];return l}function C(e,t){if(e){if(typeof e=="string")return p(e,t);var o=Object.prototype.toString.call(e).slice(8,-1);return o==="Object"&&e.constructor&&(o=e.constructor.name),o==="Map"||o==="Set"?Array.from(e):o==="Arguments"||/^(?:Ui|I)nt(?:8|16|32)(?:Clamped)?Array$/.test(o)?p(e,t):void 0}}function g(e){return function(t){if(Array.isArray(t))return p(t)}(e)||function(t){if(typeof Symbol<"u"&&t[Symbol.iterator]!=null||t["@@iterator"]!=null)return Array.from(t)}(e)||C(e)||function(){throw new TypeError(`Invalid attempt to spread non-iterable instance.
In order to be iterable, non-array objects must have a [Symbol.iterator]() method.`)}()}function q(e,t,o){return t in e?Object.defineProperty(e,t,{value:o,enumerable:!0,configurable:!0,writable:!0}):e[t]=o,e}j.r(z),j.d(z,{default:function(){return de}});var r=j(789),ie=(0,r.defineComponent)({props:{data:{required:!0,type:String},onClick:Function},render:function(){var e=this.data,t=this.onClick;return(0,r.createVNode)("span",{class:"vjs-tree-brackets",onClick:t},[e])}}),ce=(0,r.defineComponent)({emits:["change","update:modelValue"],props:{checked:{type:Boolean,default:!1},isMultiple:Boolean,onChange:Function},setup:function(e,t){var o=t.emit;return{uiType:(0,r.computed)(function(){return e.isMultiple?"checkbox":"radio"}),model:(0,r.computed)({get:function(){return e.checked},set:function(l){return o("update:modelValue",l)}})}},render:function(){var e=this.uiType,t=this.model,o=this.$emit;return(0,r.createVNode)("label",{class:["vjs-check-controller",t?"is-checked":""],onClick:function(l){return l.stopPropagation()}},[(0,r.createVNode)("span",{class:"vjs-check-controller-inner is-".concat(e)},null),(0,r.createVNode)("input",{checked:t,class:"vjs-check-controller-original is-".concat(e),type:e,onChange:function(){return o("change",t)}},null)])}}),ue=(0,r.defineComponent)({props:{nodeType:{required:!0,type:String},onClick:Function},render:function(){var e=this.nodeType,t=this.onClick,o=e==="objectStart"||e==="arrayStart";return o||e==="objectCollapsed"||e==="arrayCollapsed"?(0,r.createVNode)("span",{class:"vjs-carets vjs-carets-".concat(o?"open":"close"),onClick:t},[(0,r.createVNode)("svg",{viewBox:"0 0 1024 1024",focusable:"false","data-icon":"caret-down",width:"1em",height:"1em",fill:"currentColor","aria-hidden":"true"},[(0,r.createVNode)("path",{d:"M840.4 300H183.6c-19.7 0-30.7 20.8-18.5 35l328.4 380.8c9.4 10.9 27.5 10.9 37 0L858.9 335c12.2-14.2 1.2-35-18.5-35z"},null)])]):null}});function $(e){return $=typeof Symbol=="function"&&typeof Symbol.iterator=="symbol"?function(t){return typeof t}:function(t){return t&&typeof Symbol=="function"&&t.constructor===Symbol&&t!==Symbol.prototype?"symbol":typeof t},$(e)}function G(e){return Object.prototype.toString.call(e).slice(8,-1).toLowerCase()}function T(e){var t=arguments.length>1&&arguments[1]!==void 0?arguments[1]:"root",o=arguments.length>2&&arguments[2]!==void 0?arguments[2]:0,l=arguments.length>3?arguments[3]:void 0,h=l||{},m=h.key,S=h.index,d=h.type,w=d===void 0?"content":d,y=h.showComma,P=y!==void 0&&y,O=h.length,D=O===void 0?1:O,L=G(e);if(L==="array"){var A=W(e.map(function(b,u,n){return T(b,"".concat(t,"[").concat(u,"]"),o+1,{index:u,showComma:u!==n.length-1,length:D,type:w})}));return[T("[",t,o,{showComma:!1,key:m,length:e.length,type:"arrayStart"})[0]].concat(A,T("]",t,o,{showComma:P,length:e.length,type:"arrayEnd"})[0])}if(L==="object"){var x=Object.keys(e),E=W(x.map(function(b,u,n){return T(e[b],/^[a-zA-Z_]\w*$/.test(b)?"".concat(t,".").concat(b):"".concat(t,'["').concat(b,'"]'),o+1,{key:b,showComma:u!==n.length-1,length:D,type:w})}));return[T("{",t,o,{showComma:!1,key:m,index:S,length:x.length,type:"objectStart"})[0]].concat(E,T("}",t,o,{showComma:P,length:x.length,type:"objectEnd"})[0])}return[{content:e,level:o,key:m,index:S,path:t,showComma:P,length:D,type:w}]}function W(e){if(typeof Array.prototype.flat=="function")return e.flat();for(var t=g(e),o=[];t.length;){var l=t.shift();Array.isArray(l)?t.unshift.apply(t,g(l)):o.push(l)}return o}function Q(e){var t=arguments.length>1&&arguments[1]!==void 0?arguments[1]:new WeakMap;if(e==null)return e;if(e instanceof Date)return new Date(e);if(e instanceof RegExp)return new RegExp(e);if($(e)!=="object")return e;if(t.get(e))return t.get(e);if(Array.isArray(e)){var o=e.map(function(m){return Q(m,t)});return t.set(e,o),o}var l={};for(var h in e)l[h]=Q(e[h],t);return t.set(e,l),l}function U(e,t){var o=Object.keys(e);if(Object.getOwnPropertySymbols){var l=Object.getOwnPropertySymbols(e);t&&(l=l.filter(function(h){return Object.getOwnPropertyDescriptor(e,h).enumerable})),o.push.apply(o,l)}return o}function Z(e){for(var t=1;t<arguments.length;t++){var o=arguments[t]!=null?arguments[t]:{};t%2?U(Object(o),!0).forEach(function(l){q(e,l,o[l])}):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(o)):U(Object(o)).forEach(function(l){Object.defineProperty(e,l,Object.getOwnPropertyDescriptor(o,l))})}return e}var X={showLength:{type:Boolean,default:!1},showDoubleQuotes:{type:Boolean,default:!0},renderNodeKey:Function,renderNodeValue:Function,selectableType:String,showSelectController:{type:Boolean,default:!1},showLine:{type:Boolean,default:!0},showLineNumber:{type:Boolean,default:!1},selectOnClickNode:{type:Boolean,default:!0},nodeSelectable:{type:Function,default:function(){return!0}},highlightSelectedNode:{type:Boolean,default:!0},showIcon:{type:Boolean,default:!1},showKeyValueSpace:{type:Boolean,default:!0},editable:{type:Boolean,default:!1},editableTrigger:{type:String,default:"click"},onNodeClick:{type:Function},onBracketsClick:{type:Function},onIconClick:{type:Function},onValueChange:{type:Function}},se=(0,r.defineComponent)({name:"TreeNode",props:Z(Z({},X),{},{node:{type:Object,required:!0},collapsed:Boolean,checked:Boolean,style:Object,onSelectedChange:{type:Function}}),emits:["nodeClick","bracketsClick","iconClick","selectedChange","valueChange"],setup:function(e,t){var o=t.emit,l=(0,r.computed)(function(){return G(e.node.content)}),h=(0,r.computed)(function(){return"vjs-value vjs-value-".concat(l.value)}),m=(0,r.computed)(function(){return e.showDoubleQuotes?'"'.concat(e.node.key,'"'):e.node.key}),S=(0,r.computed)(function(){return e.selectableType==="multiple"}),d=(0,r.computed)(function(){return e.selectableType==="single"}),w=(0,r.computed)(function(){return e.nodeSelectable(e.node)&&(S.value||d.value)}),y=(0,r.reactive)({editing:!1}),P=function(u){var n,a,i=(a=(n=u.target)===null||n===void 0?void 0:n.value)==="null"?null:a==="undefined"?void 0:a==="true"||a!=="false"&&(a[0]+a[a.length-1]==='""'||a[0]+a[a.length-1]==="''"?a.slice(1,-1):typeof Number(a)=="number"&&!isNaN(Number(a))||a==="NaN"?Number(a):a);o("valueChange",i,e.node.path)},O=(0,r.computed)(function(){var u,n=(u=e.node)===null||u===void 0?void 0:u.content;return n===null?n="null":n===void 0&&(n="undefined"),l.value==="string"?'"'.concat(n,'"'):n+""}),D=function(){var u=e.renderNodeValue;return u?u({node:e.node,defaultValue:O.value}):O.value},L=function(){o("bracketsClick",!e.collapsed,e.node.path)},A=function(){o("iconClick",!e.collapsed,e.node.path)},x=function(){o("selectedChange",e.node)},E=function(){o("nodeClick",e.node),w.value&&e.selectOnClickNode&&o("selectedChange",e.node)},b=function(u){if(e.editable&&!y.editing){y.editing=!0;var n=function a(i){var c;i.target!==u.target&&((c=i.target)===null||c===void 0?void 0:c.parentElement)!==u.target&&(y.editing=!1,document.removeEventListener("click",a))};document.removeEventListener("click",n),document.addEventListener("click",n)}};return function(){var u,n=e.node;return(0,r.createVNode)("div",{class:{"vjs-tree-node":!0,"has-selector":e.showSelectController,"has-carets":e.showIcon,"is-highlight":e.highlightSelectedNode&&e.checked},onClick:E,style:e.style},[e.showLineNumber&&(0,r.createVNode)("span",{class:"vjs-node-index"},[n.id+1]),e.showSelectController&&w.value&&n.type!=="objectEnd"&&n.type!=="arrayEnd"&&(0,r.createVNode)(ce,{isMultiple:S.value,checked:e.checked,onChange:x},null),(0,r.createVNode)("div",{class:"vjs-indent"},[Array.from(Array(n.level)).map(function(a,i){return(0,r.createVNode)("div",{key:i,class:{"vjs-indent-unit":!0,"has-line":e.showLine}},null)}),e.showIcon&&(0,r.createVNode)(ue,{nodeType:n.type,onClick:A},null)]),n.key&&(0,r.createVNode)("span",{class:"vjs-key"},[(u=e.renderNodeKey,u?u({node:e.node,defaultKey:m.value||""}):m.value),(0,r.createVNode)("span",{class:"vjs-colon"},[":".concat(e.showKeyValueSpace?" ":"")])]),(0,r.createVNode)("span",null,[n.type!=="content"&&n.content?(0,r.createVNode)(ie,{data:n.content.toString(),onClick:L},null):(0,r.createVNode)("span",{class:h.value,onClick:!e.editable||e.editableTrigger&&e.editableTrigger!=="click"?void 0:b,onDblclick:e.editable&&e.editableTrigger==="dblclick"?b:void 0},[e.editable&&y.editing?(0,r.createVNode)("input",{value:O.value,onChange:P,style:{padding:"3px 8px",border:"1px solid #eee",boxShadow:"none",boxSizing:"border-box",borderRadius:5,fontFamily:"inherit"}},null):D()]),n.showComma&&(0,r.createVNode)("span",null,[","]),e.showLength&&e.collapsed&&(0,r.createVNode)("span",{class:"vjs-comment"},[(0,r.createTextVNode)(" // "),n.length,(0,r.createTextVNode)(" items ")])])])}}});function ee(e,t){var o=Object.keys(e);if(Object.getOwnPropertySymbols){var l=Object.getOwnPropertySymbols(e);t&&(l=l.filter(function(h){return Object.getOwnPropertyDescriptor(e,h).enumerable})),o.push.apply(o,l)}return o}function v(e){for(var t=1;t<arguments.length;t++){var o=arguments[t]!=null?arguments[t]:{};t%2?ee(Object(o),!0).forEach(function(l){q(e,l,o[l])}):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(o)):ee(Object(o)).forEach(function(l){Object.defineProperty(e,l,Object.getOwnPropertyDescriptor(o,l))})}return e}var de=(0,r.defineComponent)({name:"Tree",props:v(v({},X),{},{data:{type:[String,Number,Boolean,Array,Object],default:null},collapsedNodeLength:{type:Number,default:1/0},deep:{type:Number,default:1/0},pathCollapsible:{type:Function,default:function(){return!1}},rootPath:{type:String,default:"root"},virtual:{type:Boolean,default:!1},height:{type:Number,default:400},itemHeight:{type:Number,default:20},selectedValue:{type:[String,Array],default:function(){return""}},collapsedOnClickBrackets:{type:Boolean,default:!0},style:Object,onSelectedChange:{type:Function}}),slots:["renderNodeKey","renderNodeValue"],emits:["nodeClick","bracketsClick","iconClick","selectedChange","update:selectedValue","update:data"],setup:function(e,t){var o=t.emit,l=t.slots,h=(0,r.ref)(),m=(0,r.computed)(function(){return T(e.data,e.rootPath)}),S=function(n,a){return m.value.reduce(function(i,c){var s,f=c.level>=n||c.length>=a,k=(s=e.pathCollapsible)===null||s===void 0?void 0:s.call(e,c);return c.type!=="objectStart"&&c.type!=="arrayStart"||!f&&!k?i:v(v({},i),{},q({},c.path,1))},{})},d=(0,r.reactive)({translateY:0,visibleData:null,hiddenPaths:S(e.deep,e.collapsedNodeLength)}),w=(0,r.computed)(function(){for(var n=null,a=[],i=m.value.length,c=0;c<i;c++){var s=v(v({},m.value[c]),{},{id:c}),f=d.hiddenPaths[s.path];if(n&&n.path===s.path){var k=n.type==="objectStart",K=v(v(v({},s),n),{},{showComma:s.showComma,content:k?"{...}":"[...]",type:k?"objectCollapsed":"arrayCollapsed"});n=null,a.push(K)}else{if(f&&!n){n=s;continue}if(n)continue;a.push(s)}}return a}),y=(0,r.computed)(function(){var n=e.selectedValue;return n&&e.selectableType==="multiple"&&Array.isArray(n)?n:[n]}),P=(0,r.computed)(function(){return!e.selectableType||e.selectOnClickNode||e.showSelectController?"":"When selectableType is not null, selectOnClickNode and showSelectController cannot be false at the same time, because this will cause the selection to fail."}),O=function(){var n=w.value;if(e.virtual){var a,i=e.height/e.itemHeight,c=((a=h.value)===null||a===void 0?void 0:a.scrollTop)||0,s=Math.floor(c/e.itemHeight),f=s<0?0:s+i>n.length?n.length-i:s;f<0&&(f=0);var k=f+i;d.translateY=f*e.itemHeight,d.visibleData=n.filter(function(K,F){return F>=f&&F<k})}else d.visibleData=n},D=function(){O()},L=function(n){var a,i,c=n.path,s=e.selectableType;if(s==="multiple"){var f=y.value.findIndex(function(V){return V===c}),k=g(y.value);f!==-1?k.splice(f,1):k.push(c),o("update:selectedValue",k),o("selectedChange",k,g(y.value))}else if(s==="single"&&y.value[0]!==c){var K=(a=y.value,i=1,function(V){if(Array.isArray(V))return V}(a)||function(V,te){var B=V==null?null:typeof Symbol<"u"&&V[Symbol.iterator]||V["@@iterator"];if(B!=null){var ne,oe,Y=[],J=!0,re=!1;try{for(B=B.call(V);!(J=(ne=B.next()).done)&&(Y.push(ne.value),!te||Y.length!==te);J=!0);}catch(pe){re=!0,oe=pe}finally{try{J||B.return==null||B.return()}finally{if(re)throw oe}}return Y}}(a,i)||C(a,i)||function(){throw new TypeError(`Invalid attempt to destructure non-iterable instance.
In order to be iterable, non-array objects must have a [Symbol.iterator]() method.`)}())[0],F=c;o("update:selectedValue",F),o("selectedChange",F,K)}},A=function(n){o("nodeClick",n)},x=function(n,a){if(n)d.hiddenPaths=v(v({},d.hiddenPaths),{},q({},a,1));else{var i=v({},d.hiddenPaths);delete i[a],d.hiddenPaths=i}},E=function(n,a){e.collapsedOnClickBrackets&&x(n,a),o("bracketsClick",n)},b=function(n,a){x(n,a),o("iconClick",n)},u=function(n,a){var i=Q(e.data),c=e.rootPath;new Function("data","val","data".concat(a.slice(c.length),"=val"))(i,n),o("update:data",i)};return(0,r.watchEffect)(function(){P.value&&function(n){throw new Error("[VueJSONPretty] ".concat(n))}(P.value)}),(0,r.watchEffect)(function(){w.value&&O()}),(0,r.watch)(function(){return e.deep},function(n){n&&(d.hiddenPaths=S(n,e.collapsedNodeLength))}),(0,r.watch)(function(){return e.collapsedNodeLength},function(n){n&&(d.hiddenPaths=S(e.deep,n))}),function(){var n,a,i=(n=e.renderNodeKey)!==null&&n!==void 0?n:l.renderNodeKey,c=(a=e.renderNodeValue)!==null&&a!==void 0?a:l.renderNodeValue,s=d.visibleData&&d.visibleData.map(function(f){return(0,r.createVNode)(se,{key:f.id,node:f,collapsed:!!d.hiddenPaths[f.path],showDoubleQuotes:e.showDoubleQuotes,showLength:e.showLength,checked:y.value.includes(f.path),selectableType:e.selectableType,showLine:e.showLine,showLineNumber:e.showLineNumber,showSelectController:e.showSelectController,selectOnClickNode:e.selectOnClickNode,nodeSelectable:e.nodeSelectable,highlightSelectedNode:e.highlightSelectedNode,editable:e.editable,editableTrigger:e.editableTrigger,showIcon:e.showIcon,showKeyValueSpace:e.showKeyValueSpace,renderNodeKey:i,renderNodeValue:c,onNodeClick:A,onBracketsClick:E,onIconClick:b,onSelectedChange:L,onValueChange:u,style:e.itemHeight&&e.itemHeight!==20?{lineHeight:"".concat(e.itemHeight,"px")}:{}},null)});return(0,r.createVNode)("div",{ref:h,class:{"vjs-tree":!0,"is-virtual":e.virtual},onScroll:e.virtual?D:void 0,style:e.showLineNumber?v({paddingLeft:"".concat(12*Number(m.value.length.toString().length),"px")},e.style):e.style},[e.virtual?(0,r.createVNode)("div",{class:"vjs-tree-list",style:{height:"".concat(e.height,"px")}},[(0,r.createVNode)("div",{class:"vjs-tree-list-holder",style:{height:"".concat(w.value.length*e.itemHeight,"px")}},[(0,r.createVNode)("div",{class:"vjs-tree-list-holder-inner",style:{transform:"translateY(".concat(d.translateY,"px)")}},[s])])]):s])}}})}(),z}()})})(le);const Ne=fe(le.exports);export{Ne as V,Ce as a,ke as b,we as c,be as q};
