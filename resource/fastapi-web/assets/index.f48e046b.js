import{C as q,_ as D}from"./index.d6462cde.js";/* empty css                *//* empty css                *//* empty css               *//* empty css               *//* empty css                *//* empty css                */import{d as L,e as b,B as R,C as S,aH as o,aG as t,aL as u,aM as d,F as f,u as N,aK as T,aF as E,b1 as G,b2 as H,aS as K,aT as M,bO as O,aU as j,bi as z,b5 as A,bI as J,bR as P,g as Q}from"./arco.54c7388d.js";import{u as W}from"./loading.7321a6c2.js";import{f as X,g as Y}from"./corp.e8a3355a.js";import{f as Z,h as x}from"./vue.aa90ed69.js";import"./chart.f14251fc.js";import"./base.87fcf6e2.js";const ee={class:"container"},ae={class:"wrapper"},oe={class:"submit-btn"},le={name:"CorpUpdate"},te=L({...le,setup(re){const{proxy:v}=Q(),{loading:$,setLoading:n}=W(!1),y=Z(),g=x(),m=b(),l=b({id:"",name:"",code:"",sort:0,is_public:!0,remark:"",status:1});(async(a={id:g.query.id})=>{n(!0);try{const{data:e}=await X(a);l.value.id=e.id,l.value.name=e.name,l.value.code=e.code,l.value.sort=e.sort,l.value.is_public=e.is_public,l.value.remark=e.remark,l.value.status=e.status}catch{}finally{n(!1)}})();const h=async()=>{var e;if(!await((e=m.value)==null?void 0:e.validate())){n(!0);try{await Y(l.value).then(()=>{v.$message.success("\u66F4\u65B0\u6210\u529F"),y.push({name:"CorpList"})})}catch{}finally{n(!1)}}};return(a,e)=>{const c=q,p=T,V=E,i=G,s=H,C=K,k=M,w=O,_=j,B=z,I=A,U=J,F=P;return R(),S("div",ee,[o(V,{class:"container-breadcrumb"},{default:t(()=>[o(p,null,{default:t(()=>[o(c)]),_:1}),o(p,null,{default:t(()=>[u(d(a.$t("menu.corp")),1)]),_:1}),o(p,null,{default:t(()=>[u(d(a.$t("menu.corp.update")),1)]),_:1})]),_:1}),o(F,{loading:N($),style:{width:"100%"}},{default:t(()=>[o(U,{class:"general-card","body-style":{padding:"0 20px 20px 20px"},bordered:!1},{default:t(()=>[f("div",ae,[o(I,{ref_key:"formRef",ref:m,model:l.value,class:"form","label-col-props":{span:5},"wrapper-col-props":{span:18}},{default:t(()=>[o(s,{field:"name",label:a.$t("corp.label.name"),rules:[{required:!0,message:a.$t("corp.error.name.required")},{match:/^.{1,100}$/,message:a.$t("corp.error.name.pattern")}]},{default:t(()=>[o(i,{modelValue:l.value.name,"onUpdate:modelValue":e[0]||(e[0]=r=>l.value.name=r),placeholder:a.$t("corp.placeholder.name"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),o(s,{field:"code",label:a.$t("corp.label.code"),rules:[{required:!0,message:a.$t("corp.error.code.required")},{match:/^.{1,100}$/,message:a.$t("corp.error.code.pattern")}]},{default:t(()=>[o(i,{modelValue:l.value.code,"onUpdate:modelValue":e[1]||(e[1]=r=>l.value.code=r),placeholder:a.$t("corp.placeholder.code"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),o(s,{field:"sort",label:a.$t("corp.label.sort")},{default:t(()=>[o(C,{modelValue:l.value.sort,"onUpdate:modelValue":e[2]||(e[2]=r=>l.value.sort=r),placeholder:a.$t("corp.placeholder.sort"),precision:0,min:-10,max:999},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),o(s,{field:"is_public",label:a.$t("corp.label.is_public")},{default:t(()=>[o(k,{modelValue:l.value.is_public,"onUpdate:modelValue":e[3]||(e[3]=r=>l.value.is_public=r)},null,8,["modelValue"])]),_:1},8,["label"]),o(s,{field:"remark",label:a.$t("corp.label.remark")},{default:t(()=>[o(w,{modelValue:l.value.remark,"onUpdate:modelValue":e[4]||(e[4]=r=>l.value.remark=r),placeholder:a.$t("corp.placeholder.remark")},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),o(B,null,{default:t(()=>[f("div",oe,[o(_,{type:"secondary",onClick:e[5]||(e[5]=r=>a.$router.push({name:"CorpList"}))},{default:t(()=>[u(d(a.$t("button.cancel")),1)]),_:1}),o(_,{type:"primary",onClick:h},{default:t(()=>[u(d(a.$t("button.submit")),1)]),_:1})])]),_:1})]),_:1},8,["model"])])]),_:1})]),_:1},8,["loading"])])}}});const ge=D(te,[["__scopeId","data-v-a52a78dc"]]);export{ge as default};
