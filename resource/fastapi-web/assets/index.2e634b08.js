import{_ as F,E as j}from"./index.dc70f908.js";/* empty css               *//* empty css                */import{d as C,e as h,B as g,aD as $,aG as l,aH as e,aL as n,aM as v,bB as A,bC as R,b2 as B,b1 as G,bK as J,aU as S,b5 as L,C as M,aJ as P,aI as H,aE as D,aS as X,bX as Q,bi as E,aT as W,bz as Y,F as Z,bs as x,u as ee,aK as le,aF as ae,bL as te,bM as oe,bJ as re,bN as ue}from"./arco.aed15247.js";import{u as I}from"./loading.b5911e1d.js";import{d as N,e as de}from"./model.4c1fb15d.js";/* empty css               *//* empty css               *//* empty css                *//* empty css              *//* empty css              *//* empty css               */import{h as T}from"./vue.0cc5b64a.js";/* empty css               *//* empty css                */import{e as ne}from"./agent.4d9ba6a8.js";/* empty css               */import"./chart.9aa6eafa.js";import"./base.87fcf6e2.js";const se=C({__name:"base-info",emits:["changeStep"],setup(y,{emit:b}){const{setLoading:_}=I(!0),m=T(),c=h(),r=h({id:"",corp:"",name:"",model:"",type:"",remark:"",prompt:"",status:1});(async(u={id:m.query.id})=>{_(!0);try{const{data:o}=await N(u);r.value.id=o.id,r.value.corp=o.corp,r.value.name=o.name,r.value.model=o.model,r.value.type=String(o.type),r.value.remark=o.remark,r.value.prompt=o.prompt,r.value.status=o.status}catch{}finally{_(!1)}})();const a=async()=>{var o;await((o=c.value)==null?void 0:o.validate())||b("changeStep","forward",{...r.value})};return(u,o)=>{const i=A,d=R,t=B,k=G,f=J,q=S,w=L;return g(),$(w,{ref_key:"formRef",ref:c,model:r.value,class:"form","label-col-props":{span:5},"wrapper-col-props":{span:18}},{default:l(()=>[e(t,{field:"corp",label:u.$t("model.label.corp"),rules:[{required:!0,message:u.$t("model.error.corp.required")}]},{default:l(()=>[e(d,{modelValue:r.value.corp,"onUpdate:modelValue":o[0]||(o[0]=p=>r.value.corp=p),placeholder:u.$t("model.placeholder.corp")},{default:l(()=>[e(i,{value:"OpenAI"},{default:l(()=>[n("OpenAI")]),_:1}),e(i,{value:"Midjourney"},{default:l(()=>[n("Midjourney")]),_:1}),e(i,{value:"GLM"},{default:l(()=>[n("\u667A\u8C31AI")]),_:1})]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),e(t,{field:"name",label:u.$t("model.label.name"),rules:[{required:!0,message:u.$t("model.error.name.required")}]},{default:l(()=>[e(k,{modelValue:r.value.name,"onUpdate:modelValue":o[1]||(o[1]=p=>r.value.name=p),placeholder:u.$t("model.placeholder.name")},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),e(t,{field:"model",label:u.$t("model.label.model"),rules:[{required:!0,message:u.$t("model.error.model.required")},{message:u.$t("model.error.model.pattern")}]},{default:l(()=>[e(k,{modelValue:r.value.model,"onUpdate:modelValue":o[2]||(o[2]=p=>r.value.model=p),placeholder:u.$t("model.placeholder.model")},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),e(t,{field:"type",label:u.$t("model.label.type"),rules:[{required:!0,message:u.$t("model.error.type.required")}]},{default:l(()=>[e(d,{modelValue:r.value.type,"onUpdate:modelValue":o[3]||(o[3]=p=>r.value.type=p),placeholder:u.$t("model.placeholder.type")},{default:l(()=>[e(i,{value:"1"},{default:l(()=>[n("\u6587\u751F\u6587")]),_:1}),e(i,{value:"2"},{default:l(()=>[n("\u6587\u751F\u56FE")]),_:1}),e(i,{value:"3"},{default:l(()=>[n("\u56FE\u751F\u6587")]),_:1}),e(i,{value:"4"},{default:l(()=>[n("\u56FE\u751F\u56FE")]),_:1})]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),e(t,{field:"prompt",label:u.$t("model.label.prompt")},{default:l(()=>[e(f,{modelValue:r.value.prompt,"onUpdate:modelValue":o[4]||(o[4]=p=>r.value.prompt=p),placeholder:u.$t("model.placeholder.prompt")},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),e(t,{field:"remark",label:u.$t("model.label.remark")},{default:l(()=>[e(f,{modelValue:r.value.remark,"onUpdate:modelValue":o[5]||(o[5]=p=>r.value.remark=p),placeholder:u.$t("model.placeholder.remark")},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),e(t,null,{default:l(()=>[e(q,{type:"primary",onClick:a},{default:l(()=>[n(v(u.$t("model.button.next")),1)]),_:1})]),_:1})]),_:1},8,["model"])}}});const me=F(se,[["__scopeId","data-v-fc8f0082"]]),pe=C({__name:"advanced",emits:["changeStep"],setup(y,{emit:b}){const{setLoading:_}=I(!0),m=T(),c=h([]);(async()=>{_(!0);try{const{data:d}=await ne();c.value=d.items}catch{}finally{_(!1)}})();const V=h(),a=h({prompt_ratio:1,completion_ratio:1,data_format:"",is_enable_model_agent:!1,model_agents:[],is_public:!0});(async(d={id:m.query.id})=>{_(!0);try{const{data:t}=await N(d);a.value.prompt_ratio=t.prompt_ratio,a.value.completion_ratio=t.completion_ratio,a.value.data_format=String(t.data_format),a.value.is_enable_model_agent=t.is_enable_model_agent,a.value.model_agents=t.model_agents,a.value.is_public=t.is_public}catch{}finally{_(!1)}})();const o=async()=>{var t;await((t=V.value)==null?void 0:t.validate())||b("changeStep","submit",{...a.value})},i=()=>{b("changeStep","backward")};return(d,t)=>{const k=X,f=B,q=Q,w=E,p=W,z=A,K=R,U=S,O=L;return g(),$(O,{ref_key:"formRef",ref:V,model:a.value,class:"form","label-col-props":{span:5},"wrapper-col-props":{span:18}},{default:l(()=>[e(f,{field:"prompt_ratio",label:d.$t("model.label.promptRatio"),rules:[{required:!0,message:d.$t("model.error.promptRatio.required")}]},{default:l(()=>[e(k,{modelValue:a.value.prompt_ratio,"onUpdate:modelValue":t[0]||(t[0]=s=>a.value.prompt_ratio=s),min:.001,placeholder:d.$t("model.placeholder.promptRatio")},null,8,["modelValue","min","placeholder"])]),_:1},8,["label","rules"]),e(f,{field:"completion_ratio",label:d.$t("model.label.completionRatio"),rules:[{required:!0,message:d.$t("model.error.completionRatio.required")}]},{default:l(()=>[e(k,{modelValue:a.value.completion_ratio,"onUpdate:modelValue":t[1]||(t[1]=s=>a.value.completion_ratio=s),min:.001,placeholder:d.$t("model.placeholder.completionRatio")},null,8,["modelValue","min","placeholder"])]),_:1},8,["label","rules"]),e(f,{field:"data_format",label:d.$t("model.label.dataFormat"),rules:[{required:!0,message:d.$t("model.error.dataFormat.required")}]},{default:l(()=>[e(w,{size:"large"},{default:l(()=>[e(q,{modelValue:a.value.data_format,"onUpdate:modelValue":t[2]||(t[2]=s=>a.value.data_format=s),value:"1"},{default:l(()=>[n("\u7EDF\u4E00\u683C\u5F0F")]),_:1},8,["modelValue"]),e(q,{modelValue:a.value.data_format,"onUpdate:modelValue":t[3]||(t[3]=s=>a.value.data_format=s),value:"2"},{default:l(()=>[n("\u5B98\u65B9\u683C\u5F0F")]),_:1},8,["modelValue"])]),_:1})]),_:1},8,["label","rules"]),e(f,{field:"is_public",label:d.$t("model.label.isPublic"),rules:[{required:!0}]},{default:l(()=>[e(p,{modelValue:a.value.is_public,"onUpdate:modelValue":t[4]||(t[4]=s=>a.value.is_public=s)},null,8,["modelValue"])]),_:1},8,["label"]),e(f,{field:"is_enable_model_agent",label:d.$t("model.label.isEnableModelAgent")},{default:l(()=>[e(p,{modelValue:a.value.is_enable_model_agent,"onUpdate:modelValue":t[5]||(t[5]=s=>a.value.is_enable_model_agent=s)},null,8,["modelValue"])]),_:1},8,["label"]),a.value.is_enable_model_agent?(g(),$(f,{key:0,field:"model_agents",label:d.$t("model.label.modelAgents"),rules:[{required:!0,message:d.$t("model.error.modelAgents.required")}]},{default:l(()=>[e(K,{modelValue:a.value.model_agents,"onUpdate:modelValue":t[6]||(t[6]=s=>a.value.model_agents=s),placeholder:d.$t("model.placeholder.modelAgents"),"max-tag-count":3,multiple:"","allow-clear":""},{default:l(()=>[(g(!0),M(P,null,H(c.value,s=>(g(),$(z,{key:s.id,value:s.id,label:s.name},null,8,["value","label"]))),128))]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label","rules"])):D("",!0),e(f,null,{default:l(()=>[e(w,null,{default:l(()=>[e(U,{type:"secondary",onClick:i},{default:l(()=>[n(v(d.$t("model.button.prev")),1)]),_:1}),e(U,{type:"primary",onClick:o},{default:l(()=>[n(v(d.$t("model.button.next")),1)]),_:1})]),_:1})]),_:1})]),_:1},8,["model"])}}});const ie=F(pe,[["__scopeId","data-v-bbcd0827"]]);const _e={},ce={class:"success-wrap"};function fe(y,b){const _=Y,m=S,c=E;return g(),M("div",ce,[e(_,{status:"success",title:y.$t("model.success.title"),subtitle:y.$t("model.success.update.subTitle")},null,8,["title","subtitle"]),e(c,{size:16},{default:l(()=>[e(m,{key:"finish",type:"secondary",onClick:b[0]||(b[0]=r=>y.$router.push({name:"ModelList"}))},{default:l(()=>[n(v(y.$t("model.button.return")),1)]),_:1}),e(m,{key:"again",type:"primary",onClick:b[1]||(b[1]=r=>y.$router.push({name:"ModelDetail",query:{id:`${y.$route.query.id}`}}))},{default:l(()=>[n(v(y.$t("model.button.view")),1)]),_:1})]),_:1})])}const be=F(_e,[["render",fe],["__scopeId","data-v-829f16db"]]),ve={class:"container"},ge={class:"wrapper"},ye={name:"ModelUpdate"},$e=C({...ye,setup(y){const{loading:b,setLoading:_}=I(!1),m=h(1),c=h({}),r=async()=>{_(!0);try{await de(c.value),m.value=3,c.value={}}catch{}finally{_(!1)}},V=(a,u)=>{if(typeof a=="number"){m.value=a;return}if(a==="forward"||a==="submit"){if(c.value={...c.value,...u},a==="submit"){r();return}m.value+=1}else a==="backward"&&(m.value-=1)};return(a,u)=>{const o=j,i=le,d=ae,t=te,k=oe,f=re,q=ue;return g(),M("div",ve,[e(d,{class:"container-breadcrumb"},{default:l(()=>[e(i,null,{default:l(()=>[e(o)]),_:1}),e(i,null,{default:l(()=>[n(v(a.$t("menu.model")),1)]),_:1}),e(i,null,{default:l(()=>[n(v(a.$t("menu.model.update")),1)]),_:1})]),_:1}),e(q,{loading:ee(b),style:{width:"100%"}},{default:l(()=>[e(f,{class:"general-card",bordered:!1},{title:l(()=>[n(v(a.$t("model.title.update")),1)]),default:l(()=>[Z("div",ge,[e(k,{current:m.value,"onUpdate:current":u[0]||(u[0]=w=>m.value=w),style:{width:"660px"},"line-less":"",class:"steps"},{default:l(()=>[e(t,{description:a.$t("model.subTitle.baseInfo")},{default:l(()=>[n(v(a.$t("model.title.baseInfo")),1)]),_:1},8,["description"]),e(t,{description:a.$t("model.subTitle.advanced")},{default:l(()=>[n(v(a.$t("model.title.advanced")),1)]),_:1},8,["description"]),e(t,{description:a.$t("model.subTitle.update.finish")},{default:l(()=>[n(v(a.$t("model.title.update.finish")),1)]),_:1},8,["description"])]),_:1},8,["current"]),(g(),$(x,null,[m.value===1?(g(),$(me,{key:0,onChangeStep:V})):m.value===2?(g(),$(ie,{key:1,onChangeStep:V})):m.value===3?(g(),$(be,{key:2,onChangeStep:V})):D("",!0)],1024))])]),_:1})]),_:1},8,["loading"])])}}});const ze=F($e,[["__scopeId","data-v-b52a59e3"]]);export{ze as default};