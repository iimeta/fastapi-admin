import{G as re,_ as ue}from"./index.caf84b09.js";/* empty css               *//* empty css               *//* empty css               *//* empty css               *//* empty css                *//* empty css                *//* empty css                *//* empty css              *//* empty css               *//* empty css               *//* empty css                */import{d as de,e as _,B as d,C as g,aH as a,aG as r,aL as m,aM as c,F,aJ as w,aI as h,aD as s,u as A,aE as p,bu as z,bv as G,aK as te,aF as ne,bF as me,bA as ie,bB as se,b2 as fe,b1 as _e,bK as pe,bQ as ge,bi as be,aS as ve,a5 as ce,aU as we,a4 as he,aT as Ve,b4 as qe,b5 as ye,bJ as $e,bL as ke,g as Ue}from"./arco.06d431a4.js";import{u as Ce}from"./loading.403a6ba1.js";import{f as Fe}from"./vue.4e689e11.js";import{p as J}from"./common.4956e59d.js";import{q as Me,e as Ae}from"./model.d76d5ca7.js";import{q as Be}from"./corp.2f4e2793.js";import{f as De}from"./agent.0bc48605.js";import"./chart.ac3cbee9.js";import"./base.87fcf6e2.js";const Le={class:"container"},Re={class:"wrapper"},Ee={class:"submit-btn"},Ie={name:"ModelCreate"},Qe=de({...Ie,setup(Se){const{loading:P,setLoading:b}=Ce(!1),{proxy:H}=Ue(),O=Fe(),B=_([]);(async()=>{b(!0);try{const{data:l}=await Be();B.value=l.items}catch{}finally{b(!1)}})();const y=_([]);(async()=>{b(!0);try{const{data:l}=await Me();y.value=l.items}catch{}finally{b(!1)}})();const D=_([]);(async()=>{b(!0);try{const{data:l}=await De();D.value=l.items}catch{}finally{b(!1)}})();const L=_(),e=_({corp:"",name:"",model:"",type:"1",base_url:"",path:"",prompt:"",remark:"",text_quota:{billing_method:"1",prompt_ratio:1,completion_ratio:1,fixed_quota:1},image_quotas:[],data_format:"1",is_public:!0,is_enable_model_agent:!1,model_agents:[],is_enable_forward:!1,forward_config:{forward_rule:"1",match_rule:["2"],target_model:"",decision_model:"",keywords:[],target_models:[],content_length:_()},is_enable_fallback:!1,fallback_config:{fallback_model:""}}),j=async()=>{var u;if(e.value.is_enable_forward||(e.value.forward_config.forward_rule="",e.value.forward_config.match_rule=[],e.value.forward_config.target_model="",e.value.forward_config.keywords=[],e.value.forward_config.target_models=[],e.value.forward_config.content_length=_()),(e.value.forward_config.forward_rule==="1"||e.value.forward_config.forward_rule==="3")&&(e.value.forward_config.keywords=[],e.value.forward_config.target_models=[]),!await((u=L.value)==null?void 0:u.validate())){b(!0);try{await Ae(e.value).then(()=>{H.$message.success("\u65B0\u5EFA\u6210\u529F"),O.push({name:"ModelList"})})}catch{}finally{b(!1)}}},W=()=>{e.value.type==="2"&&(e.value.text_quota.billing_method="2",e.value.image_quotas.length===0&&E())},R=()=>{e.value.is_enable_forward?e.value.forward_config.forward_rule==="2"?(e.value.forward_config.keywords=[""],e.value.forward_config.target_models=[""],e.value.forward_config.target_model="",e.value.forward_config.content_length=_()):(e.value.forward_config.forward_rule==="1"||e.value.forward_config.forward_rule==="3")&&(e.value.forward_config.keywords=[],e.value.forward_config.target_models=[]):(e.value.forward_config.target_model="",e.value.forward_config.keywords=[],e.value.forward_config.target_models=[],e.value.forward_config.content_length=_())},E=()=>{const l={fixed_quota:_(),width:_(),height:_(),is_default:e.value.image_quotas.length===0?"1":""};e.value.image_quotas.push(l)},X=l=>{e.value.image_quotas.length>1&&(e.value.image_quotas[l].is_default==="1"&&(e.value.image_quotas[l===0?1:0].is_default="1"),e.value.image_quotas.splice(l,1))},Y=l=>{for(let u=0;u<e.value.image_quotas.length;u+=1)u===l?e.value.image_quotas[u].is_default="1":e.value.image_quotas[u].is_default=""},Z=()=>{e.value.forward_config.keywords.push(""),e.value.forward_config.target_models.push("")},x=l=>{e.value.forward_config.keywords.length>1&&(e.value.forward_config.keywords.splice(l,1),e.value.forward_config.target_models.splice(l,1))};return(l,u)=>{const I=re,M=te,ee=ne,Q=me,f=ie,v=se,t=fe,$=_e,S=pe,k=ge,U=be,V=ve,N=ce,q=we,K=he,C=Ve,T=qe,le=ye,ae=$e,oe=ke;return d(),g("div",Le,[a(ee,{class:"container-breadcrumb"},{default:r(()=>[a(M,null,{default:r(()=>[a(I)]),_:1}),a(M,null,{default:r(()=>[m(c(l.$t("menu.model")),1)]),_:1}),a(M,null,{default:r(()=>[m(c(l.$t("menu.model.create")),1)]),_:1})]),_:1}),a(oe,{loading:A(P),style:{width:"100%"}},{default:r(()=>[a(ae,{class:"general-card","body-style":{padding:"0 20px 20px 20px"},bordered:!1},{default:r(()=>[F("div",Re,[a(le,{ref_key:"formRef",ref:L,model:e.value,class:"form","label-col-props":{span:5},"wrapper-col-props":{span:18}},{default:r(()=>[a(Q,{orientation:"left"},{default:r(()=>[m(c(l.$t("model.title.baseInfo")),1)]),_:1}),a(t,{field:"corp",label:l.$t("model.label.corp"),rules:[{required:!0,message:l.$t("model.error.corp.required")}]},{default:r(()=>[a(v,{modelValue:e.value.corp,"onUpdate:modelValue":u[0]||(u[0]=o=>e.value.corp=o),placeholder:l.$t("model.placeholder.corp"),"allow-search":""},{default:r(()=>[(d(!0),g(w,null,h(B.value,o=>(d(),s(f,{key:o.id,value:o.id,label:o.name},null,8,["value","label"]))),128))]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),a(t,{field:"name",label:l.$t("model.label.name"),rules:[{required:!0,message:l.$t("model.error.name.required")}]},{default:r(()=>[a($,{modelValue:e.value.name,"onUpdate:modelValue":u[1]||(u[1]=o=>e.value.name=o),placeholder:l.$t("model.placeholder.name")},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),a(t,{field:"model",label:l.$t("model.label.model"),rules:[{required:!0,message:l.$t("model.error.model.required")}]},{default:r(()=>[a($,{modelValue:e.value.model,"onUpdate:modelValue":u[2]||(u[2]=o=>e.value.model=o),placeholder:l.$t("model.placeholder.model")},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),a(t,{field:"type",label:l.$t("model.label.type"),rules:[{required:!0,message:l.$t("model.error.type.required")}]},{default:r(()=>[a(v,{modelValue:e.value.type,"onUpdate:modelValue":u[3]||(u[3]=o=>e.value.type=o),placeholder:l.$t("model.placeholder.type"),"allow-search":"",onChange:W},{default:r(()=>[a(f,{value:"1"},{default:r(()=>[m("\u6587\u751F\u6587")]),_:1}),a(f,{value:"2"},{default:r(()=>[m("\u6587\u751F\u56FE")]),_:1}),a(f,{value:"3"},{default:r(()=>[m("\u56FE\u751F\u6587")]),_:1}),a(f,{value:"4"},{default:r(()=>[m("\u56FE\u751F\u56FE")]),_:1}),a(f,{value:"100"},{default:r(()=>[m("\u591A\u6A21\u6001")]),_:1})]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),a(t,{field:"base_url",label:l.$t("model.label.base_url")},{default:r(()=>[a($,{modelValue:e.value.base_url,"onUpdate:modelValue":u[4]||(u[4]=o=>e.value.base_url=o),placeholder:l.$t("model.placeholder.base_url")},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),a(t,{field:"path",label:l.$t("model.label.path")},{default:r(()=>[a($,{modelValue:e.value.path,"onUpdate:modelValue":u[5]||(u[5]=o=>e.value.path=o),placeholder:l.$t("model.placeholder.path")},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),a(t,{field:"prompt",label:l.$t("model.label.prompt")},{default:r(()=>[a(S,{modelValue:e.value.prompt,"onUpdate:modelValue":u[6]||(u[6]=o=>e.value.prompt=o),placeholder:l.$t("model.placeholder.prompt")},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),a(t,{field:"remark",label:l.$t("model.label.remark")},{default:r(()=>[a(S,{modelValue:e.value.remark,"onUpdate:modelValue":u[7]||(u[7]=o=>e.value.remark=o),placeholder:l.$t("model.placeholder.remark")},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),a(Q,{orientation:"left"},{default:r(()=>[m(c(l.$t("model.title.advanced")),1)]),_:1}),a(t,{field:"text_quota.billing_method",label:l.$t("model.label.billingMethod"),rules:[{required:!0,message:l.$t("model.error.billingMethod.required")}]},{default:r(()=>[a(U,{size:"large"},{default:r(()=>[a(k,{modelValue:e.value.text_quota.billing_method,"onUpdate:modelValue":u[8]||(u[8]=o=>e.value.text_quota.billing_method=o),value:"1","default-checked":!0,disabled:e.value.type==="2"},{default:r(()=>[m("\u500D\u7387")]),_:1},8,["modelValue","disabled"]),a(k,{modelValue:e.value.text_quota.billing_method,"onUpdate:modelValue":u[9]||(u[9]=o=>e.value.text_quota.billing_method=o),value:"2"},{default:r(()=>[m("\u56FA\u5B9A\u989D\u5EA6")]),_:1},8,["modelValue"])]),_:1})]),_:1},8,["label","rules"]),e.value.text_quota.billing_method==="1"?(d(),s(t,{key:0,field:"text_quota.prompt_ratio",label:l.$t("model.label.promptRatio"),rules:[{required:!0,message:l.$t("model.error.promptRatio.required")}]},{default:r(()=>[a(V,{modelValue:e.value.text_quota.prompt_ratio,"onUpdate:modelValue":u[10]||(u[10]=o=>e.value.text_quota.prompt_ratio=o),min:.001,placeholder:l.$t("model.placeholder.promptRatio"),style:{width:"90%","margin-right":"5px"}},null,8,["modelValue","min","placeholder"]),F("div",null," $"+c(A(J)(e.value.text_quota.prompt_ratio))+"/k ",1)]),_:1},8,["label","rules"])):p("",!0),e.value.text_quota.billing_method==="1"?(d(),s(t,{key:1,field:"text_quota.completion_ratio",label:l.$t("model.label.completionRatio"),rules:[{required:!0,message:l.$t("model.error.completionRatio.required")}]},{default:r(()=>[a(V,{modelValue:e.value.text_quota.completion_ratio,"onUpdate:modelValue":u[11]||(u[11]=o=>e.value.text_quota.completion_ratio=o),min:.001,placeholder:l.$t("model.placeholder.completionRatio"),style:{width:"90%","margin-right":"5px"}},null,8,["modelValue","min","placeholder"]),F("div",null," $"+c(A(J)(e.value.text_quota.completion_ratio))+"/k ",1)]),_:1},8,["label","rules"])):p("",!0),e.value.text_quota.billing_method==="2"&&e.value.type!=="2"?(d(),s(t,{key:2,field:"text_quota.fixed_quota",label:l.$t("model.label.fixedQuota"),rules:[{required:!0,message:l.$t("model.error.fixedQuota.required")}]},{default:r(()=>[a(V,{modelValue:e.value.text_quota.fixed_quota,"onUpdate:modelValue":u[12]||(u[12]=o=>e.value.text_quota.fixed_quota=o),min:0,max:9999999999999,placeholder:l.$t("model.placeholder.fixedQuota")},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"])):p("",!0),(d(!0),g(w,null,h(e.value.image_quotas,(o,n)=>z((d(),s(t,{key:n,field:`image_quotas[${n}].width`&&`image_quotas[${n}].height`&&`image_quotas[${n}].fixed_quota`,label:`${n+1}. `+l.$t("model.label.image_quotas"),rules:[{required:!0,message:l.$t("model.error.image_quotas.required")}]},{default:r(()=>[a(V,{modelValue:e.value.image_quotas[n].width,"onUpdate:modelValue":i=>e.value.image_quotas[n].width=i,placeholder:l.$t("model.placeholder.image_quotas.width"),style:{width:"118px","margin-right":"5px"}},null,8,["modelValue","onUpdate:modelValue","placeholder"]),m(" \xD7 "),a(V,{modelValue:e.value.image_quotas[n].height,"onUpdate:modelValue":i=>e.value.image_quotas[n].height=i,placeholder:l.$t("model.placeholder.image_quotas.height"),style:{width:"118px","margin-left":"5px","margin-right":"5px"}},null,8,["modelValue","onUpdate:modelValue","placeholder"]),a(V,{modelValue:e.value.image_quotas[n].fixed_quota,"onUpdate:modelValue":i=>e.value.image_quotas[n].fixed_quota=i,placeholder:l.$t("model.placeholder.image_quotas.fixed_quota"),style:{width:"118px","margin-right":"5px"}},null,8,["modelValue","onUpdate:modelValue","placeholder"]),a(k,{modelValue:e.value.image_quotas[n].is_default,"onUpdate:modelValue":i=>e.value.image_quotas[n].is_default=i,value:"1",style:{width:"60px"},onChange:i=>Y(n)},{default:r(()=>[m("\u9ED8\u8BA4")]),_:2},1032,["modelValue","onUpdate:modelValue","onChange"]),a(q,{type:"primary",shape:"circle",style:{margin:"0 10px 0 10px"},onClick:E},{default:r(()=>[a(N)]),_:1}),a(q,{type:"secondary",shape:"circle",onClick:i=>X(n)},{default:r(()=>[a(K)]),_:2},1032,["onClick"])]),_:2},1032,["field","label","rules"])),[[G,e.value.type==="2"]])),128)),a(t,{field:"data_format",label:l.$t("model.label.dataFormat"),rules:[{required:!0,message:l.$t("model.error.dataFormat.required")}]},{default:r(()=>[a(U,{size:"large"},{default:r(()=>[a(k,{modelValue:e.value.data_format,"onUpdate:modelValue":u[13]||(u[13]=o=>e.value.data_format=o),value:"1","default-checked":!0},{default:r(()=>[m("\u7EDF\u4E00\u683C\u5F0F")]),_:1},8,["modelValue"]),a(k,{modelValue:e.value.data_format,"onUpdate:modelValue":u[14]||(u[14]=o=>e.value.data_format=o),value:"2"},{default:r(()=>[m("\u5B98\u65B9\u683C\u5F0F")]),_:1},8,["modelValue"])]),_:1})]),_:1},8,["label","rules"]),a(t,{field:"is_public",label:l.$t("model.label.isPublic"),rules:[{required:!0}]},{default:r(()=>[a(C,{modelValue:e.value.is_public,"onUpdate:modelValue":u[15]||(u[15]=o=>e.value.is_public=o)},null,8,["modelValue"])]),_:1},8,["label"]),a(t,{field:"is_enable_model_agent",label:l.$t("model.label.isEnableModelAgent")},{default:r(()=>[a(C,{modelValue:e.value.is_enable_model_agent,"onUpdate:modelValue":u[16]||(u[16]=o=>e.value.is_enable_model_agent=o)},null,8,["modelValue"])]),_:1},8,["label"]),e.value.is_enable_model_agent?(d(),s(t,{key:3,field:"model_agents",label:l.$t("model.label.modelAgents"),rules:[{required:!0,message:l.$t("model.error.modelAgents.required")}]},{default:r(()=>[a(v,{modelValue:e.value.model_agents,"onUpdate:modelValue":u[17]||(u[17]=o=>e.value.model_agents=o),placeholder:l.$t("model.placeholder.modelAgents"),"max-tag-count":3,multiple:"","allow-search":"","allow-clear":""},{default:r(()=>[(d(!0),g(w,null,h(D.value,o=>(d(),s(f,{key:o.id,value:o.id,label:o.name},null,8,["value","label"]))),128))]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label","rules"])):p("",!0),a(t,{field:"model_forward",label:l.$t("model.label.modelForward")},{default:r(()=>[a(C,{modelValue:e.value.is_enable_forward,"onUpdate:modelValue":u[18]||(u[18]=o=>e.value.is_enable_forward=o),onChange:R},null,8,["modelValue"])]),_:1},8,["label"]),e.value.is_enable_forward?(d(),s(t,{key:4,field:"forward_config.forward_rule",label:l.$t("model.label.forwardRule"),rules:[{required:!0,message:l.$t("model.error.forwardRule.required")}]},{default:r(()=>[a(v,{modelValue:e.value.forward_config.forward_rule,"onUpdate:modelValue":u[19]||(u[19]=o=>e.value.forward_config.forward_rule=o),placeholder:l.$t("model.placeholder.forwardRule"),onChange:R},{default:r(()=>[a(f,{value:"1"},{default:r(()=>[m("\u5168\u90E8\u8F6C\u53D1")]),_:1}),a(f,{value:"2"},{default:r(()=>[m("\u6309\u5173\u952E\u5B57")]),_:1}),a(f,{value:"3"},{default:r(()=>[m("\u5185\u5BB9\u957F\u5EA6")]),_:1})]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label","rules"])):p("",!0),e.value.is_enable_forward&&e.value.forward_config.forward_rule==="3"?(d(),s(t,{key:5,field:"forward_config.content_length",label:l.$t("model.label.content_length"),rules:[{required:!0,message:l.$t("model.error.content_length.required")}]},{default:r(()=>[a(V,{modelValue:e.value.forward_config.content_length,"onUpdate:modelValue":u[20]||(u[20]=o=>e.value.forward_config.content_length=o),min:1,max:9999999999999,placeholder:l.$t("model.placeholder.content_length")},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"])):p("",!0),e.value.is_enable_forward&&(e.value.forward_config.forward_rule==="1"||e.value.forward_config.forward_rule==="3")?(d(),s(t,{key:6,field:"forward_config.target_model",label:l.$t("model.label.targetModel"),rules:[{required:!0,message:l.$t("model.error.targetModel.required")}]},{default:r(()=>[a(v,{modelValue:e.value.forward_config.target_model,"onUpdate:modelValue":u[21]||(u[21]=o=>e.value.forward_config.target_model=o),placeholder:l.$t("model.placeholder.targetModel"),"allow-search":""},{default:r(()=>[(d(!0),g(w,null,h(y.value,o=>(d(),s(f,{key:o.id,value:o.id,label:o.name},null,8,["value","label"]))),128))]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label","rules"])):p("",!0),e.value.is_enable_forward&&e.value.forward_config.forward_rule==="2"?(d(),s(t,{key:7,field:"forward_config.match_rule",label:l.$t("model.label.matchRule"),rules:[{required:!0,message:l.$t("model.error.matchRule.required")}]},{default:r(()=>[a(U,{size:"large"},{default:r(()=>[a(T,{modelValue:e.value.forward_config.match_rule,"onUpdate:modelValue":u[22]||(u[22]=o=>e.value.forward_config.match_rule=o),value:"1","default-checked":!0},{default:r(()=>[m("\u667A\u80FD\u5339\u914D")]),_:1},8,["modelValue"]),a(T,{modelValue:e.value.forward_config.match_rule,"onUpdate:modelValue":u[23]||(u[23]=o=>e.value.forward_config.match_rule=o),value:"2"},{default:r(()=>[m("\u6B63\u5219\u5339\u914D")]),_:1},8,["modelValue"])]),_:1})]),_:1},8,["label","rules"])):p("",!0),e.value.is_enable_forward&&e.value.forward_config.forward_rule==="2"&&e.value.forward_config.match_rule.includes("1")?(d(),s(t,{key:8,field:"forward_config.decision_model",label:l.$t("model.label.decisionModel"),rules:[{required:!0,message:l.$t("model.error.decisionModel.required")}]},{default:r(()=>[a(v,{modelValue:e.value.forward_config.decision_model,"onUpdate:modelValue":u[24]||(u[24]=o=>e.value.forward_config.decision_model=o),placeholder:l.$t("model.placeholder.decisionModel"),"allow-search":""},{default:r(()=>[(d(!0),g(w,null,h(y.value,o=>(d(),s(f,{key:o.id,value:o.id,label:o.name},null,8,["value","label"]))),128))]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label","rules"])):p("",!0),(d(!0),g(w,null,h(e.value.forward_config.keywords,(o,n)=>z((d(),s(t,{key:n,field:`forward_config.keywords[${n}]`&&`forward_config.target_models[${n}]`,label:`${n+1}. `+l.$t("model.label.keywords"),rules:[{required:!0,message:l.$t("model.error.keywordsAndtargetModel.required")}]},{default:r(()=>[a($,{modelValue:e.value.forward_config.keywords[n],"onUpdate:modelValue":i=>e.value.forward_config.keywords[n]=i,placeholder:l.$t("model.placeholder.keywords"),style:{width:"45%","margin-right":"5px"}},null,8,["modelValue","onUpdate:modelValue","placeholder"]),a(v,{modelValue:e.value.forward_config.target_models[n],"onUpdate:modelValue":i=>e.value.forward_config.target_models[n]=i,placeholder:l.$t("model.placeholder.targetModel"),style:{width:"40%"},"allow-search":""},{default:r(()=>[(d(!0),g(w,null,h(y.value,i=>(d(),s(f,{key:i.id,value:i.id,label:i.name},null,8,["value","label"]))),128))]),_:2},1032,["modelValue","onUpdate:modelValue","placeholder"]),a(q,{type:"primary",shape:"circle",style:{margin:"0 10px 0 10px"},onClick:Z},{default:r(()=>[a(N)]),_:1}),a(q,{type:"secondary",shape:"circle",onClick:i=>x(n)},{default:r(()=>[a(K)]),_:2},1032,["onClick"])]),_:2},1032,["field","label","rules"])),[[G,e.value.is_enable_forward&&e.value.forward_config.forward_rule==="2"]])),128)),a(t,{field:"is_enable_fallback",label:l.$t("model.label.is_enable_fallback")},{default:r(()=>[a(C,{modelValue:e.value.is_enable_fallback,"onUpdate:modelValue":u[25]||(u[25]=o=>e.value.is_enable_fallback=o)},null,8,["modelValue"])]),_:1},8,["label"]),e.value.is_enable_fallback?(d(),s(t,{key:9,field:"fallback_config.fallback_model",label:l.$t("model.label.fallback_model"),rules:[{required:!0,message:l.$t("model.error.fallback_model.required")}]},{default:r(()=>[a(v,{modelValue:e.value.fallback_config.fallback_model,"onUpdate:modelValue":u[26]||(u[26]=o=>e.value.fallback_config.fallback_model=o),placeholder:l.$t("model.placeholder.fallback_model"),"allow-search":""},{default:r(()=>[(d(!0),g(w,null,h(y.value,o=>(d(),s(f,{key:o.id,value:o.id,label:o.name},null,8,["value","label"]))),128))]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label","rules"])):p("",!0),a(U,null,{default:r(()=>[F("div",Ee,[a(q,{type:"secondary",onClick:u[27]||(u[27]=o=>l.$router.push({name:"ModelList"}))},{default:r(()=>[m(c(l.$t("form.button.cancel")),1)]),_:1}),a(q,{type:"primary",onClick:j},{default:r(()=>[m(c(l.$t("form.button.submit")),1)]),_:1})])]),_:1})]),_:1},8,["model"])])]),_:1})]),_:1},8,["loading"])])}}});const ml=ue(Qe,[["__scopeId","data-v-d078d73e"]]);export{ml as default};
