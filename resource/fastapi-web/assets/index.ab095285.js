import{_ as C,C as E}from"./index.3e0040d4.js";/* empty css               *//* empty css                */import{d as S,e as k,B as _,aD as g,aG as a,aH as e,aL as s,aM as f,bC as q,bA as D,b2 as F,bK as O,aU as I,b5 as U,C as w,aJ as M,aI as B,aT as z,bi as K,bz as G,F as j,bs as J,aE as H,u as P,aK as X,aF as Z,bL as Q,bM as W,bJ as Y,bN as x}from"./arco.47b3c23b.js";import{u as N}from"./loading.b238ab8e.js";import{c as ee}from"./key.33a08405.js";/* empty css               *//* empty css               *//* empty css                *//* empty css              *//* empty css              *//* empty css               *//* empty css               */import{q as ae}from"./model.0498f73b.js";import{f as te}from"./agent.22c91532.js";/* empty css               */import"./chart.49ffccb2.js";import"./vue.94924b34.js";import"./base.87fcf6e2.js";const le=S({__name:"base-info",emits:["changeStep"],setup(V,{emit:y}){const p=k(),l=k({corp:"",key:"",remark:""}),i=async()=>{var r;await((r=p.value)==null?void 0:r.validate())||y("changeStep","forward",{...l.value})};return(n,r)=>{const t=q,u=D,b=F,v=O,o=I,d=U;return _(),g(d,{ref_key:"formRef",ref:p,model:l.value,class:"form","label-col-props":{span:4},"wrapper-col-props":{span:18}},{default:a(()=>[e(b,{field:"corp",label:n.$t("key.label.corp"),rules:[{required:!0,message:n.$t("key.error.corp.required")}]},{default:a(()=>[e(u,{modelValue:l.value.corp,"onUpdate:modelValue":r[0]||(r[0]=m=>l.value.corp=m),placeholder:n.$t("key.placeholder.corp"),"allow-search":""},{default:a(()=>[e(t,{value:"OpenAI"},{default:a(()=>[s("OpenAI")]),_:1}),e(t,{value:"Baidu"},{default:a(()=>[s("\u767E\u5EA6")]),_:1}),e(t,{value:"Xfyun"},{default:a(()=>[s("\u79D1\u5927\u8BAF\u98DE")]),_:1}),e(t,{value:"Aliyun"},{default:a(()=>[s("\u963F\u91CC\u4E91")]),_:1}),e(t,{value:"ZhipuAI"},{default:a(()=>[s("\u667A\u8C31AI")]),_:1}),e(t,{value:"Google"},{default:a(()=>[s("Google")]),_:1}),e(t,{value:"DeepSeek"},{default:a(()=>[s("DeepSeek")]),_:1}),e(t,{value:"Midjourney"},{default:a(()=>[s("Midjourney")]),_:1})]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),e(b,{field:"key",label:n.$t("key.label.key"),rules:[{required:!0,message:n.$t("key.error.key.required")}]},{default:a(()=>[e(v,{modelValue:l.value.key,"onUpdate:modelValue":r[1]||(r[1]=m=>l.value.key=m),placeholder:n.$t("key.placeholder.key"),"auto-size":{minRows:5,maxRows:10}},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),e(b,{field:"remark",label:n.$t("key.label.remark"),rules:[{required:!1}]},{default:a(()=>[e(v,{modelValue:l.value.remark,"onUpdate:modelValue":r[2]||(r[2]=m=>l.value.remark=m),placeholder:n.$t("key.placeholder.remark")},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),e(b,null,{default:a(()=>[e(o,{type:"primary",onClick:i},{default:a(()=>[s(f(n.$t("key.button.next")),1)]),_:1})]),_:1})]),_:1},8,["model"])}}});const oe=C(le,[["__scopeId","data-v-3eabb631"]]),se=S({__name:"advanced",emits:["changeStep"],setup(V,{emit:y}){const{setLoading:p}=N(!0),l=k([]);(async()=>{p(!0);try{const{data:o}=await ae();l.value=o.items}catch{}finally{p(!1)}})();const n=k([]);(async()=>{p(!0);try{const{data:o}=await te();n.value=o.items}catch{}finally{p(!1)}})();const t=k(),u=k({models:[],model_agents:[],is_agents_only:!1}),b=async()=>{var d;await((d=t.value)==null?void 0:d.validate())||y("changeStep","submit",{...u.value})},v=()=>{y("changeStep","backward")};return(o,d)=>{const m=q,$=D,h=F,A=z,L=I,R=K,T=U;return _(),g(T,{ref_key:"formRef",ref:t,model:u.value,class:"form","label-col-props":{span:4},"wrapper-col-props":{span:18}},{default:a(()=>[e(h,{field:"models",label:o.$t("key.label.models"),rules:[{required:!1}]},{default:a(()=>[e($,{modelValue:u.value.models,"onUpdate:modelValue":d[0]||(d[0]=c=>u.value.models=c),placeholder:o.$t("key.placeholder.models"),"max-tag-count":3,multiple:"","allow-search":"","allow-clear":""},{default:a(()=>[(_(!0),w(M,null,B(l.value,c=>(_(),g(m,{key:c.id,value:c.id,label:c.name},null,8,["value","label"]))),128))]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label"]),e(h,{field:"model_agents",label:o.$t("key.label.modelAgents")},{default:a(()=>[e($,{modelValue:u.value.model_agents,"onUpdate:modelValue":d[1]||(d[1]=c=>u.value.model_agents=c),placeholder:o.$t("key.placeholder.modelAgents"),"max-tag-count":3,multiple:"","allow-search":"","allow-clear":""},{default:a(()=>[(_(!0),w(M,null,B(n.value,c=>(_(),g(m,{key:c.id,value:c.id,label:c.name},null,8,["value","label"]))),128))]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label"]),e(h,{field:"is_agents_only",label:o.$t("key.label.isAgentsOnly")},{default:a(()=>[e(A,{modelValue:u.value.is_agents_only,"onUpdate:modelValue":d[2]||(d[2]=c=>u.value.is_agents_only=c)},null,8,["modelValue"])]),_:1},8,["label"]),e(h,null,{default:a(()=>[e(R,null,{default:a(()=>[e(L,{type:"secondary",onClick:v},{default:a(()=>[s(f(o.$t("key.button.prev")),1)]),_:1}),e(L,{type:"primary",onClick:b},{default:a(()=>[s(f(o.$t("key.button.next")),1)]),_:1})]),_:1})]),_:1})]),_:1},8,["model"])}}});const ne=C(se,[["__scopeId","data-v-4abfd773"]]),re={class:"success-wrap"},ue=S({__name:"success",emits:["changeStep"],setup(V,{emit:y}){const p=()=>{y("changeStep",1)};return(l,i)=>{const n=G,r=I,t=K;return _(),w("div",re,[e(n,{status:"success",title:l.$t("key.success.title"),subtitle:l.$t("key.success.create.subTitle")},null,8,["title","subtitle"]),e(t,{size:16},{default:a(()=>[e(r,{key:"finish",type:"secondary",onClick:i[0]||(i[0]=u=>l.$router.push({name:"KeyModelList"}))},{default:a(()=>[s(f(l.$t("key.button.finish")),1)]),_:1}),e(r,{key:"again",type:"primary",onClick:p},{default:a(()=>[s(f(l.$t("key.button.again")),1)]),_:1})]),_:1})])}}});const de=C(ue,[["__scopeId","data-v-9a43fb75"]]),ce={class:"container"},pe={class:"wrapper"},_e={name:"KeyCreate"},ie=S({..._e,setup(V){const{loading:y,setLoading:p}=N(!1),l=k(1),i=k({}),n=async()=>{p(!0);try{await ee(i.value),l.value=3,i.value={}}catch{}finally{p(!1)}},r=(t,u)=>{if(typeof t=="number"){l.value=t;return}if(t==="forward"||t==="submit"){if(i.value={...i.value,...u},t==="submit"){n();return}l.value+=1}else t==="backward"&&(l.value-=1)};return(t,u)=>{const b=E,v=X,o=Z,d=Q,m=W,$=Y,h=x;return _(),w("div",ce,[e(o,{class:"container-breadcrumb"},{default:a(()=>[e(v,null,{default:a(()=>[e(b)]),_:1}),e(v,null,{default:a(()=>[s(f(t.$t("menu.key")),1)]),_:1}),e(v,null,{default:a(()=>[s(f(t.$t("menu.key.create")),1)]),_:1})]),_:1}),e(h,{loading:P(y),style:{width:"100%"}},{default:a(()=>[e($,{class:"general-card",bordered:!1},{default:a(()=>[j("div",pe,[e(m,{current:l.value,"onUpdate:current":u[0]||(u[0]=A=>l.value=A),style:{width:"660px"},"line-less":"",class:"steps"},{default:a(()=>[e(d,{description:t.$t("key.subTitle.baseInfo")},{default:a(()=>[s(f(t.$t("key.title.baseInfo")),1)]),_:1},8,["description"]),e(d,{description:t.$t("key.subTitle.advanced")},{default:a(()=>[s(f(t.$t("key.title.advanced")),1)]),_:1},8,["description"]),e(d,{description:t.$t("key.subTitle.create.finish")},{default:a(()=>[s(f(t.$t("key.title.create.finish")),1)]),_:1},8,["description"])]),_:1},8,["current"]),(_(),g(J,null,[l.value===1?(_(),g(oe,{key:0,onChangeStep:r})):l.value===2?(_(),g(ne,{key:1,onChangeStep:r})):l.value===3?(_(),g(de,{key:2,onChangeStep:r})):H("",!0)],1024))])]),_:1})]),_:1},8,["loading"])])}}});const De=C(ie,[["__scopeId","data-v-2be39207"]]);export{De as default};