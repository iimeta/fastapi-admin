import{u as J,C as X,_ as j}from"./index.46908301.js";/* empty css                *//* empty css                *//* empty css               *//* empty css               *//* empty css                *//* empty css               *//* empty css               *//* empty css               *//* empty css               *//* empty css                *//* empty css              *//* empty css                *//* empty css               *//* empty css                */import{d as Q,e as u,B as _,C as f,aH as l,aG as r,aL as m,aM as c,F as L,aJ as I,aI as M,aD as q,u as W,aK as Y,aF as Z,bE as x,bA as ee,bB as ae,b2 as le,b1 as te,aS as oe,bO as re,bP as se,aT as ne,aU as ue,bi as de,b5 as ie,bI as me,bR as ce,g as pe}from"./arco.a9260898.js";import{u as _e}from"./loading.1f346a94.js";import{f as ye,h as fe}from"./vue.ad52ddbe.js";import{e as be,c as ge}from"./key.ba43d583.js";import{q as ke}from"./corp.b46235a2.js";import{b as ve}from"./model.2294bab6.js";import{f as he}from"./agent.f3d50fc8.js";import"./chart.d103b168.js";const we={class:"container"},Ve={class:"wrapper"},$e={class:"submit-btn"},Be={name:"KeyUpdate"},Ce=Q({...Be,setup(Le){const{loading:U,setLoading:s}=_e(!1),{proxy:D}=pe(),K=ye(),S=fe(),{t:d}=J(),p=u([]),b=new Map;(async()=>{s(!0);try{const{data:a}=await ke();p.value=a.items;for(let e=0;e<p.value.length;e+=1)b.set(p.value[e].id,p.value[e])}catch{}finally{s(!1)}})();const i=u(d("key.placeholder.update.key")),g=async()=>{switch(b.get(t.value.corp).code){case"Baidu":i.value=d("key.placeholder.update.key.baidu");break;case"Xfyun":i.value=d("key.placeholder.update.key.xfyun");break;case"DeepSeek-Baidu":i.value=d("key.placeholder.update.key.deepseek.baidu");break;case"VolcEngine":i.value=d("key.placeholder.update.key.volcengine");break;default:i.value=d("key.placeholder.update.key")}},k=u([]);(async()=>{s(!0);try{const{data:a}=await ve();k.value=a.items}catch{}finally{s(!1)}})();const v=u([]);(async()=>{s(!0);try{const{data:a}=await he();v.value=a.items}catch{}finally{s(!1)}})();const h=u(),t=u({id:"",corp:"",key:"",weight:u(20),remark:"",status:1,models:[],model_agents:[],is_agents_only:!1}),A=async()=>{var e;if(!await((e=h.value)==null?void 0:e.validate())){s(!0);try{await be(t.value).then(()=>{D.$message.success("\u66F4\u65B0\u6210\u529F"),K.push({name:"ModelKeyList"})})}catch{}finally{s(!1)}}};return(async(a={id:S.query.id})=>{s(!0);try{const{data:e}=await ge(a);t.value.id=e.id,t.value.corp=e.corp,t.value.key=e.key,t.value.weight=e.weight,t.value.remark=e.remark,t.value.status=e.status,t.value.models=e.models,t.value.model_agents=e.model_agents,t.value.is_agents_only=e.is_agents_only,g()}catch{}finally{s(!1)}})(),(a,e)=>{const w=X,y=Y,F=Z,V=x,$=ee,B=ae,n=le,R=te,T=oe,N=re,E=se,O=ne,C=ue,P=de,z=ie,G=me,H=ce;return _(),f("div",we,[l(F,{class:"container-breadcrumb"},{default:r(()=>[l(y,null,{default:r(()=>[l(w)]),_:1}),l(y,null,{default:r(()=>[m(c(a.$t("menu.key")),1)]),_:1}),l(y,null,{default:r(()=>[m(c(a.$t("menu.key.update")),1)]),_:1})]),_:1}),l(H,{loading:W(U),style:{width:"100%"}},{default:r(()=>[l(G,{class:"general-card","body-style":{padding:"0 20px 20px 20px"},bordered:!1},{default:r(()=>[L("div",Ve,[l(z,{ref_key:"formRef",ref:h,model:t.value,class:"form","label-col-props":{span:5},"wrapper-col-props":{span:18}},{default:r(()=>[l(V,{orientation:"left"},{default:r(()=>[m(c(a.$t("common.title.baseInfo")),1)]),_:1}),l(n,{field:"corp",label:a.$t("key.label.corp"),rules:[{required:!0,message:a.$t("key.error.corp.required")}]},{default:r(()=>[l(B,{modelValue:t.value.corp,"onUpdate:modelValue":e[0]||(e[0]=o=>t.value.corp=o),placeholder:a.$t("key.placeholder.corp"),"allow-search":"",onChange:g},{default:r(()=>[(_(!0),f(I,null,M(p.value,o=>(_(),q($,{key:o.id,value:o.id,label:o.name},null,8,["value","label"]))),128))]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),l(n,{field:"key",label:a.$t("key.label.key"),rules:[{required:!0,message:a.$t("key.error.key.required")}]},{default:r(()=>[l(R,{modelValue:t.value.key,"onUpdate:modelValue":e[1]||(e[1]=o=>t.value.key=o),placeholder:i.value,"auto-size":{minRows:5,maxRows:10},"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),l(n,{field:"weight",label:a.$t("model.agent.label.weight")},{default:r(()=>[l(T,{modelValue:t.value.weight,"onUpdate:modelValue":e[2]||(e[2]=o=>t.value.weight=o),precision:0,min:0,max:999,placeholder:a.$t("model.agent.placeholder.weight")},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),l(n,{field:"remark",label:a.$t("key.label.remark")},{default:r(()=>[l(N,{modelValue:t.value.remark,"onUpdate:modelValue":e[3]||(e[3]=o=>t.value.remark=o),placeholder:a.$t("key.placeholder.remark")},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),l(V,{orientation:"left"},{default:r(()=>[m(c(a.$t("common.title.advanced")),1)]),_:1}),l(n,{field:"models",label:a.$t("key.label.models")},{default:r(()=>[l(E,{modelValue:t.value.models,"onUpdate:modelValue":e[4]||(e[4]=o=>t.value.models=o),"allow-search":!0,"allow-clear":!0,"tree-checkable":!0,"tree-checked-strategy":"child",data:k.value,placeholder:a.$t("key.placeholder.models"),"max-tag-count":3},null,8,["modelValue","data","placeholder"])]),_:1},8,["label"]),l(n,{field:"model_agents",label:a.$t("key.label.modelAgents")},{default:r(()=>[l(B,{modelValue:t.value.model_agents,"onUpdate:modelValue":e[5]||(e[5]=o=>t.value.model_agents=o),placeholder:a.$t("key.placeholder.modelAgents"),"max-tag-count":3,multiple:"","allow-search":"","allow-clear":""},{default:r(()=>[(_(!0),f(I,null,M(v.value,o=>(_(),q($,{key:o.id,value:o.id,label:o.name},null,8,["value","label"]))),128))]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label"]),l(n,{field:"is_agents_only",label:a.$t("key.label.isAgentsOnly")},{default:r(()=>[l(O,{modelValue:t.value.is_agents_only,"onUpdate:modelValue":e[6]||(e[6]=o=>t.value.is_agents_only=o)},null,8,["modelValue"])]),_:1},8,["label"]),l(P,null,{default:r(()=>[L("div",$e,[l(C,{type:"secondary",onClick:e[7]||(e[7]=o=>a.$router.push({name:"ModelKeyList"}))},{default:r(()=>[m(c(a.$t("button.cancel")),1)]),_:1}),l(C,{type:"primary",onClick:A},{default:r(()=>[m(c(a.$t("button.submit")),1)]),_:1})])]),_:1})]),_:1},8,["model"])])]),_:1})]),_:1},8,["loading"])])}}});const aa=j(Ce,[["__scopeId","data-v-bccd934c"]]);export{aa as default};
