import{_ as q,t as E}from"./index.74baba8a.js";import{u as S}from"./loading.dbaba456.js";/* empty css                */import{d as C,e as h,B as v,aD as k,aG as t,aH as a,aL as f,aM as b,b2 as O,b3 as U,bK as R,aV as A,b6 as B,C as I,aJ as G,aI as H,aE as N,bC as J,bD as P,aU as Q,aT as W,bj as T,bA as X,F as Y,bt as Z,u as x,aK as ee,aF as ae,bM as te,bN as le,bL as oe,bO as se}from"./arco.d2aaf5b7.js";import{d as F,e as ne}from"./app.097fee61.js";/* empty css               *//* empty css               *//* empty css                */import{h as M}from"./vue.ca65198a.js";/* empty css               *//* empty css              *//* empty css              *//* empty css               */import{q as pe}from"./model.5554004c.js";/* empty css               */import"./chart.61872c57.js";const ue=C({__name:"base-info",emits:["changeStep"],setup(y,{emit:_}){const{setLoading:d}=S(!1),s=M(),m=h(),r=h({id:"",name:"",remark:""});(async(n={id:s.query.id})=>{d(!0);try{const{data:p}=await F(n);r.value.id=p.id,r.value.name=p.name,r.value.remark=p.remark}catch{}finally{d(!1)}})();const e=async()=>{var p;await((p=m.value)==null?void 0:p.validate())||_("changeStep","forward",{...r.value})};return(n,p)=>{const w=O,l=U,o=R,$=A,u=B;return v(),k(u,{ref_key:"formRef",ref:m,model:r.value,class:"form","label-col-props":{span:6},"wrapper-col-props":{span:18}},{default:t(()=>[a(l,{field:"name",label:n.$t("app.label.name"),rules:[{required:!0,message:n.$t("app.error.name.required")},{match:/^.{1,20}$/,message:n.$t("app.error.name.pattern")}]},{default:t(()=>[a(w,{modelValue:r.value.name,"onUpdate:modelValue":p[0]||(p[0]=c=>r.value.name=c),placeholder:n.$t("app.placeholder.name")},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),a(l,{field:"remark",label:n.$t("app.label.remark"),rules:[{required:!1}]},{default:t(()=>[a(o,{modelValue:r.value.remark,"onUpdate:modelValue":p[1]||(p[1]=c=>r.value.remark=c),placeholder:n.$t("app.placeholder.remark")},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),a(l,null,{default:t(()=>[a($,{type:"primary",onClick:e},{default:t(()=>[f(b(n.$t("app.button.next")),1)]),_:1})]),_:1})]),_:1},8,["model"])}}});const re=q(ue,[["__scopeId","data-v-26c8744b"]]),ie=C({__name:"advanced",emits:["changeStep"],setup(y,{emit:_}){const{setLoading:d}=S(!0),s=M(),m=h([]);(async()=>{d(!0);try{const{data:l}=await pe();m.value=l.items}catch{}finally{d(!1)}})();const g=h(),e=h({models:[],is_limit_quota:!1,quota:h(),ip_whitelist:"",ip_blacklist:""});(async(l={id:s.query.id})=>{var o,$;d(!0);try{const{data:u}=await F(l);e.value.models=u.models,e.value.is_limit_quota=u.is_limit_quota,e.value.quota=u.quota,e.value.ip_whitelist=((o=u==null?void 0:u.ip_whitelist)==null?void 0:o.join(`
`))||"",e.value.ip_blacklist=(($=u==null?void 0:u.ip_blacklist)==null?void 0:$.join(`
`))||""}catch{}finally{d(!1)}})();const p=async()=>{var o;await((o=g.value)==null?void 0:o.validate())||_("changeStep","submit",{...e.value})},w=()=>{_("changeStep","backward")};return(l,o)=>{const $=J,u=P,c=U,V=Q,j=W,L=R,D=A,z=T,K=B;return v(),k(K,{ref_key:"formRef",ref:g,model:e.value,class:"form","label-col-props":{span:6},"wrapper-col-props":{span:18}},{default:t(()=>[a(c,{field:"models",label:l.$t("app.label.models"),rules:[{required:!1}]},{default:t(()=>[a(u,{modelValue:e.value.models,"onUpdate:modelValue":o[0]||(o[0]=i=>e.value.models=i),placeholder:l.$t("app.placeholder.models"),"max-tag-count":3,multiple:"","allow-clear":""},{default:t(()=>[(v(!0),I(G,null,H(m.value,i=>(v(),k($,{key:i.id,value:i.id,label:i.name},null,8,["value","label"]))),128))]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label"]),a(c,{field:"is_limit_quota",label:l.$t("app.label.isLimitQuota")},{default:t(()=>[a(V,{modelValue:e.value.is_limit_quota,"onUpdate:modelValue":o[1]||(o[1]=i=>e.value.is_limit_quota=i)},null,8,["modelValue"])]),_:1},8,["label"]),e.value.is_limit_quota?(v(),k(c,{key:0,field:"quota",label:l.$t("app.label.quota"),rules:[{required:!0,message:l.$t("app.error.quota.required")}]},{default:t(()=>[a(j,{modelValue:e.value.quota,"onUpdate:modelValue":o[2]||(o[2]=i=>e.value.quota=i),placeholder:l.$t("app.placeholder.quota"),min:1},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"])):N("",!0),a(c,{field:"ip_whitelist",label:l.$t("app.label.ip_whitelist"),rules:[{required:!1}]},{default:t(()=>[a(L,{modelValue:e.value.ip_whitelist,"onUpdate:modelValue":o[3]||(o[3]=i=>e.value.ip_whitelist=i),placeholder:l.$t("app.placeholder.ip_whitelist"),"auto-size":{minRows:5,maxRows:10}},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),a(c,{field:"ip_blacklist",label:l.$t("app.label.ip_blacklist"),rules:[{required:!1}]},{default:t(()=>[a(L,{modelValue:e.value.ip_blacklist,"onUpdate:modelValue":o[4]||(o[4]=i=>e.value.ip_blacklist=i),placeholder:l.$t("app.placeholder.ip_blacklist"),"auto-size":{minRows:5,maxRows:10}},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),a(c,null,{default:t(()=>[a(z,null,{default:t(()=>[a(D,{type:"secondary",onClick:w},{default:t(()=>[f(b(l.$t("model.button.prev")),1)]),_:1}),a(D,{type:"primary",onClick:p},{default:t(()=>[f(b(l.$t("model.button.next")),1)]),_:1})]),_:1})]),_:1})]),_:1},8,["model"])}}});const de=q(ie,[["__scopeId","data-v-bc1c9ada"]]);const me={},ce={class:"success-wrap"};function _e(y,_){const d=X,s=A,m=T;return v(),I("div",ce,[a(d,{status:"success",title:y.$t("app.success.title"),subtitle:y.$t("app.success.update.subTitle")},null,8,["title","subtitle"]),a(m,{size:16},{default:t(()=>[a(s,{key:"finish",type:"secondary",onClick:_[0]||(_[0]=r=>y.$router.push({name:"AppList"}))},{default:t(()=>[f(b(y.$t("app.button.return")),1)]),_:1}),a(s,{key:"again",type:"primary",onClick:_[1]||(_[1]=r=>y.$router.push({name:"AppDetail",query:{id:`${y.$route.query.id}`}}))},{default:t(()=>[f(b(y.$t("app.button.view")),1)]),_:1})]),_:1})])}const fe=q(me,[["render",_e],["__scopeId","data-v-e283e987"]]),be={class:"container"},ve={class:"wrapper"},ye={name:"AppUpdate"},$e=C({...ye,setup(y){const{loading:_,setLoading:d}=S(!1),s=h(1),m=h({}),r=async()=>{d(!0);try{await ne(m.value),s.value=3,m.value={}}catch{}finally{d(!1)}},g=(e,n)=>{if(typeof e=="number"){s.value=e;return}if(e==="forward"||e==="submit"){if(m.value={...m.value,...n},e==="submit"){r();return}s.value+=1}else e==="backward"&&(s.value-=1)};return(e,n)=>{const p=E,w=ee,l=ae,o=te,$=le,u=oe,c=se;return v(),I("div",be,[a(l,{class:"container-breadcrumb"},{default:t(()=>[a(w,null,{default:t(()=>[a(p)]),_:1}),a(w,null,{default:t(()=>[f(b(e.$t("menu.app")),1)]),_:1}),a(w,null,{default:t(()=>[f(b(e.$t("menu.app.update")),1)]),_:1})]),_:1}),a(c,{loading:x(_),style:{width:"100%"}},{default:t(()=>[a(u,{class:"general-card",bordered:!1},{title:t(()=>[f(b(e.$t("app.title.update")),1)]),default:t(()=>[Y("div",ve,[a($,{current:s.value,"onUpdate:current":n[0]||(n[0]=V=>s.value=V),style:{width:"580px"},"line-less":"",class:"steps"},{default:t(()=>[a(o,{description:e.$t("app.subTitle.baseInfo")},{default:t(()=>[f(b(e.$t("app.title.baseInfo")),1)]),_:1},8,["description"]),a(o,{description:e.$t("app.subTitle.advanced")},{default:t(()=>[f(b(e.$t("app.title.advanced")),1)]),_:1},8,["description"]),a(o,{description:e.$t("app.subTitle.update.finish")},{default:t(()=>[f(b(e.$t("app.title.update.finish")),1)]),_:1},8,["description"])]),_:1},8,["current"]),(v(),k(Z,null,[s.value===1?(v(),k(re,{key:0,onChangeStep:g})):s.value===2?(v(),k(de,{key:1,onChangeStep:g})):s.value===3?(v(),k(fe,{key:2,onChangeStep:g})):N("",!0)],1024))])]),_:1})]),_:1},8,["loading"])])}}});const Te=q($e,[["__scopeId","data-v-c37fe2dc"]]);export{Te as default};
