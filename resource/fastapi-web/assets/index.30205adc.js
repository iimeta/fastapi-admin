import{_ as T,E as ae}from"./index.36ee875b.js";/* empty css               *//* empty css                */import{d as K,e as M,B as s,aD as f,aG as a,aH as l,aL as n,aM as V,bB as Q,bC as j,b2 as J,b1 as X,bK as oe,aU as O,b5 as G,aE as k,C as S,aJ as E,aI as L,bu as re,bv as de,bX as te,bi as H,aS as ue,aT as ne,a5 as se,a4 as me,bz as ie,F as fe,bs as _e,u as pe,aK as ce,aF as ve,bL as ge,bM as be,bJ as we,bN as ye}from"./arco.aed15247.js";import{u as P}from"./loading.b5911e1d.js";import{d as Z,q as Ve,e as $e}from"./model.c407852e.js";/* empty css               *//* empty css               *//* empty css                *//* empty css              *//* empty css              *//* empty css               */import{h as W}from"./vue.0cc5b64a.js";/* empty css               *//* empty css                */import{e as he}from"./agent.4f0c5cb9.js";/* empty css               */import"./chart.9aa6eafa.js";import"./base.87fcf6e2.js";const ke=K({__name:"base-info",emits:["changeStep"],setup(y,{emit:w}){const{setLoading:g}=P(!0),p=W(),b=M(),t=M({id:"",corp:"",name:"",model:"",type:"",remark:"",base_url:"",path:"",prompt:"",status:1});(async(u={id:p.query.id})=>{g(!0);try{const{data:e}=await Z(u);t.value.id=e.id,t.value.corp=e.corp,t.value.name=e.name,t.value.model=e.model,t.value.type=String(e.type),t.value.remark=e.remark,t.value.base_url=e.base_url,t.value.path=e.path,t.value.prompt=e.prompt,t.value.status=e.status}catch{}finally{g(!1)}})();const i=async()=>{var e;await((e=b.value)==null?void 0:e.validate())||w("changeStep","forward",{...t.value})};return(u,e)=>{const c=Q,C=j,v=J,F=X,D=oe,R=O,r=G;return s(),f(r,{ref_key:"formRef",ref:b,model:t.value,class:"form","label-col-props":{span:5},"wrapper-col-props":{span:18}},{default:a(()=>[l(v,{field:"corp",label:u.$t("model.label.corp"),rules:[{required:!0,message:u.$t("model.error.corp.required")}]},{default:a(()=>[l(C,{modelValue:t.value.corp,"onUpdate:modelValue":e[0]||(e[0]=o=>t.value.corp=o),placeholder:u.$t("model.placeholder.corp")},{default:a(()=>[l(c,{value:"OpenAI"},{default:a(()=>[n("OpenAI")]),_:1}),l(c,{value:"Baidu"},{default:a(()=>[n("\u767E\u5EA6")]),_:1}),l(c,{value:"Xfyun"},{default:a(()=>[n("\u79D1\u5927\u8BAF\u98DE")]),_:1}),l(c,{value:"Aliyun"},{default:a(()=>[n("\u963F\u91CC\u4E91")]),_:1}),l(c,{value:"ZhipuAI"},{default:a(()=>[n("\u667A\u8C31AI")]),_:1}),l(c,{value:"Midjourney"},{default:a(()=>[n("Midjourney")]),_:1})]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),l(v,{field:"name",label:u.$t("model.label.name"),rules:[{required:!0,message:u.$t("model.error.name.required")}]},{default:a(()=>[l(F,{modelValue:t.value.name,"onUpdate:modelValue":e[1]||(e[1]=o=>t.value.name=o),placeholder:u.$t("model.placeholder.name")},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),l(v,{field:"model",label:u.$t("model.label.model"),rules:[{required:!0,message:u.$t("model.error.model.required")},{message:u.$t("model.error.model.pattern")}]},{default:a(()=>[l(F,{modelValue:t.value.model,"onUpdate:modelValue":e[2]||(e[2]=o=>t.value.model=o),placeholder:u.$t("model.placeholder.model")},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),l(v,{field:"type",label:u.$t("model.label.type"),rules:[{required:!0,message:u.$t("model.error.type.required")}]},{default:a(()=>[l(C,{modelValue:t.value.type,"onUpdate:modelValue":e[3]||(e[3]=o=>t.value.type=o),placeholder:u.$t("model.placeholder.type")},{default:a(()=>[l(c,{value:"1"},{default:a(()=>[n("\u6587\u751F\u6587")]),_:1}),l(c,{value:"2"},{default:a(()=>[n("\u6587\u751F\u56FE")]),_:1}),l(c,{value:"3"},{default:a(()=>[n("\u56FE\u751F\u6587")]),_:1}),l(c,{value:"4"},{default:a(()=>[n("\u56FE\u751F\u56FE")]),_:1})]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),l(v,{field:"base_url",label:u.$t("model.label.base_url")},{default:a(()=>[l(F,{modelValue:t.value.base_url,"onUpdate:modelValue":e[4]||(e[4]=o=>t.value.base_url=o),placeholder:u.$t("model.placeholder.base_url")},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),l(v,{field:"path",label:u.$t("model.label.path")},{default:a(()=>[l(F,{modelValue:t.value.path,"onUpdate:modelValue":e[5]||(e[5]=o=>t.value.path=o),placeholder:u.$t("model.placeholder.path")},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),l(v,{field:"prompt",label:u.$t("model.label.prompt")},{default:a(()=>[l(D,{modelValue:t.value.prompt,"onUpdate:modelValue":e[6]||(e[6]=o=>t.value.prompt=o),placeholder:u.$t("model.placeholder.prompt")},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),l(v,{field:"remark",label:u.$t("model.label.remark")},{default:a(()=>[l(D,{modelValue:t.value.remark,"onUpdate:modelValue":e[7]||(e[7]=o=>t.value.remark=o),placeholder:u.$t("model.placeholder.remark")},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),l(v,null,{default:a(()=>[l(R,{type:"primary",onClick:i},{default:a(()=>[n(V(u.$t("model.button.next")),1)]),_:1})]),_:1})]),_:1},8,["model"])}}});const qe=T(ke,[["__scopeId","data-v-7d50c5d3"]]),Me=K({__name:"advanced",emits:["changeStep"],setup(y,{emit:w}){const{setLoading:g}=P(!0),p=W(),b=M([]);(async()=>{g(!0);try{const{data:r}=await Ve();b.value=r.items}catch{}finally{g(!1)}})();const U=M([]);(async()=>{g(!0);try{const{data:r}=await he();U.value=r.items}catch{}finally{g(!1)}})();const u=M(),e=M({billing_method:"1",prompt_ratio:1,completion_ratio:1,fixed_quota:1,data_format:"",is_public:!0,is_enable_model_agent:!1,model_agents:[],is_forward:!1,forward_config:{forward_rule:"1",match_rule:"1",target_model:"",decision_model:"",keywords:[],target_models:[]}}),c=()=>{e.value.forward_config.keywords.push(""),e.value.forward_config.target_models.push("")},C=()=>{!e.value.is_forward&&e.value.forward_config.keywords&&(e.value.forward_config.keywords[0]===""||e.value.forward_config.target_models[0]==="")?(e.value.forward_config.keywords=[],e.value.forward_config.target_models=[]):e.value.forward_config.forward_rule==="2"&&(!e.value.forward_config.keywords||e.value.forward_config.keywords.length===0)?(e.value.forward_config.keywords=[""],e.value.forward_config.target_models=[""]):e.value.forward_config.forward_rule==="1"&&e.value.forward_config.keywords&&(e.value.forward_config.keywords[0]===""||e.value.forward_config.target_models[0]==="")&&(e.value.forward_config.keywords=[],e.value.forward_config.target_models=[])},v=r=>{e.value.forward_config.keywords.length>1&&(e.value.forward_config.keywords.splice(r,1),e.value.forward_config.target_models.splice(r,1))};(async(r={id:p.query.id})=>{var o,$,A,_;g(!0);try{const{data:m}=await Z(r);e.value.billing_method=String(m.billing_method),e.value.prompt_ratio=m.prompt_ratio,e.value.completion_ratio=m.completion_ratio,e.value.fixed_quota=m.fixed_quota,e.value.data_format=String(m.data_format),e.value.is_public=m.is_public,e.value.is_enable_model_agent=m.is_enable_model_agent,e.value.model_agents=m.model_agents,e.value.is_forward=m.is_forward,m.is_forward&&m.forward_config&&(e.value.forward_config.forward_rule=String(m.forward_config.forward_rule),e.value.forward_config.match_rule=String(m.forward_config.match_rule),e.value.forward_config.target_model=(o=m.forward_config)==null?void 0:o.target_model,e.value.forward_config.decision_model=($=m.forward_config)==null?void 0:$.decision_model,e.value.forward_config.keywords=(A=m.forward_config)==null?void 0:A.keywords,e.value.forward_config.target_models=(_=m.forward_config)==null?void 0:_.target_models)}catch{}finally{g(!1)}})();const D=async()=>{var o;await((o=u.value)==null?void 0:o.validate())||w("changeStep","submit",{...e.value})},R=()=>{w("changeStep","backward")};return(r,o)=>{const $=te,A=H,_=J,m=ue,z=ne,I=Q,B=j,Y=X,x=se,N=O,ee=me,le=G;return s(),f(le,{ref_key:"formRef",ref:u,model:e.value,class:"form","label-col-props":{span:5},"wrapper-col-props":{span:18}},{default:a(()=>[l(_,{field:"billing_method",label:r.$t("model.label.billingMethod"),rules:[{required:!0,message:r.$t("model.error.billingMethod.required")}]},{default:a(()=>[l(A,{size:"large"},{default:a(()=>[l($,{modelValue:e.value.billing_method,"onUpdate:modelValue":o[0]||(o[0]=d=>e.value.billing_method=d),value:"1"},{default:a(()=>[n("\u500D\u7387")]),_:1},8,["modelValue"]),l($,{modelValue:e.value.billing_method,"onUpdate:modelValue":o[1]||(o[1]=d=>e.value.billing_method=d),value:"2"},{default:a(()=>[n("\u56FA\u5B9A\u989D\u5EA6")]),_:1},8,["modelValue"])]),_:1})]),_:1},8,["label","rules"]),e.value.billing_method==="1"?(s(),f(_,{key:0,field:"prompt_ratio",label:r.$t("model.label.promptRatio"),rules:[{required:!0,message:r.$t("model.error.promptRatio.required")}]},{default:a(()=>[l(m,{modelValue:e.value.prompt_ratio,"onUpdate:modelValue":o[2]||(o[2]=d=>e.value.prompt_ratio=d),min:.001,placeholder:r.$t("model.placeholder.promptRatio")},null,8,["modelValue","min","placeholder"])]),_:1},8,["label","rules"])):k("",!0),e.value.billing_method==="1"?(s(),f(_,{key:1,field:"completion_ratio",label:r.$t("model.label.completionRatio"),rules:[{required:!0,message:r.$t("model.error.completionRatio.required")}]},{default:a(()=>[l(m,{modelValue:e.value.completion_ratio,"onUpdate:modelValue":o[3]||(o[3]=d=>e.value.completion_ratio=d),min:.001,placeholder:r.$t("model.placeholder.completionRatio")},null,8,["modelValue","min","placeholder"])]),_:1},8,["label","rules"])):k("",!0),e.value.billing_method==="2"?(s(),f(_,{key:2,field:"fixed_quota",label:r.$t("model.label.fixedQuota"),rules:[{required:!0,message:r.$t("model.error.fixedQuota.required")}]},{default:a(()=>[l(m,{modelValue:e.value.fixed_quota,"onUpdate:modelValue":o[4]||(o[4]=d=>e.value.fixed_quota=d),min:0,max:9999999999999,placeholder:r.$t("model.placeholder.fixedQuota")},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"])):k("",!0),l(_,{field:"data_format",label:r.$t("model.label.dataFormat"),rules:[{required:!0,message:r.$t("model.error.dataFormat.required")}]},{default:a(()=>[l(A,{size:"large"},{default:a(()=>[l($,{modelValue:e.value.data_format,"onUpdate:modelValue":o[5]||(o[5]=d=>e.value.data_format=d),value:"1"},{default:a(()=>[n("\u7EDF\u4E00\u683C\u5F0F")]),_:1},8,["modelValue"]),l($,{modelValue:e.value.data_format,"onUpdate:modelValue":o[6]||(o[6]=d=>e.value.data_format=d),value:"2"},{default:a(()=>[n("\u5B98\u65B9\u683C\u5F0F")]),_:1},8,["modelValue"])]),_:1})]),_:1},8,["label","rules"]),l(_,{field:"is_public",label:r.$t("model.label.isPublic"),rules:[{required:!0}]},{default:a(()=>[l(z,{modelValue:e.value.is_public,"onUpdate:modelValue":o[7]||(o[7]=d=>e.value.is_public=d)},null,8,["modelValue"])]),_:1},8,["label"]),l(_,{field:"is_enable_model_agent",label:r.$t("model.label.isEnableModelAgent")},{default:a(()=>[l(z,{modelValue:e.value.is_enable_model_agent,"onUpdate:modelValue":o[8]||(o[8]=d=>e.value.is_enable_model_agent=d)},null,8,["modelValue"])]),_:1},8,["label"]),e.value.is_enable_model_agent?(s(),f(_,{key:3,field:"model_agents",label:r.$t("model.label.modelAgents"),rules:[{required:!0,message:r.$t("model.error.modelAgents.required")}]},{default:a(()=>[l(B,{modelValue:e.value.model_agents,"onUpdate:modelValue":o[9]||(o[9]=d=>e.value.model_agents=d),placeholder:r.$t("model.placeholder.modelAgents"),"max-tag-count":3,multiple:"","allow-clear":""},{default:a(()=>[(s(!0),S(E,null,L(U.value,d=>(s(),f(I,{key:d.id,value:d.id,label:d.name},null,8,["value","label"]))),128))]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label","rules"])):k("",!0),l(_,{field:"model_forward",label:r.$t("model.label.modelForward")},{default:a(()=>[l(z,{modelValue:e.value.is_forward,"onUpdate:modelValue":o[10]||(o[10]=d=>e.value.is_forward=d),onChange:C},null,8,["modelValue"])]),_:1},8,["label"]),e.value.is_forward?(s(),f(_,{key:4,field:"forward_config.forward_rule",label:r.$t("model.label.forwardRule"),rules:[{required:!0,message:r.$t("model.error.forwardRule.required")}]},{default:a(()=>[l(B,{modelValue:e.value.forward_config.forward_rule,"onUpdate:modelValue":o[11]||(o[11]=d=>e.value.forward_config.forward_rule=d),placeholder:r.$t("model.placeholder.forwardRule"),onChange:C},{default:a(()=>[l(I,{value:"1"},{default:a(()=>[n("\u5168\u90E8\u8F6C\u53D1")]),_:1}),l(I,{value:"2"},{default:a(()=>[n("\u6309\u5173\u952E\u5B57")]),_:1})]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label","rules"])):k("",!0),e.value.is_forward&&e.value.forward_config.forward_rule==="1"?(s(),f(_,{key:5,field:"forward_config.target_model",label:r.$t("model.label.targetModel"),rules:[{required:!0,message:r.$t("model.error.targetModel.required")}]},{default:a(()=>[l(B,{modelValue:e.value.forward_config.target_model,"onUpdate:modelValue":o[12]||(o[12]=d=>e.value.forward_config.target_model=d),placeholder:r.$t("model.placeholder.targetModel")},{default:a(()=>[(s(!0),S(E,null,L(b.value,d=>(s(),f(I,{key:d.id,value:d.id,label:d.name},null,8,["value","label"]))),128))]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label","rules"])):k("",!0),e.value.is_forward&&e.value.forward_config.forward_rule==="2"?(s(),f(_,{key:6,field:"forward_config.match_rule",label:r.$t("model.label.matchRule"),rules:[{required:!0,message:r.$t("model.error.matchRule.required")}]},{default:a(()=>[l(A,{size:"large"},{default:a(()=>[l($,{modelValue:e.value.forward_config.match_rule,"onUpdate:modelValue":o[13]||(o[13]=d=>e.value.forward_config.match_rule=d),value:"1","default-checked":!0},{default:a(()=>[n("\u667A\u80FD\u5339\u914D")]),_:1},8,["modelValue"]),l($,{modelValue:e.value.forward_config.match_rule,"onUpdate:modelValue":o[14]||(o[14]=d=>e.value.forward_config.match_rule=d),value:"2"},{default:a(()=>[n("\u6B63\u5219\u5339\u914D")]),_:1},8,["modelValue"])]),_:1})]),_:1},8,["label","rules"])):k("",!0),e.value.is_forward&&e.value.forward_config.forward_rule==="2"&&e.value.forward_config.match_rule==="1"?(s(),f(_,{key:7,field:"forward_config.decision_model",label:r.$t("model.label.decisionModel"),rules:[{required:!0,message:r.$t("model.error.decisionModel.required")}]},{default:a(()=>[l(B,{modelValue:e.value.forward_config.decision_model,"onUpdate:modelValue":o[15]||(o[15]=d=>e.value.forward_config.decision_model=d),placeholder:r.$t("model.placeholder.decisionModel")},{default:a(()=>[(s(!0),S(E,null,L(b.value,d=>(s(),f(I,{key:d.id,value:d.id,label:d.name},null,8,["value","label"]))),128))]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label","rules"])):k("",!0),(s(!0),S(E,null,L(e.value.forward_config.keywords,(d,h)=>re((s(),f(_,{key:h,field:`forward_config.keywords[${h}]`&&`forward_config.target_models[${h}]`,label:`${h+1}. `+r.$t("model.label.keywords"),rules:[{required:!0,message:r.$t("model.error.keywordsAndtargetModel.required")}]},{default:a(()=>[l(Y,{modelValue:e.value.forward_config.keywords[h],"onUpdate:modelValue":q=>e.value.forward_config.keywords[h]=q,placeholder:r.$t("model.placeholder.keywords"),style:{width:"40%","margin-right":"5px"}},null,8,["modelValue","onUpdate:modelValue","placeholder"]),l(B,{modelValue:e.value.forward_config.target_models[h],"onUpdate:modelValue":q=>e.value.forward_config.target_models[h]=q,placeholder:r.$t("model.placeholder.targetModel"),style:{width:"40%"}},{default:a(()=>[(s(!0),S(E,null,L(b.value,q=>(s(),f(I,{key:q.id,value:q.id,label:q.name},null,8,["value","label"]))),128))]),_:2},1032,["modelValue","onUpdate:modelValue","placeholder"]),l(N,{type:"primary",shape:"circle",style:{margin:"0 10px 0 10px"},onClick:c},{default:a(()=>[l(x)]),_:1}),l(N,{type:"secondary",shape:"circle",onClick:q=>v(h)},{default:a(()=>[l(ee)]),_:2},1032,["onClick"])]),_:2},1032,["field","label","rules"])),[[de,e.value.is_forward&&e.value.forward_config.forward_rule==="2"]])),128)),l(_,null,{default:a(()=>[l(A,null,{default:a(()=>[l(N,{type:"secondary",onClick:R},{default:a(()=>[n(V(r.$t("model.button.prev")),1)]),_:1}),l(N,{type:"primary",onClick:D},{default:a(()=>[n(V(r.$t("model.button.next")),1)]),_:1})]),_:1})]),_:1})]),_:1},8,["model"])}}});const Ue=T(Me,[["__scopeId","data-v-498c8f7d"]]);const Ce={},Fe={class:"success-wrap"};function Ae(y,w){const g=ie,p=O,b=H;return s(),S("div",Fe,[l(g,{status:"success",title:y.$t("model.success.title"),subtitle:y.$t("model.success.update.subTitle")},null,8,["title","subtitle"]),l(b,{size:16},{default:a(()=>[l(p,{key:"finish",type:"secondary",onClick:w[0]||(w[0]=t=>y.$router.push({name:"ModelList"}))},{default:a(()=>[n(V(y.$t("model.button.return")),1)]),_:1}),l(p,{key:"again",type:"primary",onClick:w[1]||(w[1]=t=>y.$router.push({name:"ModelDetail",query:{id:`${y.$route.query.id}`}}))},{default:a(()=>[n(V(y.$t("model.button.view")),1)]),_:1})]),_:1})])}const Se=T(Ce,[["render",Ae],["__scopeId","data-v-829f16db"]]),De={class:"container"},Ie={class:"wrapper"},Re={name:"ModelUpdate"},Be=K({...Re,setup(y){const{loading:w,setLoading:g}=P(!1),p=M(1),b=M({}),t=async()=>{g(!0);try{await $e(b.value),p.value=3,b.value={}}catch{}finally{g(!1)}},U=(i,u)=>{if(typeof i=="number"){p.value=i;return}if(i==="forward"||i==="submit"){if(b.value={...b.value,...u},i==="submit"){t();return}p.value+=1}else i==="backward"&&(p.value-=1)};return(i,u)=>{const e=ae,c=ce,C=ve,v=ge,F=be,D=we,R=ye;return s(),S("div",De,[l(C,{class:"container-breadcrumb"},{default:a(()=>[l(c,null,{default:a(()=>[l(e)]),_:1}),l(c,null,{default:a(()=>[n(V(i.$t("menu.model")),1)]),_:1}),l(c,null,{default:a(()=>[n(V(i.$t("menu.model.update")),1)]),_:1})]),_:1}),l(R,{loading:pe(w),style:{width:"100%"}},{default:a(()=>[l(D,{class:"general-card",bordered:!1},{default:a(()=>[fe("div",Ie,[l(F,{current:p.value,"onUpdate:current":u[0]||(u[0]=r=>p.value=r),style:{width:"660px"},"line-less":"",class:"steps"},{default:a(()=>[l(v,{description:i.$t("model.subTitle.baseInfo")},{default:a(()=>[n(V(i.$t("model.title.baseInfo")),1)]),_:1},8,["description"]),l(v,{description:i.$t("model.subTitle.advanced")},{default:a(()=>[n(V(i.$t("model.title.advanced")),1)]),_:1},8,["description"]),l(v,{description:i.$t("model.subTitle.update.finish")},{default:a(()=>[n(V(i.$t("model.title.update.finish")),1)]),_:1},8,["description"])]),_:1},8,["current"]),(s(),f(_e,null,[p.value===1?(s(),f(qe,{key:0,onChangeStep:U})):p.value===2?(s(),f(Ue,{key:1,onChangeStep:U})):p.value===3?(s(),f(Se,{key:2,onChangeStep:U})):k("",!0)],1024))])]),_:1})]),_:1},8,["loading"])])}}});const ll=T(Be,[["__scopeId","data-v-18d55c77"]]);export{ll as default};
