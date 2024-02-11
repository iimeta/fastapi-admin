import{d as k,r as y,o as c,f as $,w as a,g as e,j as i,k as d,I as z,F as B,b1 as R,B as C,m as N,_ as w,q as I,s as D,ar as P,y as U,aP as Q,aQ as j,a$ as E,b0 as J,a4 as A,aJ as K,i as O,ay as G,h as H,aL as W,aM as X,aN as Y,b6 as Z,b7 as x,b3 as ee,b8 as ae}from"./index.9d2a6f93.js";import{u as T}from"./loading.2f897f78.js";/* empty css                */import{c as te}from"./app.4b161933.js";/* empty css               *//* empty css               *//* empty css                *//* empty css               *//* empty css              *//* empty css              *//* empty css               */import{q as le}from"./model.d6c87741.js";/* empty css               */const oe=k({__name:"base-info",emits:["changeStep"],setup(V,{emit:b}){const _=y(),l=y({name:"",remark:""}),m=async()=>{var t;await((t=_.value)==null?void 0:t.validate())||b("changeStep","forward",{...l.value})};return(p,t)=>{const o=z,u=B,s=R,n=C,h=N;return c(),$(h,{ref_key:"formRef",ref:_,model:l.value,class:"form","label-col-props":{span:6},"wrapper-col-props":{span:18}},{default:a(()=>[e(u,{field:"name",label:p.$t("app.label.name"),rules:[{required:!0,message:p.$t("app.error.name.required")},{match:/^.{1,20}$/,message:p.$t("app.error.name.pattern")}]},{default:a(()=>[e(o,{modelValue:l.value.name,"onUpdate:modelValue":t[0]||(t[0]=f=>l.value.name=f),placeholder:p.$t("app.placeholder.name"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),e(u,{field:"remark",label:p.$t("app.label.remark"),rules:[{required:!1}]},{default:a(()=>[e(s,{modelValue:l.value.remark,"onUpdate:modelValue":t[1]||(t[1]=f=>l.value.remark=f),placeholder:p.$t("app.placeholder.remark")},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),e(u,null,{default:a(()=>[e(n,{type:"primary",onClick:m},{default:a(()=>[i(d(p.$t("app.button.next")),1)]),_:1})]),_:1})]),_:1},8,["model"])}}});const se=w(oe,[["__scopeId","data-v-e67b024a"]]),ne=k({__name:"advanced",emits:["changeStep"],setup(V,{emit:b}){const{setLoading:_}=T(!0),l=y([]);(async()=>{_(!0);try{const{data:s}=await le();l.value=s.items}catch{}finally{_(!1)}})();const p=y(),t=y({models:[],is_limit_quota:!1,quota:y(),ip_whitelist:"",ip_blacklist:""}),o=async()=>{var n;await((n=p.value)==null?void 0:n.validate())||b("changeStep","submit",{...t.value})},u=()=>{b("changeStep","backward")};return(s,n)=>{const h=Q,f=j,v=B,q=E,S=J,g=R,L=C,F=A,M=N;return c(),$(M,{ref_key:"formRef",ref:p,model:t.value,class:"form","label-col-props":{span:6},"wrapper-col-props":{span:18}},{default:a(()=>[e(v,{field:"models",label:s.$t("app.label.models"),rules:[{required:!1}]},{default:a(()=>[e(f,{modelValue:t.value.models,"onUpdate:modelValue":n[0]||(n[0]=r=>t.value.models=r),placeholder:s.$t("app.placeholder.models"),"max-tag-count":3,multiple:"","allow-clear":""},{default:a(()=>[(c(!0),I(D,null,P(l.value,r=>(c(),$(h,{key:r.id,value:r.id,label:r.name},null,8,["value","label"]))),128))]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label"]),e(v,{field:"is_limit_quota",label:s.$t("app.label.isLimitQuota")},{default:a(()=>[e(q,{modelValue:t.value.is_limit_quota,"onUpdate:modelValue":n[1]||(n[1]=r=>t.value.is_limit_quota=r)},null,8,["modelValue"])]),_:1},8,["label"]),t.value.is_limit_quota?(c(),$(v,{key:0,field:"quota",label:s.$t("app.label.quota"),rules:[{required:!0,message:s.$t("app.error.quota.required")}]},{default:a(()=>[e(S,{modelValue:t.value.quota,"onUpdate:modelValue":n[2]||(n[2]=r=>t.value.quota=r),placeholder:s.$t("app.placeholder.quota"),min:1},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"])):U("",!0),e(v,{field:"ip_whitelist",label:s.$t("app.label.ip_whitelist"),rules:[{required:!1}]},{default:a(()=>[e(g,{modelValue:t.value.ip_whitelist,"onUpdate:modelValue":n[3]||(n[3]=r=>t.value.ip_whitelist=r),placeholder:s.$t("app.placeholder.ip_whitelist"),"auto-size":{minRows:5,maxRows:10}},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),e(v,{field:"ip_blacklist",label:s.$t("app.label.ip_blacklist"),rules:[{required:!1}]},{default:a(()=>[e(g,{modelValue:t.value.ip_blacklist,"onUpdate:modelValue":n[4]||(n[4]=r=>t.value.ip_blacklist=r),placeholder:s.$t("app.placeholder.ip_blacklist"),"auto-size":{minRows:5,maxRows:10}},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),e(v,null,{default:a(()=>[e(F,null,{default:a(()=>[e(L,{type:"secondary",onClick:u},{default:a(()=>[i(d(s.$t("app.button.prev")),1)]),_:1}),e(L,{type:"primary",onClick:o},{default:a(()=>[i(d(s.$t("app.button.next")),1)]),_:1})]),_:1})]),_:1})]),_:1},8,["model"])}}});const pe=w(ne,[["__scopeId","data-v-e1e29c1c"]]),re={class:"success-wrap"},ue=k({__name:"success",emits:["changeStep"],setup(V,{emit:b}){const _=()=>{b("changeStep",1)};return(l,m)=>{const p=K,t=C,o=A;return c(),I("div",re,[e(p,{status:"success",title:l.$t("app.success.title"),subtitle:l.$t("app.success.create.subTitle")},null,8,["title","subtitle"]),e(o,{size:16},{default:a(()=>[e(t,{key:"finish",type:"secondary",onClick:m[0]||(m[0]=u=>l.$router.push({name:"AppList"}))},{default:a(()=>[i(d(l.$t("app.button.finish")),1)]),_:1}),e(t,{key:"again",type:"primary",onClick:_},{default:a(()=>[i(d(l.$t("app.button.again")),1)]),_:1})]),_:1})])}}});const ie=w(ue,[["__scopeId","data-v-9ee08270"]]),de={class:"container"},ce={class:"wrapper"},_e={name:"AppCreate"},me=k({..._e,setup(V){const{loading:b,setLoading:_}=T(!1),l=y(1),m=y({}),p=async()=>{_(!0);try{await te(m.value),l.value=3,m.value={}}catch{}finally{_(!1)}},t=(o,u)=>{if(typeof o=="number"){l.value=o;return}if(o==="forward"||o==="submit"){if(m.value={...m.value,...u},o==="submit"){p();return}l.value+=1}else o==="backward"&&(l.value-=1)};return(o,u)=>{const s=W,n=X,h=Y,f=Z,v=x,q=ee,S=ae;return c(),I("div",de,[e(h,{class:"container-breadcrumb"},{default:a(()=>[e(n,null,{default:a(()=>[e(s)]),_:1}),e(n,null,{default:a(()=>[i(d(o.$t("menu.app")),1)]),_:1}),e(n,null,{default:a(()=>[i(d(o.$t("menu.app.create")),1)]),_:1})]),_:1}),e(S,{loading:H(b),style:{width:"100%"}},{default:a(()=>[e(q,{class:"general-card",bordered:!1},{title:a(()=>[i(d(o.$t("app.title.create")),1)]),default:a(()=>[O("div",ce,[e(v,{current:l.value,"onUpdate:current":u[0]||(u[0]=g=>l.value=g),style:{width:"580px"},"line-less":"",class:"steps"},{default:a(()=>[e(f,{description:o.$t("app.subTitle.baseInfo")},{default:a(()=>[i(d(o.$t("app.title.baseInfo")),1)]),_:1},8,["description"]),e(f,{description:o.$t("app.subTitle.advanced")},{default:a(()=>[i(d(o.$t("app.title.advanced")),1)]),_:1},8,["description"]),e(f,{description:o.$t("app.subTitle.create.finish")},{default:a(()=>[i(d(o.$t("app.title.create.finish")),1)]),_:1},8,["description"])]),_:1},8,["current"]),(c(),$(G,null,[l.value===1?(c(),$(se,{key:0,onChangeStep:t})):l.value===2?(c(),$(pe,{key:1,onChangeStep:t})):l.value===3?(c(),$(ie,{key:2,onChangeStep:t})):U("",!0)],1024))])]),_:1})]),_:1},8,["loading"])])}}});const Ie=w(me,[["__scopeId","data-v-3cb4c1fb"]]);export{Ie as default};
