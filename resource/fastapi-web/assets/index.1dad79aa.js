import{x as N,_ as O}from"./index.a8f8f038.js";/* empty css                *//* empty css                *//* empty css               *//* empty css               *//* empty css                *//* empty css               *//* empty css               *//* empty css               *//* empty css               *//* empty css                *//* empty css              *//* empty css                *//* empty css               *//* empty css                */import{d as E,e as m,B as c,C as $,aH as l,aG as r,aL as d,aM as u,F as V,aJ as z,aI as G,aD as H,u as J,aK as K,aF as P,bE as Q,bA as j,bB as W,b2 as X,b1 as Y,aS as Z,bO as x,bP as ee,aT as ae,aU as le,bi as te,b5 as oe,bI as re,bQ as ne,g as se}from"./arco.553b67be.js";import{u as de}from"./loading.9398a00f.js";import{h as ue,f as me}from"./vue.d21e79d3.js";import{e as pe,c as ie}from"./agent.102c4e9c.js";import{q as ce}from"./corp.da8cfa0b.js";import{a as _e}from"./model.78bc160f.js";import"./chart.05df75e9.js";import"./base.87fcf6e2.js";const be={class:"container"},fe={class:"wrapper"},ge={class:"submit-btn"},ve={name:"ModelAgentUpdate"},ye=E({...ve,setup(he){const{loading:k,setLoading:s}=de(!1),{proxy:w}=se(),U=ue(),q=me(),_=m([]);(async()=>{s(!0);try{const{data:e}=await ce();_.value=e.items}catch{}finally{s(!1)}})();const b=m([]);(async()=>{s(!0);try{const{data:e}=await _e();b.value=e.items}catch{}finally{s(!1)}})();const f=m(),t=m({id:"",corp:"",name:"",base_url:"",path:"",weight:m(),remark:"",status:1,models:[],key:"",is_agents_only:!0}),B=async()=>{var a;if(!await((a=f.value)==null?void 0:a.validate())){s(!0);try{await pe(t.value).then(()=>{w.$message.success("\u66F4\u65B0\u6210\u529F"),q.push({name:"ModelAgentList"})})}catch{}finally{s(!1)}}};return(async(e={id:U.query.id})=>{s(!0);try{const{data:a}=await ie(e);t.value.id=a.id,t.value.corp=a.corp,t.value.name=a.name,t.value.base_url=a.base_url,t.value.path=a.path,t.value.weight=a.weight,t.value.remark=a.remark,t.value.status=a.status,t.value.models=a.models,t.value.key=a.key}catch{}finally{s(!1)}})(),(e,a)=>{const g=N,p=K,M=P,v=Q,C=j,I=W,n=X,i=Y,A=Z,y=x,L=ee,D=ae,h=le,F=te,S=oe,T=re,R=ne;return c(),$("div",be,[l(M,{class:"container-breadcrumb"},{default:r(()=>[l(p,null,{default:r(()=>[l(g)]),_:1}),l(p,null,{default:r(()=>[d(u(e.$t("menu.agent")),1)]),_:1}),l(p,null,{default:r(()=>[d(u(e.$t("menu.model.agent.update")),1)]),_:1})]),_:1}),l(R,{loading:J(k),style:{width:"100%"}},{default:r(()=>[l(T,{class:"general-card","body-style":{padding:"0 20px 20px 20px"},bordered:!1},{default:r(()=>[V("div",fe,[l(S,{ref_key:"formRef",ref:f,model:t.value,class:"form","label-col-props":{span:5},"wrapper-col-props":{span:18}},{default:r(()=>[l(v,{orientation:"left"},{default:r(()=>[d(u(e.$t("model.title.baseInfo")),1)]),_:1}),l(n,{field:"corp",label:e.$t("model.agent.label.corp"),rules:[{required:!0,message:e.$t("model.agent.error.corp.required")}]},{default:r(()=>[l(I,{modelValue:t.value.corp,"onUpdate:modelValue":a[0]||(a[0]=o=>t.value.corp=o),placeholder:e.$t("model.agent.placeholder.corp"),"allow-search":""},{default:r(()=>[(c(!0),$(z,null,G(_.value,o=>(c(),H(C,{key:o.id,value:o.id,label:o.name},null,8,["value","label"]))),128))]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),l(n,{field:"name",label:e.$t("model.agent.label.name"),rules:[{required:!0,message:e.$t("model.agent.error.name.required")},{match:/^.{1,100}$/,message:e.$t("model.agent.error.name.pattern")}]},{default:r(()=>[l(i,{modelValue:t.value.name,"onUpdate:modelValue":a[1]||(a[1]=o=>t.value.name=o),placeholder:e.$t("model.agent.placeholder.name")},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),l(n,{field:"base_url",label:e.$t("model.agent.label.baseUrl"),rules:[{required:!0,message:e.$t("model.agent.error.baseUrl.required")}]},{default:r(()=>[l(i,{modelValue:t.value.base_url,"onUpdate:modelValue":a[2]||(a[2]=o=>t.value.base_url=o),placeholder:e.$t("model.agent.placeholder.baseUrl")},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),l(n,{field:"path",label:e.$t("model.agent.label.path")},{default:r(()=>[l(i,{modelValue:t.value.path,"onUpdate:modelValue":a[3]||(a[3]=o=>t.value.path=o),placeholder:e.$t("model.agent.placeholder.path")},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),l(n,{field:"weight",label:e.$t("model.agent.label.weight")},{default:r(()=>[l(A,{modelValue:t.value.weight,"onUpdate:modelValue":a[4]||(a[4]=o=>t.value.weight=o),precision:0,min:0,max:99999,placeholder:e.$t("model.agent.placeholder.weight")},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),l(n,{field:"remark",label:e.$t("model.agent.label.remark")},{default:r(()=>[l(y,{modelValue:t.value.remark,"onUpdate:modelValue":a[5]||(a[5]=o=>t.value.remark=o),placeholder:e.$t("model.agent.placeholder.remark")},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),l(v,{orientation:"left"},{default:r(()=>[d(u(e.$t("model.title.advanced")),1)]),_:1}),l(n,{field:"models",label:e.$t("model.agent.label.models"),rules:[{required:!1}]},{default:r(()=>[l(L,{modelValue:t.value.models,"onUpdate:modelValue":a[6]||(a[6]=o=>t.value.models=o),"allow-search":!0,"allow-clear":!0,"tree-checkable":!0,"tree-checked-strategy":"child",data:b.value,placeholder:e.$t("model.agent.placeholder.models"),"max-tag-count":3},null,8,["modelValue","data","placeholder"])]),_:1},8,["label"]),l(n,{field:"key",label:e.$t("model.agent.label.key")},{default:r(()=>[l(y,{modelValue:t.value.key,"onUpdate:modelValue":a[7]||(a[7]=o=>t.value.key=o),placeholder:e.$t("model.agent.placeholder.key"),"auto-size":{minRows:5,maxRows:10}},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),l(n,{field:"is_agents_only",label:e.$t("model.agent.label.isAgentsOnly")},{default:r(()=>[l(D,{modelValue:t.value.is_agents_only,"onUpdate:modelValue":a[8]||(a[8]=o=>t.value.is_agents_only=o)},null,8,["modelValue"])]),_:1},8,["label"]),l(F,null,{default:r(()=>[V("div",ge,[l(h,{type:"secondary",onClick:a[9]||(a[9]=o=>e.$router.push({name:"ModelAgentList"}))},{default:r(()=>[d(u(e.$t("button.cancel")),1)]),_:1}),l(h,{type:"primary",onClick:B},{default:r(()=>[d(u(e.$t("button.submit")),1)]),_:1})])]),_:1})]),_:1},8,["model"])])]),_:1})]),_:1},8,["loading"])])}}});const Qe=O(ye,[["__scopeId","data-v-7746a3f9"]]);export{Qe as default};