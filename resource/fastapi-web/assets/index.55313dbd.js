import{a as Q,B as z,_ as G}from"./index.940e37e7.js";/* empty css                *//* empty css                *//* empty css              *//* empty css               *//* empty css                *//* empty css               *//* empty css               *//* empty css              *//* empty css                *//* empty css               *//* empty css               *//* empty css               */import{d as P,e as m,bS as W,B as _,C as j,aH as e,aG as t,aL as s,aM as c,F as q,bu as K,aD as f,u as r,aE as $,bT as p,aW as O,aK as J,aF as X,aS as Y,b2 as Z,b1 as ee,bP as ae,aT as le,bQ as te,bV as oe,bW as ue,bO as se,aU as re,bi as de,b5 as ie,bI as pe,bR as ne,g as me}from"./arco.a9260898.js";import{u as _e}from"./loading.1f346a94.js";import{f as ce}from"./vue.ad52ddbe.js";import{q as be}from"./common.df364eef.js";import{f as fe}from"./app.a76a1d78.js";import{b as ve}from"./model.7ab43796.js";import"./chart.d103b168.js";const ye={class:"container"},qe={class:"wrapper"},$e={class:"submit-btn"},he={name:"AppCreate"},Ve=P({...he,setup(ke){const{loading:D,setLoading:b}=_e(!1),{proxy:C}=me(),B=ce(),U=Q(),h=m([]);(async()=>{b(!0);try{const{data:l}=await ve();h.value=l.items}catch{}finally{b(!1)}})();const V=m(),a=m({user_id:m(),name:"",models:[],is_limit_quota:!1,quota:m(),quota_expires_at:"",ip_whitelist:"",ip_blacklist:"",is_create_key:!0,remark:""}),F=async()=>{var o;if(!await((o=V.value)==null?void 0:o.validate())){b(!0);try{const{data:n}=await fe(a.value);n.key?(navigator.clipboard.writeText(n.key),O.success("\u65B0\u5EFA\u6210\u529F, \u5BC6\u94A5\u5DF2\u590D\u5236\u5230\u526A\u8D34\u677F")):C.$message.success("\u65B0\u5EFA\u6210\u529F"),B.push({name:"AppList"})}catch{}finally{b(!1)}}},R=l=>{a.value.quota=l*5e5};return(l,o)=>{const n=z,v=J,A=X,k=Y,d=Z,S=ee,T=ae,w=le,i=te,H=oe,I=ue,y=se,g=re,x=de,L=ie,M=pe,N=ne,E=W("permission");return _(),j("div",ye,[e(A,{class:"container-breadcrumb"},{default:t(()=>[e(v,null,{default:t(()=>[e(n)]),_:1}),e(v,null,{default:t(()=>[s(c(l.$t("menu.app")),1)]),_:1}),e(v,null,{default:t(()=>[s(c(l.$t("menu.app.create")),1)]),_:1})]),_:1}),e(N,{loading:r(D),style:{width:"100%"}},{default:t(()=>[e(M,{class:"general-card","body-style":{padding:"0 20px 20px 20px"},bordered:!1},{default:t(()=>[q("div",qe,[e(L,{ref_key:"formRef",ref:V,model:a.value,class:"form","label-col-props":{span:5},"wrapper-col-props":{span:18}},{default:t(()=>[K((_(),f(d,{field:"user_id",label:l.$t("app.label.user_id"),rules:[{required:r(U).role==="admin",message:l.$t("app.error.user_id.required")}]},{default:t(()=>[e(k,{modelValue:a.value.user_id,"onUpdate:modelValue":o[0]||(o[0]=u=>a.value.user_id=u),placeholder:l.$t("app.placeholder.user_id"),min:1},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"])),[[E,["admin"]]]),e(d,{field:"name",label:l.$t("app.label.name"),rules:[{required:!0,message:l.$t("app.error.name.required")},{match:/^.{1,100}$/,message:l.$t("app.error.name.pattern")}]},{default:t(()=>[e(S,{modelValue:a.value.name,"onUpdate:modelValue":o[1]||(o[1]=u=>a.value.name=u),placeholder:l.$t("app.placeholder.name")},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),e(d,{field:"models",label:l.$t("app.label.models")},{default:t(()=>[e(T,{modelValue:a.value.models,"onUpdate:modelValue":o[2]||(o[2]=u=>a.value.models=u),"allow-search":!0,"allow-clear":!0,"tree-checkable":!0,data:h.value,placeholder:l.$t("app.placeholder.models"),"max-tag-count":3,scrollbar:!1,"tree-checked-strategy":"child"},null,8,["modelValue","data","placeholder"])]),_:1},8,["label"]),e(d,{field:"is_limit_quota",label:l.$t("app.label.isLimitQuota")},{default:t(()=>[e(w,{modelValue:a.value.is_limit_quota,"onUpdate:modelValue":o[3]||(o[3]=u=>a.value.is_limit_quota=u)},null,8,["modelValue"])]),_:1},8,["label"]),a.value.is_limit_quota?(_(),f(d,{key:0,field:"quota",label:l.$t("app.label.quota"),rules:[{required:!0,message:l.$t("app.error.quota.required")}]},{default:t(()=>[e(k,{modelValue:a.value.quota,"onUpdate:modelValue":o[4]||(o[4]=u=>a.value.quota=u),placeholder:l.$t("app.placeholder.quota"),precision:0,min:0,max:9999999999999,style:{width:"492px","margin-right":"10px"}},null,8,["modelValue","placeholder"]),q("div",null," $"+c(a.value.quota?r(be)(a.value.quota):"0"),1)]),_:1},8,["label","rules"])):$("",!0),a.value.is_limit_quota?(_(),f(d,{key:1},{default:t(()=>[e(H,{type:"button",onChange:R},{default:t(()=>[e(i,{value:1},{default:t(()=>[s(" $1 ")]),_:1}),e(i,{value:5},{default:t(()=>[s(" $5 ")]),_:1}),e(i,{value:10},{default:t(()=>[s(" $10 ")]),_:1}),e(i,{value:20},{default:t(()=>[s(" $20 ")]),_:1}),e(i,{value:50},{default:t(()=>[s(" $50 ")]),_:1}),e(i,{value:100},{default:t(()=>[s(" $100 ")]),_:1}),e(i,{value:200},{default:t(()=>[s(" $200 ")]),_:1}),e(i,{value:500},{default:t(()=>[s(" $500 ")]),_:1}),e(i,{value:1e3},{default:t(()=>[s(" $1000 ")]),_:1})]),_:1},8,["onChange"])]),_:1})):$("",!0),a.value.is_limit_quota?(_(),f(d,{key:2,field:"quota_expires_at",label:l.$t("app.label.quota_expires_at")},{default:t(()=>[e(I,{modelValue:a.value.quota_expires_at,"onUpdate:modelValue":o[5]||(o[5]=u=>a.value.quota_expires_at=u),placeholder:l.$t("app.placeholder.quota_expires_at"),"time-picker-props":{defaultValue:"23:59:59"},"disabled-date":u=>r(p)(u).isBefore(r(p)()),style:{width:"100%"},"show-time":"",shortcuts:[{label:"1",value:()=>r(p)(a.value.quota_expires_at||new Date().setHours(23,59,59,999)).add(1,"day")},{label:"7",value:()=>r(p)(a.value.quota_expires_at||new Date().setHours(23,59,59,999)).add(7,"day")},{label:"15",value:()=>r(p)(a.value.quota_expires_at||new Date().setHours(23,59,59,999)).add(15,"day")},{label:"30",value:()=>r(p)(a.value.quota_expires_at||new Date().setHours(23,59,59,999)).add(30,"day")},{label:"90",value:()=>r(p)(a.value.quota_expires_at||new Date().setHours(23,59,59,999)).add(90,"day")},{label:"180",value:()=>r(p)(a.value.quota_expires_at||new Date().setHours(23,59,59,999)).add(180,"day")},{label:"365",value:()=>r(p)(a.value.quota_expires_at||new Date().setHours(23,59,59,999)).add(365,"day")}]},null,8,["modelValue","placeholder","disabled-date","shortcuts"])]),_:1},8,["label"])):$("",!0),e(d,{field:"ip_whitelist",label:l.$t("app.label.ip_whitelist")},{default:t(()=>[e(y,{modelValue:a.value.ip_whitelist,"onUpdate:modelValue":o[6]||(o[6]=u=>a.value.ip_whitelist=u),placeholder:l.$t("app.placeholder.ip_whitelist"),"auto-size":{minRows:5,maxRows:5}},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),e(d,{field:"ip_blacklist",label:l.$t("app.label.ip_blacklist")},{default:t(()=>[e(y,{modelValue:a.value.ip_blacklist,"onUpdate:modelValue":o[7]||(o[7]=u=>a.value.ip_blacklist=u),placeholder:l.$t("app.placeholder.ip_blacklist"),"auto-size":{minRows:5,maxRows:5}},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),e(d,{field:"is_create_key",label:l.$t("app.label.is_create_key")},{default:t(()=>[e(w,{modelValue:a.value.is_create_key,"onUpdate:modelValue":o[8]||(o[8]=u=>a.value.is_create_key=u)},null,8,["modelValue"])]),_:1},8,["label"]),e(d,{field:"remark",label:l.$t("app.label.remark"),rules:[{required:!1}]},{default:t(()=>[e(y,{modelValue:a.value.remark,"onUpdate:modelValue":o[9]||(o[9]=u=>a.value.remark=u),placeholder:l.$t("app.placeholder.remark")},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),e(x,null,{default:t(()=>[q("div",$e,[e(g,{type:"secondary",onClick:o[10]||(o[10]=u=>l.$router.push({name:"AppList"}))},{default:t(()=>[s(c(l.$t("button.cancel")),1)]),_:1}),e(g,{type:"primary",onClick:F},{default:t(()=>[s(c(l.$t("button.submit")),1)]),_:1})])]),_:1})]),_:1},8,["model"])])]),_:1})]),_:1},8,["loading"])])}}});const Pe=G(Ve,[["__scopeId","data-v-b3eee6ad"]]);export{Pe as default};
