import{u as Z,x,_ as ee}from"./index.d8cab1b0.js";/* empty css                *//* empty css                *//* empty css              *//* empty css               *//* empty css                *//* empty css               *//* empty css               *//* empty css               *//* empty css               *//* empty css                *//* empty css              *//* empty css                *//* empty css               *//* empty css                */import{d as le,e as i,B as g,C as h,aH as l,aG as t,aL as s,aM as c,F as R,aJ as D,aI as F,aD as L,bu as ae,bv as oe,u as te,aK as re,aF as de,bE as ne,bA as se,bB as ue,b2 as me,b1 as pe,aS as ie,bO as ce,bP as _e,aT as be,a5 as ge,aU as fe,a4 as ve,bQ as ye,bi as he,b5 as Ve,bI as ke,bR as $e,g as we}from"./arco.a9260898.js";import{u as Ue}from"./loading.1f346a94.js";import{f as Ce}from"./vue.ad52ddbe.js";import{d as Be}from"./agent.f34b296b.js";import{q as Ie}from"./corp.a5274466.js";import{b as Me}from"./model.c9c2685b.js";import"./chart.d103b168.js";const qe={class:"container"},Re={class:"wrapper"},De={class:"submit-btn"},Fe={name:"ModelAgentCreate"},Le=le({...Fe,setup(Se){const{loading:S,setLoading:u}=Ue(!1),{proxy:A}=we(),E=Ce(),{t:m}=Z(),_=i([]),V=new Map;(async()=>{u(!0);try{const{data:e}=await Ie();_.value=e.items;for(let o=0;o<_.value.length;o+=1)V.set(_.value[o].id,_.value[o])}catch{}finally{u(!1)}})();const p=i(m("key.placeholder.key")),T=async()=>{switch(V.get(a.value.corp).code){case"Baidu":p.value=m("key.placeholder.key.baidu");break;case"Xfyun":p.value=m("key.placeholder.key.xfyun");break;case"DeepSeek-Baidu":p.value=m("key.placeholder.key.deepseek.baidu");break;case"VolcEngine":p.value=m("key.placeholder.key.volcengine");break;default:p.value=m("key.placeholder.key")}},k=i([]);(async()=>{u(!0);try{const{data:e}=await Me();k.value=e.items}catch{}finally{u(!1)}})();const $=i(),a=i({corp:"",name:"",base_url:"",path:"",weight:i(20),remark:"",models:[],is_enable_model_replace:!1,replace_models:[],target_models:[],lb_strategy:"1",key:"",is_agents_only:!0}),N=async()=>{var o;if(!await((o=$.value)==null?void 0:o.validate())){u(!0);try{await Be(a.value).then(()=>{A.$message.success("\u65B0\u5EFA\u6210\u529F"),E.push({name:"ModelAgentList"})})}catch{}finally{u(!1)}}},P=()=>{a.value.is_enable_model_replace?w():(a.value.replace_models=[],a.value.target_models=[])},w=()=>{a.value.replace_models.push(""),a.value.target_models.push("")},O=e=>{a.value.replace_models.length>1&&(a.value.replace_models.splice(e,1),a.value.target_models.splice(e,1))};return(e,o)=>{const U=x,y=re,z=de,C=ne,K=se,G=ue,d=me,b=pe,H=ie,B=ce,J=_e,I=be,Q=ge,f=fe,X=ve,M=ye,q=he,j=Ve,W=ke,Y=$e;return g(),h("div",qe,[l(z,{class:"container-breadcrumb"},{default:t(()=>[l(y,null,{default:t(()=>[l(U)]),_:1}),l(y,null,{default:t(()=>[s(c(e.$t("menu.agent")),1)]),_:1}),l(y,null,{default:t(()=>[s(c(e.$t("menu.model.agent.create")),1)]),_:1})]),_:1}),l(Y,{loading:te(S),style:{width:"100%"}},{default:t(()=>[l(W,{class:"general-card","body-style":{padding:"0 20px 20px 20px"},bordered:!1},{default:t(()=>[R("div",Re,[l(j,{ref_key:"formRef",ref:$,model:a.value,class:"form","label-col-props":{span:5},"wrapper-col-props":{span:18}},{default:t(()=>[l(C,{orientation:"left"},{default:t(()=>[s(c(e.$t("common.title.baseInfo")),1)]),_:1}),l(d,{field:"corp",label:e.$t("model.agent.label.corp"),rules:[{required:!0,message:e.$t("model.agent.error.corp.required")}]},{default:t(()=>[l(G,{modelValue:a.value.corp,"onUpdate:modelValue":o[0]||(o[0]=r=>a.value.corp=r),placeholder:e.$t("model.agent.placeholder.corp"),"allow-search":"",onChange:T},{default:t(()=>[(g(!0),h(D,null,F(_.value,r=>(g(),L(K,{key:r.id,value:r.id,label:r.name},null,8,["value","label"]))),128))]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),l(d,{field:"name",label:e.$t("model.agent.label.name"),rules:[{required:!0,message:e.$t("model.agent.error.name.required")},{match:/^.{1,100}$/,message:e.$t("model.agent.error.name.pattern")}]},{default:t(()=>[l(b,{modelValue:a.value.name,"onUpdate:modelValue":o[1]||(o[1]=r=>a.value.name=r),placeholder:e.$t("model.agent.placeholder.name")},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),l(d,{field:"base_url",label:e.$t("model.agent.label.baseUrl"),rules:[{required:!0,message:e.$t("model.agent.error.baseUrl.required")}]},{default:t(()=>[l(b,{modelValue:a.value.base_url,"onUpdate:modelValue":o[2]||(o[2]=r=>a.value.base_url=r),placeholder:e.$t("model.agent.placeholder.baseUrl")},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),l(d,{field:"path",label:e.$t("model.agent.label.path")},{default:t(()=>[l(b,{modelValue:a.value.path,"onUpdate:modelValue":o[3]||(o[3]=r=>a.value.path=r),placeholder:e.$t("model.agent.placeholder.path")},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),l(d,{field:"weight",label:e.$t("model.agent.label.weight")},{default:t(()=>[l(H,{modelValue:a.value.weight,"onUpdate:modelValue":o[4]||(o[4]=r=>a.value.weight=r),precision:0,min:1,max:100,placeholder:e.$t("model.agent.placeholder.weight")},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),l(d,{field:"remark",label:e.$t("model.agent.label.remark")},{default:t(()=>[l(B,{modelValue:a.value.remark,"onUpdate:modelValue":o[5]||(o[5]=r=>a.value.remark=r),placeholder:e.$t("model.agent.placeholder.remark")},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),l(C,{orientation:"left"},{default:t(()=>[s(c(e.$t("common.title.advanced")),1)]),_:1}),l(d,{field:"models",label:e.$t("model.agent.label.models")},{default:t(()=>[l(J,{modelValue:a.value.models,"onUpdate:modelValue":o[6]||(o[6]=r=>a.value.models=r),"allow-search":!0,"allow-clear":!0,"tree-checkable":!0,"tree-checked-strategy":"child",data:k.value,placeholder:e.$t("model.agent.placeholder.models"),"max-tag-count":3},null,8,["modelValue","data","placeholder"])]),_:1},8,["label"]),l(d,{field:"is_enable_model_replace",label:e.$t("model.agent.label.is_enable_model_replace")},{default:t(()=>[l(I,{modelValue:a.value.is_enable_model_replace,"onUpdate:modelValue":o[7]||(o[7]=r=>a.value.is_enable_model_replace=r),onChange:P},null,8,["modelValue"])]),_:1},8,["label"]),(g(!0),h(D,null,F(a.value.replace_models,(r,n)=>ae((g(),L(d,{key:n,field:`replace_models[${n}]`&&`target_models[${n}]`,label:`${n+1}. `+e.$t("model.agent.label.replace_models"),rules:[{required:!0,message:e.$t("model.agent.error.replace_models.required")}]},{default:t(()=>[l(b,{modelValue:a.value.replace_models[n],"onUpdate:modelValue":v=>a.value.replace_models[n]=v,placeholder:e.$t("model.agent.placeholder.replace_models"),style:{width:"218px","margin-right":"5px"}},null,8,["modelValue","onUpdate:modelValue","placeholder"]),l(b,{modelValue:a.value.target_models[n],"onUpdate:modelValue":v=>a.value.target_models[n]=v,placeholder:e.$t("model.agent.placeholder.target_models"),style:{width:"218px"}},null,8,["modelValue","onUpdate:modelValue","placeholder"]),l(f,{type:"primary",shape:"circle",style:{margin:"0 10px 0 10px"},onClick:w},{default:t(()=>[l(Q)]),_:1}),l(f,{type:"secondary",shape:"circle",onClick:v=>O(n)},{default:t(()=>[l(X)]),_:2},1032,["onClick"])]),_:2},1032,["field","label","rules"])),[[oe,a.value.is_enable_model_replace]])),128)),l(d,{field:"lb_strategy",label:e.$t("model.agent.label.lb_strategy")},{default:t(()=>[l(q,{size:"large"},{default:t(()=>[l(M,{modelValue:a.value.lb_strategy,"onUpdate:modelValue":o[8]||(o[8]=r=>a.value.lb_strategy=r),value:"1","default-checked":!0},{default:t(()=>[s(" \u8F6E\u8BE2 ")]),_:1},8,["modelValue"]),l(M,{modelValue:a.value.lb_strategy,"onUpdate:modelValue":o[9]||(o[9]=r=>a.value.lb_strategy=r),value:"2"},{default:t(()=>[s("\u6743\u91CD")]),_:1},8,["modelValue"])]),_:1})]),_:1},8,["label"]),l(d,{field:"key",label:e.$t("model.agent.label.key")},{default:t(()=>[l(B,{modelValue:a.value.key,"onUpdate:modelValue":o[10]||(o[10]=r=>a.value.key=r),placeholder:p.value,"auto-size":{minRows:5,maxRows:10}},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),l(d,{field:"is_agents_only",label:e.$t("model.agent.label.isAgentsOnly")},{default:t(()=>[l(I,{modelValue:a.value.is_agents_only,"onUpdate:modelValue":o[11]||(o[11]=r=>a.value.is_agents_only=r)},null,8,["modelValue"])]),_:1},8,["label"]),l(q,null,{default:t(()=>[R("div",De,[l(f,{type:"secondary",onClick:o[12]||(o[12]=r=>e.$router.push({name:"ModelAgentList"}))},{default:t(()=>[s(c(e.$t("button.cancel")),1)]),_:1}),l(f,{type:"primary",onClick:N},{default:t(()=>[s(c(e.$t("button.submit")),1)]),_:1})])]),_:1})]),_:1},8,["model"])])]),_:1})]),_:1},8,["loading"])])}}});const dl=ee(Le,[["__scopeId","data-v-95ae9af9"]]);export{dl as default};
