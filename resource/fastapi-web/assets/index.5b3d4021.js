import{C as E,_ as O}from"./index.e0d6ab41.js";/* empty css                *//* empty css                *//* empty css               *//* empty css               *//* empty css                *//* empty css               *//* empty css               *//* empty css               *//* empty css               *//* empty css                *//* empty css              *//* empty css                *//* empty css               *//* empty css                */import{d as z,e as d,B as i,C as c,aH as a,aG as t,aL as u,aM as m,F as $,aJ as C,aI as L,aD as B,u as G,aK as H,aF as J,bE as P,bA as j,bB as Q,b2 as W,bO as X,aS as Y,bP as Z,aT as x,aU as ee,bi as ae,b5 as le,bI as te,bR as oe,g as re}from"./arco.91d8d802.js";import{u as ne}from"./loading.d8a03711.js";import{f as se}from"./vue.90059513.js";import{d as de}from"./key.4a72cc3f.js";import{q as ue}from"./corp.7ce589cd.js";import{a as me}from"./model.45f17314.js";import{f as ie}from"./agent.2789dbaf.js";import"./chart.1c4d013e.js";const pe={class:"container"},ce={class:"wrapper"},_e={class:"submit-btn"},fe={name:"KeyCreate"},ye=z({...fe,setup(be){const{loading:A,setLoading:n}=ne(!1),{proxy:I}=re(),M=se(),_=d([]);(async()=>{n(!0);try{const{data:e}=await ue();_.value=e.items}catch{}finally{n(!1)}})();const f=d([]);(async()=>{n(!0);try{const{data:e}=await me();f.value=e.items}catch{}finally{n(!1)}})();const y=d([]);(async()=>{n(!0);try{const{data:e}=await ie();y.value=e.items}catch{}finally{n(!1)}})();const b=d(),r=d({corp:"",key:"",weight:d(20),remark:"",models:[],model_agents:[],is_agents_only:!1}),q=async()=>{var o;if(!await((o=b.value)==null?void 0:o.validate())){n(!0);try{await de(r.value).then(()=>{I.$message.success("\u65B0\u5EFA\u6210\u529F"),M.push({name:"ModelKeyList"})})}catch{}finally{n(!1)}}};return(e,o)=>{const g=E,p=H,F=J,k=P,v=j,h=Q,s=W,w=X,S=Y,U=Z,T=x,V=ee,R=ae,D=le,K=te,N=oe;return i(),c("div",pe,[a(F,{class:"container-breadcrumb"},{default:t(()=>[a(p,null,{default:t(()=>[a(g)]),_:1}),a(p,null,{default:t(()=>[u(m(e.$t("menu.key")),1)]),_:1}),a(p,null,{default:t(()=>[u(m(e.$t("menu.key.create")),1)]),_:1})]),_:1}),a(N,{loading:G(A),style:{width:"100%"}},{default:t(()=>[a(K,{class:"general-card","body-style":{padding:"0 20px 20px 20px"},bordered:!1},{default:t(()=>[$("div",ce,[a(D,{ref_key:"formRef",ref:b,model:r.value,class:"form","label-col-props":{span:5},"wrapper-col-props":{span:18}},{default:t(()=>[a(k,{orientation:"left"},{default:t(()=>[u(m(e.$t("common.title.baseInfo")),1)]),_:1}),a(s,{field:"corp",label:e.$t("key.label.corp"),rules:[{required:!0,message:e.$t("key.error.corp.required")}]},{default:t(()=>[a(h,{modelValue:r.value.corp,"onUpdate:modelValue":o[0]||(o[0]=l=>r.value.corp=l),placeholder:e.$t("key.placeholder.corp"),"allow-search":""},{default:t(()=>[(i(!0),c(C,null,L(_.value,l=>(i(),B(v,{key:l.id,value:l.id,label:l.name},null,8,["value","label"]))),128))]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),a(s,{field:"key",label:e.$t("key.label.key"),rules:[{required:!0,message:e.$t("key.error.key.required")}]},{default:t(()=>[a(w,{modelValue:r.value.key,"onUpdate:modelValue":o[1]||(o[1]=l=>r.value.key=l),placeholder:e.$t("key.placeholder.key"),"auto-size":{minRows:5,maxRows:10}},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),a(s,{field:"weight",label:e.$t("model.agent.label.weight")},{default:t(()=>[a(S,{modelValue:r.value.weight,"onUpdate:modelValue":o[2]||(o[2]=l=>r.value.weight=l),precision:0,min:1,max:100,placeholder:e.$t("model.agent.placeholder.weight")},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),a(s,{field:"remark",label:e.$t("key.label.remark")},{default:t(()=>[a(w,{modelValue:r.value.remark,"onUpdate:modelValue":o[3]||(o[3]=l=>r.value.remark=l),placeholder:e.$t("key.placeholder.remark")},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),a(k,{orientation:"left"},{default:t(()=>[u(m(e.$t("common.title.advanced")),1)]),_:1}),a(s,{field:"models",label:e.$t("key.label.models")},{default:t(()=>[a(U,{modelValue:r.value.models,"onUpdate:modelValue":o[4]||(o[4]=l=>r.value.models=l),"allow-search":!0,"allow-clear":!0,"tree-checkable":!0,"tree-checked-strategy":"child",data:f.value,placeholder:e.$t("key.placeholder.models"),"max-tag-count":3},null,8,["modelValue","data","placeholder"])]),_:1},8,["label"]),a(s,{field:"model_agents",label:e.$t("key.label.modelAgents")},{default:t(()=>[a(h,{modelValue:r.value.model_agents,"onUpdate:modelValue":o[5]||(o[5]=l=>r.value.model_agents=l),placeholder:e.$t("key.placeholder.modelAgents"),"max-tag-count":3,multiple:"","allow-search":"","allow-clear":""},{default:t(()=>[(i(!0),c(C,null,L(y.value,l=>(i(),B(v,{key:l.id,value:l.id,label:l.name},null,8,["value","label"]))),128))]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label"]),a(s,{field:"is_agents_only",label:e.$t("key.label.isAgentsOnly")},{default:t(()=>[a(T,{modelValue:r.value.is_agents_only,"onUpdate:modelValue":o[6]||(o[6]=l=>r.value.is_agents_only=l)},null,8,["modelValue"])]),_:1},8,["label"]),a(R,null,{default:t(()=>[$("div",_e,[a(V,{type:"secondary",onClick:o[7]||(o[7]=l=>e.$router.push({name:"ModelKeyList"}))},{default:t(()=>[u(m(e.$t("button.cancel")),1)]),_:1}),a(V,{type:"primary",onClick:q},{default:t(()=>[u(m(e.$t("button.submit")),1)]),_:1})])]),_:1})]),_:1},8,["model"])])]),_:1})]),_:1},8,["loading"])])}}});const He=O(ye,[["__scopeId","data-v-095a9813"]]);export{He as default};
