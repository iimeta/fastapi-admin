import{H as ce,_ as be}from"./index.b8904415.js";/* empty css               *//* empty css               *//* empty css               *//* empty css               *//* empty css                *//* empty css                *//* empty css                *//* empty css              *//* empty css               *//* empty css               *//* empty css                */import{d as qe,e as c,B as m,C as h,aH as o,aG as t,aL as _,aM as V,F as A,aJ as $,aI as k,aD as s,aE as p,u as B,bu as T,bv as P,aK as he,aF as ye,bF as we,bA as Ve,bB as $e,b2 as ke,b1 as Ue,bK as Ce,bQ as Me,bi as je,aS as Ae,a5 as De,aU as Fe,a4 as Ee,aT as Ie,b4 as Re,b5 as Be,bJ as Qe,bL as Se,g as Le}from"./arco.17b1a46f.js";import{u as Ne}from"./loading.44762de3.js";import{h as Oe,f as Te}from"./vue.32c094a4.js";import{p as H}from"./common.ac936b7b.js";import{q as Pe,f as He,g as Ke}from"./model.65b014dc.js";import{q as ze}from"./corp.1c244aa8.js";import{f as Ge}from"./agent.260ead94.js";import"./chart.d5ce7f1f.js";import"./base.87fcf6e2.js";const Je={class:"container"},We={class:"wrapper"},Ze={class:"submit-btn"},Xe={name:"ModelUpdate"},Ye=qe({...Xe,setup(xe){const{loading:le,setLoading:y}=Ne(!0),{proxy:ae}=Le(),oe=Oe(),ue=Te(),D=c([]),Q=new Map;(async()=>{y(!0);try{const{data:l}=await ze();D.value=l.items;for(let a=0;a<D.value.length;a+=1)Q.set(D.value[a].id,D.value[a])}catch{}finally{y(!1)}})();const F=c([]);(async()=>{y(!0);try{const{data:l}=await Pe();F.value=l.items}catch{}finally{y(!1)}})();const K=c([]);(async()=>{y(!0);try{const{data:l}=await Ge();K.value=l.items}catch{}finally{y(!1)}})();const z=c(),e=c({id:"",corp:"",name:"",model:"",type:"",base_url:"",path:"",remark:"",status:1,is_enable_preset_config:!1,preset_config:{is_support_system_role:!0,system_role_prompt:"",min_tokens:c(),max_tokens:c()},text_quota:{billing_method:"1",prompt_ratio:1,completion_ratio:1,fixed_quota:1},image_quotas:[],multimodal_quota:{text_quota:{billing_method:"1",prompt_ratio:1,completion_ratio:1,fixed_quota:1},image_quotas:[]},midjourney_quotas:[],data_format:"",is_public:!0,is_enable_model_agent:!1,model_agents:[],is_enable_forward:!1,forward_config:{forward_rule:"1",match_rule:["2"],target_model:"",decision_model:"",keywords:[],target_models:[],content_length:c()},is_enable_fallback:!1,fallback_config:{fallback_model:""}}),te=async()=>{var a;if(e.value.type==="2"){const v=Q.get(e.value.corp);v&&v.code==="Midjourney"?(e.value.image_quotas=[],e.value.multimodal_quota.image_quotas=[]):(e.value.multimodal_quota.image_quotas=[],e.value.midjourney_quotas=[])}else e.value.type==="100"?(e.value.image_quotas=[],e.value.midjourney_quotas=[]):(e.value.image_quotas=[],e.value.multimodal_quota.image_quotas=[],e.value.midjourney_quotas=[]);if(!await((a=z.value)==null?void 0:a.validate())){y(!0);try{await He(e.value).then(()=>{ae.$message.success("\u66F4\u65B0\u6210\u529F"),ue.push({name:"ModelList"})})}catch{}finally{y(!1)}}};(async(l={id:oe.query.id})=>{var a,v,f,L,I,g;y(!0);try{const{data:d}=await Ke(l);if(e.value.id=d.id,e.value.corp=d.corp,e.value.name=d.name,e.value.model=d.model,e.value.type=String(d.type),e.value.base_url=d.base_url,e.value.path=d.path,e.value.remark=d.remark,e.value.status=d.status,e.value.data_format=String(d.data_format),e.value.is_public=d.is_public,e.value.is_enable_preset_config=d.is_enable_preset_config,e.value.preset_config=d.preset_config,e.value.is_enable_model_agent=d.is_enable_model_agent,e.value.model_agents=d.model_agents,e.value.is_enable_forward=d.is_enable_forward,d.text_quota&&(e.value.text_quota=d.text_quota,e.value.text_quota.billing_method=String(e.value.text_quota.billing_method)),d.image_quotas){C.value=e.value.type==="2"&&d.corp_code!=="Midjourney",e.value.image_quotas=d.image_quotas;for(let i=0;i<e.value.image_quotas.length;i+=1)if(e.value.image_quotas[i].is_default){e.value.image_quotas[i].is_default="1";break}}if(d.multimodal_quota)if(q.value=e.value.type==="100"&&d.corp_code!=="Midjourney",S.value=e.value.type==="100"&&d.corp_code!=="Midjourney",e.value.multimodal_quota=d.multimodal_quota,e.value.multimodal_quota.text_quota.billing_method=String(e.value.multimodal_quota.text_quota.billing_method),e.value.multimodal_quota.image_quotas){for(let i=0;i<e.value.multimodal_quota.image_quotas.length;i+=1)if(e.value.multimodal_quota.image_quotas[i].is_default){e.value.multimodal_quota.image_quotas[i].is_default="1";break}}else e.value.multimodal_quota.image_quotas=[];d.midjourney_quotas&&(E.value=e.value.type==="2"&&d.corp_code==="Midjourney",e.value.midjourney_quotas=d.midjourney_quotas),d.forward_config&&(d.forward_config.forward_rule&&(e.value.forward_config.forward_rule=String(d.forward_config.forward_rule)),d.forward_config.match_rule&&(e.value.forward_config.match_rule=d.forward_config.match_rule.map(String)),e.value.forward_config.content_length=(a=d.forward_config)==null?void 0:a.content_length,e.value.forward_config.target_model=(v=d.forward_config)==null?void 0:v.target_model,e.value.forward_config.decision_model=(f=d.forward_config)==null?void 0:f.decision_model,e.value.forward_config.keywords=(L=d.forward_config)==null?void 0:L.keywords,e.value.forward_config.target_models=(I=d.forward_config)==null?void 0:I.target_models),e.value.is_enable_fallback=d.is_enable_fallback,d.fallback_config&&(e.value.fallback_config.fallback_model=(g=d.fallback_config)==null?void 0:g.fallback_model)}catch{}finally{y(!1)}})();const re=()=>{E.value=!1,C.value=!1;const l=Q.get(e.value.corp);if(l&&l.code==="Midjourney"){Z();return}G()},C=c(!1),q=c(!1),S=c(!1),E=c(!1),G=()=>{if(C.value=!1,q.value=!1,S.value=!1,E.value=!1,e.value.text_quota.billing_method="1",e.value.type==="2"){const l=Q.get(e.value.corp);if(l&&l.code==="Midjourney"){Z();return}if(C.value=!0,e.value.text_quota.billing_method="2",e.value.image_quotas.length===0){const a=[256,512,1024,1024,1792],v=[256,512,1024,1792,1024];for(let f=0;f<a.length;f+=1)J(a[f],v[f])}}else if(e.value.type==="100"&&(q.value=!0,S.value=!0,e.value.multimodal_quota.image_quotas.length===0)){const l=["auto","high","low"];for(let a=0;a<l.length;a+=1)W(l[a])}},J=(l,a)=>{const v={width:l,height:a,fixed_quota:c(),is_default:e.value.image_quotas.length===0?"1":""};e.value.image_quotas.push(v)},de=l=>{e.value.image_quotas.length>1&&(e.value.image_quotas[l].is_default==="1"&&(e.value.image_quotas[l===0?1:0].is_default="1"),e.value.image_quotas.splice(l,1))},ie=l=>{for(let a=0;a<e.value.image_quotas.length;a+=1)a===l?e.value.image_quotas[a].is_default="1":e.value.image_quotas[a].is_default=""},W=l=>{const a={mode:l,fixed_quota:c(),is_default:e.value.multimodal_quota.image_quotas.length===0?"1":""};e.value.multimodal_quota.image_quotas.push(a)},me=l=>{e.value.multimodal_quota.image_quotas.length>1&&(e.value.multimodal_quota.image_quotas[l].is_default==="1"&&(e.value.multimodal_quota.image_quotas[l===0?1:0].is_default="1"),e.value.multimodal_quota.image_quotas.splice(l,1))},ne=l=>{for(let a=0;a<e.value.multimodal_quota.image_quotas.length;a+=1)a===l?e.value.multimodal_quota.image_quotas[a].is_default="1":e.value.multimodal_quota.image_quotas[a].is_default=""},Z=()=>{if(C.value=!1,E.value=!0,e.value.type="2",e.value.text_quota.billing_method="2",e.value.midjourney_quotas.length===0){const l=["\u7ED8\u56FE","\u653E\u5927","\u53D8\u6362","\u5F3A\u53D8\u6362","\u5F31\u53D8\u6362","\u63CF\u8FF0","\u6DF7\u56FE","\u91CD\u7ED8","\u5C40\u90E8\u91CD\u7ED8","\u53D8\u7126","\u81EA\u5B9A\u4E49\u53D8\u7126","\u5E73\u79FB","\u7F29\u8BCD","\u7A97\u53E3","\u6362\u8138","\u4EFB\u52A1"],a=["IMAGINE","UPSCALE","VARIATION","HIGH_VARIATION","LOW_VARIATION","DESCRIBE","BLEND","REROLL","INPAINT","ZOOM","CUSTOM_ZOOM","PAN","SHORTEN","MODAL","SWAP_FACE","TASK"],v=["/submit/imagine","/submit/change","/submit/change","/submit/action","/submit/action","/submit/describe","/submit/blend","/submit/action","/submit/action","/submit/action","/submit/action","/submit/action","/submit/shorten","/submit/modal","/insight-face/swap","/task/*"];for(let f=0;f<l.length;f+=1)X(l[f],a[f],v[f])}},X=(l,a,v)=>{const f={name:l,action:a,path:v,fixed_quota:c()};e.value.midjourney_quotas.push(f)},se=l=>{e.value.midjourney_quotas.length>1&&e.value.midjourney_quotas.splice(l,1)},Y=()=>{!e.value.is_enable_forward&&e.value.forward_config.keywords&&(e.value.forward_config.keywords[0]===""||e.value.forward_config.target_models[0]==="")?(e.value.forward_config.keywords=[],e.value.forward_config.target_models=[]):e.value.forward_config.forward_rule==="2"&&(!e.value.forward_config.keywords||e.value.forward_config.keywords.length===0)?(e.value.forward_config.keywords=[""],e.value.forward_config.target_models=[""]):(e.value.forward_config.forward_rule==="1"||e.value.forward_config.forward_rule==="3")&&e.value.forward_config.keywords&&(e.value.forward_config.keywords[0]===""||e.value.forward_config.target_models[0]==="")&&(e.value.forward_config.keywords=[],e.value.forward_config.target_models=[])},_e=()=>{e.value.forward_config.keywords.push(""),e.value.forward_config.target_models.push("")},fe=l=>{e.value.forward_config.keywords.length>1&&(e.value.forward_config.keywords.splice(l,1),e.value.forward_config.target_models.splice(l,1))};return(l,a)=>{const v=ce,f=he,L=ye,I=we,g=Ve,d=$e,i=ke,U=Ue,x=Ce,M=Me,R=je,b=Ae,N=De,w=Fe,O=Ee,j=Ie,ee=Re,pe=Be,ge=Qe,ve=Se;return m(),h("div",Je,[o(L,{class:"container-breadcrumb"},{default:t(()=>[o(f,null,{default:t(()=>[o(v)]),_:1}),o(f,null,{default:t(()=>[_(V(l.$t("menu.model")),1)]),_:1}),o(f,null,{default:t(()=>[_(V(l.$t("menu.model.update")),1)]),_:1})]),_:1}),o(ve,{loading:B(le),style:{width:"100%"}},{default:t(()=>[o(ge,{class:"general-card","body-style":{padding:"0 20px 20px 20px"},bordered:!1},{default:t(()=>[A("div",We,[o(pe,{ref_key:"formRef",ref:z,model:e.value,class:"form","label-col-props":{span:5},"wrapper-col-props":{span:18}},{default:t(()=>[o(I,{orientation:"left"},{default:t(()=>[_(V(l.$t("model.title.baseInfo")),1)]),_:1}),o(i,{field:"corp",label:l.$t("model.label.corp"),rules:[{required:!0,message:l.$t("model.error.corp.required")}]},{default:t(()=>[o(d,{modelValue:e.value.corp,"onUpdate:modelValue":a[0]||(a[0]=u=>e.value.corp=u),placeholder:l.$t("model.placeholder.corp"),"allow-search":"",onChange:re},{default:t(()=>[(m(!0),h($,null,k(D.value,u=>(m(),s(g,{key:u.id,value:u.id,label:u.name},null,8,["value","label"]))),128))]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),o(i,{field:"name",label:l.$t("model.label.name"),rules:[{required:!0,message:l.$t("model.error.name.required")}]},{default:t(()=>[o(U,{modelValue:e.value.name,"onUpdate:modelValue":a[1]||(a[1]=u=>e.value.name=u),placeholder:l.$t("model.placeholder.name")},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),o(i,{field:"model",label:l.$t("model.label.model"),rules:[{required:!0,message:l.$t("model.error.model.required")}]},{default:t(()=>[o(U,{modelValue:e.value.model,"onUpdate:modelValue":a[2]||(a[2]=u=>e.value.model=u),placeholder:l.$t("model.placeholder.model")},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),o(i,{field:"type",label:l.$t("model.label.type"),rules:[{required:!0,message:l.$t("model.error.type.required")}]},{default:t(()=>[o(d,{modelValue:e.value.type,"onUpdate:modelValue":a[3]||(a[3]=u=>e.value.type=u),placeholder:l.$t("model.placeholder.type"),"allow-search":"",onChange:G},{default:t(()=>[o(g,{value:"1"},{default:t(()=>[_("\u6587\u751F\u6587")]),_:1}),o(g,{value:"2"},{default:t(()=>[_("\u6587\u751F\u56FE")]),_:1}),o(g,{value:"3"},{default:t(()=>[_("\u56FE\u751F\u6587")]),_:1}),o(g,{value:"4"},{default:t(()=>[_("\u56FE\u751F\u56FE")]),_:1}),o(g,{value:"100"},{default:t(()=>[_("\u591A\u6A21\u6001")]),_:1})]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),o(i,{field:"base_url",label:l.$t("model.label.base_url")},{default:t(()=>[o(U,{modelValue:e.value.base_url,"onUpdate:modelValue":a[4]||(a[4]=u=>e.value.base_url=u),placeholder:l.$t("model.placeholder.base_url")},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),o(i,{field:"path",label:l.$t("model.label.path")},{default:t(()=>[o(U,{modelValue:e.value.path,"onUpdate:modelValue":a[5]||(a[5]=u=>e.value.path=u),placeholder:l.$t("model.placeholder.path")},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),o(i,{field:"remark",label:l.$t("model.label.remark")},{default:t(()=>[o(x,{modelValue:e.value.remark,"onUpdate:modelValue":a[6]||(a[6]=u=>e.value.remark=u),placeholder:l.$t("model.placeholder.remark")},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),o(I,{orientation:"left"},{default:t(()=>[_(V(l.$t("model.title.advanced")),1)]),_:1}),q.value?p("",!0):(m(),s(i,{key:0,field:"text_quota.billing_method",label:l.$t("model.label.billingMethod"),rules:[{required:!0,message:l.$t("model.error.billingMethod.required")}]},{default:t(()=>[o(R,{size:"large"},{default:t(()=>[o(M,{modelValue:e.value.text_quota.billing_method,"onUpdate:modelValue":a[7]||(a[7]=u=>e.value.text_quota.billing_method=u),value:"1","default-checked":!0,disabled:e.value.type==="2"},{default:t(()=>[_("\u500D\u7387")]),_:1},8,["modelValue","disabled"]),o(M,{modelValue:e.value.text_quota.billing_method,"onUpdate:modelValue":a[8]||(a[8]=u=>e.value.text_quota.billing_method=u),value:"2"},{default:t(()=>[_("\u56FA\u5B9A\u989D\u5EA6")]),_:1},8,["modelValue"])]),_:1})]),_:1},8,["label","rules"])),!q.value&&e.value.text_quota.billing_method==="1"?(m(),s(i,{key:1,field:"text_quota.prompt_ratio",label:l.$t("model.label.promptRatio"),rules:[{required:!0,message:l.$t("model.error.promptRatio.required")}]},{default:t(()=>[o(b,{modelValue:e.value.text_quota.prompt_ratio,"onUpdate:modelValue":a[9]||(a[9]=u=>e.value.text_quota.prompt_ratio=u),min:.001,placeholder:l.$t("model.placeholder.promptRatio"),style:{width:"90%","margin-right":"5px"}},null,8,["modelValue","min","placeholder"]),A("div",null," $"+V(B(H)(e.value.text_quota.prompt_ratio))+"/k ",1)]),_:1},8,["label","rules"])):p("",!0),!q.value&&e.value.text_quota.billing_method==="1"?(m(),s(i,{key:2,field:"text_quota.completion_ratio",label:l.$t("model.label.completionRatio"),rules:[{required:!0,message:l.$t("model.error.completionRatio.required")}]},{default:t(()=>[o(b,{modelValue:e.value.text_quota.completion_ratio,"onUpdate:modelValue":a[10]||(a[10]=u=>e.value.text_quota.completion_ratio=u),min:.001,placeholder:l.$t("model.placeholder.completionRatio"),style:{width:"90%","margin-right":"5px"}},null,8,["modelValue","min","placeholder"]),A("div",null," $"+V(B(H)(e.value.text_quota.completion_ratio))+"/k ",1)]),_:1},8,["label","rules"])):p("",!0),!q.value&&e.value.text_quota.billing_method==="2"&&e.value.type!=="2"?(m(),s(i,{key:3,field:"text_quota.fixed_quota",label:l.$t("model.label.fixedQuota"),rules:[{required:!0,message:l.$t("model.error.fixedQuota.required")}]},{default:t(()=>[o(b,{modelValue:e.value.text_quota.fixed_quota,"onUpdate:modelValue":a[11]||(a[11]=u=>e.value.text_quota.fixed_quota=u),min:0,max:9999999999999,placeholder:l.$t("model.placeholder.fixedQuota")},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"])):p("",!0),(m(!0),h($,null,k(e.value.image_quotas,(u,r)=>T((m(),s(i,{key:r,field:`image_quotas[${r}].width`&&`image_quotas[${r}].height`&&`image_quotas[${r}].fixed_quota`,label:`${r+1}. `+l.$t("model.label.image_quotas"),rules:[{required:!0,message:l.$t("model.error.image_quotas.required")}]},{default:t(()=>[o(b,{modelValue:e.value.image_quotas[r].width,"onUpdate:modelValue":n=>e.value.image_quotas[r].width=n,placeholder:l.$t("model.placeholder.image_quotas.width"),style:{width:"118px","margin-right":"5px"}},null,8,["modelValue","onUpdate:modelValue","placeholder"]),_(" \xD7 "),o(b,{modelValue:e.value.image_quotas[r].height,"onUpdate:modelValue":n=>e.value.image_quotas[r].height=n,placeholder:l.$t("model.placeholder.image_quotas.height"),style:{width:"118px","margin-left":"5px","margin-right":"5px"}},null,8,["modelValue","onUpdate:modelValue","placeholder"]),o(b,{modelValue:e.value.image_quotas[r].fixed_quota,"onUpdate:modelValue":n=>e.value.image_quotas[r].fixed_quota=n,placeholder:l.$t("model.placeholder.image_quotas.fixed_quota"),style:{width:"118px","margin-right":"5px"}},null,8,["modelValue","onUpdate:modelValue","placeholder"]),o(M,{modelValue:e.value.image_quotas[r].is_default,"onUpdate:modelValue":n=>e.value.image_quotas[r].is_default=n,value:"1",style:{width:"60px"},onChange:n=>ie(r)},{default:t(()=>[_("\u9ED8\u8BA4")]),_:2},1032,["modelValue","onUpdate:modelValue","onChange"]),o(w,{type:"primary",shape:"circle",style:{margin:"0 10px 0 10px"},onClick:a[12]||(a[12]=n=>J())},{default:t(()=>[o(N)]),_:1}),o(w,{type:"secondary",shape:"circle",onClick:n=>de(r)},{default:t(()=>[o(O)]),_:2},1032,["onClick"])]),_:2},1032,["field","label","rules"])),[[P,C.value]])),128)),q.value?(m(),s(i,{key:4,field:"multimodal_quota.text_quota.billing_method",label:l.$t("model.label.billingMethod"),rules:[{required:!0,message:l.$t("model.error.billingMethod.required")}]},{default:t(()=>[o(R,{size:"large"},{default:t(()=>[o(M,{modelValue:e.value.multimodal_quota.text_quota.billing_method,"onUpdate:modelValue":a[13]||(a[13]=u=>e.value.multimodal_quota.text_quota.billing_method=u),value:"1","default-checked":!0},{default:t(()=>[_("\u500D\u7387")]),_:1},8,["modelValue"])]),_:1})]),_:1},8,["label","rules"])):p("",!0),q.value&&e.value.multimodal_quota.text_quota.billing_method==="1"?(m(),s(i,{key:5,field:"multimodal_quota.text_quota.prompt_ratio",label:l.$t("model.label.promptRatio"),rules:[{required:!0,message:l.$t("model.error.promptRatio.required")}]},{default:t(()=>[o(b,{modelValue:e.value.multimodal_quota.text_quota.prompt_ratio,"onUpdate:modelValue":a[14]||(a[14]=u=>e.value.multimodal_quota.text_quota.prompt_ratio=u),min:.001,placeholder:l.$t("model.placeholder.promptRatio"),style:{width:"90%","margin-right":"5px"}},null,8,["modelValue","min","placeholder"]),A("div",null," $"+V(B(H)(e.value.multimodal_quota.text_quota.prompt_ratio))+"/k ",1)]),_:1},8,["label","rules"])):p("",!0),q.value&&e.value.multimodal_quota.text_quota.billing_method==="1"?(m(),s(i,{key:6,field:"multimodal_quota.text_quota.completion_ratio",label:l.$t("model.label.completionRatio"),rules:[{required:!0,message:l.$t("model.error.completionRatio.required")}]},{default:t(()=>[o(b,{modelValue:e.value.multimodal_quota.text_quota.completion_ratio,"onUpdate:modelValue":a[15]||(a[15]=u=>e.value.multimodal_quota.text_quota.completion_ratio=u),min:.001,placeholder:l.$t("model.placeholder.completionRatio"),style:{width:"90%","margin-right":"5px"}},null,8,["modelValue","min","placeholder"]),A("div",null," $"+V(B(H)(e.value.multimodal_quota.text_quota.completion_ratio))+"/k ",1)]),_:1},8,["label","rules"])):p("",!0),q.value&&e.value.multimodal_quota.text_quota.billing_method==="2"?(m(),s(i,{key:7,field:"multimodal_quota.text_quota.fixed_quota",label:l.$t("model.label.fixedQuota"),rules:[{required:!0,message:l.$t("model.error.fixedQuota.required")}]},{default:t(()=>[o(b,{modelValue:e.value.multimodal_quota.text_quota.fixed_quota,"onUpdate:modelValue":a[16]||(a[16]=u=>e.value.multimodal_quota.text_quota.fixed_quota=u),min:0,max:9999999999999,placeholder:l.$t("model.placeholder.fixedQuota")},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"])):p("",!0),(m(!0),h($,null,k(e.value.multimodal_quota.image_quotas,(u,r)=>T((m(),s(i,{key:r,field:`multimodal_quota.image_quotas[${r}].mode`&&`multimodal_quota.image_quotas[${r}].fixed_quota`,label:`${r+1}. `+l.$t("model.label.image_mode_quotas"),rules:[{required:!0,message:l.$t("model.error.image_mode_quotas.required")}]},{default:t(()=>[o(U,{modelValue:e.value.multimodal_quota.image_quotas[r].mode,"onUpdate:modelValue":n=>e.value.multimodal_quota.image_quotas[r].mode=n,placeholder:l.$t("model.placeholder.image_quotas.mode"),style:{width:"185px","margin-right":"5px"}},null,8,["modelValue","onUpdate:modelValue","placeholder"]),o(b,{modelValue:e.value.multimodal_quota.image_quotas[r].fixed_quota,"onUpdate:modelValue":n=>e.value.multimodal_quota.image_quotas[r].fixed_quota=n,placeholder:l.$t("model.placeholder.image_quotas.fixed_quota"),style:{width:"185px","margin-right":"5px"}},null,8,["modelValue","onUpdate:modelValue","placeholder"]),o(M,{modelValue:e.value.multimodal_quota.image_quotas[r].is_default,"onUpdate:modelValue":n=>e.value.multimodal_quota.image_quotas[r].is_default=n,value:"1",style:{width:"60px"},onChange:n=>ne(r)},{default:t(()=>[_("\u9ED8\u8BA4")]),_:2},1032,["modelValue","onUpdate:modelValue","onChange"]),o(w,{type:"primary",shape:"circle",style:{margin:"0 10px 0 10px"},onClick:a[17]||(a[17]=n=>W())},{default:t(()=>[o(N)]),_:1}),o(w,{type:"secondary",shape:"circle",onClick:n=>me(r)},{default:t(()=>[o(O)]),_:2},1032,["onClick"])]),_:2},1032,["field","label","rules"])),[[P,S.value]])),128)),(m(!0),h($,null,k(e.value.midjourney_quotas,(u,r)=>T((m(),s(i,{key:r,field:`midjourney_quotas[${r}].name`&&`midjourney_quotas[${r}].action`&&`midjourney_quotas[${r}].path`&&`midjourney_quotas[${r}].fixed_quota`,label:`${r+1}. `+l.$t("model.label.midjourney_quotas"),rules:[{required:!0,message:l.$t("model.error.midjourney_quotas.required")}]},{default:t(()=>[o(U,{modelValue:e.value.midjourney_quotas[r].name,"onUpdate:modelValue":n=>e.value.midjourney_quotas[r].name=n,placeholder:l.$t("model.placeholder.midjourney_quotas.name"),style:{width:"95px","margin-right":"5px"}},null,8,["modelValue","onUpdate:modelValue","placeholder"]),o(U,{modelValue:e.value.midjourney_quotas[r].action,"onUpdate:modelValue":n=>e.value.midjourney_quotas[r].action=n,placeholder:l.$t("model.placeholder.midjourney_quotas.action"),style:{width:"102px","margin-right":"5px"}},null,8,["modelValue","onUpdate:modelValue","placeholder"]),o(U,{modelValue:e.value.midjourney_quotas[r].path,"onUpdate:modelValue":n=>e.value.midjourney_quotas[r].path=n,placeholder:l.$t("model.placeholder.midjourney_quotas.path"),style:{width:"138px","margin-right":"5px"}},null,8,["modelValue","onUpdate:modelValue","placeholder"]),o(b,{modelValue:e.value.midjourney_quotas[r].fixed_quota,"onUpdate:modelValue":n=>e.value.midjourney_quotas[r].fixed_quota=n,placeholder:l.$t("model.placeholder.midjourney_quotas.fixed_quota"),style:{width:"90px"}},null,8,["modelValue","onUpdate:modelValue","placeholder"]),o(w,{type:"primary",shape:"circle",style:{margin:"0 10px 0 10px"},onClick:a[18]||(a[18]=n=>X())},{default:t(()=>[o(N)]),_:1}),o(w,{type:"secondary",shape:"circle",onClick:n=>se(r)},{default:t(()=>[o(O)]),_:2},1032,["onClick"])]),_:2},1032,["field","label","rules"])),[[P,E.value]])),128)),o(i,{field:"data_format",label:l.$t("model.label.dataFormat"),rules:[{required:!0,message:l.$t("model.error.dataFormat.required")}]},{default:t(()=>[o(R,{size:"large"},{default:t(()=>[o(M,{modelValue:e.value.data_format,"onUpdate:modelValue":a[19]||(a[19]=u=>e.value.data_format=u),value:"1","default-checked":!0},{default:t(()=>[_("\u7EDF\u4E00\u683C\u5F0F")]),_:1},8,["modelValue"])]),_:1})]),_:1},8,["label","rules"]),o(i,{field:"is_public",label:l.$t("model.label.isPublic"),rules:[{required:!0}]},{default:t(()=>[o(j,{modelValue:e.value.is_public,"onUpdate:modelValue":a[20]||(a[20]=u=>e.value.is_public=u)},null,8,["modelValue"])]),_:1},8,["label"]),o(i,{field:"is_enable_preset_config",label:l.$t("model.label.is_enable_preset_config")},{default:t(()=>[o(j,{modelValue:e.value.is_enable_preset_config,"onUpdate:modelValue":a[21]||(a[21]=u=>e.value.is_enable_preset_config=u)},null,8,["modelValue"])]),_:1},8,["label"]),e.value.is_enable_preset_config?(m(),s(i,{key:8,field:"preset_config.is_support_system_role",label:l.$t("model.label.preset_config.is_support_system_role")},{default:t(()=>[o(j,{modelValue:e.value.preset_config.is_support_system_role,"onUpdate:modelValue":a[22]||(a[22]=u=>e.value.preset_config.is_support_system_role=u)},null,8,["modelValue"])]),_:1},8,["label"])):p("",!0),e.value.is_enable_preset_config&&e.value.preset_config.is_support_system_role?(m(),s(i,{key:9,field:"preset_config.system_role_prompt",label:l.$t("model.label.preset_config.system_role_prompt")},{default:t(()=>[o(x,{modelValue:e.value.preset_config.system_role_prompt,"onUpdate:modelValue":a[23]||(a[23]=u=>e.value.preset_config.system_role_prompt=u),placeholder:l.$t("model.placeholder.preset_config.system_role_prompt")},null,8,["modelValue","placeholder"])]),_:1},8,["label"])):p("",!0),e.value.is_enable_preset_config?(m(),s(i,{key:10,field:"preset_config.max_tokens",label:l.$t("model.label.preset_config.max_tokens.range")},{default:t(()=>[o(b,{modelValue:e.value.preset_config.min_tokens,"onUpdate:modelValue":a[24]||(a[24]=u=>e.value.preset_config.min_tokens=u),placeholder:l.$t("model.placeholder.preset_config.min_tokens"),style:{width:"260px","margin-right":"5px"},min:0,max:2097152},null,8,["modelValue","placeholder"]),o(b,{modelValue:e.value.preset_config.max_tokens,"onUpdate:modelValue":a[25]||(a[25]=u=>e.value.preset_config.max_tokens=u),placeholder:l.$t("model.placeholder.preset_config.max_tokens"),style:{width:"260px"},min:0,max:2097152},null,8,["modelValue","placeholder"])]),_:1},8,["field","label"])):p("",!0),o(i,{field:"is_enable_model_agent",label:l.$t("model.label.isEnableModelAgent")},{default:t(()=>[o(j,{modelValue:e.value.is_enable_model_agent,"onUpdate:modelValue":a[26]||(a[26]=u=>e.value.is_enable_model_agent=u)},null,8,["modelValue"])]),_:1},8,["label"]),e.value.is_enable_model_agent?(m(),s(i,{key:11,field:"model_agents",label:l.$t("model.label.modelAgents"),rules:[{required:!0,message:l.$t("model.error.modelAgents.required")}]},{default:t(()=>[o(d,{modelValue:e.value.model_agents,"onUpdate:modelValue":a[27]||(a[27]=u=>e.value.model_agents=u),placeholder:l.$t("model.placeholder.modelAgents"),"max-tag-count":3,multiple:"","allow-search":"","allow-clear":""},{default:t(()=>[(m(!0),h($,null,k(K.value,u=>(m(),s(g,{key:u.id,value:u.id,label:u.name},null,8,["value","label"]))),128))]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label","rules"])):p("",!0),o(i,{field:"model_forward",label:l.$t("model.label.modelForward")},{default:t(()=>[o(j,{modelValue:e.value.is_enable_forward,"onUpdate:modelValue":a[28]||(a[28]=u=>e.value.is_enable_forward=u),onChange:Y},null,8,["modelValue"])]),_:1},8,["label"]),e.value.is_enable_forward?(m(),s(i,{key:12,field:"forward_config.forward_rule",label:l.$t("model.label.forwardRule"),rules:[{required:!0,message:l.$t("model.error.forwardRule.required")}]},{default:t(()=>[o(d,{modelValue:e.value.forward_config.forward_rule,"onUpdate:modelValue":a[29]||(a[29]=u=>e.value.forward_config.forward_rule=u),placeholder:l.$t("model.placeholder.forwardRule"),onChange:Y},{default:t(()=>[o(g,{value:"1"},{default:t(()=>[_("\u5168\u90E8\u8F6C\u53D1")]),_:1}),o(g,{value:"2"},{default:t(()=>[_("\u6309\u5173\u952E\u5B57")]),_:1}),o(g,{value:"3"},{default:t(()=>[_("\u5185\u5BB9\u957F\u5EA6")]),_:1})]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label","rules"])):p("",!0),e.value.is_enable_forward&&e.value.forward_config.forward_rule==="3"?(m(),s(i,{key:13,field:"forward_config.content_length",label:l.$t("model.label.content_length"),rules:[{required:!0,message:l.$t("model.error.content_length.required")}]},{default:t(()=>[o(b,{modelValue:e.value.forward_config.content_length,"onUpdate:modelValue":a[30]||(a[30]=u=>e.value.forward_config.content_length=u),min:1,max:9999999999999,placeholder:l.$t("model.placeholder.content_length")},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"])):p("",!0),e.value.is_enable_forward&&(e.value.forward_config.forward_rule==="1"||e.value.forward_config.forward_rule==="3")?(m(),s(i,{key:14,field:"forward_config.target_model",label:l.$t("model.label.targetModel"),rules:[{required:!0,message:l.$t("model.error.targetModel.required")}]},{default:t(()=>[o(d,{modelValue:e.value.forward_config.target_model,"onUpdate:modelValue":a[31]||(a[31]=u=>e.value.forward_config.target_model=u),placeholder:l.$t("model.placeholder.targetModel"),"allow-search":""},{default:t(()=>[(m(!0),h($,null,k(F.value,u=>(m(),s(g,{key:u.id,value:u.id,label:u.name},null,8,["value","label"]))),128))]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label","rules"])):p("",!0),e.value.is_enable_forward&&e.value.forward_config.forward_rule==="2"?(m(),s(i,{key:15,field:"forward_config.match_rule",label:l.$t("model.label.matchRule"),rules:[{required:!0,message:l.$t("model.error.matchRule.required")}]},{default:t(()=>[o(R,{size:"large"},{default:t(()=>[o(ee,{modelValue:e.value.forward_config.match_rule,"onUpdate:modelValue":a[32]||(a[32]=u=>e.value.forward_config.match_rule=u),value:"1","default-checked":!0},{default:t(()=>[_("\u667A\u80FD\u5339\u914D")]),_:1},8,["modelValue"]),o(ee,{modelValue:e.value.forward_config.match_rule,"onUpdate:modelValue":a[33]||(a[33]=u=>e.value.forward_config.match_rule=u),value:"2"},{default:t(()=>[_("\u6B63\u5219\u5339\u914D")]),_:1},8,["modelValue"])]),_:1})]),_:1},8,["label","rules"])):p("",!0),e.value.is_enable_forward&&e.value.forward_config.forward_rule==="2"&&e.value.forward_config.match_rule.includes("1")?(m(),s(i,{key:16,field:"forward_config.decision_model",label:l.$t("model.label.decisionModel"),rules:[{required:!0,message:l.$t("model.error.decisionModel.required")}]},{default:t(()=>[o(d,{modelValue:e.value.forward_config.decision_model,"onUpdate:modelValue":a[34]||(a[34]=u=>e.value.forward_config.decision_model=u),placeholder:l.$t("model.placeholder.decisionModel"),"allow-search":""},{default:t(()=>[(m(!0),h($,null,k(F.value,u=>(m(),s(g,{key:u.id,value:u.id,label:u.name},null,8,["value","label"]))),128))]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label","rules"])):p("",!0),(m(!0),h($,null,k(e.value.forward_config.keywords,(u,r)=>T((m(),s(i,{key:r,field:`forward_config.keywords[${r}]`&&`forward_config.target_models[${r}]`,label:`${r+1}. `+l.$t("model.label.keywords"),rules:[{required:!0,message:l.$t("model.error.keywordsAndtargetModel.required")}]},{default:t(()=>[o(U,{modelValue:e.value.forward_config.keywords[r],"onUpdate:modelValue":n=>e.value.forward_config.keywords[r]=n,placeholder:l.$t("model.placeholder.keywords"),style:{width:"45%","margin-right":"5px"}},null,8,["modelValue","onUpdate:modelValue","placeholder"]),o(d,{modelValue:e.value.forward_config.target_models[r],"onUpdate:modelValue":n=>e.value.forward_config.target_models[r]=n,placeholder:l.$t("model.placeholder.targetModel"),style:{width:"40%"},"allow-search":""},{default:t(()=>[(m(!0),h($,null,k(F.value,n=>(m(),s(g,{key:n.id,value:n.id,label:n.name},null,8,["value","label"]))),128))]),_:2},1032,["modelValue","onUpdate:modelValue","placeholder"]),o(w,{type:"primary",shape:"circle",style:{margin:"0 10px 0 10px"},onClick:_e},{default:t(()=>[o(N)]),_:1}),o(w,{type:"secondary",shape:"circle",onClick:n=>fe(r)},{default:t(()=>[o(O)]),_:2},1032,["onClick"])]),_:2},1032,["field","label","rules"])),[[P,e.value.is_enable_forward&&e.value.forward_config.forward_rule==="2"]])),128)),o(i,{field:"is_enable_fallback",label:l.$t("model.label.is_enable_fallback")},{default:t(()=>[o(j,{modelValue:e.value.is_enable_fallback,"onUpdate:modelValue":a[35]||(a[35]=u=>e.value.is_enable_fallback=u)},null,8,["modelValue"])]),_:1},8,["label"]),e.value.is_enable_fallback?(m(),s(i,{key:17,field:"fallback_config.fallback_model",label:l.$t("model.label.fallback_model"),rules:[{required:!0,message:l.$t("model.error.fallback_model.required")}]},{default:t(()=>[o(d,{modelValue:e.value.fallback_config.fallback_model,"onUpdate:modelValue":a[36]||(a[36]=u=>e.value.fallback_config.fallback_model=u),placeholder:l.$t("model.placeholder.fallback_model"),"allow-search":""},{default:t(()=>[(m(!0),h($,null,k(F.value,u=>(m(),s(g,{key:u.id,value:u.id,label:u.name},null,8,["value","label"]))),128))]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label","rules"])):p("",!0),o(R,null,{default:t(()=>[A("div",Ze,[o(w,{type:"secondary",onClick:a[37]||(a[37]=u=>l.$router.push({name:"ModelList"}))},{default:t(()=>[_(V(l.$t("form.button.cancel")),1)]),_:1}),o(w,{type:"primary",onClick:te},{default:t(()=>[_(V(l.$t("form.button.submit")),1)]),_:1})])]),_:1})]),_:1},8,["model"])])]),_:1})]),_:1},8,["loading"])])}}});const kl=be(Ye,[["__scopeId","data-v-3e8722cc"]]);export{kl as default};
