import{H as be,_ as qe}from"./index.387a2346.js";/* empty css               *//* empty css               *//* empty css               *//* empty css               *//* empty css                *//* empty css                *//* empty css                *//* empty css              *//* empty css               *//* empty css               *//* empty css                */import{d as ce,e as p,B as r,C as y,aH as o,aG as t,aL as s,aM as c,F as M,aJ as w,aI as $,aD as n,aE as _,u as F,bu as T,bv as P,aK as he,aF as ye,bF as Ve,bA as we,bB as $e,b2 as ke,b1 as Ue,bK as Ce,bQ as Ae,bi as Me,aS as Fe,a5 as De,aU as Ee,a4 as je,aT as Re,b4 as Ie,b5 as Qe,bJ as Be,bL as Le,g as Se}from"./arco.17b1a46f.js";import{u as Ne}from"./loading.44762de3.js";import{f as Oe}from"./vue.32c094a4.js";import{p as j}from"./common.ac936b7b.js";import{q as Te,e as Pe}from"./model.c8c8863e.js";import{q as ze}from"./corp.dc6fd269.js";import{f as He}from"./agent.bd7a5b50.js";import"./chart.d5ce7f1f.js";import"./base.87fcf6e2.js";const Ke={class:"container"},Ge={class:"wrapper"},Je={class:"submit-btn"},We={name:"ModelCreate"},Ze=ce({...We,setup(Xe){const{loading:le,setLoading:U}=Ne(!1),{proxy:ae}=Se(),oe=Oe(),R=p([]),B=new Map;(async()=>{U(!0);try{const{data:l}=await ze();R.value=l.items;for(let a=0;a<R.value.length;a+=1)B.set(R.value[a].id,R.value[a])}catch{}finally{U(!1)}})();const I=p([]);(async()=>{U(!0);try{const{data:l}=await Te();I.value=l.items}catch{}finally{U(!1)}})();const z=p([]);(async()=>{U(!0);try{const{data:l}=await He();z.value=l.items}catch{}finally{U(!1)}})();const H=p(),e=p({corp:"",name:"",model:"",type:"1",base_url:"",path:"",remark:"",is_enable_preset_config:!1,preset_config:{is_support_system_role:!0,system_role_prompt:"",min_tokens:p(),max_tokens:p()},text_quota:{billing_method:"1",prompt_ratio:1,completion_ratio:1,fixed_quota:1},image_quotas:[],audio_quota:{billing_method:"1",prompt_ratio:1,completion_ratio:1,fixed_quota:1},multimodal_quota:{text_quota:{billing_method:"1",prompt_ratio:1,completion_ratio:1,fixed_quota:1},image_quotas:[]},midjourney_quotas:[],data_format:"1",is_public:!0,is_enable_model_agent:!1,model_agents:[],is_enable_forward:!1,forward_config:{forward_rule:"1",match_rule:["2"],target_model:"",decision_model:"",keywords:[],target_models:[],content_length:p()},is_enable_fallback:!1,fallback_config:{fallback_model:""}}),ue=async()=>{var a;if(e.value.is_enable_forward||(e.value.forward_config.forward_rule="",e.value.forward_config.match_rule=[],e.value.forward_config.target_model="",e.value.forward_config.keywords=[],e.value.forward_config.target_models=[],e.value.forward_config.content_length=p()),(e.value.forward_config.forward_rule==="1"||e.value.forward_config.forward_rule==="3")&&(e.value.forward_config.keywords=[],e.value.forward_config.target_models=[]),e.value.type==="2"){const b=B.get(e.value.corp);b&&b.code==="Midjourney"?(e.value.image_quotas=[],e.value.multimodal_quota.image_quotas=[]):(e.value.multimodal_quota.image_quotas=[],e.value.midjourney_quotas=[])}else e.value.type==="100"?(e.value.image_quotas=[],e.value.midjourney_quotas=[]):(e.value.image_quotas=[],e.value.multimodal_quota.image_quotas=[],e.value.midjourney_quotas=[]);if(!await((a=H.value)==null?void 0:a.validate())){U(!0);try{await Pe(e.value).then(()=>{ae.$message.success("\u65B0\u5EFA\u6210\u529F"),oe.push({name:"ModelList"})})}catch{}finally{U(!1)}}},te=()=>{S.value=!1,Q.value=!1,q.value=!1;const l=B.get(e.value.corp);if(l&&l.code==="Midjourney"){W();return}K()},Q=p(!1),q=p(!1),h=p(!1),L=p(!1),S=p(!1),K=()=>{if(Q.value=!1,q.value=!1,h.value=!1,L.value=!1,S.value=!1,e.value.text_quota.billing_method="1",e.value.type==="2"){const l=B.get(e.value.corp);if(l&&l.code==="Midjourney"){W();return}if(Q.value=!0,e.value.text_quota.billing_method="2",e.value.image_quotas.length===0){const a=[256,512,1024,1024,1792],b=[256,512,1024,1792,1024];for(let g=0;g<a.length;g+=1)G(a[g],b[g])}}else if(e.value.type==="5"||e.value.type==="6")q.value=!0;else if(e.value.type==="100"&&(h.value=!0,L.value=!0,e.value.multimodal_quota.image_quotas.length===0)){const l=["auto","high","low"];for(let a=0;a<l.length;a+=1)J(l[a])}},G=(l,a)=>{const b={width:l,height:a,fixed_quota:p(),is_default:e.value.image_quotas.length===0?"1":""};e.value.image_quotas.push(b)},de=l=>{e.value.image_quotas.length>1&&(e.value.image_quotas[l].is_default==="1"&&(e.value.image_quotas[l===0?1:0].is_default="1"),e.value.image_quotas.splice(l,1))},re=l=>{for(let a=0;a<e.value.image_quotas.length;a+=1)a===l?e.value.image_quotas[a].is_default="1":e.value.image_quotas[a].is_default=""},J=l=>{const a={mode:l,fixed_quota:p(),is_default:e.value.multimodal_quota.image_quotas.length===0?"1":""};e.value.multimodal_quota.image_quotas.push(a)},ie=l=>{e.value.multimodal_quota.image_quotas.length>1&&(e.value.multimodal_quota.image_quotas[l].is_default==="1"&&(e.value.multimodal_quota.image_quotas[l===0?1:0].is_default="1"),e.value.multimodal_quota.image_quotas.splice(l,1))},me=l=>{for(let a=0;a<e.value.multimodal_quota.image_quotas.length;a+=1)a===l?e.value.multimodal_quota.image_quotas[a].is_default="1":e.value.multimodal_quota.image_quotas[a].is_default=""},W=()=>{if(Q.value=!1,q.value=!1,h.value=!1,L.value=!1,S.value=!0,e.value.type="2",e.value.text_quota.billing_method="2",e.value.midjourney_quotas.length===0){const l=["\u7ED8\u56FE","\u653E\u5927","\u53D8\u6362","\u5F3A\u53D8\u6362","\u5F31\u53D8\u6362","\u63CF\u8FF0","\u6DF7\u56FE","\u91CD\u7ED8","\u5C40\u90E8\u91CD\u7ED8","\u53D8\u7126","\u81EA\u5B9A\u4E49\u53D8\u7126","\u5E73\u79FB","\u7F29\u8BCD","\u7A97\u53E3","\u6362\u8138","\u4EFB\u52A1"],a=["IMAGINE","UPSCALE","VARIATION","HIGH_VARIATION","LOW_VARIATION","DESCRIBE","BLEND","REROLL","INPAINT","ZOOM","CUSTOM_ZOOM","PAN","SHORTEN","MODAL","SWAP_FACE","TASK"],b=["/submit/imagine","/submit/change","/submit/change","/submit/action","/submit/action","/submit/describe","/submit/blend","/submit/action","/submit/action","/submit/action","/submit/action","/submit/action","/submit/shorten","/submit/modal","/insight-face/swap","/task/*"];for(let g=0;g<l.length;g+=1)Z(l[g],a[g],b[g])}},Z=(l,a,b)=>{const g={name:l,action:a,path:b,fixed_quota:p()};e.value.midjourney_quotas.push(g)},ne=l=>{e.value.midjourney_quotas.length>1&&e.value.midjourney_quotas.splice(l,1)},X=()=>{e.value.is_enable_forward?e.value.forward_config.forward_rule==="2"?(e.value.forward_config.keywords=[""],e.value.forward_config.target_models=[""],e.value.forward_config.target_model="",e.value.forward_config.content_length=p()):(e.value.forward_config.forward_rule==="1"||e.value.forward_config.forward_rule==="3")&&(e.value.forward_config.keywords=[],e.value.forward_config.target_models=[]):(e.value.forward_config.target_model="",e.value.forward_config.keywords=[],e.value.forward_config.target_models=[],e.value.forward_config.content_length=p())},se=()=>{e.value.forward_config.keywords.push(""),e.value.forward_config.target_models.push("")},_e=l=>{e.value.forward_config.keywords.length>1&&(e.value.forward_config.keywords.splice(l,1),e.value.forward_config.target_models.splice(l,1))};return(l,a)=>{const b=be,g=he,pe=ye,Y=Ve,v=we,C=$e,i=ke,k=Ue,x=Ce,A=Ae,D=Me,f=Fe,N=De,V=Ee,O=je,E=Re,ee=Ie,fe=Qe,ge=Be,ve=Le;return r(),y("div",Ke,[o(pe,{class:"container-breadcrumb"},{default:t(()=>[o(g,null,{default:t(()=>[o(b)]),_:1}),o(g,null,{default:t(()=>[s(c(l.$t("menu.model")),1)]),_:1}),o(g,null,{default:t(()=>[s(c(l.$t("menu.model.create")),1)]),_:1})]),_:1}),o(ve,{loading:F(le),style:{width:"100%"}},{default:t(()=>[o(ge,{class:"general-card","body-style":{padding:"0 20px 20px 20px"},bordered:!1},{default:t(()=>[M("div",Ge,[o(fe,{ref_key:"formRef",ref:H,model:e.value,class:"form","label-col-props":{span:5},"wrapper-col-props":{span:18}},{default:t(()=>[o(Y,{orientation:"left"},{default:t(()=>[s(c(l.$t("model.title.baseInfo")),1)]),_:1}),o(i,{field:"corp",label:l.$t("model.label.corp"),rules:[{required:!0,message:l.$t("model.error.corp.required")}]},{default:t(()=>[o(C,{modelValue:e.value.corp,"onUpdate:modelValue":a[0]||(a[0]=u=>e.value.corp=u),placeholder:l.$t("model.placeholder.corp"),"allow-search":"",onChange:te},{default:t(()=>[(r(!0),y(w,null,$(R.value,u=>(r(),n(v,{key:u.id,value:u.id,label:u.name},null,8,["value","label"]))),128))]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),o(i,{field:"name",label:l.$t("model.label.name"),rules:[{required:!0,message:l.$t("model.error.name.required")}]},{default:t(()=>[o(k,{modelValue:e.value.name,"onUpdate:modelValue":a[1]||(a[1]=u=>e.value.name=u),placeholder:l.$t("model.placeholder.name")},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),o(i,{field:"model",label:l.$t("model.label.model"),rules:[{required:!0,message:l.$t("model.error.model.required")}]},{default:t(()=>[o(k,{modelValue:e.value.model,"onUpdate:modelValue":a[2]||(a[2]=u=>e.value.model=u),placeholder:l.$t("model.placeholder.model")},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),o(i,{field:"type",label:l.$t("model.label.type"),rules:[{required:!0,message:l.$t("model.error.type.required")}]},{default:t(()=>[o(C,{modelValue:e.value.type,"onUpdate:modelValue":a[3]||(a[3]=u=>e.value.type=u),placeholder:l.$t("model.placeholder.type"),"allow-search":"",onChange:K},{default:t(()=>[o(v,{value:"1"},{default:t(()=>[s("\u6587\u751F\u6587")]),_:1}),o(v,{value:"2"},{default:t(()=>[s("\u6587\u751F\u56FE")]),_:1}),o(v,{value:"3"},{default:t(()=>[s("\u56FE\u751F\u6587")]),_:1}),o(v,{value:"4"},{default:t(()=>[s("\u56FE\u751F\u56FE")]),_:1}),o(v,{value:"5"},{default:t(()=>[s("\u6587\u751F\u8BED\u97F3")]),_:1}),o(v,{value:"6"},{default:t(()=>[s("\u8BED\u97F3\u751F\u6587")]),_:1}),o(v,{value:"100"},{default:t(()=>[s("\u591A\u6A21\u6001")]),_:1})]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),o(i,{field:"base_url",label:l.$t("model.label.base_url")},{default:t(()=>[o(k,{modelValue:e.value.base_url,"onUpdate:modelValue":a[4]||(a[4]=u=>e.value.base_url=u),placeholder:l.$t("model.placeholder.base_url")},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),o(i,{field:"path",label:l.$t("model.label.path")},{default:t(()=>[o(k,{modelValue:e.value.path,"onUpdate:modelValue":a[5]||(a[5]=u=>e.value.path=u),placeholder:l.$t("model.placeholder.path")},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),o(i,{field:"remark",label:l.$t("model.label.remark")},{default:t(()=>[o(x,{modelValue:e.value.remark,"onUpdate:modelValue":a[6]||(a[6]=u=>e.value.remark=u),placeholder:l.$t("model.placeholder.remark")},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),o(Y,{orientation:"left"},{default:t(()=>[s(c(l.$t("model.title.advanced")),1)]),_:1}),!h.value&&!q.value?(r(),n(i,{key:0,field:"text_quota.billing_method",label:l.$t("model.label.billingMethod"),rules:[{required:!0,message:l.$t("model.error.billingMethod.required")}]},{default:t(()=>[o(D,{size:"large"},{default:t(()=>[o(A,{modelValue:e.value.text_quota.billing_method,"onUpdate:modelValue":a[7]||(a[7]=u=>e.value.text_quota.billing_method=u),value:"1","default-checked":!0,disabled:e.value.type==="2"},{default:t(()=>[s("\u500D\u7387")]),_:1},8,["modelValue","disabled"]),o(A,{modelValue:e.value.text_quota.billing_method,"onUpdate:modelValue":a[8]||(a[8]=u=>e.value.text_quota.billing_method=u),value:"2"},{default:t(()=>[s("\u56FA\u5B9A\u989D\u5EA6")]),_:1},8,["modelValue"])]),_:1})]),_:1},8,["label","rules"])):_("",!0),!h.value&&!q.value&&e.value.text_quota.billing_method==="1"?(r(),n(i,{key:1,field:"text_quota.prompt_ratio",label:l.$t("model.label.promptRatio"),rules:[{required:!0,message:l.$t("model.error.promptRatio.required")}]},{default:t(()=>[o(f,{modelValue:e.value.text_quota.prompt_ratio,"onUpdate:modelValue":a[9]||(a[9]=u=>e.value.text_quota.prompt_ratio=u),min:.001,placeholder:l.$t("model.placeholder.promptRatio"),style:{width:"90%","margin-right":"5px"}},null,8,["modelValue","min","placeholder"]),M("div",null," $"+c(F(j)(e.value.text_quota.prompt_ratio))+"/k ",1)]),_:1},8,["label","rules"])):_("",!0),!h.value&&!q.value&&e.value.text_quota.billing_method==="1"?(r(),n(i,{key:2,field:"text_quota.completion_ratio",label:l.$t("model.label.completionRatio"),rules:[{required:!0,message:l.$t("model.error.completionRatio.required")}]},{default:t(()=>[o(f,{modelValue:e.value.text_quota.completion_ratio,"onUpdate:modelValue":a[10]||(a[10]=u=>e.value.text_quota.completion_ratio=u),min:.001,placeholder:l.$t("model.placeholder.completionRatio"),style:{width:"90%","margin-right":"5px"}},null,8,["modelValue","min","placeholder"]),M("div",null," $"+c(F(j)(e.value.text_quota.completion_ratio))+"/k ",1)]),_:1},8,["label","rules"])):_("",!0),!h.value&&!q.value&&e.value.text_quota.billing_method==="2"&&e.value.type!=="2"?(r(),n(i,{key:3,field:"text_quota.fixed_quota",label:l.$t("model.label.fixedQuota"),rules:[{required:!0,message:l.$t("model.error.fixedQuota.required")}]},{default:t(()=>[o(f,{modelValue:e.value.text_quota.fixed_quota,"onUpdate:modelValue":a[11]||(a[11]=u=>e.value.text_quota.fixed_quota=u),min:0,max:9999999999999,placeholder:l.$t("model.placeholder.fixedQuota")},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"])):_("",!0),(r(!0),y(w,null,$(e.value.image_quotas,(u,d)=>T((r(),n(i,{key:d,field:`image_quotas[${d}].width`&&`image_quotas[${d}].height`&&`image_quotas[${d}].fixed_quota`,label:`${d+1}. `+l.$t("model.label.image_quotas"),rules:[{required:!0,message:l.$t("model.error.image_quotas.required")}]},{default:t(()=>[o(f,{modelValue:e.value.image_quotas[d].width,"onUpdate:modelValue":m=>e.value.image_quotas[d].width=m,placeholder:l.$t("model.placeholder.image_quotas.width"),style:{width:"118px","margin-right":"5px"}},null,8,["modelValue","onUpdate:modelValue","placeholder"]),s(" \xD7 "),o(f,{modelValue:e.value.image_quotas[d].height,"onUpdate:modelValue":m=>e.value.image_quotas[d].height=m,placeholder:l.$t("model.placeholder.image_quotas.height"),style:{width:"118px","margin-left":"5px","margin-right":"5px"}},null,8,["modelValue","onUpdate:modelValue","placeholder"]),o(f,{modelValue:e.value.image_quotas[d].fixed_quota,"onUpdate:modelValue":m=>e.value.image_quotas[d].fixed_quota=m,placeholder:l.$t("model.placeholder.image_quotas.fixed_quota"),style:{width:"118px","margin-right":"5px"}},null,8,["modelValue","onUpdate:modelValue","placeholder"]),o(A,{modelValue:e.value.image_quotas[d].is_default,"onUpdate:modelValue":m=>e.value.image_quotas[d].is_default=m,value:"1",style:{width:"60px"},onChange:m=>re(d)},{default:t(()=>[s("\u9ED8\u8BA4")]),_:2},1032,["modelValue","onUpdate:modelValue","onChange"]),o(V,{type:"primary",shape:"circle",style:{margin:"0 10px 0 10px"},onClick:a[12]||(a[12]=m=>G())},{default:t(()=>[o(N)]),_:1}),o(V,{type:"secondary",shape:"circle",onClick:m=>de(d)},{default:t(()=>[o(O)]),_:2},1032,["onClick"])]),_:2},1032,["field","label","rules"])),[[P,Q.value]])),128)),q.value?(r(),n(i,{key:4,field:"audio_quota.billing_method",label:l.$t("model.label.billingMethod"),rules:[{required:!0,message:l.$t("model.error.billingMethod.required")}]},{default:t(()=>[o(D,{size:"large"},{default:t(()=>[o(A,{modelValue:e.value.audio_quota.billing_method,"onUpdate:modelValue":a[13]||(a[13]=u=>e.value.audio_quota.billing_method=u),value:"1","default-checked":!0},{default:t(()=>[s("\u500D\u7387")]),_:1},8,["modelValue"]),o(A,{modelValue:e.value.audio_quota.billing_method,"onUpdate:modelValue":a[14]||(a[14]=u=>e.value.audio_quota.billing_method=u),value:"2"},{default:t(()=>[s("\u56FA\u5B9A\u989D\u5EA6")]),_:1},8,["modelValue"])]),_:1})]),_:1},8,["label","rules"])):_("",!0),q.value&&e.value.audio_quota.billing_method==="1"&&e.value.type!="6"?(r(),n(i,{key:5,field:"audio_quota.prompt_ratio",label:l.$t("model.label.promptRatio"),rules:[{required:!0,message:l.$t("model.error.promptRatio.required")}]},{default:t(()=>[o(f,{modelValue:e.value.audio_quota.prompt_ratio,"onUpdate:modelValue":a[15]||(a[15]=u=>e.value.audio_quota.prompt_ratio=u),min:.001,placeholder:l.$t("model.placeholder.promptRatio"),style:{width:"90%","margin-right":"5px"}},null,8,["modelValue","min","placeholder"]),M("div",null," $"+c(F(j)(e.value.audio_quota.prompt_ratio))+"/k ",1)]),_:1},8,["label","rules"])):_("",!0),q.value&&e.value.audio_quota.billing_method==="1"&&e.value.type!="5"?(r(),n(i,{key:6,field:"audio_quota.completion_ratio",label:l.$t("model.label.completionRatio"),rules:[{required:!0,message:l.$t("model.error.completionRatio.required")}]},{default:t(()=>[o(f,{modelValue:e.value.audio_quota.completion_ratio,"onUpdate:modelValue":a[16]||(a[16]=u=>e.value.audio_quota.completion_ratio=u),min:.001,placeholder:l.$t("model.placeholder.completionRatio"),style:{width:"90%","margin-right":"5px"}},null,8,["modelValue","min","placeholder"]),M("div",null," $"+c(F(j)(e.value.audio_quota.completion_ratio))+"/min ",1)]),_:1},8,["label","rules"])):_("",!0),q.value&&e.value.audio_quota.billing_method==="2"&&e.value.type!=="2"?(r(),n(i,{key:7,field:"audio_quota.fixed_quota",label:l.$t("model.label.fixedQuota"),rules:[{required:!0,message:l.$t("model.error.fixedQuota.required")}]},{default:t(()=>[o(f,{modelValue:e.value.audio_quota.fixed_quota,"onUpdate:modelValue":a[17]||(a[17]=u=>e.value.audio_quota.fixed_quota=u),min:0,max:9999999999999,placeholder:l.$t("model.placeholder.fixedQuota")},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"])):_("",!0),h.value?(r(),n(i,{key:8,field:"multimodal_quota.text_quota.billing_method",label:l.$t("model.label.billingMethod"),rules:[{required:!0,message:l.$t("model.error.billingMethod.required")}]},{default:t(()=>[o(D,{size:"large"},{default:t(()=>[o(A,{modelValue:e.value.multimodal_quota.text_quota.billing_method,"onUpdate:modelValue":a[18]||(a[18]=u=>e.value.multimodal_quota.text_quota.billing_method=u),value:"1","default-checked":!0},{default:t(()=>[s("\u500D\u7387")]),_:1},8,["modelValue"])]),_:1})]),_:1},8,["label","rules"])):_("",!0),h.value&&e.value.multimodal_quota.text_quota.billing_method==="1"?(r(),n(i,{key:9,field:"multimodal_quota.text_quota.prompt_ratio",label:l.$t("model.label.promptRatio"),rules:[{required:!0,message:l.$t("model.error.promptRatio.required")}]},{default:t(()=>[o(f,{modelValue:e.value.multimodal_quota.text_quota.prompt_ratio,"onUpdate:modelValue":a[19]||(a[19]=u=>e.value.multimodal_quota.text_quota.prompt_ratio=u),min:.001,placeholder:l.$t("model.placeholder.promptRatio"),style:{width:"90%","margin-right":"5px"}},null,8,["modelValue","min","placeholder"]),M("div",null," $"+c(F(j)(e.value.multimodal_quota.text_quota.prompt_ratio))+"/k ",1)]),_:1},8,["label","rules"])):_("",!0),h.value&&e.value.multimodal_quota.text_quota.billing_method==="1"?(r(),n(i,{key:10,field:"multimodal_quota.text_quota.completion_ratio",label:l.$t("model.label.completionRatio"),rules:[{required:!0,message:l.$t("model.error.completionRatio.required")}]},{default:t(()=>[o(f,{modelValue:e.value.multimodal_quota.text_quota.completion_ratio,"onUpdate:modelValue":a[20]||(a[20]=u=>e.value.multimodal_quota.text_quota.completion_ratio=u),min:.001,placeholder:l.$t("model.placeholder.completionRatio"),style:{width:"90%","margin-right":"5px"}},null,8,["modelValue","min","placeholder"]),M("div",null," $"+c(F(j)(e.value.multimodal_quota.text_quota.completion_ratio))+"/k ",1)]),_:1},8,["label","rules"])):_("",!0),h.value&&e.value.multimodal_quota.text_quota.billing_method==="2"?(r(),n(i,{key:11,field:"multimodal_quota.text_quota.fixed_quota",label:l.$t("model.label.fixedQuota"),rules:[{required:!0,message:l.$t("model.error.fixedQuota.required")}]},{default:t(()=>[o(f,{modelValue:e.value.multimodal_quota.text_quota.fixed_quota,"onUpdate:modelValue":a[21]||(a[21]=u=>e.value.multimodal_quota.text_quota.fixed_quota=u),min:0,max:9999999999999,placeholder:l.$t("model.placeholder.fixedQuota")},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"])):_("",!0),(r(!0),y(w,null,$(e.value.multimodal_quota.image_quotas,(u,d)=>T((r(),n(i,{key:d,field:`multimodal_quota.image_quotas[${d}].mode`&&`multimodal_quota.image_quotas[${d}].fixed_quota`,label:`${d+1}. `+l.$t("model.label.image_mode_quotas"),rules:[{required:!0,message:l.$t("model.error.image_mode_quotas.required")}]},{default:t(()=>[o(k,{modelValue:e.value.multimodal_quota.image_quotas[d].mode,"onUpdate:modelValue":m=>e.value.multimodal_quota.image_quotas[d].mode=m,placeholder:l.$t("model.placeholder.image_quotas.mode"),style:{width:"185px","margin-right":"5px"}},null,8,["modelValue","onUpdate:modelValue","placeholder"]),o(f,{modelValue:e.value.multimodal_quota.image_quotas[d].fixed_quota,"onUpdate:modelValue":m=>e.value.multimodal_quota.image_quotas[d].fixed_quota=m,placeholder:l.$t("model.placeholder.image_quotas.fixed_quota"),style:{width:"185px","margin-right":"5px"}},null,8,["modelValue","onUpdate:modelValue","placeholder"]),o(A,{modelValue:e.value.multimodal_quota.image_quotas[d].is_default,"onUpdate:modelValue":m=>e.value.multimodal_quota.image_quotas[d].is_default=m,value:"1",style:{width:"60px"},onChange:m=>me(d)},{default:t(()=>[s("\u9ED8\u8BA4")]),_:2},1032,["modelValue","onUpdate:modelValue","onChange"]),o(V,{type:"primary",shape:"circle",style:{margin:"0 10px 0 10px"},onClick:a[22]||(a[22]=m=>J())},{default:t(()=>[o(N)]),_:1}),o(V,{type:"secondary",shape:"circle",onClick:m=>ie(d)},{default:t(()=>[o(O)]),_:2},1032,["onClick"])]),_:2},1032,["field","label","rules"])),[[P,L.value]])),128)),(r(!0),y(w,null,$(e.value.midjourney_quotas,(u,d)=>T((r(),n(i,{key:d,field:`midjourney_quotas[${d}].name`&&`midjourney_quotas[${d}].action`&&`midjourney_quotas[${d}].path`&&`midjourney_quotas[${d}].fixed_quota`,label:`${d+1}. `+l.$t("model.label.midjourney_quotas"),rules:[{required:!0,message:l.$t("model.error.midjourney_quotas.required")}]},{default:t(()=>[o(k,{modelValue:e.value.midjourney_quotas[d].name,"onUpdate:modelValue":m=>e.value.midjourney_quotas[d].name=m,placeholder:l.$t("model.placeholder.midjourney_quotas.name"),style:{width:"95px","margin-right":"5px"}},null,8,["modelValue","onUpdate:modelValue","placeholder"]),o(k,{modelValue:e.value.midjourney_quotas[d].action,"onUpdate:modelValue":m=>e.value.midjourney_quotas[d].action=m,placeholder:l.$t("model.placeholder.midjourney_quotas.action"),style:{width:"102px","margin-right":"5px"}},null,8,["modelValue","onUpdate:modelValue","placeholder"]),o(k,{modelValue:e.value.midjourney_quotas[d].path,"onUpdate:modelValue":m=>e.value.midjourney_quotas[d].path=m,placeholder:l.$t("model.placeholder.midjourney_quotas.path"),style:{width:"138px","margin-right":"5px"}},null,8,["modelValue","onUpdate:modelValue","placeholder"]),o(f,{modelValue:e.value.midjourney_quotas[d].fixed_quota,"onUpdate:modelValue":m=>e.value.midjourney_quotas[d].fixed_quota=m,placeholder:l.$t("model.placeholder.midjourney_quotas.fixed_quota"),style:{width:"90px"}},null,8,["modelValue","onUpdate:modelValue","placeholder"]),o(V,{type:"primary",shape:"circle",style:{margin:"0 10px 0 10px"},onClick:a[23]||(a[23]=m=>Z())},{default:t(()=>[o(N)]),_:1}),o(V,{type:"secondary",shape:"circle",onClick:m=>ne(d)},{default:t(()=>[o(O)]),_:2},1032,["onClick"])]),_:2},1032,["field","label","rules"])),[[P,S.value]])),128)),o(i,{field:"data_format",label:l.$t("model.label.dataFormat"),rules:[{required:!0,message:l.$t("model.error.dataFormat.required")}]},{default:t(()=>[o(D,{size:"large"},{default:t(()=>[o(A,{modelValue:e.value.data_format,"onUpdate:modelValue":a[24]||(a[24]=u=>e.value.data_format=u),value:"1","default-checked":!0},{default:t(()=>[s("\u7EDF\u4E00\u683C\u5F0F")]),_:1},8,["modelValue"])]),_:1})]),_:1},8,["label","rules"]),o(i,{field:"is_public",label:l.$t("model.label.isPublic"),rules:[{required:!0}]},{default:t(()=>[o(E,{modelValue:e.value.is_public,"onUpdate:modelValue":a[25]||(a[25]=u=>e.value.is_public=u)},null,8,["modelValue"])]),_:1},8,["label"]),o(i,{field:"is_enable_preset_config",label:l.$t("model.label.is_enable_preset_config")},{default:t(()=>[o(E,{modelValue:e.value.is_enable_preset_config,"onUpdate:modelValue":a[26]||(a[26]=u=>e.value.is_enable_preset_config=u)},null,8,["modelValue"])]),_:1},8,["label"]),e.value.is_enable_preset_config?(r(),n(i,{key:12,field:"preset_config.is_support_system_role",label:l.$t("model.label.preset_config.is_support_system_role")},{default:t(()=>[o(E,{modelValue:e.value.preset_config.is_support_system_role,"onUpdate:modelValue":a[27]||(a[27]=u=>e.value.preset_config.is_support_system_role=u)},null,8,["modelValue"])]),_:1},8,["label"])):_("",!0),e.value.is_enable_preset_config&&e.value.preset_config.is_support_system_role?(r(),n(i,{key:13,field:"preset_config.system_role_prompt",label:l.$t("model.label.preset_config.system_role_prompt")},{default:t(()=>[o(x,{modelValue:e.value.preset_config.system_role_prompt,"onUpdate:modelValue":a[28]||(a[28]=u=>e.value.preset_config.system_role_prompt=u),placeholder:l.$t("model.placeholder.preset_config.system_role_prompt")},null,8,["modelValue","placeholder"])]),_:1},8,["label"])):_("",!0),e.value.is_enable_preset_config?(r(),n(i,{key:14,field:"preset_config.max_tokens",label:l.$t("model.label.preset_config.max_tokens.range")},{default:t(()=>[o(f,{modelValue:e.value.preset_config.min_tokens,"onUpdate:modelValue":a[29]||(a[29]=u=>e.value.preset_config.min_tokens=u),placeholder:l.$t("model.placeholder.preset_config.min_tokens"),style:{width:"260px","margin-right":"5px"},min:0,max:2097152},null,8,["modelValue","placeholder"]),o(f,{modelValue:e.value.preset_config.max_tokens,"onUpdate:modelValue":a[30]||(a[30]=u=>e.value.preset_config.max_tokens=u),placeholder:l.$t("model.placeholder.preset_config.max_tokens"),style:{width:"260px"},min:0,max:2097152},null,8,["modelValue","placeholder"])]),_:1},8,["field","label"])):_("",!0),o(i,{field:"is_enable_model_agent",label:l.$t("model.label.isEnableModelAgent")},{default:t(()=>[o(E,{modelValue:e.value.is_enable_model_agent,"onUpdate:modelValue":a[31]||(a[31]=u=>e.value.is_enable_model_agent=u)},null,8,["modelValue"])]),_:1},8,["label"]),e.value.is_enable_model_agent?(r(),n(i,{key:15,field:"model_agents",label:l.$t("model.label.modelAgents"),rules:[{required:!0,message:l.$t("model.error.modelAgents.required")}]},{default:t(()=>[o(C,{modelValue:e.value.model_agents,"onUpdate:modelValue":a[32]||(a[32]=u=>e.value.model_agents=u),placeholder:l.$t("model.placeholder.modelAgents"),"max-tag-count":3,multiple:"","allow-search":"","allow-clear":""},{default:t(()=>[(r(!0),y(w,null,$(z.value,u=>(r(),n(v,{key:u.id,value:u.id,label:u.name},null,8,["value","label"]))),128))]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label","rules"])):_("",!0),o(i,{field:"model_forward",label:l.$t("model.label.modelForward")},{default:t(()=>[o(E,{modelValue:e.value.is_enable_forward,"onUpdate:modelValue":a[33]||(a[33]=u=>e.value.is_enable_forward=u),onChange:X},null,8,["modelValue"])]),_:1},8,["label"]),e.value.is_enable_forward?(r(),n(i,{key:16,field:"forward_config.forward_rule",label:l.$t("model.label.forwardRule"),rules:[{required:!0,message:l.$t("model.error.forwardRule.required")}]},{default:t(()=>[o(C,{modelValue:e.value.forward_config.forward_rule,"onUpdate:modelValue":a[34]||(a[34]=u=>e.value.forward_config.forward_rule=u),placeholder:l.$t("model.placeholder.forwardRule"),onChange:X},{default:t(()=>[o(v,{value:"1"},{default:t(()=>[s("\u5168\u90E8\u8F6C\u53D1")]),_:1}),o(v,{value:"2"},{default:t(()=>[s("\u6309\u5173\u952E\u5B57")]),_:1}),o(v,{value:"3"},{default:t(()=>[s("\u5185\u5BB9\u957F\u5EA6")]),_:1})]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label","rules"])):_("",!0),e.value.is_enable_forward&&e.value.forward_config.forward_rule==="3"?(r(),n(i,{key:17,field:"forward_config.content_length",label:l.$t("model.label.content_length"),rules:[{required:!0,message:l.$t("model.error.content_length.required")}]},{default:t(()=>[o(f,{modelValue:e.value.forward_config.content_length,"onUpdate:modelValue":a[35]||(a[35]=u=>e.value.forward_config.content_length=u),min:1,max:9999999999999,placeholder:l.$t("model.placeholder.content_length")},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"])):_("",!0),e.value.is_enable_forward&&(e.value.forward_config.forward_rule==="1"||e.value.forward_config.forward_rule==="3")?(r(),n(i,{key:18,field:"forward_config.target_model",label:l.$t("model.label.targetModel"),rules:[{required:!0,message:l.$t("model.error.targetModel.required")}]},{default:t(()=>[o(C,{modelValue:e.value.forward_config.target_model,"onUpdate:modelValue":a[36]||(a[36]=u=>e.value.forward_config.target_model=u),placeholder:l.$t("model.placeholder.targetModel"),"allow-search":""},{default:t(()=>[(r(!0),y(w,null,$(I.value,u=>(r(),n(v,{key:u.id,value:u.id,label:u.name},null,8,["value","label"]))),128))]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label","rules"])):_("",!0),e.value.is_enable_forward&&e.value.forward_config.forward_rule==="2"?(r(),n(i,{key:19,field:"forward_config.match_rule",label:l.$t("model.label.matchRule"),rules:[{required:!0,message:l.$t("model.error.matchRule.required")}]},{default:t(()=>[o(D,{size:"large"},{default:t(()=>[o(ee,{modelValue:e.value.forward_config.match_rule,"onUpdate:modelValue":a[37]||(a[37]=u=>e.value.forward_config.match_rule=u),value:"1","default-checked":!0},{default:t(()=>[s("\u667A\u80FD\u5339\u914D")]),_:1},8,["modelValue"]),o(ee,{modelValue:e.value.forward_config.match_rule,"onUpdate:modelValue":a[38]||(a[38]=u=>e.value.forward_config.match_rule=u),value:"2"},{default:t(()=>[s("\u6B63\u5219\u5339\u914D")]),_:1},8,["modelValue"])]),_:1})]),_:1},8,["label","rules"])):_("",!0),e.value.is_enable_forward&&e.value.forward_config.forward_rule==="2"&&e.value.forward_config.match_rule.includes("1")?(r(),n(i,{key:20,field:"forward_config.decision_model",label:l.$t("model.label.decisionModel"),rules:[{required:!0,message:l.$t("model.error.decisionModel.required")}]},{default:t(()=>[o(C,{modelValue:e.value.forward_config.decision_model,"onUpdate:modelValue":a[39]||(a[39]=u=>e.value.forward_config.decision_model=u),placeholder:l.$t("model.placeholder.decisionModel"),"allow-search":""},{default:t(()=>[(r(!0),y(w,null,$(I.value,u=>(r(),n(v,{key:u.id,value:u.id,label:u.name},null,8,["value","label"]))),128))]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label","rules"])):_("",!0),(r(!0),y(w,null,$(e.value.forward_config.keywords,(u,d)=>T((r(),n(i,{key:d,field:`forward_config.keywords[${d}]`&&`forward_config.target_models[${d}]`,label:`${d+1}. `+l.$t("model.label.keywords"),rules:[{required:!0,message:l.$t("model.error.keywordsAndtargetModel.required")}]},{default:t(()=>[o(k,{modelValue:e.value.forward_config.keywords[d],"onUpdate:modelValue":m=>e.value.forward_config.keywords[d]=m,placeholder:l.$t("model.placeholder.keywords"),style:{width:"45%","margin-right":"5px"}},null,8,["modelValue","onUpdate:modelValue","placeholder"]),o(C,{modelValue:e.value.forward_config.target_models[d],"onUpdate:modelValue":m=>e.value.forward_config.target_models[d]=m,placeholder:l.$t("model.placeholder.targetModel"),style:{width:"40%"},"allow-search":""},{default:t(()=>[(r(!0),y(w,null,$(I.value,m=>(r(),n(v,{key:m.id,value:m.id,label:m.name},null,8,["value","label"]))),128))]),_:2},1032,["modelValue","onUpdate:modelValue","placeholder"]),o(V,{type:"primary",shape:"circle",style:{margin:"0 10px 0 10px"},onClick:se},{default:t(()=>[o(N)]),_:1}),o(V,{type:"secondary",shape:"circle",onClick:m=>_e(d)},{default:t(()=>[o(O)]),_:2},1032,["onClick"])]),_:2},1032,["field","label","rules"])),[[P,e.value.is_enable_forward&&e.value.forward_config.forward_rule==="2"]])),128)),o(i,{field:"is_enable_fallback",label:l.$t("model.label.is_enable_fallback")},{default:t(()=>[o(E,{modelValue:e.value.is_enable_fallback,"onUpdate:modelValue":a[40]||(a[40]=u=>e.value.is_enable_fallback=u)},null,8,["modelValue"])]),_:1},8,["label"]),e.value.is_enable_fallback?(r(),n(i,{key:21,field:"fallback_config.fallback_model",label:l.$t("model.label.fallback_model"),rules:[{required:!0,message:l.$t("model.error.fallback_model.required")}]},{default:t(()=>[o(C,{modelValue:e.value.fallback_config.fallback_model,"onUpdate:modelValue":a[41]||(a[41]=u=>e.value.fallback_config.fallback_model=u),placeholder:l.$t("model.placeholder.fallback_model"),"allow-search":""},{default:t(()=>[(r(!0),y(w,null,$(I.value,u=>(r(),n(v,{key:u.id,value:u.id,label:u.name},null,8,["value","label"]))),128))]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label","rules"])):_("",!0),o(D,null,{default:t(()=>[M("div",Je,[o(V,{type:"secondary",onClick:a[42]||(a[42]=u=>l.$router.push({name:"ModelList"}))},{default:t(()=>[s(c(l.$t("form.button.cancel")),1)]),_:1}),o(V,{type:"primary",onClick:ue},{default:t(()=>[s(c(l.$t("form.button.submit")),1)]),_:1})])]),_:1})]),_:1},8,["model"])])]),_:1})]),_:1},8,["loading"])])}}});const Vl=qe(Ze,[["__scopeId","data-v-4954f4eb"]]);export{Vl as default};
