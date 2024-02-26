import{_ as h,F as D}from"./index.82aeb519.js";import{u as U}from"./loading.b0aa7954.js";/* empty css                */import{d as F,e as $,B as _,aD as g,aG as l,aH as e,aL as s,aM as i,bC as A,bD as R,b3 as B,b2 as O,bL as K,aV as S,b6 as M,C as w,aJ as z,aI as P,aE as L,aT as j,b$ as G,bj as E,aU as H,bA as J,F as Q,bt as W,u as X,aK as Y,aF as Z,bM as x,bN as ee,bK as le,bO as ae}from"./arco.a11b8b88.js";import{c as oe}from"./model.ede023ed.js";/* empty css               *//* empty css               *//* empty css                *//* empty css              *//* empty css              *//* empty css               *//* empty css               *//* empty css                */import{e as te}from"./agent.02d2a068.js";/* empty css               */import"./chart.87d6227d.js";import"./vue.e9a5701d.js";const re=F({__name:"base-info",emits:["changeStep"],setup(C,{emit:v}){const c=$(),o=$({corp:"",name:"",model:"",type:"1",remark:""}),f=async()=>{var a;await((a=c.value)==null?void 0:a.validate())||v("changeStep","forward",{...o.value})};return(u,a)=>{const r=A,b=R,t=B,n=O,y=K,p=S,V=M;return _(),g(V,{ref_key:"formRef",ref:c,model:o.value,class:"form","label-col-props":{span:5},"wrapper-col-props":{span:18}},{default:l(()=>[e(t,{field:"corp",label:u.$t("model.label.corp"),rules:[{required:!0,message:u.$t("model.error.corp.required")}]},{default:l(()=>[e(b,{modelValue:o.value.corp,"onUpdate:modelValue":a[0]||(a[0]=m=>o.value.corp=m),placeholder:u.$t("model.placeholder.corp")},{default:l(()=>[e(r,{value:"OpenAI"},{default:l(()=>[s("OpenAI")]),_:1})]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),e(t,{field:"name",label:u.$t("model.label.name"),rules:[{required:!0,message:u.$t("model.error.name.required")}]},{default:l(()=>[e(n,{modelValue:o.value.name,"onUpdate:modelValue":a[1]||(a[1]=m=>o.value.name=m),placeholder:u.$t("model.placeholder.name")},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),e(t,{field:"model",label:u.$t("model.label.model"),rules:[{required:!0,message:u.$t("model.error.model.required")}]},{default:l(()=>[e(n,{modelValue:o.value.model,"onUpdate:modelValue":a[2]||(a[2]=m=>o.value.model=m),placeholder:u.$t("model.placeholder.model")},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),e(t,{field:"type",label:u.$t("model.label.type"),rules:[{required:!0,message:u.$t("model.error.type.required")}]},{default:l(()=>[e(b,{modelValue:o.value.type,"onUpdate:modelValue":a[3]||(a[3]=m=>o.value.type=m),placeholder:u.$t("model.placeholder.type")},{default:l(()=>[e(r,{value:"1"},{default:l(()=>[s("\u6587\u751F\u6587")]),_:1}),e(r,{value:"2"},{default:l(()=>[s("\u6587\u751F\u56FE")]),_:1}),e(r,{value:"3"},{default:l(()=>[s("\u56FE\u751F\u6587")]),_:1}),e(r,{value:"4"},{default:l(()=>[s("\u56FE\u751F\u56FE")]),_:1})]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),e(t,{field:"remark",label:u.$t("model.label.remark"),rules:[{required:!1}]},{default:l(()=>[e(y,{modelValue:o.value.remark,"onUpdate:modelValue":a[4]||(a[4]=m=>o.value.remark=m),placeholder:u.$t("model.placeholder.remark")},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),e(t,null,{default:l(()=>[e(p,{type:"primary",onClick:f},{default:l(()=>[s(i(u.$t("model.button.next")),1)]),_:1})]),_:1})]),_:1},8,["model"])}}});const ne=h(re,[["__scopeId","data-v-0ae27953"]]),ue=F({__name:"advanced",emits:["changeStep"],setup(C,{emit:v}){const{setLoading:c}=U(!0),o=$([]);(async()=>{c(!0);try{const{data:t}=await te();o.value=t.items}catch{}finally{c(!1)}})();const u=$(),a=$({prompt_ratio:1,completion_ratio:1,data_format:"1",is_enable_model_agent:!1,model_agents:[],is_public:!0}),r=async()=>{var n;await((n=u.value)==null?void 0:n.validate())||v("changeStep","submit",{...a.value})},b=()=>{v("changeStep","backward")};return(t,n)=>{const y=j,p=B,V=G,m=E,k=H,q=A,N=R,I=S,T=M;return _(),g(T,{ref_key:"formRef",ref:u,model:a.value,class:"form","label-col-props":{span:5},"wrapper-col-props":{span:18}},{default:l(()=>[e(p,{field:"prompt_ratio",label:t.$t("model.label.promptRatio"),rules:[{required:!0,message:t.$t("model.error.promptRatio.required")}]},{default:l(()=>[e(y,{modelValue:a.value.prompt_ratio,"onUpdate:modelValue":n[0]||(n[0]=d=>a.value.prompt_ratio=d),min:1,placeholder:t.$t("model.placeholder.promptRatio")},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),e(p,{field:"completion_ratio",label:t.$t("model.label.completionRatio"),rules:[{required:!0,message:t.$t("model.error.completionRatio.required")}]},{default:l(()=>[e(y,{modelValue:a.value.completion_ratio,"onUpdate:modelValue":n[1]||(n[1]=d=>a.value.completion_ratio=d),min:1,placeholder:t.$t("model.placeholder.completionRatio")},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),e(p,{field:"data_format",label:t.$t("model.label.dataFormat"),rules:[{required:!0,message:t.$t("model.error.dataFormat.required")}]},{default:l(()=>[e(m,{size:"large"},{default:l(()=>[e(V,{modelValue:a.value.data_format,"onUpdate:modelValue":n[2]||(n[2]=d=>a.value.data_format=d),value:"1","default-checked":!0},{default:l(()=>[s("\u7EDF\u4E00\u683C\u5F0F")]),_:1},8,["modelValue"]),e(V,{modelValue:a.value.data_format,"onUpdate:modelValue":n[3]||(n[3]=d=>a.value.data_format=d),value:"2"},{default:l(()=>[s("\u5B98\u65B9\u683C\u5F0F")]),_:1},8,["modelValue"])]),_:1})]),_:1},8,["label","rules"]),e(p,{field:"is_public",label:t.$t("model.label.isPublic"),rules:[{required:!0}]},{default:l(()=>[e(k,{modelValue:a.value.is_public,"onUpdate:modelValue":n[4]||(n[4]=d=>a.value.is_public=d)},null,8,["modelValue"])]),_:1},8,["label"]),e(p,{field:"is_enable_model_agent",label:t.$t("model.label.isEnableModelAgent")},{default:l(()=>[e(k,{modelValue:a.value.is_enable_model_agent,"onUpdate:modelValue":n[5]||(n[5]=d=>a.value.is_enable_model_agent=d)},null,8,["modelValue"])]),_:1},8,["label"]),a.value.is_enable_model_agent?(_(),g(p,{key:0,field:"model_agents",label:t.$t("model.label.modelAgents"),rules:[{required:!0,message:t.$t("model.error.modelAgents.required")}]},{default:l(()=>[e(N,{modelValue:a.value.model_agents,"onUpdate:modelValue":n[6]||(n[6]=d=>a.value.model_agents=d),placeholder:t.$t("model.placeholder.modelAgents"),"max-tag-count":3,multiple:"","allow-clear":""},{default:l(()=>[(_(!0),w(z,null,P(o.value,d=>(_(),g(q,{key:d.id,value:d.id,label:d.name},null,8,["value","label"]))),128))]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label","rules"])):L("",!0),e(p,null,{default:l(()=>[e(m,null,{default:l(()=>[e(I,{type:"secondary",onClick:b},{default:l(()=>[s(i(t.$t("model.button.prev")),1)]),_:1}),e(I,{type:"primary",onClick:r},{default:l(()=>[s(i(t.$t("model.button.next")),1)]),_:1})]),_:1})]),_:1})]),_:1},8,["model"])}}});const se=h(ue,[["__scopeId","data-v-8b3c354e"]]),de={class:"success-wrap"},me=F({__name:"success",emits:["changeStep"],setup(C,{emit:v}){const c=()=>{v("changeStep",1)};return(o,f)=>{const u=J,a=S,r=E;return _(),w("div",de,[e(u,{status:"success",title:o.$t("model.success.title"),subtitle:o.$t("model.success.create.subTitle")},null,8,["title","subtitle"]),e(r,{size:16},{default:l(()=>[e(a,{key:"finish",type:"secondary",onClick:f[0]||(f[0]=b=>o.$router.push({name:"ModelList"}))},{default:l(()=>[s(i(o.$t("model.button.finish")),1)]),_:1}),e(a,{key:"again",type:"primary",onClick:c},{default:l(()=>[s(i(o.$t("model.button.again")),1)]),_:1})]),_:1})])}}});const pe=h(me,[["__scopeId","data-v-77484237"]]),ie={class:"container"},_e={class:"wrapper"},ce={name:"ModelCreate"},fe=F({...ce,setup(C){const{loading:v,setLoading:c}=U(!1),o=$(1),f=$({}),u=async()=>{c(!0);try{await oe(f.value),o.value=3,f.value={}}catch{}finally{c(!1)}},a=(r,b)=>{if(typeof r=="number"){o.value=r;return}if(r==="forward"||r==="submit"){if(f.value={...f.value,...b},r==="submit"){u();return}o.value+=1}else r==="backward"&&(o.value-=1)};return(r,b)=>{const t=D,n=Y,y=Z,p=x,V=ee,m=le,k=ae;return _(),w("div",ie,[e(y,{class:"container-breadcrumb"},{default:l(()=>[e(n,null,{default:l(()=>[e(t)]),_:1}),e(n,null,{default:l(()=>[s(i(r.$t("menu.model")),1)]),_:1}),e(n,null,{default:l(()=>[s(i(r.$t("menu.model.create")),1)]),_:1})]),_:1}),e(k,{loading:X(v),style:{width:"100%"}},{default:l(()=>[e(m,{class:"general-card",bordered:!1},{title:l(()=>[s(i(r.$t("model.title.create")),1)]),default:l(()=>[Q("div",_e,[e(V,{current:o.value,"onUpdate:current":b[0]||(b[0]=q=>o.value=q),style:{width:"660px"},"line-less":"",class:"steps"},{default:l(()=>[e(p,{description:r.$t("model.subTitle.baseInfo")},{default:l(()=>[s(i(r.$t("model.title.baseInfo")),1)]),_:1},8,["description"]),e(p,{description:r.$t("model.subTitle.advanced")},{default:l(()=>[s(i(r.$t("model.title.advanced")),1)]),_:1},8,["description"]),e(p,{description:r.$t("model.subTitle.create.finish")},{default:l(()=>[s(i(r.$t("model.title.create.finish")),1)]),_:1},8,["description"])]),_:1},8,["current"]),(_(),g(W,null,[o.value===1?(_(),g(ne,{key:0,onChangeStep:a})):o.value===2?(_(),g(se,{key:1,onChangeStep:a})):o.value===3?(_(),g(pe,{key:2,onChangeStep:a})):L("",!0)],1024))])]),_:1})]),_:1},8,["loading"])])}}});const Be=h(fe,[["__scopeId","data-v-ea8fc182"]]);export{Be as default};