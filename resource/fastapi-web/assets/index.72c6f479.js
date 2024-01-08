import{_ as C,F as A}from"./index.3029f51b.js";import{u as I}from"./loading.fc5bfc3b.js";/* empty css               *//* empty css               */import{d as q,e as $,B as f,aD as k,aG as a,aH as e,aL as p,aM as _,bN as L,bO as B,b4 as U,b3 as O,bU as E,aW as K,b7 as F,C as D,aJ as H,aI as W,bn as N,bt as z,F as G,bH as J,aE as P,u as X,aK as j,aF as Q,bW as Y,bX as Z,bV as x,bs as ee}from"./arco.fd20202f.js";import{b as R,c as ae}from"./key.3d211d3d.js";/* empty css              *//* empty css               *//* empty css               *//* empty css              */import{h as T}from"./vue.70a4bb93.js";import{q as te}from"./model.943397e8.js";/* empty css               */import"./chart.57980958.js";const oe=q({__name:"base-info",emits:["changeStep"],setup(m,{emit:i}){const{setLoading:c}=I(!1),r=T(),d=$(),s=$({id:"",corp:"",key:"",remark:""});(async(l={id:r.query.id})=>{c(!0);try{const{data:o}=await R(l);s.value.id=o.id,s.value.corp=o.corp,s.value.key=o.key,s.value.remark=o.remark}catch{}finally{c(!1)}})();const t=async()=>{var o;await((o=d.value)==null?void 0:o.validate())||i("changeStep","forward",{...s.value})};return(l,o)=>{const b=L,u=B,n=U,w=O,S=E,g=K,h=F;return f(),k(h,{ref_key:"formRef",ref:d,model:s.value,class:"form","label-col-props":{span:6},"wrapper-col-props":{span:18}},{default:a(()=>[e(n,{field:"corp",label:l.$t("key.label.corp"),rules:[{required:!0,message:l.$t("key.error.corp.required")}]},{default:a(()=>[e(u,{modelValue:s.value.corp,"onUpdate:modelValue":o[0]||(o[0]=v=>s.value.corp=v),placeholder:l.$t("key.placeholder.corp")},{default:a(()=>[e(b,{value:"OpenAI"},{default:a(()=>[p("OpenAI")]),_:1})]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),e(n,{field:"key",label:l.$t("key.label.key"),rules:[{required:!0,message:l.$t("key.error.key.required")}]},{default:a(()=>[e(w,{modelValue:s.value.key,"onUpdate:modelValue":o[1]||(o[1]=v=>s.value.key=v),placeholder:l.$t("key.placeholder.update.key")},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),e(n,{field:"remark",label:l.$t("key.label.remark"),rules:[{required:!1}]},{default:a(()=>[e(S,{modelValue:s.value.remark,"onUpdate:modelValue":o[2]||(o[2]=v=>s.value.remark=v),placeholder:l.$t("key.placeholder.remark")},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),e(n,null,{default:a(()=>[e(g,{type:"primary",onClick:t},{default:a(()=>[p(_(l.$t("key.button.next")),1)]),_:1})]),_:1})]),_:1},8,["model"])}}});const le=C(oe,[["__scopeId","data-v-38b7c44b"]]),se=q({__name:"advanced",emits:["changeStep"],setup(m,{emit:i}){const{setLoading:c}=I(!0),r=T(),d=$([]);(async()=>{c(!0);try{const{data:u}=await te();d.value=u.items}catch{}finally{c(!1)}})();const y=$(),t=$({models:[]});(async(u={id:r.query.id})=>{c(!0);try{const{data:n}=await R(u);t.value.models=n.models}catch{}finally{c(!1)}})();const o=async()=>{var n;await((n=y.value)==null?void 0:n.validate())||i("changeStep","submit",{...t.value})},b=()=>{i("changeStep","backward")};return(u,n)=>{const w=L,S=B,g=U,h=K,v=N,M=F;return f(),k(M,{ref_key:"formRef",ref:y,model:t.value,class:"form","label-col-props":{span:6},"wrapper-col-props":{span:18}},{default:a(()=>[e(g,{field:"models",label:u.$t("key.label.models"),rules:[{required:!1}]},{default:a(()=>[e(S,{modelValue:t.value.models,"onUpdate:modelValue":n[0]||(n[0]=V=>t.value.models=V),placeholder:u.$t("key.placeholder.models"),"max-tag-count":3,multiple:"","allow-clear":""},{default:a(()=>[(f(!0),D(H,null,W(d.value,V=>(f(),k(w,{key:V.id,value:V.id,label:V.name},null,8,["value","label"]))),128))]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label"]),e(g,null,{default:a(()=>[e(v,null,{default:a(()=>[e(h,{type:"secondary",onClick:b},{default:a(()=>[p(_(u.$t("key.button.prev")),1)]),_:1}),e(h,{type:"primary",onClick:o},{default:a(()=>[p(_(u.$t("key.button.next")),1)]),_:1})]),_:1})]),_:1})]),_:1},8,["model"])}}});const ne=C(se,[["__scopeId","data-v-2654ac66"]]);const re={},ue={class:"success-wrap"};function ce(m,i){const c=z,r=K,d=N;return f(),D("div",ue,[e(c,{status:"success",title:m.$t("key.success.title"),subtitle:m.$t("key.success.update.subTitle")},null,8,["title","subtitle"]),e(d,{size:16},{default:a(()=>[e(r,{key:"finish",type:"secondary",onClick:i[0]||(i[0]=s=>m.$router.push({name:"KeyModelList"}))},{default:a(()=>[p(_(m.$t("key.button.return")),1)]),_:1}),e(r,{key:"again",type:"primary",onClick:i[1]||(i[1]=s=>m.$router.push({name:"KeyDetail",query:{id:`${m.$route.query.id}`}}))},{default:a(()=>[p(_(m.$t("key.button.view")),1)]),_:1})]),_:1})])}const de=C(re,[["render",ce],["__scopeId","data-v-016ee704"]]),pe={class:"container"},ie={class:"wrapper"},_e={name:"KeyUpdate"},me=q({..._e,setup(m){const{loading:i,setLoading:c}=I(!1),r=$(1),d=$({}),s=async()=>{c(!0);try{await ae(d.value),r.value=3,d.value={}}catch{}finally{c(!1)}},y=(t,l)=>{if(typeof t=="number"){r.value=t;return}if(t==="forward"||t==="submit"){if(d.value={...d.value,...l},t==="submit"){s();return}r.value+=1}else t==="backward"&&(r.value-=1)};return(t,l)=>{const o=A,b=j,u=Q,n=Y,w=Z,S=x,g=ee;return f(),D("div",pe,[e(u,{class:"container-breadcrumb"},{default:a(()=>[e(b,null,{default:a(()=>[e(o)]),_:1}),e(b,null,{default:a(()=>[p(_(t.$t("menu.key")),1)]),_:1}),e(b,null,{default:a(()=>[p(_(t.$t("menu.key.update")),1)]),_:1})]),_:1}),e(g,{loading:X(i),style:{width:"100%"}},{default:a(()=>[e(S,{class:"general-card"},{title:a(()=>[p(_(t.$t("key.title.update")),1)]),default:a(()=>[G("div",ie,[e(w,{current:r.value,"onUpdate:current":l[0]||(l[0]=h=>r.value=h),style:{width:"580px"},"line-less":"",class:"steps"},{default:a(()=>[e(n,{description:t.$t("key.subTitle.baseInfo")},{default:a(()=>[p(_(t.$t("key.title.baseInfo")),1)]),_:1},8,["description"]),e(n,{description:t.$t("key.subTitle.advanced")},{default:a(()=>[p(_(t.$t("key.title.advanced")),1)]),_:1},8,["description"]),e(n,{description:t.$t("key.subTitle.update.finish")},{default:a(()=>[p(_(t.$t("key.title.update.finish")),1)]),_:1},8,["description"])]),_:1},8,["current"]),(f(),k(J,null,[r.value===1?(f(),k(le,{key:0,onChangeStep:y})):r.value===2?(f(),k(ne,{key:1,onChangeStep:y})):r.value===3?(f(),k(de,{key:2,onChangeStep:y})):P("",!0)],1024))])]),_:1})]),_:1},8,["loading"])])}}});const Ke=C(me,[["__scopeId","data-v-13ec56e8"]]);export{Ke as default};
