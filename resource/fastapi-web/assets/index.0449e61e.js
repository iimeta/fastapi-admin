import{x as Ae,y as Me,A as He,u as Ke,B as Ue,m as Re,v as Je,I as Qe,w as We,_ as Ye}from"./index.118dd0ea.js";import{u as Ge}from"./loading.b0aa7954.js";/* empty css               *//* empty css                *//* empty css              *//* empty css                *//* empty css               *//* empty css              *//* empty css                *//* empty css                *//* empty css               *//* empty css               */import{c as Ve,S as Ze}from"./sortable.esm.16d09f12.js";/* empty css               *//* empty css               *//* empty css                *//* empty css               *//* empty css               *//* empty css              */import{aN as Xe,x as et,d as tt,r as Fe,e as U,c as xe,w as Pe,bS as at,B as n,C as v,aH as a,aG as t,aL as V,aM as c,aJ as je,aI as De,aD as m,u as g,F as Y,D as nt,bv as ot,n as lt,g as rt,aK as it,aF as st,b2 as ut,b3 as ct,bB as dt,bC as pt,bD as mt,aT as ft,bE as _t,bF as ht,b6 as yt,bG as vt,ab as gt,aV as bt,bj as kt,bk as wt,bm as Ct,bn as Nt,b5 as St,bH as Vt,bJ as xt,bP as jt,bQ as Dt,ad as Et,bT as Ot,bU as Bt,bR as It,bb as Ft,b7 as Pt,aW as Tt,bK as At}from"./arco.a11b8b88.js";import{q as Lt}from"./model.8ec249e7.js";import{u as $t}from"./vue.e9a5701d.js";import"./chart.87d6227d.js";function zt(ee){return Ae.post("/api/v1/chat/page",ee)}function qt(ee){return Ae.get("/api/v1/chat/detail",{params:ee,paramsSerializer:de=>Me.stringify(de)})}var Le={exports:{}};(function(ee,de){(function(le,pe){ee.exports=pe(He)})(et,function(le){return function(){var pe={789:function(y){y.exports=le}},re={};function _(y){var L=re[y];if(L!==void 0)return L.exports;var k=re[y]={exports:{}};return pe[y](k,k.exports,_),k.exports}_.d=function(y,L){for(var k in L)_.o(L,k)&&!_.o(y,k)&&Object.defineProperty(y,k,{enumerable:!0,get:L[k]})},_.o=function(y,L){return Object.prototype.hasOwnProperty.call(y,L)},_.r=function(y){typeof Symbol<"u"&&Symbol.toStringTag&&Object.defineProperty(y,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(y,"__esModule",{value:!0})};var te={};return function(){function y(e,o){(o==null||o>e.length)&&(o=e.length);for(var r=0,h=new Array(o);r<o;r++)h[r]=e[r];return h}function L(e,o){if(e){if(typeof e=="string")return y(e,o);var r=Object.prototype.toString.call(e).slice(8,-1);return r==="Object"&&e.constructor&&(r=e.constructor.name),r==="Map"||r==="Set"?Array.from(e):r==="Arguments"||/^(?:Ui|I)nt(?:8|16|32)(?:Clamped)?Array$/.test(r)?y(e,o):void 0}}function k(e){return function(o){if(Array.isArray(o))return y(o)}(e)||function(o){if(typeof Symbol<"u"&&o[Symbol.iterator]!=null||o["@@iterator"]!=null)return Array.from(o)}(e)||L(e)||function(){throw new TypeError(`Invalid attempt to spread non-iterable instance.
In order to be iterable, non-array objects must have a [Symbol.iterator]() method.`)}()}function z(e,o,r){return o in e?Object.defineProperty(e,o,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[o]=r,e}_.r(te),_.d(te,{default:function(){return Ce}});var u=_(789),ie=(0,u.defineComponent)({props:{data:{required:!0,type:String},onClick:Function},render:function(){var e=this.data,o=this.onClick;return(0,u.createVNode)("span",{class:"vjs-tree-brackets",onClick:o},[e])}}),ae=(0,u.defineComponent)({emits:["change","update:modelValue"],props:{checked:{type:Boolean,default:!1},isMultiple:Boolean,onChange:Function},setup:function(e,o){var r=o.emit;return{uiType:(0,u.computed)(function(){return e.isMultiple?"checkbox":"radio"}),model:(0,u.computed)({get:function(){return e.checked},set:function(h){return r("update:modelValue",h)}})}},render:function(){var e=this.uiType,o=this.model,r=this.$emit;return(0,u.createVNode)("label",{class:["vjs-check-controller",o?"is-checked":""],onClick:function(h){return h.stopPropagation()}},[(0,u.createVNode)("span",{class:"vjs-check-controller-inner is-".concat(e)},null),(0,u.createVNode)("input",{checked:o,class:"vjs-check-controller-original is-".concat(e),type:e,onChange:function(){return r("change",o)}},null)])}}),ne=(0,u.defineComponent)({props:{nodeType:{required:!0,type:String},onClick:Function},render:function(){var e=this.nodeType,o=this.onClick,r=e==="objectStart"||e==="arrayStart";return r||e==="objectCollapsed"||e==="arrayCollapsed"?(0,u.createVNode)("span",{class:"vjs-carets vjs-carets-".concat(r?"open":"close"),onClick:o},[(0,u.createVNode)("svg",{viewBox:"0 0 1024 1024",focusable:"false","data-icon":"caret-down",width:"1em",height:"1em",fill:"currentColor","aria-hidden":"true"},[(0,u.createVNode)("path",{d:"M840.4 300H183.6c-19.7 0-30.7 20.8-18.5 35l328.4 380.8c9.4 10.9 27.5 10.9 37 0L858.9 335c12.2-14.2 1.2-35-18.5-35z"},null)])]):null}});function se(e){return se=typeof Symbol=="function"&&typeof Symbol.iterator=="symbol"?function(o){return typeof o}:function(o){return o&&typeof Symbol=="function"&&o.constructor===Symbol&&o!==Symbol.prototype?"symbol":typeof o},se(e)}function me(e){return Object.prototype.toString.call(e).slice(8,-1).toLowerCase()}function R(e){var o=arguments.length>1&&arguments[1]!==void 0?arguments[1]:"root",r=arguments.length>2&&arguments[2]!==void 0?arguments[2]:0,h=arguments.length>3?arguments[3]:void 0,i=h||{},T=i.key,H=i.index,j=i.type,s=j===void 0?"content":j,p=i.showComma,I=p!==void 0&&p,O=i.length,F=O===void 0?1:O,J=me(e);if(J==="array"){var q=Z(e.map(function(A,x,l){return R(A,"".concat(o,"[").concat(x,"]"),r+1,{index:x,showComma:x!==l.length-1,length:F,type:s})}));return[R("[",o,r,{showComma:!1,key:T,length:e.length,type:"arrayStart"})[0]].concat(q,R("]",o,r,{showComma:I,length:e.length,type:"arrayEnd"})[0])}if(J==="object"){var B=Object.keys(e),X=Z(B.map(function(A,x,l){return R(e[A],/^[a-zA-Z_]\w*$/.test(A)?"".concat(o,".").concat(A):"".concat(o,'["').concat(A,'"]'),r+1,{key:A,showComma:x!==l.length-1,length:F,type:s})}));return[R("{",o,r,{showComma:!1,key:T,index:H,length:B.length,type:"objectStart"})[0]].concat(X,R("}",o,r,{showComma:I,length:B.length,type:"objectEnd"})[0])}return[{content:e,level:r,key:T,index:H,path:o,showComma:I,length:F,type:s}]}function Z(e){if(typeof Array.prototype.flat=="function")return e.flat();for(var o=k(e),r=[];o.length;){var h=o.shift();Array.isArray(h)?o.unshift.apply(o,k(h)):r.push(h)}return r}function oe(e){var o=arguments.length>1&&arguments[1]!==void 0?arguments[1]:new WeakMap;if(e==null)return e;if(e instanceof Date)return new Date(e);if(e instanceof RegExp)return new RegExp(e);if(se(e)!=="object")return e;if(o.get(e))return o.get(e);if(Array.isArray(e)){var r=e.map(function(T){return oe(T,o)});return o.set(e,r),r}var h={};for(var i in e)h[i]=oe(e[i],o);return o.set(e,h),h}function fe(e,o){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var h=Object.getOwnPropertySymbols(e);o&&(h=h.filter(function(i){return Object.getOwnPropertyDescriptor(e,i).enumerable})),r.push.apply(r,h)}return r}function _e(e){for(var o=1;o<arguments.length;o++){var r=arguments[o]!=null?arguments[o]:{};o%2?fe(Object(r),!0).forEach(function(h){z(e,h,r[h])}):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):fe(Object(r)).forEach(function(h){Object.defineProperty(e,h,Object.getOwnPropertyDescriptor(r,h))})}return e}var he={showLength:{type:Boolean,default:!1},showDoubleQuotes:{type:Boolean,default:!0},renderNodeKey:Function,renderNodeValue:Function,selectableType:String,showSelectController:{type:Boolean,default:!1},showLine:{type:Boolean,default:!0},showLineNumber:{type:Boolean,default:!1},selectOnClickNode:{type:Boolean,default:!0},nodeSelectable:{type:Function,default:function(){return!0}},highlightSelectedNode:{type:Boolean,default:!0},showIcon:{type:Boolean,default:!1},showKeyValueSpace:{type:Boolean,default:!0},editable:{type:Boolean,default:!1},editableTrigger:{type:String,default:"click"},onNodeClick:{type:Function},onBracketsClick:{type:Function},onIconClick:{type:Function},onValueChange:{type:Function}},we=(0,u.defineComponent)({name:"TreeNode",props:_e(_e({},he),{},{node:{type:Object,required:!0},collapsed:Boolean,checked:Boolean,style:Object,onSelectedChange:{type:Function}}),emits:["nodeClick","bracketsClick","iconClick","selectedChange","valueChange"],setup:function(e,o){var r=o.emit,h=(0,u.computed)(function(){return me(e.node.content)}),i=(0,u.computed)(function(){return"vjs-value vjs-value-".concat(h.value)}),T=(0,u.computed)(function(){return e.showDoubleQuotes?'"'.concat(e.node.key,'"'):e.node.key}),H=(0,u.computed)(function(){return e.selectableType==="multiple"}),j=(0,u.computed)(function(){return e.selectableType==="single"}),s=(0,u.computed)(function(){return e.nodeSelectable(e.node)&&(H.value||j.value)}),p=(0,u.reactive)({editing:!1}),I=function(x){var l,d,N=(d=(l=x.target)===null||l===void 0?void 0:l.value)==="null"?null:d==="undefined"?void 0:d==="true"||d!=="false"&&(d[0]+d[d.length-1]==='""'||d[0]+d[d.length-1]==="''"?d.slice(1,-1):typeof Number(d)=="number"&&!isNaN(Number(d))||d==="NaN"?Number(d):d);r("valueChange",N,e.node.path)},O=(0,u.computed)(function(){var x,l=(x=e.node)===null||x===void 0?void 0:x.content;return l===null?l="null":l===void 0&&(l="undefined"),h.value==="string"?'"'.concat(l,'"'):l+""}),F=function(){var x=e.renderNodeValue;return x?x({node:e.node,defaultValue:O.value}):O.value},J=function(){r("bracketsClick",!e.collapsed,e.node.path)},q=function(){r("iconClick",!e.collapsed,e.node.path)},B=function(){r("selectedChange",e.node)},X=function(){r("nodeClick",e.node),s.value&&e.selectOnClickNode&&r("selectedChange",e.node)},A=function(x){if(e.editable&&!p.editing){p.editing=!0;var l=function d(N){var S;N.target!==x.target&&((S=N.target)===null||S===void 0?void 0:S.parentElement)!==x.target&&(p.editing=!1,document.removeEventListener("click",d))};document.removeEventListener("click",l),document.addEventListener("click",l)}};return function(){var x,l=e.node;return(0,u.createVNode)("div",{class:{"vjs-tree-node":!0,"has-selector":e.showSelectController,"has-carets":e.showIcon,"is-highlight":e.highlightSelectedNode&&e.checked},onClick:X,style:e.style},[e.showLineNumber&&(0,u.createVNode)("span",{class:"vjs-node-index"},[l.id+1]),e.showSelectController&&s.value&&l.type!=="objectEnd"&&l.type!=="arrayEnd"&&(0,u.createVNode)(ae,{isMultiple:H.value,checked:e.checked,onChange:B},null),(0,u.createVNode)("div",{class:"vjs-indent"},[Array.from(Array(l.level)).map(function(d,N){return(0,u.createVNode)("div",{key:N,class:{"vjs-indent-unit":!0,"has-line":e.showLine}},null)}),e.showIcon&&(0,u.createVNode)(ne,{nodeType:l.type,onClick:q},null)]),l.key&&(0,u.createVNode)("span",{class:"vjs-key"},[(x=e.renderNodeKey,x?x({node:e.node,defaultKey:T.value||""}):T.value),(0,u.createVNode)("span",{class:"vjs-colon"},[":".concat(e.showKeyValueSpace?" ":"")])]),(0,u.createVNode)("span",null,[l.type!=="content"&&l.content?(0,u.createVNode)(ie,{data:l.content.toString(),onClick:J},null):(0,u.createVNode)("span",{class:i.value,onClick:!e.editable||e.editableTrigger&&e.editableTrigger!=="click"?void 0:A,onDblclick:e.editable&&e.editableTrigger==="dblclick"?A:void 0},[e.editable&&p.editing?(0,u.createVNode)("input",{value:O.value,onChange:I,style:{padding:"3px 8px",border:"1px solid #eee",boxShadow:"none",boxSizing:"border-box",borderRadius:5,fontFamily:"inherit"}},null):F()]),l.showComma&&(0,u.createVNode)("span",null,[","]),e.showLength&&e.collapsed&&(0,u.createVNode)("span",{class:"vjs-comment"},[(0,u.createTextVNode)(" // "),l.length,(0,u.createTextVNode)(" items ")])])])}}});function ye(e,o){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var h=Object.getOwnPropertySymbols(e);o&&(h=h.filter(function(i){return Object.getOwnPropertyDescriptor(e,i).enumerable})),r.push.apply(r,h)}return r}function P(e){for(var o=1;o<arguments.length;o++){var r=arguments[o]!=null?arguments[o]:{};o%2?ye(Object(r),!0).forEach(function(h){z(e,h,r[h])}):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):ye(Object(r)).forEach(function(h){Object.defineProperty(e,h,Object.getOwnPropertyDescriptor(r,h))})}return e}var Ce=(0,u.defineComponent)({name:"Tree",props:P(P({},he),{},{data:{type:[String,Number,Boolean,Array,Object],default:null},collapsedNodeLength:{type:Number,default:1/0},deep:{type:Number,default:1/0},pathCollapsible:{type:Function,default:function(){return!1}},rootPath:{type:String,default:"root"},virtual:{type:Boolean,default:!1},height:{type:Number,default:400},itemHeight:{type:Number,default:20},selectedValue:{type:[String,Array],default:function(){return""}},collapsedOnClickBrackets:{type:Boolean,default:!0},style:Object,onSelectedChange:{type:Function}}),slots:["renderNodeKey","renderNodeValue"],emits:["nodeClick","bracketsClick","iconClick","selectedChange","update:selectedValue","update:data"],setup:function(e,o){var r=o.emit,h=o.slots,i=(0,u.ref)(),T=(0,u.computed)(function(){return R(e.data,e.rootPath)}),H=function(l,d){return T.value.reduce(function(N,S){var D,E=S.level>=l||S.length>=d,$=(D=e.pathCollapsible)===null||D===void 0?void 0:D.call(e,S);return S.type!=="objectStart"&&S.type!=="arrayStart"||!E&&!$?N:P(P({},N),{},z({},S.path,1))},{})},j=(0,u.reactive)({translateY:0,visibleData:null,hiddenPaths:H(e.deep,e.collapsedNodeLength)}),s=(0,u.computed)(function(){for(var l=null,d=[],N=T.value.length,S=0;S<N;S++){var D=P(P({},T.value[S]),{},{id:S}),E=j.hiddenPaths[D.path];if(l&&l.path===D.path){var $=l.type==="objectStart",Q=P(P(P({},D),l),{},{showComma:D.showComma,content:$?"{...}":"[...]",type:$?"objectCollapsed":"arrayCollapsed"});l=null,d.push(Q)}else{if(E&&!l){l=D;continue}if(l)continue;d.push(D)}}return d}),p=(0,u.computed)(function(){var l=e.selectedValue;return l&&e.selectableType==="multiple"&&Array.isArray(l)?l:[l]}),I=(0,u.computed)(function(){return!e.selectableType||e.selectOnClickNode||e.showSelectController?"":"When selectableType is not null, selectOnClickNode and showSelectController cannot be false at the same time, because this will cause the selection to fail."}),O=function(){var l=s.value;if(e.virtual){var d,N=e.height/e.itemHeight,S=((d=i.value)===null||d===void 0?void 0:d.scrollTop)||0,D=Math.floor(S/e.itemHeight),E=D<0?0:D+N>l.length?l.length-N:D;E<0&&(E=0);var $=E+N;j.translateY=E*e.itemHeight,j.visibleData=l.filter(function(Q,W){return W>=E&&W<$})}else j.visibleData=l},F=function(){O()},J=function(l){var d,N,S=l.path,D=e.selectableType;if(D==="multiple"){var E=p.value.findIndex(function(K){return K===S}),$=k(p.value);E!==-1?$.splice(E,1):$.push(S),r("update:selectedValue",$),r("selectedChange",$,k(p.value))}else if(D==="single"&&p.value[0]!==S){var Q=(d=p.value,N=1,function(K){if(Array.isArray(K))return K}(d)||function(K,ve){var G=K==null?null:typeof Symbol<"u"&&K[Symbol.iterator]||K["@@iterator"];if(G!=null){var ge,be,ue=[],ce=!0,ke=!1;try{for(G=G.call(K);!(ce=(ge=G.next()).done)&&(ue.push(ge.value),!ve||ue.length!==ve);ce=!0);}catch(b){ke=!0,be=b}finally{try{ce||G.return==null||G.return()}finally{if(ke)throw be}}return ue}}(d,N)||L(d,N)||function(){throw new TypeError(`Invalid attempt to destructure non-iterable instance.
In order to be iterable, non-array objects must have a [Symbol.iterator]() method.`)}())[0],W=S;r("update:selectedValue",W),r("selectedChange",W,Q)}},q=function(l){r("nodeClick",l)},B=function(l,d){if(l)j.hiddenPaths=P(P({},j.hiddenPaths),{},z({},d,1));else{var N=P({},j.hiddenPaths);delete N[d],j.hiddenPaths=N}},X=function(l,d){e.collapsedOnClickBrackets&&B(l,d),r("bracketsClick",l)},A=function(l,d){B(l,d),r("iconClick",l)},x=function(l,d){var N=oe(e.data),S=e.rootPath;new Function("data","val","data".concat(d.slice(S.length),"=val"))(N,l),r("update:data",N)};return(0,u.watchEffect)(function(){I.value&&function(l){throw new Error("[VueJSONPretty] ".concat(l))}(I.value)}),(0,u.watchEffect)(function(){s.value&&O()}),(0,u.watch)(function(){return e.deep},function(l){l&&(j.hiddenPaths=H(l,e.collapsedNodeLength))}),(0,u.watch)(function(){return e.collapsedNodeLength},function(l){l&&(j.hiddenPaths=H(e.deep,l))}),function(){var l,d,N=(l=e.renderNodeKey)!==null&&l!==void 0?l:h.renderNodeKey,S=(d=e.renderNodeValue)!==null&&d!==void 0?d:h.renderNodeValue,D=j.visibleData&&j.visibleData.map(function(E){return(0,u.createVNode)(we,{key:E.id,node:E,collapsed:!!j.hiddenPaths[E.path],showDoubleQuotes:e.showDoubleQuotes,showLength:e.showLength,checked:p.value.includes(E.path),selectableType:e.selectableType,showLine:e.showLine,showLineNumber:e.showLineNumber,showSelectController:e.showSelectController,selectOnClickNode:e.selectOnClickNode,nodeSelectable:e.nodeSelectable,highlightSelectedNode:e.highlightSelectedNode,editable:e.editable,editableTrigger:e.editableTrigger,showIcon:e.showIcon,showKeyValueSpace:e.showKeyValueSpace,renderNodeKey:N,renderNodeValue:S,onNodeClick:q,onBracketsClick:X,onIconClick:A,onSelectedChange:J,onValueChange:x,style:e.itemHeight&&e.itemHeight!==20?{lineHeight:"".concat(e.itemHeight,"px")}:{}},null)});return(0,u.createVNode)("div",{ref:i,class:{"vjs-tree":!0,"is-virtual":e.virtual},onScroll:e.virtual?F:void 0,style:e.showLineNumber?P({paddingLeft:"".concat(12*Number(T.value.length.toString().length),"px")},e.style):e.style},[e.virtual?(0,u.createVNode)("div",{class:"vjs-tree-list",style:{height:"".concat(e.height,"px")}},[(0,u.createVNode)("div",{class:"vjs-tree-list-holder",style:{height:"".concat(s.value.length*e.itemHeight,"px")}},[(0,u.createVNode)("div",{class:"vjs-tree-list-holder-inner",style:{transform:"translateY(".concat(j.translateY,"px)")}},[D])])]):D])}}})}(),te}()})})(Le);const Te=Xe(Le.exports);const Mt={class:"container"},Ht={class:"action-icon"},Kt={class:"action-icon"},Ut={id:"tableSetting"},Rt={style:{"margin-right":"4px",cursor:"move"}},Jt={class:"title"},Qt={key:0,class:"circle red"},Wt={key:1,class:"circle"},Yt={style:{margin:"10px 0 0 10px"}},Gt={key:1},Zt={key:1},Xt={key:1},ea={key:1},ta={key:1},aa={key:1},na={key:1},oa={key:1},la={key:1},ra={key:1},ia={key:1},sa={key:1},ua={key:1},ca={key:1},da={key:1},pa={key:1},ma={key:1},fa={key:1},_a={key:1},ha={key:1},ya={key:1},va={key:1},ga={key:1},ba={key:1},ka={key:1},wa={key:1},Ca={key:1},Na={key:1},Sa={key:1},Va={key:1},xa={key:1},ja={name:"AppChatList"},Da=tt({...ja,setup(ee){const de=Fe({type:"checkbox",showCheckedAll:!0,onlyCurrent:!1}),le=U([]);(async()=>{try{const{data:s}=await Lt();le.value=s.items}catch{}})();const re=()=>({app_id:U(),key:"",models:[],total_time:U(),status:U(),req_times:[]}),{loading:_,setLoading:te}=Ge(!0),{t:y}=Ke(),L=U([]),k=U(re()),z=U([]),u=U([]),ie=U("small"),ae={current:1,pageSize:10,showTotal:!0,showPageSize:!0},ne=Fe({...ae}),se=xe(()=>[{name:y("searchTable.size.mini"),value:"mini"},{name:y("searchTable.size.small"),value:"small"},{name:y("searchTable.size.medium"),value:"medium"},{name:y("searchTable.size.large"),value:"large"}]),me=xe(()=>[{title:y("chat.columns.user_id"),dataIndex:"user_id",slotName:"user_id",align:"center",width:75},{title:y("chat.columns.app_id"),dataIndex:"app_id",slotName:"app_id",align:"center",width:75},{title:y("chat.columns.model"),dataIndex:"model",slotName:"model",align:"center"},{title:y("chat.columns.prompt_tokens"),dataIndex:"prompt_tokens",slotName:"prompt_tokens",align:"center"},{title:y("chat.columns.completion_tokens"),dataIndex:"completion_tokens",slotName:"completion_tokens",align:"center"},{title:y("chat.columns.total_tokens"),dataIndex:"total_tokens",slotName:"total_tokens",align:"center"},{title:y("chat.columns.stream"),dataIndex:"stream",slotName:"stream",align:"center"},{title:y("chat.columns.conn_time"),dataIndex:"conn_time",slotName:"conn_time",align:"center"},{title:y("chat.columns.duration"),dataIndex:"duration",slotName:"duration",align:"center"},{title:y("chat.columns.total_time"),dataIndex:"total_time",slotName:"total_time",align:"center"},{title:y("chat.columns.internal_time"),dataIndex:"internal_time",slotName:"internal_time",align:"center"},{title:y("chat.columns.status"),dataIndex:"status",slotName:"status",align:"center",width:75},{title:y("chat.columns.req_time"),dataIndex:"req_time",slotName:"req_time",align:"center",width:132},{title:y("chat.columns.operations"),dataIndex:"operations",slotName:"operations",align:"center",width:75}]),R=xe(()=>[{label:y("chat.dict.status.1"),value:1},{label:y("chat.dict.status.-1"),value:-1}]),Z=async(s={current:1,pageSize:10})=>{te(!0);try{const{data:p}=await zt(s);L.value=p.items,ne.current=s.current,ne.pageSize=s.pageSize,ne.total=p.paging.total}catch{}finally{te(!1)}},oe=()=>{Z({...ae,...k.value})},fe=s=>{Z({...ae,...k.value,current:s})},_e=s=>{Z({...ae,...k.value,pageSize:s})};Z();const he=()=>{k.value=re()},we=(s,p)=>{ie.value=s},ye=(s,p,I)=>{s?z.value.splice(I,0,p):z.value=u.value.filter(O=>O.dataIndex!==p.dataIndex)},P=(s,p,I,O=!1)=>{const F=O?Ve(s):s;return p>-1&&I>-1&&F.splice(p,1,F.splice(I,1,F[p]).pop()),F},Ce=s=>{s&&lt(()=>{const p=document.getElementById("tableSetting");new Ze(p,{onEnd(I){const{oldIndex:O,newIndex:F}=I;P(z.value,O,F),P(u.value,O,F)}})})};Pe(()=>me.value,s=>{z.value=Ve(s),z.value.forEach((p,I)=>{p.checked=!0}),u.value=Ve(z.value)},{deep:!0,immediate:!0});const e=U(!1),{copy:o,copied:r}=$t(),{proxy:h}=rt(),i=U({}),T=async s=>{e.value=!0,_.value=!0;try{const{data:p}=await qt({id:s});i.value=p}catch{}finally{_.value=!1}},H=()=>{e.value=!1},j=s=>{o(s)};return Pe(r,()=>{r.value&&h.$message.success("\u590D\u5236\u6210\u529F")}),(s,p)=>{const I=Ue,O=it,F=st,J=ut,q=ct,B=dt,X=pt,A=mt,x=ft,l=_t,d=ht,N=yt,S=vt,D=gt,E=bt,$=Re,Q=kt,W=wt,K=Je,ve=Ct,G=Nt,ge=Qe,be=We,ue=St,ce=Vt,ke=xt,b=jt,w=Dt,Ee=Et,C=Ot,M=Bt,Ne=It,Oe=Ft,Be=Pt,$e=Tt,ze=At,qe=at("permission");return n(),v("div",Mt,[a(F,{class:"container-breadcrumb"},{default:t(()=>[a(O,null,{default:t(()=>[a(I)]),_:1}),a(O,null,{default:t(()=>[V(c(s.$t("menu.chat")),1)]),_:1}),a(O,null,{default:t(()=>[V(c(s.$t("menu.chat.list")),1)]),_:1})]),_:1}),a(ze,{class:"general-card",title:s.$t("menu.chat.list"),bordered:!1},{default:t(()=>[a(d,null,{default:t(()=>[a(B,{flex:1},{default:t(()=>[a(N,{model:k.value,"label-col-props":{span:5},"wrapper-col-props":{span:18},"label-align":"left"},{default:t(()=>[a(d,{gutter:16},{default:t(()=>[a(B,{span:8},{default:t(()=>[a(q,{field:"app_id",label:s.$t("chat.form.app_id")},{default:t(()=>[a(J,{modelValue:k.value.app_id,"onUpdate:modelValue":p[0]||(p[0]=f=>k.value.app_id=f),placeholder:s.$t("chat.form.app_id.placeholder"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1}),a(B,{span:8},{default:t(()=>[a(q,{field:"key",label:s.$t("chat.form.key")},{default:t(()=>[a(J,{modelValue:k.value.key,"onUpdate:modelValue":p[1]||(p[1]=f=>k.value.key=f),placeholder:s.$t("chat.form.key.placeholder"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1}),a(B,{span:8},{default:t(()=>[a(q,{field:"models",label:s.$t("chat.form.models")},{default:t(()=>[a(A,{modelValue:k.value.models,"onUpdate:modelValue":p[2]||(p[2]=f=>k.value.models=f),placeholder:s.$t("chat.form.selectDefault"),"max-tag-count":3,multiple:"","allow-search":"","allow-clear":""},{default:t(()=>[(n(!0),v(je,null,De(le.value,f=>(n(),m(X,{key:f.id,value:f.id,label:f.name},null,8,["value","label"]))),128))]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1}),a(B,{span:8},{default:t(()=>[a(q,{field:"total_time",label:s.$t("chat.form.total_time")},{default:t(()=>[a(x,{modelValue:k.value.total_time,"onUpdate:modelValue":p[3]||(p[3]=f=>k.value.total_time=f),precision:0,placeholder:s.$t("chat.form.total_time.placeholder"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1}),a(B,{span:8},{default:t(()=>[a(q,{field:"status",label:s.$t("chat.form.status")},{default:t(()=>[a(A,{modelValue:k.value.status,"onUpdate:modelValue":p[4]||(p[4]=f=>k.value.status=f),options:g(R),placeholder:s.$t("chat.form.selectDefault"),"allow-clear":""},null,8,["modelValue","options","placeholder"])]),_:1},8,["label"])]),_:1}),a(B,{span:8},{default:t(()=>[a(q,{field:"req_times",label:s.$t("chat.form.req_times")},{default:t(()=>[a(l,{modelValue:k.value.req_times,"onUpdate:modelValue":p[5]||(p[5]=f=>k.value.req_times=f),"show-time":""},null,8,["modelValue"])]),_:1},8,["label"])]),_:1})]),_:1})]),_:1},8,["model"])]),_:1}),a(S,{style:{height:"84px"},direction:"vertical"}),a(B,{flex:"86px",style:{"text-align":"right"}},{default:t(()=>[a(Q,{direction:"vertical",size:18},{default:t(()=>[a(E,{type:"primary",onClick:oe},{icon:t(()=>[a(D)]),default:t(()=>[V(" "+c(s.$t("chat.form.search")),1)]),_:1}),a(E,{onClick:he},{icon:t(()=>[a($)]),default:t(()=>[V(" "+c(s.$t("chat.form.reset")),1)]),_:1})]),_:1})]),_:1})]),_:1}),a(S,{style:{"margin-top":"0"}}),a(d,{style:{"margin-bottom":"16px"}},{default:t(()=>[a(B,{span:12},{default:t(()=>[a(Q)]),_:1}),a(B,{span:12,style:{display:"flex","align-items":"center","justify-content":"end"}},{default:t(()=>[a(W,{content:s.$t("searchTable.actions.refresh")},{default:t(()=>[Y("div",{class:"action-icon",onClick:oe},[a($,{size:"18"})])]),_:1},8,["content"]),a(G,{onSelect:we},{content:t(()=>[(n(!0),v(je,null,De(g(se),f=>(n(),m(ve,{key:f.value,value:f.value,class:nt({active:f.value===ie.value})},{default:t(()=>[Y("span",null,c(f.name),1)]),_:2},1032,["value","class"]))),128))]),default:t(()=>[a(W,{content:s.$t("searchTable.actions.density")},{default:t(()=>[Y("div",Ht,[a(K,{size:"18"})])]),_:1},8,["content"])]),_:1}),a(W,{content:s.$t("searchTable.actions.columnSetting")},{default:t(()=>[a(ce,{trigger:"click",position:"bl",onPopupVisibleChange:Ce},{content:t(()=>[Y("div",Ut,[(n(!0),v(je,null,De(u.value,(f,Ie)=>(n(),v("div",{key:f.dataIndex,class:"setting"},[Y("div",Rt,[a(be)]),Y("div",null,[a(ue,{modelValue:f.checked,"onUpdate:modelValue":Se=>f.checked=Se,onChange:Se=>ye(Se,f,Ie)},null,8,["modelValue","onUpdate:modelValue","onChange"])]),Y("div",Jt,c(f.title==="#"?"\u5E8F\u5217\u53F7":f.title),1)]))),128))])]),default:t(()=>[Y("div",Kt,[a(ge,{size:"18"})])]),_:1})]),_:1},8,["content"])]),_:1})]),_:1}),a(ke,{"row-key":"id",loading:g(_),pagination:ne,columns:z.value,data:L.value,bordered:!1,size:ie.value,"row-selection":de,onPageChange:fe,onPageSizeChange:_e},{prompt_tokens:t(({record:f})=>[V(c(f.prompt_tokens||"-"),1)]),completion_tokens:t(({record:f})=>[V(c(f.completion_tokens||"-"),1)]),total_tokens:t(({record:f})=>[V(c(f.total_tokens||"-"),1)]),stream:t(({record:f})=>[V(c(s.$t(`chat.dict.stream.${f.stream}`)),1)]),conn_time:t(({record:f})=>[V(c(f.conn_time||"-"),1)]),duration:t(({record:f})=>[V(c(f.duration||"-"),1)]),total_time:t(({record:f})=>[V(c(f.total_time||"-"),1)]),status:t(({record:f})=>[f.status===-1?(n(),v("span",Qt)):(n(),v("span",Wt)),V(" "+c(s.$t(`chat.dict.status.${f.status}`)),1)]),operations:t(({record:f})=>[a(E,{type:"text",size:"small",onClick:Ie=>T(f.id)},{default:t(()=>[V(c(s.$t("chat.columns.operations.view")),1)]),_:2},1032,["onClick"])]),_:1},8,["loading","pagination","columns","data","size","row-selection"]),a($e,{title:"\u65E5\u5FD7\u8BE6\u60C5",visible:e.value,width:660,footer:!1,"unmount-on-close":"","render-to-body":"",onCancel:H},{default:t(()=>[Y("div",Yt,[a(Ne,{column:2,bordered:""},{default:t(()=>[a(C,{label:"\u65E5\u5FD7ID",span:2},{default:t(()=>[g(_)?(n(),m(w,{key:0,animation:!0},{default:t(()=>[a(b,{widths:["200px"],rows:1})]),_:1})):(n(),v("span",Gt,[V(c(i.value.trace_id)+" ",1),a(Ee,{class:"copy-btn",onClick:p[6]||(p[6]=f=>j(i.value.trace_id))})]))]),_:1}),a(C,{label:"\u8C03\u7528\u5BC6\u94A5",span:2},{default:t(()=>[g(_)?(n(),m(w,{key:0,animation:!0},{default:t(()=>[a(b,{widths:["200px"],rows:1})]),_:1})):(n(),v("span",Zt,[V(c(i.value.creator)+" ",1),a(Ee,{class:"copy-btn",onClick:p[7]||(p[7]=f=>j(i.value.creator))})]))]),_:1}),a(C,{label:"\u7528\u6237ID"},{default:t(()=>[g(_)?(n(),m(w,{key:0,animation:!0},{default:t(()=>[a(b,{rows:1})]),_:1})):(n(),v("span",Xt,c(i.value.user_id),1))]),_:1}),a(C,{label:"\u5E94\u7528ID"},{default:t(()=>[g(_)?(n(),m(w,{key:0,animation:!0},{default:t(()=>[a(b,{rows:1})]),_:1})):(n(),v("span",ea,c(i.value.app_id),1))]),_:1}),a(C,{label:"\u516C\u53F8"},{default:t(()=>[g(_)?(n(),m(w,{key:0,animation:!0},{default:t(()=>[a(b,{rows:1})]),_:1})):(n(),v("span",ta,c(i.value.corp),1))]),_:1}),a(C,{label:"\u6A21\u578BID"},{default:t(()=>[g(_)?(n(),m(w,{key:0,animation:!0},{default:t(()=>[a(b,{rows:1})]),_:1})):(n(),v("span",aa,c(i.value.model_id),1))]),_:1}),a(C,{label:"\u6A21\u578B\u540D\u79F0"},{default:t(()=>[g(_)?(n(),m(w,{key:0,animation:!0},{default:t(()=>[a(b,{widths:["200px"],rows:1})]),_:1})):(n(),v("span",na,c(i.value.name),1))]),_:1}),a(C,{label:"\u6A21\u578B"},{default:t(()=>[g(_)?(n(),m(w,{key:0,animation:!0},{default:t(()=>[a(b,{widths:["200px"],rows:1})]),_:1})):(n(),v("span",oa,c(i.value.model),1))]),_:1}),a(C,{label:"\u6A21\u578B\u7C7B\u578B"},{default:t(()=>[g(_)?(n(),m(w,{key:0,animation:!0},{default:t(()=>[a(b,{rows:1})]),_:1})):(n(),v("span",la,c(i.value.type),1))]),_:1}),a(C,{label:"\u6D41\u5F0F"},{default:t(()=>[g(_)?(n(),m(w,{key:0,animation:!0},{default:t(()=>[a(b,{rows:1})]),_:1})):(n(),v("span",ra,c(i.value.stream),1))]),_:1}),a(C,{label:"\u542F\u7528\u4EE3\u7406"},{default:t(()=>[g(_)?(n(),m(w,{key:0,animation:!0},{default:t(()=>[a(b,{rows:1})]),_:1})):(n(),v("span",ia,c(i.value.is_enable_model_agent),1))]),_:1}),a(C,{label:"\u4EE3\u7406ID"},{default:t(()=>[g(_)?(n(),m(w,{key:0,animation:!0},{default:t(()=>[a(b,{rows:1})]),_:1})):(n(),v("span",sa,c(i.value.model_agent_id),1))]),_:1}),a(C,{label:"\u5BC6\u94A5",span:2},{default:t(()=>[g(_)?(n(),m(w,{key:0,animation:!0},{default:t(()=>[a(b,{rows:1})]),_:1})):(n(),v("span",ua,c(i.value.key),1))]),_:1}),a(C,{label:"\u63D0\u95EE",span:2},{default:t(()=>[g(_)?(n(),m(w,{key:0,animation:!0},{default:t(()=>[a(b,{rows:1})]),_:1})):(n(),v("span",ca,c(i.value.prompt),1))]),_:1}),a(C,{label:"\u56DE\u7B54",span:2},{default:t(()=>[g(_)?(n(),m(w,{key:0,animation:!0},{default:t(()=>[a(b,{rows:1})]),_:1})):(n(),v("span",da,c(i.value.completion),1))]),_:1}),a(C,{label:"\u63D0\u95EE\u500D\u7387"},{default:t(()=>[g(_)?(n(),m(w,{key:0,animation:!0},{default:t(()=>[a(b,{rows:1})]),_:1})):(n(),v("span",pa,c(i.value.prompt_ratio),1))]),_:1}),a(C,{label:"\u56DE\u7B54\u500D\u7387"},{default:t(()=>[g(_)?(n(),m(w,{key:0,animation:!0},{default:t(()=>[a(b,{rows:1})]),_:1})):(n(),v("span",ma,c(i.value.completion_ratio),1))]),_:1}),a(C,{label:"\u63D0\u95EE\u4EE4\u724C\u6570"},{default:t(()=>[g(_)?(n(),m(w,{key:0,animation:!0},{default:t(()=>[a(b,{rows:1})]),_:1})):(n(),v("span",fa,c(i.value.prompt_tokens),1))]),_:1}),a(C,{label:"\u56DE\u7B54\u4EE4\u724C\u6570"},{default:t(()=>[g(_)?(n(),m(w,{key:0,animation:!0},{default:t(()=>[a(b,{rows:1})]),_:1})):(n(),v("span",_a,c(i.value.completion_tokens),1))]),_:1}),a(C,{label:"\u603B\u4EE4\u724C\u6570"},{default:t(()=>[g(_)?(n(),m(w,{key:0,animation:!0},{default:t(()=>[a(b,{rows:1})]),_:1})):(n(),v("span",ha,c(i.value.total_tokens),1))]),_:1}),a(C,{label:"\u72B6\u6001"},{default:t(()=>[g(_)?(n(),m(w,{key:0,animation:!0},{default:t(()=>[a(b,{widths:["200px"],rows:1})]),_:1})):(n(),v("span",ya,c(i.value.status),1))]),_:1}),a(C,{label:"\u8FDE\u63A5\u65F6\u95F4"},{default:t(()=>[g(_)?(n(),m(w,{key:0,animation:!0},{default:t(()=>[a(b,{widths:["200px"],rows:1})]),_:1})):(n(),v("span",va,[i.value.conn_time>1500?(n(),m(M,{key:0,color:"red"},{default:t(()=>[V(c(i.value.conn_time)+" ms ",1)]),_:1})):i.value.conn_time>1e3?(n(),m(M,{key:1,color:"orange"},{default:t(()=>[V(c(i.value.conn_time)+" ms ",1)]),_:1})):(n(),m(M,{key:2,color:"green"},{default:t(()=>[V(c(i.value.conn_time)+" ms",1)]),_:1}))]))]),_:1}),a(C,{label:"\u6301\u7EED\u65F6\u95F4"},{default:t(()=>[g(_)?(n(),m(w,{key:0,animation:!0},{default:t(()=>[a(b,{widths:["200px"],rows:1})]),_:1})):(n(),v("span",ga,[i.value.duration>9e4?(n(),m(M,{key:0,color:"red"},{default:t(()=>[V(c(i.value.duration)+" ms ",1)]),_:1})):i.value.duration>6e4?(n(),m(M,{key:1,color:"orange"},{default:t(()=>[V(c(i.value.duration)+" ms ",1)]),_:1})):(n(),m(M,{key:2,color:"green"},{default:t(()=>[V(c(i.value.duration)+" ms",1)]),_:1}))]))]),_:1}),a(C,{label:"\u603B\u65F6\u95F4"},{default:t(()=>[g(_)?(n(),m(w,{key:0,animation:!0},{default:t(()=>[a(b,{widths:["200px"],rows:1})]),_:1})):(n(),v("span",ba,[i.value.total_time>12e4?(n(),m(M,{key:0,color:"red"},{default:t(()=>[V(c(i.value.total_time)+" ms ",1)]),_:1})):i.value.total_time>6e4?(n(),m(M,{key:1,color:"orange"},{default:t(()=>[V(c(i.value.total_time)+" ms ",1)]),_:1})):(n(),m(M,{key:2,color:"green"},{default:t(()=>[V(c(i.value.total_time)+" ms",1)]),_:1}))]))]),_:1}),a(C,{label:"\u5185\u8017\u65F6\u95F4"},{default:t(()=>[g(_)?(n(),m(w,{key:0,animation:!0},{default:t(()=>[a(b,{widths:["200px"],rows:1})]),_:1})):(n(),v("span",ka,[i.value.internal_time>500?(n(),m(M,{key:0,color:"red"},{default:t(()=>[V(c(i.value.internal_time)+" ms ",1)]),_:1})):i.value.internal_time>100?(n(),m(M,{key:1,color:"orange"},{default:t(()=>[V(c(i.value.internal_time)+" ms ",1)]),_:1})):(n(),m(M,{key:2,color:"green"},{default:t(()=>[V(c(i.value.internal_time)+" ms",1)]),_:1}))]))]),_:1}),a(C,{label:"\u8BF7\u6C42\u65F6\u95F4"},{default:t(()=>[g(_)?(n(),m(w,{key:0,animation:!0},{default:t(()=>[a(b,{widths:["200px"],rows:1})]),_:1})):(n(),v("span",wa,c(i.value.req_time),1))]),_:1}),a(C,{label:"\u5BA2\u6237\u7AEFIP"},{default:t(()=>[g(_)?(n(),m(w,{key:0,animation:!0},{default:t(()=>[a(b,{widths:["200px"],rows:1})]),_:1})):(n(),v("span",Ca,c(i.value.client_ip),1))]),_:1}),ot((n(),m(C,{label:"\u8FDC\u7A0BIP"},{default:t(()=>[g(_)?(n(),m(w,{key:0,animation:!0},{default:t(()=>[a(b,{widths:["200px"],rows:1})]),_:1})):(n(),v("span",Na,c(i.value.remote_ip),1))]),_:1})),[[qe,["admin"]]]),a(C,{label:"\u9519\u8BEF\u4FE1\u606F",span:2},{default:t(()=>[g(_)?(n(),m(w,{key:0,animation:!0},{default:t(()=>[a(b,{widths:["200px"],rows:1})]),_:1})):(n(),v("span",Sa,c(i.value.err_msg),1))]),_:1})]),_:1}),a(Ne,{layout:"inline-vertical",column:2,style:{"margin-top":"10px",position:"relative"}},{default:t(()=>[a(C,{span:2},{default:t(()=>[a(Be,{type:"card"},{default:t(()=>[a(Oe,{key:"1",title:"\u5B8C\u6574\u63D0\u95EE"},{default:t(()=>[g(_)?(n(),m(w,{key:0,animation:!0},{default:t(()=>[a(b,{rows:3})]),_:1})):(n(),m(Q,{key:1},{default:t(()=>[i.value.messages?(n(),m(g(Te),{key:0,path:"res",data:i.value.messages,"show-length":!0},null,8,["data"])):(n(),v("span",Va,"\u65E0"))]),_:1}))]),_:1})]),_:1})]),_:1})]),_:1}),a(Ne,{layout:"inline-vertical",column:2,style:{"margin-top":"10px",position:"relative"}},{default:t(()=>[a(C,{span:2},{default:t(()=>[a(Be,{type:"card"},{default:t(()=>[a(Oe,{key:"1",title:"\u6A21\u578B\u4EE3\u7406\u4FE1\u606F"},{default:t(()=>[g(_)?(n(),m(w,{key:0,animation:!0},{default:t(()=>[a(b,{rows:3})]),_:1})):(n(),m(Q,{key:1},{default:t(()=>[i.value.model_agent?(n(),m(g(Te),{key:0,data:i.value.model_agent,"show-length":!0},null,8,["data"])):(n(),v("span",xa,"\u65E0"))]),_:1}))]),_:1})]),_:1})]),_:1})]),_:1})])]),_:1},8,["visible"])]),_:1},8,["title"])])}}});const Xa=Ye(Da,[["__scopeId","data-v-cf1535ab"]]);export{Xa as default};