import{_ as S,w as z}from"./index.05df1f52.js";/* empty css               *//* empty css                */import{d as C,e as v,B as b,aD as k,aG as a,aH as e,aL as c,aM as _,b1 as K,b2 as A,aS as E,bK as B,aU as M,b5 as L,C as I,aJ as J,aI as G,bB as H,bC as O,bi as D,bz as P,F as j,bs as Q,aE as W,u as X,aK as Y,aF as Z,bL as x,bM as ee,bJ as ae,bN as te}from"./arco.aed15247.js";import{u as U}from"./loading.b5911e1d.js";import{c as N,d as le}from"./agent.4ef33138.js";/* empty css               *//* empty css               *//* empty css                */import{h as R}from"./vue.0cc5b64a.js";/* empty css               *//* empty css              *//* empty css              *//* empty css               */import{q as oe}from"./model.89eea4c7.js";/* empty css               */import"./chart.9aa6eafa.js";import"./base.87fcf6e2.js";const ne=C({__name:"base-info",emits:["changeStep"],setup(f,{emit:i}){const{setLoading:m}=U(!1),d=R(),p=v(),o=v({id:"",name:"",base_url:"",path:"",weight:v(),remark:"",status:1});(async(n={id:d.query.id})=>{m(!0);try{const{data:l}=await N(n);o.value.id=l.id,o.value.name=l.name,o.value.base_url=l.base_url,o.value.path=l.path,o.value.weight=l.weight,o.value.remark=l.remark,o.value.status=l.status}catch{}finally{m(!1)}})();const t=async()=>{var l;await((l=p.value)==null?void 0:l.validate())||i("changeStep","forward",{...o.value})};return(n,l)=>{const g=K,s=A,r=E,w=B,V=M,h=L;return b(),k(h,{ref_key:"formRef",ref:p,model:o.value,class:"form","label-col-props":{span:4},"wrapper-col-props":{span:18}},{default:a(()=>[e(s,{field:"name",label:n.$t("model.agent.label.name"),rules:[{required:!0,message:n.$t("model.agent.error.name.required")},{match:/^.{1,100}$/,message:n.$t("model.agent.error.name.pattern")}]},{default:a(()=>[e(g,{modelValue:o.value.name,"onUpdate:modelValue":l[0]||(l[0]=u=>o.value.name=u),placeholder:n.$t("model.agent.placeholder.name")},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),e(s,{field:"base_url",label:n.$t("model.agent.label.baseUrl"),rules:[{required:!0,message:n.$t("model.agent.error.baseUrl.required")}]},{default:a(()=>[e(g,{modelValue:o.value.base_url,"onUpdate:modelValue":l[1]||(l[1]=u=>o.value.base_url=u),placeholder:n.$t("model.agent.placeholder.baseUrl")},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),e(s,{field:"path",label:n.$t("model.agent.label.path")},{default:a(()=>[e(g,{modelValue:o.value.path,"onUpdate:modelValue":l[2]||(l[2]=u=>o.value.path=u),placeholder:n.$t("model.agent.placeholder.path")},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),e(s,{field:"weight",label:n.$t("model.agent.label.weight")},{default:a(()=>[e(r,{modelValue:o.value.weight,"onUpdate:modelValue":l[3]||(l[3]=u=>o.value.weight=u),precision:0,min:0,max:99999,placeholder:n.$t("model.agent.placeholder.weight")},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),e(s,{field:"remark",label:n.$t("model.agent.label.remark")},{default:a(()=>[e(w,{modelValue:o.value.remark,"onUpdate:modelValue":l[4]||(l[4]=u=>o.value.remark=u),placeholder:n.$t("model.agent.placeholder.remark")},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),e(s,null,{default:a(()=>[e(V,{type:"primary",onClick:t},{default:a(()=>[c(_(n.$t("model.agent.button.next")),1)]),_:1})]),_:1})]),_:1},8,["model"])}}});const se=S(ne,[["__scopeId","data-v-9ce6e830"]]),re=C({__name:"advanced",emits:["changeStep"],setup(f,{emit:i}){const{setLoading:m}=U(!0),d=R(),p=v([]);(async()=>{m(!0);try{const{data:s}=await oe();p.value=s.items}catch{}finally{m(!1)}})();const y=v(),t=v({models:[],key:""});(async(s={id:d.query.id})=>{m(!0);try{const{data:r}=await N(s);t.value.models=r.models,t.value.key=r.key}catch{}finally{m(!1)}})();const l=async()=>{var r;await((r=y.value)==null?void 0:r.validate())||i("changeStep","submit",{...t.value})},g=()=>{i("changeStep","backward")};return(s,r)=>{const w=H,V=O,h=A,u=B,q=M,F=D,T=L;return b(),k(T,{ref_key:"formRef",ref:y,model:t.value,class:"form","label-col-props":{span:5},"wrapper-col-props":{span:18}},{default:a(()=>[e(h,{field:"models",label:s.$t("model.agent.label.models"),rules:[{required:!1}]},{default:a(()=>[e(V,{modelValue:t.value.models,"onUpdate:modelValue":r[0]||(r[0]=$=>t.value.models=$),placeholder:s.$t("model.agent.placeholder.models"),"max-tag-count":3,multiple:"","allow-clear":""},{default:a(()=>[(b(!0),I(J,null,G(p.value,$=>(b(),k(w,{key:$.id,value:$.id,label:$.name},null,8,["value","label"]))),128))]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label"]),e(h,{field:"key",label:s.$t("model.agent.label.key")},{default:a(()=>[e(u,{modelValue:t.value.key,"onUpdate:modelValue":r[1]||(r[1]=$=>t.value.key=$),placeholder:s.$t("model.agent.placeholder.key"),"auto-size":{minRows:5,maxRows:10}},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),e(h,null,{default:a(()=>[e(F,null,{default:a(()=>[e(q,{type:"secondary",onClick:g},{default:a(()=>[c(_(s.$t("model.button.prev")),1)]),_:1}),e(q,{type:"primary",onClick:l},{default:a(()=>[c(_(s.$t("model.button.next")),1)]),_:1})]),_:1})]),_:1})]),_:1},8,["model"])}}});const de=S(re,[["__scopeId","data-v-923731d5"]]);const ue={},me={class:"success-wrap"};function pe(f,i){const m=P,d=M,p=D;return b(),I("div",me,[e(m,{status:"success",title:f.$t("model.agent.success.title"),subtitle:f.$t("model.agent.success.update.subTitle")},null,8,["title","subtitle"]),e(p,{size:16},{default:a(()=>[e(d,{key:"finish",type:"secondary",onClick:i[0]||(i[0]=o=>f.$router.push({name:"ModelAgentList"}))},{default:a(()=>[c(_(f.$t("model.agent.button.return")),1)]),_:1}),e(d,{key:"again",type:"primary",onClick:i[1]||(i[1]=o=>f.$router.push({name:"ModelAgentDetail",query:{id:`${f.$route.query.id}`}}))},{default:a(()=>[c(_(f.$t("model.agent.button.view")),1)]),_:1})]),_:1})])}const ie=S(ue,[["render",pe],["__scopeId","data-v-6fea42ce"]]),ce={class:"container"},_e={class:"wrapper"},fe={name:"ModelAgentUpdate"},be=C({...fe,setup(f){const{loading:i,setLoading:m}=U(!1),d=v(1),p=v({}),o=async()=>{m(!0);try{await le(p.value),d.value=3,p.value={}}catch{}finally{m(!1)}},y=(t,n)=>{if(typeof t=="number"){d.value=t;return}if(t==="forward"||t==="submit"){if(p.value={...p.value,...n},t==="submit"){o();return}d.value+=1}else t==="backward"&&(d.value-=1)};return(t,n)=>{const l=z,g=Y,s=Z,r=x,w=ee,V=ae,h=te;return b(),I("div",ce,[e(s,{class:"container-breadcrumb"},{default:a(()=>[e(g,null,{default:a(()=>[e(l)]),_:1}),e(g,null,{default:a(()=>[c(_(t.$t("menu.agent")),1)]),_:1}),e(g,null,{default:a(()=>[c(_(t.$t("menu.model.agent.update")),1)]),_:1})]),_:1}),e(h,{loading:X(i),style:{width:"100%"}},{default:a(()=>[e(V,{class:"general-card",bordered:!1},{title:a(()=>[c(_(t.$t("model.agent.title.update")),1)]),default:a(()=>[j("div",_e,[e(w,{current:d.value,"onUpdate:current":n[0]||(n[0]=u=>d.value=u),style:{width:"660px"},"line-less":"",class:"steps"},{default:a(()=>[e(r,{description:t.$t("model.agent.subTitle.baseInfo")},{default:a(()=>[c(_(t.$t("model.agent.title.baseInfo")),1)]),_:1},8,["description"]),e(r,{description:t.$t("model.agent.subTitle.advanced")},{default:a(()=>[c(_(t.$t("model.agent.title.advanced")),1)]),_:1},8,["description"]),e(r,{description:t.$t("model.agent.subTitle.update.finish")},{default:a(()=>[c(_(t.$t("model.agent.title.update.finish")),1)]),_:1},8,["description"])]),_:1},8,["current"]),(b(),k(Q,null,[d.value===1?(b(),k(se,{key:0,onChangeStep:y})):d.value===2?(b(),k(de,{key:1,onChangeStep:y})):d.value===3?(b(),k(ie,{key:2,onChangeStep:y})):W("",!0)],1024))])]),_:1})]),_:1},8,["loading"])])}}});const Ne=S(be,[["__scopeId","data-v-46b68998"]]);export{Ne as default};
