import{u as G,C as H,_ as J}from"./index.b24b4575.js";/* empty css                *//* empty css                *//* empty css               *//* empty css               *//* empty css                *//* empty css               *//* empty css               *//* empty css               *//* empty css               *//* empty css                *//* empty css              *//* empty css                *//* empty css               *//* empty css                */import{d as X,e as d,B as _,C as y,aH as a,aG as o,aL as i,aM as c,F as L,aJ as M,aI as I,aD as A,u as j,aK as Q,aF as W,bE as Y,bA as Z,bB as x,b2 as ee,bO as ae,aS as le,bP as te,aT as oe,aU as re,bi as ne,b5 as se,bI as de,bR as ue,g as me}from"./arco.a9260898.js";import{u as ie}from"./loading.1f346a94.js";import{f as ce}from"./vue.ad52ddbe.js";import{d as pe}from"./key.bf048eb9.js";import{q as _e}from"./corp.4e269a74.js";import{b as fe}from"./model.13b6297c.js";import{f as ye}from"./agent.53baf448.js";import"./chart.d103b168.js";const be={class:"container"},ke={class:"wrapper"},ge={class:"submit-btn"},ve={name:"KeyCreate"},he=X({...ve,setup(we){const{loading:S,setLoading:n}=ie(!1),{proxy:q}=me(),F=ce(),{t:u}=G(),p=d([]),b=new Map;(async()=>{n(!0);try{const{data:e}=await _e();p.value=e.items;for(let l=0;l<p.value.length;l+=1)b.set(p.value[l].id,p.value[l])}catch{}finally{n(!1)}})();const m=d(u("key.placeholder.key")),U=async()=>{switch(b.get(r.value.corp).code){case"Baidu":m.value=u("key.placeholder.key.baidu");break;case"Xfyun":m.value=u("key.placeholder.key.xfyun");break;case"DeepSeek-Baidu":m.value=u("key.placeholder.key.deepseek.baidu");break;case"VolcEngine":m.value=u("key.placeholder.key.volcengine");break;default:m.value=u("key.placeholder.key")}},k=d([]);(async()=>{n(!0);try{const{data:e}=await fe();k.value=e.items}catch{}finally{n(!1)}})();const g=d([]);(async()=>{n(!0);try{const{data:e}=await ye();g.value=e.items}catch{}finally{n(!1)}})();const v=d(),r=d({corp:"",key:"",weight:d(20),remark:"",models:[],model_agents:[],is_agents_only:!1}),T=async()=>{var l;if(!await((l=v.value)==null?void 0:l.validate())){n(!0);try{await pe(r.value).then(()=>{q.$message.success("\u65B0\u5EFA\u6210\u529F"),F.push({name:"ModelKeyList"})})}catch{}finally{n(!1)}}};return(e,l)=>{const h=H,f=Q,D=W,w=Y,V=Z,$=x,s=ee,C=ae,K=le,R=te,E=oe,B=re,N=ne,O=se,P=de,z=ue;return _(),y("div",be,[a(D,{class:"container-breadcrumb"},{default:o(()=>[a(f,null,{default:o(()=>[a(h)]),_:1}),a(f,null,{default:o(()=>[i(c(e.$t("menu.key")),1)]),_:1}),a(f,null,{default:o(()=>[i(c(e.$t("menu.key.create")),1)]),_:1})]),_:1}),a(z,{loading:j(S),style:{width:"100%"}},{default:o(()=>[a(P,{class:"general-card","body-style":{padding:"0 20px 20px 20px"},bordered:!1},{default:o(()=>[L("div",ke,[a(O,{ref_key:"formRef",ref:v,model:r.value,class:"form","label-col-props":{span:5},"wrapper-col-props":{span:18}},{default:o(()=>[a(w,{orientation:"left"},{default:o(()=>[i(c(e.$t("common.title.baseInfo")),1)]),_:1}),a(s,{field:"corp",label:e.$t("key.label.corp"),rules:[{required:!0,message:e.$t("key.error.corp.required")}]},{default:o(()=>[a($,{modelValue:r.value.corp,"onUpdate:modelValue":l[0]||(l[0]=t=>r.value.corp=t),placeholder:e.$t("key.placeholder.corp"),"allow-search":"",onChange:U},{default:o(()=>[(_(!0),y(M,null,I(p.value,t=>(_(),A(V,{key:t.id,value:t.id,label:t.name},null,8,["value","label"]))),128))]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),a(s,{field:"key",label:e.$t("key.label.key"),rules:[{required:!0,message:e.$t("key.error.key.required")}]},{default:o(()=>[a(C,{modelValue:r.value.key,"onUpdate:modelValue":l[1]||(l[1]=t=>r.value.key=t),placeholder:m.value,"auto-size":{minRows:5,maxRows:10}},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),a(s,{field:"weight",label:e.$t("model.agent.label.weight")},{default:o(()=>[a(K,{modelValue:r.value.weight,"onUpdate:modelValue":l[2]||(l[2]=t=>r.value.weight=t),precision:0,min:1,max:100,placeholder:e.$t("model.agent.placeholder.weight")},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),a(s,{field:"remark",label:e.$t("key.label.remark")},{default:o(()=>[a(C,{modelValue:r.value.remark,"onUpdate:modelValue":l[3]||(l[3]=t=>r.value.remark=t),placeholder:e.$t("key.placeholder.remark")},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),a(w,{orientation:"left"},{default:o(()=>[i(c(e.$t("common.title.advanced")),1)]),_:1}),a(s,{field:"models",label:e.$t("key.label.models")},{default:o(()=>[a(R,{modelValue:r.value.models,"onUpdate:modelValue":l[4]||(l[4]=t=>r.value.models=t),"allow-search":!0,"allow-clear":!0,"tree-checkable":!0,"tree-checked-strategy":"child",data:k.value,placeholder:e.$t("key.placeholder.models"),"max-tag-count":3},null,8,["modelValue","data","placeholder"])]),_:1},8,["label"]),a(s,{field:"model_agents",label:e.$t("key.label.modelAgents")},{default:o(()=>[a($,{modelValue:r.value.model_agents,"onUpdate:modelValue":l[5]||(l[5]=t=>r.value.model_agents=t),placeholder:e.$t("key.placeholder.modelAgents"),"max-tag-count":3,multiple:"","allow-search":"","allow-clear":""},{default:o(()=>[(_(!0),y(M,null,I(g.value,t=>(_(),A(V,{key:t.id,value:t.id,label:t.name},null,8,["value","label"]))),128))]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label"]),a(s,{field:"is_agents_only",label:e.$t("key.label.isAgentsOnly")},{default:o(()=>[a(E,{modelValue:r.value.is_agents_only,"onUpdate:modelValue":l[6]||(l[6]=t=>r.value.is_agents_only=t)},null,8,["modelValue"])]),_:1},8,["label"]),a(N,null,{default:o(()=>[L("div",ge,[a(B,{type:"secondary",onClick:l[7]||(l[7]=t=>e.$router.push({name:"ModelKeyList"}))},{default:o(()=>[i(c(e.$t("button.cancel")),1)]),_:1}),a(B,{type:"primary",onClick:T},{default:o(()=>[i(c(e.$t("button.submit")),1)]),_:1})])]),_:1})]),_:1},8,["model"])])]),_:1})]),_:1},8,["loading"])])}}});const Qe=J(he,[["__scopeId","data-v-ec308a8a"]]);export{Qe as default};
