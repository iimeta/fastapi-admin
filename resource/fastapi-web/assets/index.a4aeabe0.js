import{_ as R,E as ee}from"./index.dee4ecd9.js";/* empty css               *//* empty css                */import{d as D,e as $,B as i,aD as f,aG as a,aH as l,aL as s,aM as v,bB as P,bC as Q,b2 as j,b1 as J,bK as le,aU as K,b5 as X,F as z,aE as y,C as q,aJ as S,aI as I,bu as ae,bv as oe,bX as re,bi as G,aS as de,aT as te,b4 as ue,a5 as ne,a4 as se,bz as me,bs as ie,u as fe,aK as pe,aF as _e,bL as ce,bM as ve,bJ as be,bN as ge}from"./arco.aed15247.js";import{u as H}from"./loading.b5911e1d.js";import{q as we,d as ye}from"./model.a2e6dd96.js";/* empty css               *//* empty css               *//* empty css                *//* empty css              *//* empty css              *//* empty css               *//* empty css               *//* empty css                */import{f as Ve}from"./agent.5ebb6491.js";/* empty css               */import"./chart.9aa6eafa.js";import"./vue.0cc5b64a.js";import"./base.87fcf6e2.js";const he=D({__name:"base-info",emits:["changeStep"],setup(L,{emit:V}){const c=$(),d=$({corp:"",name:"",model:"",type:"1",base_url:"",path:"",prompt:"",remark:""}),g=async()=>{var m;await((m=c.value)==null?void 0:m.validate())||V("changeStep","forward",{...d.value})};return(n,m)=>{const u=P,e=Q,_=j,b=J,C=le,k=K,M=X;return i(),f(M,{ref_key:"formRef",ref:c,model:d.value,class:"form","label-col-props":{span:5},"wrapper-col-props":{span:18}},{default:a(()=>[l(_,{field:"corp",label:n.$t("model.label.corp"),rules:[{required:!0,message:n.$t("model.error.corp.required")}]},{default:a(()=>[l(e,{modelValue:d.value.corp,"onUpdate:modelValue":m[0]||(m[0]=o=>d.value.corp=o),placeholder:n.$t("model.placeholder.corp")},{default:a(()=>[l(u,{value:"OpenAI"},{default:a(()=>[s("OpenAI")]),_:1}),l(u,{value:"Baidu"},{default:a(()=>[s("\u767E\u5EA6")]),_:1}),l(u,{value:"Xfyun"},{default:a(()=>[s("\u79D1\u5927\u8BAF\u98DE")]),_:1}),l(u,{value:"Aliyun"},{default:a(()=>[s("\u963F\u91CC\u4E91")]),_:1}),l(u,{value:"ZhipuAI"},{default:a(()=>[s("\u667A\u8C31AI")]),_:1}),l(u,{value:"Midjourney"},{default:a(()=>[s("Midjourney")]),_:1})]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),l(_,{field:"name",label:n.$t("model.label.name"),rules:[{required:!0,message:n.$t("model.error.name.required")}]},{default:a(()=>[l(b,{modelValue:d.value.name,"onUpdate:modelValue":m[1]||(m[1]=o=>d.value.name=o),placeholder:n.$t("model.placeholder.name")},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),l(_,{field:"model",label:n.$t("model.label.model"),rules:[{required:!0,message:n.$t("model.error.model.required")}]},{default:a(()=>[l(b,{modelValue:d.value.model,"onUpdate:modelValue":m[2]||(m[2]=o=>d.value.model=o),placeholder:n.$t("model.placeholder.model")},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),l(_,{field:"type",label:n.$t("model.label.type"),rules:[{required:!0,message:n.$t("model.error.type.required")}]},{default:a(()=>[l(e,{modelValue:d.value.type,"onUpdate:modelValue":m[3]||(m[3]=o=>d.value.type=o),placeholder:n.$t("model.placeholder.type")},{default:a(()=>[l(u,{value:"1"},{default:a(()=>[s("\u6587\u751F\u6587")]),_:1}),l(u,{value:"2"},{default:a(()=>[s("\u6587\u751F\u56FE")]),_:1}),l(u,{value:"3"},{default:a(()=>[s("\u56FE\u751F\u6587")]),_:1}),l(u,{value:"4"},{default:a(()=>[s("\u56FE\u751F\u56FE")]),_:1})]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),l(_,{field:"base_url",label:n.$t("model.label.base_url")},{default:a(()=>[l(b,{modelValue:d.value.base_url,"onUpdate:modelValue":m[4]||(m[4]=o=>d.value.base_url=o),placeholder:n.$t("model.placeholder.base_url")},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),l(_,{field:"path",label:n.$t("model.label.path")},{default:a(()=>[l(b,{modelValue:d.value.path,"onUpdate:modelValue":m[5]||(m[5]=o=>d.value.path=o),placeholder:n.$t("model.placeholder.path")},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),l(_,{field:"prompt",label:n.$t("model.label.prompt")},{default:a(()=>[l(C,{modelValue:d.value.prompt,"onUpdate:modelValue":m[6]||(m[6]=o=>d.value.prompt=o),placeholder:n.$t("model.placeholder.prompt")},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),l(_,{field:"remark",label:n.$t("model.label.remark")},{default:a(()=>[l(C,{modelValue:d.value.remark,"onUpdate:modelValue":m[7]||(m[7]=o=>d.value.remark=o),placeholder:n.$t("model.placeholder.remark")},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),l(_,null,{default:a(()=>[l(k,{type:"primary",onClick:g},{default:a(()=>[s(v(n.$t("model.button.next")),1)]),_:1})]),_:1})]),_:1},8,["model"])}}});const $e=R(he,[["__scopeId","data-v-f50e3504"]]),ke=D({__name:"advanced",emits:["changeStep"],setup(L,{emit:V}){const{setLoading:c}=H(!0),d=$([]);(async()=>{c(!0);try{const{data:o}=await we();d.value=o.items}catch{}finally{c(!1)}})();const n=$([]);(async()=>{c(!0);try{const{data:o}=await Ve();n.value=o.items}catch{}finally{c(!1)}})();const u=$(),e=$({billing_method:"1",prompt_ratio:1,completion_ratio:1,fixed_quota:1,data_format:"1",is_public:!0,is_enable_model_agent:!1,model_agents:[],is_forward:!1,forward_config:{forward_rule:"1",match_rule:["2"],target_model:"",decision_model:"",keywords:[],target_models:[]}}),_=()=>{e.value.forward_config.keywords.push(""),e.value.forward_config.target_models.push("")},b=()=>{e.value.is_forward?e.value.forward_config.forward_rule==="2"?(e.value.forward_config.keywords=[""],e.value.forward_config.target_models=[""],e.value.forward_config.target_model=""):e.value.forward_config.forward_rule==="1"&&(e.value.forward_config.keywords=[],e.value.forward_config.target_models=[]):(e.value.forward_config.target_model="",e.value.forward_config.keywords=[],e.value.forward_config.target_models=[])},C=o=>{e.value.forward_config.keywords.length>1&&(e.value.forward_config.keywords.splice(o,1),e.value.forward_config.target_models.splice(o,1))},k=async()=>{var t;e.value.is_forward||(e.value.forward_config.forward_rule="",e.value.forward_config.match_rule=[],e.value.forward_config.target_model="",e.value.forward_config.keywords=[],e.value.forward_config.target_models=[]),e.value.forward_config.forward_rule==="1"&&(e.value.forward_config.keywords=[],e.value.forward_config.target_models=[]),await((t=u.value)==null?void 0:t.validate())||V("changeStep","submit",{...e.value})},M=()=>{V("changeStep","backward")};return(o,t)=>{const F=re,B=G,p=j,N=de,T=te,U=P,A=Q,O=ue,Z=J,W=ne,E=K,Y=se,x=X;return i(),f(x,{ref_key:"formRef",ref:u,model:e.value,class:"form","label-col-props":{span:5},"wrapper-col-props":{span:18}},{default:a(()=>[l(p,{field:"billing_method",label:o.$t("model.label.billingMethod"),rules:[{required:!0,message:o.$t("model.error.billingMethod.required")}]},{default:a(()=>[l(B,{size:"large"},{default:a(()=>[l(F,{modelValue:e.value.billing_method,"onUpdate:modelValue":t[0]||(t[0]=r=>e.value.billing_method=r),value:"1","default-checked":!0},{default:a(()=>[s("\u500D\u7387")]),_:1},8,["modelValue"]),l(F,{modelValue:e.value.billing_method,"onUpdate:modelValue":t[1]||(t[1]=r=>e.value.billing_method=r),value:"2"},{default:a(()=>[s("\u56FA\u5B9A\u989D\u5EA6")]),_:1},8,["modelValue"])]),_:1})]),_:1},8,["label","rules"]),e.value.billing_method==="1"?(i(),f(p,{key:0,field:"prompt_ratio",label:o.$t("model.label.promptRatio"),rules:[{required:!0,message:o.$t("model.error.promptRatio.required")}]},{default:a(()=>[l(N,{modelValue:e.value.prompt_ratio,"onUpdate:modelValue":t[2]||(t[2]=r=>e.value.prompt_ratio=r),min:.001,placeholder:o.$t("model.placeholder.promptRatio"),style:{width:"80%","margin-right":"5px"}},null,8,["modelValue","min","placeholder"]),z("div",null," $"+v(parseFloat((1e3/(5e5/e.value.prompt_ratio)).toFixed(6)))+"/k ",1)]),_:1},8,["label","rules"])):y("",!0),e.value.billing_method==="1"?(i(),f(p,{key:1,field:"completion_ratio",label:o.$t("model.label.completionRatio"),rules:[{required:!0,message:o.$t("model.error.completionRatio.required")}]},{default:a(()=>[l(N,{modelValue:e.value.completion_ratio,"onUpdate:modelValue":t[3]||(t[3]=r=>e.value.completion_ratio=r),min:.001,placeholder:o.$t("model.placeholder.completionRatio"),style:{width:"80%","margin-right":"5px"}},null,8,["modelValue","min","placeholder"]),z("div",null," $"+v(parseFloat((1e3/(5e5/e.value.completion_ratio)).toFixed(6)))+"/k ",1)]),_:1},8,["label","rules"])):y("",!0),e.value.billing_method==="2"?(i(),f(p,{key:2,field:"fixed_quota",label:o.$t("model.label.fixedQuota"),rules:[{required:!0,message:o.$t("model.error.fixedQuota.required")}]},{default:a(()=>[l(N,{modelValue:e.value.fixed_quota,"onUpdate:modelValue":t[4]||(t[4]=r=>e.value.fixed_quota=r),min:0,max:9999999999999,placeholder:o.$t("model.placeholder.fixedQuota")},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"])):y("",!0),l(p,{field:"data_format",label:o.$t("model.label.dataFormat"),rules:[{required:!0,message:o.$t("model.error.dataFormat.required")}]},{default:a(()=>[l(B,{size:"large"},{default:a(()=>[l(F,{modelValue:e.value.data_format,"onUpdate:modelValue":t[5]||(t[5]=r=>e.value.data_format=r),value:"1","default-checked":!0},{default:a(()=>[s("\u7EDF\u4E00\u683C\u5F0F")]),_:1},8,["modelValue"]),l(F,{modelValue:e.value.data_format,"onUpdate:modelValue":t[6]||(t[6]=r=>e.value.data_format=r),value:"2"},{default:a(()=>[s("\u5B98\u65B9\u683C\u5F0F")]),_:1},8,["modelValue"])]),_:1})]),_:1},8,["label","rules"]),l(p,{field:"is_public",label:o.$t("model.label.isPublic"),rules:[{required:!0}]},{default:a(()=>[l(T,{modelValue:e.value.is_public,"onUpdate:modelValue":t[7]||(t[7]=r=>e.value.is_public=r)},null,8,["modelValue"])]),_:1},8,["label"]),l(p,{field:"is_enable_model_agent",label:o.$t("model.label.isEnableModelAgent")},{default:a(()=>[l(T,{modelValue:e.value.is_enable_model_agent,"onUpdate:modelValue":t[8]||(t[8]=r=>e.value.is_enable_model_agent=r)},null,8,["modelValue"])]),_:1},8,["label"]),e.value.is_enable_model_agent?(i(),f(p,{key:3,field:"model_agents",label:o.$t("model.label.modelAgents"),rules:[{required:!0,message:o.$t("model.error.modelAgents.required")}]},{default:a(()=>[l(A,{modelValue:e.value.model_agents,"onUpdate:modelValue":t[9]||(t[9]=r=>e.value.model_agents=r),placeholder:o.$t("model.placeholder.modelAgents"),"max-tag-count":3,multiple:"","allow-clear":""},{default:a(()=>[(i(!0),q(S,null,I(n.value,r=>(i(),f(U,{key:r.id,value:r.id,label:r.name},null,8,["value","label"]))),128))]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label","rules"])):y("",!0),l(p,{field:"model_forward",label:o.$t("model.label.modelForward")},{default:a(()=>[l(T,{modelValue:e.value.is_forward,"onUpdate:modelValue":t[10]||(t[10]=r=>e.value.is_forward=r),onChange:b},null,8,["modelValue"])]),_:1},8,["label"]),e.value.is_forward?(i(),f(p,{key:4,field:"forward_config.forward_rule",label:o.$t("model.label.forwardRule"),rules:[{required:!0,message:o.$t("model.error.forwardRule.required")}]},{default:a(()=>[l(A,{modelValue:e.value.forward_config.forward_rule,"onUpdate:modelValue":t[11]||(t[11]=r=>e.value.forward_config.forward_rule=r),placeholder:o.$t("model.placeholder.forwardRule"),onChange:b},{default:a(()=>[l(U,{value:"1"},{default:a(()=>[s("\u5168\u90E8\u8F6C\u53D1")]),_:1}),l(U,{value:"2"},{default:a(()=>[s("\u6309\u5173\u952E\u5B57")]),_:1})]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label","rules"])):y("",!0),e.value.is_forward&&e.value.forward_config.forward_rule==="1"?(i(),f(p,{key:5,field:"forward_config.target_model",label:o.$t("model.label.targetModel"),rules:[{required:!0,message:o.$t("model.error.targetModel.required")}]},{default:a(()=>[l(A,{modelValue:e.value.forward_config.target_model,"onUpdate:modelValue":t[12]||(t[12]=r=>e.value.forward_config.target_model=r),placeholder:o.$t("model.placeholder.targetModel")},{default:a(()=>[(i(!0),q(S,null,I(d.value,r=>(i(),f(U,{key:r.id,value:r.id,label:r.name},null,8,["value","label"]))),128))]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label","rules"])):y("",!0),e.value.is_forward&&e.value.forward_config.forward_rule==="2"?(i(),f(p,{key:6,field:"forward_config.match_rule",label:o.$t("model.label.matchRule"),rules:[{required:!0,message:o.$t("model.error.matchRule.required")}]},{default:a(()=>[l(B,{size:"large"},{default:a(()=>[l(O,{modelValue:e.value.forward_config.match_rule,"onUpdate:modelValue":t[13]||(t[13]=r=>e.value.forward_config.match_rule=r),value:"1","default-checked":!0},{default:a(()=>[s("\u667A\u80FD\u5339\u914D")]),_:1},8,["modelValue"]),l(O,{modelValue:e.value.forward_config.match_rule,"onUpdate:modelValue":t[14]||(t[14]=r=>e.value.forward_config.match_rule=r),value:"2"},{default:a(()=>[s("\u6B63\u5219\u5339\u914D")]),_:1},8,["modelValue"])]),_:1})]),_:1},8,["label","rules"])):y("",!0),e.value.is_forward&&e.value.forward_config.forward_rule==="2"&&e.value.forward_config.match_rule.includes("1")?(i(),f(p,{key:7,field:"forward_config.decision_model",label:o.$t("model.label.decisionModel"),rules:[{required:!0,message:o.$t("model.error.decisionModel.required")}]},{default:a(()=>[l(A,{modelValue:e.value.forward_config.decision_model,"onUpdate:modelValue":t[15]||(t[15]=r=>e.value.forward_config.decision_model=r),placeholder:o.$t("model.placeholder.decisionModel")},{default:a(()=>[(i(!0),q(S,null,I(d.value,r=>(i(),f(U,{key:r.id,value:r.id,label:r.name},null,8,["value","label"]))),128))]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label","rules"])):y("",!0),(i(!0),q(S,null,I(e.value.forward_config.keywords,(r,w)=>ae((i(),f(p,{key:w,field:`forward_config.keywords[${w}]`&&`forward_config.target_models[${w}]`,label:`${w+1}. `+o.$t("model.label.keywords"),rules:[{required:!0,message:o.$t("model.error.keywordsAndtargetModel.required")}]},{default:a(()=>[l(Z,{modelValue:e.value.forward_config.keywords[w],"onUpdate:modelValue":h=>e.value.forward_config.keywords[w]=h,placeholder:o.$t("model.placeholder.keywords"),style:{width:"40%","margin-right":"5px"}},null,8,["modelValue","onUpdate:modelValue","placeholder"]),l(A,{modelValue:e.value.forward_config.target_models[w],"onUpdate:modelValue":h=>e.value.forward_config.target_models[w]=h,placeholder:o.$t("model.placeholder.targetModel"),style:{width:"40%"}},{default:a(()=>[(i(!0),q(S,null,I(d.value,h=>(i(),f(U,{key:h.id,value:h.id,label:h.name},null,8,["value","label"]))),128))]),_:2},1032,["modelValue","onUpdate:modelValue","placeholder"]),l(E,{type:"primary",shape:"circle",style:{margin:"0 10px 0 10px"},onClick:_},{default:a(()=>[l(W)]),_:1}),l(E,{type:"secondary",shape:"circle",onClick:h=>C(w)},{default:a(()=>[l(Y)]),_:2},1032,["onClick"])]),_:2},1032,["field","label","rules"])),[[oe,e.value.is_forward&&e.value.forward_config.forward_rule==="2"]])),128)),l(p,null,{default:a(()=>[l(B,null,{default:a(()=>[l(E,{type:"secondary",onClick:M},{default:a(()=>[s(v(o.$t("model.button.prev")),1)]),_:1}),l(E,{type:"primary",onClick:k},{default:a(()=>[s(v(o.$t("model.button.next")),1)]),_:1})]),_:1})]),_:1})]),_:1},8,["model"])}}});const qe=R(ke,[["__scopeId","data-v-b10761de"]]),Ce={class:"success-wrap"},Fe=D({__name:"success",emits:["changeStep"],setup(L,{emit:V}){const c=()=>{V("changeStep",1)};return(d,g)=>{const n=me,m=K,u=G;return i(),q("div",Ce,[l(n,{status:"success",title:d.$t("model.success.title"),subtitle:d.$t("model.success.create.subTitle")},null,8,["title","subtitle"]),l(u,{size:16},{default:a(()=>[l(m,{key:"finish",type:"secondary",onClick:g[0]||(g[0]=e=>d.$router.push({name:"ModelList"}))},{default:a(()=>[s(v(d.$t("model.button.finish")),1)]),_:1}),l(m,{key:"again",type:"primary",onClick:c},{default:a(()=>[s(v(d.$t("model.button.again")),1)]),_:1})]),_:1})])}}});const Ue=R(Fe,[["__scopeId","data-v-77484237"]]),Me={class:"container"},Ae={class:"wrapper"},Se={name:"ModelCreate"},Ie=D({...Se,setup(L){const{loading:V,setLoading:c}=H(!1),d=$(1),g=$({}),n=async()=>{c(!0);try{await ye(g.value),d.value=3,g.value={}}catch{}finally{c(!1)}},m=(u,e)=>{if(typeof u=="number"){d.value=u;return}if(u==="forward"||u==="submit"){if(g.value={...g.value,...e},u==="submit"){n();return}d.value+=1}else u==="backward"&&(d.value-=1)};return(u,e)=>{const _=ee,b=pe,C=_e,k=ce,M=ve,o=be,t=ge;return i(),q("div",Me,[l(C,{class:"container-breadcrumb"},{default:a(()=>[l(b,null,{default:a(()=>[l(_)]),_:1}),l(b,null,{default:a(()=>[s(v(u.$t("menu.model")),1)]),_:1}),l(b,null,{default:a(()=>[s(v(u.$t("menu.model.create")),1)]),_:1})]),_:1}),l(t,{loading:fe(V),style:{width:"100%"}},{default:a(()=>[l(o,{class:"general-card",bordered:!1},{default:a(()=>[z("div",Ae,[l(M,{current:d.value,"onUpdate:current":e[0]||(e[0]=F=>d.value=F),style:{width:"660px"},"line-less":"",class:"steps"},{default:a(()=>[l(k,{description:u.$t("model.subTitle.baseInfo")},{default:a(()=>[s(v(u.$t("model.title.baseInfo")),1)]),_:1},8,["description"]),l(k,{description:u.$t("model.subTitle.advanced")},{default:a(()=>[s(v(u.$t("model.title.advanced")),1)]),_:1},8,["description"]),l(k,{description:u.$t("model.subTitle.create.finish")},{default:a(()=>[s(v(u.$t("model.title.create.finish")),1)]),_:1},8,["description"])]),_:1},8,["current"]),(i(),f(ie,null,[d.value===1?(i(),f($e,{key:0,onChangeStep:m})):d.value===2?(i(),f(qe,{key:1,onChangeStep:m})):d.value===3?(i(),f(Ue,{key:2,onChangeStep:m})):y("",!0)],1024))])]),_:1})]),_:1},8,["loading"])])}}});const Ye=R(Ie,[["__scopeId","data-v-9fb2a161"]]);export{Ye as default};
