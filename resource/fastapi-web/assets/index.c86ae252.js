import{_ as U,G as L}from"./index.3029f51b.js";import{u as w}from"./loading.fc5bfc3b.js";/* empty css               *//* empty css               */import{d as S,e as k,B as v,aD as h,aG as o,aH as e,aL as n,aM as c,bN as A,bO as O,b4 as I,b3 as R,bU as z,aW as C,b7 as B,aT as G,c2 as H,bn as D,aU as K,C as M,bt as P,F as W,bH as X,aE as j,u as J,aK as Q,aF as Y,bW as Z,bX as x,bV as ee,bs as le}from"./arco.fd20202f.js";import{c as E,d as ae}from"./model.943397e8.js";/* empty css              *//* empty css               *//* empty css               *//* empty css              */import{h as N}from"./vue.70a4bb93.js";/* empty css               *//* empty css               */import"./chart.57980958.js";const oe=S({__name:"base-info",emits:["changeStep"],setup(f,{emit:_}){const{setLoading:b}=w(!0),s=N(),p=k(),l=k({id:"",corp:"",name:"",model:"",type:"",remark:""});(async(r={id:s.query.id})=>{b(!0);try{const{data:a}=await E(r);l.value.id=a.id,l.value.corp=a.corp,l.value.name=a.name,l.value.model=a.model,l.value.type=String(a.type),l.value.remark=a.remark}catch{}finally{b(!1)}})();const u=async()=>{var a;await((a=p.value)==null?void 0:a.validate())||_("changeStep","forward",{...l.value})};return(r,a)=>{const t=A,y=O,d=I,$=R,g=z,V=C,F=B;return v(),h(F,{ref_key:"formRef",ref:p,model:l.value,class:"form","label-col-props":{span:6},"wrapper-col-props":{span:18}},{default:o(()=>[e(d,{field:"corp",label:r.$t("model.label.corp"),rules:[{required:!0,message:r.$t("model.error.corp.required")}]},{default:o(()=>[e(y,{modelValue:l.value.corp,"onUpdate:modelValue":a[0]||(a[0]=i=>l.value.corp=i),placeholder:r.$t("model.placeholder.corp")},{default:o(()=>[e(t,{value:"OpenAI"},{default:o(()=>[n("OpenAI")]),_:1})]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),e(d,{field:"name",label:r.$t("model.label.name"),rules:[{required:!0,message:r.$t("model.error.name.required")}]},{default:o(()=>[e($,{modelValue:l.value.name,"onUpdate:modelValue":a[1]||(a[1]=i=>l.value.name=i),placeholder:r.$t("model.placeholder.name")},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),e(d,{field:"model",label:r.$t("model.label.model"),rules:[{required:!0,message:r.$t("model.error.model.required")},{message:r.$t("model.error.model.pattern")}]},{default:o(()=>[e($,{modelValue:l.value.model,"onUpdate:modelValue":a[2]||(a[2]=i=>l.value.model=i),placeholder:r.$t("model.placeholder.model")},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),e(d,{field:"type",label:r.$t("model.label.type"),rules:[{required:!0,message:r.$t("model.error.type.required")}]},{default:o(()=>[e(y,{modelValue:l.value.type,"onUpdate:modelValue":a[3]||(a[3]=i=>l.value.type=i),placeholder:r.$t("model.placeholder.type")},{default:o(()=>[e(t,{value:"1"},{default:o(()=>[n("\u6587\u751F\u6587")]),_:1}),e(t,{value:"2"},{default:o(()=>[n("\u6587\u751F\u56FE")]),_:1}),e(t,{value:"3"},{default:o(()=>[n("\u56FE\u751F\u6587")]),_:1}),e(t,{value:"4"},{default:o(()=>[n("\u56FE\u751F\u56FE")]),_:1})]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),e(d,{field:"remark",label:r.$t("model.label.remark"),rules:[{required:!1}]},{default:o(()=>[e(g,{modelValue:l.value.remark,"onUpdate:modelValue":a[4]||(a[4]=i=>l.value.remark=i),placeholder:r.$t("model.placeholder.remark")},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),e(d,null,{default:o(()=>[e(V,{type:"primary",onClick:u},{default:o(()=>[n(c(r.$t("model.button.next")),1)]),_:1})]),_:1})]),_:1},8,["model"])}}});const te=U(oe,[["__scopeId","data-v-8f99e743"]]),re=S({__name:"advanced",emits:["changeStep"],setup(f,{emit:_}){const{setLoading:b}=w(!0),s=N(),p=k(),l=k({prompt_ratio:1,completion_ratio:1,data_format:"",base_url:"",path:"",proxy:"",is_public:!0});(async(a={id:s.query.id})=>{b(!0);try{const{data:t}=await E(a);l.value.prompt_ratio=t.prompt_ratio,l.value.completion_ratio=t.completion_ratio,l.value.data_format=String(t.data_format),l.value.base_url=t.base_url,l.value.path=t.path,l.value.proxy=t.proxy,l.value.is_public=t.is_public}catch{}finally{b(!1)}})();const u=async()=>{var t;await((t=p.value)==null?void 0:t.validate())||_("changeStep","submit",{...l.value})},r=()=>{_("changeStep","backward")};return(a,t)=>{const y=G,d=I,$=H,g=D,V=R,F=K,i=C,T=B;return v(),h(T,{ref_key:"formRef",ref:p,model:l.value,class:"form","label-col-props":{span:6},"wrapper-col-props":{span:18}},{default:o(()=>[e(d,{field:"prompt_ratio",label:a.$t("model.label.promptRatio"),rules:[{required:!0,message:a.$t("model.error.promptRatio.required")}]},{default:o(()=>[e(y,{modelValue:l.value.prompt_ratio,"onUpdate:modelValue":t[0]||(t[0]=m=>l.value.prompt_ratio=m),min:1,placeholder:a.$t("model.placeholder.promptRatio")},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),e(d,{field:"completion_ratio",label:a.$t("model.label.completionRatio"),rules:[{required:!0,message:a.$t("model.error.completionRatio.required")}]},{default:o(()=>[e(y,{modelValue:l.value.completion_ratio,"onUpdate:modelValue":t[1]||(t[1]=m=>l.value.completion_ratio=m),min:1,placeholder:a.$t("model.placeholder.completionRatio")},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),e(d,{field:"data_format",label:a.$t("model.label.dataFormat"),rules:[{required:!0,message:a.$t("model.error.dataFormat.required")}]},{default:o(()=>[e(g,{size:"large"},{default:o(()=>[e($,{modelValue:l.value.data_format,"onUpdate:modelValue":t[2]||(t[2]=m=>l.value.data_format=m),value:"1"},{default:o(()=>[n("\u7EDF\u4E00\u683C\u5F0F")]),_:1},8,["modelValue"]),e($,{modelValue:l.value.data_format,"onUpdate:modelValue":t[3]||(t[3]=m=>l.value.data_format=m),value:"2"},{default:o(()=>[n("\u5B98\u65B9\u683C\u5F0F")]),_:1},8,["modelValue"])]),_:1})]),_:1},8,["label","rules"]),e(d,{field:"base_url",label:a.$t("model.label.baseUrl"),rules:[{required:!1}]},{default:o(()=>[e(V,{modelValue:l.value.base_url,"onUpdate:modelValue":t[4]||(t[4]=m=>l.value.base_url=m),placeholder:a.$t("model.placeholder.baseUrl")},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),e(d,{field:"path",label:a.$t("model.label.path"),rules:[{required:!1}]},{default:o(()=>[e(V,{modelValue:l.value.path,"onUpdate:modelValue":t[5]||(t[5]=m=>l.value.path=m),placeholder:a.$t("model.placeholder.path")},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),e(d,{field:"proxy",label:a.$t("model.label.proxy"),rules:[{required:!1}]},{default:o(()=>[e(V,{modelValue:l.value.proxy,"onUpdate:modelValue":t[6]||(t[6]=m=>l.value.proxy=m),placeholder:a.$t("model.placeholder.proxy")},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),e(d,{field:"is_public",label:a.$t("model.label.isPublic"),rules:[{required:!0}]},{default:o(()=>[e(F,{modelValue:l.value.is_public,"onUpdate:modelValue":t[7]||(t[7]=m=>l.value.is_public=m)},null,8,["modelValue"])]),_:1},8,["label"]),e(d,null,{default:o(()=>[e(g,null,{default:o(()=>[e(i,{type:"secondary",onClick:r},{default:o(()=>[n(c(a.$t("model.button.prev")),1)]),_:1}),e(i,{type:"primary",onClick:u},{default:o(()=>[n(c(a.$t("model.button.next")),1)]),_:1})]),_:1})]),_:1})]),_:1},8,["model"])}}});const ue=U(re,[["__scopeId","data-v-20229978"]]);const de={},ne={class:"success-wrap"};function se(f,_){const b=P,s=C,p=D;return v(),M("div",ne,[e(b,{status:"success",title:f.$t("model.success.title"),subtitle:f.$t("model.success.update.subTitle")},null,8,["title","subtitle"]),e(p,{size:16},{default:o(()=>[e(s,{key:"finish",type:"secondary",onClick:_[0]||(_[0]=l=>f.$router.push({name:"ModelList"}))},{default:o(()=>[n(c(f.$t("model.button.return")),1)]),_:1}),e(s,{key:"again",type:"primary",onClick:_[1]||(_[1]=l=>f.$router.push({name:"ModelDetail",query:{id:`${f.$route.query.id}`}}))},{default:o(()=>[n(c(f.$t("model.button.view")),1)]),_:1})]),_:1})])}const me=U(de,[["render",se],["__scopeId","data-v-829f16db"]]),pe={class:"container"},ie={class:"wrapper"},_e={name:"ModelUpdate"},ce=S({..._e,setup(f){const{loading:_,setLoading:b}=w(!1),s=k(1),p=k({}),l=async()=>{b(!0);try{await ae(p.value),s.value=3,p.value={}}catch{}finally{b(!1)}},q=(u,r)=>{if(typeof u=="number"){s.value=u;return}if(u==="forward"||u==="submit"){if(p.value={...p.value,...r},u==="submit"){l();return}s.value+=1}else u==="backward"&&(s.value-=1)};return(u,r)=>{const a=L,t=Q,y=Y,d=Z,$=x,g=ee,V=le;return v(),M("div",pe,[e(y,{class:"container-breadcrumb"},{default:o(()=>[e(t,null,{default:o(()=>[e(a)]),_:1}),e(t,null,{default:o(()=>[n(c(u.$t("menu.model")),1)]),_:1}),e(t,null,{default:o(()=>[n(c(u.$t("menu.model.update")),1)]),_:1})]),_:1}),e(V,{loading:J(_),style:{width:"100%"}},{default:o(()=>[e(g,{class:"general-card"},{title:o(()=>[n(c(u.$t("model.title.update")),1)]),default:o(()=>[W("div",ie,[e($,{current:s.value,"onUpdate:current":r[0]||(r[0]=F=>s.value=F),style:{width:"580px"},"line-less":"",class:"steps"},{default:o(()=>[e(d,{description:u.$t("model.subTitle.baseInfo")},{default:o(()=>[n(c(u.$t("model.title.baseInfo")),1)]),_:1},8,["description"]),e(d,{description:u.$t("model.subTitle.advanced")},{default:o(()=>[n(c(u.$t("model.title.advanced")),1)]),_:1},8,["description"]),e(d,{description:u.$t("model.subTitle.update.finish")},{default:o(()=>[n(c(u.$t("model.title.update.finish")),1)]),_:1},8,["description"])]),_:1},8,["current"]),(v(),h(X,null,[s.value===1?(v(),h(te,{key:0,onChangeStep:q})):s.value===2?(v(),h(ue,{key:1,onChangeStep:q})):s.value===3?(v(),h(me,{key:2,onChangeStep:q})):j("",!0)],1024))])]),_:1})]),_:1},8,["loading"])])}}});const Ce=U(ce,[["__scopeId","data-v-eec543e9"]]);export{Ce as default};
