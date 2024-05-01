import{_ as k,z}from"./index.36ee875b.js";/* empty css               *//* empty css                */import{d as w,e as y,B as i,aD as $,aG as a,aH as e,aL as c,aM as _,b1 as D,b2 as B,bK as R,aU as C,b5 as U,C as I,aJ as K,aI as E,aE as N,bB as J,bC as G,aT as H,aS as O,bi as A,bz as P,F as Q,bs as j,u as W,aK as X,aF as Y,bL as Z,bM as x,bJ as ee,bN as ae}from"./arco.aed15247.js";import{u as T}from"./loading.b5911e1d.js";import{d as te}from"./app.c5fa0b7b.js";/* empty css               *//* empty css               *//* empty css                *//* empty css               *//* empty css              *//* empty css              *//* empty css               */import{q as le}from"./model.c407852e.js";/* empty css               */import"./chart.9aa6eafa.js";import"./vue.0cc5b64a.js";import"./base.87fcf6e2.js";const oe=w({__name:"base-info",emits:["changeStep"],setup(V,{emit:b}){const d=y(),l=y({name:"",remark:""}),m=async()=>{var t;await((t=d.value)==null?void 0:t.validate())||b("changeStep","forward",{...l.value})};return(p,t)=>{const o=D,u=B,s=R,n=C,h=U;return i(),$(h,{ref_key:"formRef",ref:d,model:l.value,class:"form","label-col-props":{span:5},"wrapper-col-props":{span:18}},{default:a(()=>[e(u,{field:"name",label:p.$t("app.label.name"),rules:[{required:!0,message:p.$t("app.error.name.required")},{match:/^.{1,100}$/,message:p.$t("app.error.name.pattern")}]},{default:a(()=>[e(o,{modelValue:l.value.name,"onUpdate:modelValue":t[0]||(t[0]=f=>l.value.name=f),placeholder:p.$t("app.placeholder.name")},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),e(u,{field:"remark",label:p.$t("app.label.remark"),rules:[{required:!1}]},{default:a(()=>[e(s,{modelValue:l.value.remark,"onUpdate:modelValue":t[1]||(t[1]=f=>l.value.remark=f),placeholder:p.$t("app.placeholder.remark")},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),e(u,null,{default:a(()=>[e(n,{type:"primary",onClick:m},{default:a(()=>[c(_(p.$t("app.button.next")),1)]),_:1})]),_:1})]),_:1},8,["model"])}}});const se=k(oe,[["__scopeId","data-v-e5d3d922"]]),ne=w({__name:"advanced",emits:["changeStep"],setup(V,{emit:b}){const{setLoading:d}=T(!0),l=y([]);(async()=>{d(!0);try{const{data:s}=await le();l.value=s.items}catch{}finally{d(!1)}})();const p=y(),t=y({models:[],is_limit_quota:!1,quota:y(),ip_whitelist:"",ip_blacklist:""}),o=async()=>{var n;await((n=p.value)==null?void 0:n.validate())||b("changeStep","submit",{...t.value})},u=()=>{b("changeStep","backward")};return(s,n)=>{const h=J,f=G,v=B,q=H,S=O,g=R,L=C,F=A,M=U;return i(),$(M,{ref_key:"formRef",ref:p,model:t.value,class:"form","label-col-props":{span:4},"wrapper-col-props":{span:18}},{default:a(()=>[e(v,{field:"models",label:s.$t("app.label.models"),rules:[{required:!1}]},{default:a(()=>[e(f,{modelValue:t.value.models,"onUpdate:modelValue":n[0]||(n[0]=r=>t.value.models=r),placeholder:s.$t("app.placeholder.models"),"max-tag-count":3,multiple:"","allow-clear":""},{default:a(()=>[(i(!0),I(K,null,E(l.value,r=>(i(),$(h,{key:r.id,value:r.id,label:r.name},null,8,["value","label"]))),128))]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label"]),e(v,{field:"is_limit_quota",label:s.$t("app.label.isLimitQuota")},{default:a(()=>[e(q,{modelValue:t.value.is_limit_quota,"onUpdate:modelValue":n[1]||(n[1]=r=>t.value.is_limit_quota=r)},null,8,["modelValue"])]),_:1},8,["label"]),t.value.is_limit_quota?(i(),$(v,{key:0,field:"quota",label:s.$t("app.label.quota"),rules:[{required:!0,message:s.$t("app.error.quota.required")}]},{default:a(()=>[e(S,{modelValue:t.value.quota,"onUpdate:modelValue":n[2]||(n[2]=r=>t.value.quota=r),placeholder:s.$t("app.placeholder.quota"),precision:0,min:0,max:9999999999999},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"])):N("",!0),e(v,{field:"ip_whitelist",label:s.$t("app.label.ip_whitelist"),rules:[{required:!1}]},{default:a(()=>[e(g,{modelValue:t.value.ip_whitelist,"onUpdate:modelValue":n[3]||(n[3]=r=>t.value.ip_whitelist=r),placeholder:s.$t("app.placeholder.ip_whitelist"),"auto-size":{minRows:5,maxRows:10}},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),e(v,{field:"ip_blacklist",label:s.$t("app.label.ip_blacklist"),rules:[{required:!1}]},{default:a(()=>[e(g,{modelValue:t.value.ip_blacklist,"onUpdate:modelValue":n[4]||(n[4]=r=>t.value.ip_blacklist=r),placeholder:s.$t("app.placeholder.ip_blacklist"),"auto-size":{minRows:5,maxRows:10}},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),e(v,null,{default:a(()=>[e(F,null,{default:a(()=>[e(L,{type:"secondary",onClick:u},{default:a(()=>[c(_(s.$t("app.button.prev")),1)]),_:1}),e(L,{type:"primary",onClick:o},{default:a(()=>[c(_(s.$t("app.button.next")),1)]),_:1})]),_:1})]),_:1})]),_:1},8,["model"])}}});const pe=k(ne,[["__scopeId","data-v-ac6c04cd"]]),re={class:"success-wrap"},ue=w({__name:"success",emits:["changeStep"],setup(V,{emit:b}){const d=()=>{b("changeStep",1)};return(l,m)=>{const p=P,t=C,o=A;return i(),I("div",re,[e(p,{status:"success",title:l.$t("app.success.title"),subtitle:l.$t("app.success.create.subTitle")},null,8,["title","subtitle"]),e(o,{size:16},{default:a(()=>[e(t,{key:"finish",type:"secondary",onClick:m[0]||(m[0]=u=>l.$router.push({name:"AppList"}))},{default:a(()=>[c(_(l.$t("app.button.finish")),1)]),_:1}),e(t,{key:"again",type:"primary",onClick:d},{default:a(()=>[c(_(l.$t("app.button.again")),1)]),_:1})]),_:1})])}}});const ie=k(ue,[["__scopeId","data-v-9ee08270"]]),de={class:"container"},me={class:"wrapper"},ce={name:"AppCreate"},_e=w({...ce,setup(V){const{loading:b,setLoading:d}=T(!1),l=y(1),m=y({}),p=async()=>{d(!0);try{await te(m.value),l.value=3,m.value={}}catch{}finally{d(!1)}},t=(o,u)=>{if(typeof o=="number"){l.value=o;return}if(o==="forward"||o==="submit"){if(m.value={...m.value,...u},o==="submit"){p();return}l.value+=1}else o==="backward"&&(l.value-=1)};return(o,u)=>{const s=z,n=X,h=Y,f=Z,v=x,q=ee,S=ae;return i(),I("div",de,[e(h,{class:"container-breadcrumb"},{default:a(()=>[e(n,null,{default:a(()=>[e(s)]),_:1}),e(n,null,{default:a(()=>[c(_(o.$t("menu.app")),1)]),_:1}),e(n,null,{default:a(()=>[c(_(o.$t("menu.app.create")),1)]),_:1})]),_:1}),e(S,{loading:W(b),style:{width:"100%"}},{default:a(()=>[e(q,{class:"general-card",bordered:!1},{default:a(()=>[Q("div",me,[e(v,{current:l.value,"onUpdate:current":u[0]||(u[0]=g=>l.value=g),style:{width:"660px"},"line-less":"",class:"steps"},{default:a(()=>[e(f,{description:o.$t("app.subTitle.baseInfo")},{default:a(()=>[c(_(o.$t("app.title.baseInfo")),1)]),_:1},8,["description"]),e(f,{description:o.$t("app.subTitle.advanced")},{default:a(()=>[c(_(o.$t("app.title.advanced")),1)]),_:1},8,["description"]),e(f,{description:o.$t("app.subTitle.create.finish")},{default:a(()=>[c(_(o.$t("app.title.create.finish")),1)]),_:1},8,["description"])]),_:1},8,["current"]),(i(),$(j,null,[l.value===1?(i(),$(se,{key:0,onChangeStep:t})):l.value===2?(i(),$(pe,{key:1,onChangeStep:t})):l.value===3?(i(),$(ie,{key:2,onChangeStep:t})):N("",!0)],1024))])]),_:1})]),_:1},8,["loading"])])}}});const Ne=k(_e,[["__scopeId","data-v-90f2dd71"]]);export{Ne as default};